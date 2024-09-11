/*
Copyright The Clusternet Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// GlobalizationSpecApplyConfiguration represents an declarative configuration of the GlobalizationSpec type for use
// with apply.
type GlobalizationSpecApplyConfiguration struct {
	OverridePolicy          *v1alpha1.OverridePolicy            `json:"overridePolicy,omitempty"`
	ClusterAffinity         *v1.LabelSelectorApplyConfiguration `json:"clusterAffinity,omitempty"`
	Overrides               []OverrideConfigApplyConfiguration  `json:"overrides,omitempty"`
	Priority                *int32                              `json:"priority,omitempty"`
	*FeedApplyConfiguration `json:"feed,omitempty"`
}

// GlobalizationSpecApplyConfiguration constructs an declarative configuration of the GlobalizationSpec type for use with
// apply.
func GlobalizationSpec() *GlobalizationSpecApplyConfiguration {
	return &GlobalizationSpecApplyConfiguration{}
}

// WithOverridePolicy sets the OverridePolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OverridePolicy field is set to the value of the last call.
func (b *GlobalizationSpecApplyConfiguration) WithOverridePolicy(value v1alpha1.OverridePolicy) *GlobalizationSpecApplyConfiguration {
	b.OverridePolicy = &value
	return b
}

// WithClusterAffinity sets the ClusterAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterAffinity field is set to the value of the last call.
func (b *GlobalizationSpecApplyConfiguration) WithClusterAffinity(value *v1.LabelSelectorApplyConfiguration) *GlobalizationSpecApplyConfiguration {
	b.ClusterAffinity = value
	return b
}

// WithOverrides adds the given value to the Overrides field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Overrides field.
func (b *GlobalizationSpecApplyConfiguration) WithOverrides(values ...*OverrideConfigApplyConfiguration) *GlobalizationSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOverrides")
		}
		b.Overrides = append(b.Overrides, *values[i])
	}
	return b
}

// WithPriority sets the Priority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Priority field is set to the value of the last call.
func (b *GlobalizationSpecApplyConfiguration) WithPriority(value int32) *GlobalizationSpecApplyConfiguration {
	b.Priority = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *GlobalizationSpecApplyConfiguration) WithKind(value string) *GlobalizationSpecApplyConfiguration {
	b.ensureFeedApplyConfigurationExists()
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *GlobalizationSpecApplyConfiguration) WithAPIVersion(value string) *GlobalizationSpecApplyConfiguration {
	b.ensureFeedApplyConfigurationExists()
	b.APIVersion = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *GlobalizationSpecApplyConfiguration) WithNamespace(value string) *GlobalizationSpecApplyConfiguration {
	b.ensureFeedApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *GlobalizationSpecApplyConfiguration) WithName(value string) *GlobalizationSpecApplyConfiguration {
	b.ensureFeedApplyConfigurationExists()
	b.Name = &value
	return b
}

func (b *GlobalizationSpecApplyConfiguration) ensureFeedApplyConfigurationExists() {
	if b.FeedApplyConfiguration == nil {
		b.FeedApplyConfiguration = &FeedApplyConfiguration{}
	}
}
