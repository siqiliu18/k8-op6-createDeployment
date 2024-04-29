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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Op6ApiKindSpec defines the desired state of Op6ApiKind
type Op6ApiKindSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Name string `json:"name,omitempty"`
}

// Op6ApiKindStatus defines the observed state of Op6ApiKind
type Op6ApiKindStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Op6ApiKind is the Schema for the op6apikinds API
type Op6ApiKind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Op6ApiKindSpec   `json:"spec,omitempty"`
	Status Op6ApiKindStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Op6ApiKindList contains a list of Op6ApiKind
type Op6ApiKindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Op6ApiKind `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Op6ApiKind{}, &Op6ApiKindList{})
}
