package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/flightctl/flightctl/lib/apipublic/v1alpha1"
	"github.com/flightctl/flightctl/lib/cli"
)

var (
	server string
)

func ptr(s string) *string {
	return &s
}

// go run cmd/flightctl-lib-go-test/main.go --server="https://api.flightctl.apps.server-foundation-sno-lite-f5c4m.dev04.red-chesterfield.com"
func main() {
	flag.StringVar(&server, "server", "", "server")
	flag.Parse()

	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImppOUx3TXdKQ3RFQjEtOGMyVlNpd3hZaVVlbmR6T1ZlMWNmQno0NDdvRzQifQ.eyJhdWQiOlsiaHR0cHM6Ly9rdWJlcm5ldGVzLmRlZmF1bHQuc3ZjIl0sImV4cCI6MTczMjYyMzIyMywiaWF0IjoxNzMyNTM2ODIzLCJpc3MiOiJodHRwczovL2t1YmVybmV0ZXMuZGVmYXVsdC5zdmMiLCJrdWJlcm5ldGVzLmlvIjp7Im5hbWVzcGFjZSI6ImRlZmF1bHQiLCJzZXJ2aWNlYWNjb3VudCI6eyJuYW1lIjoiZmxpZ2h0Y3RsLXRlc3Qtc2EiLCJ1aWQiOiIwZGQzOWUzNS05YjJmLTQzMWMtYjJlMC1iNTQxNDk2NzRhM2MifX0sIm5iZiI6MTczMjUzNjgyMywic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6ZmxpZ2h0Y3RsLXRlc3Qtc2EifQ.TbC7QG9RS9oHQvO28Qmm1EywJGthxIQvymBIlrPHhSnV8-pGzITYFMH22AZi6RySf_3OYQ1XGgfsURCgTMvAPpvm6pg9fSLxFSOpNPI3ye1xbnpWil9f3bA0diBiAXHaDhLZk5KZbkLeFMxFEIyr9uBfzK62PzdVGtvlNYkrNdzeQVzypgda1SmELLdY8vdEV2pjerwVoi1I_PgeWbV4HmtC4lWATfdQ6_nMY6y3dvTdR8dpqjnHtIb04YOOebCBbtGHWR23tcI1uID5eMABY--599TDFLg4RbUQDvJtSHBSzyoo-KnoVEtlwbgJVN7PVXRR9DHrWLtgsDgL4bZJKca-WheT8OdOevaUo28Y8HXpyVJ6tceoyubpIoycmABY6p16arsKdF6y6WBgYHeQPIQB8NMQ1qIO91qJq0Bs1rff3SzQmV7DLVeQdpiMPTjCcJXN4-_1yqaI0ITOtQOspgemSS0oGkqYKJG2WjTitSPfX9oN-WAzSX5_jHab82YhHJLY-TND-yqgREjukToU-ANf4GzCEYk-ss5CADDMS-GN316gFwt-5SdrwgIukEwXgA_E-WP0CdsHSnHbRLfb1vb80HX91Begab9gIKRwLTMEVbVkKksNLtM-gn_qZmzsVTq5Jzc0v-x9aFabsUQ0Ovk2u6lWJQWeRoyp2O6igqY"

	var err error
	var ctx = context.Background()

	repoName := "test-repo5"

	// apply repository
	expectRepo := v1alpha1.Repository{
		ApiVersion: "v1alpha1",
		Kind:       "Repository",
		Metadata: v1alpha1.ObjectMeta{
			Name: ptr(repoName),
		},
		Spec: v1alpha1.RepositorySpec{},
	}
	expectRepo.Spec.MergeHttpRepoSpec(v1alpha1.HttpRepoSpec{
		Type: "http",
		Url:  "https://github.com/flightctl/flightctl",
	})
	err = cli.ApplyRepository(ctx, token, server, &expectRepo)
	if err != nil {
		panic(err)
	}

	// get repository
	resultRepo, err := cli.GetRepository(ctx, token, server, repoName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("repository:%v\n", resultRepo)

	// get device
	deviceName := "kfqn5m8j0r8eg2mk0mindl2vhmovgj1q91c9s84eaqpp5pgdvl00"
	deviceResponse, err := cli.GetDevice(ctx, token, server, deviceName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("device:%v\n", deviceResponse)
}
