package platform

import (
	"fmt"
	"log"

	"github.com/AlexsJones/kubebuilder/src/log"

	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

// operation represents a Kubernetes operation.
type operation interface {
	Do(c *kubernetes.Clientset)
}

type versionOperation struct{}

func (op *versionOperation) Do(c *kubernetes.Clientset) {
	info, err := c.Discovery().ServerVersion()
	if err != nil {
		logger.GetInstance().Fatal(fmt.Sprintf("failed to retrieve server API version: %s\n", err))
	}

	logger.GetInstance().Log(fmt.Sprintf("server API version information: %s\n", info))

}

type deployOperation struct {
	image string
	name  string
	port  int
}

func (op *deployOperation) Do(c *kubernetes.Clientset) {
	if err := op.doDeployment(c); err != nil {
		log.Fatal(err)
	}

	if err := op.doService(c); err != nil {
		log.Fatal(err)
	}
}

func (op *deployOperation) doDeployment(c *kubernetes.Clientset) error {
	appName := op.name

	// Define Deployments spec.
	deploySpec := &v1beta1.Deployment{
		TypeMeta: unversioned.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "extensions/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: appName,
		},
		Spec: v1beta1.DeploymentSpec{
			Replicas: int32p(1),
			Strategy: v1beta1.DeploymentStrategy{
				Type: v1beta1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &v1beta1.RollingUpdateDeployment{
					MaxUnavailable: &intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(0),
					},
					MaxSurge: &intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(1),
					},
				},
			},
			RevisionHistoryLimit: int32p(10),
			Template: v1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Name:   appName,
					Labels: map[string]string{"app": appName},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						v1.Container{
							Name:  op.name,
							Image: op.image,
							Ports: []v1.ContainerPort{
								v1.ContainerPort{ContainerPort: int32(op.port), Protocol: v1.ProtocolTCP},
							},
							Resources: v1.ResourceRequirements{
								Limits: v1.ResourceList{
									v1.ResourceCPU:    resource.MustParse("100m"),
									v1.ResourceMemory: resource.MustParse("256Mi"),
								},
							},
							ImagePullPolicy: v1.PullIfNotPresent,
						},
					},
					RestartPolicy: v1.RestartPolicyAlways,
					DNSPolicy:     v1.DNSClusterFirst,
				},
			},
		},
	}

	// Implement deployment update-or-create semantics.
	deploy := c.Extensions().Deployments(namespace)
	_, err := deploy.Update(deploySpec)
	switch {
	case err == nil:
		logger.Println("deployment controller updated")
	case !errors.IsNotFound(err):
		return fmt.Errorf("could not update deployment controller: %s", err)
	default:
		_, err = deploy.Create(deploySpec)
		if err != nil {
			return fmt.Errorf("could not create deployment controller: %s", err)
		}
		logger.Println("deployment controller created")
	}

	return nil
}

func (op *deployOperation) doService(c *kubernetes.Clientset) error {
	appName := op.name

	// Define service spec.
	serviceSpec := &v1.Service{
		TypeMeta: unversioned.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: appName,
		},
		Spec: v1.ServiceSpec{
			Type:     v1.ServiceTypeClusterIP,
			Selector: map[string]string{"app": appName},
			Ports: []v1.ServicePort{
				v1.ServicePort{
					Protocol: v1.ProtocolTCP,
					Port:     80,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(op.port),
					},
				},
			},
		},
	}

	// Implement service update-or-create semantics.
	service := c.Core().Services(namespace)
	svc, err := service.Get(appName)
	switch {
	case err == nil:
		serviceSpec.ObjectMeta.ResourceVersion = svc.ObjectMeta.ResourceVersion
		serviceSpec.Spec.ClusterIP = svc.Spec.ClusterIP
		_, err = service.Update(serviceSpec)
		if err != nil {
			return fmt.Errorf("failed to update service: %s", err)
		}
		logger.Println("service updated")
	case errors.IsNotFound(err):
		_, err = service.Create(serviceSpec)
		if err != nil {
			return fmt.Errorf("failed to create service: %s", err)
		}
		logger.Println("service created")
	default:
		return fmt.Errorf("unexpected error: %s", err)
	}

	return nil
}

func int32p(i int32) *int32 {
	return &i
}
