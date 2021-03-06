/**
 * Copyright (c) 2016 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package k8s

type K8sClusterCredentials struct {
	ClusterName    string `json:"cluster_name" envconfig:"name"`
	Server         string `json:"api_server" envconfig:"api_server"`
	Username       string `json:"username" envconfig:"username"`
	Password       string `json:"password" envconfig:"password"`
	CaCert         string `json:"ca_cert" envconfig:"ca_cert"`
	AdminKey       string `json:"admin_key" envconfig:"admin_key"`
	AdminCert      string `json:"admin_cert" envconfig:"admin_cert"`
	ConsulEndpoint string `json:"consul_http_api" envconfig:"consul_http_api"`
	ConsulDomain   string `json:"consul_domain" envconfig:"consul_domain" default:"consul"`
}

type K8sConnector interface {
	DeleteCluster(org string) error
	GetCluster(org string) (int, K8sClusterCredentials, error)
	GetOrCreateCluster(org string) (K8sClusterCredentials, error)
	PostCluster(org string) (int, error)
	GetClusters() ([]K8sClusterCredentials, error)
}
