package main

import (
	"flag"

	"k8s.io/client-go/util/workqueue"

	"k8s.io/client-go/tools/cache"

	log "github.com/Sirupsen/logrus"

	snowballresourceclientset "github.com/robel-yemane/snowball-controller/pkg/client/clientset/versioned"
	snowballresourceinformerv1 "github.com/robel-yemane/snowball-controller/pkg/client/informers/externalversions/snowballresource/v1"
	"github.com/robel-yemane/snowball-controller/pkg/signals"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	masterURL  string
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

// retrieve the Kubernetes cluster client from kubeconfig if running outisde the cluster
// or use tokens if running from within the cluster
func getKubernetesClient() (kubernetes.Interface, snowballresourceclientset.Interface) {
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	snowballClient, err := snowballresourceclientset.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building snowball clientset: %s", err.Error())
	}
	return kubeClient, snowballClient
}

func main() {
	flag.Parse()

	//get the kuberenetes client for connectivity
	k8sClient, snowbalClient := getKubernetesClient()

	// retrieve our custom resource informer which was generated from
	// the code generator and pass it to the custom resource client, specifying
	// we should be looking through all namespaces for listing and watching
	snowballInformer := snowballresourceinformerv1.NewSnowballResourceInformer(
		snowbalClient,
		metav1.NamespaceAll,
		0,
		cache.Indexers{},
	)

	// create a new queue so that when the informer gets a resource that is either
	// a result of listing or watrching, we can add an identifying key to the queue
	// so that it can be handled in the handler
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// add event handler to handle the three types of events for the resources:
	// - adding new resources
	// - updating existing resources
	// - deleting resources

	snowballInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			// convert the resource object into a key (in this case
			// we are just doing it in the format of 'namespace/name')
			key, err := cache.MetaNamespaceKeyFunc(obj)
			log.Infof("Add my resource: %s", key)
			if err == nil {
				// add teh key to the queue for the handler to get
				queue.Add(key)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			log.Infof("Update my resource: %s", key)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			// DeletionHandlingMetaNameSpaceKeyFunc is a helper function that allows
			// us to check the DeletedFinalStateUnknown existence in teh event that
			// a resource was deleted but it is still contained in the index
			//
			// this then in turn calls MetaNameSpaceKeyFun
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			log.Infof("Deleted myresource: %s", key)
			if err == nil {
				queue.Add(key)
			}
		},
	})

	// construct the Controller object which has all of the necessary components to
	// handle logging, connections, informing(listing and watching), the queue,
	// and the handler
	controller := Controller{
		logger:    log.NewEntry(log.New()),
		clientset: k8sClient,
		informer:  snowballInformer,
		queue:     queue,
		handler:   &TestHandler{},
	}

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()
	// run the controller loop to process items
	controller.Run(1, stopCh)
	// select {}

	// // use a channel to synchronize the finalization for a graceful shutdown
	// stopCh := make(chan struct{})
	// defer close(stopCh)

	// // run the controller loop to process items
	// go controller.Run(stopCh)

	// // use a channel to handle OS signals to terminate and gracefully shut
	// // down processing
	// sigTerm := make(chan os.Signal, 1)
	// signal.Notify(sigTerm, syscall.SIGTERM)
	// signal.Notify(sigTerm, syscall.SIGINT)
	// <-sigTerm
}
