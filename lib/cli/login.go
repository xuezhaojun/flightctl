package cli

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	apiClient "github.com/flightctl/flightctl/lib/api/client"
	"github.com/flightctl/flightctl/lib/apipublic/v1alpha1"
	"github.com/flightctl/flightctl/lib/reqid"
	"github.com/go-chi/chi/middleware"
)

func LoginWithOpenshiftToken(ctx context.Context, token string, server string) error {
	httpClient, err := NewHTTPClientFromConfig(Config{
		Server: server,
		Token:  token,
	})
	if err != nil {
		return fmt.Errorf("failed to create http client: %w", err)
	}

	c, err := apiClient.NewClientWithResponses(server, apiClient.WithHTTPClient(httpClient))
	if err != nil {
		return fmt.Errorf("creating client: %w", err)
	}

	resp, err := c.AuthConfigWithResponse(ctx)
	if err != nil {
		return fmt.Errorf("failed to get auth info: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("unexpected response code: %v", resp.StatusCode())
	}

	headerVal := "Bearer " + token
	res, err := c.AuthValidateWithResponse(ctx, &v1alpha1.AuthValidateParams{Authentication: &headerVal})
	if err != nil {
		return fmt.Errorf("validating token: %w", err)
	}

	if res.StatusCode() != http.StatusOK {
		return fmt.Errorf("unexpected response code: %v", res.StatusCode())
	}

	return nil
}

func apiClientFromToken(token string, server string) (*apiClient.ClientWithResponses, error) {
	// Create client configuration
	config := Config{
		Server: server,
		Token:  token,
	}

	// Create new client
	c, err := NewFromConfig(config)
	if err != nil {
		return nil, fmt.Errorf("creating client: %w", err)
	}

	return c, nil
}

type Config struct {
	Server string `json:"server"`
	Token  string `json:"token"`
}

// NewHTTPClientFromConfig returns a new HTTP Client from the given config.
func NewHTTPClientFromConfig(config Config) (*http.Client, error) {
	u, err := url.Parse(config.Server)
	if err != nil {
		return nil, fmt.Errorf("NewHTTPClientFromConfig: parsing CA certs: %w", err)
	}

	tlsConfig := tls.Config{
		ServerName:         u.Hostname(),
		MinVersion:         tls.VersionTLS13,
		InsecureSkipVerify: true, //nolint:gosec
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tlsConfig,
		},
	}
	return httpClient, nil
}

// NewFromConfig returns a new FlightCtl API client from the given config.
func NewFromConfig(config Config) (*apiClient.ClientWithResponses, error) {
	httpClient, err := NewHTTPClientFromConfig(config)
	if err != nil {
		return nil, fmt.Errorf("NewFromConfig: creating HTTP client %w", err)
	}
	ref := apiClient.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set(middleware.RequestIDHeader, reqid.GetReqID())
		if config.Token != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Token))
		}
		return nil
	})
	return apiClient.NewClientWithResponses(config.Server, apiClient.WithHTTPClient(httpClient), ref)
}
