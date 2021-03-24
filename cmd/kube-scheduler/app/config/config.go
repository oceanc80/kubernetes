/*
Copyright 2018 The Kubernetes Authors.

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

package config

import (
	apiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/events"
	"k8s.io/client-go/tools/leaderelection"
	kubeschedulerconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
)

// Config has all the context to run a Scheduler
type Config struct {
	// ComponentConfig is the scheduler server's configuration object.
	ComponentConfig kubeschedulerconfig.KubeSchedulerConfiguration

	// LoopbackClientConfig is a config for a privileged loopback connection
	LoopbackClientConfig *restclient.Config

	Authentication apiserver.AuthenticationInfo
	Authorization  apiserver.AuthorizationInfo
	SecureServing  *apiserver.SecureServingInfo

	Client             clientset.Interface
	KubeConfig         *restclient.Config
	InformerFactory    informers.SharedInformerFactory
	DynInformerFactory dynamicinformer.DynamicSharedInformerFactory

	//nolint:staticcheck // SA1019 this deprecated field still needs to be used for now. It will be removed once the migration is done.
	EventBroadcaster events.EventBroadcasterAdapter

	// LeaderElection is optional.
	LeaderElection *leaderelection.LeaderElectionConfig

	// OpenShiftContext is additional context that we need to launch the kube-scheduler for openshift
	OpenShiftContext OpenShiftContext
}

type completedConfig struct {
	*Config
}

// CompletedConfig same as Config, just to swap private object.
type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (c *Config) Complete() CompletedConfig {
	cc := completedConfig{c}

	apiserver.AuthorizeClientBearerToken(c.LoopbackClientConfig, &c.Authentication, &c.Authorization)

	return CompletedConfig{&cc}
}
