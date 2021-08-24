package providers

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const RKE2_WINDOWS = "rke2.windows"


func IsRKE2Windows(ctx context.Context, k8sClient kubernetes.Interface) (bool, error) {
	// Check if the cluster is rke2 to avoid false positive of rke1 windows clusters
	if isRKE2, err := IsRKE2(ctx, k8sClient); err != nil || !isRKE2 {
		return false, err
	}
	if isHarvester, err := IsHarvester(ctx, k8sClient); err != nil || isHarvester {
		return false, err
	}
	// Look for nodes that have a Windows specific label
	listOpts := metav1.ListOptions{
		Limit:         1,
		LabelSelector: "kubernetes.io/os=windows",
	}

	windowsNodes, err := k8sClient.CoreV1().Nodes().List(ctx, listOpts)
	if err != nil {
		return false, err
	}
	if len(windowsNodes.Items) == 0 {
		return false, nil
	}

	return len(windowsNodes.Items) > 0, nil
}
