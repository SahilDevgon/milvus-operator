// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	in.ComponentSpec.DeepCopyInto(&out.ComponentSpec)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDeployStatus) DeepCopyInto(out *ComponentDeployStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDeployStatus.
func (in *ComponentDeployStatus) DeepCopy() *ComponentDeployStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentDeployStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSpec) DeepCopyInto(out *ComponentSpec) {
	*out = *in
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]Values, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpec.
func (in *ComponentSpec) DeepCopy() *ComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InClusterConfig) DeepCopyInto(out *InClusterConfig) {
	*out = *in
	in.Values.DeepCopyInto(&out.Values)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InClusterConfig.
func (in *InClusterConfig) DeepCopy() *InClusterConfig {
	if in == nil {
		return nil
	}
	out := new(InClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Milvus) DeepCopyInto(out *Milvus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Milvus.
func (in *Milvus) DeepCopy() *Milvus {
	if in == nil {
		return nil
	}
	out := new(Milvus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Milvus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusComponents) DeepCopyInto(out *MilvusComponents) {
	*out = *in
	in.ComponentSpec.DeepCopyInto(&out.ComponentSpec)
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(MilvusProxy)
		(*in).DeepCopyInto(*out)
	}
	if in.MixCoord != nil {
		in, out := &in.MixCoord, &out.MixCoord
		*out = new(MilvusMixCoord)
		(*in).DeepCopyInto(*out)
	}
	if in.RootCoord != nil {
		in, out := &in.RootCoord, &out.RootCoord
		*out = new(MilvusRootCoord)
		(*in).DeepCopyInto(*out)
	}
	if in.IndexCoord != nil {
		in, out := &in.IndexCoord, &out.IndexCoord
		*out = new(MilvusIndexCoord)
		(*in).DeepCopyInto(*out)
	}
	if in.DataCoord != nil {
		in, out := &in.DataCoord, &out.DataCoord
		*out = new(MilvusDataCoord)
		(*in).DeepCopyInto(*out)
	}
	if in.QueryCoord != nil {
		in, out := &in.QueryCoord, &out.QueryCoord
		*out = new(MilvusQueryCoord)
		(*in).DeepCopyInto(*out)
	}
	if in.IndexNode != nil {
		in, out := &in.IndexNode, &out.IndexNode
		*out = new(MilvusIndexNode)
		(*in).DeepCopyInto(*out)
	}
	if in.DataNode != nil {
		in, out := &in.DataNode, &out.DataNode
		*out = new(MilvusDataNode)
		(*in).DeepCopyInto(*out)
	}
	if in.QueryNode != nil {
		in, out := &in.QueryNode, &out.QueryNode
		*out = new(MilvusQueryNode)
		(*in).DeepCopyInto(*out)
	}
	if in.Standalone != nil {
		in, out := &in.Standalone, &out.Standalone
		*out = new(MilvusStandalone)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusComponents.
func (in *MilvusComponents) DeepCopy() *MilvusComponents {
	if in == nil {
		return nil
	}
	out := new(MilvusComponents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusCondition) DeepCopyInto(out *MilvusCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusCondition.
func (in *MilvusCondition) DeepCopy() *MilvusCondition {
	if in == nil {
		return nil
	}
	out := new(MilvusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusDataCoord) DeepCopyInto(out *MilvusDataCoord) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusDataCoord.
func (in *MilvusDataCoord) DeepCopy() *MilvusDataCoord {
	if in == nil {
		return nil
	}
	out := new(MilvusDataCoord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusDataNode) DeepCopyInto(out *MilvusDataNode) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusDataNode.
func (in *MilvusDataNode) DeepCopy() *MilvusDataNode {
	if in == nil {
		return nil
	}
	out := new(MilvusDataNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusDependencies) DeepCopyInto(out *MilvusDependencies) {
	*out = *in
	in.Etcd.DeepCopyInto(&out.Etcd)
	in.Pulsar.DeepCopyInto(&out.Pulsar)
	in.Kafka.DeepCopyInto(&out.Kafka)
	in.RocksMQ.DeepCopyInto(&out.RocksMQ)
	in.Storage.DeepCopyInto(&out.Storage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusDependencies.
func (in *MilvusDependencies) DeepCopy() *MilvusDependencies {
	if in == nil {
		return nil
	}
	out := new(MilvusDependencies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusEtcd) DeepCopyInto(out *MilvusEtcd) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InCluster != nil {
		in, out := &in.InCluster, &out.InCluster
		*out = new(InClusterConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusEtcd.
func (in *MilvusEtcd) DeepCopy() *MilvusEtcd {
	if in == nil {
		return nil
	}
	out := new(MilvusEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusIndexCoord) DeepCopyInto(out *MilvusIndexCoord) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusIndexCoord.
func (in *MilvusIndexCoord) DeepCopy() *MilvusIndexCoord {
	if in == nil {
		return nil
	}
	out := new(MilvusIndexCoord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusIndexNode) DeepCopyInto(out *MilvusIndexNode) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusIndexNode.
func (in *MilvusIndexNode) DeepCopy() *MilvusIndexNode {
	if in == nil {
		return nil
	}
	out := new(MilvusIndexNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusIngress) DeepCopyInto(out *MilvusIngress) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IngressClassName != nil {
		in, out := &in.IngressClassName, &out.IngressClassName
		*out = new(string)
		**out = **in
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLSSecretRefs != nil {
		in, out := &in.TLSSecretRefs, &out.TLSSecretRefs
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusIngress.
func (in *MilvusIngress) DeepCopy() *MilvusIngress {
	if in == nil {
		return nil
	}
	out := new(MilvusIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusKafka) DeepCopyInto(out *MilvusKafka) {
	*out = *in
	if in.InCluster != nil {
		in, out := &in.InCluster, &out.InCluster
		*out = new(InClusterConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.BrokerList != nil {
		in, out := &in.BrokerList, &out.BrokerList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusKafka.
func (in *MilvusKafka) DeepCopy() *MilvusKafka {
	if in == nil {
		return nil
	}
	out := new(MilvusKafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusList) DeepCopyInto(out *MilvusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Milvus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusList.
func (in *MilvusList) DeepCopy() *MilvusList {
	if in == nil {
		return nil
	}
	out := new(MilvusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MilvusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusMixCoord) DeepCopyInto(out *MilvusMixCoord) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusMixCoord.
func (in *MilvusMixCoord) DeepCopy() *MilvusMixCoord {
	if in == nil {
		return nil
	}
	out := new(MilvusMixCoord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusProxy) DeepCopyInto(out *MilvusProxy) {
	*out = *in
	in.ServiceComponent.DeepCopyInto(&out.ServiceComponent)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusProxy.
func (in *MilvusProxy) DeepCopy() *MilvusProxy {
	if in == nil {
		return nil
	}
	out := new(MilvusProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusPulsar) DeepCopyInto(out *MilvusPulsar) {
	*out = *in
	if in.InCluster != nil {
		in, out := &in.InCluster, &out.InCluster
		*out = new(InClusterConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusPulsar.
func (in *MilvusPulsar) DeepCopy() *MilvusPulsar {
	if in == nil {
		return nil
	}
	out := new(MilvusPulsar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusQueryCoord) DeepCopyInto(out *MilvusQueryCoord) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusQueryCoord.
func (in *MilvusQueryCoord) DeepCopy() *MilvusQueryCoord {
	if in == nil {
		return nil
	}
	out := new(MilvusQueryCoord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusQueryNode) DeepCopyInto(out *MilvusQueryNode) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusQueryNode.
func (in *MilvusQueryNode) DeepCopy() *MilvusQueryNode {
	if in == nil {
		return nil
	}
	out := new(MilvusQueryNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusReplicas) DeepCopyInto(out *MilvusReplicas) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusReplicas.
func (in *MilvusReplicas) DeepCopy() *MilvusReplicas {
	if in == nil {
		return nil
	}
	out := new(MilvusReplicas)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusRocksMQ) DeepCopyInto(out *MilvusRocksMQ) {
	*out = *in
	in.Persistence.DeepCopyInto(&out.Persistence)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusRocksMQ.
func (in *MilvusRocksMQ) DeepCopy() *MilvusRocksMQ {
	if in == nil {
		return nil
	}
	out := new(MilvusRocksMQ)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusRootCoord) DeepCopyInto(out *MilvusRootCoord) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusRootCoord.
func (in *MilvusRootCoord) DeepCopy() *MilvusRootCoord {
	if in == nil {
		return nil
	}
	out := new(MilvusRootCoord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusSpec) DeepCopyInto(out *MilvusSpec) {
	*out = *in
	in.Com.DeepCopyInto(&out.Com)
	in.Dep.DeepCopyInto(&out.Dep)
	in.Conf.DeepCopyInto(&out.Conf)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusSpec.
func (in *MilvusSpec) DeepCopy() *MilvusSpec {
	if in == nil {
		return nil
	}
	out := new(MilvusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusStandalone) DeepCopyInto(out *MilvusStandalone) {
	*out = *in
	in.ServiceComponent.DeepCopyInto(&out.ServiceComponent)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusStandalone.
func (in *MilvusStandalone) DeepCopy() *MilvusStandalone {
	if in == nil {
		return nil
	}
	out := new(MilvusStandalone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusStatus) DeepCopyInto(out *MilvusStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MilvusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.IngressStatus.DeepCopyInto(&out.IngressStatus)
	out.DeprecatedReplicas = in.DeprecatedReplicas
	if in.ComponentsDeployStatus != nil {
		in, out := &in.ComponentsDeployStatus, &out.ComponentsDeployStatus
		*out = make(map[string]ComponentDeployStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusStatus.
func (in *MilvusStatus) DeepCopy() *MilvusStatus {
	if in == nil {
		return nil
	}
	out := new(MilvusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusStorage) DeepCopyInto(out *MilvusStorage) {
	*out = *in
	if in.InCluster != nil {
		in, out := &in.InCluster, &out.InCluster
		*out = new(InClusterConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusStorage.
func (in *MilvusStorage) DeepCopy() *MilvusStorage {
	if in == nil {
		return nil
	}
	out := new(MilvusStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusUpgrade) DeepCopyInto(out *MilvusUpgrade) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusUpgrade.
func (in *MilvusUpgrade) DeepCopy() *MilvusUpgrade {
	if in == nil {
		return nil
	}
	out := new(MilvusUpgrade)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MilvusUpgrade) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusUpgradeList) DeepCopyInto(out *MilvusUpgradeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MilvusUpgrade, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusUpgradeList.
func (in *MilvusUpgradeList) DeepCopy() *MilvusUpgradeList {
	if in == nil {
		return nil
	}
	out := new(MilvusUpgradeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MilvusUpgradeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusUpgradeSpec) DeepCopyInto(out *MilvusUpgradeSpec) {
	*out = *in
	out.Milvus = in.Milvus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusUpgradeSpec.
func (in *MilvusUpgradeSpec) DeepCopy() *MilvusUpgradeSpec {
	if in == nil {
		return nil
	}
	out := new(MilvusUpgradeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MilvusUpgradeStatus) DeepCopyInto(out *MilvusUpgradeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReplicasBeforeUpgrade != nil {
		in, out := &in.ReplicasBeforeUpgrade, &out.ReplicasBeforeUpgrade
		*out = new(MilvusReplicas)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MilvusUpgradeStatus.
func (in *MilvusUpgradeStatus) DeepCopy() *MilvusUpgradeStatus {
	if in == nil {
		return nil
	}
	out := new(MilvusUpgradeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence) DeepCopyInto(out *Persistence) {
	*out = *in
	in.PersistentVolumeClaim.DeepCopyInto(&out.PersistentVolumeClaim)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence.
func (in *Persistence) DeepCopy() *Persistence {
	if in == nil {
		return nil
	}
	out := new(Persistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaim) DeepCopyInto(out *PersistentVolumeClaim) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaim.
func (in *PersistentVolumeClaim) DeepCopy() *PersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceComponent) DeepCopyInto(out *ServiceComponent) {
	*out = *in
	in.Component.DeepCopyInto(&out.Component)
	if in.ServiceLabels != nil {
		in, out := &in.ServiceLabels, &out.ServiceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceAnnotations != nil {
		in, out := &in.ServiceAnnotations, &out.ServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(MilvusIngress)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceComponent.
func (in *ServiceComponent) DeepCopy() *ServiceComponent {
	if in == nil {
		return nil
	}
	out := new(ServiceComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Values.
func (in *Values) DeepCopy() *Values {
	if in == nil {
		return nil
	}
	out := new(Values)
	in.DeepCopyInto(out)
	return out
}