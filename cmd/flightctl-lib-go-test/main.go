package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/flightctl/flightctl/api/v1alpha1"
	"github.com/flightctl/flightctl/lib/cli"
)

var (
	token  string
	server string
)

func ptr(s string) *string {
	return &s
}

func main() {
	flag.StringVar(&token, "token", "", "token")
	flag.StringVar(&server, "server", "", "server")
	flag.Parse()

	fmt.Println("token:", token)
	fmt.Println("server:", server)

	var err error
	var ctx = context.Background()

	// login
	err = cli.LoginWithOpenshiftToken(ctx, token, server)
	if err != nil {
		panic(err)
	}

	repoName := "test-repo3"

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
	fmt.Printf("repository:%v", resultRepo)
}
