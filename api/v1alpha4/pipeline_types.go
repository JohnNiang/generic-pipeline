/*
Copyright 2021.

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

package v1alpha4

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PipelineSpec defines the desired state of Pipeline
type PipelineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// EventProvider string `json:"provider,omitempty"`
	// +optional
	Event *Event `json:"event,omitempty"`

	Jobs []Job `json:"jobs"`
}

// Event refers to event provider and condition.
type Event struct {
	// EventProvider indicates what's structure of Event.
	// +optional
	Provider string `json:"provider,omitempty"`
	// Filter refers to configuration of event filter.
	// +optional
	Filter apiextensionsv1.JSON `json:"filter,omitempty"`
}

type Job struct {
	Name string `json:"name"`
	// +optional
	If     string `json:"if,omitempty"`
	RunsOn string `json:"runs-on"`
	Steps  []Step `json:"steps"`
}

type Step struct {
	Name string `json:"name"`
	// +optional
	If string `json:"if,omitempty"`
	// +optional
	Use string `json:"use,omitempty"`
	// +optional
	With map[string]apiextensionsv1.JSON `json:"with,omitempty"`
	// +optional
	Run string `json:"run,omitempty"`
}

// PipelineStatus defines the observed state of Pipeline
type PipelineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Pipeline is the Schema for the pipelines API
type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelineSpec   `json:"spec,omitempty"`
	Status PipelineStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PipelineList contains a list of Pipeline
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pipeline `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pipeline{}, &PipelineList{})
}
