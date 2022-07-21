package api

import "github.com/labstack/echo/v4"

type Agent interface {
	Get(context echo.Context) error
	GetK8sObjs(context echo.Context) error
	GetPodsByDaemonSet(context echo.Context) error
	GetPodsByDeployment(context echo.Context) error
	GetPodsByReplicaSet(context echo.Context) error
	GetPodsByStatefulSet(context echo.Context) error
}
