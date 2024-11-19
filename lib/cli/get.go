package cli

import (
	"context"
	"fmt"
	"net/http"

	"github.com/flightctl/flightctl/api/v1alpha1"
)

func GetRepository(ctx context.Context, token string, server string, name string) (*v1alpha1.Repository, error) {
	c, err := apiClientFromToken(token, server)
	if err != nil {
		return nil, fmt.Errorf("creating client: %w", err)
	}

	// Call API to get repository
	response, err := c.ReadRepositoryWithResponse(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("reading repository %s: %w", name, err)
	}

	// Check response status
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("reading repository %s: status code %d", name, response.StatusCode())
	}

	return response.JSON200, nil
}

func GetDevice(ctx context.Context, token string, server string, name string) (*v1alpha1.Device, error) {
	c, err := apiClientFromToken(token, server)
	if err != nil {
		return nil, fmt.Errorf("creating client: %w", err)
	}

	// Call API to get device
	response, err := c.ReadDeviceWithResponse(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("reading device %s: %w", name, err)
	}

	// Check response status
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("reading device %s: status code %d", name, response.StatusCode())
	}
	if response.JSON200 == nil {
		return nil, fmt.Errorf("reading device %s: status code %d", name, response.StatusCode())
	}

	return response.JSON200, nil
}
