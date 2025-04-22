/*
Copyright 2025.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	migrationv1alpha1 "github.com/vaaandark/kaleido/api/v1alpha1"
)

// MigrationJobReconciler reconciles a MigrationJob object
type MigrationJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=migration.kaleido.io,resources=migrationjobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=migration.kaleido.io,resources=migrationjobs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=migration.kaleido.io,resources=migrationjobs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MigrationJob object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.4/pkg/reconcile
func (r *MigrationJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = logf.FromContext(ctx)

	migrationJob := &migrationv1alpha1.MigrationJob{}
	if err := r.Get(ctx, req.NamespacedName, migrationJob); err != nil {
		logf.Log.Info("Failed to get MigrationJob", "error", err.Error())
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if migrationJob.Status.Phase == "" {
		migrationJob.Status.Phase = migrationv1alpha1.MigrationJobPhaseWaiting
		if err := r.Status().Update(ctx, migrationJob); err != nil {
			logf.Log.Info("Failed to update MigrationJob status", "error", err.Error())
			return ctrl.Result{}, err
		}
		logf.Log.Info("MigrationJob status updated",
			"job", migrationJob.Name,
			"phase", migrationv1alpha1.MigrationJobPhaseWaiting)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MigrationJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&migrationv1alpha1.MigrationJob{}).
		Named("migrationjob").
		Complete(r)
}
