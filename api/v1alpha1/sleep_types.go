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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SleepSpec defines the desired state of Sleep
type SleepSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Sleep. Edit sleep_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// SleepStatus defines the observed state of Sleep
type SleepStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Sleep is the Schema for the sleeps API
type Sleep struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SleepSpec   `json:"spec,omitempty"`
	Status SleepStatus `json:"status,omitempty"`
}

func (s Sleep) DeepCopyObject() runtime.Object {
	//TODO implement me
	panic("implement me")
}

//+kubebuilder:object:root=true

// SleepList contains a list of Sleep
type SleepList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Sleep `json:"items"`
}

func (s SleepList) DeepCopyObject() runtime.Object {
	//TODO implement me
	panic("implement me")
}

func init() {
	SchemeBuilder.Register(&Sleep{}, &SleepList{})
}
