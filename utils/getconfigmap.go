package utils

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"strings"
)

type ClusterInfo struct {
	Kubeconfig string
	Tokens     map[string]string
}

func GetClusterInfo() *ClusterInfo {
	// get kube-public cluster-info cm from k8s

	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	cm, err := clientset.CoreV1().ConfigMaps("kube-public").Get(context.TODO(), "cluster-info", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	clusterInfo := &ClusterInfo{}
	sigs := map[string]string{}
	for k, v := range cm.Data {
		if k == "kubeconfig" {
			clusterInfo.Kubeconfig = v
			continue
		}

		if strings.HasPrefix(k, "jws-kubeconfig-") {
			tokenId := strings.TrimPrefix(k, "jws-kubeconfig-")
			sigs[tokenId] = v
		}
	}
	clusterInfo.Tokens = sigs
	return clusterInfo
}
