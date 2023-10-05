/*
Copyright 2023.

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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	scalityv1alpha1 "github.com/example/pod-ui/api/v1alpha1"
)

// PodUIReconciler reconciles a PodUI object
type PodUIReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=scality.pod-ui.com,resources=poduis,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=scality.pod-ui.com,resources=poduis/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=scality.pod-ui.com,resources=poduis/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PodUI object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *PodUIReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// retrieve the CR

	podui, err := r.GetMethod(ctx, req)
	if err != nil {
		log.Log.Error(err, "unable to get CR")
		return ctrl.Result{}, err
	}
	// TODO(user): your logic here

	// update the status
	// if status.state is empty then update
	if podui.Status.State != "" {
		return ctrl.Result{}, nil
	}

	podui.Status.State = "Success"
	if err := r.Status().Update(ctx, &podui); err != nil {
		log.Log.Error(err, "unable to update PodUI status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PodUIReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&scalityv1alpha1.PodUI{}).
		Complete(r)
}

func (r *PodUIReconciler) GetMethod(ctx context.Context, req ctrl.Request) (scalityv1alpha1.PodUI, error) {
	_ = log.FromContext(ctx)

	// retrieve the CR
	var podui scalityv1alpha1.PodUI

	if err := r.Get(ctx, req.NamespacedName, &podui); err != nil {
		log.Log.Error(err, "unable to fetch PodUI")
		return podui, client.IgnoreNotFound(err)
	}

	return podui, nil
}
