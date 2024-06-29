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
)

// StepSpec defines a step to be executed
type StepSpec struct {
	// Command is the full command to be executed
	Command []string `json:"command,omitempty"`
	// Description defines the step definition
	Description string `json:"description,omitempty"`
	// Previous stdout - todo(knabben) must go to step status
	PreviousStdout string
	// Previous stderr - todo(knabben) must go to step status
	PreviousStderr string
}

// DemoSpec defines the desired state of Demo
type DemoSpec struct {
	// Steps are a list of executable steps for a Demo
	Steps []StepSpec `json:"steps,omitempty"`
}

// DemoStatus defines the observed state of Demo
type DemoStatus struct {
	// Runs are the number of usage of this resource by demonstrators
	Runs int `json:"runs,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Demo is the Schema for the demo API
type Demo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoSpec   `json:"spec,omitempty"`
	Status DemoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DemoList contains a list of Demo
type DemoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Demo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Demo{}, &DemoList{})
}
