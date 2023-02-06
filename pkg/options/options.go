/*
Copyright 2022 The Kubernetes Authors.

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

package options

// ProviderIDFormatType enum attribute to identify Power VS or VPC ProviderID format.
type ProviderIDFormatType string

const (
	// PowerVSProviderIDFormatV1 will set provider id to machine as ibmpowervs://<cluster_name>/<vm_hostname>
	// Deprecated: Use ProviderIDFormatV1.
	PowerVSProviderIDFormatV1 ProviderIDFormatType = "v1"

	// PowerVSProviderIDFormatV2 will set provider id to machine as ibmpowervs://<region>/<zone>/<service_instance_id>/<powervs_machine_id>
<<<<<<< Updated upstream
	// Deprecated: Use ProviderIDFormatV2.
	PowerVSProviderIDFormatV2 ProviderIDFormatType = "v2"

	// ProviderIDFormatV1 will set provider id to machine as follows
	// For VPC machines: ibmvpc://<cluster_name>/<vm_hostname>
	// For Power VS machines: ibmpowervs://<cluster_name>/<vm_hostname>
	ProviderIDFormatV1 ProviderIDFormatType = "v1"

	// ProviderIDFormatV2 will set provider id to machine as follows
	// For VPC machines: ibm://<account_id>///<cluster_id>/<vpc_machine_id>
	// For Power VS machines: ibmpowervs://<region>/<zone>/<service_instance_id>/<powervs_machine_id>
	ProviderIDFormatV2 ProviderIDFormatType = "v2"
=======
	PowerVSProviderIDFormatV2 PowerVSProviderIDFormatType = "v2"

	// PowerVSProviderIDFormatV3 will set provider id to machine as ibmpowervs://<zone>/<service_instance_id>/<powervs_machine_id>
	PowerVSProviderIDFormatV3 PowerVSProviderIDFormatType = "v3"
>>>>>>> Stashed changes
)

var (
	// PowerVSProviderIDFormat is used to identify the Provider ID format for Power VS Machine.
	// Deprecated: Instead use ProviderIDFormat.
	PowerVSProviderIDFormat string
	// ProviderIDFormat is used to identify the Provider ID format for Machine.
	ProviderIDFormat string
)
