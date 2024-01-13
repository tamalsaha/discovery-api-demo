package main

import (
	"fmt"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"
)

// /Users/tamal/go/src/github.com/tamalsaha/secret-projects/ocm/multicluster-controlplane.kubeconfig

func main() {
	ctrl.SetLogger(klogr.New())
	if err := useGeneratedClient(); err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Minute)
}

func useGeneratedClient() error {
	fmt.Println("Using Generated client")
	cfg := ctrl.GetConfigOrDie()
	cfg.QPS = 100
	cfg.Burst = 100

	kc, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return err
	}

	rest.SetDefaultWarningHandler(rest.NoWarnings{})

	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(kc.Discovery()))

	var pglist *core.NodeList
	pglist, err = kc.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, db := range pglist.Items {
		fmt.Println(client.ObjectKeyFromObject(&db))
	}
	return nil
}
