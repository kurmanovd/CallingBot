package service

import (
	"context"

	//"sky-backend/service/gcloud"
	"server/service/web"
)

// Manager is just a collection of all services we have in the project
type Manager struct {
	Call CallService
}

// NewManager creates new service manager
func NewManager(ctx context.Context) (*Manager, error) {
	return &Manager{
		Call: web.NewCallWebService(ctx),
	}, nil
}
