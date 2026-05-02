package config

import (
	"context"
	"fmt"

	"cloud.google.com/go/compute/metadata"

	"github.com/otakakot/research-cloud-functions-metadata/env"
)

type Project interface {
	GetID(context.Context) (string, error)
}

func NewProject(env env.Env) Project {
	if env.IsLocal() {
		return LocalProject{}
	}

	return CloudProject{}
}

type LocalProject struct{}

func (lp LocalProject) GetID(_ context.Context) (string, error) {
	return "local", nil
}

type CloudProject struct{}

func (cp CloudProject) GetID(_ context.Context) (string, error) {
	pid, err := metadata.ProjectID()
	if err != nil {
		return "", fmt.Errorf("failed to get project id: %w", err)
	}

	return pid, nil
}
