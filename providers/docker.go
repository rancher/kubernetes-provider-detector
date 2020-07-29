package providers

import (
	"context"

	"k8s.io/client-go/kubernetes"
)

func IsDocker(ctx context.Context, k8sClient kubernetes.Interface) (bool, error) {
	return true, nil
}
