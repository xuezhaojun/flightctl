package apiserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"slices"
	"time"

	api "github.com/flightctl/flightctl/api/v1alpha1"
	"github.com/flightctl/flightctl/internal/api/server"
	tlsmiddleware "github.com/flightctl/flightctl/internal/api_server/middleware"
	"github.com/flightctl/flightctl/internal/auth"
	"github.com/flightctl/flightctl/internal/config"
	"github.com/flightctl/flightctl/internal/crypto"
	"github.com/flightctl/flightctl/internal/instrumentation"
	"github.com/flightctl/flightctl/internal/service"
	"github.com/flightctl/flightctl/internal/store"
	"github.com/flightctl/flightctl/internal/tasks"
	"github.com/flightctl/flightctl/pkg/queues"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	oapimiddleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/sirupsen/logrus"
)

const (
	gracefulShutdownTimeout = 5 * time.Second
)

type Server struct {
	log      logrus.FieldLogger
	cfg      *config.Config
	store    store.Store
	ca       *crypto.CA
	listener net.Listener
	provider queues.Provider
	metrics  *instrumentation.ApiMetrics
}

// New returns a new instance of a flightctl server.
func New(
	log logrus.FieldLogger,
	cfg *config.Config,
	store store.Store,
	ca *crypto.CA,
	listener net.Listener,
	provider queues.Provider,
	metrics *instrumentation.ApiMetrics,
) *Server {
	return &Server{
		log:      log,
		cfg:      cfg,
		store:    store,
		ca:       ca,
		listener: listener,
		provider: provider,
		metrics:  metrics,
	}
}

func oapiErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, fmt.Sprintf("API Error: %s", message), statusCode)
}

func (s *Server) Run(ctx context.Context) error {
	s.log.Println("Initializing async jobs")
	publisher, err := tasks.TaskQueuePublisher(s.provider)
	if err != nil {
		return err
	}
	configStorage, err := tasks.NewConfigStorage(s.cfg.KV.Hostname, s.cfg.KV.Port)
	if err != nil {
		return err
	}
	callbackManager := tasks.NewCallbackManager(publisher, s.log)

	s.log.Println("Initializing API server")
	swagger, err := api.GetSwagger()
	if err != nil {
		return fmt.Errorf("failed loading swagger spec: %w", err)
	}
	// Skip server name validation
	swagger.Servers = nil

	oapiOpts := oapimiddleware.Options{
		ErrorHandler: oapiErrorHandler,
	}

	authMiddleware, err := auth.CreateAuthMiddleware(s.cfg, s.log)
	if err != nil {
		return err
	}

	router := chi.NewRouter()

	middlewares := [](func(http.Handler) http.Handler){
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		authMiddleware,
		oapimiddleware.OapiRequestValidatorWithOptions(swagger, &oapiOpts),
	}

	if s.metrics != nil {
		middlewares = slices.Insert(middlewares, 0, s.metrics.ApiServerMiddleware)
	}

	router.Use(middlewares...)

	h := service.NewServiceHandler(s.store, callbackManager, configStorage, s.ca, s.log, s.cfg.Service.BaseAgentGrpcUrl, s.cfg.Service.BaseAgentEndpointUrl, s.cfg.Service.BaseUIUrl)
	server.HandlerFromMux(server.NewStrictHandler(h, nil), router)

	srv := tlsmiddleware.NewHTTPServer(router, s.log, s.cfg.Service.Address)

	go func() {
		<-ctx.Done()
		s.log.Println("Shutdown signal received:", ctx.Err())
		ctxTimeout, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		_ = srv.Shutdown(ctxTimeout)
		configStorage.Close()
		s.provider.Stop()
		s.provider.Wait()
	}()

	s.log.Printf("Listening on %s...", s.listener.Addr().String())
	if err := srv.Serve(s.listener); err != nil && !errors.Is(err, net.ErrClosed) {
		return err
	}

	return nil
}
