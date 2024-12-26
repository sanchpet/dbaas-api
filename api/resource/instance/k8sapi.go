package instance

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8sAPI struct {
}

func NewK8sAPI() *K8sAPI {
	return &K8sAPI{}
}

func (k *K8sAPI) List() (Instances, error) {
	instance1 := &Instance{
		ID:   uuid.New(),
		Type: "pgsql",
		CPU:  1,
		RAM:  1024,
	}

	instance2 := &Instance{
		ID:   uuid.New(),
		Type: "pgsql",
		CPU:  2,
		RAM:  2048,
	}

	return Instances{instance1, instance2}, nil
}

func (k *K8sAPI) Create(instance *Instance) (*Instance, error) {
	return instance, nil
}

func (k *K8sAPI) Get(id uuid.UUID) (*Instance, error) {
	instance := &Instance{
		ID:   id,
		Type: "pgsql",
		CPU:  1,
		RAM:  1024,
	}

	return instance, nil
}

func (k *K8sAPI) Delete(id uuid.UUID) (bool, error) {
	return true, nil
}

func Test() (Instances, error) {
	config, err := rest.InClusterConfig()

	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod example-xxxxx not found in default namespace\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found example-xxxxx pod in default namespace\n")
		}

		time.Sleep(10 * time.Second)
	}
}
