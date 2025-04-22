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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MigrationJobPhase string

const (
	MigrationJobPhaseWaiting      MigrationJobPhase = "Waiting"
	MigrationJobPhaseCheckpointed MigrationJobPhase = "Checkpointed"
	MigrationJobPhaseTransferred  MigrationJobPhase = "Transferred"
	MigrationJobPhaseDone         MigrationJobPhase = "Done"
)

// MigrationJobSpec defines the desired state of MigrationJob.
type MigrationJobSpec struct {
	// The source node to migrate from.
	SourceNode string `json:"sourceNode"`

	// The target node to migrate to.
	TargetNode string `json:"targetNode"`

	// The name of pod to migrate.
	PodName string `json:"podName"`

	// The namespace of the pod to migrate.
	PodNamespace string `json:"podNamespace,omitempty"`
}

// MigrationJobStatus defines the observed state of MigrationJob.
type MigrationJobStatus struct {
	// The current phase of the migration job.
	Phase MigrationJobPhase `json:"phase"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MigrationJob is the Schema for the migrationjobs API.
type MigrationJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MigrationJobSpec   `json:"spec,omitempty"`
	Status MigrationJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MigrationJobList contains a list of MigrationJob.
type MigrationJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MigrationJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MigrationJob{}, &MigrationJobList{})
}
