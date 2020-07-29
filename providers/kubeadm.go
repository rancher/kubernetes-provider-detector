package providers

import (
	"context"

	"k8s.io/client-go/kubernetes"
)

func IsKubeadm(ctx context.Context, k8sClient kubernetes.Interface) (bool, error) {
	return true, nil
}
