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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServerClassSpec defines the desired state of ServerClass
type ServerClassSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ServerClass. Edit serverclass_types.go to remove/update
	ServerClassName string                  `json:"serverClassName"`
	GPUEnabled      bool                    `json:"gpuEnabled"`
	Replicas        int32                   `json:"replicas"`
	PodTemplate     *corev1.PodTemplateSpec `json:"podTemplateSpec,omitempty"`
}

// ServerClassStatus defines the observed state of ServerClass
type ServerClassStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Ready bool `json:"ready"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ServerClass is the Schema for the serverclasses API
type ServerClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServerClassSpec   `json:"spec,omitempty"`
	Status ServerClassStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerClassList contains a list of ServerClass
type ServerClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServerClass{}, &ServerClassList{})
}
