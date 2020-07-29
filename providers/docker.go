package providers

import (
	"context"
	"errors"

	"k8s.io/client-go/kubernetes"
)

func IsDocker(ctx context.Context, k8sClient kubernetes.Interface) (bool, error) {
	return false, errors.New("provider docker not implemented")
}
