/*
Copyright 2023 Red Hat, Inc.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	openshiftclusterv1 "github.com/openshift/cluster-api-provider-openshift/api/cluster/v1alpha1"
)

// OpenShiftBootstrapConfigReconciler reconciles a OpenShiftBootstrapConfig object.
type OpenShiftBootstrapConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// SetupWithManager sets up the controller with the Manager.
func (r *OpenShiftBootstrapConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := ctrl.NewControllerManagedBy(mgr).
		For(&openshiftclusterv1.OpenShiftBootstrapConfig{}).
		Complete(r); err != nil {
		return fmt.Errorf("failed to create controller: %w", err)
	}

	return nil
}

//+kubebuilder:rbac:groups=cluster.openshift.io,resources=openshiftbootstrapconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cluster.openshift.io,resources=openshiftbootstrapconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cluster.openshift.io,resources=openshiftbootstrapconfigs/finalizers,verbs=update

// Reconcile reconciles a OpenShiftBootstrapConfig object.
func (r *OpenShiftBootstrapConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}
