package instance

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
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
	config, err := rest.InClusterConfig()

	if err != nil {
		return nil, fmt.Errorf("failed to get in-cluster config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	ns := instance.User
	if ns == "" {
		ns = "default"
	}

	size := strconv.Itoa(instance.Storage) + "Mi"

	if err := CreateNamespaceIfNotExist(clientset, ns); err != nil {
		return nil, err
	}

	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create dynamic client: %w", err)
	}

	clusterRes := schema.GroupVersionResource{
		Group:    "postgresql.cnpg.io",
		Version:  "v1",
		Resource: "clusters",
	}

	cluster := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "postgresql.cnpg.io/v1",
			"kind":       "Cluster",
			"metadata": map[string]interface{}{
				"name":      instance.User + "-cluster",
				"namespace": instance.User,
			},
			"spec": map[string]interface{}{
				"instances": int64(1), // use int64 for numbers in unstructured
				"storage": map[string]interface{}{
					"size":         size,
					"storageClass": "local-path",
				},
				"managed": map[string]interface{}{
					"services": map[string]interface{}{
						"additional": []interface{}{
							map[string]interface{}{
								"selectorType": "rw",
								"serviceTemplate": map[string]interface{}{
									"metadata": map[string]interface{}{
										"name": "test-database",
										"labels": map[string]interface{}{
											"test-label": "true",
										},
										"annotations": map[string]interface{}{
											"test-annotation": "true",
										},
									},
									"spec": map[string]interface{}{
										"type": "NodePort",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	cluster.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "postgresql.cnpg.io",
		Version: "v1",
		Kind:    "Cluster",
	})

	result, err := dynClient.Resource(clusterRes).Namespace(instance.User).Create(context.TODO(), cluster, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create cluster resource: %w", err)
	}

	fmt.Printf("Created Cluster %s\n", result.GetName())

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

func CreateNamespaceIfNotExist(clientset *kubernetes.Clientset, namespace string) error {
	_, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err == nil {
		// Namespace exists
		return nil
	}
	// If error is not NotFound, return it
	if !errors.IsNotFound(err) {
		return fmt.Errorf("failed to get namespace %s: %w", namespace, err)
	}

	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create namespace %s: %w", namespace, err)
	}
	fmt.Printf("Created namespace %s\n", namespace)
	return nil
}
