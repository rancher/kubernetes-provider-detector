package detector

import (
	"context"

	"github.com/rancher/provider-detector/providers"
	"k8s.io/client-go/kubernetes"
)

var allProviders = make(map[string]IsProvider)

// IsProvider is the interface all providers need to implement
type IsProvider func(ctx context.Context, k8sClient kubernetes.Interface) (bool, error)

func init() {
	allProviders["aks"] = providers.IsAKS
	// allProviders["docker"] = providers.IsDocker
	allProviders["eks"] = providers.IsEKS
	allProviders["gke"] = providers.IsGKE
	allProviders["k3s"] = providers.IsK3s
	// allProviders["kubeadm"] = providers.IsKubeadm
	// allProviders["minikube"] = providers.IsMinikube
	allProviders["rke"] = providers.IsRKE
	allProviders["rke2"] = providers.IsRKE2
}

// DetectProvider accepts a k8s interface and checks all registered providers for a match
// Returning an emptry string indicates the provider is unknown
func DetectProvider(ctx context.Context, k8sClient kubernetes.Interface) (string, error) {
	for name, p := range allProviders {
		// Check the context before calling the provider
		if err := ctx.Err(); err != nil {
			return "", err
		}

		if ok, err := p(ctx, k8sClient); err != nil {
			return "", err
		} else if ok {
			return name, nil
		}
	}
	return "", nil
}
