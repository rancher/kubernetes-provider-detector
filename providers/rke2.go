package providers

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"

	"k8s.io/client-go/kubernetes"
)

const RKE2 = "rke2"

func IsRKE2(ctx context.Context, k8sClient kubernetes.Interface) (bool, error) {
	// if there are windows nodes then this should not be counted as rke2.linux
	windowsNodes, err := k8sClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{
		Limit:         1,
		LabelSelector: "kubernetes.io/os=windows",
	})
	if err != nil {
		return false, err
	}
	if len(windowsNodes.Items) != 0 {
		return false, nil
	}

	if isHarvester, err := IsHarvester(ctx, k8sClient); err != nil || isHarvester {
		return false, err
	}

	v, err := k8sClient.Discovery().ServerVersion()
	if err != nil {
		return false, err
	}
	if strings.Contains(v.GitVersion, "+rke2") {
		return true, nil
	}
	return false, nil
}
