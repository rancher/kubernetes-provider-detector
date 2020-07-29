package providers

import (
	"context"
	"errors"

	"k8s.io/client-go/kubernetes"
)

func IsKubeadm(ctx context.Context, k8sClient kubernetes.Interface) (bool, error) {
	return false, errors.New("provider kubeadm not implemented")
}
