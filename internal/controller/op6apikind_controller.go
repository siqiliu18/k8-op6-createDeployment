/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	op6apigroupv1 "op6repo/api/v1"

	appsv1 "k8s.io/api/apps/v1" // https://pkg.go.dev/k8s.io/api/apps/v1
	corev1 "k8s.io/api/core/v1"
)

// Op6ApiKindReconciler reconciles a Op6ApiKind object
type Op6ApiKindReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=op6apigroup.op6domain,resources=op6apikinds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=op6apigroup.op6domain,resources=op6apikinds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=op6apigroup.op6domain,resources=op6apikinds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Op6ApiKind object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *Op6ApiKindReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// TODO(user): your logic here
	myProxy := &op6apigroupv1.Op6ApiKind{}
	err := r.Get(ctx, req.NamespacedName, myProxy)
	if err != nil {
		// If it's another error, we return it
		log.Error(err, "Error while retrieving MyProxy instance")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	existingDeployments := &appsv1.Deployment{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: myProxy.Spec.Name, Namespace: req.Namespace}, existingDeployments)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		dep := r.deploymentForExample(myProxy, req.Namespace)
		log.Info("Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
		err = r.Create(ctx, dep)
		if err != nil {
			log.Error(err, "Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}
		// Deployment created successfully - return and requeue
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		log.Error(err, "Failed to get Deployment")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var size int32 = 2
	if *existingDeployments.Spec.Replicas != size {
		fmt.Println(">>>>>>>>>>>>>>>>> here >>>>>>>>>>>>>>>>>>>>>>>>>")
		existingDeployments.Spec.Replicas = &size
		err = r.Update(ctx, existingDeployments)
		if err != nil {
			log.Error(err, "Failed to update Deployment", "Deployment.Namespace", existingDeployments.Namespace, "Deployment.Name", existingDeployments.Name)
			return ctrl.Result{}, err
		}
		// Spec updated - return and requeue
		// return ctrl.Result{Requeue: true}, nil
	}

	return ctrl.Result{}, nil
}

func (r *Op6ApiKindReconciler) deploymentForExample(myproxy *op6apigroupv1.Op6ApiKind, ns string) *appsv1.Deployment {
	dep := &appsv1.Deployment{}

	var replicas int32 = 2

	labels := map[string]string{
		"app": myproxy.Spec.Name,
	}
	dep.Namespace = ns
	dep.Name = myproxy.Spec.Name
	dep.Labels = labels

	dep.Spec = appsv1.DeploymentSpec{
		Selector: &v1.LabelSelector{
			MatchLabels: labels,
		},
		Replicas: &replicas,
		Template: corev1.PodTemplateSpec{
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:            "nginx",
						Image:           "nginx",
						ImagePullPolicy: corev1.PullNever,
					},
				},
			},
		},
	}
	dep.Spec.Template.Labels = labels

	return dep
}

// SetupWithManager sets up the controller with the Manager.
func (r *Op6ApiKindReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&op6apigroupv1.Op6ApiKind{}).
		Complete(r)
}
