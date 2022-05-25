package api

import "github.com/labstack/echo/v4"

// ReplicaSet api operations
type ReplicaSet interface {
	Get(context echo.Context) error
	GetByOwnerReference(context echo.Context) error
}
