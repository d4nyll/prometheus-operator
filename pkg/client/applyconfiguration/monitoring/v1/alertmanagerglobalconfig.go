// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
)

// AlertmanagerGlobalConfigApplyConfiguration represents an declarative configuration of the AlertmanagerGlobalConfig type for use
// with apply.
type AlertmanagerGlobalConfigApplyConfiguration struct {
	ResolveTimeout *v1.Duration                  `json:"resolveTimeout,omitempty"`
	HTTPConfig     *HTTPConfigApplyConfiguration `json:"httpConfig,omitempty"`
	SlackAPIURL    *corev1.SecretKeySelector     `json:"slackApiUrl,omitempty"`
	OpsGenieAPIURL *corev1.SecretKeySelector     `json:"opsGenieApiUrl,omitempty"`
	OpsGenieAPIKey *corev1.SecretKeySelector     `json:"opsGenieApiKey,omitempty"`
	PagerdutyURL   *corev1.SecretKeySelector     `json:"pagerdutyUrl,omitempty"`
}

// AlertmanagerGlobalConfigApplyConfiguration constructs an declarative configuration of the AlertmanagerGlobalConfig type for use with
// apply.
func AlertmanagerGlobalConfig() *AlertmanagerGlobalConfigApplyConfiguration {
	return &AlertmanagerGlobalConfigApplyConfiguration{}
}

// WithResolveTimeout sets the ResolveTimeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResolveTimeout field is set to the value of the last call.
func (b *AlertmanagerGlobalConfigApplyConfiguration) WithResolveTimeout(value v1.Duration) *AlertmanagerGlobalConfigApplyConfiguration {
	b.ResolveTimeout = &value
	return b
}

// WithHTTPConfig sets the HTTPConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPConfig field is set to the value of the last call.
func (b *AlertmanagerGlobalConfigApplyConfiguration) WithHTTPConfig(value *HTTPConfigApplyConfiguration) *AlertmanagerGlobalConfigApplyConfiguration {
	b.HTTPConfig = value
	return b
}

// WithSlackAPIURL sets the SlackAPIURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SlackAPIURL field is set to the value of the last call.
func (b *AlertmanagerGlobalConfigApplyConfiguration) WithSlackAPIURL(value corev1.SecretKeySelector) *AlertmanagerGlobalConfigApplyConfiguration {
	b.SlackAPIURL = &value
	return b
}

// WithOpsGenieAPIURL sets the OpsGenieAPIURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OpsGenieAPIURL field is set to the value of the last call.
func (b *AlertmanagerGlobalConfigApplyConfiguration) WithOpsGenieAPIURL(value corev1.SecretKeySelector) *AlertmanagerGlobalConfigApplyConfiguration {
	b.OpsGenieAPIURL = &value
	return b
}

// WithOpsGenieAPIKey sets the OpsGenieAPIKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OpsGenieAPIKey field is set to the value of the last call.
func (b *AlertmanagerGlobalConfigApplyConfiguration) WithOpsGenieAPIKey(value corev1.SecretKeySelector) *AlertmanagerGlobalConfigApplyConfiguration {
	b.OpsGenieAPIKey = &value
	return b
}

// WithPagerdutyURL sets the PagerdutyURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PagerdutyURL field is set to the value of the last call.
func (b *AlertmanagerGlobalConfigApplyConfiguration) WithPagerdutyURL(value corev1.SecretKeySelector) *AlertmanagerGlobalConfigApplyConfiguration {
	b.PagerdutyURL = &value
	return b
}
