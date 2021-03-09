package service

import (
	"context"
)

// CallService is a service for call
type CallService interface {
	GetCall(context.Context) error
}
