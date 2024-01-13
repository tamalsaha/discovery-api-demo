package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
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
	mapping, err := mapper.RESTMapping(schema.GroupKind{
		Group: "certificates.k8s.io",
		Kind:  "CertificateSigningRequest",
	})
	if err != nil {
		return err
	}

	fmt.Println(mapping.GroupVersionKind)
	return nil
}
