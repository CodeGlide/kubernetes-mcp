package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Iok8sapicorev1ConfigMapKeySelector represents the Iok8sapicorev1ConfigMapKeySelector schema from the OpenAPI specification
type Iok8sapicorev1ConfigMapKeySelector struct {
	Key string `json:"key"` // The key to select.
	Name string `json:"name,omitempty"` // Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the ConfigMap or its key must be defined
}

// Iok8sapicorev1RBDVolumeSource represents the Iok8sapicorev1RBDVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1RBDVolumeSource struct {
	User string `json:"user,omitempty"` // user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	Image string `json:"image"` // image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Keyring string `json:"keyring,omitempty"` // keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Monitors []string `json:"monitors"` // monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Pool string `json:"pool,omitempty"` // pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Readonly bool `json:"readOnly,omitempty"` // readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
}

// Iok8sapiauthorizationv1NonResourceAttributes represents the Iok8sapiauthorizationv1NonResourceAttributes schema from the OpenAPI specification
type Iok8sapiauthorizationv1NonResourceAttributes struct {
	Path string `json:"path,omitempty"` // Path is the URL path of the request
	Verb string `json:"verb,omitempty"` // Verb is the standard HTTP verb
}

// Iok8sapicoordinationv1beta1LeaseCandidateList represents the Iok8sapicoordinationv1beta1LeaseCandidateList schema from the OpenAPI specification
type Iok8sapicoordinationv1beta1LeaseCandidateList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicoordinationv1beta1LeaseCandidate `json:"items"` // items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1ContainerUser represents the Iok8sapicorev1ContainerUser schema from the OpenAPI specification
type Iok8sapicorev1ContainerUser struct {
	Linux Iok8sapicorev1LinuxContainerUser `json:"linux,omitempty"` // LinuxContainerUser represents user identity information in Linux containers
}

// Iok8sapidiscoveryv1EndpointConditions represents the Iok8sapidiscoveryv1EndpointConditions schema from the OpenAPI specification
type Iok8sapidiscoveryv1EndpointConditions struct {
	Terminating bool `json:"terminating,omitempty"` // terminating indicates that this endpoint is terminating. A nil value should be interpreted as "false".
	Ready bool `json:"ready,omitempty"` // ready indicates that this endpoint is ready to receive traffic, according to whatever system is managing the endpoint. A nil value should be interpreted as "true". In general, an endpoint should be marked ready if it is serving and not terminating, though this can be overridden in some cases, such as when the associated Service has set the publishNotReadyAddresses flag.
	Serving bool `json:"serving,omitempty"` // serving indicates that this endpoint is able to receive traffic, according to whatever system is managing the endpoint. For endpoints backed by pods, the EndpointSlice controller will mark the endpoint as serving if the pod's Ready condition is True. A nil value should be interpreted as "true".
}

// Iok8sapinetworkingv1IngressPortStatus represents the Iok8sapinetworkingv1IngressPortStatus schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressPortStatus struct {
	ErrorField string `json:"error,omitempty"` // error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use CamelCase names - cloud provider specific error values must have names that comply with the format foo.example.com/CamelCase.
	Port int `json:"port"` // port is the port number of the ingress port.
	Protocol string `json:"protocol"` // protocol is the protocol of the ingress port. The supported values are: "TCP", "UDP", "SCTP"
}

// Iok8sapicorev1ResourceRequirements represents the Iok8sapicorev1ResourceRequirements schema from the OpenAPI specification
type Iok8sapicorev1ResourceRequirements struct {
	Claims []Iok8sapicorev1ResourceClaim `json:"claims,omitempty"` // Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.
	Limits map[string]interface{} `json:"limits,omitempty"` // Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests map[string]interface{} `json:"requests,omitempty"` // Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
}

// Iok8sapiresourcev1beta2ResourceSliceSpec represents the Iok8sapiresourcev1beta2ResourceSliceSpec schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ResourceSliceSpec struct {
	Nodename string `json:"nodeName,omitempty"` // NodeName identifies the node which provides the resources in this pool. A field selector can be used to list only ResourceSlice objects belonging to a certain node. This field can be used to limit access from nodes to ResourceSlices with the same node name. It also indicates to autoscalers that adding new nodes of the same type as some old node might also make new resources available. Exactly one of NodeName, NodeSelector, AllNodes, and PerDeviceNodeSelection must be set. This field is immutable.
	Nodeselector Iok8sapicorev1NodeSelector `json:"nodeSelector,omitempty"` // A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
	Perdevicenodeselection bool `json:"perDeviceNodeSelection,omitempty"` // PerDeviceNodeSelection defines whether the access from nodes to resources in the pool is set on the ResourceSlice level or on each device. If it is set to true, every device defined the ResourceSlice must specify this individually. Exactly one of NodeName, NodeSelector, AllNodes, and PerDeviceNodeSelection must be set.
	Pool Iok8sapiresourcev1beta2ResourcePool `json:"pool"` // ResourcePool describes the pool that ResourceSlices belong to.
	Sharedcounters []Iok8sapiresourcev1beta2CounterSet `json:"sharedCounters,omitempty"` // SharedCounters defines a list of counter sets, each of which has a name and a list of counters available. The names of the SharedCounters must be unique in the ResourceSlice. The maximum number of counters in all sets is 32.
	Allnodes bool `json:"allNodes,omitempty"` // AllNodes indicates that all nodes have access to the resources in the pool. Exactly one of NodeName, NodeSelector, AllNodes, and PerDeviceNodeSelection must be set.
	Devices []Iok8sapiresourcev1beta2Device `json:"devices,omitempty"` // Devices lists some or all of the devices in this pool. Must not have more than 128 entries.
	Driver string `json:"driver"` // Driver identifies the DRA driver providing the capacity information. A field selector can be used to list only ResourceSlice objects with a certain driver name. Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver. This field is immutable.
}

// Iok8sapicorev1ReplicationController represents the Iok8sapicorev1ReplicationController schema from the OpenAPI specification
type Iok8sapicorev1ReplicationController struct {
	Status Iok8sapicorev1ReplicationControllerStatus `json:"status,omitempty"` // ReplicationControllerStatus represents the current status of a replication controller.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1ReplicationControllerSpec `json:"spec,omitempty"` // ReplicationControllerSpec is the specification of a replication controller.
}

// Iok8sapiadmissionregistrationv1NamedRuleWithOperations represents the Iok8sapiadmissionregistrationv1NamedRuleWithOperations schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1NamedRuleWithOperations struct {
	Resourcenames []string `json:"resourceNames,omitempty"` // ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources. If wildcard is present, the validation rule will ensure resources do not overlap with each other. Depending on the enclosing object, subresources might not be allowed. Required.
	Scope string `json:"scope,omitempty"` // scope specifies the scope of this rule. Valid values are "Cluster", "Namespaced", and "*" "Cluster" means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. "Namespaced" means that only namespaced resources will match this rule. "*" means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is "*".
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.
	Apiversions []string `json:"apiVersions,omitempty"` // APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.
	Operations []string `json:"operations,omitempty"` // Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If '*' is present, the length of the slice must be one. Required.
}

// Iok8sapibatchv1PodFailurePolicyRule represents the Iok8sapibatchv1PodFailurePolicyRule schema from the OpenAPI specification
type Iok8sapibatchv1PodFailurePolicyRule struct {
	Action string `json:"action"` // Specifies the action taken on a pod failure when the requirements are satisfied. Possible values are: - FailJob: indicates that the pod's job is marked as Failed and all running pods are terminated. - FailIndex: indicates that the pod's index is marked as Failed and will not be restarted. - Ignore: indicates that the counter towards the .backoffLimit is not incremented and a replacement pod is created. - Count: indicates that the pod is handled in the default way - the counter towards the .backoffLimit is incremented. Additional values are considered to be added in the future. Clients should react to an unknown action by skipping the rule.
	Onexitcodes Iok8sapibatchv1PodFailurePolicyOnExitCodesRequirement `json:"onExitCodes,omitempty"` // PodFailurePolicyOnExitCodesRequirement describes the requirement for handling a failed pod based on its container exit codes. In particular, it lookups the .state.terminated.exitCode for each app container and init container status, represented by the .status.containerStatuses and .status.initContainerStatuses fields in the Pod status, respectively. Containers completed with success (exit code 0) are excluded from the requirement check.
	Onpodconditions []Iok8sapibatchv1PodFailurePolicyOnPodConditionsPattern `json:"onPodConditions,omitempty"` // Represents the requirement on the pod conditions. The requirement is represented as a list of pod condition patterns. The requirement is satisfied if at least one pattern matches an actual pod condition. At most 20 elements are allowed.
}

// Iok8sapicorev1ResourceQuota represents the Iok8sapicorev1ResourceQuota schema from the OpenAPI specification
type Iok8sapicorev1ResourceQuota struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1ResourceQuotaSpec `json:"spec,omitempty"` // ResourceQuotaSpec defines the desired hard limits to enforce for Quota.
	Status Iok8sapicorev1ResourceQuotaStatus `json:"status,omitempty"` // ResourceQuotaStatus defines the enforced hard limits and observed use.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapicertificatesv1CertificateSigningRequestList represents the Iok8sapicertificatesv1CertificateSigningRequestList schema from the OpenAPI specification
type Iok8sapicertificatesv1CertificateSigningRequestList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicertificatesv1CertificateSigningRequest `json:"items"` // items is a collection of CertificateSigningRequest objects
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiappsv1DaemonSetUpdateStrategy represents the Iok8sapiappsv1DaemonSetUpdateStrategy schema from the OpenAPI specification
type Iok8sapiappsv1DaemonSetUpdateStrategy struct {
	Rollingupdate Iok8sapiappsv1RollingUpdateDaemonSet `json:"rollingUpdate,omitempty"` // Spec to control the desired behavior of daemon set rolling update.
	TypeField string `json:"type,omitempty"` // Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
}

// Iok8sapiautoscalingv1HorizontalPodAutoscaler represents the Iok8sapiautoscalingv1HorizontalPodAutoscaler schema from the OpenAPI specification
type Iok8sapiautoscalingv1HorizontalPodAutoscaler struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiautoscalingv1HorizontalPodAutoscalerSpec `json:"spec,omitempty"` // specification of a horizontal pod autoscaler.
	Status Iok8sapiautoscalingv1HorizontalPodAutoscalerStatus `json:"status,omitempty"` // current status of a horizontal pod autoscaler
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1EndpointsList represents the Iok8sapicorev1EndpointsList schema from the OpenAPI specification
type Iok8sapicorev1EndpointsList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1Endpoints `json:"items"` // List of endpoints.
}

// Iok8sapiresourcev1beta1DeviceToleration represents the Iok8sapiresourcev1beta1DeviceToleration schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceToleration struct {
	Key string `json:"key,omitempty"` // Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys. Must be a label name.
	Operator string `json:"operator,omitempty"` // Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a ResourceClaim can tolerate all taints of a particular category.
	Tolerationseconds int64 `json:"tolerationSeconds,omitempty"` // TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system. If larger than zero, the time when the pod needs to be evicted is calculated as <time when taint was adedd> + <toleration seconds>.
	Value string `json:"value,omitempty"` // Value is the taint value the toleration matches to. If the operator is Exists, the value must be empty, otherwise just a regular string. Must be a label value.
	Effect string `json:"effect,omitempty"` // Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule and NoExecute.
}

// Iok8sapiautoscalingv2MetricStatus represents the Iok8sapiautoscalingv2MetricStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2MetricStatus struct {
	Pods Iok8sapiautoscalingv2PodsMetricStatus `json:"pods,omitempty"` // PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
	Resource Iok8sapiautoscalingv2ResourceMetricStatus `json:"resource,omitempty"` // ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
	TypeField string `json:"type"` // type is the type of metric source. It will be one of "ContainerResource", "External", "Object", "Pods" or "Resource", each corresponds to a matching field in the object.
	Containerresource Iok8sapiautoscalingv2ContainerResourceMetricStatus `json:"containerResource,omitempty"` // ContainerResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing a single container in each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
	External Iok8sapiautoscalingv2ExternalMetricStatus `json:"external,omitempty"` // ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
	Object Iok8sapiautoscalingv2ObjectMetricStatus `json:"object,omitempty"` // ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
}

// Iok8sapicorev1FlexVolumeSource represents the Iok8sapicorev1FlexVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1FlexVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	Options map[string]interface{} `json:"options,omitempty"` // options is Optional: this field holds extra command options if any.
	Readonly bool `json:"readOnly,omitempty"` // readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	Driver string `json:"driver"` // driver is the name of the driver to use for this volume.
}

// Iok8sapiadmissionregistrationv1alpha1ParamRef represents the Iok8sapiadmissionregistrationv1alpha1ParamRef schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1ParamRef struct {
	Name string `json:"name,omitempty"` // `name` is the name of the resource being referenced. `name` and `selector` are mutually exclusive properties. If one is set, the other must be unset.
	Namespace string `json:"namespace,omitempty"` // namespace is the namespace of the referenced resource. Allows limiting the search for params to a specific namespace. Applies to both `name` and `selector` fields. A per-namespace parameter may be used by specifying a namespace-scoped `paramKind` in the policy and leaving this field empty. - If `paramKind` is cluster-scoped, this field MUST be unset. Setting this field results in a configuration error. - If `paramKind` is namespace-scoped, the namespace of the object being evaluated for admission will be used when this field is left unset. Take care that if this is left empty the binding must not match any cluster-scoped resources, which will result in an error.
	Parameternotfoundaction string `json:"parameterNotFoundAction,omitempty"` // `parameterNotFoundAction` controls the behavior of the binding when the resource exists, and name or selector is valid, but there are no parameters matched by the binding. If the value is set to `Allow`, then no matched parameters will be treated as successful validation by the binding. If set to `Deny`, then no matched parameters will be subject to the `failurePolicy` of the policy. Allowed values are `Allow` or `Deny` Default to `Deny`
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapicorev1Binding represents the Iok8sapicorev1Binding schema from the OpenAPI specification
type Iok8sapicorev1Binding struct {
	Target Iok8sapicorev1ObjectReference `json:"target"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1ScaleIOPersistentVolumeSource represents the Iok8sapicorev1ScaleIOPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ScaleIOPersistentVolumeSource struct {
	System string `json:"system"` // system is the name of the storage system as configured in ScaleIO.
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs"
	Secretref Iok8sapicorev1SecretReference `json:"secretRef"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	Volumename string `json:"volumeName,omitempty"` // volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.
	Protectiondomain string `json:"protectionDomain,omitempty"` // protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.
	Sslenabled bool `json:"sslEnabled,omitempty"` // sslEnabled is the flag to enable/disable SSL communication with Gateway, default false
	Storagemode string `json:"storageMode,omitempty"` // storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.
	Storagepool string `json:"storagePool,omitempty"` // storagePool is the ScaleIO Storage Pool associated with the protection domain.
	Gateway string `json:"gateway"` // gateway is the host address of the ScaleIO API Gateway.
	Readonly bool `json:"readOnly,omitempty"` // readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
}

// Iok8sapicorev1AttachedVolume represents the Iok8sapicorev1AttachedVolume schema from the OpenAPI specification
type Iok8sapicorev1AttachedVolume struct {
	Devicepath string `json:"devicePath"` // DevicePath represents the device path where the volume should be available
	Name string `json:"name"` // Name of the attached volume
}

// Iok8sapicorev1Node represents the Iok8sapicorev1Node schema from the OpenAPI specification
type Iok8sapicorev1Node struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1NodeSpec `json:"spec,omitempty"` // NodeSpec describes the attributes that a node is created with.
	Status Iok8sapicorev1NodeStatus `json:"status,omitempty"` // NodeStatus is information about the current status of a node.
}

// Iok8sapicorev1PodTemplateSpec represents the Iok8sapicorev1PodTemplateSpec schema from the OpenAPI specification
type Iok8sapicorev1PodTemplateSpec struct {
	Spec Iok8sapicorev1PodSpec `json:"spec,omitempty"` // PodSpec is a description of a pod.
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1HostAlias represents the Iok8sapicorev1HostAlias schema from the OpenAPI specification
type Iok8sapicorev1HostAlias struct {
	Ip string `json:"ip"` // IP address of the host file entry.
	Hostnames []string `json:"hostnames,omitempty"` // Hostnames for the above IP address.
}

// Iok8sapicoordinationv1beta1LeaseCandidateSpec represents the Iok8sapicoordinationv1beta1LeaseCandidateSpec schema from the OpenAPI specification
type Iok8sapicoordinationv1beta1LeaseCandidateSpec struct {
	Binaryversion string `json:"binaryVersion"` // BinaryVersion is the binary version. It must be in a semver format without leading `v`. This field is required.
	Emulationversion string `json:"emulationVersion,omitempty"` // EmulationVersion is the emulation version. It must be in a semver format without leading `v`. EmulationVersion must be less than or equal to BinaryVersion. This field is required when strategy is "OldestEmulationVersion"
	Leasename string `json:"leaseName"` // LeaseName is the name of the lease for which this candidate is contending. The limits on this field are the same as on Lease.name. Multiple lease candidates may reference the same Lease.name. This field is immutable.
	Pingtime string `json:"pingTime,omitempty"` // MicroTime is version of Time with microsecond level precision.
	Renewtime string `json:"renewTime,omitempty"` // MicroTime is version of Time with microsecond level precision.
	Strategy string `json:"strategy"` // Strategy is the strategy that coordinated leader election will use for picking the leader. If multiple candidates for the same Lease return different strategies, the strategy provided by the candidate with the latest BinaryVersion will be used. If there is still conflict, this is a user error and coordinated leader election will not operate the Lease until resolved.
}

// Iok8sapinetworkingv1IngressList represents the Iok8sapinetworkingv1IngressList schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapinetworkingv1Ingress `json:"items"` // items is the list of Ingress.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiauthenticationv1TokenReview represents the Iok8sapiauthenticationv1TokenReview schema from the OpenAPI specification
type Iok8sapiauthenticationv1TokenReview struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthenticationv1TokenReviewSpec `json:"spec"` // TokenReviewSpec is a description of the token authentication request.
	Status Iok8sapiauthenticationv1TokenReviewStatus `json:"status,omitempty"` // TokenReviewStatus is the result of the token authentication request.
}

// Iok8sapiappsv1StatefulSetCondition represents the Iok8sapiappsv1StatefulSetCondition schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSetCondition struct {
	TypeField string `json:"type"` // Type of statefulset condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
}

// Iok8sapiflowcontrolv1ExemptPriorityLevelConfiguration represents the Iok8sapiflowcontrolv1ExemptPriorityLevelConfiguration schema from the OpenAPI specification
type Iok8sapiflowcontrolv1ExemptPriorityLevelConfiguration struct {
	Lendablepercent int `json:"lendablePercent,omitempty"` // `lendablePercent` prescribes the fraction of the level's NominalCL that can be borrowed by other priority levels. This value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level's LendableConcurrencyLimit (LendableCL), is defined as follows. LendableCL(i) = round( NominalCL(i) * lendablePercent(i)/100.0 )
	Nominalconcurrencyshares int `json:"nominalConcurrencyShares,omitempty"` // `nominalConcurrencyShares` (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats nominally reserved for this priority level. This DOES NOT limit the dispatching from this priority level but affects the other priority levels through the borrowing mechanism. The server's concurrency limit (ServerCL) is divided among all the priority levels in proportion to their NCS values: NominalCL(i) = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[priority level k] NCS(k) Bigger numbers mean a larger nominal concurrency limit, at the expense of every other priority level. This field has a default value of zero.
}

// Iok8sapiflowcontrolv1GroupSubject represents the Iok8sapiflowcontrolv1GroupSubject schema from the OpenAPI specification
type Iok8sapiflowcontrolv1GroupSubject struct {
	Name string `json:"name"` // name is the user group that matches, or "*" to match all user groups. See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required.
}

// Iok8sapicorev1NodeConfigSource represents the Iok8sapicorev1NodeConfigSource schema from the OpenAPI specification
type Iok8sapicorev1NodeConfigSource struct {
	Configmap Iok8sapicorev1ConfigMapNodeConfigSource `json:"configMap,omitempty"` // ConfigMapNodeConfigSource contains the information to reference a ConfigMap as a config source for the Node. This API is deprecated since 1.22: https://git.k8s.io/enhancements/keps/sig-node/281-dynamic-kubelet-configuration
}

// Iok8sapinetworkingv1IngressClassSpec represents the Iok8sapinetworkingv1IngressClassSpec schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressClassSpec struct {
	Controller string `json:"controller,omitempty"` // controller refers to the name of the controller that should handle this class. This allows for different "flavors" that are controlled by the same controller. For example, you may have different parameters for the same implementing controller. This should be specified as a domain-prefixed path no more than 250 characters in length, e.g. "acme.io/ingress-controller". This field is immutable.
	Parameters Iok8sapinetworkingv1IngressClassParametersReference `json:"parameters,omitempty"` // IngressClassParametersReference identifies an API object. This can be used to specify a cluster or namespace-scoped resource.
}

// Iok8sapiadmissionregistrationv1MutatingWebhookConfigurationList represents the Iok8sapiadmissionregistrationv1MutatingWebhookConfigurationList schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1MutatingWebhookConfigurationList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiadmissionregistrationv1MutatingWebhookConfiguration `json:"items"` // List of MutatingWebhookConfiguration.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapirbacv1Role represents the Iok8sapirbacv1Role schema from the OpenAPI specification
type Iok8sapirbacv1Role struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Rules []Iok8sapirbacv1PolicyRule `json:"rules,omitempty"` // Rules holds all the PolicyRules for this Role
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapiappsv1ReplicaSetStatus represents the Iok8sapiappsv1ReplicaSetStatus schema from the OpenAPI specification
type Iok8sapiappsv1ReplicaSetStatus struct {
	Terminatingreplicas int `json:"terminatingReplicas,omitempty"` // The number of terminating pods for this replica set. Terminating pods have a non-null .metadata.deletionTimestamp and have not yet reached the Failed or Succeeded .status.phase. This is an alpha field. Enable DeploymentReplicaSetTerminatingReplicas to be able to use this field.
	Availablereplicas int `json:"availableReplicas,omitempty"` // The number of available non-terminating pods (ready for at least minReadySeconds) for this replica set.
	Conditions []Iok8sapiappsv1ReplicaSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a replica set's current state.
	Fullylabeledreplicas int `json:"fullyLabeledReplicas,omitempty"` // The number of non-terminating pods that have labels matching the labels of the pod template of the replicaset.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
	Readyreplicas int `json:"readyReplicas,omitempty"` // The number of non-terminating pods targeted by this ReplicaSet with a Ready Condition.
	Replicas int `json:"replicas"` // Replicas is the most recently observed number of non-terminating pods. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicaset
}

// Iok8sapicorev1ServiceAccountList represents the Iok8sapicorev1ServiceAccountList schema from the OpenAPI specification
type Iok8sapicorev1ServiceAccountList struct {
	Items []Iok8sapicorev1ServiceAccount `json:"items"` // List of ServiceAccounts. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1Secret represents the Iok8sapicorev1Secret schema from the OpenAPI specification
type Iok8sapicorev1Secret struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Data map[string]interface{} `json:"data,omitempty"` // Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
	Immutable bool `json:"immutable,omitempty"` // Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Stringdata map[string]interface{} `json:"stringData,omitempty"` // stringData allows specifying non-binary secret data in string form. It is provided as a write-only input field for convenience. All keys and values are merged into the data field on write, overwriting any existing values. The stringData field is never output when reading from the API.
	TypeField string `json:"type,omitempty"` // Used to facilitate programmatic handling of secret data. More info: https://kubernetes.io/docs/concepts/configuration/secret/#secret-types
}

// Iok8sapicorev1HTTPHeader represents the Iok8sapicorev1HTTPHeader schema from the OpenAPI specification
type Iok8sapicorev1HTTPHeader struct {
	Name string `json:"name"` // The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.
	Value string `json:"value"` // The header field value
}

// Iok8sapiadmissionregistrationv1alpha1MatchResources represents the Iok8sapiadmissionregistrationv1alpha1MatchResources schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1MatchResources struct {
	Resourcerules []Iok8sapiadmissionregistrationv1alpha1NamedRuleWithOperations `json:"resourceRules,omitempty"` // ResourceRules describes what operations on what resources/subresources the admission policy matches. The policy cares about an operation if it matches _any_ Rule.
	Excluderesourcerules []Iok8sapiadmissionregistrationv1alpha1NamedRuleWithOperations `json:"excludeResourceRules,omitempty"` // ExcludeResourceRules describes what operations on what resources/subresources the policy should not care about. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)
	Matchpolicy string `json:"matchPolicy,omitempty"` // matchPolicy defines how the "MatchResources" list is used to match incoming requests. Allowed values are "Exact" or "Equivalent". - Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, the admission policy does not consider requests to apps/v1beta1 or extensions/v1beta1 API groups. - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, the admission policy **does** consider requests made to apps/v1beta1 or extensions/v1beta1 API groups. The API server translates the request to a matched resource API if necessary. Defaults to "Equivalent"
	Namespaceselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"namespaceSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Objectselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"objectSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapicorev1PersistentVolumeClaimVolumeSource represents the Iok8sapicorev1PersistentVolumeClaimVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaimVolumeSource struct {
	Claimname string `json:"claimName"` // claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Readonly bool `json:"readOnly,omitempty"` // readOnly Will force the ReadOnly setting in VolumeMounts. Default false.
}

// Iok8sapimachinerypkgapismetav1LabelSelector represents the Iok8sapimachinerypkgapismetav1LabelSelector schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1LabelSelector struct {
	Matchexpressions []Iok8sapimachinerypkgapismetav1LabelSelectorRequirement `json:"matchExpressions,omitempty"` // matchExpressions is a list of label selector requirements. The requirements are ANDed.
	Matchlabels map[string]interface{} `json:"matchLabels,omitempty"` // matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
}

// Iok8sapiresourcev1beta1DeviceAttribute represents the Iok8sapiresourcev1beta1DeviceAttribute schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceAttribute struct {
	BoolField bool `json:"bool,omitempty"` // BoolValue is a true/false value.
	IntField int64 `json:"int,omitempty"` // IntValue is a number.
	StringField string `json:"string,omitempty"` // StringValue is a string. Must not be longer than 64 characters.
	Version string `json:"version,omitempty"` // VersionValue is a semantic version according to semver.org spec 2.0.0. Must not be longer than 64 characters.
}

// Iok8sapiadmissionregistrationv1alpha1Variable represents the Iok8sapiadmissionregistrationv1alpha1Variable schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1Variable struct {
	Expression string `json:"expression"` // Expression is the expression that will be evaluated as the value of the variable. The CEL expression has access to the same identifiers as the CEL expressions in Validation.
	Name string `json:"name"` // Name is the name of the variable. The name must be a valid CEL identifier and unique among all variables. The variable can be accessed in other expressions through `variables` For example, if name is "foo", the variable will be available as `variables.foo`
}

// Iok8sapicorev1DownwardAPIVolumeFile represents the Iok8sapicorev1DownwardAPIVolumeFile schema from the OpenAPI specification
type Iok8sapicorev1DownwardAPIVolumeFile struct {
	Fieldref Iok8sapicorev1ObjectFieldSelector `json:"fieldRef,omitempty"` // ObjectFieldSelector selects an APIVersioned field of an object.
	Mode int `json:"mode,omitempty"` // Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Path string `json:"path"` // Required: Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	Resourcefieldref Iok8sapicorev1ResourceFieldSelector `json:"resourceFieldRef,omitempty"` // ResourceFieldSelector represents container resources (cpu, memory) and their output format
}

// Iok8sapicorev1ObjectFieldSelector represents the Iok8sapicorev1ObjectFieldSelector schema from the OpenAPI specification
type Iok8sapicorev1ObjectFieldSelector struct {
	Apiversion string `json:"apiVersion,omitempty"` // Version of the schema the FieldPath is written in terms of, defaults to "v1".
	Fieldpath string `json:"fieldPath"` // Path of the field to select in the specified API version.
}

// Iok8sapiappsv1DeploymentList represents the Iok8sapiappsv1DeploymentList schema from the OpenAPI specification
type Iok8sapiappsv1DeploymentList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiappsv1Deployment `json:"items"` // Items is the list of Deployments.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapistoragev1VolumeAttachmentSource represents the Iok8sapistoragev1VolumeAttachmentSource schema from the OpenAPI specification
type Iok8sapistoragev1VolumeAttachmentSource struct {
	Inlinevolumespec Iok8sapicorev1PersistentVolumeSpec `json:"inlineVolumeSpec,omitempty"` // PersistentVolumeSpec is the specification of a persistent volume.
	Persistentvolumename string `json:"persistentVolumeName,omitempty"` // persistentVolumeName represents the name of the persistent volume to attach.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path,omitempty"` // path is an optional URL path at which the webhook will be contacted.
	Port int `json:"port,omitempty"` // port is an optional service port at which the webhook will be contacted. `port` should be a valid port number (1-65535, inclusive). Defaults to 443 for backward compatibility.
	Name string `json:"name"` // name is the name of the service. Required
	Namespace string `json:"namespace"` // namespace is the namespace of the service. Required
}

// Iok8sapiautoscalingv2CrossVersionObjectReference represents the Iok8sapiautoscalingv2CrossVersionObjectReference schema from the OpenAPI specification
type Iok8sapiautoscalingv2CrossVersionObjectReference struct {
	Apiversion string `json:"apiVersion,omitempty"` // apiVersion is the API version of the referent
	Kind string `json:"kind"` // kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Name string `json:"name"` // name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
}

// Iok8sapicorev1TopologySelectorTerm represents the Iok8sapicorev1TopologySelectorTerm schema from the OpenAPI specification
type Iok8sapicorev1TopologySelectorTerm struct {
	Matchlabelexpressions []Iok8sapicorev1TopologySelectorLabelRequirement `json:"matchLabelExpressions,omitempty"` // A list of topology selector requirements by labels.
}

// Iok8sapicorev1ObjectReference represents the Iok8sapicorev1ObjectReference schema from the OpenAPI specification
type Iok8sapicorev1ObjectReference struct {
	Name string `json:"name,omitempty"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Namespace string `json:"namespace,omitempty"` // Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Resourceversion string `json:"resourceVersion,omitempty"` // Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
	Uid string `json:"uid,omitempty"` // UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
	Apiversion string `json:"apiVersion,omitempty"` // API version of the referent.
	Fieldpath string `json:"fieldPath,omitempty"` // If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.
	Kind string `json:"kind,omitempty"` // Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapimachinerypkgapismetav1DeleteOptions represents the Iok8sapimachinerypkgapismetav1DeleteOptions schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1DeleteOptions struct {
	Propagationpolicy string `json:"propagationPolicy,omitempty"` // Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Dryrun []string `json:"dryRun,omitempty"` // When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed
	Graceperiodseconds int64 `json:"gracePeriodSeconds,omitempty"` // The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.
	Ignorestorereaderrorwithclusterbreakingpotential bool `json:"ignoreStoreReadErrorWithClusterBreakingPotential,omitempty"` // if set to true, it will trigger an unsafe deletion of the resource in case the normal deletion flow fails with a corrupt object error. A resource is considered corrupt if it can not be retrieved from the underlying storage successfully because of a) its data can not be transformed e.g. decryption failure, or b) it fails to decode into an object. NOTE: unsafe deletion ignores finalizer constraints, skips precondition checks, and removes the object from the storage. WARNING: This may potentially break the cluster if the workload associated with the resource being unsafe-deleted relies on normal deletion flow. Use only if you REALLY know what you are doing. The default value is false, and the user must opt in to enable it
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Orphandependents bool `json:"orphanDependents,omitempty"` // Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the "orphan" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.
	Preconditions Iok8sapimachinerypkgapismetav1Preconditions `json:"preconditions,omitempty"` // Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out.
}

// Iok8sapiflowcontrolv1ServiceAccountSubject represents the Iok8sapiflowcontrolv1ServiceAccountSubject schema from the OpenAPI specification
type Iok8sapiflowcontrolv1ServiceAccountSubject struct {
	Name string `json:"name"` // `name` is the name of matching ServiceAccount objects, or "*" to match regardless of name. Required.
	Namespace string `json:"namespace"` // `namespace` is the namespace of matching ServiceAccount objects. Required.
}

// Iok8sapiresourcev1beta2ResourceClaimStatus represents the Iok8sapiresourcev1beta2ResourceClaimStatus schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ResourceClaimStatus struct {
	Allocation Iok8sapiresourcev1beta2AllocationResult `json:"allocation,omitempty"` // AllocationResult contains attributes of an allocated resource.
	Devices []Iok8sapiresourcev1beta2AllocatedDeviceStatus `json:"devices,omitempty"` // Devices contains the status of each device allocated for this claim, as reported by the driver. This can include driver-specific information. Entries are owned by their respective drivers.
	Reservedfor []Iok8sapiresourcev1beta2ResourceClaimConsumerReference `json:"reservedFor,omitempty"` // ReservedFor indicates which entities are currently allowed to use the claim. A Pod which references a ResourceClaim which is not reserved for that Pod will not be started. A claim that is in use or might be in use because it has been reserved must not get deallocated. In a cluster with multiple scheduler instances, two pods might get scheduled concurrently by different schedulers. When they reference the same ResourceClaim which already has reached its maximum number of consumers, only one pod can be scheduled. Both schedulers try to add their pod to the claim.status.reservedFor field, but only the update that reaches the API server first gets stored. The other one fails with an error and the scheduler which issued it knows that it must put the pod back into the queue, waiting for the ResourceClaim to become usable again. There can be at most 256 such reservations. This may get increased in the future, but not reduced.
}

// Iok8sapinetworkingv1IngressLoadBalancerStatus represents the Iok8sapinetworkingv1IngressLoadBalancerStatus schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressLoadBalancerStatus struct {
	Ingress []Iok8sapinetworkingv1IngressLoadBalancerIngress `json:"ingress,omitempty"` // ingress is a list containing ingress points for the load-balancer.
}

// Iok8sapirbacv1Subject represents the Iok8sapirbacv1Subject schema from the OpenAPI specification
type Iok8sapirbacv1Subject struct {
	Namespace string `json:"namespace,omitempty"` // Namespace of the referenced object. If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
	Apigroup string `json:"apiGroup,omitempty"` // APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
	Kind string `json:"kind"` // Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
	Name string `json:"name"` // Name of the object being referenced.
}

// Iok8sapiautoscalingv1ScaleStatus represents the Iok8sapiautoscalingv1ScaleStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv1ScaleStatus struct {
	Replicas int `json:"replicas"` // replicas is the actual number of observed instances of the scaled object.
	Selector string `json:"selector,omitempty"` // selector is the label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
}

// Iok8sapiresourcev1beta2ResourceClaimTemplate represents the Iok8sapiresourcev1beta2ResourceClaimTemplate schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ResourceClaimTemplate struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiresourcev1beta2ResourceClaimTemplateSpec `json:"spec"` // ResourceClaimTemplateSpec contains the metadata and fields for a ResourceClaim.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapinetworkingv1IngressBackend represents the Iok8sapinetworkingv1IngressBackend schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressBackend struct {
	Resource Iok8sapicorev1TypedLocalObjectReference `json:"resource,omitempty"` // TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.
	Service Iok8sapinetworkingv1IngressServiceBackend `json:"service,omitempty"` // IngressServiceBackend references a Kubernetes Service as a Backend.
}

// Iok8sapimachinerypkgapismetav1ListMeta represents the Iok8sapimachinerypkgapismetav1ListMeta schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1ListMeta struct {
	ContinueField string `json:"continue,omitempty"` // continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.
	Remainingitemcount int64 `json:"remainingItemCount,omitempty"` // remainingItemCount is the number of subsequent items in the list which are not included in this list response. If the list request contained label or field selectors, then the number of remaining items is unknown and the field will be left unset and omitted during serialization. If the list is complete (either because it is not chunking or because this is the last chunk), then there are no more remaining items and this field will be left unset and omitted during serialization. Servers older than v1.15 do not set this field. The intended use of the remainingItemCount is *estimating* the size of a collection. Clients should not rely on the remainingItemCount to be set or to be exact.
	Resourceversion string `json:"resourceVersion,omitempty"` // String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
	Selflink string `json:"selfLink,omitempty"` // Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []GeneratedType `json:"items"` // Items is the list of APIService
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiresourcev1beta2CELDeviceSelector represents the Iok8sapiresourcev1beta2CELDeviceSelector schema from the OpenAPI specification
type Iok8sapiresourcev1beta2CELDeviceSelector struct {
	Expression string `json:"expression"` // Expression is a CEL expression which evaluates a single device. It must evaluate to true when the device under consideration satisfies the desired criteria, and false when it does not. Any other result is an error and causes allocation of devices to abort. The expression's input is an object named "device", which carries the following properties: - driver (string): the name of the driver which defines this device. - attributes (map[string]object): the device's attributes, grouped by prefix (e.g. device.attributes["dra.example.com"] evaluates to an object with all of the attributes which were prefixed by "dra.example.com". - capacity (map[string]object): the device's capacities, grouped by prefix. Example: Consider a device with driver="dra.example.com", which exposes two attributes named "model" and "ext.example.com/family" and which exposes one capacity named "modules". This input to this expression would have the following fields: device.driver device.attributes["dra.example.com"].model device.attributes["ext.example.com"].family device.capacity["dra.example.com"].modules The device.driver field can be used to check for a specific driver, either as a high-level precondition (i.e. you only want to consider devices from this driver) or as part of a multi-clause expression that is meant to consider devices from different drivers. The value type of each attribute is defined by the device definition, and users who write these expressions must consult the documentation for their specific drivers. The value type of each capacity is Quantity. If an unknown prefix is used as a lookup in either device.attributes or device.capacity, an empty map will be returned. Any reference to an unknown field will cause an evaluation error and allocation to abort. A robust expression should check for the existence of attributes before referencing them. For ease of use, the cel.bind() function is enabled, and can be used to simplify expressions that access multiple attributes with the same domain. For example: cel.bind(dra, device.attributes["dra.example.com"], dra.someBool && dra.anotherBool) The length of the expression must be smaller or equal to 10 Ki. The cost of evaluating it is also limited based on the estimated number of logical steps.
}

// Iok8sapiadmissionregistrationv1MatchCondition represents the Iok8sapiadmissionregistrationv1MatchCondition schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1MatchCondition struct {
	Expression string `json:"expression"` // Expression represents the expression which will be evaluated by CEL. Must evaluate to bool. CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables: 'object' - The object from the incoming request. The value is null for DELETE requests. 'oldObject' - The existing object. The value is null for CREATE requests. 'request' - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest). 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request. See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz 'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the request resource. Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/ Required.
	Name string `json:"name"` // Name is an identifier for this match condition, used for strategic merging of MatchConditions, as well as providing an identifier for logging purposes. A good name should be descriptive of the associated expression. Name must be a qualified name consisting of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character (e.g. 'MyName', or 'my.name', or '123-abc', regex used for validation is '([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]') with an optional DNS subdomain prefix and '/' (e.g. 'example.com/MyName') Required.
}

// Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyList represents the Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyList schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicy `json:"items"` // List of ValidatingAdmissionPolicy.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapipolicyv1PodDisruptionBudgetSpec represents the Iok8sapipolicyv1PodDisruptionBudgetSpec schema from the OpenAPI specification
type Iok8sapipolicyv1PodDisruptionBudgetSpec struct {
	Maxunavailable string `json:"maxUnavailable,omitempty"` // IntOrString is a type that can hold an int32 or a string. When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type. This allows you to have, for example, a JSON field that can accept a name or number.
	Minavailable string `json:"minAvailable,omitempty"` // IntOrString is a type that can hold an int32 or a string. When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type. This allows you to have, for example, a JSON field that can accept a name or number.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Unhealthypodevictionpolicy string `json:"unhealthyPodEvictionPolicy,omitempty"` // UnhealthyPodEvictionPolicy defines the criteria for when unhealthy pods should be considered for eviction. Current implementation considers healthy pods, as pods that have status.conditions item with type="Ready",status="True". Valid policies are IfHealthyBudget and AlwaysAllow. If no policy is specified, the default behavior will be used, which corresponds to the IfHealthyBudget policy. IfHealthyBudget policy means that running pods (status.phase="Running"), but not yet healthy can be evicted only if the guarded application is not disrupted (status.currentHealthy is at least equal to status.desiredHealthy). Healthy pods will be subject to the PDB for eviction. AlwaysAllow policy means that all running pods (status.phase="Running"), but not yet healthy are considered disrupted and can be evicted regardless of whether the criteria in a PDB is met. This means perspective running pods of a disrupted application might not get a chance to become healthy. Healthy pods will be subject to the PDB for eviction. Additional policies may be added in the future. Clients making eviction decisions should disallow eviction of unhealthy pods if they encounter an unrecognized policy in this field.
}

// Iok8sapiresourcev1beta1CounterSet represents the Iok8sapiresourcev1beta1CounterSet schema from the OpenAPI specification
type Iok8sapiresourcev1beta1CounterSet struct {
	Name string `json:"name"` // Name defines the name of the counter set. It must be a DNS label.
	Counters map[string]interface{} `json:"counters"` // Counters defines the set of counters for this CounterSet The name of each counter must be unique in that set and must be a DNS label. The maximum number of counters is 32.
}

// Iok8sapiappsv1StatefulSetPersistentVolumeClaimRetentionPolicy represents the Iok8sapiappsv1StatefulSetPersistentVolumeClaimRetentionPolicy schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSetPersistentVolumeClaimRetentionPolicy struct {
	Whenscaled string `json:"whenScaled,omitempty"` // WhenScaled specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is scaled down. The default policy of `Retain` causes PVCs to not be affected by a scaledown. The `Delete` policy causes the associated PVCs for any excess pods above the replica count to be deleted.
	Whendeleted string `json:"whenDeleted,omitempty"` // WhenDeleted specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is deleted. The default policy of `Retain` causes PVCs to not be affected by StatefulSet deletion. The `Delete` policy causes those PVCs to be deleted.
}

// Iok8sapinetworkingv1IPBlock represents the Iok8sapinetworkingv1IPBlock schema from the OpenAPI specification
type Iok8sapinetworkingv1IPBlock struct {
	Cidr string `json:"cidr"` // cidr is a string representing the IPBlock Valid examples are "192.168.1.0/24" or "2001:db8::/64"
	Except []string `json:"except,omitempty"` // except is a slice of CIDRs that should not be included within an IPBlock Valid examples are "192.168.1.0/24" or "2001:db8::/64" Except values will be rejected if they are outside the cidr range
}

// Iok8sapiresourcev1beta1DeviceSelector represents the Iok8sapiresourcev1beta1DeviceSelector schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceSelector struct {
	Cel Iok8sapiresourcev1beta1CELDeviceSelector `json:"cel,omitempty"` // CELDeviceSelector contains a CEL expression for selecting a device.
}

// Iok8sapiresourcev1beta2ResourceClaimTemplateSpec represents the Iok8sapiresourcev1beta2ResourceClaimTemplateSpec schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ResourceClaimTemplateSpec struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiresourcev1beta2ResourceClaimSpec `json:"spec"` // ResourceClaimSpec defines what is being requested in a ResourceClaim and how to configure it.
}

// Iok8sapiautoscalingv2ExternalMetricStatus represents the Iok8sapiautoscalingv2ExternalMetricStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2ExternalMetricStatus struct {
	Current Iok8sapiautoscalingv2MetricValueStatus `json:"current"` // MetricValueStatus holds the current value for a metric
	Metric Iok8sapiautoscalingv2MetricIdentifier `json:"metric"` // MetricIdentifier defines the name and optionally selector for a metric
}

// Iok8sapicorev1AzureDiskVolumeSource represents the Iok8sapicorev1AzureDiskVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1AzureDiskVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Kind string `json:"kind,omitempty"` // kind expected values are Shared: multiple blob disks per storage account Dedicated: single blob disk per storage account Managed: azure managed data disk (only in managed availability set). defaults to shared
	Readonly bool `json:"readOnly,omitempty"` // readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Cachingmode string `json:"cachingMode,omitempty"` // cachingMode is the Host Caching mode: None, Read Only, Read Write.
	Diskname string `json:"diskName"` // diskName is the Name of the data disk in the blob storage
	Diskuri string `json:"diskURI"` // diskURI is the URI of data disk in the blob storage
}

// Iok8sapicorev1EventSeries represents the Iok8sapicorev1EventSeries schema from the OpenAPI specification
type Iok8sapicorev1EventSeries struct {
	Count int `json:"count,omitempty"` // Number of occurrences in this series up to the last heartbeat time
	Lastobservedtime string `json:"lastObservedTime,omitempty"` // MicroTime is version of Time with microsecond level precision.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Format string `json:"format,omitempty"` // format is an optional OpenAPI type definition for this column. The 'name' format is applied to the primary identifier column to assist in clients identifying column is the resource name. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.
	Jsonpath string `json:"jsonPath"` // jsonPath is a simple JSON path (i.e. with array notation) which is evaluated against each custom resource to produce the value for this column.
	Name string `json:"name"` // name is a human readable name for the column.
	Priority int `json:"priority,omitempty"` // priority is an integer defining the relative importance of this column compared to others. Lower numbers are considered higher priority. Columns that may be omitted in limited space scenarios should be given a priority greater than 0.
	TypeField string `json:"type"` // type is an OpenAPI type definition for this column. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.
	Description string `json:"description,omitempty"` // description is a human readable description of this column.
}

// Iok8sapicorev1ComponentCondition represents the Iok8sapicorev1ComponentCondition schema from the OpenAPI specification
type Iok8sapicorev1ComponentCondition struct {
	Status string `json:"status"` // Status of the condition for a component. Valid values for "Healthy": "True", "False", or "Unknown".
	TypeField string `json:"type"` // Type of condition for a component. Valid value: "Healthy"
	ErrorField string `json:"error,omitempty"` // Condition error code for a component. For example, a health check error code.
	Message string `json:"message,omitempty"` // Message about the condition for a component. For example, information about a health check.
}

// Iok8sapicorev1TCPSocketAction represents the Iok8sapicorev1TCPSocketAction schema from the OpenAPI specification
type Iok8sapicorev1TCPSocketAction struct {
	Host string `json:"host,omitempty"` // Optional: Host name to connect to, defaults to the pod IP.
	Port string `json:"port"` // IntOrString is a type that can hold an int32 or a string. When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type. This allows you to have, for example, a JSON field that can accept a name or number.
}

// Iok8sapicertificatesv1CertificateSigningRequestCondition represents the Iok8sapicertificatesv1CertificateSigningRequestCondition schema from the OpenAPI specification
type Iok8sapicertificatesv1CertificateSigningRequestCondition struct {
	TypeField string `json:"type"` // type of the condition. Known conditions are "Approved", "Denied", and "Failed". An "Approved" condition is added via the /approval subresource, indicating the request was approved and should be issued by the signer. A "Denied" condition is added via the /approval subresource, indicating the request was denied and should not be issued by the signer. A "Failed" condition is added via the /status subresource, indicating the signer failed to issue the certificate. Approved and Denied conditions are mutually exclusive. Approved, Denied, and Failed conditions cannot be removed once added. Only one condition of a given type is allowed.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Lastupdatetime string `json:"lastUpdateTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // message contains a human readable message with details about the request state
	Reason string `json:"reason,omitempty"` // reason indicates a brief reason for the request state
	Status string `json:"status"` // status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or "Unknown".
}

// Iok8sapiresourcev1alpha3CELDeviceSelector represents the Iok8sapiresourcev1alpha3CELDeviceSelector schema from the OpenAPI specification
type Iok8sapiresourcev1alpha3CELDeviceSelector struct {
	Expression string `json:"expression"` // Expression is a CEL expression which evaluates a single device. It must evaluate to true when the device under consideration satisfies the desired criteria, and false when it does not. Any other result is an error and causes allocation of devices to abort. The expression's input is an object named "device", which carries the following properties: - driver (string): the name of the driver which defines this device. - attributes (map[string]object): the device's attributes, grouped by prefix (e.g. device.attributes["dra.example.com"] evaluates to an object with all of the attributes which were prefixed by "dra.example.com". - capacity (map[string]object): the device's capacities, grouped by prefix. Example: Consider a device with driver="dra.example.com", which exposes two attributes named "model" and "ext.example.com/family" and which exposes one capacity named "modules". This input to this expression would have the following fields: device.driver device.attributes["dra.example.com"].model device.attributes["ext.example.com"].family device.capacity["dra.example.com"].modules The device.driver field can be used to check for a specific driver, either as a high-level precondition (i.e. you only want to consider devices from this driver) or as part of a multi-clause expression that is meant to consider devices from different drivers. The value type of each attribute is defined by the device definition, and users who write these expressions must consult the documentation for their specific drivers. The value type of each capacity is Quantity. If an unknown prefix is used as a lookup in either device.attributes or device.capacity, an empty map will be returned. Any reference to an unknown field will cause an evaluation error and allocation to abort. A robust expression should check for the existence of attributes before referencing them. For ease of use, the cel.bind() function is enabled, and can be used to simplify expressions that access multiple attributes with the same domain. For example: cel.bind(dra, device.attributes["dra.example.com"], dra.someBool && dra.anotherBool) The length of the expression must be smaller or equal to 10 Ki. The cost of evaluating it is also limited based on the estimated number of logical steps.
}

// Iok8sapiresourcev1beta1DeviceCounterConsumption represents the Iok8sapiresourcev1beta1DeviceCounterConsumption schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceCounterConsumption struct {
	Counters map[string]interface{} `json:"counters"` // Counters defines the counters that will be consumed by the device. The maximum number counters in a device is 32. In addition, the maximum number of all counters in all devices is 1024 (for example, 64 devices with 16 counters each).
	Counterset string `json:"counterSet"` // CounterSet is the name of the set from which the counters defined will be consumed.
}

// Iok8sapicorev1Event represents the Iok8sapicorev1Event schema from the OpenAPI specification
type Iok8sapicorev1Event struct {
	Eventtime string `json:"eventTime,omitempty"` // MicroTime is version of Time with microsecond level precision.
	Reason string `json:"reason,omitempty"` // This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
	Series Iok8sapicorev1EventSeries `json:"series,omitempty"` // EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Firsttimestamp string `json:"firstTimestamp,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Related Iok8sapicorev1ObjectReference `json:"related,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Reportingcomponent string `json:"reportingComponent,omitempty"` // Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	Reportinginstance string `json:"reportingInstance,omitempty"` // ID of the controller instance, e.g. `kubelet-xyzf`.
	Source Iok8sapicorev1EventSource `json:"source,omitempty"` // EventSource contains information for an event.
	Message string `json:"message,omitempty"` // A human-readable description of the status of this operation.
	TypeField string `json:"type,omitempty"` // Type of this event (Normal, Warning), new types could be added in the future
	Lasttimestamp string `json:"lastTimestamp,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Count int `json:"count,omitempty"` // The number of times this event has occurred.
	Involvedobject Iok8sapicorev1ObjectReference `json:"involvedObject"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Action string `json:"action,omitempty"` // What action was taken/failed regarding to the Regarding object.
}

// Iok8sapistoragev1beta1VolumeAttributesClassList represents the Iok8sapistoragev1beta1VolumeAttributesClassList schema from the OpenAPI specification
type Iok8sapistoragev1beta1VolumeAttributesClassList struct {
	Items []Iok8sapistoragev1beta1VolumeAttributesClass `json:"items"` // items is the list of VolumeAttributesClass objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1ContainerStateTerminated represents the Iok8sapicorev1ContainerStateTerminated schema from the OpenAPI specification
type Iok8sapicorev1ContainerStateTerminated struct {
	Containerid string `json:"containerID,omitempty"` // Container's ID in the format '<type>://<container_id>'
	Exitcode int `json:"exitCode"` // Exit status from the last termination of the container
	Finishedat string `json:"finishedAt,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // Message regarding the last termination of the container
	Reason string `json:"reason,omitempty"` // (brief) reason from the last termination of the container
	Signal int `json:"signal,omitempty"` // Signal from the last termination of the container
	Startedat string `json:"startedAt,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
}

// Iok8sapiauthenticationv1TokenRequestStatus represents the Iok8sapiauthenticationv1TokenRequestStatus schema from the OpenAPI specification
type Iok8sapiauthenticationv1TokenRequestStatus struct {
	Expirationtimestamp string `json:"expirationTimestamp"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Token string `json:"token"` // Token is the opaque bearer token.
}

// Iok8sapiresourcev1beta1NetworkDeviceData represents the Iok8sapiresourcev1beta1NetworkDeviceData schema from the OpenAPI specification
type Iok8sapiresourcev1beta1NetworkDeviceData struct {
	Hardwareaddress string `json:"hardwareAddress,omitempty"` // HardwareAddress represents the hardware address (e.g. MAC Address) of the device's network interface. Must not be longer than 128 characters.
	Interfacename string `json:"interfaceName,omitempty"` // InterfaceName specifies the name of the network interface associated with the allocated device. This might be the name of a physical or virtual network interface being configured in the pod. Must not be longer than 256 characters.
	Ips []string `json:"ips,omitempty"` // IPs lists the network addresses assigned to the device's network interface. This can include both IPv4 and IPv6 addresses. The IPs are in the CIDR notation, which includes both the address and the associated subnet mask. e.g.: "192.0.2.5/24" for IPv4 and "2001:db8::5/64" for IPv6. Must not contain more than 16 entries.
}

// Iok8sapicorev1PodTemplate represents the Iok8sapicorev1PodTemplate schema from the OpenAPI specification
type Iok8sapicorev1PodTemplate struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Template Iok8sapicorev1PodTemplateSpec `json:"template,omitempty"` // PodTemplateSpec describes the data a pod should have when created from a template
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapiresourcev1beta1AllocationResult represents the Iok8sapiresourcev1beta1AllocationResult schema from the OpenAPI specification
type Iok8sapiresourcev1beta1AllocationResult struct {
	Devices Iok8sapiresourcev1beta1DeviceAllocationResult `json:"devices,omitempty"` // DeviceAllocationResult is the result of allocating devices.
	Nodeselector Iok8sapicorev1NodeSelector `json:"nodeSelector,omitempty"` // A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
}

// Iok8sapiadmissionregistrationv1AuditAnnotation represents the Iok8sapiadmissionregistrationv1AuditAnnotation schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1AuditAnnotation struct {
	Key string `json:"key"` // key specifies the audit annotation key. The audit annotation keys of a ValidatingAdmissionPolicy must be unique. The key must be a qualified name ([A-Za-z0-9][-A-Za-z0-9_.]*) no more than 63 bytes in length. The key is combined with the resource name of the ValidatingAdmissionPolicy to construct an audit annotation key: "{ValidatingAdmissionPolicy name}/{key}". If an admission webhook uses the same resource name as this ValidatingAdmissionPolicy and the same audit annotation key, the annotation key will be identical. In this case, the first annotation written with the key will be included in the audit event and all subsequent annotations with the same key will be discarded. Required.
	Valueexpression string `json:"valueExpression"` // valueExpression represents the expression which is evaluated by CEL to produce an audit annotation value. The expression must evaluate to either a string or null value. If the expression evaluates to a string, the audit annotation is included with the string value. If the expression evaluates to null or empty string the audit annotation will be omitted. The valueExpression may be no longer than 5kb in length. If the result of the valueExpression is more than 10kb in length, it will be truncated to 10kb. If multiple ValidatingAdmissionPolicyBinding resources match an API request, then the valueExpression will be evaluated for each binding. All unique values produced by the valueExpressions will be joined together in a comma-separated list. Required.
}

// Iok8sapicorev1ExecAction represents the Iok8sapicorev1ExecAction schema from the OpenAPI specification
type Iok8sapicorev1ExecAction struct {
	Command []string `json:"command,omitempty"` // Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
}

// Iok8sapirbacv1ClusterRole represents the Iok8sapirbacv1ClusterRole schema from the OpenAPI specification
type Iok8sapirbacv1ClusterRole struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Rules []Iok8sapirbacv1PolicyRule `json:"rules,omitempty"` // Rules holds all the PolicyRules for this ClusterRole
	Aggregationrule Iok8sapirbacv1AggregationRule `json:"aggregationRule,omitempty"` // AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapiresourcev1beta2DeviceConstraint represents the Iok8sapiresourcev1beta2DeviceConstraint schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceConstraint struct {
	Matchattribute string `json:"matchAttribute,omitempty"` // MatchAttribute requires that all devices in question have this attribute and that its type and value are the same across those devices. For example, if you specified "dra.example.com/numa" (a hypothetical example!), then only devices in the same NUMA node will be chosen. A device which does not have that attribute will not be chosen. All devices should use a value of the same type for this attribute because that is part of its specification, but if one device doesn't, then it also will not be chosen. Must include the domain qualifier.
	Requests []string `json:"requests,omitempty"` // Requests is a list of the one or more requests in this claim which must co-satisfy this constraint. If a request is fulfilled by multiple devices, then all of the devices must satisfy the constraint. If this is not specified, this constraint applies to all requests in this claim. References to subrequests must include the name of the main request and may include the subrequest using the format <main request>[/<subrequest>]. If just the main request is given, the constraint applies to all subrequests.
}

// Iok8sapistoragev1CSIStorageCapacityList represents the Iok8sapistoragev1CSIStorageCapacityList schema from the OpenAPI specification
type Iok8sapistoragev1CSIStorageCapacityList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapistoragev1CSIStorageCapacity `json:"items"` // items is the list of CSIStorageCapacity objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapiautoscalingv1CrossVersionObjectReference represents the Iok8sapiautoscalingv1CrossVersionObjectReference schema from the OpenAPI specification
type Iok8sapiautoscalingv1CrossVersionObjectReference struct {
	Apiversion string `json:"apiVersion,omitempty"` // apiVersion is the API version of the referent
	Kind string `json:"kind"` // kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Name string `json:"name"` // name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
}

// Iok8sapistoragev1CSINodeDriver represents the Iok8sapistoragev1CSINodeDriver schema from the OpenAPI specification
type Iok8sapistoragev1CSINodeDriver struct {
	Name string `json:"name"` // name represents the name of the CSI driver that this object refers to. This MUST be the same name returned by the CSI GetPluginName() call for that driver.
	Nodeid string `json:"nodeID"` // nodeID of the node from the driver point of view. This field enables Kubernetes to communicate with storage systems that do not share the same nomenclature for nodes. For example, Kubernetes may refer to a given node as "node1", but the storage system may refer to the same node as "nodeA". When Kubernetes issues a command to the storage system to attach a volume to a specific node, it can use this field to refer to the node name using the ID that the storage system will understand, e.g. "nodeA" instead of "node1". This field is required.
	Topologykeys []string `json:"topologyKeys,omitempty"` // topologyKeys is the list of keys supported by the driver. When a driver is initialized on a cluster, it provides a set of topology keys that it understands (e.g. "company.com/zone", "company.com/region"). When a driver is initialized on a node, it provides the same topology keys along with values. Kubelet will expose these topology keys as labels on its own node object. When Kubernetes does topology aware provisioning, it can use this list to determine which labels it should retrieve from the node object and pass back to the driver. It is possible for different nodes to use different topology keys. This can be empty if driver does not support topology.
	Allocatable Iok8sapistoragev1VolumeNodeResources `json:"allocatable,omitempty"` // VolumeNodeResources is a set of resource limits for scheduling of volumes.
}

// Iok8sapicoordinationv1LeaseList represents the Iok8sapicoordinationv1LeaseList schema from the OpenAPI specification
type Iok8sapicoordinationv1LeaseList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicoordinationv1Lease `json:"items"` // items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapibatchv1JobSpec represents the Iok8sapibatchv1JobSpec schema from the OpenAPI specification
type Iok8sapibatchv1JobSpec struct {
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Podfailurepolicy Iok8sapibatchv1PodFailurePolicy `json:"podFailurePolicy,omitempty"` // PodFailurePolicy describes how failed pods influence the backoffLimit.
	Completions int `json:"completions,omitempty"` // Specifies the desired number of successfully finished pods the job should be run with. Setting to null means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value. Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Completionmode string `json:"completionMode,omitempty"` // completionMode specifies how Pod completions are tracked. It can be `NonIndexed` (default) or `Indexed`. `NonIndexed` means that the Job is considered complete when there have been .spec.completions successfully completed Pods. Each Pod completion is homologous to each other. `Indexed` means that the Pods of a Job get an associated completion index from 0 to (.spec.completions - 1), available in the annotation batch.kubernetes.io/job-completion-index. The Job is considered complete when there is one successfully completed Pod for each index. When value is `Indexed`, .spec.completions must be specified and `.spec.parallelism` must be less than or equal to 10^5. In addition, The Pod name takes the form `$(job-name)-$(index)-$(random-string)`, the Pod hostname takes the form `$(job-name)-$(index)`. More completion modes can be added in the future. If the Job controller observes a mode that it doesn't recognize, which is possible during upgrades due to version skew, the controller skips updates for the Job.
	Managedby string `json:"managedBy,omitempty"` // ManagedBy field indicates the controller that manages a Job. The k8s Job controller reconciles jobs which don't have this field at all or the field value is the reserved string `kubernetes.io/job-controller`, but skips reconciling Jobs with a custom value for this field. The value must be a valid domain-prefixed path (e.g. acme.io/foo) - all characters before the first "/" must be a valid subdomain as defined by RFC 1123. All characters trailing the first "/" must be valid HTTP Path characters as defined by RFC 3986. The value cannot exceed 63 characters. This field is immutable. This field is beta-level. The job controller accepts setting the field when the feature gate JobManagedBy is enabled (enabled by default).
	Maxfailedindexes int `json:"maxFailedIndexes,omitempty"` // Specifies the maximal number of failed indexes before marking the Job as failed, when backoffLimitPerIndex is set. Once the number of failed indexes exceeds this number the entire Job is marked as Failed and its execution is terminated. When left as null the job continues execution of all of its indexes and is marked with the `Complete` Job condition. It can only be specified when backoffLimitPerIndex is set. It can be null or up to completions. It is required and must be less than or equal to 10^4 when is completions greater than 10^5.
	Podreplacementpolicy string `json:"podReplacementPolicy,omitempty"` // podReplacementPolicy specifies when to create replacement Pods. Possible values are: - TerminatingOrFailed means that we recreate pods when they are terminating (has a metadata.deletionTimestamp) or failed. - Failed means to wait until a previously created Pod is fully terminated (has phase Failed or Succeeded) before creating a replacement Pod. When using podFailurePolicy, Failed is the the only allowed value. TerminatingOrFailed and Failed are allowed values when podFailurePolicy is not in use.
	Backofflimit int `json:"backoffLimit,omitempty"` // Specifies the number of retries before marking this job failed. Defaults to 6, unless backoffLimitPerIndex (only Indexed Job) is specified. When backoffLimitPerIndex is specified, backoffLimit defaults to 2147483647.
	Suspend bool `json:"suspend,omitempty"` // suspend specifies whether the Job controller should create Pods or not. If a Job is created with suspend set to true, no Pods are created by the Job controller. If a Job is suspended after creation (i.e. the flag goes from false to true), the Job controller will delete all active Pods associated with this Job. Users must design their workload to gracefully handle this. Suspending a Job will reset the StartTime field of the Job, effectively resetting the ActiveDeadlineSeconds timer too. Defaults to false.
	Successpolicy Iok8sapibatchv1SuccessPolicy `json:"successPolicy,omitempty"` // SuccessPolicy describes when a Job can be declared as succeeded based on the success of some indexes.
	Ttlsecondsafterfinished int `json:"ttlSecondsAfterFinished,omitempty"` // ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes.
	Parallelism int `json:"parallelism,omitempty"` // Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Activedeadlineseconds int64 `json:"activeDeadlineSeconds,omitempty"` // Specifies the duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it; value must be positive integer. If a Job is suspended (at creation or through an update), this timer will effectively be stopped and reset when the Job is resumed again.
	Manualselector bool `json:"manualSelector,omitempty"` // manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template. When true, the user is responsible for picking unique labels and specifying the selector. Failure to pick a unique label may cause this and other jobs to not function correctly. However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
	Backofflimitperindex int `json:"backoffLimitPerIndex,omitempty"` // Specifies the limit for the number of retries within an index before marking this index as failed. When enabled the number of failures per index is kept in the pod's batch.kubernetes.io/job-index-failure-count annotation. It can only be set when Job's completionMode=Indexed, and the Pod's restart policy is Never. The field is immutable.
}

// Iok8sapinetworkingv1HTTPIngressPath represents the Iok8sapinetworkingv1HTTPIngressPath schema from the OpenAPI specification
type Iok8sapinetworkingv1HTTPIngressPath struct {
	Backend Iok8sapinetworkingv1IngressBackend `json:"backend"` // IngressBackend describes all endpoints for a given service and port.
	Path string `json:"path,omitempty"` // path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value "Exact" or "Prefix".
	Pathtype string `json:"pathType"` // pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is done on a path element by element basis. A path element refers is the list of labels in the path split by the '/' separator. A request is a match for path p if every p is an element-wise prefix of p of the request path. Note that if the last element of the path is a substring of the last element in request path, it is not a match (e.g. /foo/bar matches /foo/bar/baz, but does not match /foo/barbaz). * ImplementationSpecific: Interpretation of the Path matching is up to the IngressClass. Implementations can treat this as a separate PathType or treat it identically to Prefix or Exact path types. Implementations are required to support all path types.
}

// Iok8sapicorev1ModifyVolumeStatus represents the Iok8sapicorev1ModifyVolumeStatus schema from the OpenAPI specification
type Iok8sapicorev1ModifyVolumeStatus struct {
	Targetvolumeattributesclassname string `json:"targetVolumeAttributesClassName,omitempty"` // targetVolumeAttributesClassName is the name of the VolumeAttributesClass the PVC currently being reconciled
	Status string `json:"status"` // status is the status of the ControllerModifyVolume operation. It can be in any of following states: - Pending Pending indicates that the PersistentVolumeClaim cannot be modified due to unmet requirements, such as the specified VolumeAttributesClass not existing. - InProgress InProgress indicates that the volume is being modified. - Infeasible Infeasible indicates that the request has been rejected as invalid by the CSI driver. To 	 resolve the error, a valid VolumeAttributesClass needs to be specified. Note: New statuses can be added in the future. Consumers should check for unknown statuses and fail appropriately.
}

// Iok8sapicorev1ISCSIVolumeSource represents the Iok8sapicorev1ISCSIVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ISCSIVolumeSource struct {
	Initiatorname string `json:"initiatorName,omitempty"` // initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
	Portals []string `json:"portals,omitempty"` // portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Targetportal string `json:"targetPortal"` // targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Chapauthsession bool `json:"chapAuthSession,omitempty"` // chapAuthSession defines whether support iSCSI Session CHAP authentication
	Iqn string `json:"iqn"` // iqn is the target iSCSI Qualified Name.
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	Iscsiinterface string `json:"iscsiInterface,omitempty"` // iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
	Lun int `json:"lun"` // lun represents iSCSI Target Lun number.
	Readonly bool `json:"readOnly,omitempty"` // readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	Chapauthdiscovery bool `json:"chapAuthDiscovery,omitempty"` // chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication
}

// Iok8sapirbacv1AggregationRule represents the Iok8sapirbacv1AggregationRule schema from the OpenAPI specification
type Iok8sapirbacv1AggregationRule struct {
	Clusterroleselectors []Iok8sapimachinerypkgapismetav1LabelSelector `json:"clusterRoleSelectors,omitempty"` // ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
}

// Iok8sapicorev1Toleration represents the Iok8sapicorev1Toleration schema from the OpenAPI specification
type Iok8sapicorev1Toleration struct {
	Tolerationseconds int64 `json:"tolerationSeconds,omitempty"` // TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
	Value string `json:"value,omitempty"` // Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.
	Effect string `json:"effect,omitempty"` // Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
	Key string `json:"key,omitempty"` // Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
	Operator string `json:"operator,omitempty"` // Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Jsonpath string `json:"jsonPath"` // jsonPath is a simple JSON path which is evaluated against each custom resource to produce a field selector value. Only JSON paths without the array notation are allowed. Must point to a field of type string, boolean or integer. Types with enum values and strings with formats are allowed. If jsonPath refers to absent field in a resource, the jsonPath evaluates to an empty string. Must not point to metdata fields. Required.
}

// Iok8sapibatchv1JobList represents the Iok8sapibatchv1JobList schema from the OpenAPI specification
type Iok8sapibatchv1JobList struct {
	Items []Iok8sapibatchv1Job `json:"items"` // items is the list of Jobs.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapiresourcev1beta2ResourcePool represents the Iok8sapiresourcev1beta2ResourcePool schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ResourcePool struct {
	Generation int64 `json:"generation"` // Generation tracks the change in a pool over time. Whenever a driver changes something about one or more of the resources in a pool, it must change the generation in all ResourceSlices which are part of that pool. Consumers of ResourceSlices should only consider resources from the pool with the highest generation number. The generation may be reset by drivers, which should be fine for consumers, assuming that all ResourceSlices in a pool are updated to match or deleted. Combined with ResourceSliceCount, this mechanism enables consumers to detect pools which are comprised of multiple ResourceSlices and are in an incomplete state.
	Name string `json:"name"` // Name is used to identify the pool. For node-local devices, this is often the node name, but this is not required. It must not be longer than 253 characters and must consist of one or more DNS sub-domains separated by slashes. This field is immutable.
	Resourceslicecount int64 `json:"resourceSliceCount"` // ResourceSliceCount is the total number of ResourceSlices in the pool at this generation number. Must be greater than zero. Consumers can use this to check whether they have seen all ResourceSlices belonging to the same pool.
}

// Iok8sapicertificatesv1alpha1ClusterTrustBundleSpec represents the Iok8sapicertificatesv1alpha1ClusterTrustBundleSpec schema from the OpenAPI specification
type Iok8sapicertificatesv1alpha1ClusterTrustBundleSpec struct {
	Signername string `json:"signerName,omitempty"` // signerName indicates the associated signer, if any. In order to create or update a ClusterTrustBundle that sets signerName, you must have the following cluster-scoped permission: group=certificates.k8s.io resource=signers resourceName=<the signer name> verb=attest. If signerName is not empty, then the ClusterTrustBundle object must be named with the signer name as a prefix (translating slashes to colons). For example, for the signer name `example.com/foo`, valid ClusterTrustBundle object names include `example.com:foo:abc` and `example.com:foo:v1`. If signerName is empty, then the ClusterTrustBundle object's name must not have such a prefix. List/watch requests for ClusterTrustBundles can filter on this field using a `spec.signerName=NAME` field selector.
	Trustbundle string `json:"trustBundle"` // trustBundle contains the individual X.509 trust anchors for this bundle, as PEM bundle of PEM-wrapped, DER-formatted X.509 certificates. The data must consist only of PEM certificate blocks that parse as valid X.509 certificates. Each certificate must include a basic constraints extension with the CA bit set. The API server will reject objects that contain duplicate certificates, or that use PEM block headers. Users of ClusterTrustBundles, including Kubelet, are free to reorder and deduplicate certificate blocks in this file according to their own logic, as well as to drop PEM block headers and inter-block data.
}

// Iok8sapicorev1ReplicationControllerCondition represents the Iok8sapicorev1ReplicationControllerCondition schema from the OpenAPI specification
type Iok8sapicorev1ReplicationControllerCondition struct {
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of replication controller condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
}

// Iok8sapiappsv1DeploymentCondition represents the Iok8sapiappsv1DeploymentCondition schema from the OpenAPI specification
type Iok8sapiappsv1DeploymentCondition struct {
	Lastupdatetime string `json:"lastUpdateTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of deployment condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
}

// Iok8sapiresourcev1beta1ResourceSliceList represents the Iok8sapiresourcev1beta1ResourceSliceList schema from the OpenAPI specification
type Iok8sapiresourcev1beta1ResourceSliceList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiresourcev1beta1ResourceSlice `json:"items"` // Items is the list of resource ResourceSlices.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicoordinationv1LeaseSpec represents the Iok8sapicoordinationv1LeaseSpec schema from the OpenAPI specification
type Iok8sapicoordinationv1LeaseSpec struct {
	Renewtime string `json:"renewTime,omitempty"` // MicroTime is version of Time with microsecond level precision.
	Strategy string `json:"strategy,omitempty"` // Strategy indicates the strategy for picking the leader for coordinated leader election. If the field is not specified, there is no active coordination for this lease. (Alpha) Using this field requires the CoordinatedLeaderElection feature gate to be enabled.
	Acquiretime string `json:"acquireTime,omitempty"` // MicroTime is version of Time with microsecond level precision.
	Holderidentity string `json:"holderIdentity,omitempty"` // holderIdentity contains the identity of the holder of a current lease. If Coordinated Leader Election is used, the holder identity must be equal to the elected LeaseCandidate.metadata.name field.
	Leasedurationseconds int `json:"leaseDurationSeconds,omitempty"` // leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measured against the time of last observed renewTime.
	Leasetransitions int `json:"leaseTransitions,omitempty"` // leaseTransitions is the number of transitions of a lease between holders.
	Preferredholder string `json:"preferredHolder,omitempty"` // PreferredHolder signals to a lease holder that the lease has a more optimal holder and should be given up. This field can only be set if Strategy is also set.
}

// Iok8sapinodev1Scheduling represents the Iok8sapinodev1Scheduling schema from the OpenAPI specification
type Iok8sapinodev1Scheduling struct {
	Nodeselector map[string]interface{} `json:"nodeSelector,omitempty"` // nodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing nodeSelector. Any conflicts will cause the pod to be rejected in admission.
	Tolerations []Iok8sapicorev1Toleration `json:"tolerations,omitempty"` // tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass.
}

// Iok8sapiapiserverinternalv1alpha1StorageVersionList represents the Iok8sapiapiserverinternalv1alpha1StorageVersionList schema from the OpenAPI specification
type Iok8sapiapiserverinternalv1alpha1StorageVersionList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiapiserverinternalv1alpha1StorageVersion `json:"items"` // Items holds a list of StorageVersion
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiautoscalingv1HorizontalPodAutoscalerList represents the Iok8sapiautoscalingv1HorizontalPodAutoscalerList schema from the OpenAPI specification
type Iok8sapiautoscalingv1HorizontalPodAutoscalerList struct {
	Items []Iok8sapiautoscalingv1HorizontalPodAutoscaler `json:"items"` // items is the list of horizontal pod autoscaler objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapimachinerypkgapismetav1APIResourceList represents the Iok8sapimachinerypkgapismetav1APIResourceList schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1APIResourceList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Resources []Iok8sapimachinerypkgapismetav1APIResource `json:"resources"` // resources contains the name of the resources and if they are namespaced.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Groupversion string `json:"groupVersion"` // groupVersion is the group and version this APIResourceList is for.
}

// Iok8sapibatchv1CronJob represents the Iok8sapibatchv1CronJob schema from the OpenAPI specification
type Iok8sapibatchv1CronJob struct {
	Spec Iok8sapibatchv1CronJobSpec `json:"spec,omitempty"` // CronJobSpec describes how the job execution will look like and when it will actually run.
	Status Iok8sapibatchv1CronJobStatus `json:"status,omitempty"` // CronJobStatus represents the current state of a cron job.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1PodSecurityContext represents the Iok8sapicorev1PodSecurityContext schema from the OpenAPI specification
type Iok8sapicorev1PodSecurityContext struct {
	Windowsoptions Iok8sapicorev1WindowsSecurityContextOptions `json:"windowsOptions,omitempty"` // WindowsSecurityContextOptions contain Windows-specific options and credentials.
	Fsgroup int64 `json:"fsGroup,omitempty"` // A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw---- If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.
	Supplementalgroups []int64 `json:"supplementalGroups,omitempty"` // A list of groups applied to the first process run in each container, in addition to the container's primary GID and fsGroup (if specified). If the SupplementalGroupsPolicy feature is enabled, the supplementalGroupsPolicy field determines whether these are in addition to or instead of any group memberships defined in the container image. If unspecified, no additional groups are added, though group memberships defined in the container image may still be used, depending on the supplementalGroupsPolicy field. Note that this field cannot be set when spec.os.name is windows.
	Apparmorprofile Iok8sapicorev1AppArmorProfile `json:"appArmorProfile,omitempty"` // AppArmorProfile defines a pod or container's AppArmor settings.
	Fsgroupchangepolicy string `json:"fsGroupChangePolicy,omitempty"` // fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are "OnRootMismatch" and "Always". If not specified, "Always" is used. Note that this field cannot be set when spec.os.name is windows.
	Runasnonroot bool `json:"runAsNonRoot,omitempty"` // Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	Selinuxoptions Iok8sapicorev1SELinuxOptions `json:"seLinuxOptions,omitempty"` // SELinuxOptions are the labels to be applied to the container
	Runasgroup int64 `json:"runAsGroup,omitempty"` // The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.
	Runasuser int64 `json:"runAsUser,omitempty"` // The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.
	Selinuxchangepolicy string `json:"seLinuxChangePolicy,omitempty"` // seLinuxChangePolicy defines how the container's SELinux label is applied to all volumes used by the Pod. It has no effect on nodes that do not support SELinux or to volumes does not support SELinux. Valid values are "MountOption" and "Recursive". "Recursive" means relabeling of all files on all Pod volumes by the container runtime. This may be slow for large volumes, but allows mixing privileged and unprivileged Pods sharing the same volume on the same node. "MountOption" mounts all eligible Pod volumes with `-o context` mount option. This requires all Pods that share the same volume to use the same SELinux label. It is not possible to share the same volume among privileged and unprivileged Pods. Eligible volumes are in-tree FibreChannel and iSCSI volumes, and all CSI volumes whose CSI driver announces SELinux support by setting spec.seLinuxMount: true in their CSIDriver instance. Other volumes are always re-labelled recursively. "MountOption" value is allowed only when SELinuxMount feature gate is enabled. If not specified and SELinuxMount feature gate is enabled, "MountOption" is used. If not specified and SELinuxMount feature gate is disabled, "MountOption" is used for ReadWriteOncePod volumes and "Recursive" for all other volumes. This field affects only Pods that have SELinux label set, either in PodSecurityContext or in SecurityContext of all containers. All Pods that use the same volume should use the same seLinuxChangePolicy, otherwise some pods can get stuck in ContainerCreating state. Note that this field cannot be set when spec.os.name is windows.
	Seccompprofile Iok8sapicorev1SeccompProfile `json:"seccompProfile,omitempty"` // SeccompProfile defines a pod/container's seccomp profile settings. Only one profile source may be set.
	Supplementalgroupspolicy string `json:"supplementalGroupsPolicy,omitempty"` // Defines how supplemental groups of the first container processes are calculated. Valid values are "Merge" and "Strict". If not specified, "Merge" is used. (Alpha) Using the field requires the SupplementalGroupsPolicy feature gate to be enabled and the container runtime must implement support for this feature. Note that this field cannot be set when spec.os.name is windows.
	Sysctls []Iok8sapicorev1Sysctl `json:"sysctls,omitempty"` // Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.
}

// Iok8sapinetworkingv1beta1ServiceCIDR represents the Iok8sapinetworkingv1beta1ServiceCIDR schema from the OpenAPI specification
type Iok8sapinetworkingv1beta1ServiceCIDR struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapinetworkingv1beta1ServiceCIDRSpec `json:"spec,omitempty"` // ServiceCIDRSpec define the CIDRs the user wants to use for allocating ClusterIPs for Services.
	Status Iok8sapinetworkingv1beta1ServiceCIDRStatus `json:"status,omitempty"` // ServiceCIDRStatus describes the current state of the ServiceCIDR.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1ConfigMapNodeConfigSource represents the Iok8sapicorev1ConfigMapNodeConfigSource schema from the OpenAPI specification
type Iok8sapicorev1ConfigMapNodeConfigSource struct {
	Name string `json:"name"` // Name is the metadata.name of the referenced ConfigMap. This field is required in all cases.
	Namespace string `json:"namespace"` // Namespace is the metadata.namespace of the referenced ConfigMap. This field is required in all cases.
	Resourceversion string `json:"resourceVersion,omitempty"` // ResourceVersion is the metadata.ResourceVersion of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status.
	Uid string `json:"uid,omitempty"` // UID is the metadata.UID of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status.
	Kubeletconfigkey string `json:"kubeletConfigKey"` // KubeletConfigKey declares which key of the referenced ConfigMap corresponds to the KubeletConfiguration structure This field is required in all cases.
}

// Iok8sapiresourcev1beta1ResourceClaimTemplate represents the Iok8sapiresourcev1beta1ResourceClaimTemplate schema from the OpenAPI specification
type Iok8sapiresourcev1beta1ResourceClaimTemplate struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiresourcev1beta1ResourceClaimTemplateSpec `json:"spec"` // ResourceClaimTemplateSpec contains the metadata and fields for a ResourceClaim.
}

// Iok8sapicorev1ContainerState represents the Iok8sapicorev1ContainerState schema from the OpenAPI specification
type Iok8sapicorev1ContainerState struct {
	Terminated Iok8sapicorev1ContainerStateTerminated `json:"terminated,omitempty"` // ContainerStateTerminated is a terminated state of a container.
	Waiting Iok8sapicorev1ContainerStateWaiting `json:"waiting,omitempty"` // ContainerStateWaiting is a waiting state of a container.
	Running Iok8sapicorev1ContainerStateRunning `json:"running,omitempty"` // ContainerStateRunning is a running state of a container.
}

// Iok8sapicorev1TypedLocalObjectReference represents the Iok8sapicorev1TypedLocalObjectReference schema from the OpenAPI specification
type Iok8sapicorev1TypedLocalObjectReference struct {
	Apigroup string `json:"apiGroup,omitempty"` // APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
	Kind string `json:"kind"` // Kind is the type of resource being referenced
	Name string `json:"name"` // Name is the name of resource being referenced
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Items []GeneratedType `json:"items"` // items list individual CustomResourceDefinition objects
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Scope string `json:"scope"` // scope indicates whether the defined custom resource is cluster- or namespace-scoped. Allowed values are `Cluster` and `Namespaced`.
	Versions []GeneratedType `json:"versions"` // versions is the list of all API versions of the defined custom resource. Version names are used to compute the order in which served versions are listed in API discovery. If the version string is "kube-like", it will sort above non "kube-like" version strings, which are ordered lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the major version), then optionally the string "alpha" or "beta" and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.
	Conversion GeneratedType `json:"conversion,omitempty"` // CustomResourceConversion describes how to convert different versions of a CR.
	Group string `json:"group"` // group is the API group of the defined custom resource. The custom resources are served under `/apis/<group>/...`. Must match the name of the CustomResourceDefinition (in the form `<names.plural>.<group>`).
	Names GeneratedType `json:"names"` // CustomResourceDefinitionNames indicates the names to serve this CustomResourceDefinition
	Preserveunknownfields bool `json:"preserveUnknownFields,omitempty"` // preserveUnknownFields indicates that object fields which are not specified in the OpenAPI schema should be preserved when persisting to storage. apiVersion, kind, metadata and known fields inside metadata are always preserved. This field is deprecated in favor of setting `x-preserve-unknown-fields` to true in `spec.versions[*].schema.openAPIV3Schema`. See https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#field-pruning for details.
}

// Iok8sapicorev1PodResourceClaimStatus represents the Iok8sapicorev1PodResourceClaimStatus schema from the OpenAPI specification
type Iok8sapicorev1PodResourceClaimStatus struct {
	Name string `json:"name"` // Name uniquely identifies this resource claim inside the pod. This must match the name of an entry in pod.spec.resourceClaims, which implies that the string must be a DNS_LABEL.
	Resourceclaimname string `json:"resourceClaimName,omitempty"` // ResourceClaimName is the name of the ResourceClaim that was generated for the Pod in the namespace of the Pod. If this is unset, then generating a ResourceClaim was not necessary. The pod.spec.resourceClaims entry can be ignored in this case.
}

// Iok8sapimachinerypkgapismetav1APIGroup represents the Iok8sapimachinerypkgapismetav1APIGroup schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1APIGroup struct {
	Serveraddressbyclientcidrs []Iok8sapimachinerypkgapismetav1ServerAddressByClientCIDR `json:"serverAddressByClientCIDRs,omitempty"` // a map of client CIDR to server address that is serving this group. This is to help clients reach servers in the most network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case of multiple matches, clients should use the longest matching CIDR. The server returns only those CIDRs that it thinks that the client can match. For example: the master will return an internal IP CIDR only, if the client reaches the server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that order) to get the client IP.
	Versions []Iok8sapimachinerypkgapismetav1GroupVersionForDiscovery `json:"versions"` // versions are the versions supported in this group.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Name string `json:"name"` // name is the name of the group.
	Preferredversion Iok8sapimachinerypkgapismetav1GroupVersionForDiscovery `json:"preferredVersion,omitempty"` // GroupVersion contains the "group/version" and "version" string of a version. It is made a struct to keep extensibility.
}

// Iok8sapistoragemigrationv1alpha1StorageVersionMigration represents the Iok8sapistoragemigrationv1alpha1StorageVersionMigration schema from the OpenAPI specification
type Iok8sapistoragemigrationv1alpha1StorageVersionMigration struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapistoragemigrationv1alpha1StorageVersionMigrationSpec `json:"spec,omitempty"` // Spec of the storage version migration.
	Status Iok8sapistoragemigrationv1alpha1StorageVersionMigrationStatus `json:"status,omitempty"` // Status of the storage version migration.
}

// Iok8sapinetworkingv1beta1ServiceCIDRStatus represents the Iok8sapinetworkingv1beta1ServiceCIDRStatus schema from the OpenAPI specification
type Iok8sapinetworkingv1beta1ServiceCIDRStatus struct {
	Conditions []Iok8sapimachinerypkgapismetav1Condition `json:"conditions,omitempty"` // conditions holds an array of metav1.Condition that describe the state of the ServiceCIDR. Current service state
}

// Iok8sapistoragemigrationv1alpha1StorageVersionMigrationList represents the Iok8sapistoragemigrationv1alpha1StorageVersionMigrationList schema from the OpenAPI specification
type Iok8sapistoragemigrationv1alpha1StorageVersionMigrationList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapistoragemigrationv1alpha1StorageVersionMigration `json:"items"` // Items is the list of StorageVersionMigration
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1EphemeralVolumeSource represents the Iok8sapicorev1EphemeralVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1EphemeralVolumeSource struct {
	Volumeclaimtemplate Iok8sapicorev1PersistentVolumeClaimTemplate `json:"volumeClaimTemplate,omitempty"` // PersistentVolumeClaimTemplate is used to produce PersistentVolumeClaim objects as part of an EphemeralVolumeSource.
}

// Iok8sapistoragev1CSIDriverList represents the Iok8sapistoragev1CSIDriverList schema from the OpenAPI specification
type Iok8sapistoragev1CSIDriverList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapistoragev1CSIDriver `json:"items"` // items is the list of CSIDriver
}

// Iok8sapiresourcev1beta2NetworkDeviceData represents the Iok8sapiresourcev1beta2NetworkDeviceData schema from the OpenAPI specification
type Iok8sapiresourcev1beta2NetworkDeviceData struct {
	Interfacename string `json:"interfaceName,omitempty"` // InterfaceName specifies the name of the network interface associated with the allocated device. This might be the name of a physical or virtual network interface being configured in the pod. Must not be longer than 256 characters.
	Ips []string `json:"ips,omitempty"` // IPs lists the network addresses assigned to the device's network interface. This can include both IPv4 and IPv6 addresses. The IPs are in the CIDR notation, which includes both the address and the associated subnet mask. e.g.: "192.0.2.5/24" for IPv4 and "2001:db8::5/64" for IPv6.
	Hardwareaddress string `json:"hardwareAddress,omitempty"` // HardwareAddress represents the hardware address (e.g. MAC Address) of the device's network interface. Must not be longer than 128 characters.
}

// Iok8sapibatchv1PodFailurePolicyOnExitCodesRequirement represents the Iok8sapibatchv1PodFailurePolicyOnExitCodesRequirement schema from the OpenAPI specification
type Iok8sapibatchv1PodFailurePolicyOnExitCodesRequirement struct {
	Containername string `json:"containerName,omitempty"` // Restricts the check for exit codes to the container with the specified name. When null, the rule applies to all containers. When specified, it should match one the container or initContainer names in the pod template.
	Operator string `json:"operator"` // Represents the relationship between the container exit code(s) and the specified values. Containers completed with success (exit code 0) are excluded from the requirement check. Possible values are: - In: the requirement is satisfied if at least one container exit code (might be multiple if there are multiple containers not restricted by the 'containerName' field) is in the set of specified values. - NotIn: the requirement is satisfied if at least one container exit code (might be multiple if there are multiple containers not restricted by the 'containerName' field) is not in the set of specified values. Additional values are considered to be added in the future. Clients should react to an unknown operator by assuming the requirement is not satisfied.
	Values []int `json:"values"` // Specifies the set of values. Each returned container exit code (might be multiple in case of multiple containers) is checked against this set of values with respect to the operator. The list of values must be ordered and must not contain duplicates. Value '0' cannot be used for the In operator. At least one element is required. At most 255 elements are allowed.
}

// Iok8sapiautoscalingv1Scale represents the Iok8sapiautoscalingv1Scale schema from the OpenAPI specification
type Iok8sapiautoscalingv1Scale struct {
	Spec Iok8sapiautoscalingv1ScaleSpec `json:"spec,omitempty"` // ScaleSpec describes the attributes of a scale subresource.
	Status Iok8sapiautoscalingv1ScaleStatus `json:"status,omitempty"` // ScaleStatus represents the current status of a scale subresource.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapistoragev1VolumeNodeResources represents the Iok8sapistoragev1VolumeNodeResources schema from the OpenAPI specification
type Iok8sapistoragev1VolumeNodeResources struct {
	Count int `json:"count,omitempty"` // count indicates the maximum number of unique volumes managed by the CSI driver that can be used on a node. A volume that is both attached and mounted on a node is considered to be used once, not twice. The same rule applies for a unique volume that is shared among multiple pods on the same node. If this field is not specified, then the supported number of volumes on this node is unbounded.
}

// Iok8sapiflowcontrolv1ResourcePolicyRule represents the Iok8sapiflowcontrolv1ResourcePolicyRule schema from the OpenAPI specification
type Iok8sapiflowcontrolv1ResourcePolicyRule struct {
	Apigroups []string `json:"apiGroups"` // `apiGroups` is a list of matching API groups and may not be empty. "*" matches all API groups and, if present, must be the only entry. Required.
	Clusterscope bool `json:"clusterScope,omitempty"` // `clusterScope` indicates whether to match requests that do not specify a namespace (which happens either because the resource is not namespaced or the request targets all namespaces). If this field is omitted or false then the `namespaces` field must contain a non-empty list.
	Namespaces []string `json:"namespaces,omitempty"` // `namespaces` is a list of target namespaces that restricts matches. A request that specifies a target namespace matches only if either (a) this list contains that target namespace or (b) this list contains "*". Note that "*" matches any specified namespace but does not match a request that _does not specify_ a namespace (see the `clusterScope` field for that). This list may be empty, but only if `clusterScope` is true.
	Resources []string `json:"resources"` // `resources` is a list of matching resources (i.e., lowercase and plural) with, if desired, subresource. For example, [ "services", "nodes/status" ]. This list may not be empty. "*" matches all resources and, if present, must be the only entry. Required.
	Verbs []string `json:"verbs"` // `verbs` is a list of matching verbs and may not be empty. "*" matches all verbs and, if present, must be the only entry. Required.
}

// Iok8sapiautoscalingv2HPAScalingRules represents the Iok8sapiautoscalingv2HPAScalingRules schema from the OpenAPI specification
type Iok8sapiautoscalingv2HPAScalingRules struct {
	Stabilizationwindowseconds int `json:"stabilizationWindowSeconds,omitempty"` // stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).
	Tolerance string `json:"tolerance,omitempty"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
	Policies []Iok8sapiautoscalingv2HPAScalingPolicy `json:"policies,omitempty"` // policies is a list of potential scaling polices which can be used during scaling. If not set, use the default values: - For scale up: allow doubling the number of pods, or an absolute change of 4 pods in a 15s window. - For scale down: allow all pods to be removed in a 15s window.
	Selectpolicy string `json:"selectPolicy,omitempty"` // selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.
}

// Iok8sapiadmissionregistrationv1RuleWithOperations represents the Iok8sapiadmissionregistrationv1RuleWithOperations schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1RuleWithOperations struct {
	Apiversions []string `json:"apiVersions,omitempty"` // APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.
	Operations []string `json:"operations,omitempty"` // Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If '*' is present, the length of the slice must be one. Required.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources. If wildcard is present, the validation rule will ensure resources do not overlap with each other. Depending on the enclosing object, subresources might not be allowed. Required.
	Scope string `json:"scope,omitempty"` // scope specifies the scope of this rule. Valid values are "Cluster", "Namespaced", and "*" "Cluster" means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. "Namespaced" means that only namespaced resources will match this rule. "*" means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is "*".
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.
}

// Iok8sapipolicyv1PodDisruptionBudget represents the Iok8sapipolicyv1PodDisruptionBudget schema from the OpenAPI specification
type Iok8sapipolicyv1PodDisruptionBudget struct {
	Spec Iok8sapipolicyv1PodDisruptionBudgetSpec `json:"spec,omitempty"` // PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
	Status Iok8sapipolicyv1PodDisruptionBudgetStatus `json:"status,omitempty"` // PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiauthenticationv1TokenReviewSpec represents the Iok8sapiauthenticationv1TokenReviewSpec schema from the OpenAPI specification
type Iok8sapiauthenticationv1TokenReviewSpec struct {
	Audiences []string `json:"audiences,omitempty"` // Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
	Token string `json:"token,omitempty"` // Token is the opaque bearer token.
}

// Iok8sapiresourcev1beta1DeviceClassSpec represents the Iok8sapiresourcev1beta1DeviceClassSpec schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceClassSpec struct {
	Config []Iok8sapiresourcev1beta1DeviceClassConfiguration `json:"config,omitempty"` // Config defines configuration parameters that apply to each device that is claimed via this class. Some classses may potentially be satisfied by multiple drivers, so each instance of a vendor configuration applies to exactly one driver. They are passed to the driver, but are not considered while allocating the claim.
	Selectors []Iok8sapiresourcev1beta1DeviceSelector `json:"selectors,omitempty"` // Each selector must be satisfied by a device which is claimed via this class.
}

// Iok8sapicorev1PodTemplateList represents the Iok8sapicorev1PodTemplateList schema from the OpenAPI specification
type Iok8sapicorev1PodTemplateList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1PodTemplate `json:"items"` // List of pod templates
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1WeightedPodAffinityTerm represents the Iok8sapicorev1WeightedPodAffinityTerm schema from the OpenAPI specification
type Iok8sapicorev1WeightedPodAffinityTerm struct {
	Podaffinityterm Iok8sapicorev1PodAffinityTerm `json:"podAffinityTerm"` // Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose value of the label with key <topologyKey> matches that of any node on which a pod of the set of pods is running
	Weight int `json:"weight"` // weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
}

// Iok8sapicorev1Probe represents the Iok8sapicorev1Probe schema from the OpenAPI specification
type Iok8sapicorev1Probe struct {
	Tcpsocket Iok8sapicorev1TCPSocketAction `json:"tcpSocket,omitempty"` // TCPSocketAction describes an action based on opening a socket
	Timeoutseconds int `json:"timeoutSeconds,omitempty"` // Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	Httpget Iok8sapicorev1HTTPGetAction `json:"httpGet,omitempty"` // HTTPGetAction describes an action based on HTTP Get requests.
	Initialdelayseconds int `json:"initialDelaySeconds,omitempty"` // Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	Exec Iok8sapicorev1ExecAction `json:"exec,omitempty"` // ExecAction describes a "run in container" action.
	Failurethreshold int `json:"failureThreshold,omitempty"` // Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
	Grpc Iok8sapicorev1GRPCAction `json:"grpc,omitempty"` // GRPCAction specifies an action involving a GRPC service.
	Periodseconds int `json:"periodSeconds,omitempty"` // How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
	Terminationgraceperiodseconds int64 `json:"terminationGracePeriodSeconds,omitempty"` // Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.
	Successthreshold int `json:"successThreshold,omitempty"` // Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // name is the version name, e.g. “v1”, “v2beta1”, etc. The custom resources are served under this version at `/apis/<group>/<version>/...` if `served` is true.
	Additionalprintercolumns []GeneratedType `json:"additionalPrinterColumns,omitempty"` // additionalPrinterColumns specifies additional columns returned in Table output. See https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables for details. If no columns are specified, a single column displaying the age of the custom resource is used.
	Schema GeneratedType `json:"schema,omitempty"` // CustomResourceValidation is a list of validation methods for CustomResources.
	Served bool `json:"served"` // served is a flag enabling/disabling this version from being served via REST APIs
	Subresources GeneratedType `json:"subresources,omitempty"` // CustomResourceSubresources defines the status and scale subresources for CustomResources.
	Deprecated bool `json:"deprecated,omitempty"` // deprecated indicates this version of the custom resource API is deprecated. When set to true, API requests to this version receive a warning header in the server response. Defaults to false.
	Storage bool `json:"storage"` // storage indicates this version should be used when persisting custom resources to storage. There must be exactly one version with storage=true.
	Selectablefields []GeneratedType `json:"selectableFields,omitempty"` // selectableFields specifies paths to fields that may be used as field selectors. A maximum of 8 selectable fields are allowed. See https://kubernetes.io/docs/concepts/overview/working-with-objects/field-selectors
	Deprecationwarning string `json:"deprecationWarning,omitempty"` // deprecationWarning overrides the default warning returned to API clients. May only be set when `deprecated` is true. The default warning indicates this version is deprecated and recommends use of the newest served version of equal or greater stability, if one exists.
}

// Iok8sapicorev1GCEPersistentDiskVolumeSource represents the Iok8sapicorev1GCEPersistentDiskVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1GCEPersistentDiskVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	Partition int `json:"partition,omitempty"` // partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	Pdname string `json:"pdName"` // pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	Readonly bool `json:"readOnly,omitempty"` // readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
}

// Iok8sapicorev1ServiceList represents the Iok8sapicorev1ServiceList schema from the OpenAPI specification
type Iok8sapicorev1ServiceList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1Service `json:"items"` // List of services
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapistoragemigrationv1alpha1GroupVersionResource represents the Iok8sapistoragemigrationv1alpha1GroupVersionResource schema from the OpenAPI specification
type Iok8sapistoragemigrationv1alpha1GroupVersionResource struct {
	Resource string `json:"resource,omitempty"` // The name of the resource.
	Version string `json:"version,omitempty"` // The name of the version.
	Group string `json:"group,omitempty"` // The name of the group.
}

// Iok8sapirbacv1PolicyRule represents the Iok8sapirbacv1PolicyRule schema from the OpenAPI specification
type Iok8sapirbacv1PolicyRule struct {
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. '*' represents all resources.
	Verbs []string `json:"verbs"` // Verbs is a list of Verbs that apply to ALL the ResourceKinds contained in this rule. '*' represents all verbs.
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the name of the APIGroup that contains the resources. If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed. "" represents the core API group and "*" represents all API groups.
	Nonresourceurls []string `json:"nonResourceURLs,omitempty"` // NonResourceURLs is a set of partial urls that a user should have access to. *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"), but not both.
	Resourcenames []string `json:"resourceNames,omitempty"` // ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.
}

// Iok8sapiresourcev1beta2DeviceTaint represents the Iok8sapiresourcev1beta2DeviceTaint schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceTaint struct {
	Key string `json:"key"` // The taint key to be applied to a device. Must be a label name.
	Timeadded string `json:"timeAdded,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Value string `json:"value,omitempty"` // The taint value corresponding to the taint key. Must be a label value.
	Effect string `json:"effect"` // The effect of the taint on claims that do not tolerate the taint and through such claims on the pods using them. Valid effects are NoSchedule and NoExecute. PreferNoSchedule as used for nodes is not valid here.
}

// Iok8sapicorev1VolumeProjection represents the Iok8sapicorev1VolumeProjection schema from the OpenAPI specification
type Iok8sapicorev1VolumeProjection struct {
	Configmap Iok8sapicorev1ConfigMapProjection `json:"configMap,omitempty"` // Adapts a ConfigMap into a projected volume. The contents of the target ConfigMap's Data field will be presented in a projected volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. Note that this is identical to a configmap volume source without the default mode.
	Downwardapi Iok8sapicorev1DownwardAPIProjection `json:"downwardAPI,omitempty"` // Represents downward API info for projecting into a projected volume. Note that this is identical to a downwardAPI volume source without the default mode.
	Secret Iok8sapicorev1SecretProjection `json:"secret,omitempty"` // Adapts a secret into a projected volume. The contents of the target Secret's Data field will be presented in a projected volume as files using the keys in the Data field as the file names. Note that this is identical to a secret volume source without the default mode.
	Serviceaccounttoken Iok8sapicorev1ServiceAccountTokenProjection `json:"serviceAccountToken,omitempty"` // ServiceAccountTokenProjection represents a projected service account token volume. This projection can be used to insert a service account token into the pods runtime filesystem for use against APIs (Kubernetes API Server or otherwise).
	Clustertrustbundle Iok8sapicorev1ClusterTrustBundleProjection `json:"clusterTrustBundle,omitempty"` // ClusterTrustBundleProjection describes how to select a set of ClusterTrustBundle objects and project their contents into the pod filesystem.
}

// Iok8sapibatchv1Job represents the Iok8sapibatchv1Job schema from the OpenAPI specification
type Iok8sapibatchv1Job struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapibatchv1JobSpec `json:"spec,omitempty"` // JobSpec describes how the job execution will look like.
	Status Iok8sapibatchv1JobStatus `json:"status,omitempty"` // JobStatus represents the current state of a Job.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1ContainerStateRunning represents the Iok8sapicorev1ContainerStateRunning schema from the OpenAPI specification
type Iok8sapicorev1ContainerStateRunning struct {
	Startedat string `json:"startedAt,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
}

// Iok8sapinetworkingv1NetworkPolicyList represents the Iok8sapinetworkingv1NetworkPolicyList schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicyList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapinetworkingv1NetworkPolicy `json:"items"` // items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapicorev1AWSElasticBlockStoreVolumeSource represents the Iok8sapicorev1AWSElasticBlockStoreVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1AWSElasticBlockStoreVolumeSource struct {
	Partition int `json:"partition,omitempty"` // partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	Readonly bool `json:"readOnly,omitempty"` // readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	Volumeid string `json:"volumeID"` // volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
}

// Iok8sapimachinerypkgapismetav1OwnerReference represents the Iok8sapimachinerypkgapismetav1OwnerReference schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1OwnerReference struct {
	Controller bool `json:"controller,omitempty"` // If true, this reference points to the managing controller.
	Kind string `json:"kind"` // Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Name string `json:"name"` // Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names
	Uid string `json:"uid"` // UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids
	Apiversion string `json:"apiVersion"` // API version of the referent.
	Blockownerdeletion bool `json:"blockOwnerDeletion,omitempty"` // If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion for how the garbage collector interacts with this field and enforces the foreground deletion. Defaults to false. To set this field, a user needs "delete" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.
}

// Iok8sapistoragev1beta1VolumeAttributesClass represents the Iok8sapistoragev1beta1VolumeAttributesClass schema from the OpenAPI specification
type Iok8sapistoragev1beta1VolumeAttributesClass struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Parameters map[string]interface{} `json:"parameters,omitempty"` // parameters hold volume attributes defined by the CSI driver. These values are opaque to the Kubernetes and are passed directly to the CSI driver. The underlying storage provider supports changing these attributes on an existing volume, however the parameters field itself is immutable. To invoke a volume update, a new VolumeAttributesClass should be created with new parameters, and the PersistentVolumeClaim should be updated to reference the new VolumeAttributesClass. This field is required and must contain at least one key/value pair. The keys cannot be empty, and the maximum number of parameters is 512, with a cumulative max size of 256K. If the CSI driver rejects invalid parameters, the target PersistentVolumeClaim will be set to an "Infeasible" state in the modifyVolumeStatus field.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Drivername string `json:"driverName"` // Name of the CSI driver This field is immutable.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapiresourcev1beta2OpaqueDeviceConfiguration represents the Iok8sapiresourcev1beta2OpaqueDeviceConfiguration schema from the OpenAPI specification
type Iok8sapiresourcev1beta2OpaqueDeviceConfiguration struct {
	Driver string `json:"driver"` // Driver is used to determine which kubelet plugin needs to be passed these configuration parameters. An admission policy provided by the driver developer could use this to decide whether it needs to validate them. Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver.
	Parameters Iok8sapimachinerypkgruntimeRawExtension `json:"parameters"` // RawExtension is used to hold extensions in external versions. To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types. // Internal package: 	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.Object `json:"myPlugin"` 	} 	type PluginA struct { 		AOption string `json:"aOption"` 	} // External package: 	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.RawExtension `json:"myPlugin"` 	} 	type PluginA struct { 		AOption string `json:"aOption"` 	} // On the wire, the JSON will look something like this: 	{ 		"kind":"MyAPIObject", 		"apiVersion":"v1", 		"myPlugin": { 			"kind":"PluginA", 			"aOption":"foo", 		}, 	} So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)
}

// Iok8sapinetworkingv1HTTPIngressRuleValue represents the Iok8sapinetworkingv1HTTPIngressRuleValue schema from the OpenAPI specification
type Iok8sapinetworkingv1HTTPIngressRuleValue struct {
	Paths []Iok8sapinetworkingv1HTTPIngressPath `json:"paths"` // paths is a collection of paths that map requests to backends.
}

// Iok8sapicorev1PodResourceClaim represents the Iok8sapicorev1PodResourceClaim schema from the OpenAPI specification
type Iok8sapicorev1PodResourceClaim struct {
	Name string `json:"name"` // Name uniquely identifies this resource claim inside the pod. This must be a DNS_LABEL.
	Resourceclaimname string `json:"resourceClaimName,omitempty"` // ResourceClaimName is the name of a ResourceClaim object in the same namespace as this pod. Exactly one of ResourceClaimName and ResourceClaimTemplateName must be set.
	Resourceclaimtemplatename string `json:"resourceClaimTemplateName,omitempty"` // ResourceClaimTemplateName is the name of a ResourceClaimTemplate object in the same namespace as this pod. The template will be used to create a new ResourceClaim, which will be bound to this pod. When this pod is deleted, the ResourceClaim will also be deleted. The pod name and resource name, along with a generated component, will be used to form a unique name for the ResourceClaim, which will be recorded in pod.status.resourceClaimStatuses. This field is immutable and no changes will be made to the corresponding ResourceClaim by the control plane after creating the ResourceClaim. Exactly one of ResourceClaimName and ResourceClaimTemplateName must be set.
}

// Iok8sapiresourcev1beta2DeviceSubRequest represents the Iok8sapiresourcev1beta2DeviceSubRequest schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceSubRequest struct {
	Name string `json:"name"` // Name can be used to reference this subrequest in the list of constraints or the list of configurations for the claim. References must use the format <main request>/<subrequest>. Must be a DNS label.
	Selectors []Iok8sapiresourcev1beta2DeviceSelector `json:"selectors,omitempty"` // Selectors define criteria which must be satisfied by a specific device in order for that device to be considered for this subrequest. All selectors must be satisfied for a device to be considered.
	Tolerations []Iok8sapiresourcev1beta2DeviceToleration `json:"tolerations,omitempty"` // If specified, the request's tolerations. Tolerations for NoSchedule are required to allocate a device which has a taint with that effect. The same applies to NoExecute. In addition, should any of the allocated devices get tainted with NoExecute after allocation and that effect is not tolerated, then all pods consuming the ResourceClaim get deleted to evict them. The scheduler will not let new pods reserve the claim while it has these tainted devices. Once all pods are evicted, the claim will get deallocated. The maximum number of tolerations is 16. This is an alpha field and requires enabling the DRADeviceTaints feature gate.
	Allocationmode string `json:"allocationMode,omitempty"` // AllocationMode and its related fields define how devices are allocated to satisfy this subrequest. Supported values are: - ExactCount: This request is for a specific number of devices. This is the default. The exact number is provided in the count field. - All: This subrequest is for all of the matching devices in a pool. Allocation will fail if some devices are already allocated, unless adminAccess is requested. If AllocationMode is not specified, the default mode is ExactCount. If the mode is ExactCount and count is not specified, the default count is one. Any other subrequests must specify this field. More modes may get added in the future. Clients must refuse to handle requests with unknown modes.
	Count int64 `json:"count,omitempty"` // Count is used only when the count mode is "ExactCount". Must be greater than zero. If AllocationMode is ExactCount and this field is not specified, the default is one.
	Deviceclassname string `json:"deviceClassName"` // DeviceClassName references a specific DeviceClass, which can define additional configuration and selectors to be inherited by this subrequest. A class is required. Which classes are available depends on the cluster. Administrators may use this to restrict which devices may get requested by only installing classes with selectors for permitted devices. If users are free to request anything without restrictions, then administrators can create an empty DeviceClass for users to reference.
}

// Iok8sapicertificatesv1CertificateSigningRequest represents the Iok8sapicertificatesv1CertificateSigningRequest schema from the OpenAPI specification
type Iok8sapicertificatesv1CertificateSigningRequest struct {
	Status Iok8sapicertificatesv1CertificateSigningRequestStatus `json:"status,omitempty"` // CertificateSigningRequestStatus contains conditions used to indicate approved/denied/failed status of the request, and the issued certificate.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicertificatesv1CertificateSigningRequestSpec `json:"spec"` // CertificateSigningRequestSpec contains the certificate request.
}

// Iok8sapiflowcontrolv1PriorityLevelConfigurationCondition represents the Iok8sapiflowcontrolv1PriorityLevelConfigurationCondition schema from the OpenAPI specification
type Iok8sapiflowcontrolv1PriorityLevelConfigurationCondition struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // `message` is a human-readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // `reason` is a unique, one-word, CamelCase reason for the condition's last transition.
	Status string `json:"status,omitempty"` // `status` is the status of the condition. Can be True, False, Unknown. Required.
	TypeField string `json:"type,omitempty"` // `type` is the type of the condition. Required.
}

// Iok8sapiappsv1DaemonSetList represents the Iok8sapiappsv1DaemonSetList schema from the OpenAPI specification
type Iok8sapiappsv1DaemonSetList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiappsv1DaemonSet `json:"items"` // A list of daemon sets.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapiappsv1StatefulSetStatus represents the Iok8sapiappsv1StatefulSetStatus schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSetStatus struct {
	Updatedreplicas int `json:"updatedReplicas,omitempty"` // updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.
	Conditions []Iok8sapiappsv1StatefulSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a statefulset's current state.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet's generation, which is updated on mutation by the API Server.
	Currentrevision string `json:"currentRevision,omitempty"` // currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).
	Readyreplicas int `json:"readyReplicas,omitempty"` // readyReplicas is the number of pods created for this StatefulSet with a Ready Condition.
	Collisioncount int `json:"collisionCount,omitempty"` // collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
	Currentreplicas int `json:"currentReplicas,omitempty"` // currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.
	Availablereplicas int `json:"availableReplicas,omitempty"` // Total number of available pods (ready for at least minReadySeconds) targeted by this statefulset.
	Replicas int `json:"replicas"` // replicas is the number of Pods created by the StatefulSet controller.
	Updaterevision string `json:"updateRevision,omitempty"` // updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)
}

// Iok8sapinetworkingv1IngressServiceBackend represents the Iok8sapinetworkingv1IngressServiceBackend schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressServiceBackend struct {
	Port Iok8sapinetworkingv1ServiceBackendPort `json:"port,omitempty"` // ServiceBackendPort is the service port being referenced.
	Name string `json:"name"` // name is the referenced service. The service must exist in the same namespace as the Ingress object.
}

// Iok8sapicorev1NodeAddress represents the Iok8sapicorev1NodeAddress schema from the OpenAPI specification
type Iok8sapicorev1NodeAddress struct {
	Address string `json:"address"` // The node address.
	TypeField string `json:"type"` // Node address type, one of Hostname, ExternalIP or InternalIP.
}

// Iok8sapimachinerypkgapismetav1Patch represents the Iok8sapimachinerypkgapismetav1Patch schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1Patch struct {
}

// Iok8sapiresourcev1beta1ResourceClaimConsumerReference represents the Iok8sapiresourcev1beta1ResourceClaimConsumerReference schema from the OpenAPI specification
type Iok8sapiresourcev1beta1ResourceClaimConsumerReference struct {
	Name string `json:"name"` // Name is the name of resource being referenced.
	Resource string `json:"resource"` // Resource is the type of resource being referenced, for example "pods".
	Uid string `json:"uid"` // UID identifies exactly one incarnation of the resource.
	Apigroup string `json:"apiGroup,omitempty"` // APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
}

// Iok8sapiflowcontrolv1LimitedPriorityLevelConfiguration represents the Iok8sapiflowcontrolv1LimitedPriorityLevelConfiguration schema from the OpenAPI specification
type Iok8sapiflowcontrolv1LimitedPriorityLevelConfiguration struct {
	Borrowinglimitpercent int `json:"borrowingLimitPercent,omitempty"` // `borrowingLimitPercent`, if present, configures a limit on how many seats this priority level can borrow from other priority levels. The limit is known as this level's BorrowingConcurrencyLimit (BorrowingCL) and is a limit on the total number of seats that this level may borrow at any one time. This field holds the ratio of that limit to the level's nominal concurrency limit. When this field is non-nil, it must hold a non-negative integer and the limit is calculated as follows. BorrowingCL(i) = round( NominalCL(i) * borrowingLimitPercent(i)/100.0 ) The value of this field can be more than 100, implying that this priority level can borrow a number of seats that is greater than its own nominal concurrency limit (NominalCL). When this field is left `nil`, the limit is effectively infinite.
	Lendablepercent int `json:"lendablePercent,omitempty"` // `lendablePercent` prescribes the fraction of the level's NominalCL that can be borrowed by other priority levels. The value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level's LendableConcurrencyLimit (LendableCL), is defined as follows. LendableCL(i) = round( NominalCL(i) * lendablePercent(i)/100.0 )
	Limitresponse Iok8sapiflowcontrolv1LimitResponse `json:"limitResponse,omitempty"` // LimitResponse defines how to handle requests that can not be executed right now.
	Nominalconcurrencyshares int `json:"nominalConcurrencyShares,omitempty"` // `nominalConcurrencyShares` (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats available at this priority level. This is used both for requests dispatched from this priority level as well as requests dispatched from other priority levels borrowing seats from this level. The server's concurrency limit (ServerCL) is divided among the Limited priority levels in proportion to their NCS values: NominalCL(i) = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[priority level k] NCS(k) Bigger numbers mean a larger nominal concurrency limit, at the expense of every other priority level. If not specified, this field defaults to a value of 30. Setting this field to zero supports the construction of a "jail" for this priority level that is used to hold some request(s)
}

// Iok8sapicorev1NodeSwapStatus represents the Iok8sapicorev1NodeSwapStatus schema from the OpenAPI specification
type Iok8sapicorev1NodeSwapStatus struct {
	Capacity int64 `json:"capacity,omitempty"` // Total amount of swap memory in bytes.
}

// Iok8sapicorev1ReplicationControllerList represents the Iok8sapicorev1ReplicationControllerList schema from the OpenAPI specification
type Iok8sapicorev1ReplicationControllerList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1ReplicationController `json:"items"` // List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapistoragev1CSINode represents the Iok8sapistoragev1CSINode schema from the OpenAPI specification
type Iok8sapistoragev1CSINode struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapistoragev1CSINodeSpec `json:"spec"` // CSINodeSpec holds information about the specification of all CSI drivers installed on a node
}

// Iok8sapiadmissionregistrationv1alpha1MatchCondition represents the Iok8sapiadmissionregistrationv1alpha1MatchCondition schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1MatchCondition struct {
	Expression string `json:"expression"` // Expression represents the expression which will be evaluated by CEL. Must evaluate to bool. CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables: 'object' - The object from the incoming request. The value is null for DELETE requests. 'oldObject' - The existing object. The value is null for CREATE requests. 'request' - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest). 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request. See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz 'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the request resource. Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/ Required.
	Name string `json:"name"` // Name is an identifier for this match condition, used for strategic merging of MatchConditions, as well as providing an identifier for logging purposes. A good name should be descriptive of the associated expression. Name must be a qualified name consisting of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character (e.g. 'MyName', or 'my.name', or '123-abc', regex used for validation is '([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]') with an optional DNS subdomain prefix and '/' (e.g. 'example.com/MyName') Required.
}

// Iok8sapicorev1ResourceFieldSelector represents the Iok8sapicorev1ResourceFieldSelector schema from the OpenAPI specification
type Iok8sapicorev1ResourceFieldSelector struct {
	Containername string `json:"containerName,omitempty"` // Container name: required for volumes, optional for env vars
	Divisor string `json:"divisor,omitempty"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
	Resource string `json:"resource"` // Required: resource to select
}

// Iok8sapicorev1PersistentVolumeClaimCondition represents the Iok8sapicorev1PersistentVolumeClaimCondition schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaimCondition struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // message is the human-readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports "Resizing" that means the underlying persistent volume is being resized.
	Status string `json:"status"` // Status is the status of the condition. Can be True, False, Unknown. More info: https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/persistent-volume-claim-v1/#:~:text=state%20of%20pvc-,conditions.status,-(string)%2C%20required
	TypeField string `json:"type"` // Type is the type of the condition. More info: https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/persistent-volume-claim-v1/#:~:text=set%20to%20%27ResizeStarted%27.-,PersistentVolumeClaimCondition,-contains%20details%20about
	Lastprobetime string `json:"lastProbeTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
}

// Iok8sapiauthenticationv1SelfSubjectReviewStatus represents the Iok8sapiauthenticationv1SelfSubjectReviewStatus schema from the OpenAPI specification
type Iok8sapiauthenticationv1SelfSubjectReviewStatus struct {
	Userinfo Iok8sapiauthenticationv1UserInfo `json:"userInfo,omitempty"` // UserInfo holds the information about the user needed to implement the user.Info interface.
}

// Iok8sapicorev1LoadBalancerIngress represents the Iok8sapicorev1LoadBalancerIngress schema from the OpenAPI specification
type Iok8sapicorev1LoadBalancerIngress struct {
	Ipmode string `json:"ipMode,omitempty"` // IPMode specifies how the load-balancer IP behaves, and may only be specified when the ip field is specified. Setting this to "VIP" indicates that traffic is delivered to the node with the destination set to the load-balancer's IP and port. Setting this to "Proxy" indicates that traffic is delivered to the node or pod with the destination set to the node's IP and node port or the pod's IP and port. Service implementations may use this information to adjust traffic routing.
	Ports []Iok8sapicorev1PortStatus `json:"ports,omitempty"` // Ports is a list of records of service ports If used, every port defined in the service should have an entry in it
	Hostname string `json:"hostname,omitempty"` // Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)
	Ip string `json:"ip,omitempty"` // IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)
}

// Iok8sapicorev1PodAffinity represents the Iok8sapicorev1PodAffinity schema from the OpenAPI specification
type Iok8sapicorev1PodAffinity struct {
	Preferredduringschedulingignoredduringexecution []Iok8sapicorev1WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"` // The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
	Requiredduringschedulingignoredduringexecution []Iok8sapicorev1PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"` // If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
}

// Iok8sapimachinerypkgruntimeRawExtension represents the Iok8sapimachinerypkgruntimeRawExtension schema from the OpenAPI specification
type Iok8sapimachinerypkgruntimeRawExtension struct {
}

// Iok8sapinetworkingv1Ingress represents the Iok8sapinetworkingv1Ingress schema from the OpenAPI specification
type Iok8sapinetworkingv1Ingress struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapinetworkingv1IngressSpec `json:"spec,omitempty"` // IngressSpec describes the Ingress the user wishes to exist.
	Status Iok8sapinetworkingv1IngressStatus `json:"status,omitempty"` // IngressStatus describe the current state of the Ingress.
}

// Iok8sapinetworkingv1IngressClass represents the Iok8sapinetworkingv1IngressClass schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressClass struct {
	Spec Iok8sapinetworkingv1IngressClassSpec `json:"spec,omitempty"` // IngressClassSpec provides information about the class of an Ingress.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1NodeStatus represents the Iok8sapicorev1NodeStatus schema from the OpenAPI specification
type Iok8sapicorev1NodeStatus struct {
	Nodeinfo Iok8sapicorev1NodeSystemInfo `json:"nodeInfo,omitempty"` // NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
	Runtimehandlers []Iok8sapicorev1NodeRuntimeHandler `json:"runtimeHandlers,omitempty"` // The available runtime handlers.
	Volumesattached []Iok8sapicorev1AttachedVolume `json:"volumesAttached,omitempty"` // List of volumes that are attached to the node.
	Addresses []Iok8sapicorev1NodeAddress `json:"addresses,omitempty"` // List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/reference/node/node-status/#addresses Note: This field is declared as mergeable, but the merge key is not sufficiently unique, which can cause data corruption when it is merged. Callers should instead use a full-replacement patch. See https://pr.k8s.io/79391 for an example. Consumers should assume that addresses can change during the lifetime of a Node. However, there are some exceptions where this may not be possible, such as Pods that inherit a Node's address in its own status or consumers of the downward API (status.hostIP).
	Conditions []Iok8sapicorev1NodeCondition `json:"conditions,omitempty"` // Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/reference/node/node-status/#condition
	Config Iok8sapicorev1NodeConfigStatus `json:"config,omitempty"` // NodeConfigStatus describes the status of the config assigned by Node.Spec.ConfigSource.
	Images []Iok8sapicorev1ContainerImage `json:"images,omitempty"` // List of container images on this node
	Allocatable map[string]interface{} `json:"allocatable,omitempty"` // Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
	Volumesinuse []string `json:"volumesInUse,omitempty"` // List of attachable volumes in use (mounted) by the node.
	Phase string `json:"phase,omitempty"` // NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated.
	Capacity map[string]interface{} `json:"capacity,omitempty"` // Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/reference/node/node-status/#capacity
	Daemonendpoints Iok8sapicorev1NodeDaemonEndpoints `json:"daemonEndpoints,omitempty"` // NodeDaemonEndpoints lists ports opened by daemons running on the Node.
	Features Iok8sapicorev1NodeFeatures `json:"features,omitempty"` // NodeFeatures describes the set of features implemented by the CRI implementation. The features contained in the NodeFeatures should depend only on the cri implementation independent of runtime handlers.
}

// Iok8sapiauthenticationv1BoundObjectReference represents the Iok8sapiauthenticationv1BoundObjectReference schema from the OpenAPI specification
type Iok8sapiauthenticationv1BoundObjectReference struct {
	Kind string `json:"kind,omitempty"` // Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
	Name string `json:"name,omitempty"` // Name of the referent.
	Uid string `json:"uid,omitempty"` // UID of the referent.
	Apiversion string `json:"apiVersion,omitempty"` // API version of the referent.
}

// Iok8sapicorev1RBDPersistentVolumeSource represents the Iok8sapicorev1RBDPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1RBDPersistentVolumeSource struct {
	Monitors []string `json:"monitors"` // monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Pool string `json:"pool,omitempty"` // pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Readonly bool `json:"readOnly,omitempty"` // readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Secretref Iok8sapicorev1SecretReference `json:"secretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	User string `json:"user,omitempty"` // user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	Image string `json:"image"` // image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Keyring string `json:"keyring,omitempty"` // keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
}

// Iok8sapiresourcev1alpha3DeviceTaint represents the Iok8sapiresourcev1alpha3DeviceTaint schema from the OpenAPI specification
type Iok8sapiresourcev1alpha3DeviceTaint struct {
	Timeadded string `json:"timeAdded,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Value string `json:"value,omitempty"` // The taint value corresponding to the taint key. Must be a label value.
	Effect string `json:"effect"` // The effect of the taint on claims that do not tolerate the taint and through such claims on the pods using them. Valid effects are NoSchedule and NoExecute. PreferNoSchedule as used for nodes is not valid here.
	Key string `json:"key"` // The taint key to be applied to a device. Must be a label name.
}

// Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyBindingSpec represents the Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyBindingSpec schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyBindingSpec struct {
	Matchresources Iok8sapiadmissionregistrationv1alpha1MatchResources `json:"matchResources,omitempty"` // MatchResources decides whether to run the admission control policy on an object based on whether it meets the match criteria. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)
	Paramref Iok8sapiadmissionregistrationv1alpha1ParamRef `json:"paramRef,omitempty"` // ParamRef describes how to locate the params to be used as input to expressions of rules applied by a policy binding.
	Policyname string `json:"policyName,omitempty"` // policyName references a MutatingAdmissionPolicy name which the MutatingAdmissionPolicyBinding binds to. If the referenced resource does not exist, this binding is considered invalid and will be ignored Required.
}

// Iok8sapinetworkingv1beta1ParentReference represents the Iok8sapinetworkingv1beta1ParentReference schema from the OpenAPI specification
type Iok8sapinetworkingv1beta1ParentReference struct {
	Group string `json:"group,omitempty"` // Group is the group of the object being referenced.
	Name string `json:"name"` // Name is the name of the object being referenced.
	Namespace string `json:"namespace,omitempty"` // Namespace is the namespace of the object being referenced.
	Resource string `json:"resource"` // Resource is the resource of the object being referenced.
}

// Iok8sapicorev1HTTPGetAction represents the Iok8sapicorev1HTTPGetAction schema from the OpenAPI specification
type Iok8sapicorev1HTTPGetAction struct {
	Httpheaders []Iok8sapicorev1HTTPHeader `json:"httpHeaders,omitempty"` // Custom headers to set in the request. HTTP allows repeated headers.
	Path string `json:"path,omitempty"` // Path to access on the HTTP server.
	Port string `json:"port"` // IntOrString is a type that can hold an int32 or a string. When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type. This allows you to have, for example, a JSON field that can accept a name or number.
	Scheme string `json:"scheme,omitempty"` // Scheme to use for connecting to the host. Defaults to HTTP.
	Host string `json:"host,omitempty"` // Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
}

// Iok8sapicorev1DownwardAPIVolumeSource represents the Iok8sapicorev1DownwardAPIVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1DownwardAPIVolumeSource struct {
	Defaultmode int `json:"defaultMode,omitempty"` // Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Items []Iok8sapicorev1DownwardAPIVolumeFile `json:"items,omitempty"` // Items is a list of downward API volume file
}

// Iok8sapinetworkingv1IngressSpec represents the Iok8sapinetworkingv1IngressSpec schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressSpec struct {
	Defaultbackend Iok8sapinetworkingv1IngressBackend `json:"defaultBackend,omitempty"` // IngressBackend describes all endpoints for a given service and port.
	Ingressclassname string `json:"ingressClassName,omitempty"` // ingressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -> IngressClass -> Ingress resource). Although the `kubernetes.io/ingress.class` annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present.
	Rules []Iok8sapinetworkingv1IngressRule `json:"rules,omitempty"` // rules is a list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
	Tls []Iok8sapinetworkingv1IngressTLS `json:"tls,omitempty"` // tls represents the TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
}

// Iok8sapicorev1EnvFromSource represents the Iok8sapicorev1EnvFromSource schema from the OpenAPI specification
type Iok8sapicorev1EnvFromSource struct {
	Configmapref Iok8sapicorev1ConfigMapEnvSource `json:"configMapRef,omitempty"` // ConfigMapEnvSource selects a ConfigMap to populate the environment variables with. The contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.
	Prefix string `json:"prefix,omitempty"` // Optional text to prepend to the name of each environment variable. May consist of any printable ASCII characters except '='.
	Secretref Iok8sapicorev1SecretEnvSource `json:"secretRef,omitempty"` // SecretEnvSource selects a Secret to populate the environment variables with. The contents of the target Secret's Data field will represent the key-value pairs as environment variables.
}

// Iok8sapiresourcev1beta1ResourcePool represents the Iok8sapiresourcev1beta1ResourcePool schema from the OpenAPI specification
type Iok8sapiresourcev1beta1ResourcePool struct {
	Name string `json:"name"` // Name is used to identify the pool. For node-local devices, this is often the node name, but this is not required. It must not be longer than 253 characters and must consist of one or more DNS sub-domains separated by slashes. This field is immutable.
	Resourceslicecount int64 `json:"resourceSliceCount"` // ResourceSliceCount is the total number of ResourceSlices in the pool at this generation number. Must be greater than zero. Consumers can use this to check whether they have seen all ResourceSlices belonging to the same pool.
	Generation int64 `json:"generation"` // Generation tracks the change in a pool over time. Whenever a driver changes something about one or more of the resources in a pool, it must change the generation in all ResourceSlices which are part of that pool. Consumers of ResourceSlices should only consider resources from the pool with the highest generation number. The generation may be reset by drivers, which should be fine for consumers, assuming that all ResourceSlices in a pool are updated to match or deleted. Combined with ResourceSliceCount, this mechanism enables consumers to detect pools which are comprised of multiple ResourceSlices and are in an incomplete state.
}

// Iok8sapicorev1PodIP represents the Iok8sapicorev1PodIP schema from the OpenAPI specification
type Iok8sapicorev1PodIP struct {
	Ip string `json:"ip"` // IP is the IP address assigned to the pod
}

// Iok8sapiauthorizationv1ResourceAttributes represents the Iok8sapiauthorizationv1ResourceAttributes schema from the OpenAPI specification
type Iok8sapiauthorizationv1ResourceAttributes struct {
	Labelselector Iok8sapiauthorizationv1LabelSelectorAttributes `json:"labelSelector,omitempty"` // LabelSelectorAttributes indicates a label limited access. Webhook authors are encouraged to * ensure rawSelector and requirements are not both set * consider the requirements field if set * not try to parse or consider the rawSelector field if set. This is to avoid another CVE-2022-2880 (i.e. getting different systems to agree on how exactly to parse a query is not something we want), see https://www.oxeye.io/resources/golang-parameter-smuggling-attack for more details. For the *SubjectAccessReview endpoints of the kube-apiserver: * If rawSelector is empty and requirements are empty, the request is not limited. * If rawSelector is present and requirements are empty, the rawSelector will be parsed and limited if the parsing succeeds. * If rawSelector is empty and requirements are present, the requirements should be honored * If rawSelector is present and requirements are present, the request is invalid.
	Name string `json:"name,omitempty"` // Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
	Version string `json:"version,omitempty"` // Version is the API Version of the Resource. "*" means all.
	Fieldselector Iok8sapiauthorizationv1FieldSelectorAttributes `json:"fieldSelector,omitempty"` // FieldSelectorAttributes indicates a field limited access. Webhook authors are encouraged to * ensure rawSelector and requirements are not both set * consider the requirements field if set * not try to parse or consider the rawSelector field if set. This is to avoid another CVE-2022-2880 (i.e. getting different systems to agree on how exactly to parse a query is not something we want), see https://www.oxeye.io/resources/golang-parameter-smuggling-attack for more details. For the *SubjectAccessReview endpoints of the kube-apiserver: * If rawSelector is empty and requirements are empty, the request is not limited. * If rawSelector is present and requirements are empty, the rawSelector will be parsed and limited if the parsing succeeds. * If rawSelector is empty and requirements are present, the requirements should be honored * If rawSelector is present and requirements are present, the request is invalid.
	Namespace string `json:"namespace,omitempty"` // Namespace is the namespace of the action being requested. Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
	Resource string `json:"resource,omitempty"` // Resource is one of the existing resource types. "*" means all.
	Verb string `json:"verb,omitempty"` // Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy. "*" means all.
	Group string `json:"group,omitempty"` // Group is the API Group of the Resource. "*" means all.
	Subresource string `json:"subresource,omitempty"` // Subresource is one of the existing resource types. "" means none.
}

// Iok8sapistoragev1alpha1VolumeAttributesClass represents the Iok8sapistoragev1alpha1VolumeAttributesClass schema from the OpenAPI specification
type Iok8sapistoragev1alpha1VolumeAttributesClass struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Parameters map[string]interface{} `json:"parameters,omitempty"` // parameters hold volume attributes defined by the CSI driver. These values are opaque to the Kubernetes and are passed directly to the CSI driver. The underlying storage provider supports changing these attributes on an existing volume, however the parameters field itself is immutable. To invoke a volume update, a new VolumeAttributesClass should be created with new parameters, and the PersistentVolumeClaim should be updated to reference the new VolumeAttributesClass. This field is required and must contain at least one key/value pair. The keys cannot be empty, and the maximum number of parameters is 512, with a cumulative max size of 256K. If the CSI driver rejects invalid parameters, the target PersistentVolumeClaim will be set to an "Infeasible" state in the modifyVolumeStatus field.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Drivername string `json:"driverName"` // Name of the CSI driver This field is immutable.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapicorev1PodDNSConfigOption represents the Iok8sapicorev1PodDNSConfigOption schema from the OpenAPI specification
type Iok8sapicorev1PodDNSConfigOption struct {
	Name string `json:"name,omitempty"` // Name is this DNS resolver option's name. Required.
	Value string `json:"value,omitempty"` // Value is this DNS resolver option's value.
}

// Iok8sapiautoscalingv1HorizontalPodAutoscalerSpec represents the Iok8sapiautoscalingv1HorizontalPodAutoscalerSpec schema from the OpenAPI specification
type Iok8sapiautoscalingv1HorizontalPodAutoscalerSpec struct {
	Scaletargetref Iok8sapiautoscalingv1CrossVersionObjectReference `json:"scaleTargetRef"` // CrossVersionObjectReference contains enough information to let you identify the referred resource.
	Targetcpuutilizationpercentage int `json:"targetCPUUtilizationPercentage,omitempty"` // targetCPUUtilizationPercentage is the target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
	Maxreplicas int `json:"maxReplicas"` // maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
	Minreplicas int `json:"minReplicas,omitempty"` // minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod. minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured. Scaling is active as long as at least one metric value is available.
}

// Iok8sapicorev1ResourceQuotaSpec represents the Iok8sapicorev1ResourceQuotaSpec schema from the OpenAPI specification
type Iok8sapicorev1ResourceQuotaSpec struct {
	Hard map[string]interface{} `json:"hard,omitempty"` // hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Scopeselector Iok8sapicorev1ScopeSelector `json:"scopeSelector,omitempty"` // A scope selector represents the AND of the selectors represented by the scoped-resource selector requirements.
	Scopes []string `json:"scopes,omitempty"` // A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.
}

// Iok8sapiautoscalingv2HorizontalPodAutoscaler represents the Iok8sapiautoscalingv2HorizontalPodAutoscaler schema from the OpenAPI specification
type Iok8sapiautoscalingv2HorizontalPodAutoscaler struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiautoscalingv2HorizontalPodAutoscalerSpec `json:"spec,omitempty"` // HorizontalPodAutoscalerSpec describes the desired functionality of the HorizontalPodAutoscaler.
	Status Iok8sapiautoscalingv2HorizontalPodAutoscalerStatus `json:"status,omitempty"` // HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
}

// Iok8sapiresourcev1beta2ExactDeviceRequest represents the Iok8sapiresourcev1beta2ExactDeviceRequest schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ExactDeviceRequest struct {
	Selectors []Iok8sapiresourcev1beta2DeviceSelector `json:"selectors,omitempty"` // Selectors define criteria which must be satisfied by a specific device in order for that device to be considered for this request. All selectors must be satisfied for a device to be considered.
	Tolerations []Iok8sapiresourcev1beta2DeviceToleration `json:"tolerations,omitempty"` // If specified, the request's tolerations. Tolerations for NoSchedule are required to allocate a device which has a taint with that effect. The same applies to NoExecute. In addition, should any of the allocated devices get tainted with NoExecute after allocation and that effect is not tolerated, then all pods consuming the ResourceClaim get deleted to evict them. The scheduler will not let new pods reserve the claim while it has these tainted devices. Once all pods are evicted, the claim will get deallocated. The maximum number of tolerations is 16. This is an alpha field and requires enabling the DRADeviceTaints feature gate.
	Adminaccess bool `json:"adminAccess,omitempty"` // AdminAccess indicates that this is a claim for administrative access to the device(s). Claims with AdminAccess are expected to be used for monitoring or other management services for a device. They ignore all ordinary claims to the device with respect to access modes and any resource allocations. This is an alpha field and requires enabling the DRAAdminAccess feature gate. Admin access is disabled if this field is unset or set to false, otherwise it is enabled.
	Allocationmode string `json:"allocationMode,omitempty"` // AllocationMode and its related fields define how devices are allocated to satisfy this request. Supported values are: - ExactCount: This request is for a specific number of devices. This is the default. The exact number is provided in the count field. - All: This request is for all of the matching devices in a pool. At least one device must exist on the node for the allocation to succeed. Allocation will fail if some devices are already allocated, unless adminAccess is requested. If AllocationMode is not specified, the default mode is ExactCount. If the mode is ExactCount and count is not specified, the default count is one. Any other requests must specify this field. More modes may get added in the future. Clients must refuse to handle requests with unknown modes.
	Count int64 `json:"count,omitempty"` // Count is used only when the count mode is "ExactCount". Must be greater than zero. If AllocationMode is ExactCount and this field is not specified, the default is one.
	Deviceclassname string `json:"deviceClassName"` // DeviceClassName references a specific DeviceClass, which can define additional configuration and selectors to be inherited by this request. A DeviceClassName is required. Administrators may use this to restrict which devices may get requested by only installing classes with selectors for permitted devices. If users are free to request anything without restrictions, then administrators can create an empty DeviceClass for users to reference.
}

// Iok8sapiauthenticationv1SelfSubjectReview represents the Iok8sapiauthenticationv1SelfSubjectReview schema from the OpenAPI specification
type Iok8sapiauthenticationv1SelfSubjectReview struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Status Iok8sapiauthenticationv1SelfSubjectReviewStatus `json:"status,omitempty"` // SelfSubjectReviewStatus is filled by the kube-apiserver and sent back to a user.
}

// Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicySpec represents the Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicySpec schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicySpec struct {
	Matchconditions []Iok8sapiadmissionregistrationv1alpha1MatchCondition `json:"matchConditions,omitempty"` // matchConditions is a list of conditions that must be met for a request to be validated. Match conditions filter requests that have already been matched by the matchConstraints. An empty list of matchConditions matches all requests. There are a maximum of 64 match conditions allowed. If a parameter object is provided, it can be accessed via the `params` handle in the same manner as validation expressions. The exact matching logic is (in order): 1. If ANY matchCondition evaluates to FALSE, the policy is skipped. 2. If ALL matchConditions evaluate to TRUE, the policy is evaluated. 3. If any matchCondition evaluates to an error (but none are FALSE): - If failurePolicy=Fail, reject the request - If failurePolicy=Ignore, the policy is skipped
	Matchconstraints Iok8sapiadmissionregistrationv1alpha1MatchResources `json:"matchConstraints,omitempty"` // MatchResources decides whether to run the admission control policy on an object based on whether it meets the match criteria. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)
	Mutations []Iok8sapiadmissionregistrationv1alpha1Mutation `json:"mutations,omitempty"` // mutations contain operations to perform on matching objects. mutations may not be empty; a minimum of one mutation is required. mutations are evaluated in order, and are reinvoked according to the reinvocationPolicy. The mutations of a policy are invoked for each binding of this policy and reinvocation of mutations occurs on a per binding basis.
	Paramkind Iok8sapiadmissionregistrationv1alpha1ParamKind `json:"paramKind,omitempty"` // ParamKind is a tuple of Group Kind and Version.
	Reinvocationpolicy string `json:"reinvocationPolicy,omitempty"` // reinvocationPolicy indicates whether mutations may be called multiple times per MutatingAdmissionPolicyBinding as part of a single admission evaluation. Allowed values are "Never" and "IfNeeded". Never: These mutations will not be called more than once per binding in a single admission evaluation. IfNeeded: These mutations may be invoked more than once per binding for a single admission request and there is no guarantee of order with respect to other admission plugins, admission webhooks, bindings of this policy and admission policies. Mutations are only reinvoked when mutations change the object after this mutation is invoked. Required.
	Variables []Iok8sapiadmissionregistrationv1alpha1Variable `json:"variables,omitempty"` // variables contain definitions of variables that can be used in composition of other expressions. Each variable is defined as a named CEL expression. The variables defined here will be available under `variables` in other expressions of the policy except matchConditions because matchConditions are evaluated before the rest of the policy. The expression of a variable can refer to other variables defined earlier in the list but not those after. Thus, variables must be sorted by the order of first appearance and acyclic.
	Failurepolicy string `json:"failurePolicy,omitempty"` // failurePolicy defines how to handle failures for the admission policy. Failures can occur from CEL expression parse errors, type check errors, runtime errors and invalid or mis-configured policy definitions or bindings. A policy is invalid if paramKind refers to a non-existent Kind. A binding is invalid if paramRef.name refers to a non-existent resource. failurePolicy does not define how validations that evaluate to false are handled. Allowed values are Ignore or Fail. Defaults to Fail.
}

// Iok8sapistoragev1VolumeAttachmentSpec represents the Iok8sapistoragev1VolumeAttachmentSpec schema from the OpenAPI specification
type Iok8sapistoragev1VolumeAttachmentSpec struct {
	Nodename string `json:"nodeName"` // nodeName represents the node that the volume should be attached to.
	Source Iok8sapistoragev1VolumeAttachmentSource `json:"source"` // VolumeAttachmentSource represents a volume that should be attached. Right now only PersistentVolumes can be attached via external attacher, in the future we may allow also inline volumes in pods. Exactly one member can be set.
	Attacher string `json:"attacher"` // attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
}

// Iok8sapiadmissionregistrationv1alpha1ApplyConfiguration represents the Iok8sapiadmissionregistrationv1alpha1ApplyConfiguration schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1ApplyConfiguration struct {
	Expression string `json:"expression,omitempty"` // expression will be evaluated by CEL to create an apply configuration. ref: https://github.com/google/cel-spec Apply configurations are declared in CEL using object initialization. For example, this CEL expression returns an apply configuration to set a single field: 	Object{ 	 spec: Object.spec{ 	 serviceAccountName: "example" 	 } 	} Apply configurations may not modify atomic structs, maps or arrays due to the risk of accidental deletion of values not included in the apply configuration. CEL expressions have access to the object types needed to create apply configurations: - 'Object' - CEL type of the resource object. - 'Object.<fieldName>' - CEL type of object field (such as 'Object.spec') - 'Object.<fieldName1>.<fieldName2>...<fieldNameN>` - CEL type of nested field (such as 'Object.spec.containers') CEL expressions have access to the contents of the API request, organized into CEL variables as well as some other useful variables: - 'object' - The object from the incoming request. The value is null for DELETE requests. - 'oldObject' - The existing object. The value is null for CREATE requests. - 'request' - Attributes of the API request([ref](/pkg/apis/admission/types.go#AdmissionRequest)). - 'params' - Parameter resource referred to by the policy binding being evaluated. Only populated if the policy has a ParamKind. - 'namespaceObject' - The namespace object that the incoming object belongs to. The value is null for cluster-scoped resources. - 'variables' - Map of composited variables, from its name to its lazily evaluated value. For example, a variable named 'foo' can be accessed as 'variables.foo'. - 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request. See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz - 'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the request resource. The `apiVersion`, `kind`, `metadata.name` and `metadata.generateName` are always accessible from the root of the object. No other metadata properties are accessible. Only property names of the form `[a-zA-Z_.-/][a-zA-Z0-9_.-/]*` are accessible. Required.
}

// Iok8sapistoragev1alpha1VolumeAttributesClassList represents the Iok8sapistoragev1alpha1VolumeAttributesClassList schema from the OpenAPI specification
type Iok8sapistoragev1alpha1VolumeAttributesClassList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapistoragev1alpha1VolumeAttributesClass `json:"items"` // items is the list of VolumeAttributesClass objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiappsv1ReplicaSetCondition represents the Iok8sapiappsv1ReplicaSetCondition schema from the OpenAPI specification
type Iok8sapiappsv1ReplicaSetCondition struct {
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of replica set condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
}

// Iok8sapiflowcontrolv1FlowSchemaList represents the Iok8sapiflowcontrolv1FlowSchemaList schema from the OpenAPI specification
type Iok8sapiflowcontrolv1FlowSchemaList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiflowcontrolv1FlowSchema `json:"items"` // `items` is a list of FlowSchemas.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapicorev1ISCSIPersistentVolumeSource represents the Iok8sapicorev1ISCSIPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ISCSIPersistentVolumeSource struct {
	Portals []string `json:"portals,omitempty"` // portals is the iSCSI Target Portal List. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Secretref Iok8sapicorev1SecretReference `json:"secretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	Targetportal string `json:"targetPortal"` // targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Chapauthdiscovery bool `json:"chapAuthDiscovery,omitempty"` // chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication
	Chapauthsession bool `json:"chapAuthSession,omitempty"` // chapAuthSession defines whether support iSCSI Session CHAP authentication
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	Iqn string `json:"iqn"` // iqn is Target iSCSI Qualified Name.
	Readonly bool `json:"readOnly,omitempty"` // readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
	Lun int `json:"lun"` // lun is iSCSI Target Lun number.
	Initiatorname string `json:"initiatorName,omitempty"` // initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
	Iscsiinterface string `json:"iscsiInterface,omitempty"` // iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
}

// Iok8sapiflowcontrolv1PriorityLevelConfigurationStatus represents the Iok8sapiflowcontrolv1PriorityLevelConfigurationStatus schema from the OpenAPI specification
type Iok8sapiflowcontrolv1PriorityLevelConfigurationStatus struct {
	Conditions []Iok8sapiflowcontrolv1PriorityLevelConfigurationCondition `json:"conditions,omitempty"` // `conditions` is the current state of "request-priority".
}

// Iok8sapicorev1PodReadinessGate represents the Iok8sapicorev1PodReadinessGate schema from the OpenAPI specification
type Iok8sapicorev1PodReadinessGate struct {
	Conditiontype string `json:"conditionType"` // ConditionType refers to a condition in the pod's condition list with matching type.
}

// Iok8sapiappsv1RollingUpdateDeployment represents the Iok8sapiappsv1RollingUpdateDeployment schema from the OpenAPI specification
type Iok8sapiappsv1RollingUpdateDeployment struct {
	Maxsurge string `json:"maxSurge,omitempty"` // IntOrString is a type that can hold an int32 or a string. When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type. This allows you to have, for example, a JSON field that can accept a name or number.
	Maxunavailable string `json:"maxUnavailable,omitempty"` // IntOrString is a type that can hold an int32 or a string. When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type. This allows you to have, for example, a JSON field that can accept a name or number.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// Iok8sapicorev1PersistentVolumeList represents the Iok8sapicorev1PersistentVolumeList schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1PersistentVolume `json:"items"` // items is a list of persistent volumes. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapimachinerypkgapismetav1Condition represents the Iok8sapimachinerypkgapismetav1Condition schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1Condition struct {
	Message string `json:"message"` // message is a human readable message indicating details about the transition. This may be an empty string.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance.
	Reason string `json:"reason"` // reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty.
	Status string `json:"status"` // status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // type of condition in CamelCase or in foo.example.com/CamelCase.
	Lasttransitiontime string `json:"lastTransitionTime"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
}

// Iok8sapicoordinationv1alpha2LeaseCandidateSpec represents the Iok8sapicoordinationv1alpha2LeaseCandidateSpec schema from the OpenAPI specification
type Iok8sapicoordinationv1alpha2LeaseCandidateSpec struct {
	Renewtime string `json:"renewTime,omitempty"` // MicroTime is version of Time with microsecond level precision.
	Strategy string `json:"strategy"` // Strategy is the strategy that coordinated leader election will use for picking the leader. If multiple candidates for the same Lease return different strategies, the strategy provided by the candidate with the latest BinaryVersion will be used. If there is still conflict, this is a user error and coordinated leader election will not operate the Lease until resolved.
	Binaryversion string `json:"binaryVersion"` // BinaryVersion is the binary version. It must be in a semver format without leading `v`. This field is required.
	Emulationversion string `json:"emulationVersion,omitempty"` // EmulationVersion is the emulation version. It must be in a semver format without leading `v`. EmulationVersion must be less than or equal to BinaryVersion. This field is required when strategy is "OldestEmulationVersion"
	Leasename string `json:"leaseName"` // LeaseName is the name of the lease for which this candidate is contending. This field is immutable.
	Pingtime string `json:"pingTime,omitempty"` // MicroTime is version of Time with microsecond level precision.
}

// Iok8sapiauthorizationv1LocalSubjectAccessReview represents the Iok8sapiauthorizationv1LocalSubjectAccessReview schema from the OpenAPI specification
type Iok8sapiauthorizationv1LocalSubjectAccessReview struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1SubjectAccessReviewSpec `json:"spec"` // SubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
	Status Iok8sapiauthorizationv1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1PersistentVolumeClaimStatus represents the Iok8sapicorev1PersistentVolumeClaimStatus schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaimStatus struct {
	Currentvolumeattributesclassname string `json:"currentVolumeAttributesClassName,omitempty"` // currentVolumeAttributesClassName is the current name of the VolumeAttributesClass the PVC is using. When unset, there is no VolumeAttributeClass applied to this PersistentVolumeClaim This is a beta field and requires enabling VolumeAttributesClass feature (off by default).
	Modifyvolumestatus Iok8sapicorev1ModifyVolumeStatus `json:"modifyVolumeStatus,omitempty"` // ModifyVolumeStatus represents the status object of ControllerModifyVolume operation
	Phase string `json:"phase,omitempty"` // phase represents the current phase of PersistentVolumeClaim.
	Accessmodes []string `json:"accessModes,omitempty"` // accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	Allocatedresourcestatuses map[string]interface{} `json:"allocatedResourceStatuses,omitempty"` // allocatedResourceStatuses stores status of resource being resized for the given PVC. Key names follow standard Kubernetes label syntax. Valid values are either: 	* Un-prefixed keys: 		- storage - the capacity of the volume. 	* Custom resources must use implementation-defined prefixed names such as "example.com/my-custom-resource" Apart from above values - keys that are unprefixed or have kubernetes.io prefix are considered reserved and hence may not be used. ClaimResourceStatus can be in any of following states: 	- ControllerResizeInProgress: 		State set when resize controller starts resizing the volume in control-plane. 	- ControllerResizeFailed: 		State set when resize has failed in resize controller with a terminal error. 	- NodeResizePending: 		State set when resize controller has finished resizing the volume but further resizing of 		volume is needed on the node. 	- NodeResizeInProgress: 		State set when kubelet starts resizing the volume. 	- NodeResizeFailed: 		State set when resizing has failed in kubelet with a terminal error. Transient errors don't set 		NodeResizeFailed. For example: if expanding a PVC for more capacity - this field can be one of the following states: 	- pvc.status.allocatedResourceStatus['storage'] = "ControllerResizeInProgress" - pvc.status.allocatedResourceStatus['storage'] = "ControllerResizeFailed" - pvc.status.allocatedResourceStatus['storage'] = "NodeResizePending" - pvc.status.allocatedResourceStatus['storage'] = "NodeResizeInProgress" - pvc.status.allocatedResourceStatus['storage'] = "NodeResizeFailed" When this field is not set, it means that no resize operation is in progress for the given PVC. A controller that receives PVC update with previously unknown resourceName or ClaimResourceStatus should ignore the update for the purpose it was designed. For example - a controller that only is responsible for resizing capacity of the volume, should ignore PVC updates that change other valid resources associated with PVC. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.
	Allocatedresources map[string]interface{} `json:"allocatedResources,omitempty"` // allocatedResources tracks the resources allocated to a PVC including its capacity. Key names follow standard Kubernetes label syntax. Valid values are either: 	* Un-prefixed keys: 		- storage - the capacity of the volume. 	* Custom resources must use implementation-defined prefixed names such as "example.com/my-custom-resource" Apart from above values - keys that are unprefixed or have kubernetes.io prefix are considered reserved and hence may not be used. Capacity reported here may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity. A controller that receives PVC update with previously unknown resourceName should ignore the update for the purpose it was designed. For example - a controller that only is responsible for resizing capacity of the volume, should ignore PVC updates that change other valid resources associated with PVC. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.
	Capacity map[string]interface{} `json:"capacity,omitempty"` // capacity represents the actual resources of the underlying volume.
	Conditions []Iok8sapicorev1PersistentVolumeClaimCondition `json:"conditions,omitempty"` // conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'Resizing'.
}

// Iok8sapiresourcev1alpha3DeviceTaintRuleSpec represents the Iok8sapiresourcev1alpha3DeviceTaintRuleSpec schema from the OpenAPI specification
type Iok8sapiresourcev1alpha3DeviceTaintRuleSpec struct {
	Taint Iok8sapiresourcev1alpha3DeviceTaint `json:"taint"` // The device this taint is attached to has the "effect" on any claim which does not tolerate the taint and, through the claim, to pods using the claim.
	Deviceselector Iok8sapiresourcev1alpha3DeviceTaintSelector `json:"deviceSelector,omitempty"` // DeviceTaintSelector defines which device(s) a DeviceTaintRule applies to. The empty selector matches all devices. Without a selector, no devices are matched.
}

// Iok8sapiadmissionregistrationv1ExpressionWarning represents the Iok8sapiadmissionregistrationv1ExpressionWarning schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ExpressionWarning struct {
	Warning string `json:"warning"` // The content of type checking information in a human-readable form. Each line of the warning contains the type that the expression is checked against, followed by the type check error from the compiler.
	Fieldref string `json:"fieldRef"` // The path to the field that refers the expression. For example, the reference to the expression of the first item of validations is "spec.validations[0].expression"
}

// Iok8sapiresourcev1beta2DeviceAllocationConfiguration represents the Iok8sapiresourcev1beta2DeviceAllocationConfiguration schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceAllocationConfiguration struct {
	Requests []string `json:"requests,omitempty"` // Requests lists the names of requests where the configuration applies. If empty, its applies to all requests. References to subrequests must include the name of the main request and may include the subrequest using the format <main request>[/<subrequest>]. If just the main request is given, the configuration applies to all subrequests.
	Source string `json:"source"` // Source records whether the configuration comes from a class and thus is not something that a normal user would have been able to set or from a claim.
	Opaque Iok8sapiresourcev1beta2OpaqueDeviceConfiguration `json:"opaque,omitempty"` // OpaqueDeviceConfiguration contains configuration parameters for a driver in a format defined by the driver vendor.
}

// Iok8sapiresourcev1beta2ResourceClaimSpec represents the Iok8sapiresourcev1beta2ResourceClaimSpec schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ResourceClaimSpec struct {
	Devices Iok8sapiresourcev1beta2DeviceClaim `json:"devices,omitempty"` // DeviceClaim defines how to request devices with a ResourceClaim.
}

// Iok8sapicorev1ContainerStateWaiting represents the Iok8sapicorev1ContainerStateWaiting schema from the OpenAPI specification
type Iok8sapicorev1ContainerStateWaiting struct {
	Message string `json:"message,omitempty"` // Message regarding why the container is not yet running.
	Reason string `json:"reason,omitempty"` // (brief) reason the container is not yet running.
}

// Iok8sapiflowcontrolv1PriorityLevelConfiguration represents the Iok8sapiflowcontrolv1PriorityLevelConfiguration schema from the OpenAPI specification
type Iok8sapiflowcontrolv1PriorityLevelConfiguration struct {
	Status Iok8sapiflowcontrolv1PriorityLevelConfigurationStatus `json:"status,omitempty"` // PriorityLevelConfigurationStatus represents the current state of a "request-priority".
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiflowcontrolv1PriorityLevelConfigurationSpec `json:"spec,omitempty"` // PriorityLevelConfigurationSpec specifies the configuration of a priority level.
}

// Iok8sapicorev1ClusterTrustBundleProjection represents the Iok8sapicorev1ClusterTrustBundleProjection schema from the OpenAPI specification
type Iok8sapicorev1ClusterTrustBundleProjection struct {
	Labelselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"labelSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Name string `json:"name,omitempty"` // Select a single ClusterTrustBundle by object name. Mutually-exclusive with signerName and labelSelector.
	Optional bool `json:"optional,omitempty"` // If true, don't block pod startup if the referenced ClusterTrustBundle(s) aren't available. If using name, then the named ClusterTrustBundle is allowed not to exist. If using signerName, then the combination of signerName and labelSelector is allowed to match zero ClusterTrustBundles.
	Path string `json:"path"` // Relative path from the volume root to write the bundle.
	Signername string `json:"signerName,omitempty"` // Select all ClusterTrustBundles that match this signer name. Mutually-exclusive with name. The contents of all selected ClusterTrustBundles will be unified and deduplicated.
}

// Iok8sapistoragev1TokenRequest represents the Iok8sapistoragev1TokenRequest schema from the OpenAPI specification
type Iok8sapistoragev1TokenRequest struct {
	Audience string `json:"audience"` // audience is the intended audience of the token in "TokenRequestSpec". It will default to the audiences of kube apiserver.
	Expirationseconds int64 `json:"expirationSeconds,omitempty"` // expirationSeconds is the duration of validity of the token in "TokenRequestSpec". It has the same default value of "ExpirationSeconds" in "TokenRequestSpec".
}

// Iok8sapinetworkingv1IPAddress represents the Iok8sapinetworkingv1IPAddress schema from the OpenAPI specification
type Iok8sapinetworkingv1IPAddress struct {
	Spec Iok8sapinetworkingv1IPAddressSpec `json:"spec,omitempty"` // IPAddressSpec describe the attributes in an IP Address.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiresourcev1beta2DeviceRequestAllocationResult represents the Iok8sapiresourcev1beta2DeviceRequestAllocationResult schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceRequestAllocationResult struct {
	Request string `json:"request"` // Request is the name of the request in the claim which caused this device to be allocated. If it references a subrequest in the firstAvailable list on a DeviceRequest, this field must include both the name of the main request and the subrequest using the format <main request>/<subrequest>. Multiple devices may have been allocated per request.
	Tolerations []Iok8sapiresourcev1beta2DeviceToleration `json:"tolerations,omitempty"` // A copy of all tolerations specified in the request at the time when the device got allocated. The maximum number of tolerations is 16. This is an alpha field and requires enabling the DRADeviceTaints feature gate.
	Adminaccess bool `json:"adminAccess,omitempty"` // AdminAccess indicates that this device was allocated for administrative access. See the corresponding request field for a definition of mode. This is an alpha field and requires enabling the DRAAdminAccess feature gate. Admin access is disabled if this field is unset or set to false, otherwise it is enabled.
	Device string `json:"device"` // Device references one device instance via its name in the driver's resource pool. It must be a DNS label.
	Driver string `json:"driver"` // Driver specifies the name of the DRA driver whose kubelet plugin should be invoked to process the allocation once the claim is needed on a node. Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver.
	Pool string `json:"pool"` // This name together with the driver name and the device name field identify which device was allocated (`<driver name>/<pool name>/<device name>`). Must not be longer than 253 characters and may contain one or more DNS sub-domains separated by slashes.
}

// Iok8sapiadmissionregistrationv1MutatingWebhookConfiguration represents the Iok8sapiadmissionregistrationv1MutatingWebhookConfiguration schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1MutatingWebhookConfiguration struct {
	Webhooks []Iok8sapiadmissionregistrationv1MutatingWebhook `json:"webhooks,omitempty"` // Webhooks is a list of webhooks and the affected resources and operations.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1PodDNSConfig represents the Iok8sapicorev1PodDNSConfig schema from the OpenAPI specification
type Iok8sapicorev1PodDNSConfig struct {
	Searches []string `json:"searches,omitempty"` // A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.
	Nameservers []string `json:"nameservers,omitempty"` // A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.
	Options []Iok8sapicorev1PodDNSConfigOption `json:"options,omitempty"` // A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.
}

// Iok8sapicorev1EnvVarSource represents the Iok8sapicorev1EnvVarSource schema from the OpenAPI specification
type Iok8sapicorev1EnvVarSource struct {
	Configmapkeyref Iok8sapicorev1ConfigMapKeySelector `json:"configMapKeyRef,omitempty"` // Selects a key from a ConfigMap.
	Fieldref Iok8sapicorev1ObjectFieldSelector `json:"fieldRef,omitempty"` // ObjectFieldSelector selects an APIVersioned field of an object.
	Resourcefieldref Iok8sapicorev1ResourceFieldSelector `json:"resourceFieldRef,omitempty"` // ResourceFieldSelector represents container resources (cpu, memory) and their output format
	Secretkeyref Iok8sapicorev1SecretKeySelector `json:"secretKeyRef,omitempty"` // SecretKeySelector selects a key of a Secret.
}

// Iok8sapicorev1NodeRuntimeHandlerFeatures represents the Iok8sapicorev1NodeRuntimeHandlerFeatures schema from the OpenAPI specification
type Iok8sapicorev1NodeRuntimeHandlerFeatures struct {
	Usernamespaces bool `json:"userNamespaces,omitempty"` // UserNamespaces is set to true if the runtime handler supports UserNamespaces, including for volumes.
	Recursivereadonlymounts bool `json:"recursiveReadOnlyMounts,omitempty"` // RecursiveReadOnlyMounts is set to true if the runtime handler supports RecursiveReadOnlyMounts.
}

// Iok8sapicorev1LinuxContainerUser represents the Iok8sapicorev1LinuxContainerUser schema from the OpenAPI specification
type Iok8sapicorev1LinuxContainerUser struct {
	Gid int64 `json:"gid"` // GID is the primary gid initially attached to the first process in the container
	Supplementalgroups []int64 `json:"supplementalGroups,omitempty"` // SupplementalGroups are the supplemental groups initially attached to the first process in the container
	Uid int64 `json:"uid"` // UID is the primary uid initially attached to the first process in the container
}

// Iok8sapicorev1VolumeMount represents the Iok8sapicorev1VolumeMount schema from the OpenAPI specification
type Iok8sapicorev1VolumeMount struct {
	Recursivereadonly string `json:"recursiveReadOnly,omitempty"` // RecursiveReadOnly specifies whether read-only mounts should be handled recursively. If ReadOnly is false, this field has no meaning and must be unspecified. If ReadOnly is true, and this field is set to Disabled, the mount is not made recursively read-only. If this field is set to IfPossible, the mount is made recursively read-only, if it is supported by the container runtime. If this field is set to Enabled, the mount is made recursively read-only if it is supported by the container runtime, otherwise the pod will not be started and an error will be generated to indicate the reason. If this field is set to IfPossible or Enabled, MountPropagation must be set to None (or be unspecified, which defaults to None). If this field is not specified, it is treated as an equivalent of Disabled.
	Subpath string `json:"subPath,omitempty"` // Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).
	Subpathexpr string `json:"subPathExpr,omitempty"` // Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to "" (volume's root). SubPathExpr and SubPath are mutually exclusive.
	Mountpath string `json:"mountPath"` // Path within the container at which the volume should be mounted. Must not contain ':'.
	Mountpropagation string `json:"mountPropagation,omitempty"` // mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified (which defaults to None).
	Name string `json:"name"` // This must match the Name of a Volume.
	Readonly bool `json:"readOnly,omitempty"` // Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
}

// Iok8sapiresourcev1beta2ResourceSlice represents the Iok8sapiresourcev1beta2ResourceSlice schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ResourceSlice struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiresourcev1beta2ResourceSliceSpec `json:"spec"` // ResourceSliceSpec contains the information published by the driver in one ResourceSlice.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapibatchv1JobStatus represents the Iok8sapibatchv1JobStatus schema from the OpenAPI specification
type Iok8sapibatchv1JobStatus struct {
	Failedindexes string `json:"failedIndexes,omitempty"` // FailedIndexes holds the failed indexes when spec.backoffLimitPerIndex is set. The indexes are represented in the text format analogous as for the `completedIndexes` field, ie. they are kept as decimal integers separated by commas. The numbers are listed in increasing order. Three or more consecutive numbers are compressed and represented by the first and last element of the series, separated by a hyphen. For example, if the failed indexes are 1, 3, 4, 5 and 7, they are represented as "1,3-5,7". The set of failed indexes cannot overlap with the set of completed indexes.
	Uncountedterminatedpods Iok8sapibatchv1UncountedTerminatedPods `json:"uncountedTerminatedPods,omitempty"` // UncountedTerminatedPods holds UIDs of Pods that have terminated but haven't been accounted in Job status counters.
	Failed int `json:"failed,omitempty"` // The number of pods which reached phase Failed. The value increases monotonically.
	Active int `json:"active,omitempty"` // The number of pending and running pods which are not terminating (without a deletionTimestamp). The value is zero for finished jobs.
	Conditions []Iok8sapibatchv1JobCondition `json:"conditions,omitempty"` // The latest available observations of an object's current state. When a Job fails, one of the conditions will have type "Failed" and status true. When a Job is suspended, one of the conditions will have type "Suspended" and status true; when the Job is resumed, the status of this condition will become false. When a Job is completed, one of the conditions will have type "Complete" and status true. A job is considered finished when it is in a terminal condition, either "Complete" or "Failed". A Job cannot have both the "Complete" and "Failed" conditions. Additionally, it cannot be in the "Complete" and "FailureTarget" conditions. The "Complete", "Failed" and "FailureTarget" conditions cannot be disabled. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Ready int `json:"ready,omitempty"` // The number of active pods which have a Ready condition and are not terminating (without a deletionTimestamp).
	Starttime string `json:"startTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Succeeded int `json:"succeeded,omitempty"` // The number of pods which reached phase Succeeded. The value increases monotonically for a given spec. However, it may decrease in reaction to scale down of elastic indexed jobs.
	Completiontime string `json:"completionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Terminating int `json:"terminating,omitempty"` // The number of pods which are terminating (in phase Pending or Running and have a deletionTimestamp). This field is beta-level. The job controller populates the field when the feature gate JobPodReplacementPolicy is enabled (enabled by default).
	Completedindexes string `json:"completedIndexes,omitempty"` // completedIndexes holds the completed indexes when .spec.completionMode = "Indexed" in a text format. The indexes are represented as decimal integers separated by commas. The numbers are listed in increasing order. Three or more consecutive numbers are compressed and represented by the first and last element of the series, separated by a hyphen. For example, if the completed indexes are 1, 3, 4, 5 and 7, they are represented as "1,3-5,7".
}

// Iok8sapicorev1DaemonEndpoint represents the Iok8sapicorev1DaemonEndpoint schema from the OpenAPI specification
type Iok8sapicorev1DaemonEndpoint struct {
	Port int `json:"Port"` // Port number of the given endpoint.
}

// Iok8sapicorev1LocalObjectReference represents the Iok8sapicorev1LocalObjectReference schema from the OpenAPI specification
type Iok8sapicorev1LocalObjectReference struct {
	Name string `json:"name,omitempty"` // Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
}

// Iok8sapicorev1ResourceHealth represents the Iok8sapicorev1ResourceHealth schema from the OpenAPI specification
type Iok8sapicorev1ResourceHealth struct {
	Health string `json:"health,omitempty"` // Health of the resource. can be one of: - Healthy: operates as normal - Unhealthy: reported unhealthy. We consider this a temporary health issue since we do not have a mechanism today to distinguish temporary and permanent issues. - Unknown: The status cannot be determined. For example, Device Plugin got unregistered and hasn't been re-registered since. In future we may want to introduce the PermanentlyUnhealthy Status.
	Resourceid string `json:"resourceID"` // ResourceID is the unique identifier of the resource. See the ResourceID type for more information.
}

// Iok8sapiflowcontrolv1FlowDistinguisherMethod represents the Iok8sapiflowcontrolv1FlowDistinguisherMethod schema from the OpenAPI specification
type Iok8sapiflowcontrolv1FlowDistinguisherMethod struct {
	TypeField string `json:"type"` // `type` is the type of flow distinguisher method The supported types are "ByUser" and "ByNamespace". Required.
}

// Iok8sapicorev1StorageOSPersistentVolumeSource represents the Iok8sapicorev1StorageOSPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1StorageOSPersistentVolumeSource struct {
	Secretref Iok8sapicorev1ObjectReference `json:"secretRef,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Volumename string `json:"volumeName,omitempty"` // volumeName is the human-readable name of the StorageOS volume. Volume names are only unique within a namespace.
	Volumenamespace string `json:"volumeNamespace,omitempty"` // volumeNamespace specifies the scope of the volume within StorageOS. If no namespace is specified then the Pod's namespace will be used. This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Readonly bool `json:"readOnly,omitempty"` // readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
}

// Iok8sapicorev1NodeCondition represents the Iok8sapicorev1NodeCondition schema from the OpenAPI specification
type Iok8sapicorev1NodeCondition struct {
	Lastheartbeattime string `json:"lastHeartbeatTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // Human readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // (brief) reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of node condition.
}

// Iok8sapiappsv1StatefulSetOrdinals represents the Iok8sapiappsv1StatefulSetOrdinals schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSetOrdinals struct {
	Start int `json:"start,omitempty"` // start is the number representing the first replica's index. It may be used to number replicas from an alternate index (eg: 1-indexed) over the default 0-indexed names, or to orchestrate progressive movement of replicas from one StatefulSet to another. If set, replica indices will be in the range: [.spec.ordinals.start, .spec.ordinals.start + .spec.replicas). If unset, defaults to 0. Replica indices will be in the range: [0, .spec.replicas).
}

// Iok8sapinetworkingv1IngressClassList represents the Iok8sapinetworkingv1IngressClassList schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressClassList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapinetworkingv1IngressClass `json:"items"` // items is the list of IngressClasses.
}

// Iok8sapicorev1NodeSelector represents the Iok8sapicorev1NodeSelector schema from the OpenAPI specification
type Iok8sapicorev1NodeSelector struct {
	Nodeselectorterms []Iok8sapicorev1NodeSelectorTerm `json:"nodeSelectorTerms"` // Required. A list of node selector terms. The terms are ORed.
}

// Iok8sapiappsv1ReplicaSet represents the Iok8sapiappsv1ReplicaSet schema from the OpenAPI specification
type Iok8sapiappsv1ReplicaSet struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1ReplicaSetSpec `json:"spec,omitempty"` // ReplicaSetSpec is the specification of a ReplicaSet.
	Status Iok8sapiappsv1ReplicaSetStatus `json:"status,omitempty"` // ReplicaSetStatus represents the current status of a ReplicaSet.
}

// Iok8sapinetworkingv1beta1IPAddressList represents the Iok8sapinetworkingv1beta1IPAddressList schema from the OpenAPI specification
type Iok8sapinetworkingv1beta1IPAddressList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapinetworkingv1beta1IPAddress `json:"items"` // items is the list of IPAddresses.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1NodeRuntimeHandler represents the Iok8sapicorev1NodeRuntimeHandler schema from the OpenAPI specification
type Iok8sapicorev1NodeRuntimeHandler struct {
	Name string `json:"name,omitempty"` // Runtime handler name. Empty for the default runtime handler.
	Features Iok8sapicorev1NodeRuntimeHandlerFeatures `json:"features,omitempty"` // NodeRuntimeHandlerFeatures is a set of features implemented by the runtime handler.
}

// Iok8sapicoordinationv1Lease represents the Iok8sapicoordinationv1Lease schema from the OpenAPI specification
type Iok8sapicoordinationv1Lease struct {
	Spec Iok8sapicoordinationv1LeaseSpec `json:"spec,omitempty"` // LeaseSpec is a specification of a Lease.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1ContainerResizePolicy represents the Iok8sapicorev1ContainerResizePolicy schema from the OpenAPI specification
type Iok8sapicorev1ContainerResizePolicy struct {
	Resourcename string `json:"resourceName"` // Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.
	Restartpolicy string `json:"restartPolicy"` // Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.
}

// Iok8sapiappsv1ReplicaSetList represents the Iok8sapiappsv1ReplicaSetList schema from the OpenAPI specification
type Iok8sapiappsv1ReplicaSetList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiappsv1ReplicaSet `json:"items"` // List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicaset
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1EndpointAddress represents the Iok8sapicorev1EndpointAddress schema from the OpenAPI specification
type Iok8sapicorev1EndpointAddress struct {
	Targetref Iok8sapicorev1ObjectReference `json:"targetRef,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Hostname string `json:"hostname,omitempty"` // The Hostname of this endpoint
	Ip string `json:"ip"` // The IP of this endpoint. May not be loopback (127.0.0.0/8 or ::1), link-local (169.254.0.0/16 or fe80::/10), or link-local multicast (224.0.0.0/24 or ff02::/16).
	Nodename string `json:"nodeName,omitempty"` // Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.
}

// Iok8sapiautoscalingv2HPAScalingPolicy represents the Iok8sapiautoscalingv2HPAScalingPolicy schema from the OpenAPI specification
type Iok8sapiautoscalingv2HPAScalingPolicy struct {
	Periodseconds int `json:"periodSeconds"` // periodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).
	TypeField string `json:"type"` // type is used to specify the scaling policy.
	Value int `json:"value"` // value contains the amount of change which is permitted by the policy. It must be greater than zero
}

// Iok8sapiflowcontrolv1QueuingConfiguration represents the Iok8sapiflowcontrolv1QueuingConfiguration schema from the OpenAPI specification
type Iok8sapiflowcontrolv1QueuingConfiguration struct {
	Queues int `json:"queues,omitempty"` // `queues` is the number of queues for this priority level. The queues exist independently at each apiserver. The value must be positive. Setting it to 1 effectively precludes shufflesharding and thus makes the distinguisher method of associated flow schemas irrelevant. This field has a default value of 64.
	Handsize int `json:"handSize,omitempty"` // `handSize` is a small positive number that configures the shuffle sharding of requests into queues. When enqueuing a request at this priority level the request's flow identifier (a string pair) is hashed and the hash value is used to shuffle the list of queues and deal a hand of the size specified here. The request is put into one of the shortest queues in that hand. `handSize` must be no larger than `queues`, and should be significantly smaller (so that a few heavy flows do not saturate most of the queues). See the user-facing documentation for more extensive guidance on setting this field. This field has a default value of 8.
	Queuelengthlimit int `json:"queueLengthLimit,omitempty"` // `queueLengthLimit` is the maximum number of requests allowed to be waiting in a given queue of this priority level at a time; excess requests are rejected. This value must be positive. If not specified, it will be defaulted to 50.
}

// Iok8sapicorev1PhotonPersistentDiskVolumeSource represents the Iok8sapicorev1PhotonPersistentDiskVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1PhotonPersistentDiskVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Pdid string `json:"pdID"` // pdID is the ID that identifies Photon Controller persistent disk
}

// Iok8sapicorev1NodeSystemInfo represents the Iok8sapicorev1NodeSystemInfo schema from the OpenAPI specification
type Iok8sapicorev1NodeSystemInfo struct {
	Kubeletversion string `json:"kubeletVersion"` // Kubelet Version reported by the node.
	Kernelversion string `json:"kernelVersion"` // Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	Containerruntimeversion string `json:"containerRuntimeVersion"` // ContainerRuntime Version reported by the node through runtime remote API (e.g. containerd://1.4.2).
	Bootid string `json:"bootID"` // Boot ID reported by the node.
	Kubeproxyversion string `json:"kubeProxyVersion"` // Deprecated: KubeProxy Version reported by the node.
	Machineid string `json:"machineID"` // MachineID reported by the node. For unique machine identification in the cluster this field is preferred. Learn more from man(5) machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html
	Operatingsystem string `json:"operatingSystem"` // The Operating System reported by the node
	Osimage string `json:"osImage"` // OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	Swap Iok8sapicorev1NodeSwapStatus `json:"swap,omitempty"` // NodeSwapStatus represents swap memory information.
	Systemuuid string `json:"systemUUID"` // SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-us/red_hat_subscription_management/1/html/rhsm/uuid
	Architecture string `json:"architecture"` // The Architecture reported by the node
}

// Iok8sapistoragev1CSINodeList represents the Iok8sapistoragev1CSINodeList schema from the OpenAPI specification
type Iok8sapistoragev1CSINodeList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapistoragev1CSINode `json:"items"` // items is the list of CSINode
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cabundle string `json:"caBundle,omitempty"` // caBundle is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.
	Service GeneratedType `json:"service,omitempty"` // ServiceReference holds a reference to Service.legacy.k8s.io
	Url string `json:"url,omitempty"` // url gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified. The `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address. Please note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster. The scheme must be "https"; the URL must begin with "https://". A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier. Attempting to use a user or basic auth e.g. "user:password@" is not allowed. Fragments ("#...") and query parameters ("?...") are not allowed, either.
}

// Iok8sapiauthenticationv1TokenRequestSpec represents the Iok8sapiauthenticationv1TokenRequestSpec schema from the OpenAPI specification
type Iok8sapiauthenticationv1TokenRequestSpec struct {
	Boundobjectref Iok8sapiauthenticationv1BoundObjectReference `json:"boundObjectRef,omitempty"` // BoundObjectReference is a reference to an object that a token is bound to.
	Expirationseconds int64 `json:"expirationSeconds,omitempty"` // ExpirationSeconds is the requested duration of validity of the request. The token issuer may return a token with a different validity duration so a client needs to check the 'expiration' field in a response.
	Audiences []string `json:"audiences"` // Audiences are the intendend audiences of the token. A recipient of a token must identify themself with an identifier in the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be used to authenticate against any of the audiences listed but implies a high degree of trust between the target audiences.
}

// Iok8sapicorev1LoadBalancerStatus represents the Iok8sapicorev1LoadBalancerStatus schema from the OpenAPI specification
type Iok8sapicorev1LoadBalancerStatus struct {
	Ingress []Iok8sapicorev1LoadBalancerIngress `json:"ingress,omitempty"` // Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points.
}

// Iok8sapicorev1NamespaceSpec represents the Iok8sapicorev1NamespaceSpec schema from the OpenAPI specification
type Iok8sapicorev1NamespaceSpec struct {
	Finalizers []string `json:"finalizers,omitempty"` // Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
}

// Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyBindingList represents the Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyBindingList schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyBindingList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyBinding `json:"items"` // List of PolicyBinding.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapinetworkingv1IngressLoadBalancerIngress represents the Iok8sapinetworkingv1IngressLoadBalancerIngress schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressLoadBalancerIngress struct {
	Hostname string `json:"hostname,omitempty"` // hostname is set for load-balancer ingress points that are DNS based.
	Ip string `json:"ip,omitempty"` // ip is set for load-balancer ingress points that are IP based.
	Ports []Iok8sapinetworkingv1IngressPortStatus `json:"ports,omitempty"` // ports provides information about the ports exposed by this LoadBalancer.
}

// Iok8sapinetworkingv1IPAddressList represents the Iok8sapinetworkingv1IPAddressList schema from the OpenAPI specification
type Iok8sapinetworkingv1IPAddressList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapinetworkingv1IPAddress `json:"items"` // items is the list of IPAddresses.
}

// Iok8sapiauthorizationv1SelfSubjectRulesReview represents the Iok8sapiauthorizationv1SelfSubjectRulesReview schema from the OpenAPI specification
type Iok8sapiauthorizationv1SelfSubjectRulesReview struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1SelfSubjectRulesReviewSpec `json:"spec"` // SelfSubjectRulesReviewSpec defines the specification for SelfSubjectRulesReview.
	Status Iok8sapiauthorizationv1SubjectRulesReviewStatus `json:"status,omitempty"` // SubjectRulesReviewStatus contains the result of a rules check. This check can be incomplete depending on the set of authorizers the server is configured with and any errors experienced during evaluation. Because authorization rules are additive, if a rule appears in a list it's safe to assume the subject has that permission, even if that list is incomplete.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1SecretReference represents the Iok8sapicorev1SecretReference schema from the OpenAPI specification
type Iok8sapicorev1SecretReference struct {
	Name string `json:"name,omitempty"` // name is unique within a namespace to reference a secret resource.
	Namespace string `json:"namespace,omitempty"` // namespace defines the space within which the secret name must be unique.
}

// Iok8sapiadmissionregistrationv1ValidatingWebhookConfigurationList represents the Iok8sapiadmissionregistrationv1ValidatingWebhookConfigurationList schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ValidatingWebhookConfigurationList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiadmissionregistrationv1ValidatingWebhookConfiguration `json:"items"` // List of ValidatingWebhookConfiguration.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1EventSource represents the Iok8sapicorev1EventSource schema from the OpenAPI specification
type Iok8sapicorev1EventSource struct {
	Component string `json:"component,omitempty"` // Component from which the event is generated.
	Host string `json:"host,omitempty"` // Node name on which the event is generated.
}

// Iok8sapiresourcev1beta1DeviceRequest represents the Iok8sapiresourcev1beta1DeviceRequest schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceRequest struct {
	Firstavailable []Iok8sapiresourcev1beta1DeviceSubRequest `json:"firstAvailable,omitempty"` // FirstAvailable contains subrequests, of which exactly one will be satisfied by the scheduler to satisfy this request. It tries to satisfy them in the order in which they are listed here. So if there are two entries in the list, the scheduler will only check the second one if it determines that the first one cannot be used. This field may only be set in the entries of DeviceClaim.Requests. DRA does not yet implement scoring, so the scheduler will select the first set of devices that satisfies all the requests in the claim. And if the requirements can be satisfied on more than one node, other scheduling features will determine which node is chosen. This means that the set of devices allocated to a claim might not be the optimal set available to the cluster. Scoring will be implemented later.
	Name string `json:"name"` // Name can be used to reference this request in a pod.spec.containers[].resources.claims entry and in a constraint of the claim. Must be a DNS label and unique among all DeviceRequests in a ResourceClaim.
	Selectors []Iok8sapiresourcev1beta1DeviceSelector `json:"selectors,omitempty"` // Selectors define criteria which must be satisfied by a specific device in order for that device to be considered for this request. All selectors must be satisfied for a device to be considered. This field can only be set when deviceClassName is set and no subrequests are specified in the firstAvailable list.
	Tolerations []Iok8sapiresourcev1beta1DeviceToleration `json:"tolerations,omitempty"` // If specified, the request's tolerations. Tolerations for NoSchedule are required to allocate a device which has a taint with that effect. The same applies to NoExecute. In addition, should any of the allocated devices get tainted with NoExecute after allocation and that effect is not tolerated, then all pods consuming the ResourceClaim get deleted to evict them. The scheduler will not let new pods reserve the claim while it has these tainted devices. Once all pods are evicted, the claim will get deallocated. The maximum number of tolerations is 16. This field can only be set when deviceClassName is set and no subrequests are specified in the firstAvailable list. This is an alpha field and requires enabling the DRADeviceTaints feature gate.
	Adminaccess bool `json:"adminAccess,omitempty"` // AdminAccess indicates that this is a claim for administrative access to the device(s). Claims with AdminAccess are expected to be used for monitoring or other management services for a device. They ignore all ordinary claims to the device with respect to access modes and any resource allocations. This field can only be set when deviceClassName is set and no subrequests are specified in the firstAvailable list. This is an alpha field and requires enabling the DRAAdminAccess feature gate. Admin access is disabled if this field is unset or set to false, otherwise it is enabled.
	Allocationmode string `json:"allocationMode,omitempty"` // AllocationMode and its related fields define how devices are allocated to satisfy this request. Supported values are: - ExactCount: This request is for a specific number of devices. This is the default. The exact number is provided in the count field. - All: This request is for all of the matching devices in a pool. At least one device must exist on the node for the allocation to succeed. Allocation will fail if some devices are already allocated, unless adminAccess is requested. If AllocationMode is not specified, the default mode is ExactCount. If the mode is ExactCount and count is not specified, the default count is one. Any other requests must specify this field. This field can only be set when deviceClassName is set and no subrequests are specified in the firstAvailable list. More modes may get added in the future. Clients must refuse to handle requests with unknown modes.
	Count int64 `json:"count,omitempty"` // Count is used only when the count mode is "ExactCount". Must be greater than zero. If AllocationMode is ExactCount and this field is not specified, the default is one. This field can only be set when deviceClassName is set and no subrequests are specified in the firstAvailable list.
	Deviceclassname string `json:"deviceClassName,omitempty"` // DeviceClassName references a specific DeviceClass, which can define additional configuration and selectors to be inherited by this request. A class is required if no subrequests are specified in the firstAvailable list and no class can be set if subrequests are specified in the firstAvailable list. Which classes are available depends on the cluster. Administrators may use this to restrict which devices may get requested by only installing classes with selectors for permitted devices. If users are free to request anything without restrictions, then administrators can create an empty DeviceClass for users to reference.
}

// Iok8sapiresourcev1beta1Counter represents the Iok8sapiresourcev1beta1Counter schema from the OpenAPI specification
type Iok8sapiresourcev1beta1Counter struct {
	Value string `json:"value"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
}

// Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyList represents the Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyList schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicy `json:"items"` // List of ValidatingAdmissionPolicy.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Kind string `json:"kind"` // kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use this value as the `kind` attribute in API calls.
	Listkind string `json:"listKind,omitempty"` // listKind is the serialized kind of the list for this resource. Defaults to "`kind`List".
	Plural string `json:"plural"` // plural is the plural name of the resource to serve. The custom resources are served under `/apis/<group>/<version>/.../<plural>`. Must match the name of the CustomResourceDefinition (in the form `<names.plural>.<group>`). Must be all lowercase.
	Shortnames []string `json:"shortNames,omitempty"` // shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like `kubectl get <shortname>`. It must be all lowercase.
	Singular string `json:"singular,omitempty"` // singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased `kind`.
	Categories []string `json:"categories,omitempty"` // categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API discovery documents, and used by clients to support invocations like `kubectl get all`.
}

// Iok8sapinetworkingv1NetworkPolicyPeer represents the Iok8sapinetworkingv1NetworkPolicyPeer schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicyPeer struct {
	Podselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"podSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Ipblock Iok8sapinetworkingv1IPBlock `json:"ipBlock,omitempty"` // IPBlock describes a particular CIDR (Ex. "192.168.1.0/24","2001:db8::/64") that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.
	Namespaceselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"namespaceSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapiresourcev1alpha3DeviceTaintRuleList represents the Iok8sapiresourcev1alpha3DeviceTaintRuleList schema from the OpenAPI specification
type Iok8sapiresourcev1alpha3DeviceTaintRuleList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiresourcev1alpha3DeviceTaintRule `json:"items"` // Items is the list of DeviceTaintRules.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1SecretVolumeSource represents the Iok8sapicorev1SecretVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1SecretVolumeSource struct {
	Defaultmode int `json:"defaultMode,omitempty"` // defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Items []Iok8sapicorev1KeyToPath `json:"items,omitempty"` // items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
	Optional bool `json:"optional,omitempty"` // optional field specify whether the Secret or its keys must be defined
	Secretname string `json:"secretName,omitempty"` // secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
}

// Iok8sapicorev1ServiceAccountTokenProjection represents the Iok8sapicorev1ServiceAccountTokenProjection schema from the OpenAPI specification
type Iok8sapicorev1ServiceAccountTokenProjection struct {
	Audience string `json:"audience,omitempty"` // audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.
	Expirationseconds int64 `json:"expirationSeconds,omitempty"` // expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.
	Path string `json:"path"` // path is the path relative to the mount point of the file to project the token into.
}

// Iok8sapicorev1KeyToPath represents the Iok8sapicorev1KeyToPath schema from the OpenAPI specification
type Iok8sapicorev1KeyToPath struct {
	Path string `json:"path"` // path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
	Key string `json:"key"` // key is the key to project.
	Mode int `json:"mode,omitempty"` // mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
}

// Iok8sapicorev1PersistentVolumeClaimSpec represents the Iok8sapicorev1PersistentVolumeClaimSpec schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaimSpec struct {
	Storageclassname string `json:"storageClassName,omitempty"` // storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	Accessmodes []string `json:"accessModes,omitempty"` // accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	Datasource Iok8sapicorev1TypedLocalObjectReference `json:"dataSource,omitempty"` // TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.
	Resources Iok8sapicorev1VolumeResourceRequirements `json:"resources,omitempty"` // VolumeResourceRequirements describes the storage resource requirements for a volume.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Datasourceref Iok8sapicorev1TypedObjectReference `json:"dataSourceRef,omitempty"` // TypedObjectReference contains enough information to let you locate the typed referenced object
	Volumeattributesclassname string `json:"volumeAttributesClassName,omitempty"` // volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim. If specified, the CSI driver will create or update the volume with the attributes defined in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName, it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass will be applied to the claim but it's not allowed to reset this field to empty string once it is set. If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass will be set by the persistentvolume controller if it exists. If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource exists. More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/ (Beta) Using this field requires the VolumeAttributesClass feature gate to be enabled (off by default).
	Volumemode string `json:"volumeMode,omitempty"` // volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
	Volumename string `json:"volumeName,omitempty"` // volumeName is the binding reference to the PersistentVolume backing this claim.
}

// Iok8sapicorev1ConfigMapEnvSource represents the Iok8sapicorev1ConfigMapEnvSource schema from the OpenAPI specification
type Iok8sapicorev1ConfigMapEnvSource struct {
	Name string `json:"name,omitempty"` // Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the ConfigMap must be defined
}

// Iok8sapiresourcev1beta1DeviceSubRequest represents the Iok8sapiresourcev1beta1DeviceSubRequest schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceSubRequest struct {
	Tolerations []Iok8sapiresourcev1beta1DeviceToleration `json:"tolerations,omitempty"` // If specified, the request's tolerations. Tolerations for NoSchedule are required to allocate a device which has a taint with that effect. The same applies to NoExecute. In addition, should any of the allocated devices get tainted with NoExecute after allocation and that effect is not tolerated, then all pods consuming the ResourceClaim get deleted to evict them. The scheduler will not let new pods reserve the claim while it has these tainted devices. Once all pods are evicted, the claim will get deallocated. The maximum number of tolerations is 16. This is an alpha field and requires enabling the DRADeviceTaints feature gate.
	Allocationmode string `json:"allocationMode,omitempty"` // AllocationMode and its related fields define how devices are allocated to satisfy this subrequest. Supported values are: - ExactCount: This request is for a specific number of devices. This is the default. The exact number is provided in the count field. - All: This subrequest is for all of the matching devices in a pool. Allocation will fail if some devices are already allocated, unless adminAccess is requested. If AllocationMode is not specified, the default mode is ExactCount. If the mode is ExactCount and count is not specified, the default count is one. Any other subrequests must specify this field. More modes may get added in the future. Clients must refuse to handle requests with unknown modes.
	Count int64 `json:"count,omitempty"` // Count is used only when the count mode is "ExactCount". Must be greater than zero. If AllocationMode is ExactCount and this field is not specified, the default is one.
	Deviceclassname string `json:"deviceClassName"` // DeviceClassName references a specific DeviceClass, which can define additional configuration and selectors to be inherited by this subrequest. A class is required. Which classes are available depends on the cluster. Administrators may use this to restrict which devices may get requested by only installing classes with selectors for permitted devices. If users are free to request anything without restrictions, then administrators can create an empty DeviceClass for users to reference.
	Name string `json:"name"` // Name can be used to reference this subrequest in the list of constraints or the list of configurations for the claim. References must use the format <main request>/<subrequest>. Must be a DNS label.
	Selectors []Iok8sapiresourcev1beta1DeviceSelector `json:"selectors,omitempty"` // Selectors define criteria which must be satisfied by a specific device in order for that device to be considered for this subrequest. All selectors must be satisfied for a device to be considered.
}

// Iok8sapiauthorizationv1ResourceRule represents the Iok8sapiauthorizationv1ResourceRule schema from the OpenAPI specification
type Iok8sapiauthorizationv1ResourceRule struct {
	Resourcenames []string `json:"resourceNames,omitempty"` // ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed. "*" means all.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. "*" means all in the specified apiGroups. "*/foo" represents the subresource 'foo' for all resources in the specified apiGroups.
	Verbs []string `json:"verbs"` // Verb is a list of kubernetes resource API verbs, like: get, list, watch, create, update, delete, proxy. "*" means all.
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the name of the APIGroup that contains the resources. If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed. "*" means all.
}

// Iok8sapiadmissionregistrationv1alpha1Mutation represents the Iok8sapiadmissionregistrationv1alpha1Mutation schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1Mutation struct {
	Jsonpatch Iok8sapiadmissionregistrationv1alpha1JSONPatch `json:"jsonPatch,omitempty"` // JSONPatch defines a JSON Patch.
	Patchtype string `json:"patchType"` // patchType indicates the patch strategy used. Allowed values are "ApplyConfiguration" and "JSONPatch". Required.
	Applyconfiguration Iok8sapiadmissionregistrationv1alpha1ApplyConfiguration `json:"applyConfiguration,omitempty"` // ApplyConfiguration defines the desired configuration values of an object.
}

// Iok8sapicorev1ComponentStatusList represents the Iok8sapicorev1ComponentStatusList schema from the OpenAPI specification
type Iok8sapicorev1ComponentStatusList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1ComponentStatus `json:"items"` // List of ComponentStatus objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiappsv1DaemonSetCondition represents the Iok8sapiappsv1DaemonSetCondition schema from the OpenAPI specification
type Iok8sapiappsv1DaemonSetCondition struct {
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of DaemonSet condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
}

// Iok8sapiauthorizationv1SubjectRulesReviewStatus represents the Iok8sapiauthorizationv1SubjectRulesReviewStatus schema from the OpenAPI specification
type Iok8sapiauthorizationv1SubjectRulesReviewStatus struct {
	Evaluationerror string `json:"evaluationError,omitempty"` // EvaluationError can appear in combination with Rules. It indicates an error occurred during rule evaluation, such as an authorizer that doesn't support rule evaluation, and that ResourceRules and/or NonResourceRules may be incomplete.
	Incomplete bool `json:"incomplete"` // Incomplete is true when the rules returned by this call are incomplete. This is most commonly encountered when an authorizer, such as an external authorizer, doesn't support rules evaluation.
	Nonresourcerules []Iok8sapiauthorizationv1NonResourceRule `json:"nonResourceRules"` // NonResourceRules is the list of actions the subject is allowed to perform on non-resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
	Resourcerules []Iok8sapiauthorizationv1ResourceRule `json:"resourceRules"` // ResourceRules is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"` // Type is the type of the condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // Human-readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // Unique, one-word, CamelCase reason for the condition's last transition.
	Status string `json:"status"` // Status is the status of the condition. Can be True, False, Unknown.
}

// Iok8sapicorev1Service represents the Iok8sapicorev1Service schema from the OpenAPI specification
type Iok8sapicorev1Service struct {
	Status Iok8sapicorev1ServiceStatus `json:"status,omitempty"` // ServiceStatus represents the current status of a service.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1ServiceSpec `json:"spec,omitempty"` // ServiceSpec describes the attributes that a user creates on a service.
}

// Iok8sapimachinerypkgversionInfo represents the Iok8sapimachinerypkgversionInfo schema from the OpenAPI specification
type Iok8sapimachinerypkgversionInfo struct {
	Platform string `json:"platform"`
	Gitcommit string `json:"gitCommit"`
	Minor string `json:"minor"` // Minor is the minor version of the binary version
	Builddate string `json:"buildDate"`
	Emulationminor string `json:"emulationMinor,omitempty"` // EmulationMinor is the minor version of the emulation version
	Gittreestate string `json:"gitTreeState"`
	Mincompatibilityminor string `json:"minCompatibilityMinor,omitempty"` // MinCompatibilityMinor is the minor version of the minimum compatibility version
	Emulationmajor string `json:"emulationMajor,omitempty"` // EmulationMajor is the major version of the emulation version
	Goversion string `json:"goVersion"`
	Mincompatibilitymajor string `json:"minCompatibilityMajor,omitempty"` // MinCompatibilityMajor is the major version of the minimum compatibility version
	Compiler string `json:"compiler"`
	Gitversion string `json:"gitVersion"`
	Major string `json:"major"` // Major is the major version of the binary version
}

// Iok8sapiadmissionregistrationv1alpha1ParamKind represents the Iok8sapiadmissionregistrationv1alpha1ParamKind schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1ParamKind struct {
	Kind string `json:"kind,omitempty"` // Kind is the API kind the resources belong to. Required.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion is the API group version the resources belong to. In format of "group/version". Required.
}

// Iok8sapinetworkingv1IngressClassParametersReference represents the Iok8sapinetworkingv1IngressClassParametersReference schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressClassParametersReference struct {
	Name string `json:"name"` // name is the name of resource being referenced.
	Namespace string `json:"namespace,omitempty"` // namespace is the namespace of the resource being referenced. This field is required when scope is set to "Namespace" and must be unset when scope is set to "Cluster".
	Scope string `json:"scope,omitempty"` // scope represents if this refers to a cluster or namespace scoped resource. This may be set to "Cluster" (default) or "Namespace".
	Apigroup string `json:"apiGroup,omitempty"` // apiGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
	Kind string `json:"kind"` // kind is the type of resource being referenced.
}

// Iok8sapiresourcev1beta1OpaqueDeviceConfiguration represents the Iok8sapiresourcev1beta1OpaqueDeviceConfiguration schema from the OpenAPI specification
type Iok8sapiresourcev1beta1OpaqueDeviceConfiguration struct {
	Parameters Iok8sapimachinerypkgruntimeRawExtension `json:"parameters"` // RawExtension is used to hold extensions in external versions. To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types. // Internal package: 	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.Object `json:"myPlugin"` 	} 	type PluginA struct { 		AOption string `json:"aOption"` 	} // External package: 	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.RawExtension `json:"myPlugin"` 	} 	type PluginA struct { 		AOption string `json:"aOption"` 	} // On the wire, the JSON will look something like this: 	{ 		"kind":"MyAPIObject", 		"apiVersion":"v1", 		"myPlugin": { 			"kind":"PluginA", 			"aOption":"foo", 		}, 	} So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)
	Driver string `json:"driver"` // Driver is used to determine which kubelet plugin needs to be passed these configuration parameters. An admission policy provided by the driver developer could use this to decide whether it needs to validate them. Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver.
}

// Iok8sapiautoscalingv2ExternalMetricSource represents the Iok8sapiautoscalingv2ExternalMetricSource schema from the OpenAPI specification
type Iok8sapiautoscalingv2ExternalMetricSource struct {
	Metric Iok8sapiautoscalingv2MetricIdentifier `json:"metric"` // MetricIdentifier defines the name and optionally selector for a metric
	Target Iok8sapiautoscalingv2MetricTarget `json:"target"` // MetricTarget defines the target value, average value, or average utilization of a specific metric
}

// Iok8sapibatchv1CronJobSpec represents the Iok8sapibatchv1CronJobSpec schema from the OpenAPI specification
type Iok8sapibatchv1CronJobSpec struct {
	Startingdeadlineseconds int64 `json:"startingDeadlineSeconds,omitempty"` // Optional deadline in seconds for starting the job if it misses scheduled time for any reason. Missed jobs executions will be counted as failed ones.
	Successfuljobshistorylimit int `json:"successfulJobsHistoryLimit,omitempty"` // The number of successful finished jobs to retain. Value must be non-negative integer. Defaults to 3.
	Suspend bool `json:"suspend,omitempty"` // This flag tells the controller to suspend subsequent executions, it does not apply to already started executions. Defaults to false.
	Timezone string `json:"timeZone,omitempty"` // The time zone name for the given schedule, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones. If not specified, this will default to the time zone of the kube-controller-manager process. The set of valid time zone names and the time zone offset is loaded from the system-wide time zone database by the API server during CronJob validation and the controller manager during execution. If no system-wide time zone database can be found a bundled version of the database is used instead. If the time zone name becomes invalid during the lifetime of a CronJob or due to a change in host configuration, the controller will stop creating new new Jobs and will create a system event with the reason UnknownTimeZone. More information can be found in https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#time-zones
	Concurrencypolicy string `json:"concurrencyPolicy,omitempty"` // Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
	Failedjobshistorylimit int `json:"failedJobsHistoryLimit,omitempty"` // The number of failed finished jobs to retain. Value must be non-negative integer. Defaults to 1.
	Jobtemplate Iok8sapibatchv1JobTemplateSpec `json:"jobTemplate"` // JobTemplateSpec describes the data a Job should have when created from a template
	Schedule string `json:"schedule"` // The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
}

// Iok8sapicorev1PortworxVolumeSource represents the Iok8sapicorev1PortworxVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1PortworxVolumeSource struct {
	Volumeid string `json:"volumeID"` // volumeID uniquely identifies a Portworx volume
	Fstype string `json:"fsType,omitempty"` // fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
	Readonly bool `json:"readOnly,omitempty"` // readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
}

// Iok8sapicertificatesv1CertificateSigningRequestSpec represents the Iok8sapicertificatesv1CertificateSigningRequestSpec schema from the OpenAPI specification
type Iok8sapicertificatesv1CertificateSigningRequestSpec struct {
	Usages []string `json:"usages,omitempty"` // usages specifies a set of key usages requested in the issued certificate. Requests for TLS client certificates typically request: "digital signature", "key encipherment", "client auth". Requests for TLS serving certificates typically request: "key encipherment", "digital signature", "server auth". Valid values are: "signing", "digital signature", "content commitment", "key encipherment", "key agreement", "data encipherment", "cert sign", "crl sign", "encipher only", "decipher only", "any", "server auth", "client auth", "code signing", "email protection", "s/mime", "ipsec end system", "ipsec tunnel", "ipsec user", "timestamping", "ocsp signing", "microsoft sgc", "netscape sgc"
	Username string `json:"username,omitempty"` // username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
	Expirationseconds int `json:"expirationSeconds,omitempty"` // expirationSeconds is the requested duration of validity of the issued certificate. The certificate signer may issue a certificate with a different validity duration so a client must check the delta between the notBefore and and notAfter fields in the issued certificate to determine the actual duration. The v1.22+ in-tree implementations of the well-known Kubernetes signers will honor this field as long as the requested duration is not greater than the maximum duration they will honor per the --cluster-signing-duration CLI flag to the Kubernetes controller manager. Certificate signers may not honor this field for various reasons: 1. Old signer that is unaware of the field (such as the in-tree implementations prior to v1.22) 2. Signer whose configured maximum is shorter than the requested duration 3. Signer whose configured minimum is longer than the requested duration The minimum valid value for expirationSeconds is 600, i.e. 10 minutes.
	Extra map[string]interface{} `json:"extra,omitempty"` // extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
	Groups []string `json:"groups,omitempty"` // groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
	Request string `json:"request"` // request contains an x509 certificate signing request encoded in a "CERTIFICATE REQUEST" PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded.
	Signername string `json:"signerName"` // signerName indicates the requested signer, and is a qualified name. List/watch requests for CertificateSigningRequests can filter on this field using a "spec.signerName=NAME" fieldSelector. Well-known Kubernetes signers are: 1. "kubernetes.io/kube-apiserver-client": issues client certificates that can be used to authenticate to kube-apiserver. Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the "csrsigning" controller in kube-controller-manager. 2. "kubernetes.io/kube-apiserver-client-kubelet": issues client certificates that kubelets use to authenticate to kube-apiserver. Requests for this signer can be auto-approved by the "csrapproving" controller in kube-controller-manager, and can be issued by the "csrsigning" controller in kube-controller-manager. 3. "kubernetes.io/kubelet-serving" issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely. Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the "csrsigning" controller in kube-controller-manager. More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signers Custom signerNames can also be specified. The signer defines: 1. Trust distribution: how trust (CA bundles) are distributed. 2. Permitted subjects: and behavior when a disallowed subject is requested. 3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested. 4. Required, permitted, or forbidden key usages / extended key usages. 5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin. 6. Whether or not requests for CA certificates are allowed.
	Uid string `json:"uid,omitempty"` // uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
}

// Iok8sapipolicyv1PodDisruptionBudgetStatus represents the Iok8sapipolicyv1PodDisruptionBudgetStatus schema from the OpenAPI specification
type Iok8sapipolicyv1PodDisruptionBudgetStatus struct {
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // Most recent generation observed when updating this PDB status. DisruptionsAllowed and other status information is valid only if observedGeneration equals to PDB's object generation.
	Conditions []Iok8sapimachinerypkgapismetav1Condition `json:"conditions,omitempty"` // Conditions contain conditions for PDB. The disruption controller sets the DisruptionAllowed condition. The following are known values for the reason field (additional reasons could be added in the future): - SyncFailed: The controller encountered an error and wasn't able to compute the number of allowed disruptions. Therefore no disruptions are allowed and the status of the condition will be False. - InsufficientPods: The number of pods are either at or below the number required by the PodDisruptionBudget. No disruptions are allowed and the status of the condition will be False. - SufficientPods: There are more pods than required by the PodDisruptionBudget. The condition will be True, and the number of allowed disruptions are provided by the disruptionsAllowed property.
	Currenthealthy int `json:"currentHealthy"` // current number of healthy pods
	Desiredhealthy int `json:"desiredHealthy"` // minimum desired number of healthy pods
	Disruptedpods map[string]interface{} `json:"disruptedPods,omitempty"` // DisruptedPods contains information about pods whose eviction was processed by the API server eviction subresource handler but has not yet been observed by the PodDisruptionBudget controller. A pod will be in this map from the time when the API server processed the eviction request to the time when the pod is seen by PDB controller as having been marked for deletion (or after a timeout). The key in the map is the name of the pod and the value is the time when the API server processed the eviction request. If the deletion didn't occur and a pod is still there it will be removed from the list automatically by PodDisruptionBudget controller after some time. If everything goes smooth this map should be empty for the most of the time. Large number of entries in the map may indicate problems with pod deletions.
	Disruptionsallowed int `json:"disruptionsAllowed"` // Number of pod disruptions that are currently allowed.
	Expectedpods int `json:"expectedPods"` // total number of pods counted by this disruption budget
}

// Iok8sapicoordinationv1alpha2LeaseCandidate represents the Iok8sapicoordinationv1alpha2LeaseCandidate schema from the OpenAPI specification
type Iok8sapicoordinationv1alpha2LeaseCandidate struct {
	Spec Iok8sapicoordinationv1alpha2LeaseCandidateSpec `json:"spec,omitempty"` // LeaseCandidateSpec is a specification of a Lease.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiautoscalingv2HorizontalPodAutoscalerStatus represents the Iok8sapiautoscalingv2HorizontalPodAutoscalerStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2HorizontalPodAutoscalerStatus struct {
	Currentmetrics []Iok8sapiautoscalingv2MetricStatus `json:"currentMetrics,omitempty"` // currentMetrics is the last read state of the metrics used by this autoscaler.
	Currentreplicas int `json:"currentReplicas,omitempty"` // currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
	Desiredreplicas int `json:"desiredReplicas"` // desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
	Lastscaletime string `json:"lastScaleTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // observedGeneration is the most recent generation observed by this autoscaler.
	Conditions []Iok8sapiautoscalingv2HorizontalPodAutoscalerCondition `json:"conditions,omitempty"` // conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.
}

// Iok8sapiappsv1DaemonSetStatus represents the Iok8sapiappsv1DaemonSetStatus schema from the OpenAPI specification
type Iok8sapiappsv1DaemonSetStatus struct {
	Collisioncount int `json:"collisionCount,omitempty"` // Count of hash collisions for the DaemonSet. The DaemonSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
	Numberready int `json:"numberReady"` // numberReady is the number of nodes that should be running the daemon pod and have one or more of the daemon pod running with a Ready Condition.
	Numberunavailable int `json:"numberUnavailable,omitempty"` // The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds)
	Conditions []Iok8sapiappsv1DaemonSetCondition `json:"conditions,omitempty"` // Represents the latest available observations of a DaemonSet's current state.
	Updatednumberscheduled int `json:"updatedNumberScheduled,omitempty"` // The total number of nodes that are running updated daemon pod
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The most recent generation observed by the daemon set controller.
	Currentnumberscheduled int `json:"currentNumberScheduled"` // The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Desirednumberscheduled int `json:"desiredNumberScheduled"` // The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	Numberavailable int `json:"numberAvailable,omitempty"` // The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds)
	Numbermisscheduled int `json:"numberMisscheduled"` // The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
}

// Iok8sapicorev1ServiceSpec represents the Iok8sapicorev1ServiceSpec schema from the OpenAPI specification
type Iok8sapicorev1ServiceSpec struct {
	Loadbalancerip string `json:"loadBalancerIP,omitempty"` // Only applies to Service Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations. Using it is non-portable and it may not support dual-stack. Users are encouraged to use implementation-specific annotations when available.
	Ipfamilypolicy string `json:"ipFamilyPolicy,omitempty"` // IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be "SingleStack" (a single IP family), "PreferDualStack" (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or "RequireDualStack" (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.
	Ports []Iok8sapicorev1ServicePort `json:"ports,omitempty"` // The list of ports that are exposed by this service. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Externalips []string `json:"externalIPs,omitempty"` // externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system.
	Loadbalancerclass string `json:"loadBalancerClass,omitempty"` // loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. "internal-vip" or "example.com/internal-vip". Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.
	Sessionaffinity string `json:"sessionAffinity,omitempty"` // Supports "ClientIP" and "None". Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Trafficdistribution string `json:"trafficDistribution,omitempty"` // TrafficDistribution offers a way to express preferences for how traffic is distributed to Service endpoints. Implementations can use this field as a hint, but are not required to guarantee strict adherence. If the field is not set, the implementation will apply its default routing strategy. If set to "PreferClose", implementations should prioritize endpoints that are in the same zone.
	TypeField string `json:"type,omitempty"` // type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. "ClusterIP" allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is "None", no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. "NodePort" builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. "LoadBalancer" builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. "ExternalName" aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	Healthchecknodeport int `json:"healthCheckNodePort,omitempty"` // healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used. If not specified, a value will be automatically allocated. External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not. If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type). This field cannot be updated once set.
	Clusterip string `json:"clusterIP,omitempty"` // clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are "None", empty string (""), or a valid IP address. Setting this to "None" makes a "headless service" (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Externaltrafficpolicy string `json:"externalTrafficPolicy,omitempty"` // externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the Service's "externally-facing" addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to "Local", the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, "Cluster", uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get "Cluster" semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node.
	Sessionaffinityconfig Iok8sapicorev1SessionAffinityConfig `json:"sessionAffinityConfig,omitempty"` // SessionAffinityConfig represents the configurations of session affinity.
	Clusterips []string `json:"clusterIPs,omitempty"` // ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are "None", empty string (""), or a valid IP address. Setting this to "None" makes a "headless service" (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. If this field is not specified, it will be initialized from the clusterIP field. If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value. This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Externalname string `json:"externalName,omitempty"` // externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires `type` to be "ExternalName".
	Publishnotreadyaddresses bool `json:"publishNotReadyAddresses,omitempty"` // publishNotReadyAddresses indicates that any agent which deals with endpoints for this Service should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet's Headless Service to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered "ready" even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior.
	Internaltrafficpolicy string `json:"internalTrafficPolicy,omitempty"` // InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to "Local", the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, "Cluster", uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features).
	Ipfamilies []string `json:"ipFamilies,omitempty"` // IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are "IPv4" and "IPv6". This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to "headless" services. This field will be wiped when updating a Service to type ExternalName. This field may hold a maximum of two entries (dual-stack families, in either order). These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.
	Selector map[string]interface{} `json:"selector,omitempty"` // Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/
	Loadbalancersourceranges []string `json:"loadBalancerSourceRanges,omitempty"` // If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature." More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/
	Allocateloadbalancernodeports bool `json:"allocateLoadBalancerNodePorts,omitempty"` // allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer. Default is "true". It may be set to "false" if the cluster load-balancer does not rely on NodePorts. If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name,omitempty"` // Name is the name of the service
	Namespace string `json:"namespace,omitempty"` // Namespace is the namespace of the service
	Port int `json:"port,omitempty"` // If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
}

// Iok8sapicorev1GitRepoVolumeSource represents the Iok8sapicorev1GitRepoVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1GitRepoVolumeSource struct {
	Directory string `json:"directory,omitempty"` // directory is the target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
	Repository string `json:"repository"` // repository is the URL
	Revision string `json:"revision,omitempty"` // revision is the commit hash for the specified revision.
}

// Iok8sapinetworkingv1ServiceCIDRSpec represents the Iok8sapinetworkingv1ServiceCIDRSpec schema from the OpenAPI specification
type Iok8sapinetworkingv1ServiceCIDRSpec struct {
	Cidrs []string `json:"cidrs,omitempty"` // CIDRs defines the IP blocks in CIDR notation (e.g. "192.168.0.0/24" or "2001:db8::/64") from which to assign service cluster IPs. Max of two CIDRs is allowed, one of each IP family. This field is immutable.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Strategy string `json:"strategy"` // strategy specifies how custom resources are converted between versions. Allowed values are: - `"None"`: The converter only change the apiVersion and would not touch any other field in the custom resource. - `"Webhook"`: API Server will call to an external webhook to do the conversion. Additional information is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set.
	Webhook GeneratedType `json:"webhook,omitempty"` // WebhookConversion describes how to call a conversion webhook
}

// Iok8sapinetworkingv1IngressRule represents the Iok8sapinetworkingv1IngressRule schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressRule struct {
	Host string `json:"host,omitempty"` // host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the "host" part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the IP in the Spec of the parent Ingress. 2. The `:` delimiter is not respected because ports are not allowed. 	 Currently the port of an Ingress is implicitly :80 for http and 	 :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue. host can be "precise" which is a domain name without the terminating dot of a network host (e.g. "foo.bar.com") or "wildcard", which is a domain name prefixed with a single wildcard label (e.g. "*.foo.com"). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == "*"). Requests will be matched against the Host field in the following way: 1. If host is precise, the request matches this rule if the http host header is equal to Host. 2. If host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.
	Http Iok8sapinetworkingv1HTTPIngressRuleValue `json:"http,omitempty"` // HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.
}

// Iok8sapicorev1Container represents the Iok8sapicorev1Container schema from the OpenAPI specification
type Iok8sapicorev1Container struct {
	Restartpolicy string `json:"restartPolicy,omitempty"` // RestartPolicy defines the restart behavior of individual containers in a pod. This field may only be set for init containers, and the only allowed value is "Always". For non-init containers or when this field is not specified, the restart behavior is defined by the Pod's restart policy and the container type. Setting the RestartPolicy as "Always" for the init container will have the following effect: this init container will be continually restarted on exit until all regular containers have terminated. Once all regular containers have completed, all init containers with restartPolicy "Always" will be shut down. This lifecycle differs from normal init containers and is often referred to as a "sidecar" container. Although this init container still starts in the init container sequence, it does not wait for the container to complete before proceeding to the next init container. Instead, the next init container starts immediately after this init container is started, or after any startupProbe has successfully completed.
	Resources Iok8sapicorev1ResourceRequirements `json:"resources,omitempty"` // ResourceRequirements describes the compute resource requirements.
	Ports []Iok8sapicorev1ContainerPort `json:"ports,omitempty"` // List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.
	Workingdir string `json:"workingDir,omitempty"` // Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	Startupprobe Iok8sapicorev1Probe `json:"startupProbe,omitempty"` // Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
	Terminationmessagepath string `json:"terminationMessagePath,omitempty"` // Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
	Args []string `json:"args,omitempty"` // Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Envfrom []Iok8sapicorev1EnvFromSource `json:"envFrom,omitempty"` // List of sources to populate environment variables in the container. The keys defined within a source may consist of any printable ASCII characters except '='. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
	Lifecycle Iok8sapicorev1Lifecycle `json:"lifecycle,omitempty"` // Lifecycle describes actions that the management system should take in response to container lifecycle events. For the PostStart and PreStop lifecycle handlers, management of the container blocks until the action is complete, unless the container process fails, in which case the handler is aborted.
	Livenessprobe Iok8sapicorev1Probe `json:"livenessProbe,omitempty"` // Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
	Readinessprobe Iok8sapicorev1Probe `json:"readinessProbe,omitempty"` // Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
	Stdin bool `json:"stdin,omitempty"` // Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
	Terminationmessagepolicy string `json:"terminationMessagePolicy,omitempty"` // Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
	Securitycontext Iok8sapicorev1SecurityContext `json:"securityContext,omitempty"` // SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext. When both are set, the values in SecurityContext take precedence.
	Volumedevices []Iok8sapicorev1VolumeDevice `json:"volumeDevices,omitempty"` // volumeDevices is the list of block devices to be used by the container.
	Tty bool `json:"tty,omitempty"` // Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
	Volumemounts []Iok8sapicorev1VolumeMount `json:"volumeMounts,omitempty"` // Pod volumes to mount into the container's filesystem. Cannot be updated.
	Env []Iok8sapicorev1EnvVar `json:"env,omitempty"` // List of environment variables to set in the container. Cannot be updated.
	Imagepullpolicy string `json:"imagePullPolicy,omitempty"` // Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	Resizepolicy []Iok8sapicorev1ContainerResizePolicy `json:"resizePolicy,omitempty"` // Resources resize policy for the container.
	Command []string `json:"command,omitempty"` // Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Image string `json:"image,omitempty"` // Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.
	Name string `json:"name"` // Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
	Stdinonce bool `json:"stdinOnce,omitempty"` // Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
}

// Iok8sapiflowcontrolv1PriorityLevelConfigurationReference represents the Iok8sapiflowcontrolv1PriorityLevelConfigurationReference schema from the OpenAPI specification
type Iok8sapiflowcontrolv1PriorityLevelConfigurationReference struct {
	Name string `json:"name"` // `name` is the name of the priority level configuration being referenced Required.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Clientconfig GeneratedType `json:"clientConfig,omitempty"` // WebhookClientConfig contains the information to make a TLS connection with the webhook.
	Conversionreviewversions []string `json:"conversionReviewVersions"` // conversionReviewVersions is an ordered list of preferred `ConversionReview` versions the Webhook expects. The API server will use the first version in the list which it supports. If none of the versions specified in this list are supported by API server, conversion will fail for the custom resource. If a persisted Webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail.
}

// Iok8sapischedulingv1PriorityClassList represents the Iok8sapischedulingv1PriorityClassList schema from the OpenAPI specification
type Iok8sapischedulingv1PriorityClassList struct {
	Items []Iok8sapischedulingv1PriorityClass `json:"items"` // items is the list of PriorityClasses
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapiresourcev1beta2Counter represents the Iok8sapiresourcev1beta2Counter schema from the OpenAPI specification
type Iok8sapiresourcev1beta2Counter struct {
	Value string `json:"value"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
}

// Iok8sapiresourcev1beta2ResourceClaimConsumerReference represents the Iok8sapiresourcev1beta2ResourceClaimConsumerReference schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ResourceClaimConsumerReference struct {
	Resource string `json:"resource"` // Resource is the type of resource being referenced, for example "pods".
	Uid string `json:"uid"` // UID identifies exactly one incarnation of the resource.
	Apigroup string `json:"apiGroup,omitempty"` // APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
	Name string `json:"name"` // Name is the name of resource being referenced.
}

// Iok8sapiresourcev1beta2AllocationResult represents the Iok8sapiresourcev1beta2AllocationResult schema from the OpenAPI specification
type Iok8sapiresourcev1beta2AllocationResult struct {
	Devices Iok8sapiresourcev1beta2DeviceAllocationResult `json:"devices,omitempty"` // DeviceAllocationResult is the result of allocating devices.
	Nodeselector Iok8sapicorev1NodeSelector `json:"nodeSelector,omitempty"` // A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
}

// Iok8sapicorev1NamespaceCondition represents the Iok8sapicorev1NamespaceCondition schema from the OpenAPI specification
type Iok8sapicorev1NamespaceCondition struct {
	Reason string `json:"reason,omitempty"` // Unique, one-word, CamelCase reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of namespace controller condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // Human-readable message indicating details about last transition.
}

// Iok8sapinetworkingv1beta1ServiceCIDRSpec represents the Iok8sapinetworkingv1beta1ServiceCIDRSpec schema from the OpenAPI specification
type Iok8sapinetworkingv1beta1ServiceCIDRSpec struct {
	Cidrs []string `json:"cidrs,omitempty"` // CIDRs defines the IP blocks in CIDR notation (e.g. "192.168.0.0/24" or "2001:db8::/64") from which to assign service cluster IPs. Max of two CIDRs is allowed, one of each IP family. This field is immutable.
}

// Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicy represents the Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicy schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicy struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicySpec `json:"spec,omitempty"` // MutatingAdmissionPolicySpec is the specification of the desired behavior of the admission policy.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapiflowcontrolv1PriorityLevelConfigurationList represents the Iok8sapiflowcontrolv1PriorityLevelConfigurationList schema from the OpenAPI specification
type Iok8sapiflowcontrolv1PriorityLevelConfigurationList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiflowcontrolv1PriorityLevelConfiguration `json:"items"` // `items` is a list of request-priorities.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapirbacv1ClusterRoleList represents the Iok8sapirbacv1ClusterRoleList schema from the OpenAPI specification
type Iok8sapirbacv1ClusterRoleList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapirbacv1ClusterRole `json:"items"` // Items is a list of ClusterRoles
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiflowcontrolv1PriorityLevelConfigurationSpec represents the Iok8sapiflowcontrolv1PriorityLevelConfigurationSpec schema from the OpenAPI specification
type Iok8sapiflowcontrolv1PriorityLevelConfigurationSpec struct {
	TypeField string `json:"type"` // `type` indicates whether this priority level is subject to limitation on request execution. A value of `"Exempt"` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels. A value of `"Limited"` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.
	Exempt Iok8sapiflowcontrolv1ExemptPriorityLevelConfiguration `json:"exempt,omitempty"` // ExemptPriorityLevelConfiguration describes the configurable aspects of the handling of exempt requests. In the mandatory exempt configuration object the values in the fields here can be modified by authorized users, unlike the rest of the `spec`.
	Limited Iok8sapiflowcontrolv1LimitedPriorityLevelConfiguration `json:"limited,omitempty"` // LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues: - How are requests for this priority level limited? - What should be done with requests that exceed the limit?
}

// Iok8sapicorev1GlusterfsPersistentVolumeSource represents the Iok8sapicorev1GlusterfsPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1GlusterfsPersistentVolumeSource struct {
	Readonly bool `json:"readOnly,omitempty"` // readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Endpoints string `json:"endpoints"` // endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Endpointsnamespace string `json:"endpointsNamespace,omitempty"` // endpointsNamespace is the namespace that contains Glusterfs endpoint. If this field is empty, the EndpointNamespace defaults to the same namespace as the bound PVC. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Path string `json:"path"` // path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
}

// Iok8sapiresourcev1alpha3DeviceTaintSelector represents the Iok8sapiresourcev1alpha3DeviceTaintSelector schema from the OpenAPI specification
type Iok8sapiresourcev1alpha3DeviceTaintSelector struct {
	Pool string `json:"pool,omitempty"` // If pool is set, only devices in that pool are selected. Also setting the driver name may be useful to avoid ambiguity when different drivers use the same pool name, but this is not required because selecting pools from different drivers may also be useful, for example when drivers with node-local devices use the node name as their pool name.
	Selectors []Iok8sapiresourcev1alpha3DeviceSelector `json:"selectors,omitempty"` // Selectors contains the same selection criteria as a ResourceClaim. Currently, CEL expressions are supported. All of these selectors must be satisfied.
	Device string `json:"device,omitempty"` // If device is set, only devices with that name are selected. This field corresponds to slice.spec.devices[].name. Setting also driver and pool may be required to avoid ambiguity, but is not required.
	Deviceclassname string `json:"deviceClassName,omitempty"` // If DeviceClassName is set, the selectors defined there must be satisfied by a device to be selected. This field corresponds to class.metadata.name.
	Driver string `json:"driver,omitempty"` // If driver is set, only devices from that driver are selected. This fields corresponds to slice.spec.driver.
}

// Iok8sapiresourcev1beta2DeviceClaim represents the Iok8sapiresourcev1beta2DeviceClaim schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceClaim struct {
	Requests []Iok8sapiresourcev1beta2DeviceRequest `json:"requests,omitempty"` // Requests represent individual requests for distinct devices which must all be satisfied. If empty, nothing needs to be allocated.
	Config []Iok8sapiresourcev1beta2DeviceClaimConfiguration `json:"config,omitempty"` // This field holds configuration for multiple potential drivers which could satisfy requests in this claim. It is ignored while allocating the claim.
	Constraints []Iok8sapiresourcev1beta2DeviceConstraint `json:"constraints,omitempty"` // These constraints must be satisfied by the set of devices that get allocated for the claim.
}

// Iok8sapirbacv1ClusterRoleBindingList represents the Iok8sapirbacv1ClusterRoleBindingList schema from the OpenAPI specification
type Iok8sapirbacv1ClusterRoleBindingList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapirbacv1ClusterRoleBinding `json:"items"` // Items is a list of ClusterRoleBindings
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiresourcev1beta2DeviceSelector represents the Iok8sapiresourcev1beta2DeviceSelector schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceSelector struct {
	Cel Iok8sapiresourcev1beta2CELDeviceSelector `json:"cel,omitempty"` // CELDeviceSelector contains a CEL expression for selecting a device.
}

// Iok8sapiauthenticationv1TokenRequest represents the Iok8sapiauthenticationv1TokenRequest schema from the OpenAPI specification
type Iok8sapiauthenticationv1TokenRequest struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthenticationv1TokenRequestSpec `json:"spec"` // TokenRequestSpec contains client provided parameters of a token request.
	Status Iok8sapiauthenticationv1TokenRequestStatus `json:"status,omitempty"` // TokenRequestStatus is the result of a token request.
}

// Iok8sapicorev1Lifecycle represents the Iok8sapicorev1Lifecycle schema from the OpenAPI specification
type Iok8sapicorev1Lifecycle struct {
	Prestop Iok8sapicorev1LifecycleHandler `json:"preStop,omitempty"` // LifecycleHandler defines a specific action that should be taken in a lifecycle hook. One and only one of the fields, except TCPSocket must be specified.
	Stopsignal string `json:"stopSignal,omitempty"` // StopSignal defines which signal will be sent to a container when it is being stopped. If not specified, the default is defined by the container runtime in use. StopSignal can only be set for Pods with a non-empty .spec.os.name
	Poststart Iok8sapicorev1LifecycleHandler `json:"postStart,omitempty"` // LifecycleHandler defines a specific action that should be taken in a lifecycle hook. One and only one of the fields, except TCPSocket must be specified.
}

// Iok8sapiresourcev1beta2DeviceCounterConsumption represents the Iok8sapiresourcev1beta2DeviceCounterConsumption schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceCounterConsumption struct {
	Counterset string `json:"counterSet"` // CounterSet is the name of the set from which the counters defined will be consumed.
	Counters map[string]interface{} `json:"counters"` // Counters defines the counters that will be consumed by the device. The maximum number counters in a device is 32. In addition, the maximum number of all counters in all devices is 1024 (for example, 64 devices with 16 counters each).
}

// Iok8sapicorev1SecretList represents the Iok8sapicorev1SecretList schema from the OpenAPI specification
type Iok8sapicorev1SecretList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1Secret `json:"items"` // Items is a list of secret objects. More info: https://kubernetes.io/docs/concepts/configuration/secret
}

// Iok8sapimachinerypkgapismetav1WatchEvent represents the Iok8sapimachinerypkgapismetav1WatchEvent schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1WatchEvent struct {
	Object Iok8sapimachinerypkgruntimeRawExtension `json:"object"` // RawExtension is used to hold extensions in external versions. To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types. // Internal package: 	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.Object `json:"myPlugin"` 	} 	type PluginA struct { 		AOption string `json:"aOption"` 	} // External package: 	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.RawExtension `json:"myPlugin"` 	} 	type PluginA struct { 		AOption string `json:"aOption"` 	} // On the wire, the JSON will look something like this: 	{ 		"kind":"MyAPIObject", 		"apiVersion":"v1", 		"myPlugin": { 			"kind":"PluginA", 			"aOption":"foo", 		}, 	} So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)
	TypeField string `json:"type"`
}

// Iok8sapicorev1NodeDaemonEndpoints represents the Iok8sapicorev1NodeDaemonEndpoints schema from the OpenAPI specification
type Iok8sapicorev1NodeDaemonEndpoints struct {
	Kubeletendpoint Iok8sapicorev1DaemonEndpoint `json:"kubeletEndpoint,omitempty"` // DaemonEndpoint contains information about a single Daemon endpoint.
}

// Iok8sapiresourcev1beta1ResourceSlice represents the Iok8sapiresourcev1beta1ResourceSlice schema from the OpenAPI specification
type Iok8sapiresourcev1beta1ResourceSlice struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiresourcev1beta1ResourceSliceSpec `json:"spec"` // ResourceSliceSpec contains the information published by the driver in one ResourceSlice.
}

// Iok8sapistoragev1VolumeAttachmentStatus represents the Iok8sapistoragev1VolumeAttachmentStatus schema from the OpenAPI specification
type Iok8sapistoragev1VolumeAttachmentStatus struct {
	Attacherror Iok8sapistoragev1VolumeError `json:"attachError,omitempty"` // VolumeError captures an error encountered during a volume operation.
	Attached bool `json:"attached"` // attached indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	Attachmentmetadata map[string]interface{} `json:"attachmentMetadata,omitempty"` // attachmentMetadata is populated with any information returned by the attach operation, upon successful attach, that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	Detacherror Iok8sapistoragev1VolumeError `json:"detachError,omitempty"` // VolumeError captures an error encountered during a volume operation.
}

// Iok8sapiapiserverinternalv1alpha1StorageVersionStatus represents the Iok8sapiapiserverinternalv1alpha1StorageVersionStatus schema from the OpenAPI specification
type Iok8sapiapiserverinternalv1alpha1StorageVersionStatus struct {
	Commonencodingversion string `json:"commonEncodingVersion,omitempty"` // If all API server instances agree on the same encoding storage version, then this field is set to that version. Otherwise this field is left empty. API servers should finish updating its storageVersionStatus entry before serving write operations, so that this field will be in sync with the reality.
	Conditions []Iok8sapiapiserverinternalv1alpha1StorageVersionCondition `json:"conditions,omitempty"` // The latest available observations of the storageVersion's state.
	Storageversions []Iok8sapiapiserverinternalv1alpha1ServerStorageVersion `json:"storageVersions,omitempty"` // The reported versions per API server instance.
}

// Iok8sapirbacv1RoleRef represents the Iok8sapirbacv1RoleRef schema from the OpenAPI specification
type Iok8sapirbacv1RoleRef struct {
	Kind string `json:"kind"` // Kind is the type of resource being referenced
	Name string `json:"name"` // Name is the name of resource being referenced
	Apigroup string `json:"apiGroup"` // APIGroup is the group for the resource being referenced
}

// Iok8sapimachinerypkgapismetav1StatusCause represents the Iok8sapimachinerypkgapismetav1StatusCause schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1StatusCause struct {
	Reason string `json:"reason,omitempty"` // A machine-readable description of the cause of the error. If this value is empty there is no information available.
	Field string `json:"field,omitempty"` // The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed. Fields may appear more than once in an array of causes due to fields having multiple errors. Optional. Examples: "name" - the field "name" on the current resource "items[0].name" - the field "name" on the first array entry in "items"
	Message string `json:"message,omitempty"` // A human-readable description of the cause of the error. This field may be presented as-is to a reader.
}

// Iok8sapiauthorizationv1SubjectAccessReviewStatus represents the Iok8sapiauthorizationv1SubjectAccessReviewStatus schema from the OpenAPI specification
type Iok8sapiauthorizationv1SubjectAccessReviewStatus struct {
	Allowed bool `json:"allowed"` // Allowed is required. True if the action would be allowed, false otherwise.
	Denied bool `json:"denied,omitempty"` // Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.
	Evaluationerror string `json:"evaluationError,omitempty"` // EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.
	Reason string `json:"reason,omitempty"` // Reason is optional. It indicates why a request was allowed or denied.
}

// Iok8sapiflowcontrolv1FlowSchemaSpec represents the Iok8sapiflowcontrolv1FlowSchemaSpec schema from the OpenAPI specification
type Iok8sapiflowcontrolv1FlowSchemaSpec struct {
	Rules []Iok8sapiflowcontrolv1PolicyRulesWithSubjects `json:"rules,omitempty"` // `rules` describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.
	Distinguishermethod Iok8sapiflowcontrolv1FlowDistinguisherMethod `json:"distinguisherMethod,omitempty"` // FlowDistinguisherMethod specifies the method of a flow distinguisher.
	Matchingprecedence int `json:"matchingPrecedence,omitempty"` // `matchingPrecedence` is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence. Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.
	Prioritylevelconfiguration Iok8sapiflowcontrolv1PriorityLevelConfigurationReference `json:"priorityLevelConfiguration"` // PriorityLevelConfigurationReference contains information that points to the "request-priority" being used.
}

// Iok8sapiresourcev1beta1DeviceClaim represents the Iok8sapiresourcev1beta1DeviceClaim schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceClaim struct {
	Constraints []Iok8sapiresourcev1beta1DeviceConstraint `json:"constraints,omitempty"` // These constraints must be satisfied by the set of devices that get allocated for the claim.
	Requests []Iok8sapiresourcev1beta1DeviceRequest `json:"requests,omitempty"` // Requests represent individual requests for distinct devices which must all be satisfied. If empty, nothing needs to be allocated.
	Config []Iok8sapiresourcev1beta1DeviceClaimConfiguration `json:"config,omitempty"` // This field holds configuration for multiple potential drivers which could satisfy requests in this claim. It is ignored while allocating the claim.
}

// Iok8sapicorev1CinderVolumeSource represents the Iok8sapicorev1CinderVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1CinderVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	Readonly bool `json:"readOnly,omitempty"` // readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	Volumeid string `json:"volumeID"` // volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
}

// Iok8sapicorev1VolumeResourceRequirements represents the Iok8sapicorev1VolumeResourceRequirements schema from the OpenAPI specification
type Iok8sapicorev1VolumeResourceRequirements struct {
	Limits map[string]interface{} `json:"limits,omitempty"` // Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests map[string]interface{} `json:"requests,omitempty"` // Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec GeneratedType `json:"spec"` // CustomResourceDefinitionSpec describes how a user wants their resource to appear
	Status GeneratedType `json:"status,omitempty"` // CustomResourceDefinitionStatus indicates the state of the CustomResourceDefinition
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapinodev1RuntimeClass represents the Iok8sapinodev1RuntimeClass schema from the OpenAPI specification
type Iok8sapinodev1RuntimeClass struct {
	Overhead Iok8sapinodev1Overhead `json:"overhead,omitempty"` // Overhead structure represents the resource overhead associated with running a pod.
	Scheduling Iok8sapinodev1Scheduling `json:"scheduling,omitempty"` // Scheduling specifies the scheduling constraints for nodes supporting a RuntimeClass.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Handler string `json:"handler"` // handler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration. It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The Handler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicoordinationv1beta1LeaseCandidate represents the Iok8sapicoordinationv1beta1LeaseCandidate schema from the OpenAPI specification
type Iok8sapicoordinationv1beta1LeaseCandidate struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicoordinationv1beta1LeaseCandidateSpec `json:"spec,omitempty"` // LeaseCandidateSpec is a specification of a Lease.
}

// Iok8sapicorev1FlexPersistentVolumeSource represents the Iok8sapicorev1FlexPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1FlexPersistentVolumeSource struct {
	Driver string `json:"driver"` // driver is the name of the driver to use for this volume.
	Fstype string `json:"fsType,omitempty"` // fsType is the Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	Options map[string]interface{} `json:"options,omitempty"` // options is Optional: this field holds extra command options if any.
	Readonly bool `json:"readOnly,omitempty"` // readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretref Iok8sapicorev1SecretReference `json:"secretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
}

// Iok8sapischedulingv1PriorityClass represents the Iok8sapischedulingv1PriorityClass schema from the OpenAPI specification
type Iok8sapischedulingv1PriorityClass struct {
	Description string `json:"description,omitempty"` // description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	Globaldefault bool `json:"globalDefault,omitempty"` // globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Preemptionpolicy string `json:"preemptionPolicy,omitempty"` // preemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.
	Value int `json:"value"` // value represents the integer value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapiadmissionregistrationv1ParamRef represents the Iok8sapiadmissionregistrationv1ParamRef schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ParamRef struct {
	Name string `json:"name,omitempty"` // name is the name of the resource being referenced. One of `name` or `selector` must be set, but `name` and `selector` are mutually exclusive properties. If one is set, the other must be unset. A single parameter used for all admission requests can be configured by setting the `name` field, leaving `selector` blank, and setting namespace if `paramKind` is namespace-scoped.
	Namespace string `json:"namespace,omitempty"` // namespace is the namespace of the referenced resource. Allows limiting the search for params to a specific namespace. Applies to both `name` and `selector` fields. A per-namespace parameter may be used by specifying a namespace-scoped `paramKind` in the policy and leaving this field empty. - If `paramKind` is cluster-scoped, this field MUST be unset. Setting this field results in a configuration error. - If `paramKind` is namespace-scoped, the namespace of the object being evaluated for admission will be used when this field is left unset. Take care that if this is left empty the binding must not match any cluster-scoped resources, which will result in an error.
	Parameternotfoundaction string `json:"parameterNotFoundAction,omitempty"` // `parameterNotFoundAction` controls the behavior of the binding when the resource exists, and name or selector is valid, but there are no parameters matched by the binding. If the value is set to `Allow`, then no matched parameters will be treated as successful validation by the binding. If set to `Deny`, then no matched parameters will be subject to the `failurePolicy` of the policy. Allowed values are `Allow` or `Deny` Required
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapiappsv1StatefulSetUpdateStrategy represents the Iok8sapiappsv1StatefulSetUpdateStrategy schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSetUpdateStrategy struct {
	Rollingupdate Iok8sapiappsv1RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty"` // RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.
	TypeField string `json:"type,omitempty"` // Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.
}

// Iok8sapicorev1LifecycleHandler represents the Iok8sapicorev1LifecycleHandler schema from the OpenAPI specification
type Iok8sapicorev1LifecycleHandler struct {
	Exec Iok8sapicorev1ExecAction `json:"exec,omitempty"` // ExecAction describes a "run in container" action.
	Httpget Iok8sapicorev1HTTPGetAction `json:"httpGet,omitempty"` // HTTPGetAction describes an action based on HTTP Get requests.
	Sleep Iok8sapicorev1SleepAction `json:"sleep,omitempty"` // SleepAction describes a "sleep" action.
	Tcpsocket Iok8sapicorev1TCPSocketAction `json:"tcpSocket,omitempty"` // TCPSocketAction describes an action based on opening a socket
}

// Iok8sapicorev1LimitRangeList represents the Iok8sapicorev1LimitRangeList schema from the OpenAPI specification
type Iok8sapicorev1LimitRangeList struct {
	Items []Iok8sapicorev1LimitRange `json:"items"` // Items is a list of LimitRange objects. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1NodeSelectorRequirement represents the Iok8sapicorev1NodeSelectorRequirement schema from the OpenAPI specification
type Iok8sapicorev1NodeSelectorRequirement struct {
	Key string `json:"key"` // The label key that the selector applies to.
	Operator string `json:"operator"` // Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
	Values []string `json:"values,omitempty"` // An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
}

// Iok8sapiresourcev1beta1AllocatedDeviceStatus represents the Iok8sapiresourcev1beta1AllocatedDeviceStatus schema from the OpenAPI specification
type Iok8sapiresourcev1beta1AllocatedDeviceStatus struct {
	Conditions []Iok8sapimachinerypkgapismetav1Condition `json:"conditions,omitempty"` // Conditions contains the latest observation of the device's state. If the device has been configured according to the class and claim config references, the `Ready` condition should be True. Must not contain more than 8 entries.
	Data Iok8sapimachinerypkgruntimeRawExtension `json:"data,omitempty"` // RawExtension is used to hold extensions in external versions. To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types. // Internal package: 	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.Object `json:"myPlugin"` 	} 	type PluginA struct { 		AOption string `json:"aOption"` 	} // External package: 	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.RawExtension `json:"myPlugin"` 	} 	type PluginA struct { 		AOption string `json:"aOption"` 	} // On the wire, the JSON will look something like this: 	{ 		"kind":"MyAPIObject", 		"apiVersion":"v1", 		"myPlugin": { 			"kind":"PluginA", 			"aOption":"foo", 		}, 	} So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)
	Device string `json:"device"` // Device references one device instance via its name in the driver's resource pool. It must be a DNS label.
	Driver string `json:"driver"` // Driver specifies the name of the DRA driver whose kubelet plugin should be invoked to process the allocation once the claim is needed on a node. Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver.
	Networkdata Iok8sapiresourcev1beta1NetworkDeviceData `json:"networkData,omitempty"` // NetworkDeviceData provides network-related details for the allocated device. This information may be filled by drivers or other components to configure or identify the device within a network context.
	Pool string `json:"pool"` // This name together with the driver name and the device name field identify which device was allocated (`<driver name>/<pool name>/<device name>`). Must not be longer than 253 characters and may contain one or more DNS sub-domains separated by slashes.
}

// Iok8sapiadmissionregistrationv1ServiceReference represents the Iok8sapiadmissionregistrationv1ServiceReference schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ServiceReference struct {
	Port int `json:"port,omitempty"` // If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
	Name string `json:"name"` // `name` is the name of the service. Required
	Namespace string `json:"namespace"` // `namespace` is the namespace of the service. Required
	Path string `json:"path,omitempty"` // `path` is an optional URL path which will be sent in any request to this service.
}

// Iok8sapimachinerypkgapismetav1Preconditions represents the Iok8sapimachinerypkgapismetav1Preconditions schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1Preconditions struct {
	Resourceversion string `json:"resourceVersion,omitempty"` // Specifies the target ResourceVersion
	Uid string `json:"uid,omitempty"` // Specifies the target UID.
}

// Iok8sapiresourcev1beta1Device represents the Iok8sapiresourcev1beta1Device schema from the OpenAPI specification
type Iok8sapiresourcev1beta1Device struct {
	Name string `json:"name"` // Name is unique identifier among all devices managed by the driver in the pool. It must be a DNS label.
	Basic Iok8sapiresourcev1beta1BasicDevice `json:"basic,omitempty"` // BasicDevice defines one device instance.
}

// Iok8sapiresourcev1beta1DeviceTaint represents the Iok8sapiresourcev1beta1DeviceTaint schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceTaint struct {
	Key string `json:"key"` // The taint key to be applied to a device. Must be a label name.
	Timeadded string `json:"timeAdded,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Value string `json:"value,omitempty"` // The taint value corresponding to the taint key. Must be a label value.
	Effect string `json:"effect"` // The effect of the taint on claims that do not tolerate the taint and through such claims on the pods using them. Valid effects are NoSchedule and NoExecute. PreferNoSchedule as used for nodes is not valid here.
}

// Iok8sapicorev1SessionAffinityConfig represents the Iok8sapicorev1SessionAffinityConfig schema from the OpenAPI specification
type Iok8sapicorev1SessionAffinityConfig struct {
	Clientip Iok8sapicorev1ClientIPConfig `json:"clientIP,omitempty"` // ClientIPConfig represents the configurations of Client IP based session affinity.
}

// Iok8sapiresourcev1beta1ResourceSliceSpec represents the Iok8sapiresourcev1beta1ResourceSliceSpec schema from the OpenAPI specification
type Iok8sapiresourcev1beta1ResourceSliceSpec struct {
	Driver string `json:"driver"` // Driver identifies the DRA driver providing the capacity information. A field selector can be used to list only ResourceSlice objects with a certain driver name. Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver. This field is immutable.
	Nodename string `json:"nodeName,omitempty"` // NodeName identifies the node which provides the resources in this pool. A field selector can be used to list only ResourceSlice objects belonging to a certain node. This field can be used to limit access from nodes to ResourceSlices with the same node name. It also indicates to autoscalers that adding new nodes of the same type as some old node might also make new resources available. Exactly one of NodeName, NodeSelector, AllNodes, and PerDeviceNodeSelection must be set. This field is immutable.
	Nodeselector Iok8sapicorev1NodeSelector `json:"nodeSelector,omitempty"` // A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
	Perdevicenodeselection bool `json:"perDeviceNodeSelection,omitempty"` // PerDeviceNodeSelection defines whether the access from nodes to resources in the pool is set on the ResourceSlice level or on each device. If it is set to true, every device defined the ResourceSlice must specify this individually. Exactly one of NodeName, NodeSelector, AllNodes, and PerDeviceNodeSelection must be set.
	Pool Iok8sapiresourcev1beta1ResourcePool `json:"pool"` // ResourcePool describes the pool that ResourceSlices belong to.
	Sharedcounters []Iok8sapiresourcev1beta1CounterSet `json:"sharedCounters,omitempty"` // SharedCounters defines a list of counter sets, each of which has a name and a list of counters available. The names of the SharedCounters must be unique in the ResourceSlice. The maximum number of SharedCounters is 32.
	Allnodes bool `json:"allNodes,omitempty"` // AllNodes indicates that all nodes have access to the resources in the pool. Exactly one of NodeName, NodeSelector, AllNodes, and PerDeviceNodeSelection must be set.
	Devices []Iok8sapiresourcev1beta1Device `json:"devices,omitempty"` // Devices lists some or all of the devices in this pool. Must not have more than 128 entries.
}

// Iok8sapimachinerypkgapismetav1APIResource represents the Iok8sapimachinerypkgapismetav1APIResource schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1APIResource struct {
	Version string `json:"version,omitempty"` // version is the preferred version of the resource. Empty implies the version of the containing resource list For subresources, this may have a different value, for example: v1 (while inside a v1beta1 version of the core resource's group)".
	Verbs []string `json:"verbs"` // verbs is a list of supported kube verbs (this includes get, list, watch, create, update, patch, delete, deletecollection, and proxy)
	Storageversionhash string `json:"storageVersionHash,omitempty"` // The hash value of the storage version, the version this resource is converted to when written to the data store. Value must be treated as opaque by clients. Only equality comparison on the value is valid. This is an alpha feature and may change or be removed in the future. The field is populated by the apiserver only if the StorageVersionHash feature gate is enabled. This field will remain optional even if it graduates.
	Categories []string `json:"categories,omitempty"` // categories is a list of the grouped resources this resource belongs to (e.g. 'all')
	Group string `json:"group,omitempty"` // group is the preferred group of the resource. Empty implies the group of the containing resource list. For subresources, this may have a different value, for example: Scale".
	Kind string `json:"kind"` // kind is the kind for the resource (e.g. 'Foo' is the kind for a resource 'foo')
	Shortnames []string `json:"shortNames,omitempty"` // shortNames is a list of suggested short names of the resource.
	Name string `json:"name"` // name is the plural name of the resource.
	Namespaced bool `json:"namespaced"` // namespaced indicates if a resource is namespaced or not.
	Singularname string `json:"singularName"` // singularName is the singular name of the resource. This allows clients to handle plural and singular opaquely. The singularName is more correct for reporting status on a single item and both singular and plural are allowed from the kubectl CLI interface.
}

// Iok8sapiadmissionregistrationv1WebhookClientConfig represents the Iok8sapiadmissionregistrationv1WebhookClientConfig schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1WebhookClientConfig struct {
	Url string `json:"url,omitempty"` // `url` gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified. The `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address. Please note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster. The scheme must be "https"; the URL must begin with "https://". A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier. Attempting to use a user or basic auth e.g. "user:password@" is not allowed. Fragments ("#...") and query parameters ("?...") are not allowed, either.
	Cabundle string `json:"caBundle,omitempty"` // `caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.
	Service Iok8sapiadmissionregistrationv1ServiceReference `json:"service,omitempty"` // ServiceReference holds a reference to Service.legacy.k8s.io
}

// Iok8sapiresourcev1beta1DeviceAllocationResult represents the Iok8sapiresourcev1beta1DeviceAllocationResult schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceAllocationResult struct {
	Config []Iok8sapiresourcev1beta1DeviceAllocationConfiguration `json:"config,omitempty"` // This field is a combination of all the claim and class configuration parameters. Drivers can distinguish between those based on a flag. This includes configuration parameters for drivers which have no allocated devices in the result because it is up to the drivers which configuration parameters they support. They can silently ignore unknown configuration parameters.
	Results []Iok8sapiresourcev1beta1DeviceRequestAllocationResult `json:"results,omitempty"` // Results lists all allocated devices.
}

// Iok8sapimachinerypkgapismetav1ServerAddressByClientCIDR represents the Iok8sapimachinerypkgapismetav1ServerAddressByClientCIDR schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1ServerAddressByClientCIDR struct {
	Clientcidr string `json:"clientCIDR"` // The CIDR with which clients can match their IP to figure out the server address that they should use.
	Serveraddress string `json:"serverAddress"` // Address of this server, suitable for a client that matches the above CIDR. This can be a hostname, hostname:port, IP or IP:port.
}

// Iok8sapimachinerypkgapismetav1APIVersions represents the Iok8sapimachinerypkgapismetav1APIVersions schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1APIVersions struct {
	Versions []string `json:"versions"` // versions are the api versions that are available.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Serveraddressbyclientcidrs []Iok8sapimachinerypkgapismetav1ServerAddressByClientCIDR `json:"serverAddressByClientCIDRs"` // a map of client CIDR to server address that is serving this group. This is to help clients reach servers in the most network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case of multiple matches, clients should use the longest matching CIDR. The server returns only those CIDRs that it thinks that the client can match. For example: the master will return an internal IP CIDR only, if the client reaches the server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that order) to get the client IP.
}

// Iok8sapinetworkingv1NetworkPolicy represents the Iok8sapinetworkingv1NetworkPolicy schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicy struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapinetworkingv1NetworkPolicySpec `json:"spec,omitempty"` // NetworkPolicySpec provides the specification of a NetworkPolicy
}

// Iok8sapicorev1ComponentStatus represents the Iok8sapicorev1ComponentStatus schema from the OpenAPI specification
type Iok8sapicorev1ComponentStatus struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Conditions []Iok8sapicorev1ComponentCondition `json:"conditions,omitempty"` // List of component conditions observed
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiautoscalingv2PodsMetricSource represents the Iok8sapiautoscalingv2PodsMetricSource schema from the OpenAPI specification
type Iok8sapiautoscalingv2PodsMetricSource struct {
	Metric Iok8sapiautoscalingv2MetricIdentifier `json:"metric"` // MetricIdentifier defines the name and optionally selector for a metric
	Target Iok8sapiautoscalingv2MetricTarget `json:"target"` // MetricTarget defines the target value, average value, or average utilization of a specific metric
}

// Iok8sapiflowcontrolv1FlowSchemaCondition represents the Iok8sapiflowcontrolv1FlowSchemaCondition schema from the OpenAPI specification
type Iok8sapiflowcontrolv1FlowSchemaCondition struct {
	TypeField string `json:"type,omitempty"` // `type` is the type of the condition. Required.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // `message` is a human-readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // `reason` is a unique, one-word, CamelCase reason for the condition's last transition.
	Status string `json:"status,omitempty"` // `status` is the status of the condition. Can be True, False, Unknown. Required.
}

// Iok8sapiauthorizationv1FieldSelectorAttributes represents the Iok8sapiauthorizationv1FieldSelectorAttributes schema from the OpenAPI specification
type Iok8sapiauthorizationv1FieldSelectorAttributes struct {
	Rawselector string `json:"rawSelector,omitempty"` // rawSelector is the serialization of a field selector that would be included in a query parameter. Webhook implementations are encouraged to ignore rawSelector. The kube-apiserver's *SubjectAccessReview will parse the rawSelector as long as the requirements are not present.
	Requirements []Iok8sapimachinerypkgapismetav1FieldSelectorRequirement `json:"requirements,omitempty"` // requirements is the parsed interpretation of a field selector. All requirements must be met for a resource instance to match the selector. Webhook implementations should handle requirements, but how to handle them is up to the webhook. Since requirements can only limit the request, it is safe to authorize as unlimited request if the requirements are not understood.
}

// Iok8sapiadmissionregistrationv1alpha1JSONPatch represents the Iok8sapiadmissionregistrationv1alpha1JSONPatch schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1JSONPatch struct {
	Expression string `json:"expression,omitempty"` // expression will be evaluated by CEL to create a [JSON patch](https://jsonpatch.com/). ref: https://github.com/google/cel-spec expression must return an array of JSONPatch values. For example, this CEL expression returns a JSON patch to conditionally modify a value: 	 [ 	 JSONPatch{op: "test", path: "/spec/example", value: "Red"}, 	 JSONPatch{op: "replace", path: "/spec/example", value: "Green"} 	 ] To define an object for the patch value, use Object types. For example: 	 [ 	 JSONPatch{ 	 op: "add", 	 path: "/spec/selector", 	 value: Object.spec.selector{matchLabels: {"environment": "test"}} 	 } 	 ] To use strings containing '/' and '~' as JSONPatch path keys, use "jsonpatch.escapeKey". For example: 	 [ 	 JSONPatch{ 	 op: "add", 	 path: "/metadata/labels/" + jsonpatch.escapeKey("example.com/environment"), 	 value: "test" 	 }, 	 ] CEL expressions have access to the types needed to create JSON patches and objects: - 'JSONPatch' - CEL type of JSON Patch operations. JSONPatch has the fields 'op', 'from', 'path' and 'value'. See [JSON patch](https://jsonpatch.com/) for more details. The 'value' field may be set to any of: string, integer, array, map or object. If set, the 'path' and 'from' fields must be set to a [JSON pointer](https://datatracker.ietf.org/doc/html/rfc6901/) string, where the 'jsonpatch.escapeKey()' CEL function may be used to escape path keys containing '/' and '~'. - 'Object' - CEL type of the resource object. - 'Object.<fieldName>' - CEL type of object field (such as 'Object.spec') - 'Object.<fieldName1>.<fieldName2>...<fieldNameN>` - CEL type of nested field (such as 'Object.spec.containers') CEL expressions have access to the contents of the API request, organized into CEL variables as well as some other useful variables: - 'object' - The object from the incoming request. The value is null for DELETE requests. - 'oldObject' - The existing object. The value is null for CREATE requests. - 'request' - Attributes of the API request([ref](/pkg/apis/admission/types.go#AdmissionRequest)). - 'params' - Parameter resource referred to by the policy binding being evaluated. Only populated if the policy has a ParamKind. - 'namespaceObject' - The namespace object that the incoming object belongs to. The value is null for cluster-scoped resources. - 'variables' - Map of composited variables, from its name to its lazily evaluated value. For example, a variable named 'foo' can be accessed as 'variables.foo'. - 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request. See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz - 'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the request resource. CEL expressions have access to [Kubernetes CEL function libraries](https://kubernetes.io/docs/reference/using-api/cel/#cel-options-language-features-and-libraries) as well as: - 'jsonpatch.escapeKey' - Performs JSONPatch key escaping. '~' and '/' are escaped as '~0' and `~1' respectively). Only property names of the form `[a-zA-Z_.-/][a-zA-Z0-9_.-/]*` are accessible. Required.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Scale GeneratedType `json:"scale,omitempty"` // CustomResourceSubresourceScale defines how to serve the scale subresource for CustomResources.
	Status GeneratedType `json:"status,omitempty"` // CustomResourceSubresourceStatus defines how to serve the status subresource for CustomResources. Status is represented by the `.status` JSON path inside of a CustomResource. When set, * exposes a /status subresource for the custom resource * PUT requests to the /status subresource take a custom resource object, and ignore changes to anything except the status stanza * PUT/POST/PATCH requests to the custom resource ignore changes to the status stanza
}

// Iok8sapistoragev1VolumeError represents the Iok8sapistoragev1VolumeError schema from the OpenAPI specification
type Iok8sapistoragev1VolumeError struct {
	Errorcode int `json:"errorCode,omitempty"` // errorCode is a numeric gRPC code representing the error encountered during Attach or Detach operations. This is an optional, alpha field that requires the MutableCSINodeAllocatableCount feature gate being enabled to be set.
	Message string `json:"message,omitempty"` // message represents the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information.
	Time string `json:"time,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
}

// Iok8sapiadmissionregistrationv1Variable represents the Iok8sapiadmissionregistrationv1Variable schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1Variable struct {
	Name string `json:"name"` // Name is the name of the variable. The name must be a valid CEL identifier and unique among all variables. The variable can be accessed in other expressions through `variables` For example, if name is "foo", the variable will be available as `variables.foo`
	Expression string `json:"expression"` // Expression is the expression that will be evaluated as the value of the variable. The CEL expression has access to the same identifiers as the CEL expressions in Validation.
}

// Iok8sapicorev1PodSpec represents the Iok8sapicorev1PodSpec schema from the OpenAPI specification
type Iok8sapicorev1PodSpec struct {
	Nodename string `json:"nodeName,omitempty"` // NodeName indicates in which node this pod is scheduled. If empty, this pod is a candidate for scheduling by the scheduler defined in schedulerName. Once this field is set, the kubelet for this node becomes responsible for the lifecycle of this pod. This field should not be used to express a desire for the pod to be scheduled on a specific node. https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodename
	Priorityclassname string `json:"priorityClassName,omitempty"` // If specified, indicates the pod's priority. "system-node-critical" and "system-cluster-critical" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.
	Schedulinggates []Iok8sapicorev1PodSchedulingGate `json:"schedulingGates,omitempty"` // SchedulingGates is an opaque list of values that if specified will block scheduling the pod. If schedulingGates is not empty, the pod will stay in the SchedulingGated state and the scheduler will not attempt to schedule the pod. SchedulingGates can only be set at pod creation time, and be removed only afterwards.
	Serviceaccountname string `json:"serviceAccountName,omitempty"` // ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	Volumes []Iok8sapicorev1Volume `json:"volumes,omitempty"` // List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes
	Hostpid bool `json:"hostPID,omitempty"` // Use the host's pid namespace. Optional: Default to false.
	Tolerations []Iok8sapicorev1Toleration `json:"tolerations,omitempty"` // If specified, the pod's tolerations.
	Dnspolicy string `json:"dnsPolicy,omitempty"` // Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
	Enableservicelinks bool `json:"enableServiceLinks,omitempty"` // EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true.
	Hostaliases []Iok8sapicorev1HostAlias `json:"hostAliases,omitempty"` // HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified.
	Priority int `json:"priority,omitempty"` // The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.
	Automountserviceaccounttoken bool `json:"automountServiceAccountToken,omitempty"` // AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.
	Topologyspreadconstraints []Iok8sapicorev1TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"` // TopologySpreadConstraints describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.
	Ephemeralcontainers []Iok8sapicorev1EphemeralContainer `json:"ephemeralContainers,omitempty"` // List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing pod to perform user-initiated actions such as debugging. This list cannot be specified when creating a pod, and it cannot be modified by updating the pod spec. In order to add an ephemeral container to an existing pod, use the pod's ephemeralcontainers subresource.
	Securitycontext Iok8sapicorev1PodSecurityContext `json:"securityContext,omitempty"` // PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext. Field values of container.securityContext take precedence over field values of PodSecurityContext.
	Affinity Iok8sapicorev1Affinity `json:"affinity,omitempty"` // Affinity is a group of affinity scheduling rules.
	Imagepullsecrets []Iok8sapicorev1LocalObjectReference `json:"imagePullSecrets,omitempty"` // ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	Sethostnameasfqdn bool `json:"setHostnameAsFQDN,omitempty"` // If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default). In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname). In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\Tcpip\\Parameters to FQDN. If a pod does not have FQDN, this has no effect. Default to false.
	Dnsconfig Iok8sapicorev1PodDNSConfig `json:"dnsConfig,omitempty"` // PodDNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.
	Restartpolicy string `json:"restartPolicy,omitempty"` // Restart policy for all containers within the pod. One of Always, OnFailure, Never. In some contexts, only a subset of those values may be permitted. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	Hostusers bool `json:"hostUsers,omitempty"` // Use the host's user namespace. Optional: Default to true. If set to true or not present, the pod will be run in the host user namespace, useful for when the pod needs a feature only available to the host user namespace, such as loading a kernel module with CAP_SYS_MODULE. When set to false, a new userns is created for the pod. Setting false is useful for mitigating container breakout vulnerabilities even allowing users to run their containers as root without actually having root privileges on the host. This field is alpha-level and is only honored by servers that enable the UserNamespacesSupport feature.
	Resourceclaims []Iok8sapicorev1PodResourceClaim `json:"resourceClaims,omitempty"` // ResourceClaims defines which ResourceClaims must be allocated and reserved before the Pod is allowed to start. The resources will be made available to those containers which consume them by name. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable.
	Activedeadlineseconds int64 `json:"activeDeadlineSeconds,omitempty"` // Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.
	Subdomain string `json:"subdomain,omitempty"` // If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". If not specified, the pod will not have a domainname at all.
	Runtimeclassname string `json:"runtimeClassName,omitempty"` // RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod. If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the "legacy" RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class
	Shareprocessnamespace bool `json:"shareProcessNamespace,omitempty"` // Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false.
	Preemptionpolicy string `json:"preemptionPolicy,omitempty"` // PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.
	Readinessgates []Iok8sapicorev1PodReadinessGate `json:"readinessGates,omitempty"` // If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to "True" More info: https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates
	Overhead map[string]interface{} `json:"overhead,omitempty"` // Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests. The RuntimeClass admission controller will reject Pod create requests which have the overhead already set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero. More info: https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md
	Terminationgraceperiodseconds int64 `json:"terminationGracePeriodSeconds,omitempty"` // Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.
	Nodeselector map[string]interface{} `json:"nodeSelector,omitempty"` // NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	Os Iok8sapicorev1PodOS `json:"os,omitempty"` // PodOS defines the OS parameters of a pod.
	Schedulername string `json:"schedulerName,omitempty"` // If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.
	Hostnetwork bool `json:"hostNetwork,omitempty"` // Host networking requested for this pod. Use the host's network namespace. When using HostNetwork you should specify ports so the scheduler is aware. When `hostNetwork` is true, specified `hostPort` fields in port definitions must match `containerPort`, and unspecified `hostPort` fields in port definitions are defaulted to match `containerPort`. Default to false.
	Hostname string `json:"hostname,omitempty"` // Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.
	Containers []Iok8sapicorev1Container `json:"containers"` // List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated.
	Serviceaccount string `json:"serviceAccount,omitempty"` // DeprecatedServiceAccount is a deprecated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.
	Initcontainers []Iok8sapicorev1Container `json:"initContainers,omitempty"` // List of initialization containers belonging to the pod. Init containers are executed in order prior to containers being started. If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy. The name for an init container or normal container must be unique among all containers. Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes. The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of that value or the sum of the normal containers. Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	Resources Iok8sapicorev1ResourceRequirements `json:"resources,omitempty"` // ResourceRequirements describes the compute resource requirements.
	Hostipc bool `json:"hostIPC,omitempty"` // Use the host's ipc namespace. Optional: Default to false.
}

// Iok8sapiresourcev1beta2CounterSet represents the Iok8sapiresourcev1beta2CounterSet schema from the OpenAPI specification
type Iok8sapiresourcev1beta2CounterSet struct {
	Counters map[string]interface{} `json:"counters"` // Counters defines the set of counters for this CounterSet The name of each counter must be unique in that set and must be a DNS label. The maximum number of counters in all sets is 32.
	Name string `json:"name"` // Name defines the name of the counter set. It must be a DNS label.
}

// Iok8sapicertificatesv1beta1ClusterTrustBundleSpec represents the Iok8sapicertificatesv1beta1ClusterTrustBundleSpec schema from the OpenAPI specification
type Iok8sapicertificatesv1beta1ClusterTrustBundleSpec struct {
	Trustbundle string `json:"trustBundle"` // trustBundle contains the individual X.509 trust anchors for this bundle, as PEM bundle of PEM-wrapped, DER-formatted X.509 certificates. The data must consist only of PEM certificate blocks that parse as valid X.509 certificates. Each certificate must include a basic constraints extension with the CA bit set. The API server will reject objects that contain duplicate certificates, or that use PEM block headers. Users of ClusterTrustBundles, including Kubelet, are free to reorder and deduplicate certificate blocks in this file according to their own logic, as well as to drop PEM block headers and inter-block data.
	Signername string `json:"signerName,omitempty"` // signerName indicates the associated signer, if any. In order to create or update a ClusterTrustBundle that sets signerName, you must have the following cluster-scoped permission: group=certificates.k8s.io resource=signers resourceName=<the signer name> verb=attest. If signerName is not empty, then the ClusterTrustBundle object must be named with the signer name as a prefix (translating slashes to colons). For example, for the signer name `example.com/foo`, valid ClusterTrustBundle object names include `example.com:foo:abc` and `example.com:foo:v1`. If signerName is empty, then the ClusterTrustBundle object's name must not have such a prefix. List/watch requests for ClusterTrustBundles can filter on this field using a `spec.signerName=NAME` field selector.
}

// Iok8sapicorev1NodeSelectorTerm represents the Iok8sapicorev1NodeSelectorTerm schema from the OpenAPI specification
type Iok8sapicorev1NodeSelectorTerm struct {
	Matchexpressions []Iok8sapicorev1NodeSelectorRequirement `json:"matchExpressions,omitempty"` // A list of node selector requirements by node's labels.
	Matchfields []Iok8sapicorev1NodeSelectorRequirement `json:"matchFields,omitempty"` // A list of node selector requirements by node's fields.
}

// Iok8sapiflowcontrolv1FlowSchemaStatus represents the Iok8sapiflowcontrolv1FlowSchemaStatus schema from the OpenAPI specification
type Iok8sapiflowcontrolv1FlowSchemaStatus struct {
	Conditions []Iok8sapiflowcontrolv1FlowSchemaCondition `json:"conditions,omitempty"` // `conditions` is a list of the current states of FlowSchema.
}

// Iok8sapicorev1LimitRangeSpec represents the Iok8sapicorev1LimitRangeSpec schema from the OpenAPI specification
type Iok8sapicorev1LimitRangeSpec struct {
	Limits []Iok8sapicorev1LimitRangeItem `json:"limits"` // Limits is the list of LimitRangeItem objects that are enforced.
}

// Iok8sapibatchv1CronJobList represents the Iok8sapibatchv1CronJobList schema from the OpenAPI specification
type Iok8sapibatchv1CronJobList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapibatchv1CronJob `json:"items"` // items is the list of CronJobs.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapiautoscalingv1HorizontalPodAutoscalerStatus represents the Iok8sapiautoscalingv1HorizontalPodAutoscalerStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv1HorizontalPodAutoscalerStatus struct {
	Currentreplicas int `json:"currentReplicas"` // currentReplicas is the current number of replicas of pods managed by this autoscaler.
	Desiredreplicas int `json:"desiredReplicas"` // desiredReplicas is the desired number of replicas of pods managed by this autoscaler.
	Lastscaletime string `json:"lastScaleTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // observedGeneration is the most recent generation observed by this autoscaler.
	Currentcpuutilizationpercentage int `json:"currentCPUUtilizationPercentage,omitempty"` // currentCPUUtilizationPercentage is the current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
}

// Iok8sapiresourcev1beta1ResourceClaimList represents the Iok8sapiresourcev1beta1ResourceClaimList schema from the OpenAPI specification
type Iok8sapiresourcev1beta1ResourceClaimList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiresourcev1beta1ResourceClaim `json:"items"` // Items is the list of resource claims.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapistoragev1VolumeAttachmentList represents the Iok8sapistoragev1VolumeAttachmentList schema from the OpenAPI specification
type Iok8sapistoragev1VolumeAttachmentList struct {
	Items []Iok8sapistoragev1VolumeAttachment `json:"items"` // items is the list of VolumeAttachments
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1NamespaceList represents the Iok8sapicorev1NamespaceList schema from the OpenAPI specification
type Iok8sapicorev1NamespaceList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1Namespace `json:"items"` // Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
}

// Iok8sapicorev1TopologySpreadConstraint represents the Iok8sapicorev1TopologySpreadConstraint schema from the OpenAPI specification
type Iok8sapicorev1TopologySpreadConstraint struct {
	Labelselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"labelSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Matchlabelkeys []string `json:"matchLabelKeys,omitempty"` // MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector. This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).
	Maxskew int `json:"maxSkew"` // MaxSkew describes the degree to which pods may be unevenly distributed. When `whenUnsatisfiable=DoNotSchedule`, it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | | P P | P P | P | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When `whenUnsatisfiable=ScheduleAnyway`, it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.
	Mindomains int `json:"minDomains,omitempty"` // MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats "global minimum" as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule. For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | | P P | P P | P P | The number of domains is less than 5(MinDomains), so "global minimum" is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.
	Nodeaffinitypolicy string `json:"nodeAffinityPolicy,omitempty"` // NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations. If this value is nil, the behavior is equivalent to the Honor policy.
	Nodetaintspolicy string `json:"nodeTaintsPolicy,omitempty"` // NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included. If this value is nil, the behavior is equivalent to the Ignore policy.
	Topologykey string `json:"topologyKey"` // TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a "bucket", and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is "kubernetes.io/hostname", each Node is a domain of that topology. And, if TopologyKey is "topology.kubernetes.io/zone", each zone is a domain of that topology. It's a required field.
	Whenunsatisfiable string `json:"whenUnsatisfiable"` // WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered "Unsatisfiable" for an incoming pod if and only if every possible node assignment for that pod would violate "MaxSkew" on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P | P | P | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.
}

// Iok8sapieventsv1Event represents the Iok8sapieventsv1Event schema from the OpenAPI specification
type Iok8sapieventsv1Event struct {
	Action string `json:"action,omitempty"` // action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.
	Note string `json:"note,omitempty"` // note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Eventtime string `json:"eventTime"` // MicroTime is version of Time with microsecond level precision.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Reason string `json:"reason,omitempty"` // reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.
	Deprecatedfirsttimestamp string `json:"deprecatedFirstTimestamp,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Series Iok8sapieventsv1EventSeries `json:"series,omitempty"` // EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time. How often to update the EventSeries is up to the event reporters. The default event reporter in "k8s.io/client-go/tools/events/event_broadcaster.go" shows how this struct is updated on heartbeats and can guide customized reporter implementations.
	Deprecatedcount int `json:"deprecatedCount,omitempty"` // deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Deprecatedlasttimestamp string `json:"deprecatedLastTimestamp,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Related Iok8sapicorev1ObjectReference `json:"related,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Reportinginstance string `json:"reportingInstance,omitempty"` // reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
	Deprecatedsource Iok8sapicorev1EventSource `json:"deprecatedSource,omitempty"` // EventSource contains information for an event.
	Regarding Iok8sapicorev1ObjectReference `json:"regarding,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Reportingcontroller string `json:"reportingController,omitempty"` // reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
	TypeField string `json:"type,omitempty"` // type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.
}

// Iok8sapicorev1NodeSpec represents the Iok8sapicorev1NodeSpec schema from the OpenAPI specification
type Iok8sapicorev1NodeSpec struct {
	Configsource Iok8sapicorev1NodeConfigSource `json:"configSource,omitempty"` // NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil. This API is deprecated since 1.22
	Externalid string `json:"externalID,omitempty"` // Deprecated. Not all kubelets will set this field. Remove field after 1.13. see: https://issues.k8s.io/61966
	Podcidr string `json:"podCIDR,omitempty"` // PodCIDR represents the pod IP range assigned to the node.
	Podcidrs []string `json:"podCIDRs,omitempty"` // podCIDRs represents the IP ranges assigned to the node for usage by Pods on that node. If this field is specified, the 0th entry must match the podCIDR field. It may contain at most 1 value for each of IPv4 and IPv6.
	Providerid string `json:"providerID,omitempty"` // ID of the node assigned by the cloud provider in the format: <ProviderName>://<ProviderSpecificNodeID>
	Taints []Iok8sapicorev1Taint `json:"taints,omitempty"` // If specified, the node's taints.
	Unschedulable bool `json:"unschedulable,omitempty"` // Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration
}

// Iok8sapicorev1SecretKeySelector represents the Iok8sapicorev1SecretKeySelector schema from the OpenAPI specification
type Iok8sapicorev1SecretKeySelector struct {
	Key string `json:"key"` // The key of the secret to select from. Must be a valid secret key.
	Name string `json:"name,omitempty"` // Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the Secret or its key must be defined
}

// Iok8sapicorev1PreferredSchedulingTerm represents the Iok8sapicorev1PreferredSchedulingTerm schema from the OpenAPI specification
type Iok8sapicorev1PreferredSchedulingTerm struct {
	Preference Iok8sapicorev1NodeSelectorTerm `json:"preference"` // A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
	Weight int `json:"weight"` // Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
}

// Iok8sapiautoscalingv2MetricSpec represents the Iok8sapiautoscalingv2MetricSpec schema from the OpenAPI specification
type Iok8sapiautoscalingv2MetricSpec struct {
	Pods Iok8sapiautoscalingv2PodsMetricSource `json:"pods,omitempty"` // PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.
	Resource Iok8sapiautoscalingv2ResourceMetricSource `json:"resource,omitempty"` // ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory). The values will be averaged together before being compared to the target. Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source. Only one "target" type should be set.
	TypeField string `json:"type"` // type is the type of metric source. It should be one of "ContainerResource", "External", "Object", "Pods" or "Resource", each mapping to a matching field in the object.
	Containerresource Iok8sapiautoscalingv2ContainerResourceMetricSource `json:"containerResource,omitempty"` // ContainerResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory). The values will be averaged together before being compared to the target. Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source. Only one "target" type should be set.
	External Iok8sapiautoscalingv2ExternalMetricSource `json:"external,omitempty"` // ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
	Object Iok8sapiautoscalingv2ObjectMetricSource `json:"object,omitempty"` // ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
}

// Iok8sapicertificatesv1beta1ClusterTrustBundleList represents the Iok8sapicertificatesv1beta1ClusterTrustBundleList schema from the OpenAPI specification
type Iok8sapicertificatesv1beta1ClusterTrustBundleList struct {
	Items []Iok8sapicertificatesv1beta1ClusterTrustBundle `json:"items"` // items is a collection of ClusterTrustBundle objects
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapiresourcev1beta2DeviceClassSpec represents the Iok8sapiresourcev1beta2DeviceClassSpec schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceClassSpec struct {
	Selectors []Iok8sapiresourcev1beta2DeviceSelector `json:"selectors,omitempty"` // Each selector must be satisfied by a device which is claimed via this class.
	Config []Iok8sapiresourcev1beta2DeviceClassConfiguration `json:"config,omitempty"` // Config defines configuration parameters that apply to each device that is claimed via this class. Some classses may potentially be satisfied by multiple drivers, so each instance of a vendor configuration applies to exactly one driver. They are passed to the driver, but are not considered while allocating the claim.
}

// Iok8sapiappsv1ControllerRevisionList represents the Iok8sapiappsv1ControllerRevisionList schema from the OpenAPI specification
type Iok8sapiappsv1ControllerRevisionList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiappsv1ControllerRevision `json:"items"` // Items is the list of ControllerRevisions
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapibatchv1UncountedTerminatedPods represents the Iok8sapibatchv1UncountedTerminatedPods schema from the OpenAPI specification
type Iok8sapibatchv1UncountedTerminatedPods struct {
	Failed []string `json:"failed,omitempty"` // failed holds UIDs of failed Pods.
	Succeeded []string `json:"succeeded,omitempty"` // succeeded holds UIDs of succeeded Pods.
}

// Iok8sapiadmissionregistrationv1ParamKind represents the Iok8sapiadmissionregistrationv1ParamKind schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ParamKind struct {
	Kind string `json:"kind,omitempty"` // Kind is the API kind the resources belong to. Required.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion is the API group version the resources belong to. In format of "group/version". Required.
}

// Iok8sapicorev1CSIVolumeSource represents the Iok8sapicorev1CSIVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1CSIVolumeSource struct {
	Driver string `json:"driver"` // driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.
	Fstype string `json:"fsType,omitempty"` // fsType to mount. Ex. "ext4", "xfs", "ntfs". If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.
	Nodepublishsecretref Iok8sapicorev1LocalObjectReference `json:"nodePublishSecretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	Readonly bool `json:"readOnly,omitempty"` // readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).
	Volumeattributes map[string]interface{} `json:"volumeAttributes,omitempty"` // volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.
}

// Iok8sapicorev1SleepAction represents the Iok8sapicorev1SleepAction schema from the OpenAPI specification
type Iok8sapicorev1SleepAction struct {
	Seconds int64 `json:"seconds"` // Seconds is the number of seconds to sleep.
}

// Iok8sapicorev1PortStatus represents the Iok8sapicorev1PortStatus schema from the OpenAPI specification
type Iok8sapicorev1PortStatus struct {
	ErrorField string `json:"error,omitempty"` // Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use CamelCase names - cloud provider specific error values must have names that comply with the format foo.example.com/CamelCase.
	Port int `json:"port"` // Port is the port number of the service port of which status is recorded here
	Protocol string `json:"protocol"` // Protocol is the protocol of the service port of which status is recorded here The supported values are: "TCP", "UDP", "SCTP"
}

// Iok8sapicorev1ResourceClaim represents the Iok8sapicorev1ResourceClaim schema from the OpenAPI specification
type Iok8sapicorev1ResourceClaim struct {
	Name string `json:"name"` // Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.
	Request string `json:"request,omitempty"` // Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.
}

// Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyBinding represents the Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyBinding schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyBinding struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyBindingSpec `json:"spec,omitempty"` // ValidatingAdmissionPolicyBindingSpec is the specification of the ValidatingAdmissionPolicyBinding.
}

// Iok8sapiflowcontrolv1Subject represents the Iok8sapiflowcontrolv1Subject schema from the OpenAPI specification
type Iok8sapiflowcontrolv1Subject struct {
	User Iok8sapiflowcontrolv1UserSubject `json:"user,omitempty"` // UserSubject holds detailed information for user-kind subject.
	Group Iok8sapiflowcontrolv1GroupSubject `json:"group,omitempty"` // GroupSubject holds detailed information for group-kind subject.
	Kind string `json:"kind"` // `kind` indicates which one of the other fields is non-empty. Required
	Serviceaccount Iok8sapiflowcontrolv1ServiceAccountSubject `json:"serviceAccount,omitempty"` // ServiceAccountSubject holds detailed information for service-account-kind subject.
}

// Iok8sapiappsv1DeploymentStatus represents the Iok8sapiappsv1DeploymentStatus schema from the OpenAPI specification
type Iok8sapiappsv1DeploymentStatus struct {
	Availablereplicas int `json:"availableReplicas,omitempty"` // Total number of available non-terminating pods (ready for at least minReadySeconds) targeted by this deployment.
	Readyreplicas int `json:"readyReplicas,omitempty"` // Total number of non-terminating pods targeted by this Deployment with a Ready Condition.
	Replicas int `json:"replicas,omitempty"` // Total number of non-terminating pods targeted by this deployment (their labels match the selector).
	Updatedreplicas int `json:"updatedReplicas,omitempty"` // Total number of non-terminating pods targeted by this deployment that have the desired template spec.
	Conditions []Iok8sapiappsv1DeploymentCondition `json:"conditions,omitempty"` // Represents the latest available observations of a deployment's current state.
	Terminatingreplicas int `json:"terminatingReplicas,omitempty"` // Total number of terminating pods targeted by this deployment. Terminating pods have a non-null .metadata.deletionTimestamp and have not yet reached the Failed or Succeeded .status.phase. This is an alpha field. Enable DeploymentReplicaSetTerminatingReplicas to be able to use this field.
	Unavailablereplicas int `json:"unavailableReplicas,omitempty"` // Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created.
	Collisioncount int `json:"collisionCount,omitempty"` // Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The generation observed by the deployment controller.
}

// Iok8sapistoragev1CSIDriverSpec represents the Iok8sapistoragev1CSIDriverSpec schema from the OpenAPI specification
type Iok8sapistoragev1CSIDriverSpec struct {
	Fsgrouppolicy string `json:"fsGroupPolicy,omitempty"` // fsGroupPolicy defines if the underlying volume supports changing ownership and permission of the volume before being mounted. Refer to the specific FSGroupPolicy values for additional details. This field was immutable in Kubernetes < 1.29 and now is mutable. Defaults to ReadWriteOnceWithFSType, which will examine each volume to determine if Kubernetes should modify ownership and permissions of the volume. With the default policy the defined fsGroup will only be applied if a fstype is defined and the volume's access mode contains ReadWriteOnce.
	Podinfoonmount bool `json:"podInfoOnMount,omitempty"` // podInfoOnMount indicates this CSI volume driver requires additional pod information (like podName, podUID, etc.) during mount operations, if set to true. If set to false, pod information will not be passed on mount. Default is false. The CSI driver specifies podInfoOnMount as part of driver deployment. If true, Kubelet will pass pod information as VolumeContext in the CSI NodePublishVolume() calls. The CSI driver is responsible for parsing and validating the information passed in as VolumeContext. The following VolumeContext will be passed if podInfoOnMount is set to true. This list might grow, but the prefix will be used. "csi.storage.k8s.io/pod.name": pod.Name "csi.storage.k8s.io/pod.namespace": pod.Namespace "csi.storage.k8s.io/pod.uid": string(pod.UID) "csi.storage.k8s.io/ephemeral": "true" if the volume is an ephemeral inline volume defined by a CSIVolumeSource, otherwise "false" "csi.storage.k8s.io/ephemeral" is a new feature in Kubernetes 1.16. It is only required for drivers which support both the "Persistent" and "Ephemeral" VolumeLifecycleMode. Other drivers can leave pod info disabled and/or ignore this field. As Kubernetes 1.15 doesn't support this field, drivers can only support one mode when deployed on such a cluster and the deployment determines which mode that is, for example via a command line parameter of the driver. This field was immutable in Kubernetes < 1.29 and now is mutable.
	Storagecapacity bool `json:"storageCapacity,omitempty"` // storageCapacity indicates that the CSI volume driver wants pod scheduling to consider the storage capacity that the driver deployment will report by creating CSIStorageCapacity objects with capacity information, if set to true. The check can be enabled immediately when deploying a driver. In that case, provisioning new volumes with late binding will pause until the driver deployment has published some suitable CSIStorageCapacity object. Alternatively, the driver can be deployed with the field unset or false and it can be flipped later when storage capacity information has been published. This field was immutable in Kubernetes <= 1.22 and now is mutable.
	Volumelifecyclemodes []string `json:"volumeLifecycleModes,omitempty"` // volumeLifecycleModes defines what kind of volumes this CSI volume driver supports. The default if the list is empty is "Persistent", which is the usage defined by the CSI specification and implemented in Kubernetes via the usual PV/PVC mechanism. The other mode is "Ephemeral". In this mode, volumes are defined inline inside the pod spec with CSIVolumeSource and their lifecycle is tied to the lifecycle of that pod. A driver has to be aware of this because it is only going to get a NodePublishVolume call for such a volume. For more information about implementing this mode, see https://kubernetes-csi.github.io/docs/ephemeral-local-volumes.html A driver can support one or more of these modes and more modes may be added in the future. This field is beta. This field is immutable.
	Selinuxmount bool `json:"seLinuxMount,omitempty"` // seLinuxMount specifies if the CSI driver supports "-o context" mount option. When "true", the CSI driver must ensure that all volumes provided by this CSI driver can be mounted separately with different `-o context` options. This is typical for storage backends that provide volumes as filesystems on block devices or as independent shared volumes. Kubernetes will call NodeStage / NodePublish with "-o context=xyz" mount option when mounting a ReadWriteOncePod volume used in Pod that has explicitly set SELinux context. In the future, it may be expanded to other volume AccessModes. In any case, Kubernetes will ensure that the volume is mounted only with a single SELinux context. When "false", Kubernetes won't pass any special SELinux mount options to the driver. This is typical for volumes that represent subdirectories of a bigger shared filesystem. Default is "false".
	Tokenrequests []Iok8sapistoragev1TokenRequest `json:"tokenRequests,omitempty"` // tokenRequests indicates the CSI driver needs pods' service account tokens it is mounting volume for to do necessary authentication. Kubelet will pass the tokens in VolumeContext in the CSI NodePublishVolume calls. The CSI driver should parse and validate the following VolumeContext: "csi.storage.k8s.io/serviceAccount.tokens": { "<audience>": { "token": <token>, "expirationTimestamp": <expiration timestamp in RFC3339>, }, ... } Note: Audience in each TokenRequest should be different and at most one token is empty string. To receive a new token after expiry, RequiresRepublish can be used to trigger NodePublishVolume periodically.
	Attachrequired bool `json:"attachRequired,omitempty"` // attachRequired indicates this CSI volume driver requires an attach operation (because it implements the CSI ControllerPublishVolume() method), and that the Kubernetes attach detach controller should call the attach volume interface which checks the volumeattachment status and waits until the volume is attached before proceeding to mounting. The CSI external-attacher coordinates with CSI volume driver and updates the volumeattachment status when the attach operation is complete. If the value is specified to false, the attach operation will be skipped. Otherwise the attach operation will be called. This field is immutable.
	Nodeallocatableupdateperiodseconds int64 `json:"nodeAllocatableUpdatePeriodSeconds,omitempty"` // nodeAllocatableUpdatePeriodSeconds specifies the interval between periodic updates of the CSINode allocatable capacity for this driver. When set, both periodic updates and updates triggered by capacity-related failures are enabled. If not set, no updates occur (neither periodic nor upon detecting capacity-related failures), and the allocatable.count remains static. The minimum allowed value for this field is 10 seconds. This is an alpha feature and requires the MutableCSINodeAllocatableCount feature gate to be enabled. This field is mutable.
	Requiresrepublish bool `json:"requiresRepublish,omitempty"` // requiresRepublish indicates the CSI driver wants `NodePublishVolume` being periodically called to reflect any possible change in the mounted volume. This field defaults to false. Note: After a successful initial NodePublishVolume call, subsequent calls to NodePublishVolume should only update the contents of the volume. New mount points will not be seen by a running container.
}

// Iok8sapiadmissionregistrationv1ValidatingWebhookConfiguration represents the Iok8sapiadmissionregistrationv1ValidatingWebhookConfiguration schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ValidatingWebhookConfiguration struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Webhooks []Iok8sapiadmissionregistrationv1ValidatingWebhook `json:"webhooks,omitempty"` // Webhooks is a list of webhooks and the affected resources and operations.
}

// Iok8sapiapiserverinternalv1alpha1StorageVersion represents the Iok8sapiapiserverinternalv1alpha1StorageVersion schema from the OpenAPI specification
type Iok8sapiapiserverinternalv1alpha1StorageVersion struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiapiserverinternalv1alpha1StorageVersionSpec `json:"spec"` // StorageVersionSpec is an empty spec.
	Status Iok8sapiapiserverinternalv1alpha1StorageVersionStatus `json:"status"` // API server instances report the versions they can decode and the version they encode objects to when persisting objects in the backend.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapicorev1SELinuxOptions represents the Iok8sapicorev1SELinuxOptions schema from the OpenAPI specification
type Iok8sapicorev1SELinuxOptions struct {
	Level string `json:"level,omitempty"` // Level is SELinux level label that applies to the container.
	Role string `json:"role,omitempty"` // Role is a SELinux role label that applies to the container.
	TypeField string `json:"type,omitempty"` // Type is a SELinux type label that applies to the container.
	User string `json:"user,omitempty"` // User is a SELinux user label that applies to the container.
}

// Iok8sapicorev1WindowsSecurityContextOptions represents the Iok8sapicorev1WindowsSecurityContextOptions schema from the OpenAPI specification
type Iok8sapicorev1WindowsSecurityContextOptions struct {
	Gmsacredentialspecname string `json:"gmsaCredentialSpecName,omitempty"` // GMSACredentialSpecName is the name of the GMSA credential spec to use.
	Hostprocess bool `json:"hostProcess,omitempty"` // HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.
	Runasusername string `json:"runAsUserName,omitempty"` // The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	Gmsacredentialspec string `json:"gmsaCredentialSpec,omitempty"` // GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.
}

// Iok8sapiflowcontrolv1LimitResponse represents the Iok8sapiflowcontrolv1LimitResponse schema from the OpenAPI specification
type Iok8sapiflowcontrolv1LimitResponse struct {
	TypeField string `json:"type"` // `type` is "Queue" or "Reject". "Queue" means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. "Reject" means that requests that can not be executed upon arrival are rejected. Required.
	Queuing Iok8sapiflowcontrolv1QueuingConfiguration `json:"queuing,omitempty"` // QueuingConfiguration holds the configuration parameters for queuing
}

// Iok8sapirbacv1ClusterRoleBinding represents the Iok8sapirbacv1ClusterRoleBinding schema from the OpenAPI specification
type Iok8sapirbacv1ClusterRoleBinding struct {
	Roleref Iok8sapirbacv1RoleRef `json:"roleRef"` // RoleRef contains information that points to the role being used
	Subjects []Iok8sapirbacv1Subject `json:"subjects,omitempty"` // Subjects holds references to the objects the role applies to.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiauthorizationv1SubjectAccessReviewSpec represents the Iok8sapiauthorizationv1SubjectAccessReviewSpec schema from the OpenAPI specification
type Iok8sapiauthorizationv1SubjectAccessReviewSpec struct {
	User string `json:"user,omitempty"` // User is the user you're testing for. If you specify "User" but not "Groups", then is it interpreted as "What if User were not a member of any groups
	Extra map[string]interface{} `json:"extra,omitempty"` // Extra corresponds to the user.Info.GetExtra() method from the authenticator. Since that is input to the authorizer it needs a reflection here.
	Groups []string `json:"groups,omitempty"` // Groups is the groups you're testing for.
	Nonresourceattributes Iok8sapiauthorizationv1NonResourceAttributes `json:"nonResourceAttributes,omitempty"` // NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
	Resourceattributes Iok8sapiauthorizationv1ResourceAttributes `json:"resourceAttributes,omitempty"` // ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
	Uid string `json:"uid,omitempty"` // UID information about the requesting user.
}

// Iok8sapinetworkingv1ServiceBackendPort represents the Iok8sapinetworkingv1ServiceBackendPort schema from the OpenAPI specification
type Iok8sapinetworkingv1ServiceBackendPort struct {
	Number int `json:"number,omitempty"` // number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with "Name".
	Name string `json:"name,omitempty"` // name is the name of the port on the Service. This is a mutually exclusive setting with "Number".
}

// Iok8sapiautoscalingv2MetricIdentifier represents the Iok8sapiautoscalingv2MetricIdentifier schema from the OpenAPI specification
type Iok8sapiautoscalingv2MetricIdentifier struct {
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Name string `json:"name"` // name is the name of the given metric
}

// Iok8sapicorev1Capabilities represents the Iok8sapicorev1Capabilities schema from the OpenAPI specification
type Iok8sapicorev1Capabilities struct {
	Add []string `json:"add,omitempty"` // Added capabilities
	Drop []string `json:"drop,omitempty"` // Removed capabilities
}

// Iok8sapiappsv1DeploymentStrategy represents the Iok8sapiappsv1DeploymentStrategy schema from the OpenAPI specification
type Iok8sapiappsv1DeploymentStrategy struct {
	Rollingupdate Iok8sapiappsv1RollingUpdateDeployment `json:"rollingUpdate,omitempty"` // Spec to control the desired behavior of rolling update.
	TypeField string `json:"type,omitempty"` // Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
}

// Iok8sapicertificatesv1alpha1ClusterTrustBundleList represents the Iok8sapicertificatesv1alpha1ClusterTrustBundleList schema from the OpenAPI specification
type Iok8sapicertificatesv1alpha1ClusterTrustBundleList struct {
	Items []Iok8sapicertificatesv1alpha1ClusterTrustBundle `json:"items"` // items is a collection of ClusterTrustBundle objects
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1ReplicationControllerStatus represents the Iok8sapicorev1ReplicationControllerStatus schema from the OpenAPI specification
type Iok8sapicorev1ReplicationControllerStatus struct {
	Replicas int `json:"replicas"` // Replicas is the most recently observed number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
	Availablereplicas int `json:"availableReplicas,omitempty"` // The number of available replicas (ready for at least minReadySeconds) for this replication controller.
	Conditions []Iok8sapicorev1ReplicationControllerCondition `json:"conditions,omitempty"` // Represents the latest available observations of a replication controller's current state.
	Fullylabeledreplicas int `json:"fullyLabeledReplicas,omitempty"` // The number of pods that have labels matching the labels of the pod template of the replication controller.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // ObservedGeneration reflects the generation of the most recently observed replication controller.
	Readyreplicas int `json:"readyReplicas,omitempty"` // The number of ready replicas for this replication controller.
}

// Iok8sapiautoscalingv2HorizontalPodAutoscalerList represents the Iok8sapiautoscalingv2HorizontalPodAutoscalerList schema from the OpenAPI specification
type Iok8sapiautoscalingv2HorizontalPodAutoscalerList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiautoscalingv2HorizontalPodAutoscaler `json:"items"` // items is the list of horizontal pod autoscaler objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiautoscalingv2ContainerResourceMetricSource represents the Iok8sapiautoscalingv2ContainerResourceMetricSource schema from the OpenAPI specification
type Iok8sapiautoscalingv2ContainerResourceMetricSource struct {
	Container string `json:"container"` // container is the name of the container in the pods of the scaling target
	Name string `json:"name"` // name is the name of the resource in question.
	Target Iok8sapiautoscalingv2MetricTarget `json:"target"` // MetricTarget defines the target value, average value, or average utilization of a specific metric
}

// Iok8sapinetworkingv1IngressTLS represents the Iok8sapinetworkingv1IngressTLS schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressTLS struct {
	Hosts []string `json:"hosts,omitempty"` // hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
	Secretname string `json:"secretName,omitempty"` // secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the "Host" header is used for routing.
}

// Iok8sapinetworkingv1beta1IPAddress represents the Iok8sapinetworkingv1beta1IPAddress schema from the OpenAPI specification
type Iok8sapinetworkingv1beta1IPAddress struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapinetworkingv1beta1IPAddressSpec `json:"spec,omitempty"` // IPAddressSpec describe the attributes in an IP Address.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status GeneratedType `json:"status,omitempty"` // APIServiceStatus contains derived information about an API server
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec GeneratedType `json:"spec,omitempty"` // APIServiceSpec contains information for locating and communicating with a server. Only https is supported, though you are able to disable certificate verification.
}

// Iok8sapicorev1PodList represents the Iok8sapicorev1PodList schema from the OpenAPI specification
type Iok8sapicorev1PodList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1Pod `json:"items"` // List of pods. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiflowcontrolv1UserSubject represents the Iok8sapiflowcontrolv1UserSubject schema from the OpenAPI specification
type Iok8sapiflowcontrolv1UserSubject struct {
	Name string `json:"name"` // `name` is the username that matches, or "*" to match all usernames. Required.
}

// Iok8sapicorev1NFSVolumeSource represents the Iok8sapicorev1NFSVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1NFSVolumeSource struct {
	Path string `json:"path"` // path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Readonly bool `json:"readOnly,omitempty"` // readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Server string `json:"server"` // server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
}

// Iok8sapicorev1Namespace represents the Iok8sapicorev1Namespace schema from the OpenAPI specification
type Iok8sapicorev1Namespace struct {
	Spec Iok8sapicorev1NamespaceSpec `json:"spec,omitempty"` // NamespaceSpec describes the attributes on a Namespace.
	Status Iok8sapicorev1NamespaceStatus `json:"status,omitempty"` // NamespaceStatus is information about the current status of a Namespace.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1ProjectedVolumeSource represents the Iok8sapicorev1ProjectedVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ProjectedVolumeSource struct {
	Defaultmode int `json:"defaultMode,omitempty"` // defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Sources []Iok8sapicorev1VolumeProjection `json:"sources,omitempty"` // sources is the list of volume projections. Each entry in this list handles one source.
}

// Iok8sapiresourcev1beta1DeviceClassConfiguration represents the Iok8sapiresourcev1beta1DeviceClassConfiguration schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceClassConfiguration struct {
	Opaque Iok8sapiresourcev1beta1OpaqueDeviceConfiguration `json:"opaque,omitempty"` // OpaqueDeviceConfiguration contains configuration parameters for a driver in a format defined by the driver vendor.
}

// Iok8sapistoragev1VolumeAttachment represents the Iok8sapistoragev1VolumeAttachment schema from the OpenAPI specification
type Iok8sapistoragev1VolumeAttachment struct {
	Status Iok8sapistoragev1VolumeAttachmentStatus `json:"status,omitempty"` // VolumeAttachmentStatus is the status of a VolumeAttachment request.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapistoragev1VolumeAttachmentSpec `json:"spec"` // VolumeAttachmentSpec is the specification of a VolumeAttachment request.
}

// Iok8sapicorev1Taint represents the Iok8sapicorev1Taint schema from the OpenAPI specification
type Iok8sapicorev1Taint struct {
	Timeadded string `json:"timeAdded,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Value string `json:"value,omitempty"` // The taint value corresponding to the taint key.
	Effect string `json:"effect"` // Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
	Key string `json:"key"` // Required. The taint key to be applied to a node.
}

// Iok8sapiautoscalingv2HorizontalPodAutoscalerCondition represents the Iok8sapiautoscalingv2HorizontalPodAutoscalerCondition schema from the OpenAPI specification
type Iok8sapiautoscalingv2HorizontalPodAutoscalerCondition struct {
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // message is a human-readable explanation containing details about the transition
	Reason string `json:"reason,omitempty"` // reason is the reason for the condition's last transition.
	Status string `json:"status"` // status is the status of the condition (True, False, Unknown)
	TypeField string `json:"type"` // type describes the current condition
}

// Iok8sapicorev1VolumeDevice represents the Iok8sapicorev1VolumeDevice schema from the OpenAPI specification
type Iok8sapicorev1VolumeDevice struct {
	Devicepath string `json:"devicePath"` // devicePath is the path inside of the container that the device will be mapped to.
	Name string `json:"name"` // name must match the name of a persistentVolumeClaim in the pod
}

// Iok8sapiresourcev1beta2DeviceClassConfiguration represents the Iok8sapiresourcev1beta2DeviceClassConfiguration schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceClassConfiguration struct {
	Opaque Iok8sapiresourcev1beta2OpaqueDeviceConfiguration `json:"opaque,omitempty"` // OpaqueDeviceConfiguration contains configuration parameters for a driver in a format defined by the driver vendor.
}

// Iok8sapimachinerypkgapismetav1Status represents the Iok8sapimachinerypkgapismetav1Status schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1Status struct {
	Details Iok8sapimachinerypkgapismetav1StatusDetails `json:"details,omitempty"` // StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Message string `json:"message,omitempty"` // A human-readable description of the status of this operation.
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Reason string `json:"reason,omitempty"` // A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
	Status string `json:"status,omitempty"` // Status of the operation. One of: "Success" or "Failure". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Code int `json:"code,omitempty"` // Suggested HTTP return code for this status, 0 if not set.
}

// Iok8sapidiscoveryv1EndpointSlice represents the Iok8sapidiscoveryv1EndpointSlice schema from the OpenAPI specification
type Iok8sapidiscoveryv1EndpointSlice struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Ports []Iok8sapidiscoveryv1EndpointPort `json:"ports,omitempty"` // ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. Each slice may include a maximum of 100 ports. Services always have at least 1 port, so EndpointSlices generated by the EndpointSlice controller will likewise always have at least 1 port. EndpointSlices used for other purposes may have an empty ports list.
	Addresstype string `json:"addressType"` // addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name. (Deprecated) The EndpointSlice controller only generates, and kube-proxy only processes, slices of addressType "IPv4" and "IPv6". No semantics are defined for the "FQDN" type.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Endpoints []Iok8sapidiscoveryv1Endpoint `json:"endpoints"` // endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
}

// Iok8sapicorev1ResourceStatus represents the Iok8sapicorev1ResourceStatus schema from the OpenAPI specification
type Iok8sapicorev1ResourceStatus struct {
	Name string `json:"name"` // Name of the resource. Must be unique within the pod and in case of non-DRA resource, match one of the resources from the pod spec. For DRA resources, the value must be "claim:<claim_name>/<request>". When this status is reported about a container, the "claim_name" and "request" must match one of the claims of this container.
	Resources []Iok8sapicorev1ResourceHealth `json:"resources,omitempty"` // List of unique resources health. Each element in the list contains an unique resource ID and its health. At a minimum, for the lifetime of a Pod, resource ID must uniquely identify the resource allocated to the Pod on the Node. If other Pod on the same Node reports the status with the same resource ID, it must be the same resource they share. See ResourceID type definition for a specific format it has in various use cases.
}

// Iok8sapistoragev1CSINodeSpec represents the Iok8sapistoragev1CSINodeSpec schema from the OpenAPI specification
type Iok8sapistoragev1CSINodeSpec struct {
	Drivers []Iok8sapistoragev1CSINodeDriver `json:"drivers"` // drivers is a list of information of all CSI Drivers existing on a node. If all drivers in the list are uninstalled, this can become empty.
}

// Iok8sapiresourcev1beta2DeviceRequest represents the Iok8sapiresourcev1beta2DeviceRequest schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceRequest struct {
	Exactly Iok8sapiresourcev1beta2ExactDeviceRequest `json:"exactly,omitempty"` // ExactDeviceRequest is a request for one or more identical devices.
	Firstavailable []Iok8sapiresourcev1beta2DeviceSubRequest `json:"firstAvailable,omitempty"` // FirstAvailable contains subrequests, of which exactly one will be selected by the scheduler. It tries to satisfy them in the order in which they are listed here. So if there are two entries in the list, the scheduler will only check the second one if it determines that the first one can not be used. DRA does not yet implement scoring, so the scheduler will select the first set of devices that satisfies all the requests in the claim. And if the requirements can be satisfied on more than one node, other scheduling features will determine which node is chosen. This means that the set of devices allocated to a claim might not be the optimal set available to the cluster. Scoring will be implemented later.
	Name string `json:"name"` // Name can be used to reference this request in a pod.spec.containers[].resources.claims entry and in a constraint of the claim. References using the name in the DeviceRequest will uniquely identify a request when the Exactly field is set. When the FirstAvailable field is set, a reference to the name of the DeviceRequest will match whatever subrequest is chosen by the scheduler. Must be a DNS label.
}

// Iok8sapicorev1ServiceStatus represents the Iok8sapicorev1ServiceStatus schema from the OpenAPI specification
type Iok8sapicorev1ServiceStatus struct {
	Conditions []Iok8sapimachinerypkgapismetav1Condition `json:"conditions,omitempty"` // Current service state
	Loadbalancer Iok8sapicorev1LoadBalancerStatus `json:"loadBalancer,omitempty"` // LoadBalancerStatus represents the status of a load-balancer.
}

// Iok8sapicorev1SeccompProfile represents the Iok8sapicorev1SeccompProfile schema from the OpenAPI specification
type Iok8sapicorev1SeccompProfile struct {
	Localhostprofile string `json:"localhostProfile,omitempty"` // localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is "Localhost". Must NOT be set for any other type.
	TypeField string `json:"type"` // type indicates which kind of seccomp profile will be applied. Valid options are: Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.
}

// Iok8sapiresourcev1beta1BasicDevice represents the Iok8sapiresourcev1beta1BasicDevice schema from the OpenAPI specification
type Iok8sapiresourcev1beta1BasicDevice struct {
	Nodeselector Iok8sapicorev1NodeSelector `json:"nodeSelector,omitempty"` // A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
	Taints []Iok8sapiresourcev1beta1DeviceTaint `json:"taints,omitempty"` // If specified, these are the driver-defined taints. The maximum number of taints is 4. This is an alpha field and requires enabling the DRADeviceTaints feature gate.
	Allnodes bool `json:"allNodes,omitempty"` // AllNodes indicates that all nodes have access to the device. Must only be set if Spec.PerDeviceNodeSelection is set to true. At most one of NodeName, NodeSelector and AllNodes can be set.
	Attributes map[string]interface{} `json:"attributes,omitempty"` // Attributes defines the set of attributes for this device. The name of each attribute must be unique in that set. The maximum number of attributes and capacities combined is 32.
	Capacity map[string]interface{} `json:"capacity,omitempty"` // Capacity defines the set of capacities for this device. The name of each capacity must be unique in that set. The maximum number of attributes and capacities combined is 32.
	Consumescounters []Iok8sapiresourcev1beta1DeviceCounterConsumption `json:"consumesCounters,omitempty"` // ConsumesCounters defines a list of references to sharedCounters and the set of counters that the device will consume from those counter sets. There can only be a single entry per counterSet. The total number of device counter consumption entries must be <= 32. In addition, the total number in the entire ResourceSlice must be <= 1024 (for example, 64 devices with 16 counters each).
	Nodename string `json:"nodeName,omitempty"` // NodeName identifies the node where the device is available. Must only be set if Spec.PerDeviceNodeSelection is set to true. At most one of NodeName, NodeSelector and AllNodes can be set.
}

// Iok8sapicorev1NodeConfigStatus represents the Iok8sapicorev1NodeConfigStatus schema from the OpenAPI specification
type Iok8sapicorev1NodeConfigStatus struct {
	Active Iok8sapicorev1NodeConfigSource `json:"active,omitempty"` // NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil. This API is deprecated since 1.22
	Assigned Iok8sapicorev1NodeConfigSource `json:"assigned,omitempty"` // NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil. This API is deprecated since 1.22
	ErrorField string `json:"error,omitempty"` // Error describes any problems reconciling the Spec.ConfigSource to the Active config. Errors may occur, for example, attempting to checkpoint Spec.ConfigSource to the local Assigned record, attempting to checkpoint the payload associated with Spec.ConfigSource, attempting to load or validate the Assigned config, etc. Errors may occur at different points while syncing config. Earlier errors (e.g. download or checkpointing errors) will not result in a rollback to LastKnownGood, and may resolve across Kubelet retries. Later errors (e.g. loading or validating a checkpointed config) will result in a rollback to LastKnownGood. In the latter case, it is usually possible to resolve the error by fixing the config assigned in Spec.ConfigSource. You can find additional information for debugging by searching the error message in the Kubelet log. Error is a human-readable description of the error state; machines can check whether or not Error is empty, but should not rely on the stability of the Error text across Kubelet versions.
	Lastknowngood Iok8sapicorev1NodeConfigSource `json:"lastKnownGood,omitempty"` // NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil. This API is deprecated since 1.22
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Conditions []GeneratedType `json:"conditions,omitempty"` // conditions indicate state for particular aspects of a CustomResourceDefinition
	Storedversions []string `json:"storedVersions,omitempty"` // storedVersions lists all versions of CustomResources that were ever persisted. Tracking these versions allows a migration path for stored versions in etcd. The field is mutable so a migration controller can finish a migration to another version (ensuring no old objects are left in storage), and then remove the rest of the versions from this list. Versions may not be removed from `spec.versions` while they exist in this list.
	Acceptednames GeneratedType `json:"acceptedNames,omitempty"` // CustomResourceDefinitionNames indicates the names to serve this CustomResourceDefinition
}

// Iok8sapicorev1ReplicationControllerSpec represents the Iok8sapicorev1ReplicationControllerSpec schema from the OpenAPI specification
type Iok8sapicorev1ReplicationControllerSpec struct {
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	Replicas int `json:"replicas,omitempty"` // Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
	Selector map[string]interface{} `json:"selector,omitempty"` // Selector is a label query over pods that should match the Replicas count. If Selector is empty, it is defaulted to the labels present on the Pod template. Label keys and values that must match in order to be controlled by this replication controller, if empty defaulted to labels on Pod template. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Template Iok8sapicorev1PodTemplateSpec `json:"template,omitempty"` // PodTemplateSpec describes the data a pod should have when created from a template
}

// Iok8sapiresourcev1beta1ResourceClaimTemplateSpec represents the Iok8sapiresourcev1beta1ResourceClaimTemplateSpec schema from the OpenAPI specification
type Iok8sapiresourcev1beta1ResourceClaimTemplateSpec struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiresourcev1beta1ResourceClaimSpec `json:"spec"` // ResourceClaimSpec defines what is being requested in a ResourceClaim and how to configure it.
}

// Iok8sapibatchv1SuccessPolicy represents the Iok8sapibatchv1SuccessPolicy schema from the OpenAPI specification
type Iok8sapibatchv1SuccessPolicy struct {
	Rules []Iok8sapibatchv1SuccessPolicyRule `json:"rules"` // rules represents the list of alternative rules for the declaring the Jobs as successful before `.status.succeeded >= .spec.completions`. Once any of the rules are met, the "SuccessCriteriaMet" condition is added, and the lingering pods are removed. The terminal state for such a Job has the "Complete" condition. Additionally, these rules are evaluated in order; Once the Job meets one of the rules, other rules are ignored. At most 20 elements are allowed.
}

// Iok8sapiresourcev1beta2DeviceCapacity represents the Iok8sapiresourcev1beta2DeviceCapacity schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceCapacity struct {
	Value string `json:"value"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
}

// Iok8sapidiscoveryv1Endpoint represents the Iok8sapidiscoveryv1Endpoint schema from the OpenAPI specification
type Iok8sapidiscoveryv1Endpoint struct {
	Zone string `json:"zone,omitempty"` // zone is the name of the Zone this endpoint exists in.
	Addresses []string `json:"addresses"` // addresses of this endpoint. For EndpointSlices of addressType "IPv4" or "IPv6", the values are IP addresses in canonical form. The syntax and semantics of other addressType values are not defined. This must contain at least one address but no more than 100. EndpointSlices generated by the EndpointSlice controller will always have exactly 1 address. No semantics are defined for additional addresses beyond the first, and kube-proxy does not look at them.
	Conditions Iok8sapidiscoveryv1EndpointConditions `json:"conditions,omitempty"` // EndpointConditions represents the current condition of an endpoint.
	Deprecatedtopology map[string]interface{} `json:"deprecatedTopology,omitempty"` // deprecatedTopology contains topology information part of the v1beta1 API. This field is deprecated, and will be removed when the v1beta1 API is removed (no sooner than kubernetes v1.24). While this field can hold values, it is not writable through the v1 API, and any attempts to write to it will be silently ignored. Topology information can be found in the zone and nodeName fields instead.
	Hints Iok8sapidiscoveryv1EndpointHints `json:"hints,omitempty"` // EndpointHints provides hints describing how an endpoint should be consumed.
	Hostname string `json:"hostname,omitempty"` // hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must be lowercase and pass DNS Label (RFC 1123) validation.
	Nodename string `json:"nodeName,omitempty"` // nodeName represents the name of the Node hosting this endpoint. This can be used to determine endpoints local to a Node.
	Targetref Iok8sapicorev1ObjectReference `json:"targetRef,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
}

// Iok8sapicertificatesv1beta1ClusterTrustBundle represents the Iok8sapicertificatesv1beta1ClusterTrustBundle schema from the OpenAPI specification
type Iok8sapicertificatesv1beta1ClusterTrustBundle struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicertificatesv1beta1ClusterTrustBundleSpec `json:"spec"` // ClusterTrustBundleSpec contains the signer and trust anchors.
}

// Iok8sapimachinerypkgapismetav1ObjectMeta represents the Iok8sapimachinerypkgapismetav1ObjectMeta schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1ObjectMeta struct {
	Generation int64 `json:"generation,omitempty"` // A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.
	Deletiontimestamp string `json:"deletionTimestamp,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Namespace string `json:"namespace,omitempty"` // Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the "default" namespace, but "default" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty. Must be a DNS_LABEL. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces
	Annotations map[string]interface{} `json:"annotations,omitempty"` // Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations
	Name string `json:"name,omitempty"` // Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names
	Resourceversion string `json:"resourceVersion,omitempty"` // An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources. Populated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
	Labels map[string]interface{} `json:"labels,omitempty"` // Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels
	Ownerreferences []Iok8sapimachinerypkgapismetav1OwnerReference `json:"ownerReferences,omitempty"` // List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.
	Selflink string `json:"selfLink,omitempty"` // Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.
	Creationtimestamp string `json:"creationTimestamp,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Deletiongraceperiodseconds int64 `json:"deletionGracePeriodSeconds,omitempty"` // Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.
	Managedfields []Iok8sapimachinerypkgapismetav1ManagedFieldsEntry `json:"managedFields,omitempty"` // ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like "ci-cd". The set of fields is always in the version that the workflow used when modifying the object.
	Uid string `json:"uid,omitempty"` // UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations. Populated by the system. Read-only. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids
	Finalizers []string `json:"finalizers,omitempty"` // Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.
	Generatename string `json:"generateName,omitempty"` // GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server. If this field is specified and the generated name exists, the server will return a 409. Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency
}

// Iok8sapibatchv1CronJobStatus represents the Iok8sapibatchv1CronJobStatus schema from the OpenAPI specification
type Iok8sapibatchv1CronJobStatus struct {
	Active []Iok8sapicorev1ObjectReference `json:"active,omitempty"` // A list of pointers to currently running jobs.
	Lastscheduletime string `json:"lastScheduleTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Lastsuccessfultime string `json:"lastSuccessfulTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
}

// Iok8sapinetworkingv1NetworkPolicyPort represents the Iok8sapinetworkingv1NetworkPolicyPort schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicyPort struct {
	Port string `json:"port,omitempty"` // IntOrString is a type that can hold an int32 or a string. When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type. This allows you to have, for example, a JSON field that can accept a name or number.
	Protocol string `json:"protocol,omitempty"` // protocol represents the protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.
	Endport int `json:"endPort,omitempty"` // endPort indicates that the range of ports from port to endPort if set, inclusive, should be allowed by the policy. This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port.
}

// Iok8sapicorev1EphemeralContainer represents the Iok8sapicorev1EphemeralContainer schema from the OpenAPI specification
type Iok8sapicorev1EphemeralContainer struct {
	Args []string `json:"args,omitempty"` // Arguments to the entrypoint. The image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Resources Iok8sapicorev1ResourceRequirements `json:"resources,omitempty"` // ResourceRequirements describes the compute resource requirements.
	Stdinonce bool `json:"stdinOnce,omitempty"` // Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
	Tty bool `json:"tty,omitempty"` // Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
	Securitycontext Iok8sapicorev1SecurityContext `json:"securityContext,omitempty"` // SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext. When both are set, the values in SecurityContext take precedence.
	Command []string `json:"command,omitempty"` // Entrypoint array. Not executed within a shell. The image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Envfrom []Iok8sapicorev1EnvFromSource `json:"envFrom,omitempty"` // List of sources to populate environment variables in the container. The keys defined within a source may consist of any printable ASCII characters except '='. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
	Readinessprobe Iok8sapicorev1Probe `json:"readinessProbe,omitempty"` // Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
	Startupprobe Iok8sapicorev1Probe `json:"startupProbe,omitempty"` // Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
	Livenessprobe Iok8sapicorev1Probe `json:"livenessProbe,omitempty"` // Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
	Image string `json:"image,omitempty"` // Container image name. More info: https://kubernetes.io/docs/concepts/containers/images
	Imagepullpolicy string `json:"imagePullPolicy,omitempty"` // Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	Lifecycle Iok8sapicorev1Lifecycle `json:"lifecycle,omitempty"` // Lifecycle describes actions that the management system should take in response to container lifecycle events. For the PostStart and PreStop lifecycle handlers, management of the container blocks until the action is complete, unless the container process fails, in which case the handler is aborted.
	Resizepolicy []Iok8sapicorev1ContainerResizePolicy `json:"resizePolicy,omitempty"` // Resources resize policy for the container.
	Terminationmessagepath string `json:"terminationMessagePath,omitempty"` // Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
	Name string `json:"name"` // Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.
	Targetcontainername string `json:"targetContainerName,omitempty"` // If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container uses the namespaces configured in the Pod spec. The container runtime must implement support for this feature. If the runtime does not support namespace targeting then the result of setting this field is undefined.
	Terminationmessagepolicy string `json:"terminationMessagePolicy,omitempty"` // Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
	Env []Iok8sapicorev1EnvVar `json:"env,omitempty"` // List of environment variables to set in the container. Cannot be updated.
	Restartpolicy string `json:"restartPolicy,omitempty"` // Restart policy for the container to manage the restart behavior of each container within a pod. This may only be set for init containers. You cannot set this field on ephemeral containers.
	Stdin bool `json:"stdin,omitempty"` // Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
	Volumedevices []Iok8sapicorev1VolumeDevice `json:"volumeDevices,omitempty"` // volumeDevices is the list of block devices to be used by the container.
	Volumemounts []Iok8sapicorev1VolumeMount `json:"volumeMounts,omitempty"` // Pod volumes to mount into the container's filesystem. Subpath mounts are not allowed for ephemeral containers. Cannot be updated.
	Workingdir string `json:"workingDir,omitempty"` // Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	Ports []Iok8sapicorev1ContainerPort `json:"ports,omitempty"` // Ports are not allowed for ephemeral containers.
}

// Iok8sapicertificatesv1CertificateSigningRequestStatus represents the Iok8sapicertificatesv1CertificateSigningRequestStatus schema from the OpenAPI specification
type Iok8sapicertificatesv1CertificateSigningRequestStatus struct {
	Certificate string `json:"certificate,omitempty"` // certificate is populated with an issued certificate by the signer after an Approved condition is present. This field is set via the /status subresource. Once populated, this field is immutable. If the certificate signing request is denied, a condition of type "Denied" is added and this field remains empty. If the signer cannot issue the certificate, a condition of type "Failed" is added and this field remains empty. Validation requirements: 1. certificate must contain one or more PEM blocks. 2. All PEM blocks must have the "CERTIFICATE" label, contain no headers, and the encoded data must be a BER-encoded ASN.1 Certificate structure as described in section 4 of RFC5280. 3. Non-PEM content may appear before or after the "CERTIFICATE" PEM blocks and is unvalidated, to allow for explanatory text as described in section 5.2 of RFC7468. If more than one PEM block is present, and the definition of the requested spec.signerName does not indicate otherwise, the first block is the issued certificate, and subsequent blocks should be treated as intermediate certificates and presented in TLS handshakes. The certificate is encoded in PEM format. When serialized as JSON or YAML, the data is additionally base64-encoded, so it consists of: base64( -----BEGIN CERTIFICATE----- ... -----END CERTIFICATE----- )
	Conditions []Iok8sapicertificatesv1CertificateSigningRequestCondition `json:"conditions,omitempty"` // conditions applied to the request. Known conditions are "Approved", "Denied", and "Failed".
}

// Iok8sapicorev1VolumeNodeAffinity represents the Iok8sapicorev1VolumeNodeAffinity schema from the OpenAPI specification
type Iok8sapicorev1VolumeNodeAffinity struct {
	Required Iok8sapicorev1NodeSelector `json:"required,omitempty"` // A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
}

// Iok8sapirbacv1RoleList represents the Iok8sapirbacv1RoleList schema from the OpenAPI specification
type Iok8sapirbacv1RoleList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapirbacv1Role `json:"items"` // Items is a list of Roles
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapiadmissionregistrationv1Validation represents the Iok8sapiadmissionregistrationv1Validation schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1Validation struct {
	Reason string `json:"reason,omitempty"` // Reason represents a machine-readable description of why this validation failed. If this is the first validation in the list to fail, this reason, as well as the corresponding HTTP response code, are used in the HTTP response to the client. The currently supported reasons are: "Unauthorized", "Forbidden", "Invalid", "RequestEntityTooLarge". If not set, StatusReasonInvalid is used in the response to the client.
	Expression string `json:"expression"` // Expression represents the expression which will be evaluated by CEL. ref: https://github.com/google/cel-spec CEL expressions have access to the contents of the API request/response, organized into CEL variables as well as some other useful variables: - 'object' - The object from the incoming request. The value is null for DELETE requests. - 'oldObject' - The existing object. The value is null for CREATE requests. - 'request' - Attributes of the API request([ref](/pkg/apis/admission/types.go#AdmissionRequest)). - 'params' - Parameter resource referred to by the policy binding being evaluated. Only populated if the policy has a ParamKind. - 'namespaceObject' - The namespace object that the incoming object belongs to. The value is null for cluster-scoped resources. - 'variables' - Map of composited variables, from its name to its lazily evaluated value. For example, a variable named 'foo' can be accessed as 'variables.foo'. - 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request. See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz - 'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the request resource. The `apiVersion`, `kind`, `metadata.name` and `metadata.generateName` are always accessible from the root of the object. No other metadata properties are accessible. Only property names of the form `[a-zA-Z_.-/][a-zA-Z0-9_.-/]*` are accessible. Accessible property names are escaped according to the following rules when accessed in the expression: - '__' escapes to '__underscores__' - '.' escapes to '__dot__' - '-' escapes to '__dash__' - '/' escapes to '__slash__' - Property names that exactly match a CEL RESERVED keyword escape to '__{keyword}__'. The keywords are: 	 "true", "false", "null", "in", "as", "break", "const", "continue", "else", "for", "function", "if", 	 "import", "let", "loop", "package", "namespace", "return". Examples: - Expression accessing a property named "namespace": {"Expression": "object.__namespace__ > 0"} - Expression accessing a property named "x-prop": {"Expression": "object.x__dash__prop > 0"} - Expression accessing a property named "redact__d": {"Expression": "object.redact__underscores__d > 0"} Equality on arrays with list type of 'set' or 'map' ignores element order, i.e. [1, 2] == [2, 1]. Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type: - 'set': `X + Y` performs a union where the array positions of all elements in `X` are preserved and non-intersecting elements in `Y` are appended, retaining their partial order. - 'map': `X + Y` performs a merge where the array positions of all keys in `X` are preserved but the values are overwritten by values in `Y` when the key sets of `X` and `Y` intersect. Elements in `Y` with non-intersecting keys are appended, retaining their partial order. Required.
	Message string `json:"message,omitempty"` // Message represents the message displayed when validation fails. The message is required if the Expression contains line breaks. The message must not contain line breaks. If unset, the message is "failed rule: {Rule}". e.g. "must be a URL with the host matching spec.host" If the Expression contains line breaks. Message is required. The message must not contain line breaks. If unset, the message is "failed Expression: {Expression}".
	Messageexpression string `json:"messageExpression,omitempty"` // messageExpression declares a CEL expression that evaluates to the validation failure message that is returned when this rule fails. Since messageExpression is used as a failure message, it must evaluate to a string. If both message and messageExpression are present on a validation, then messageExpression will be used if validation fails. If messageExpression results in a runtime error, the runtime error is logged, and the validation failure message is produced as if the messageExpression field were unset. If messageExpression evaluates to an empty string, a string with only spaces, or a string that contains line breaks, then the validation failure message will also be produced as if the messageExpression field were unset, and the fact that messageExpression produced an empty string/string with only spaces/string with line breaks will be logged. messageExpression has access to all the same variables as the `expression` except for 'authorizer' and 'authorizer.requestResource'. Example: "object.x must be less than max ("+string(params.max)+")"
}

// Iok8sapirbacv1RoleBindingList represents the Iok8sapirbacv1RoleBindingList schema from the OpenAPI specification
type Iok8sapirbacv1RoleBindingList struct {
	Items []Iok8sapirbacv1RoleBinding `json:"items"` // Items is a list of RoleBindings
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapiadmissionregistrationv1ValidatingWebhook represents the Iok8sapiadmissionregistrationv1ValidatingWebhook schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ValidatingWebhook struct {
	Objectselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"objectSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Sideeffects string `json:"sideEffects"` // SideEffects states whether this webhook has side effects. Acceptable values are: None, NoneOnDryRun (webhooks created via v1beta1 may also specify Some or Unknown). Webhooks with side effects MUST implement a reconciliation system, since a request may be rejected by a future step in the admission chain and the side effects therefore need to be undone. Requests with the dryRun attribute will be auto-rejected if they match a webhook with sideEffects == Unknown or Some.
	Admissionreviewversions []string `json:"admissionReviewVersions"` // AdmissionReviewVersions is an ordered list of preferred `AdmissionReview` versions the Webhook expects. API server will try to use first version in the list which it supports. If none of the versions specified in this list supported by API server, validation will fail for this object. If a persisted webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail and be subject to the failure policy.
	Matchconditions []Iok8sapiadmissionregistrationv1MatchCondition `json:"matchConditions,omitempty"` // MatchConditions is a list of conditions that must be met for a request to be sent to this webhook. Match conditions filter requests that have already been matched by the rules, namespaceSelector, and objectSelector. An empty list of matchConditions matches all requests. There are a maximum of 64 match conditions allowed. The exact matching logic is (in order): 1. If ANY matchCondition evaluates to FALSE, the webhook is skipped. 2. If ALL matchConditions evaluate to TRUE, the webhook is called. 3. If any matchCondition evaluates to an error (but none are FALSE): - If failurePolicy=Fail, reject the request - If failurePolicy=Ignore, the error is ignored and the webhook is skipped
	Timeoutseconds int `json:"timeoutSeconds,omitempty"` // TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook call will be ignored or the API call will fail based on the failure policy. The timeout value must be between 1 and 30 seconds. Default to 10 seconds.
	Clientconfig Iok8sapiadmissionregistrationv1WebhookClientConfig `json:"clientConfig"` // WebhookClientConfig contains the information to make a TLS connection with the webhook
	Failurepolicy string `json:"failurePolicy,omitempty"` // FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed values are Ignore or Fail. Defaults to Fail.
	Matchpolicy string `json:"matchPolicy,omitempty"` // matchPolicy defines how the "rules" list is used to match incoming requests. Allowed values are "Exact" or "Equivalent". - Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the webhook. - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the webhook. Defaults to "Equivalent"
	Namespaceselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"namespaceSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Rules []Iok8sapiadmissionregistrationv1RuleWithOperations `json:"rules,omitempty"` // Rules describes what operations on what resources/subresources the webhook cares about. The webhook cares about an operation if it matches _any_ Rule. However, in order to prevent ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks from putting the cluster in a state which cannot be recovered from without completely disabling the plugin, ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called on admission requests for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.
	Name string `json:"name"` // The name of the admission webhook. Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where "imagepolicy" is the name of the webhook, and kubernetes.io is the name of the organization. Required.
}

// Iok8sapiauthorizationv1SelfSubjectAccessReviewSpec represents the Iok8sapiauthorizationv1SelfSubjectAccessReviewSpec schema from the OpenAPI specification
type Iok8sapiauthorizationv1SelfSubjectAccessReviewSpec struct {
	Nonresourceattributes Iok8sapiauthorizationv1NonResourceAttributes `json:"nonResourceAttributes,omitempty"` // NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
	Resourceattributes Iok8sapiauthorizationv1ResourceAttributes `json:"resourceAttributes,omitempty"` // ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
}

// Iok8sapicorev1NodeAffinity represents the Iok8sapicorev1NodeAffinity schema from the OpenAPI specification
type Iok8sapicorev1NodeAffinity struct {
	Preferredduringschedulingignoredduringexecution []Iok8sapicorev1PreferredSchedulingTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"` // The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
	Requiredduringschedulingignoredduringexecution Iok8sapicorev1NodeSelector `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"` // A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
}

// Iok8sapicorev1EventList represents the Iok8sapicorev1EventList schema from the OpenAPI specification
type Iok8sapicorev1EventList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1Event `json:"items"` // List of events
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapicorev1LimitRange represents the Iok8sapicorev1LimitRange schema from the OpenAPI specification
type Iok8sapicorev1LimitRange struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1LimitRangeSpec `json:"spec,omitempty"` // LimitRangeSpec defines a min/max usage limit for resources that match on kind.
}

// Iok8sapidiscoveryv1ForNode represents the Iok8sapidiscoveryv1ForNode schema from the OpenAPI specification
type Iok8sapidiscoveryv1ForNode struct {
	Name string `json:"name"` // name represents the name of the node.
}

// Iok8sapinetworkingv1NetworkPolicyEgressRule represents the Iok8sapinetworkingv1NetworkPolicyEgressRule schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicyEgressRule struct {
	Ports []Iok8sapinetworkingv1NetworkPolicyPort `json:"ports,omitempty"` // ports is a list of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.
	To []Iok8sapinetworkingv1NetworkPolicyPeer `json:"to,omitempty"` // to is a list of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.
}

// Iok8sapiautoscalingv1ScaleSpec represents the Iok8sapiautoscalingv1ScaleSpec schema from the OpenAPI specification
type Iok8sapiautoscalingv1ScaleSpec struct {
	Replicas int `json:"replicas,omitempty"` // replicas is the desired number of instances for the scaled object.
}

// Iok8sapinetworkingv1beta1IPAddressSpec represents the Iok8sapinetworkingv1beta1IPAddressSpec schema from the OpenAPI specification
type Iok8sapinetworkingv1beta1IPAddressSpec struct {
	Parentref Iok8sapinetworkingv1beta1ParentReference `json:"parentRef"` // ParentReference describes a reference to a parent object.
}

// Iok8sapistoragemigrationv1alpha1MigrationCondition represents the Iok8sapistoragemigrationv1alpha1MigrationCondition schema from the OpenAPI specification
type Iok8sapistoragemigrationv1alpha1MigrationCondition struct {
	Reason string `json:"reason,omitempty"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of the condition.
	Lastupdatetime string `json:"lastUpdateTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // A human readable message indicating details about the transition.
}

// Iok8sapicorev1NamespaceStatus represents the Iok8sapicorev1NamespaceStatus schema from the OpenAPI specification
type Iok8sapicorev1NamespaceStatus struct {
	Conditions []Iok8sapicorev1NamespaceCondition `json:"conditions,omitempty"` // Represents the latest available observations of a namespace's current state.
	Phase string `json:"phase,omitempty"` // Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
}

// Iok8sapiflowcontrolv1PolicyRulesWithSubjects represents the Iok8sapiflowcontrolv1PolicyRulesWithSubjects schema from the OpenAPI specification
type Iok8sapiflowcontrolv1PolicyRulesWithSubjects struct {
	Subjects []Iok8sapiflowcontrolv1Subject `json:"subjects"` // subjects is the list of normal user, serviceaccount, or group that this rule cares about. There must be at least one member in this slice. A slice that includes both the system:authenticated and system:unauthenticated user groups matches every request. Required.
	Nonresourcerules []Iok8sapiflowcontrolv1NonResourcePolicyRule `json:"nonResourceRules,omitempty"` // `nonResourceRules` is a list of NonResourcePolicyRules that identify matching requests according to their verb and the target non-resource URL.
	Resourcerules []Iok8sapiflowcontrolv1ResourcePolicyRule `json:"resourceRules,omitempty"` // `resourceRules` is a slice of ResourcePolicyRules that identify matching requests according to their verb and the target resource. At least one of `resourceRules` and `nonResourceRules` has to be non-empty.
}

// Iok8sapicorev1HostIP represents the Iok8sapicorev1HostIP schema from the OpenAPI specification
type Iok8sapicorev1HostIP struct {
	Ip string `json:"ip"` // IP is the IP address assigned to the host
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Specreplicaspath string `json:"specReplicasPath"` // specReplicasPath defines the JSON path inside of a custom resource that corresponds to Scale `spec.replicas`. Only JSON paths without the array notation are allowed. Must be a JSON Path under `.spec`. If there is no value under the given path in the custom resource, the `/scale` subresource will return an error on GET.
	Statusreplicaspath string `json:"statusReplicasPath"` // statusReplicasPath defines the JSON path inside of a custom resource that corresponds to Scale `status.replicas`. Only JSON paths without the array notation are allowed. Must be a JSON Path under `.status`. If there is no value under the given path in the custom resource, the `status.replicas` value in the `/scale` subresource will default to 0.
	Labelselectorpath string `json:"labelSelectorPath,omitempty"` // labelSelectorPath defines the JSON path inside of a custom resource that corresponds to Scale `status.selector`. Only JSON paths without the array notation are allowed. Must be a JSON Path under `.status` or `.spec`. Must be set to work with HorizontalPodAutoscaler. The field pointed by this JSON path must be a string field (not a complex selector struct) which contains a serialized label selector in string form. More info: https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions#scale-subresource If there is no value under the given path in the custom resource, the `status.selector` value in the `/scale` subresource will default to the empty string.
}

// Iok8sapiresourcev1beta2Device represents the Iok8sapiresourcev1beta2Device schema from the OpenAPI specification
type Iok8sapiresourcev1beta2Device struct {
	Consumescounters []Iok8sapiresourcev1beta2DeviceCounterConsumption `json:"consumesCounters,omitempty"` // ConsumesCounters defines a list of references to sharedCounters and the set of counters that the device will consume from those counter sets. There can only be a single entry per counterSet. The total number of device counter consumption entries must be <= 32. In addition, the total number in the entire ResourceSlice must be <= 1024 (for example, 64 devices with 16 counters each).
	Name string `json:"name"` // Name is unique identifier among all devices managed by the driver in the pool. It must be a DNS label.
	Nodename string `json:"nodeName,omitempty"` // NodeName identifies the node where the device is available. Must only be set if Spec.PerDeviceNodeSelection is set to true. At most one of NodeName, NodeSelector and AllNodes can be set.
	Nodeselector Iok8sapicorev1NodeSelector `json:"nodeSelector,omitempty"` // A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
	Taints []Iok8sapiresourcev1beta2DeviceTaint `json:"taints,omitempty"` // If specified, these are the driver-defined taints. The maximum number of taints is 4. This is an alpha field and requires enabling the DRADeviceTaints feature gate.
	Allnodes bool `json:"allNodes,omitempty"` // AllNodes indicates that all nodes have access to the device. Must only be set if Spec.PerDeviceNodeSelection is set to true. At most one of NodeName, NodeSelector and AllNodes can be set.
	Attributes map[string]interface{} `json:"attributes,omitempty"` // Attributes defines the set of attributes for this device. The name of each attribute must be unique in that set. The maximum number of attributes and capacities combined is 32.
	Capacity map[string]interface{} `json:"capacity,omitempty"` // Capacity defines the set of capacities for this device. The name of each capacity must be unique in that set. The maximum number of attributes and capacities combined is 32.
}

// Iok8sapicorev1FlockerVolumeSource represents the Iok8sapicorev1FlockerVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1FlockerVolumeSource struct {
	Datasetname string `json:"datasetName,omitempty"` // datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated
	Datasetuuid string `json:"datasetUUID,omitempty"` // datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset
}

// Iok8sapiappsv1RollingUpdateDaemonSet represents the Iok8sapiappsv1RollingUpdateDaemonSet schema from the OpenAPI specification
type Iok8sapiappsv1RollingUpdateDaemonSet struct {
	Maxsurge string `json:"maxSurge,omitempty"` // IntOrString is a type that can hold an int32 or a string. When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type. This allows you to have, for example, a JSON field that can accept a name or number.
	Maxunavailable string `json:"maxUnavailable,omitempty"` // IntOrString is a type that can hold an int32 or a string. When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type. This allows you to have, for example, a JSON field that can accept a name or number.
}

// Iok8sapiresourcev1beta2AllocatedDeviceStatus represents the Iok8sapiresourcev1beta2AllocatedDeviceStatus schema from the OpenAPI specification
type Iok8sapiresourcev1beta2AllocatedDeviceStatus struct {
	Driver string `json:"driver"` // Driver specifies the name of the DRA driver whose kubelet plugin should be invoked to process the allocation once the claim is needed on a node. Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver.
	Networkdata Iok8sapiresourcev1beta2NetworkDeviceData `json:"networkData,omitempty"` // NetworkDeviceData provides network-related details for the allocated device. This information may be filled by drivers or other components to configure or identify the device within a network context.
	Pool string `json:"pool"` // This name together with the driver name and the device name field identify which device was allocated (`<driver name>/<pool name>/<device name>`). Must not be longer than 253 characters and may contain one or more DNS sub-domains separated by slashes.
	Conditions []Iok8sapimachinerypkgapismetav1Condition `json:"conditions,omitempty"` // Conditions contains the latest observation of the device's state. If the device has been configured according to the class and claim config references, the `Ready` condition should be True. Must not contain more than 8 entries.
	Data Iok8sapimachinerypkgruntimeRawExtension `json:"data,omitempty"` // RawExtension is used to hold extensions in external versions. To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types. // Internal package: 	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.Object `json:"myPlugin"` 	} 	type PluginA struct { 		AOption string `json:"aOption"` 	} // External package: 	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.RawExtension `json:"myPlugin"` 	} 	type PluginA struct { 		AOption string `json:"aOption"` 	} // On the wire, the JSON will look something like this: 	{ 		"kind":"MyAPIObject", 		"apiVersion":"v1", 		"myPlugin": { 			"kind":"PluginA", 			"aOption":"foo", 		}, 	} So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)
	Device string `json:"device"` // Device references one device instance via its name in the driver's resource pool. It must be a DNS label.
}

// Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicySpec represents the Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicySpec schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicySpec struct {
	Matchconstraints Iok8sapiadmissionregistrationv1MatchResources `json:"matchConstraints,omitempty"` // MatchResources decides whether to run the admission control policy on an object based on whether it meets the match criteria. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)
	Paramkind Iok8sapiadmissionregistrationv1ParamKind `json:"paramKind,omitempty"` // ParamKind is a tuple of Group Kind and Version.
	Validations []Iok8sapiadmissionregistrationv1Validation `json:"validations,omitempty"` // Validations contain CEL expressions which is used to apply the validation. Validations and AuditAnnotations may not both be empty; a minimum of one Validations or AuditAnnotations is required.
	Variables []Iok8sapiadmissionregistrationv1Variable `json:"variables,omitempty"` // Variables contain definitions of variables that can be used in composition of other expressions. Each variable is defined as a named CEL expression. The variables defined here will be available under `variables` in other expressions of the policy except MatchConditions because MatchConditions are evaluated before the rest of the policy. The expression of a variable can refer to other variables defined earlier in the list but not those after. Thus, Variables must be sorted by the order of first appearance and acyclic.
	Auditannotations []Iok8sapiadmissionregistrationv1AuditAnnotation `json:"auditAnnotations,omitempty"` // auditAnnotations contains CEL expressions which are used to produce audit annotations for the audit event of the API request. validations and auditAnnotations may not both be empty; a least one of validations or auditAnnotations is required.
	Failurepolicy string `json:"failurePolicy,omitempty"` // failurePolicy defines how to handle failures for the admission policy. Failures can occur from CEL expression parse errors, type check errors, runtime errors and invalid or mis-configured policy definitions or bindings. A policy is invalid if spec.paramKind refers to a non-existent Kind. A binding is invalid if spec.paramRef.name refers to a non-existent resource. failurePolicy does not define how validations that evaluate to false are handled. When failurePolicy is set to Fail, ValidatingAdmissionPolicyBinding validationActions define how failures are enforced. Allowed values are Ignore or Fail. Defaults to Fail.
	Matchconditions []Iok8sapiadmissionregistrationv1MatchCondition `json:"matchConditions,omitempty"` // MatchConditions is a list of conditions that must be met for a request to be validated. Match conditions filter requests that have already been matched by the rules, namespaceSelector, and objectSelector. An empty list of matchConditions matches all requests. There are a maximum of 64 match conditions allowed. If a parameter object is provided, it can be accessed via the `params` handle in the same manner as validation expressions. The exact matching logic is (in order): 1. If ANY matchCondition evaluates to FALSE, the policy is skipped. 2. If ALL matchConditions evaluate to TRUE, the policy is evaluated. 3. If any matchCondition evaluates to an error (but none are FALSE): - If failurePolicy=Fail, reject the request - If failurePolicy=Ignore, the policy is skipped
}

// Iok8sapiresourcev1beta1DeviceClass represents the Iok8sapiresourcev1beta1DeviceClass schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceClass struct {
	Spec Iok8sapiresourcev1beta1DeviceClassSpec `json:"spec"` // DeviceClassSpec is used in a [DeviceClass] to define what can be allocated and how to configure it.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiappsv1DaemonSet represents the Iok8sapiappsv1DaemonSet schema from the OpenAPI specification
type Iok8sapiappsv1DaemonSet struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1DaemonSetSpec `json:"spec,omitempty"` // DaemonSetSpec is the specification of a daemon set.
	Status Iok8sapiappsv1DaemonSetStatus `json:"status,omitempty"` // DaemonSetStatus represents the current status of a daemon set.
}

// Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyStatus represents the Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyStatus schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyStatus struct {
	Conditions []Iok8sapimachinerypkgapismetav1Condition `json:"conditions,omitempty"` // The conditions represent the latest available observations of a policy's current state.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // The generation observed by the controller.
	Typechecking Iok8sapiadmissionregistrationv1TypeChecking `json:"typeChecking,omitempty"` // TypeChecking contains results of type checking the expressions in the ValidatingAdmissionPolicy
}

// Iok8sapiappsv1ReplicaSetSpec represents the Iok8sapiappsv1ReplicaSetSpec schema from the OpenAPI specification
type Iok8sapiappsv1ReplicaSetSpec struct {
	Replicas int `json:"replicas,omitempty"` // Replicas is the number of desired pods. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicaset
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Template Iok8sapicorev1PodTemplateSpec `json:"template,omitempty"` // PodTemplateSpec describes the data a pod should have when created from a template
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
}

// Iok8sapiresourcev1beta2DeviceAttribute represents the Iok8sapiresourcev1beta2DeviceAttribute schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceAttribute struct {
	BoolField bool `json:"bool,omitempty"` // BoolValue is a true/false value.
	IntField int64 `json:"int,omitempty"` // IntValue is a number.
	StringField string `json:"string,omitempty"` // StringValue is a string. Must not be longer than 64 characters.
	Version string `json:"version,omitempty"` // VersionValue is a semantic version according to semver.org spec 2.0.0. Must not be longer than 64 characters.
}

// Iok8sapinetworkingv1ServiceCIDRStatus represents the Iok8sapinetworkingv1ServiceCIDRStatus schema from the OpenAPI specification
type Iok8sapinetworkingv1ServiceCIDRStatus struct {
	Conditions []Iok8sapimachinerypkgapismetav1Condition `json:"conditions,omitempty"` // conditions holds an array of metav1.Condition that describe the state of the ServiceCIDR. Current service state
}

// Iok8sapicorev1ConfigMapVolumeSource represents the Iok8sapicorev1ConfigMapVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ConfigMapVolumeSource struct {
	Items []Iok8sapicorev1KeyToPath `json:"items,omitempty"` // items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
	Name string `json:"name,omitempty"` // Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // optional specify whether the ConfigMap or its keys must be defined
	Defaultmode int `json:"defaultMode,omitempty"` // defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
}

// Iok8sapicorev1ContainerImage represents the Iok8sapicorev1ContainerImage schema from the OpenAPI specification
type Iok8sapicorev1ContainerImage struct {
	Names []string `json:"names,omitempty"` // Names by which this image is known. e.g. ["kubernetes.example/hyperkube:v1.0.7", "cloud-vendor.registry.example/cloud-vendor/hyperkube:v1.0.7"]
	Sizebytes int64 `json:"sizeBytes,omitempty"` // The size of the image in bytes.
}

// Iok8sapimachinerypkgapismetav1FieldsV1 represents the Iok8sapimachinerypkgapismetav1FieldsV1 schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1FieldsV1 struct {
}

// Iok8sapicorev1PodCondition represents the Iok8sapicorev1PodCondition schema from the OpenAPI specification
type Iok8sapicorev1PodCondition struct {
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // If set, this represents the .metadata.generation that the pod condition was set based upon. This is an alpha field. Enable PodObservedGenerationTracking to be able to use this field.
	Reason string `json:"reason,omitempty"` // Unique, one-word, CamelCase reason for the condition's last transition.
	Status string `json:"status"` // Status is the status of the condition. Can be True, False, Unknown. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	TypeField string `json:"type"` // Type is the type of the condition. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Lastprobetime string `json:"lastProbeTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // Human-readable message indicating details about last transition.
}

// Iok8sapiappsv1StatefulSetList represents the Iok8sapiappsv1StatefulSetList schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSetList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiappsv1StatefulSet `json:"items"` // Items is the list of stateful sets.
}

// Iok8sapiadmissionregistrationv1TypeChecking represents the Iok8sapiadmissionregistrationv1TypeChecking schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1TypeChecking struct {
	Expressionwarnings []Iok8sapiadmissionregistrationv1ExpressionWarning `json:"expressionWarnings,omitempty"` // The type checking warnings for each expression.
}

// Iok8sapicorev1ConfigMap represents the Iok8sapicorev1ConfigMap schema from the OpenAPI specification
type Iok8sapicorev1ConfigMap struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Binarydata map[string]interface{} `json:"binaryData,omitempty"` // BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
	Data map[string]interface{} `json:"data,omitempty"` // Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
	Immutable bool `json:"immutable,omitempty"` // Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reason string `json:"reason,omitempty"` // reason is a unique, one-word, CamelCase reason for the condition's last transition.
	Status string `json:"status"` // status is the status of the condition. Can be True, False, Unknown.
	TypeField string `json:"type"` // type is the type of the condition. Types include Established, NamesAccepted and Terminating.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // message is a human-readable message indicating details about last transition.
}

// Iok8sapicorev1LocalVolumeSource represents the Iok8sapicorev1LocalVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1LocalVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type to mount. It applies only when the Path is a block device. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default value is to auto-select a filesystem if unspecified.
	Path string `json:"path"` // path of the full path to the volume on the node. It can be either a directory or block device (disk, partition, ...).
}

// Iok8sapicorev1Volume represents the Iok8sapicorev1Volume schema from the OpenAPI specification
type Iok8sapicorev1Volume struct {
	Image Iok8sapicorev1ImageVolumeSource `json:"image,omitempty"` // ImageVolumeSource represents a image volume resource.
	Csi Iok8sapicorev1CSIVolumeSource `json:"csi,omitempty"` // Represents a source location of a volume to mount, managed by an external CSI driver
	Azuredisk Iok8sapicorev1AzureDiskVolumeSource `json:"azureDisk,omitempty"` // AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	Iscsi Iok8sapicorev1ISCSIVolumeSource `json:"iscsi,omitempty"` // Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
	Cephfs Iok8sapicorev1CephFSVolumeSource `json:"cephfs,omitempty"` // Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
	Persistentvolumeclaim Iok8sapicorev1PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"` // PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).
	Awselasticblockstore Iok8sapicorev1AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"` // Represents a Persistent Disk resource in AWS. An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.
	Hostpath Iok8sapicorev1HostPathVolumeSource `json:"hostPath,omitempty"` // Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.
	Storageos Iok8sapicorev1StorageOSVolumeSource `json:"storageos,omitempty"` // Represents a StorageOS persistent volume resource.
	Flocker Iok8sapicorev1FlockerVolumeSource `json:"flocker,omitempty"` // Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
	Gitrepo Iok8sapicorev1GitRepoVolumeSource `json:"gitRepo,omitempty"` // Represents a volume that is populated with the contents of a git repository. Git repo volumes do not support ownership management. Git repo volumes support SELinux relabeling. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.
	Configmap Iok8sapicorev1ConfigMapVolumeSource `json:"configMap,omitempty"` // Adapts a ConfigMap into a volume. The contents of the target ConfigMap's Data field will be presented in a volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. ConfigMap volumes support ownership management and SELinux relabeling.
	Emptydir Iok8sapicorev1EmptyDirVolumeSource `json:"emptyDir,omitempty"` // Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.
	Projected Iok8sapicorev1ProjectedVolumeSource `json:"projected,omitempty"` // Represents a projected volume source
	Downwardapi Iok8sapicorev1DownwardAPIVolumeSource `json:"downwardAPI,omitempty"` // DownwardAPIVolumeSource represents a volume containing downward API info. Downward API volumes support ownership management and SELinux relabeling.
	Quobyte Iok8sapicorev1QuobyteVolumeSource `json:"quobyte,omitempty"` // Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
	Flexvolume Iok8sapicorev1FlexVolumeSource `json:"flexVolume,omitempty"` // FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
	Gcepersistentdisk Iok8sapicorev1GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty"` // Represents a Persistent Disk resource in Google Compute Engine. A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.
	Vspherevolume Iok8sapicorev1VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty"` // Represents a vSphere volume resource.
	Name string `json:"name"` // name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Nfs Iok8sapicorev1NFSVolumeSource `json:"nfs,omitempty"` // Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.
	Secret Iok8sapicorev1SecretVolumeSource `json:"secret,omitempty"` // Adapts a Secret into a volume. The contents of the target Secret's Data field will be presented in a volume as files using the keys in the Data field as the file names. Secret volumes support ownership management and SELinux relabeling.
	Portworxvolume Iok8sapicorev1PortworxVolumeSource `json:"portworxVolume,omitempty"` // PortworxVolumeSource represents a Portworx volume resource.
	Rbd Iok8sapicorev1RBDVolumeSource `json:"rbd,omitempty"` // Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
	Glusterfs Iok8sapicorev1GlusterfsVolumeSource `json:"glusterfs,omitempty"` // Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
	Ephemeral Iok8sapicorev1EphemeralVolumeSource `json:"ephemeral,omitempty"` // Represents an ephemeral volume that is handled by a normal storage driver.
	Azurefile Iok8sapicorev1AzureFileVolumeSource `json:"azureFile,omitempty"` // AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	Cinder Iok8sapicorev1CinderVolumeSource `json:"cinder,omitempty"` // Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
	Fc Iok8sapicorev1FCVolumeSource `json:"fc,omitempty"` // Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
	Scaleio Iok8sapicorev1ScaleIOVolumeSource `json:"scaleIO,omitempty"` // ScaleIOVolumeSource represents a persistent ScaleIO volume
	Photonpersistentdisk Iok8sapicorev1PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty"` // Represents a Photon Controller persistent disk resource.
}

// Iok8sapicorev1FCVolumeSource represents the Iok8sapicorev1FCVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1FCVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Lun int `json:"lun,omitempty"` // lun is Optional: FC target lun number
	Readonly bool `json:"readOnly,omitempty"` // readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Targetwwns []string `json:"targetWWNs,omitempty"` // targetWWNs is Optional: FC target worldwide names (WWNs)
	Wwids []string `json:"wwids,omitempty"` // wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
}

// Iok8sapicorev1DownwardAPIProjection represents the Iok8sapicorev1DownwardAPIProjection schema from the OpenAPI specification
type Iok8sapicorev1DownwardAPIProjection struct {
	Items []Iok8sapicorev1DownwardAPIVolumeFile `json:"items,omitempty"` // Items is a list of DownwardAPIVolume file
}

// Iok8sapicorev1PersistentVolumeStatus represents the Iok8sapicorev1PersistentVolumeStatus schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeStatus struct {
	Phase string `json:"phase,omitempty"` // phase indicates if a volume is available, bound to a claim, or released by a claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#phase
	Reason string `json:"reason,omitempty"` // reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI.
	Lastphasetransitiontime string `json:"lastPhaseTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // message is a human-readable message indicating details about why the volume is in this state.
}

// Iok8sapicorev1PersistentVolumeClaimTemplate represents the Iok8sapicorev1PersistentVolumeClaimTemplate schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaimTemplate struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1PersistentVolumeClaimSpec `json:"spec"` // PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes
}

// Iok8sapicorev1EnvVar represents the Iok8sapicorev1EnvVar schema from the OpenAPI specification
type Iok8sapicorev1EnvVar struct {
	Name string `json:"name"` // Name of the environment variable. May consist of any printable ASCII characters except '='.
	Value string `json:"value,omitempty"` // Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
	Valuefrom Iok8sapicorev1EnvVarSource `json:"valueFrom,omitempty"` // EnvVarSource represents a source for the value of an EnvVar.
}

// Iok8sapicorev1PodOS represents the Iok8sapicorev1PodOS schema from the OpenAPI specification
type Iok8sapicorev1PodOS struct {
	Name string `json:"name"` // Name is the name of the operating system. The currently supported values are linux and windows. Additional value may be defined in future and can be one of: https://github.com/opencontainers/runtime-spec/blob/master/config.md#platform-specific-configuration Clients should expect to handle additional values and treat unrecognized values in this field as os: null
}

// Iok8sapicorev1PodAffinityTerm represents the Iok8sapicorev1PodAffinityTerm schema from the OpenAPI specification
type Iok8sapicorev1PodAffinityTerm struct {
	Matchlabelkeys []string `json:"matchLabelKeys,omitempty"` // MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)` to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set.
	Mismatchlabelkeys []string `json:"mismatchLabelKeys,omitempty"` // MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)` to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set.
	Namespaceselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"namespaceSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Namespaces []string `json:"namespaces,omitempty"` // namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means "this pod's namespace".
	Topologykey string `json:"topologyKey"` // This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.
	Labelselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"labelSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapistoragev1CSIStorageCapacity represents the Iok8sapistoragev1CSIStorageCapacity schema from the OpenAPI specification
type Iok8sapistoragev1CSIStorageCapacity struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Maximumvolumesize string `json:"maximumVolumeSize,omitempty"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Nodetopology Iok8sapimachinerypkgapismetav1LabelSelector `json:"nodeTopology,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Storageclassname string `json:"storageClassName"` // storageClassName represents the name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Capacity string `json:"capacity,omitempty"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
}

// Iok8sapicorev1CSIPersistentVolumeSource represents the Iok8sapicorev1CSIPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1CSIPersistentVolumeSource struct {
	Volumeattributes map[string]interface{} `json:"volumeAttributes,omitempty"` // volumeAttributes of the volume to publish.
	Nodepublishsecretref Iok8sapicorev1SecretReference `json:"nodePublishSecretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	Driver string `json:"driver"` // driver is the name of the driver to use for this volume. Required.
	Controllerexpandsecretref Iok8sapicorev1SecretReference `json:"controllerExpandSecretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	Controllerpublishsecretref Iok8sapicorev1SecretReference `json:"controllerPublishSecretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	Fstype string `json:"fsType,omitempty"` // fsType to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
	Nodeexpandsecretref Iok8sapicorev1SecretReference `json:"nodeExpandSecretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	Readonly bool `json:"readOnly,omitempty"` // readOnly value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).
	Volumehandle string `json:"volumeHandle"` // volumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.
	Nodestagesecretref Iok8sapicorev1SecretReference `json:"nodeStageSecretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
}

// Iok8sapiresourcev1beta1ResourceClaimSpec represents the Iok8sapiresourcev1beta1ResourceClaimSpec schema from the OpenAPI specification
type Iok8sapiresourcev1beta1ResourceClaimSpec struct {
	Devices Iok8sapiresourcev1beta1DeviceClaim `json:"devices,omitempty"` // DeviceClaim defines how to request devices with a ResourceClaim.
}

// Iok8sapicorev1LimitRangeItem represents the Iok8sapicorev1LimitRangeItem schema from the OpenAPI specification
type Iok8sapicorev1LimitRangeItem struct {
	DefaultField map[string]interface{} `json:"default,omitempty"` // Default resource requirement limit value by resource name if resource limit is omitted.
	Defaultrequest map[string]interface{} `json:"defaultRequest,omitempty"` // DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.
	Max map[string]interface{} `json:"max,omitempty"` // Max usage constraints on this kind by resource name.
	Maxlimitrequestratio map[string]interface{} `json:"maxLimitRequestRatio,omitempty"` // MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.
	Min map[string]interface{} `json:"min,omitempty"` // Min usage constraints on this kind by resource name.
	TypeField string `json:"type"` // Type of resource that this limit applies to.
}

// Iok8sapiresourcev1beta1CELDeviceSelector represents the Iok8sapiresourcev1beta1CELDeviceSelector schema from the OpenAPI specification
type Iok8sapiresourcev1beta1CELDeviceSelector struct {
	Expression string `json:"expression"` // Expression is a CEL expression which evaluates a single device. It must evaluate to true when the device under consideration satisfies the desired criteria, and false when it does not. Any other result is an error and causes allocation of devices to abort. The expression's input is an object named "device", which carries the following properties: - driver (string): the name of the driver which defines this device. - attributes (map[string]object): the device's attributes, grouped by prefix (e.g. device.attributes["dra.example.com"] evaluates to an object with all of the attributes which were prefixed by "dra.example.com". - capacity (map[string]object): the device's capacities, grouped by prefix. Example: Consider a device with driver="dra.example.com", which exposes two attributes named "model" and "ext.example.com/family" and which exposes one capacity named "modules". This input to this expression would have the following fields: device.driver device.attributes["dra.example.com"].model device.attributes["ext.example.com"].family device.capacity["dra.example.com"].modules The device.driver field can be used to check for a specific driver, either as a high-level precondition (i.e. you only want to consider devices from this driver) or as part of a multi-clause expression that is meant to consider devices from different drivers. The value type of each attribute is defined by the device definition, and users who write these expressions must consult the documentation for their specific drivers. The value type of each capacity is Quantity. If an unknown prefix is used as a lookup in either device.attributes or device.capacity, an empty map will be returned. Any reference to an unknown field will cause an evaluation error and allocation to abort. A robust expression should check for the existence of attributes before referencing them. For ease of use, the cel.bind() function is enabled, and can be used to simplify expressions that access multiple attributes with the same domain. For example: cel.bind(dra, device.attributes["dra.example.com"], dra.someBool && dra.anotherBool) The length of the expression must be smaller or equal to 10 Ki. The cost of evaluating it is also limited based on the estimated number of logical steps.
}

// Iok8sapiauthorizationv1SubjectAccessReview represents the Iok8sapiauthorizationv1SubjectAccessReview schema from the OpenAPI specification
type Iok8sapiauthorizationv1SubjectAccessReview struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1SubjectAccessReviewSpec `json:"spec"` // SubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
	Status Iok8sapiauthorizationv1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapiapiserverinternalv1alpha1StorageVersionCondition represents the Iok8sapiapiserverinternalv1alpha1StorageVersionCondition schema from the OpenAPI specification
type Iok8sapiapiserverinternalv1alpha1StorageVersionCondition struct {
	Reason string `json:"reason"` // The reason for the condition's last transition.
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of the condition.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message"` // A human readable message indicating details about the transition.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // If set, this represents the .metadata.generation that the condition was set based upon.
}

// Iok8sapicorev1TopologySelectorLabelRequirement represents the Iok8sapicorev1TopologySelectorLabelRequirement schema from the OpenAPI specification
type Iok8sapicorev1TopologySelectorLabelRequirement struct {
	Key string `json:"key"` // The label key that the selector applies to.
	Values []string `json:"values"` // An array of string values. One value must match the label to be selected. Each entry in Values is ORed.
}

// Iok8sapinodev1Overhead represents the Iok8sapinodev1Overhead schema from the OpenAPI specification
type Iok8sapinodev1Overhead struct {
	Podfixed map[string]interface{} `json:"podFixed,omitempty"` // podFixed represents the fixed resource overhead associated with running a pod.
}

// Iok8sapiautoscalingv2ObjectMetricStatus represents the Iok8sapiautoscalingv2ObjectMetricStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2ObjectMetricStatus struct {
	Describedobject Iok8sapiautoscalingv2CrossVersionObjectReference `json:"describedObject"` // CrossVersionObjectReference contains enough information to let you identify the referred resource.
	Metric Iok8sapiautoscalingv2MetricIdentifier `json:"metric"` // MetricIdentifier defines the name and optionally selector for a metric
	Current Iok8sapiautoscalingv2MetricValueStatus `json:"current"` // MetricValueStatus holds the current value for a metric
}

// Iok8sapinetworkingv1IPAddressSpec represents the Iok8sapinetworkingv1IPAddressSpec schema from the OpenAPI specification
type Iok8sapinetworkingv1IPAddressSpec struct {
	Parentref Iok8sapinetworkingv1ParentReference `json:"parentRef"` // ParentReference describes a reference to a parent object.
}

// Iok8sapiresourcev1beta1DeviceRequestAllocationResult represents the Iok8sapiresourcev1beta1DeviceRequestAllocationResult schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceRequestAllocationResult struct {
	Tolerations []Iok8sapiresourcev1beta1DeviceToleration `json:"tolerations,omitempty"` // A copy of all tolerations specified in the request at the time when the device got allocated. The maximum number of tolerations is 16. This is an alpha field and requires enabling the DRADeviceTaints feature gate.
	Adminaccess bool `json:"adminAccess,omitempty"` // AdminAccess indicates that this device was allocated for administrative access. See the corresponding request field for a definition of mode. This is an alpha field and requires enabling the DRAAdminAccess feature gate. Admin access is disabled if this field is unset or set to false, otherwise it is enabled.
	Device string `json:"device"` // Device references one device instance via its name in the driver's resource pool. It must be a DNS label.
	Driver string `json:"driver"` // Driver specifies the name of the DRA driver whose kubelet plugin should be invoked to process the allocation once the claim is needed on a node. Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver.
	Pool string `json:"pool"` // This name together with the driver name and the device name field identify which device was allocated (`<driver name>/<pool name>/<device name>`). Must not be longer than 253 characters and may contain one or more DNS sub-domains separated by slashes.
	Request string `json:"request"` // Request is the name of the request in the claim which caused this device to be allocated. If it references a subrequest in the firstAvailable list on a DeviceRequest, this field must include both the name of the main request and the subrequest using the format <main request>/<subrequest>. Multiple devices may have been allocated per request.
}

// Iok8sapistoragev1StorageClass represents the Iok8sapistoragev1StorageClass schema from the OpenAPI specification
type Iok8sapistoragev1StorageClass struct {
	Mountoptions []string `json:"mountOptions,omitempty"` // mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	Volumebindingmode string `json:"volumeBindingMode,omitempty"` // volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound. When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	Allowvolumeexpansion bool `json:"allowVolumeExpansion,omitempty"` // allowVolumeExpansion shows whether the storage class allow volume expand.
	Allowedtopologies []Iok8sapicorev1TopologySelectorTerm `json:"allowedTopologies,omitempty"` // allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Parameters map[string]interface{} `json:"parameters,omitempty"` // parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Provisioner string `json:"provisioner"` // provisioner indicates the type of the provisioner.
	Reclaimpolicy string `json:"reclaimPolicy,omitempty"` // reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
}

// Iok8sapicertificatesv1alpha1ClusterTrustBundle represents the Iok8sapicertificatesv1alpha1ClusterTrustBundle schema from the OpenAPI specification
type Iok8sapicertificatesv1alpha1ClusterTrustBundle struct {
	Spec Iok8sapicertificatesv1alpha1ClusterTrustBundleSpec `json:"spec"` // ClusterTrustBundleSpec contains the signer and trust anchors.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiautoscalingv2MetricValueStatus represents the Iok8sapiautoscalingv2MetricValueStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2MetricValueStatus struct {
	Averagevalue string `json:"averageValue,omitempty"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
	Value string `json:"value,omitempty"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
	Averageutilization int `json:"averageUtilization,omitempty"` // currentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
}

// Iok8sapicorev1AzureFileVolumeSource represents the Iok8sapicorev1AzureFileVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1AzureFileVolumeSource struct {
	Readonly bool `json:"readOnly,omitempty"` // readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretname string `json:"secretName"` // secretName is the name of secret that contains Azure Storage Account Name and Key
	Sharename string `json:"shareName"` // shareName is the azure share Name
}

// Iok8sapicorev1ImageVolumeSource represents the Iok8sapicorev1ImageVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ImageVolumeSource struct {
	Pullpolicy string `json:"pullPolicy,omitempty"` // Policy for pulling OCI objects. Possible values are: Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails. Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present. IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.
	Reference string `json:"reference,omitempty"` // Required: Image or artifact reference to be used. Behaves in the same way as pod.spec.containers[*].image. Pull secrets will be assembled in the same way as for the container image by looking up node credentials, SA image pull secrets, and pod spec image pull secrets. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.
}

// Iok8sapiappsv1ControllerRevision represents the Iok8sapiappsv1ControllerRevision schema from the OpenAPI specification
type Iok8sapiappsv1ControllerRevision struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Data Iok8sapimachinerypkgruntimeRawExtension `json:"data,omitempty"` // RawExtension is used to hold extensions in external versions. To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types. // Internal package: 	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.Object `json:"myPlugin"` 	} 	type PluginA struct { 		AOption string `json:"aOption"` 	} // External package: 	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.RawExtension `json:"myPlugin"` 	} 	type PluginA struct { 		AOption string `json:"aOption"` 	} // On the wire, the JSON will look something like this: 	{ 		"kind":"MyAPIObject", 		"apiVersion":"v1", 		"myPlugin": { 			"kind":"PluginA", 			"aOption":"foo", 		}, 	} So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Revision int64 `json:"revision"` // Revision indicates the revision of the state represented by Data.
}

// Iok8sapicorev1ScopeSelector represents the Iok8sapicorev1ScopeSelector schema from the OpenAPI specification
type Iok8sapicorev1ScopeSelector struct {
	Matchexpressions []Iok8sapicorev1ScopedResourceSelectorRequirement `json:"matchExpressions,omitempty"` // A list of scope selector requirements by scope of the resources.
}

// Iok8sapiapiserverinternalv1alpha1StorageVersionSpec represents the Iok8sapiapiserverinternalv1alpha1StorageVersionSpec schema from the OpenAPI specification
type Iok8sapiapiserverinternalv1alpha1StorageVersionSpec struct {
}

// Iok8sapicorev1PodStatus represents the Iok8sapicorev1PodStatus schema from the OpenAPI specification
type Iok8sapicorev1PodStatus struct {
	Containerstatuses []Iok8sapicorev1ContainerStatus `json:"containerStatuses,omitempty"` // Statuses of containers in this pod. Each container in the pod should have at most one status in this list, and all statuses should be for containers in the pod. However this is not enforced. If a status for a non-existent container is present in the list, or the list has duplicate names, the behavior of various Kubernetes components is not defined and those statuses might be ignored. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
	Starttime string `json:"startTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Hostip string `json:"hostIP,omitempty"` // hostIP holds the IP address of the host to which the pod is assigned. Empty if the pod has not started yet. A pod can be assigned to a node that has a problem in kubelet which in turns mean that HostIP will not be updated even if there is a node is assigned to pod
	Resize string `json:"resize,omitempty"` // Status of resources resize desired for pod's containers. It is empty if no resources resize is pending. Any changes to container resources will automatically set this to "Proposed" Deprecated: Resize status is moved to two pod conditions PodResizePending and PodResizeInProgress. PodResizePending will track states where the spec has been resized, but the Kubelet has not yet allocated the resources. PodResizeInProgress will track in-progress resizes, and should be present whenever allocated resources != acknowledged resources.
	Phase string `json:"phase,omitempty"` // The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle. The conditions array, the reason and message fields, and the individual container status arrays contain more detail about the pod's status. There are five possible phase values: Pending: The pod has been accepted by the Kubernetes system, but one or more of the container images has not been created. This includes time before being scheduled as well as time spent downloading images over the network, which could take a while. Running: The pod has been bound to a node, and all of the containers have been created. At least one container is still running, or is in the process of starting or restarting. Succeeded: All containers in the pod have terminated in success, and will not be restarted. Failed: All containers in the pod have terminated, and at least one container has terminated in failure. The container either exited with non-zero status or was terminated by the system. Unknown: For some reason the state of the pod could not be obtained, typically due to an error in communicating with the host of the pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase
	Reason string `json:"reason,omitempty"` // A brief CamelCase message indicating details about why the pod is in this state. e.g. 'Evicted'
	Hostips []Iok8sapicorev1HostIP `json:"hostIPs,omitempty"` // hostIPs holds the IP addresses allocated to the host. If this field is specified, the first entry must match the hostIP field. This list is empty if the pod has not started yet. A pod can be assigned to a node that has a problem in kubelet which in turns means that HostIPs will not be updated even if there is a node is assigned to this pod.
	Ephemeralcontainerstatuses []Iok8sapicorev1ContainerStatus `json:"ephemeralContainerStatuses,omitempty"` // Statuses for any ephemeral containers that have run in this pod. Each ephemeral container in the pod should have at most one status in this list, and all statuses should be for containers in the pod. However this is not enforced. If a status for a non-existent container is present in the list, or the list has duplicate names, the behavior of various Kubernetes components is not defined and those statuses might be ignored. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
	Initcontainerstatuses []Iok8sapicorev1ContainerStatus `json:"initContainerStatuses,omitempty"` // Statuses of init containers in this pod. The most recent successful non-restartable init container will have ready = true, the most recently started container will have startTime set. Each init container in the pod should have at most one status in this list, and all statuses should be for containers in the pod. However this is not enforced. If a status for a non-existent container is present in the list, or the list has duplicate names, the behavior of various Kubernetes components is not defined and those statuses might be ignored. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-and-container-status
	Message string `json:"message,omitempty"` // A human readable message indicating details about why the pod is in this condition.
	Podip string `json:"podIP,omitempty"` // podIP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated.
	Podips []Iok8sapicorev1PodIP `json:"podIPs,omitempty"` // podIPs holds the IP addresses allocated to the pod. If this field is specified, the 0th entry must match the podIP field. Pods may be allocated at most 1 value for each of IPv4 and IPv6. This list is empty if no IPs have been allocated yet.
	Resourceclaimstatuses []Iok8sapicorev1PodResourceClaimStatus `json:"resourceClaimStatuses,omitempty"` // Status of resource claims.
	Conditions []Iok8sapicorev1PodCondition `json:"conditions,omitempty"` // Current service state of pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Nominatednodename string `json:"nominatedNodeName,omitempty"` // nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods. This field does not guarantee that the pod will be scheduled on this node. Scheduler may decide to place the pod elsewhere if other nodes become available sooner. Scheduler may also decide to give the resources on this node to a higher priority pod that is created after preemption. As a result, this field may be different than PodSpec.nodeName when the pod is scheduled.
	Observedgeneration int64 `json:"observedGeneration,omitempty"` // If set, this represents the .metadata.generation that the pod status was set based upon. This is an alpha field. Enable PodObservedGenerationTracking to be able to use this field.
	Qosclass string `json:"qosClass,omitempty"` // The Quality of Service (QOS) classification assigned to the pod based on resource requirements See PodQOSClass type for available QOS classes More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-qos/#quality-of-service-classes
}

// Iok8sapiresourcev1beta1ResourceClaimStatus represents the Iok8sapiresourcev1beta1ResourceClaimStatus schema from the OpenAPI specification
type Iok8sapiresourcev1beta1ResourceClaimStatus struct {
	Reservedfor []Iok8sapiresourcev1beta1ResourceClaimConsumerReference `json:"reservedFor,omitempty"` // ReservedFor indicates which entities are currently allowed to use the claim. A Pod which references a ResourceClaim which is not reserved for that Pod will not be started. A claim that is in use or might be in use because it has been reserved must not get deallocated. In a cluster with multiple scheduler instances, two pods might get scheduled concurrently by different schedulers. When they reference the same ResourceClaim which already has reached its maximum number of consumers, only one pod can be scheduled. Both schedulers try to add their pod to the claim.status.reservedFor field, but only the update that reaches the API server first gets stored. The other one fails with an error and the scheduler which issued it knows that it must put the pod back into the queue, waiting for the ResourceClaim to become usable again. There can be at most 256 such reservations. This may get increased in the future, but not reduced.
	Allocation Iok8sapiresourcev1beta1AllocationResult `json:"allocation,omitempty"` // AllocationResult contains attributes of an allocated resource.
	Devices []Iok8sapiresourcev1beta1AllocatedDeviceStatus `json:"devices,omitempty"` // Devices contains the status of each device allocated for this claim, as reported by the driver. This can include driver-specific information. Entries are owned by their respective drivers.
}

// Iok8sapiautoscalingv2HorizontalPodAutoscalerSpec represents the Iok8sapiautoscalingv2HorizontalPodAutoscalerSpec schema from the OpenAPI specification
type Iok8sapiautoscalingv2HorizontalPodAutoscalerSpec struct {
	Behavior Iok8sapiautoscalingv2HorizontalPodAutoscalerBehavior `json:"behavior,omitempty"` // HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).
	Maxreplicas int `json:"maxReplicas"` // maxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas.
	Metrics []Iok8sapiautoscalingv2MetricSpec `json:"metrics,omitempty"` // metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used). The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods. Ergo, metrics used must decrease as the pod count is increased, and vice-versa. See the individual metric source types for more information about how each type of metric must respond. If not set, the default metric will be set to 80% average CPU utilization.
	Minreplicas int `json:"minReplicas,omitempty"` // minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod. minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured. Scaling is active as long as at least one metric value is available.
	Scaletargetref Iok8sapiautoscalingv2CrossVersionObjectReference `json:"scaleTargetRef"` // CrossVersionObjectReference contains enough information to let you identify the referred resource.
}

// Iok8sapicorev1GlusterfsVolumeSource represents the Iok8sapicorev1GlusterfsVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1GlusterfsVolumeSource struct {
	Readonly bool `json:"readOnly,omitempty"` // readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Endpoints string `json:"endpoints"` // endpoints is the endpoint name that details Glusterfs topology.
	Path string `json:"path"` // path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
}

// Iok8sapiresourcev1beta1DeviceClaimConfiguration represents the Iok8sapiresourcev1beta1DeviceClaimConfiguration schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceClaimConfiguration struct {
	Opaque Iok8sapiresourcev1beta1OpaqueDeviceConfiguration `json:"opaque,omitempty"` // OpaqueDeviceConfiguration contains configuration parameters for a driver in a format defined by the driver vendor.
	Requests []string `json:"requests,omitempty"` // Requests lists the names of requests where the configuration applies. If empty, it applies to all requests. References to subrequests must include the name of the main request and may include the subrequest using the format <main request>[/<subrequest>]. If just the main request is given, the configuration applies to all subrequests.
}

// Iok8sapiresourcev1beta1DeviceCapacity represents the Iok8sapiresourcev1beta1DeviceCapacity schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceCapacity struct {
	Value string `json:"value"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
}

// Iok8sapiauthenticationv1UserInfo represents the Iok8sapiauthenticationv1UserInfo schema from the OpenAPI specification
type Iok8sapiauthenticationv1UserInfo struct {
	Groups []string `json:"groups,omitempty"` // The names of groups this user is a part of.
	Uid string `json:"uid,omitempty"` // A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
	Username string `json:"username,omitempty"` // The name that uniquely identifies this user among all active users.
	Extra map[string]interface{} `json:"extra,omitempty"` // Any additional information provided by the authenticator.
}

// Iok8sapiappsv1StatefulSetSpec represents the Iok8sapiappsv1StatefulSetSpec schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSetSpec struct {
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Servicename string `json:"serviceName,omitempty"` // serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where "pod-specific-string" is managed by the StatefulSet controller.
	Volumeclaimtemplates []Iok8sapicorev1PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty"` // volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	Podmanagementpolicy string `json:"podManagementPolicy,omitempty"` // podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is `OrderedReady`, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is `Parallel` which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.
	Updatestrategy Iok8sapiappsv1StatefulSetUpdateStrategy `json:"updateStrategy,omitempty"` // StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.
	Ordinals Iok8sapiappsv1StatefulSetOrdinals `json:"ordinals,omitempty"` // StatefulSetOrdinals describes the policy used for replica ordinal assignment in this StatefulSet.
	Persistentvolumeclaimretentionpolicy Iok8sapiappsv1StatefulSetPersistentVolumeClaimRetentionPolicy `json:"persistentVolumeClaimRetentionPolicy,omitempty"` // StatefulSetPersistentVolumeClaimRetentionPolicy describes the policy used for PVCs created from the StatefulSet VolumeClaimTemplates.
	Replicas int `json:"replicas,omitempty"` // replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Insecureskiptlsverify bool `json:"insecureSkipTLSVerify,omitempty"` // InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged. You should use the CABundle instead.
	Service GeneratedType `json:"service,omitempty"` // ServiceReference holds a reference to Service.legacy.k8s.io
	Version string `json:"version,omitempty"` // Version is the API version this server hosts. For example, "v1"
	Versionpriority int `json:"versionPriority"` // VersionPriority controls the ordering of this API version inside of its group. Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). Since it's inside of a group, the number can be small, probably in the 10s. In case of equal version priorities, the version string will be used to compute the order inside a group. If the version string is "kube-like", it will sort above non "kube-like" version strings, which are ordered lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the major version), then optionally the string "alpha" or "beta" and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.
	Cabundle string `json:"caBundle,omitempty"` // CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate. If unspecified, system trust roots on the apiserver are used.
	Group string `json:"group,omitempty"` // Group is the API group name this server hosts
	Grouppriorityminimum int `json:"groupPriorityMinimum"` // GroupPriorityMinimum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMinimum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object. (v1.bar before v1.foo) We'd recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s
}

// Iok8sapicorev1QuobyteVolumeSource represents the Iok8sapicorev1QuobyteVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1QuobyteVolumeSource struct {
	Group string `json:"group,omitempty"` // group to map volume access to Default is no group
	Readonly bool `json:"readOnly,omitempty"` // readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
	Registry string `json:"registry"` // registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes
	Tenant string `json:"tenant,omitempty"` // tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin
	User string `json:"user,omitempty"` // user to map volume access to Defaults to serivceaccount user
	Volume string `json:"volume"` // volume is a string that references an already created Quobyte volume by name.
}

// Iok8sapicorev1TypedObjectReference represents the Iok8sapicorev1TypedObjectReference schema from the OpenAPI specification
type Iok8sapicorev1TypedObjectReference struct {
	Apigroup string `json:"apiGroup,omitempty"` // APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
	Kind string `json:"kind"` // Kind is the type of resource being referenced
	Name string `json:"name"` // Name is the name of resource being referenced
	Namespace string `json:"namespace,omitempty"` // Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.
}

// Iok8sapidiscoveryv1EndpointHints represents the Iok8sapidiscoveryv1EndpointHints schema from the OpenAPI specification
type Iok8sapidiscoveryv1EndpointHints struct {
	Fornodes []Iok8sapidiscoveryv1ForNode `json:"forNodes,omitempty"` // forNodes indicates the node(s) this endpoint should be consumed by when using topology aware routing. May contain a maximum of 8 entries. This is an Alpha feature and is only used when the PreferSameTrafficDistribution feature gate is enabled.
	Forzones []Iok8sapidiscoveryv1ForZone `json:"forZones,omitempty"` // forZones indicates the zone(s) this endpoint should be consumed by when using topology aware routing. May contain a maximum of 8 entries.
}

// Iok8sapiresourcev1alpha3DeviceTaintRule represents the Iok8sapiresourcev1alpha3DeviceTaintRule schema from the OpenAPI specification
type Iok8sapiresourcev1alpha3DeviceTaintRule struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiresourcev1alpha3DeviceTaintRuleSpec `json:"spec"` // DeviceTaintRuleSpec specifies the selector and one taint.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1VolumeMountStatus represents the Iok8sapicorev1VolumeMountStatus schema from the OpenAPI specification
type Iok8sapicorev1VolumeMountStatus struct {
	Mountpath string `json:"mountPath"` // MountPath corresponds to the original VolumeMount.
	Name string `json:"name"` // Name corresponds to the name of the original VolumeMount.
	Readonly bool `json:"readOnly,omitempty"` // ReadOnly corresponds to the original VolumeMount.
	Recursivereadonly string `json:"recursiveReadOnly,omitempty"` // RecursiveReadOnly must be set to Disabled, Enabled, or unspecified (for non-readonly mounts). An IfPossible value in the original VolumeMount must be translated to Disabled or Enabled, depending on the mount result.
}

// Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyBindingList represents the Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyBindingList schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyBindingList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyBinding `json:"items"` // List of PolicyBinding.
}

// Iok8sapiresourcev1beta1DeviceConstraint represents the Iok8sapiresourcev1beta1DeviceConstraint schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceConstraint struct {
	Matchattribute string `json:"matchAttribute,omitempty"` // MatchAttribute requires that all devices in question have this attribute and that its type and value are the same across those devices. For example, if you specified "dra.example.com/numa" (a hypothetical example!), then only devices in the same NUMA node will be chosen. A device which does not have that attribute will not be chosen. All devices should use a value of the same type for this attribute because that is part of its specification, but if one device doesn't, then it also will not be chosen. Must include the domain qualifier.
	Requests []string `json:"requests,omitempty"` // Requests is a list of the one or more requests in this claim which must co-satisfy this constraint. If a request is fulfilled by multiple devices, then all of the devices must satisfy the constraint. If this is not specified, this constraint applies to all requests in this claim. References to subrequests must include the name of the main request and may include the subrequest using the format <main request>[/<subrequest>]. If just the main request is given, the constraint applies to all subrequests.
}

// Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyBindingSpec represents the Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyBindingSpec schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyBindingSpec struct {
	Validationactions []string `json:"validationActions,omitempty"` // validationActions declares how Validations of the referenced ValidatingAdmissionPolicy are enforced. If a validation evaluates to false it is always enforced according to these actions. Failures defined by the ValidatingAdmissionPolicy's FailurePolicy are enforced according to these actions only if the FailurePolicy is set to Fail, otherwise the failures are ignored. This includes compilation errors, runtime errors and misconfigurations of the policy. validationActions is declared as a set of action values. Order does not matter. validationActions may not contain duplicates of the same action. The supported actions values are: "Deny" specifies that a validation failure results in a denied request. "Warn" specifies that a validation failure is reported to the request client in HTTP Warning headers, with a warning code of 299. Warnings can be sent both for allowed or denied admission responses. "Audit" specifies that a validation failure is included in the published audit event for the request. The audit event will contain a `validation.policy.admission.k8s.io/validation_failure` audit annotation with a value containing the details of the validation failures, formatted as a JSON list of objects, each with the following fields: - message: The validation failure message string - policy: The resource name of the ValidatingAdmissionPolicy - binding: The resource name of the ValidatingAdmissionPolicyBinding - expressionIndex: The index of the failed validations in the ValidatingAdmissionPolicy - validationActions: The enforcement actions enacted for the validation failure Example audit annotation: `"validation.policy.admission.k8s.io/validation_failure": "[{\"message\": \"Invalid value\", {\"policy\": \"policy.example.com\", {\"binding\": \"policybinding.example.com\", {\"expressionIndex\": \"1\", {\"validationActions\": [\"Audit\"]}]"` Clients should expect to handle additional values by ignoring any values not recognized. "Deny" and "Warn" may not be used together since this combination needlessly duplicates the validation failure both in the API response body and the HTTP warning headers. Required.
	Matchresources Iok8sapiadmissionregistrationv1MatchResources `json:"matchResources,omitempty"` // MatchResources decides whether to run the admission control policy on an object based on whether it meets the match criteria. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)
	Paramref Iok8sapiadmissionregistrationv1ParamRef `json:"paramRef,omitempty"` // ParamRef describes how to locate the params to be used as input to expressions of rules applied by a policy binding.
	Policyname string `json:"policyName,omitempty"` // PolicyName references a ValidatingAdmissionPolicy name which the ValidatingAdmissionPolicyBinding binds to. If the referenced resource does not exist, this binding is considered invalid and will be ignored Required.
}

// Iok8sapiauthorizationv1LabelSelectorAttributes represents the Iok8sapiauthorizationv1LabelSelectorAttributes schema from the OpenAPI specification
type Iok8sapiauthorizationv1LabelSelectorAttributes struct {
	Rawselector string `json:"rawSelector,omitempty"` // rawSelector is the serialization of a field selector that would be included in a query parameter. Webhook implementations are encouraged to ignore rawSelector. The kube-apiserver's *SubjectAccessReview will parse the rawSelector as long as the requirements are not present.
	Requirements []Iok8sapimachinerypkgapismetav1LabelSelectorRequirement `json:"requirements,omitempty"` // requirements is the parsed interpretation of a label selector. All requirements must be met for a resource instance to match the selector. Webhook implementations should handle requirements, but how to handle them is up to the webhook. Since requirements can only limit the request, it is safe to authorize as unlimited request if the requirements are not understood.
}

// Iok8sapicorev1StorageOSVolumeSource represents the Iok8sapicorev1StorageOSVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1StorageOSVolumeSource struct {
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	Volumename string `json:"volumeName,omitempty"` // volumeName is the human-readable name of the StorageOS volume. Volume names are only unique within a namespace.
	Volumenamespace string `json:"volumeNamespace,omitempty"` // volumeNamespace specifies the scope of the volume within StorageOS. If no namespace is specified then the Pod's namespace will be used. This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Readonly bool `json:"readOnly,omitempty"` // readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
}

// Iok8sapiadmissionregistrationv1MutatingWebhook represents the Iok8sapiadmissionregistrationv1MutatingWebhook schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1MutatingWebhook struct {
	Sideeffects string `json:"sideEffects"` // SideEffects states whether this webhook has side effects. Acceptable values are: None, NoneOnDryRun (webhooks created via v1beta1 may also specify Some or Unknown). Webhooks with side effects MUST implement a reconciliation system, since a request may be rejected by a future step in the admission chain and the side effects therefore need to be undone. Requests with the dryRun attribute will be auto-rejected if they match a webhook with sideEffects == Unknown or Some.
	Timeoutseconds int `json:"timeoutSeconds,omitempty"` // TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook call will be ignored or the API call will fail based on the failure policy. The timeout value must be between 1 and 30 seconds. Default to 10 seconds.
	Matchconditions []Iok8sapiadmissionregistrationv1MatchCondition `json:"matchConditions,omitempty"` // MatchConditions is a list of conditions that must be met for a request to be sent to this webhook. Match conditions filter requests that have already been matched by the rules, namespaceSelector, and objectSelector. An empty list of matchConditions matches all requests. There are a maximum of 64 match conditions allowed. The exact matching logic is (in order): 1. If ANY matchCondition evaluates to FALSE, the webhook is skipped. 2. If ALL matchConditions evaluate to TRUE, the webhook is called. 3. If any matchCondition evaluates to an error (but none are FALSE): - If failurePolicy=Fail, reject the request - If failurePolicy=Ignore, the error is ignored and the webhook is skipped
	Reinvocationpolicy string `json:"reinvocationPolicy,omitempty"` // reinvocationPolicy indicates whether this webhook should be called multiple times as part of a single admission evaluation. Allowed values are "Never" and "IfNeeded". Never: the webhook will not be called more than once in a single admission evaluation. IfNeeded: the webhook will be called at least one additional time as part of the admission evaluation if the object being admitted is modified by other admission plugins after the initial webhook call. Webhooks that specify this option *must* be idempotent, able to process objects they previously admitted. Note: * the number of additional invocations is not guaranteed to be exactly one. * if additional invocations result in further modifications to the object, webhooks are not guaranteed to be invoked again. * webhooks that use this option may be reordered to minimize the number of additional invocations. * to validate an object after all mutations are guaranteed complete, use a validating admission webhook instead. Defaults to "Never".
	Rules []Iok8sapiadmissionregistrationv1RuleWithOperations `json:"rules,omitempty"` // Rules describes what operations on what resources/subresources the webhook cares about. The webhook cares about an operation if it matches _any_ Rule. However, in order to prevent ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks from putting the cluster in a state which cannot be recovered from without completely disabling the plugin, ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called on admission requests for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.
	Name string `json:"name"` // The name of the admission webhook. Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where "imagepolicy" is the name of the webhook, and kubernetes.io is the name of the organization. Required.
	Admissionreviewversions []string `json:"admissionReviewVersions"` // AdmissionReviewVersions is an ordered list of preferred `AdmissionReview` versions the Webhook expects. API server will try to use first version in the list which it supports. If none of the versions specified in this list supported by API server, validation will fail for this object. If a persisted webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail and be subject to the failure policy.
	Failurepolicy string `json:"failurePolicy,omitempty"` // FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed values are Ignore or Fail. Defaults to Fail.
	Objectselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"objectSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Clientconfig Iok8sapiadmissionregistrationv1WebhookClientConfig `json:"clientConfig"` // WebhookClientConfig contains the information to make a TLS connection with the webhook
	Matchpolicy string `json:"matchPolicy,omitempty"` // matchPolicy defines how the "rules" list is used to match incoming requests. Allowed values are "Exact" or "Equivalent". - Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the webhook. - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the webhook. Defaults to "Equivalent"
	Namespaceselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"namespaceSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
}

// Iok8sapiresourcev1beta2ResourceClaim represents the Iok8sapiresourcev1beta2ResourceClaim schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ResourceClaim struct {
	Status Iok8sapiresourcev1beta2ResourceClaimStatus `json:"status,omitempty"` // ResourceClaimStatus tracks whether the resource has been allocated and what the result of that was.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiresourcev1beta2ResourceClaimSpec `json:"spec"` // ResourceClaimSpec defines what is being requested in a ResourceClaim and how to configure it.
}

// Iok8sapicorev1EndpointSubset represents the Iok8sapicorev1EndpointSubset schema from the OpenAPI specification
type Iok8sapicorev1EndpointSubset struct {
	Notreadyaddresses []Iok8sapicorev1EndpointAddress `json:"notReadyAddresses,omitempty"` // IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check.
	Ports []Iok8sapicorev1EndpointPort `json:"ports,omitempty"` // Port numbers available on the related IP addresses.
	Addresses []Iok8sapicorev1EndpointAddress `json:"addresses,omitempty"` // IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize.
}

// Iok8sapicorev1CinderPersistentVolumeSource represents the Iok8sapicorev1CinderPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1CinderPersistentVolumeSource struct {
	Fstype string `json:"fsType,omitempty"` // fsType Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	Readonly bool `json:"readOnly,omitempty"` // readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	Secretref Iok8sapicorev1SecretReference `json:"secretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	Volumeid string `json:"volumeID"` // volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
}

// Iok8sapieventsv1EventList represents the Iok8sapieventsv1EventList schema from the OpenAPI specification
type Iok8sapieventsv1EventList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapieventsv1Event `json:"items"` // items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiresourcev1beta2DeviceClaimConfiguration represents the Iok8sapiresourcev1beta2DeviceClaimConfiguration schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceClaimConfiguration struct {
	Opaque Iok8sapiresourcev1beta2OpaqueDeviceConfiguration `json:"opaque,omitempty"` // OpaqueDeviceConfiguration contains configuration parameters for a driver in a format defined by the driver vendor.
	Requests []string `json:"requests,omitempty"` // Requests lists the names of requests where the configuration applies. If empty, it applies to all requests. References to subrequests must include the name of the main request and may include the subrequest using the format <main request>[/<subrequest>]. If just the main request is given, the configuration applies to all subrequests.
}

// Iok8sapiautoscalingv2ContainerResourceMetricStatus represents the Iok8sapiautoscalingv2ContainerResourceMetricStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2ContainerResourceMetricStatus struct {
	Current Iok8sapiautoscalingv2MetricValueStatus `json:"current"` // MetricValueStatus holds the current value for a metric
	Name string `json:"name"` // name is the name of the resource in question.
	Container string `json:"container"` // container is the name of the container in the pods of the scaling target
}

// Iok8sapiauthorizationv1SelfSubjectRulesReviewSpec represents the Iok8sapiauthorizationv1SelfSubjectRulesReviewSpec schema from the OpenAPI specification
type Iok8sapiauthorizationv1SelfSubjectRulesReviewSpec struct {
	Namespace string `json:"namespace,omitempty"` // Namespace to evaluate rules for. Required.
}

// Iok8sapicorev1ContainerStatus represents the Iok8sapicorev1ContainerStatus schema from the OpenAPI specification
type Iok8sapicorev1ContainerStatus struct {
	Containerid string `json:"containerID,omitempty"` // ContainerID is the ID of the container in the format '<type>://<container_id>'. Where type is a container runtime identifier, returned from Version call of CRI API (for example "containerd").
	Image string `json:"image"` // Image is the name of container image that the container is running. The container image may not match the image used in the PodSpec, as it may have been resolved by the runtime. More info: https://kubernetes.io/docs/concepts/containers/images.
	Resources Iok8sapicorev1ResourceRequirements `json:"resources,omitempty"` // ResourceRequirements describes the compute resource requirements.
	State Iok8sapicorev1ContainerState `json:"state,omitempty"` // ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.
	Laststate Iok8sapicorev1ContainerState `json:"lastState,omitempty"` // ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.
	Imageid string `json:"imageID"` // ImageID is the image ID of the container's image. The image ID may not match the image ID of the image used in the PodSpec, as it may have been resolved by the runtime.
	Stopsignal string `json:"stopSignal,omitempty"` // StopSignal reports the effective stop signal for this container
	Name string `json:"name"` // Name is a DNS_LABEL representing the unique name of the container. Each container in a pod must have a unique name across all container types. Cannot be updated.
	Restartcount int `json:"restartCount"` // RestartCount holds the number of times the container has been restarted. Kubelet makes an effort to always increment the value, but there are cases when the state may be lost due to node restarts and then the value may be reset to 0. The value is never negative.
	Started bool `json:"started,omitempty"` // Started indicates whether the container has finished its postStart lifecycle hook and passed its startup probe. Initialized as false, becomes true after startupProbe is considered successful. Resets to false when the container is restarted, or if kubelet loses state temporarily. In both cases, startup probes will run again. Is always true when no startupProbe is defined and container is running and has passed the postStart lifecycle hook. The null value must be treated the same as false.
	Volumemounts []Iok8sapicorev1VolumeMountStatus `json:"volumeMounts,omitempty"` // Status of volume mounts.
	Ready bool `json:"ready"` // Ready specifies whether the container is currently passing its readiness check. The value will change as readiness probes keep executing. If no readiness probes are specified, this field defaults to true once the container is fully started (see Started field). The value is typically used to determine whether a container is ready to accept traffic.
	Allocatedresourcesstatus []Iok8sapicorev1ResourceStatus `json:"allocatedResourcesStatus,omitempty"` // AllocatedResourcesStatus represents the status of various resources allocated for this Pod.
	Allocatedresources map[string]interface{} `json:"allocatedResources,omitempty"` // AllocatedResources represents the compute resources allocated for this container by the node. Kubelet sets this value to Container.Resources.Requests upon successful pod admission and after successfully admitting desired pod resize.
	User Iok8sapicorev1ContainerUser `json:"user,omitempty"` // ContainerUser represents user identity information
}

// Iok8sapicorev1SecretProjection represents the Iok8sapicorev1SecretProjection schema from the OpenAPI specification
type Iok8sapicorev1SecretProjection struct {
	Items []Iok8sapicorev1KeyToPath `json:"items,omitempty"` // items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
	Name string `json:"name,omitempty"` // Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // optional field specify whether the Secret or its key must be defined
}

// Iok8sapiautoscalingv2HorizontalPodAutoscalerBehavior represents the Iok8sapiautoscalingv2HorizontalPodAutoscalerBehavior schema from the OpenAPI specification
type Iok8sapiautoscalingv2HorizontalPodAutoscalerBehavior struct {
	Scaledown Iok8sapiautoscalingv2HPAScalingRules `json:"scaleDown,omitempty"` // HPAScalingRules configures the scaling behavior for one direction via scaling Policy Rules and a configurable metric tolerance. Scaling Policy Rules are applied after calculating DesiredReplicas from metrics for the HPA. They can limit the scaling velocity by specifying scaling policies. They can prevent flapping by specifying the stabilization window, so that the number of replicas is not set instantly, instead, the safest value from the stabilization window is chosen. The tolerance is applied to the metric values and prevents scaling too eagerly for small metric variations. (Note that setting a tolerance requires enabling the alpha HPAConfigurableTolerance feature gate.)
	Scaleup Iok8sapiautoscalingv2HPAScalingRules `json:"scaleUp,omitempty"` // HPAScalingRules configures the scaling behavior for one direction via scaling Policy Rules and a configurable metric tolerance. Scaling Policy Rules are applied after calculating DesiredReplicas from metrics for the HPA. They can limit the scaling velocity by specifying scaling policies. They can prevent flapping by specifying the stabilization window, so that the number of replicas is not set instantly, instead, the safest value from the stabilization window is chosen. The tolerance is applied to the metric values and prevents scaling too eagerly for small metric variations. (Note that setting a tolerance requires enabling the alpha HPAConfigurableTolerance feature gate.)
}

// Iok8sapistoragev1StorageClassList represents the Iok8sapistoragev1StorageClassList schema from the OpenAPI specification
type Iok8sapistoragev1StorageClassList struct {
	Items []Iok8sapistoragev1StorageClass `json:"items"` // items is the list of StorageClasses
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1VsphereVirtualDiskVolumeSource represents the Iok8sapicorev1VsphereVirtualDiskVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1VsphereVirtualDiskVolumeSource struct {
	Volumepath string `json:"volumePath"` // volumePath is the path that identifies vSphere volume vmdk
	Fstype string `json:"fsType,omitempty"` // fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	Storagepolicyid string `json:"storagePolicyID,omitempty"` // storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
	Storagepolicyname string `json:"storagePolicyName,omitempty"` // storagePolicyName is the storage Policy Based Management (SPBM) profile name.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Openapiv3schema GeneratedType `json:"openAPIV3Schema,omitempty"` // JSONSchemaProps is a JSON-Schema following Specification Draft 4 (http://json-schema.org/).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Fieldpath string `json:"fieldPath,omitempty"` // fieldPath represents the field path returned when the validation fails. It must be a relative JSON path (i.e. with array notation) scoped to the location of this x-kubernetes-validations extension in the schema and refer to an existing field. e.g. when validation checks if a specific attribute `foo` under a map `testMap`, the fieldPath could be set to `.testMap.foo` If the validation checks two lists must have unique attributes, the fieldPath could be set to either of the list: e.g. `.testList` It does not support list numeric index. It supports child operation to refer to an existing field currently. Refer to [JSONPath support in Kubernetes](https://kubernetes.io/docs/reference/kubectl/jsonpath/) for more info. Numeric index of array is not supported. For field name which contains special characters, use `['specialName']` to refer the field name. e.g. for attribute `foo.34$` appears in a list `testList`, the fieldPath could be set to `.testList['foo.34$']`
	Message string `json:"message,omitempty"` // Message represents the message displayed when validation fails. The message is required if the Rule contains line breaks. The message must not contain line breaks. If unset, the message is "failed rule: {Rule}". e.g. "must be a URL with the host matching spec.host"
	Messageexpression string `json:"messageExpression,omitempty"` // MessageExpression declares a CEL expression that evaluates to the validation failure message that is returned when this rule fails. Since messageExpression is used as a failure message, it must evaluate to a string. If both message and messageExpression are present on a rule, then messageExpression will be used if validation fails. If messageExpression results in a runtime error, the runtime error is logged, and the validation failure message is produced as if the messageExpression field were unset. If messageExpression evaluates to an empty string, a string with only spaces, or a string that contains line breaks, then the validation failure message will also be produced as if the messageExpression field were unset, and the fact that messageExpression produced an empty string/string with only spaces/string with line breaks will be logged. messageExpression has access to all the same variables as the rule; the only difference is the return type. Example: "x must be less than max ("+string(self.max)+")"
	Optionaloldself bool `json:"optionalOldSelf,omitempty"` // optionalOldSelf is used to opt a transition rule into evaluation even when the object is first created, or if the old object is missing the value. When enabled `oldSelf` will be a CEL optional whose value will be `None` if there is no old value, or when the object is initially created. You may check for presence of oldSelf using `oldSelf.hasValue()` and unwrap it after checking using `oldSelf.value()`. Check the CEL documentation for Optional types for more information: https://pkg.go.dev/github.com/google/cel-go/cel#OptionalTypes May not be set unless `oldSelf` is used in `rule`.
	Reason string `json:"reason,omitempty"` // reason provides a machine-readable validation failure reason that is returned to the caller when a request fails this validation rule. The HTTP status code returned to the caller will match the reason of the reason of the first failed validation rule. The currently supported reasons are: "FieldValueInvalid", "FieldValueForbidden", "FieldValueRequired", "FieldValueDuplicate". If not set, default to use "FieldValueInvalid". All future added reasons must be accepted by clients when reading this value and unknown reasons should be treated as FieldValueInvalid.
	Rule string `json:"rule"` // Rule represents the expression which will be evaluated by CEL. ref: https://github.com/google/cel-spec The Rule is scoped to the location of the x-kubernetes-validations extension in the schema. The `self` variable in the CEL expression is bound to the scoped value. Example: - Rule scoped to the root of a resource with a status subresource: {"rule": "self.status.actual <= self.spec.maxDesired"} If the Rule is scoped to an object with properties, the accessible properties of the object are field selectable via `self.field` and field presence can be checked via `has(self.field)`. Null valued fields are treated as absent fields in CEL expressions. If the Rule is scoped to an object with additionalProperties (i.e. a map) the value of the map are accessible via `self[mapKey]`, map containment can be checked via `mapKey in self` and all entries of the map are accessible via CEL macros and functions such as `self.all(...)`. If the Rule is scoped to an array, the elements of the array are accessible via `self[i]` and also by macros and functions. If the Rule is scoped to a scalar, `self` is bound to the scalar value. Examples: - Rule scoped to a map of objects: {"rule": "self.components['Widget'].priority < 10"} - Rule scoped to a list of integers: {"rule": "self.values.all(value, value >= 0 && value < 100)"} - Rule scoped to a string value: {"rule": "self.startsWith('kube')"} The `apiVersion`, `kind`, `metadata.name` and `metadata.generateName` are always accessible from the root of the object and from any x-kubernetes-embedded-resource annotated objects. No other metadata properties are accessible. Unknown data preserved in custom resources via x-kubernetes-preserve-unknown-fields is not accessible in CEL expressions. This includes: - Unknown field values that are preserved by object schemas with x-kubernetes-preserve-unknown-fields. - Object properties where the property schema is of an "unknown type". An "unknown type" is recursively defined as: - A schema with no type and x-kubernetes-preserve-unknown-fields set to true - An array where the items schema is of an "unknown type" - An object where the additionalProperties schema is of an "unknown type" Only property names of the form `[a-zA-Z_.-/][a-zA-Z0-9_.-/]*` are accessible. Accessible property names are escaped according to the following rules when accessed in the expression: - '__' escapes to '__underscores__' - '.' escapes to '__dot__' - '-' escapes to '__dash__' - '/' escapes to '__slash__' - Property names that exactly match a CEL RESERVED keyword escape to '__{keyword}__'. The keywords are: 	 "true", "false", "null", "in", "as", "break", "const", "continue", "else", "for", "function", "if", 	 "import", "let", "loop", "package", "namespace", "return". Examples: - Rule accessing a property named "namespace": {"rule": "self.__namespace__ > 0"} - Rule accessing a property named "x-prop": {"rule": "self.x__dash__prop > 0"} - Rule accessing a property named "redact__d": {"rule": "self.redact__underscores__d > 0"} Equality on arrays with x-kubernetes-list-type of 'set' or 'map' ignores element order, i.e. [1, 2] == [2, 1]. Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type: - 'set': `X + Y` performs a union where the array positions of all elements in `X` are preserved and non-intersecting elements in `Y` are appended, retaining their partial order. - 'map': `X + Y` performs a merge where the array positions of all keys in `X` are preserved but the values are overwritten by values in `Y` when the key sets of `X` and `Y` intersect. Elements in `Y` with non-intersecting keys are appended, retaining their partial order. If `rule` makes use of the `oldSelf` variable it is implicitly a `transition rule`. By default, the `oldSelf` variable is the same type as `self`. When `optionalOldSelf` is true, the `oldSelf` variable is a CEL optional variable whose value() is the same type as `self`. See the documentation for the `optionalOldSelf` field for details. Transition rules by default are applied only on UPDATE requests and are skipped if an old value could not be found. You can opt a transition rule into unconditional evaluation by setting `optionalOldSelf` to true.
}

// Iok8sapiresourcev1beta2ResourceSliceList represents the Iok8sapiresourcev1beta2ResourceSliceList schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ResourceSliceList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiresourcev1beta2ResourceSlice `json:"items"` // Items is the list of resource ResourceSlices.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiapiserverinternalv1alpha1ServerStorageVersion represents the Iok8sapiapiserverinternalv1alpha1ServerStorageVersion schema from the OpenAPI specification
type Iok8sapiapiserverinternalv1alpha1ServerStorageVersion struct {
	Apiserverid string `json:"apiServerID,omitempty"` // The ID of the reporting API server.
	Decodableversions []string `json:"decodableVersions,omitempty"` // The API server can decode objects encoded in these versions. The encodingVersion must be included in the decodableVersions.
	Encodingversion string `json:"encodingVersion,omitempty"` // The API server encodes the object to this version when persisting it in the backend (e.g., etcd).
	Servedversions []string `json:"servedVersions,omitempty"` // The API server can serve these versions. DecodableVersions must include all ServedVersions.
}

// Iok8sapicorev1ResourceQuotaStatus represents the Iok8sapicorev1ResourceQuotaStatus schema from the OpenAPI specification
type Iok8sapicorev1ResourceQuotaStatus struct {
	Hard map[string]interface{} `json:"hard,omitempty"` // Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Used map[string]interface{} `json:"used,omitempty"` // Used is the current observed total usage of the resource in the namespace.
}

// Iok8sapicorev1ServicePort represents the Iok8sapicorev1ServicePort schema from the OpenAPI specification
type Iok8sapicorev1ServicePort struct {
	Protocol string `json:"protocol,omitempty"` // The IP protocol for this port. Supports "TCP", "UDP", and "SCTP". Default is TCP.
	Targetport string `json:"targetPort,omitempty"` // IntOrString is a type that can hold an int32 or a string. When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type. This allows you to have, for example, a JSON field that can accept a name or number.
	Appprotocol string `json:"appProtocol,omitempty"` // The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either: * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior- * 'kubernetes.io/ws' - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455 * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.
	Name string `json:"name,omitempty"` // The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.
	Nodeport int `json:"nodePort,omitempty"` // The port on each node on which this service is exposed when type is NodePort or LoadBalancer. Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail. If not specified, a port will be allocated if this Service requires one. If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	Port int `json:"port"` // The port that will be exposed by this service.
}

// Iok8sapicorev1PodSchedulingGate represents the Iok8sapicorev1PodSchedulingGate schema from the OpenAPI specification
type Iok8sapicorev1PodSchedulingGate struct {
	Name string `json:"name"` // Name of the scheduling gate. Each scheduling gate must have a unique name field.
}

// Iok8sapiauthorizationv1SelfSubjectAccessReview represents the Iok8sapiauthorizationv1SelfSubjectAccessReview schema from the OpenAPI specification
type Iok8sapiauthorizationv1SelfSubjectAccessReview struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiauthorizationv1SelfSubjectAccessReviewSpec `json:"spec"` // SelfSubjectAccessReviewSpec is a description of the access request. Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
	Status Iok8sapiauthorizationv1SubjectAccessReviewStatus `json:"status,omitempty"` // SubjectAccessReviewStatus
}

// Iok8sapinetworkingv1NetworkPolicySpec represents the Iok8sapinetworkingv1NetworkPolicySpec schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicySpec struct {
	Podselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"podSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Policytypes []string `json:"policyTypes,omitempty"` // policyTypes is a list of rule types that the NetworkPolicy relates to. Valid options are ["Ingress"], ["Egress"], or ["Ingress", "Egress"]. If this field is not specified, it will default based on the existence of ingress or egress rules; policies that contain an egress section are assumed to affect egress, and all policies (whether or not they contain an ingress section) are assumed to affect ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ "Egress" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include "Egress" (since such a policy would not include an egress section and would otherwise default to just [ "Ingress" ]). This field is beta-level in 1.8
	Egress []Iok8sapinetworkingv1NetworkPolicyEgressRule `json:"egress,omitempty"` // egress is a list of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8
	Ingress []Iok8sapinetworkingv1NetworkPolicyIngressRule `json:"ingress,omitempty"` // ingress is a list of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)
}

// Iok8sapicorev1Sysctl represents the Iok8sapicorev1Sysctl schema from the OpenAPI specification
type Iok8sapicorev1Sysctl struct {
	Name string `json:"name"` // Name of a property to set
	Value string `json:"value"` // Value of a property to set
}

// Iok8sapicorev1ContainerPort represents the Iok8sapicorev1ContainerPort schema from the OpenAPI specification
type Iok8sapicorev1ContainerPort struct {
	Hostip string `json:"hostIP,omitempty"` // What host IP to bind the external port to.
	Hostport int `json:"hostPort,omitempty"` // Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
	Name string `json:"name,omitempty"` // If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.
	Protocol string `json:"protocol,omitempty"` // Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
	Containerport int `json:"containerPort"` // Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.
}

// Iok8sapiappsv1DeploymentSpec represents the Iok8sapiappsv1DeploymentSpec schema from the OpenAPI specification
type Iok8sapiappsv1DeploymentSpec struct {
	Replicas int `json:"replicas,omitempty"` // Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Strategy Iok8sapiappsv1DeploymentStrategy `json:"strategy,omitempty"` // DeploymentStrategy describes how to replace existing pods with new ones.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	Paused bool `json:"paused,omitempty"` // Indicates that the deployment is paused.
	Progressdeadlineseconds int `json:"progressDeadlineSeconds,omitempty"` // The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.
}

// Iok8sapicorev1PersistentVolumeClaimList represents the Iok8sapicorev1PersistentVolumeClaimList schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaimList struct {
	Items []Iok8sapicorev1PersistentVolumeClaim `json:"items"` // items is a list of persistent volume claims. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapiauthenticationv1TokenReviewStatus represents the Iok8sapiauthenticationv1TokenReviewStatus schema from the OpenAPI specification
type Iok8sapiauthenticationv1TokenReviewStatus struct {
	Audiences []string `json:"audiences,omitempty"` // Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token. An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token's audiences. A client of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns an empty status.audience field where status.authenticated is "true", the token is valid against the audience of the Kubernetes API server.
	Authenticated bool `json:"authenticated,omitempty"` // Authenticated indicates that the token was associated with a known user.
	ErrorField string `json:"error,omitempty"` // Error indicates that the token couldn't be checked
	User Iok8sapiauthenticationv1UserInfo `json:"user,omitempty"` // UserInfo holds the information about the user needed to implement the user.Info interface.
}

// Iok8sapiappsv1Deployment represents the Iok8sapiappsv1Deployment schema from the OpenAPI specification
type Iok8sapiappsv1Deployment struct {
	Status Iok8sapiappsv1DeploymentStatus `json:"status,omitempty"` // DeploymentStatus is the most recently observed status of the Deployment.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiappsv1DeploymentSpec `json:"spec,omitempty"` // DeploymentSpec is the specification of the desired behavior of the Deployment.
}

// Iok8sapiflowcontrolv1NonResourcePolicyRule represents the Iok8sapiflowcontrolv1NonResourcePolicyRule schema from the OpenAPI specification
type Iok8sapiflowcontrolv1NonResourcePolicyRule struct {
	Nonresourceurls []string `json:"nonResourceURLs"` // `nonResourceURLs` is a set of url prefixes that a user should have access to and may not be empty. For example: - "/healthz" is legal - "/hea*" is illegal - "/hea" is legal but matches nothing - "/hea/*" also matches nothing - "/healthz/*" matches all per-component health checks. "*" matches all non-resource urls. if it is present, it must be the only entry. Required.
	Verbs []string `json:"verbs"` // `verbs` is a list of matching verbs and may not be empty. "*" matches all verbs. If it is present, it must be the only entry. Required.
}

// Iok8sapicorev1HostPathVolumeSource represents the Iok8sapicorev1HostPathVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1HostPathVolumeSource struct {
	Path string `json:"path"` // path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	TypeField string `json:"type,omitempty"` // type for HostPath Volume Defaults to "" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
}

// Iok8sapistoragemigrationv1alpha1StorageVersionMigrationSpec represents the Iok8sapistoragemigrationv1alpha1StorageVersionMigrationSpec schema from the OpenAPI specification
type Iok8sapistoragemigrationv1alpha1StorageVersionMigrationSpec struct {
	Resource Iok8sapistoragemigrationv1alpha1GroupVersionResource `json:"resource"` // The names of the group, the version, and the resource.
	Continuetoken string `json:"continueToken,omitempty"` // The token used in the list options to get the next chunk of objects to migrate. When the .status.conditions indicates the migration is "Running", users can use this token to check the progress of the migration.
}

// Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyBinding represents the Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyBinding schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyBinding struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiadmissionregistrationv1alpha1MutatingAdmissionPolicyBindingSpec `json:"spec,omitempty"` // MutatingAdmissionPolicyBindingSpec is the specification of the MutatingAdmissionPolicyBinding.
}

// Iok8sapiresourcev1beta1ResourceClaimTemplateList represents the Iok8sapiresourcev1beta1ResourceClaimTemplateList schema from the OpenAPI specification
type Iok8sapiresourcev1beta1ResourceClaimTemplateList struct {
	Items []Iok8sapiresourcev1beta1ResourceClaimTemplate `json:"items"` // Items is the list of resource claim templates.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1Affinity represents the Iok8sapicorev1Affinity schema from the OpenAPI specification
type Iok8sapicorev1Affinity struct {
	Podaffinity Iok8sapicorev1PodAffinity `json:"podAffinity,omitempty"` // Pod affinity is a group of inter pod affinity scheduling rules.
	Podantiaffinity Iok8sapicorev1PodAntiAffinity `json:"podAntiAffinity,omitempty"` // Pod anti affinity is a group of inter pod anti affinity scheduling rules.
	Nodeaffinity Iok8sapicorev1NodeAffinity `json:"nodeAffinity,omitempty"` // Node affinity is a group of node affinity scheduling rules.
}

// Iok8sapiadmissionregistrationv1MatchResources represents the Iok8sapiadmissionregistrationv1MatchResources schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1MatchResources struct {
	Excluderesourcerules []Iok8sapiadmissionregistrationv1NamedRuleWithOperations `json:"excludeResourceRules,omitempty"` // ExcludeResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy should not care about. The exclude rules take precedence over include rules (if a resource matches both, it is excluded)
	Matchpolicy string `json:"matchPolicy,omitempty"` // matchPolicy defines how the "MatchResources" list is used to match incoming requests. Allowed values are "Exact" or "Equivalent". - Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the ValidatingAdmissionPolicy. - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the ValidatingAdmissionPolicy. Defaults to "Equivalent"
	Namespaceselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"namespaceSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Objectselector Iok8sapimachinerypkgapismetav1LabelSelector `json:"objectSelector,omitempty"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Resourcerules []Iok8sapiadmissionregistrationv1NamedRuleWithOperations `json:"resourceRules,omitempty"` // ResourceRules describes what operations on what resources/subresources the ValidatingAdmissionPolicy matches. The policy cares about an operation if it matches _any_ Rule.
}

// Iok8sapiresourcev1beta2DeviceToleration represents the Iok8sapiresourcev1beta2DeviceToleration schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceToleration struct {
	Operator string `json:"operator,omitempty"` // Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a ResourceClaim can tolerate all taints of a particular category.
	Tolerationseconds int64 `json:"tolerationSeconds,omitempty"` // TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system. If larger than zero, the time when the pod needs to be evicted is calculated as <time when taint was adedd> + <toleration seconds>.
	Value string `json:"value,omitempty"` // Value is the taint value the toleration matches to. If the operator is Exists, the value must be empty, otherwise just a regular string. Must be a label value.
	Effect string `json:"effect,omitempty"` // Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule and NoExecute.
	Key string `json:"key,omitempty"` // Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys. Must be a label name.
}

// Iok8sapicorev1Pod represents the Iok8sapicorev1Pod schema from the OpenAPI specification
type Iok8sapicorev1Pod struct {
	Spec Iok8sapicorev1PodSpec `json:"spec,omitempty"` // PodSpec is a description of a pod.
	Status Iok8sapicorev1PodStatus `json:"status,omitempty"` // PodStatus represents information about the status of a pod. Status may trail the actual state of a system, especially if the node that hosts the pod cannot contact the control plane.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1ResourceQuotaList represents the Iok8sapicorev1ResourceQuotaList schema from the OpenAPI specification
type Iok8sapicorev1ResourceQuotaList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1ResourceQuota `json:"items"` // Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapibatchv1JobTemplateSpec represents the Iok8sapibatchv1JobTemplateSpec schema from the OpenAPI specification
type Iok8sapibatchv1JobTemplateSpec struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapibatchv1JobSpec `json:"spec,omitempty"` // JobSpec describes how the job execution will look like.
}

// Iok8sapiresourcev1beta1DeviceClassList represents the Iok8sapiresourcev1beta1DeviceClassList schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceClassList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiresourcev1beta1DeviceClass `json:"items"` // Items is the list of resource classes.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapidiscoveryv1EndpointSliceList represents the Iok8sapidiscoveryv1EndpointSliceList schema from the OpenAPI specification
type Iok8sapidiscoveryv1EndpointSliceList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapidiscoveryv1EndpointSlice `json:"items"` // items is the list of endpoint slices
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiflowcontrolv1FlowSchema represents the Iok8sapiflowcontrolv1FlowSchema schema from the OpenAPI specification
type Iok8sapiflowcontrolv1FlowSchema struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiflowcontrolv1FlowSchemaSpec `json:"spec,omitempty"` // FlowSchemaSpec describes how the FlowSchema's specification looks like.
	Status Iok8sapiflowcontrolv1FlowSchemaStatus `json:"status,omitempty"` // FlowSchemaStatus represents the current state of a FlowSchema.
}

// Iok8sapicorev1ConfigMapProjection represents the Iok8sapicorev1ConfigMapProjection schema from the OpenAPI specification
type Iok8sapicorev1ConfigMapProjection struct {
	Name string `json:"name,omitempty"` // Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // optional specify whether the ConfigMap or its keys must be defined
	Items []Iok8sapicorev1KeyToPath `json:"items,omitempty"` // items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
}

// Iok8sapipolicyv1Eviction represents the Iok8sapipolicyv1Eviction schema from the OpenAPI specification
type Iok8sapipolicyv1Eviction struct {
	Deleteoptions Iok8sapimachinerypkgapismetav1DeleteOptions `json:"deleteOptions,omitempty"` // DeleteOptions may be provided when deleting an API object.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapicorev1PersistentVolumeSpec represents the Iok8sapicorev1PersistentVolumeSpec schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeSpec struct {
	Volumeattributesclassname string `json:"volumeAttributesClassName,omitempty"` // Name of VolumeAttributesClass to which this persistent volume belongs. Empty value is not allowed. When this field is not set, it indicates that this volume does not belong to any VolumeAttributesClass. This field is mutable and can be changed by the CSI driver after a volume has been updated successfully to a new class. For an unbound PersistentVolume, the volumeAttributesClassName will be matched with unbound PersistentVolumeClaims during the binding process. This is a beta field and requires enabling VolumeAttributesClass feature (off by default).
	Flexvolume Iok8sapicorev1FlexPersistentVolumeSource `json:"flexVolume,omitempty"` // FlexPersistentVolumeSource represents a generic persistent volume resource that is provisioned/attached using an exec based plugin.
	Scaleio Iok8sapicorev1ScaleIOPersistentVolumeSource `json:"scaleIO,omitempty"` // ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume
	Persistentvolumereclaimpolicy string `json:"persistentVolumeReclaimPolicy,omitempty"` // persistentVolumeReclaimPolicy defines what happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	Storageclassname string `json:"storageClassName,omitempty"` // storageClassName is the name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
	Azuredisk Iok8sapicorev1AzureDiskVolumeSource `json:"azureDisk,omitempty"` // AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	Cephfs Iok8sapicorev1CephFSPersistentVolumeSource `json:"cephfs,omitempty"` // Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
	Glusterfs Iok8sapicorev1GlusterfsPersistentVolumeSource `json:"glusterfs,omitempty"` // Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
	Claimref Iok8sapicorev1ObjectReference `json:"claimRef,omitempty"` // ObjectReference contains enough information to let you inspect or modify the referred object.
	Volumemode string `json:"volumeMode,omitempty"` // volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.
	Rbd Iok8sapicorev1RBDPersistentVolumeSource `json:"rbd,omitempty"` // Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
	Csi Iok8sapicorev1CSIPersistentVolumeSource `json:"csi,omitempty"` // Represents storage that is managed by an external CSI volume driver
	Gcepersistentdisk Iok8sapicorev1GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty"` // Represents a Persistent Disk resource in Google Compute Engine. A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.
	Hostpath Iok8sapicorev1HostPathVolumeSource `json:"hostPath,omitempty"` // Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.
	Nodeaffinity Iok8sapicorev1VolumeNodeAffinity `json:"nodeAffinity,omitempty"` // VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.
	Portworxvolume Iok8sapicorev1PortworxVolumeSource `json:"portworxVolume,omitempty"` // PortworxVolumeSource represents a Portworx volume resource.
	Awselasticblockstore Iok8sapicorev1AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"` // Represents a Persistent Disk resource in AWS. An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.
	Iscsi Iok8sapicorev1ISCSIPersistentVolumeSource `json:"iscsi,omitempty"` // ISCSIPersistentVolumeSource represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
	Flocker Iok8sapicorev1FlockerVolumeSource `json:"flocker,omitempty"` // Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
	Accessmodes []string `json:"accessModes,omitempty"` // accessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	Photonpersistentdisk Iok8sapicorev1PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty"` // Represents a Photon Controller persistent disk resource.
	Azurefile Iok8sapicorev1AzureFilePersistentVolumeSource `json:"azureFile,omitempty"` // AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	Vspherevolume Iok8sapicorev1VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty"` // Represents a vSphere volume resource.
	Mountoptions []string `json:"mountOptions,omitempty"` // mountOptions is the list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	Local Iok8sapicorev1LocalVolumeSource `json:"local,omitempty"` // Local represents directly-attached storage with node affinity
	Nfs Iok8sapicorev1NFSVolumeSource `json:"nfs,omitempty"` // Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.
	Capacity map[string]interface{} `json:"capacity,omitempty"` // capacity is the description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Cinder Iok8sapicorev1CinderPersistentVolumeSource `json:"cinder,omitempty"` // Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
	Fc Iok8sapicorev1FCVolumeSource `json:"fc,omitempty"` // Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
	Quobyte Iok8sapicorev1QuobyteVolumeSource `json:"quobyte,omitempty"` // Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
	Storageos Iok8sapicorev1StorageOSPersistentVolumeSource `json:"storageos,omitempty"` // Represents a StorageOS persistent volume resource.
}

// Iok8sapicorev1GRPCAction represents the Iok8sapicorev1GRPCAction schema from the OpenAPI specification
type Iok8sapicorev1GRPCAction struct {
	Port int `json:"port"` // Port number of the gRPC service. Number must be in the range 1 to 65535.
	Service string `json:"service,omitempty"` // Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.
}

// Iok8sapimachinerypkgapismetav1StatusDetails represents the Iok8sapimachinerypkgapismetav1StatusDetails schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1StatusDetails struct {
	Kind string `json:"kind,omitempty"` // The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Name string `json:"name,omitempty"` // The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).
	Retryafterseconds int `json:"retryAfterSeconds,omitempty"` // If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.
	Uid string `json:"uid,omitempty"` // UID of the resource. (when there is a single resource which can be described). More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids
	Causes []Iok8sapimachinerypkgapismetav1StatusCause `json:"causes,omitempty"` // The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.
	Group string `json:"group,omitempty"` // The group attribute of the resource associated with the status StatusReason.
}

// Iok8sapicorev1Endpoints represents the Iok8sapicorev1Endpoints schema from the OpenAPI specification
type Iok8sapicorev1Endpoints struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Subsets []Iok8sapicorev1EndpointSubset `json:"subsets,omitempty"` // The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service.
}

// Iok8sapinetworkingv1beta1ServiceCIDRList represents the Iok8sapinetworkingv1beta1ServiceCIDRList schema from the OpenAPI specification
type Iok8sapinetworkingv1beta1ServiceCIDRList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapinetworkingv1beta1ServiceCIDR `json:"items"` // items is the list of ServiceCIDRs.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapinodev1RuntimeClassList represents the Iok8sapinodev1RuntimeClassList schema from the OpenAPI specification
type Iok8sapinodev1RuntimeClassList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapinodev1RuntimeClass `json:"items"` // items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1NodeList represents the Iok8sapicorev1NodeList schema from the OpenAPI specification
type Iok8sapicorev1NodeList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1Node `json:"items"` // List of nodes
}

// Iok8sapiappsv1DaemonSetSpec represents the Iok8sapiappsv1DaemonSetSpec schema from the OpenAPI specification
type Iok8sapiappsv1DaemonSetSpec struct {
	Minreadyseconds int `json:"minReadySeconds,omitempty"` // The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
	Revisionhistorylimit int `json:"revisionHistoryLimit,omitempty"` // The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	Selector Iok8sapimachinerypkgapismetav1LabelSelector `json:"selector"` // A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Template Iok8sapicorev1PodTemplateSpec `json:"template"` // PodTemplateSpec describes the data a pod should have when created from a template
	Updatestrategy Iok8sapiappsv1DaemonSetUpdateStrategy `json:"updateStrategy,omitempty"` // DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
}

// Iok8sapicorev1CephFSVolumeSource represents the Iok8sapicorev1CephFSVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1CephFSVolumeSource struct {
	User string `json:"user,omitempty"` // user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	Monitors []string `json:"monitors"` // monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	Path string `json:"path,omitempty"` // path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	Readonly bool `json:"readOnly,omitempty"` // readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	Secretfile string `json:"secretFile,omitempty"` // secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef,omitempty"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
}

// Iok8sapicorev1SecurityContext represents the Iok8sapicorev1SecurityContext schema from the OpenAPI specification
type Iok8sapicorev1SecurityContext struct {
	Apparmorprofile Iok8sapicorev1AppArmorProfile `json:"appArmorProfile,omitempty"` // AppArmorProfile defines a pod or container's AppArmor settings.
	Privileged bool `json:"privileged,omitempty"` // Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.
	Readonlyrootfilesystem bool `json:"readOnlyRootFilesystem,omitempty"` // Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.
	Seccompprofile Iok8sapicorev1SeccompProfile `json:"seccompProfile,omitempty"` // SeccompProfile defines a pod/container's seccomp profile settings. Only one profile source may be set.
	Allowprivilegeescalation bool `json:"allowPrivilegeEscalation,omitempty"` // AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.
	Procmount string `json:"procMount,omitempty"` // procMount denotes the type of proc mount to use for the containers. The default value is Default which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.
	Runasgroup int64 `json:"runAsGroup,omitempty"` // The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.
	Runasnonroot bool `json:"runAsNonRoot,omitempty"` // Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	Selinuxoptions Iok8sapicorev1SELinuxOptions `json:"seLinuxOptions,omitempty"` // SELinuxOptions are the labels to be applied to the container
	Capabilities Iok8sapicorev1Capabilities `json:"capabilities,omitempty"` // Adds and removes POSIX capabilities from running containers.
	Runasuser int64 `json:"runAsUser,omitempty"` // The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.
	Windowsoptions Iok8sapicorev1WindowsSecurityContextOptions `json:"windowsOptions,omitempty"` // WindowsSecurityContextOptions contain Windows-specific options and credentials.
}

// Iok8sapicorev1ClientIPConfig represents the Iok8sapicorev1ClientIPConfig schema from the OpenAPI specification
type Iok8sapicorev1ClientIPConfig struct {
	Timeoutseconds int `json:"timeoutSeconds,omitempty"` // timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == "ClientIP". Default value is 10800(for 3 hours).
}

// Iok8sapibatchv1PodFailurePolicyOnPodConditionsPattern represents the Iok8sapibatchv1PodFailurePolicyOnPodConditionsPattern schema from the OpenAPI specification
type Iok8sapibatchv1PodFailurePolicyOnPodConditionsPattern struct {
	Status string `json:"status"` // Specifies the required Pod condition status. To match a pod condition it is required that the specified status equals the pod condition status. Defaults to True.
	TypeField string `json:"type"` // Specifies the required Pod condition type. To match a pod condition it is required that specified type equals the pod condition type.
}

// Iok8sapimachinerypkgapismetav1FieldSelectorRequirement represents the Iok8sapimachinerypkgapismetav1FieldSelectorRequirement schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1FieldSelectorRequirement struct {
	Operator string `json:"operator"` // operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. The list of operators may grow in the future.
	Values []string `json:"values,omitempty"` // values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.
	Key string `json:"key"` // key is the field selector key that the requirement applies to.
}

// Iok8sapimachinerypkgapismetav1ManagedFieldsEntry represents the Iok8sapimachinerypkgapismetav1ManagedFieldsEntry schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1ManagedFieldsEntry struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the version of this resource that this field set applies to. The format is "group/version" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted.
	Fieldstype string `json:"fieldsType,omitempty"` // FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: "FieldsV1"
	Fieldsv1 Iok8sapimachinerypkgapismetav1FieldsV1 `json:"fieldsV1,omitempty"` // FieldsV1 stores a set of fields in a data structure like a Trie, in JSON format. Each key is either a '.' representing the field itself, and will always map to an empty set, or a string representing a sub-field or item. The string will follow one of these four formats: 'f:<name>', where <name> is the name of a field in a struct, or key in a map 'v:<value>', where <value> is the exact json formatted value of a list item 'i:<index>', where <index> is position of a item in a list 'k:<keys>', where <keys> is a map of a list item's key fields to their unique values If a key maps to an empty Fields value, the field that key represents is part of the set. The exact format is defined in sigs.k8s.io/structured-merge-diff
	Manager string `json:"manager,omitempty"` // Manager is an identifier of the workflow managing these fields.
	Operation string `json:"operation,omitempty"` // Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.
	Subresource string `json:"subresource,omitempty"` // Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource.
	Time string `json:"time,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
}

// Iok8sapinetworkingv1IngressStatus represents the Iok8sapinetworkingv1IngressStatus schema from the OpenAPI specification
type Iok8sapinetworkingv1IngressStatus struct {
	Loadbalancer Iok8sapinetworkingv1IngressLoadBalancerStatus `json:"loadBalancer,omitempty"` // IngressLoadBalancerStatus represents the status of a load-balancer.
}

// Iok8sapiresourcev1beta1ResourceClaim represents the Iok8sapiresourcev1beta1ResourceClaim schema from the OpenAPI specification
type Iok8sapiresourcev1beta1ResourceClaim struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiresourcev1beta1ResourceClaimSpec `json:"spec"` // ResourceClaimSpec defines what is being requested in a ResourceClaim and how to configure it.
	Status Iok8sapiresourcev1beta1ResourceClaimStatus `json:"status,omitempty"` // ResourceClaimStatus tracks whether the resource has been allocated and what the result of that was.
}

// Iok8sapiresourcev1beta2ResourceClaimList represents the Iok8sapiresourcev1beta2ResourceClaimList schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ResourceClaimList struct {
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiresourcev1beta2ResourceClaim `json:"items"` // Items is the list of resource claims.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapiautoscalingv2PodsMetricStatus represents the Iok8sapiautoscalingv2PodsMetricStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2PodsMetricStatus struct {
	Current Iok8sapiautoscalingv2MetricValueStatus `json:"current"` // MetricValueStatus holds the current value for a metric
	Metric Iok8sapiautoscalingv2MetricIdentifier `json:"metric"` // MetricIdentifier defines the name and optionally selector for a metric
}

// Iok8sapibatchv1JobCondition represents the Iok8sapibatchv1JobCondition schema from the OpenAPI specification
type Iok8sapibatchv1JobCondition struct {
	Status string `json:"status"` // Status of the condition, one of True, False, Unknown.
	TypeField string `json:"type"` // Type of job condition, Complete or Failed.
	Lastprobetime string `json:"lastProbeTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Lasttransitiontime string `json:"lastTransitionTime,omitempty"` // Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.
	Message string `json:"message,omitempty"` // Human readable message indicating details about last transition.
	Reason string `json:"reason,omitempty"` // (brief) reason for the condition's last transition.
}

// Iok8sapiappsv1RollingUpdateStatefulSetStrategy represents the Iok8sapiappsv1RollingUpdateStatefulSetStrategy schema from the OpenAPI specification
type Iok8sapiappsv1RollingUpdateStatefulSetStrategy struct {
	Maxunavailable string `json:"maxUnavailable,omitempty"` // IntOrString is a type that can hold an int32 or a string. When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type. This allows you to have, for example, a JSON field that can accept a name or number.
	Partition int `json:"partition,omitempty"` // Partition indicates the ordinal at which the StatefulSet should be partitioned for updates. During a rolling update, all pods from ordinal Replicas-1 to Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched. This is helpful in being able to do a canary based deployment. The default value is 0.
}

// Iok8sapinetworkingv1ServiceCIDR represents the Iok8sapinetworkingv1ServiceCIDR schema from the OpenAPI specification
type Iok8sapinetworkingv1ServiceCIDR struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapinetworkingv1ServiceCIDRSpec `json:"spec,omitempty"` // ServiceCIDRSpec define the CIDRs the user wants to use for allocating ClusterIPs for Services.
	Status Iok8sapinetworkingv1ServiceCIDRStatus `json:"status,omitempty"` // ServiceCIDRStatus describes the current state of the ServiceCIDR.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapimachinerypkgapismetav1GroupVersionForDiscovery represents the Iok8sapimachinerypkgapismetav1GroupVersionForDiscovery schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1GroupVersionForDiscovery struct {
	Groupversion string `json:"groupVersion"` // groupVersion specifies the API group and version in the form "group/version"
	Version string `json:"version"` // version specifies the version in the form of "version". This is to save the clients the trouble of splitting the GroupVersion.
}

// Iok8sapimachinerypkgapismetav1APIGroupList represents the Iok8sapimachinerypkgapismetav1APIGroupList schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1APIGroupList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Groups []Iok8sapimachinerypkgapismetav1APIGroup `json:"groups"` // groups is a list of APIGroup.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapiresourcev1beta2DeviceClassList represents the Iok8sapiresourcev1beta2DeviceClassList schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceClassList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiresourcev1beta2DeviceClass `json:"items"` // Items is the list of resource classes.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapiautoscalingv2ResourceMetricStatus represents the Iok8sapiautoscalingv2ResourceMetricStatus schema from the OpenAPI specification
type Iok8sapiautoscalingv2ResourceMetricStatus struct {
	Name string `json:"name"` // name is the name of the resource in question.
	Current Iok8sapiautoscalingv2MetricValueStatus `json:"current"` // MetricValueStatus holds the current value for a metric
}

// Iok8sapicorev1PersistentVolumeClaim represents the Iok8sapicorev1PersistentVolumeClaim schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolumeClaim struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1PersistentVolumeClaimSpec `json:"spec,omitempty"` // PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes
	Status Iok8sapicorev1PersistentVolumeClaimStatus `json:"status,omitempty"` // PersistentVolumeClaimStatus is the current status of a persistent volume claim.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapibatchv1PodFailurePolicy represents the Iok8sapibatchv1PodFailurePolicy schema from the OpenAPI specification
type Iok8sapibatchv1PodFailurePolicy struct {
	Rules []Iok8sapibatchv1PodFailurePolicyRule `json:"rules"` // A list of pod failure policy rules. The rules are evaluated in order. Once a rule matches a Pod failure, the remaining of the rules are ignored. When no rule matches the Pod failure, the default handling applies - the counter of pod failures is incremented and it is checked against the backoffLimit. At most 20 elements are allowed.
}

// Iok8sapibatchv1SuccessPolicyRule represents the Iok8sapibatchv1SuccessPolicyRule schema from the OpenAPI specification
type Iok8sapibatchv1SuccessPolicyRule struct {
	Succeededindexes string `json:"succeededIndexes,omitempty"` // succeededIndexes specifies the set of indexes which need to be contained in the actual set of the succeeded indexes for the Job. The list of indexes must be within 0 to ".spec.completions-1" and must not contain duplicates. At least one element is required. The indexes are represented as intervals separated by commas. The intervals can be a decimal integer or a pair of decimal integers separated by a hyphen. The number are listed in represented by the first and last element of the series, separated by a hyphen. For example, if the completed indexes are 1, 3, 4, 5 and 7, they are represented as "1,3-5,7". When this field is null, this field doesn't default to any value and is never evaluated at any time.
	Succeededcount int `json:"succeededCount,omitempty"` // succeededCount specifies the minimal required size of the actual set of the succeeded indexes for the Job. When succeededCount is used along with succeededIndexes, the check is constrained only to the set of indexes specified by succeededIndexes. For example, given that succeededIndexes is "1-4", succeededCount is "3", and completed indexes are "1", "3", and "5", the Job isn't declared as succeeded because only "1" and "3" indexes are considered in that rules. When this field is null, this doesn't default to any value and is never evaluated at any time. When specified it needs to be a positive integer.
}

// Iok8sapiauthorizationv1NonResourceRule represents the Iok8sapiauthorizationv1NonResourceRule schema from the OpenAPI specification
type Iok8sapiauthorizationv1NonResourceRule struct {
	Nonresourceurls []string `json:"nonResourceURLs,omitempty"` // NonResourceURLs is a set of partial urls that a user should have access to. *s are allowed, but only as the full, final step in the path. "*" means all.
	Verbs []string `json:"verbs"` // Verb is a list of kubernetes non-resource API verbs, like: get, post, put, delete, patch, head, options. "*" means all.
}

// Iok8sapinetworkingv1NetworkPolicyIngressRule represents the Iok8sapinetworkingv1NetworkPolicyIngressRule schema from the OpenAPI specification
type Iok8sapinetworkingv1NetworkPolicyIngressRule struct {
	Ports []Iok8sapinetworkingv1NetworkPolicyPort `json:"ports,omitempty"` // ports is a list of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.
	From []Iok8sapinetworkingv1NetworkPolicyPeer `json:"from,omitempty"` // from is a list of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the from list.
}

// Iok8sapicorev1ConfigMapList represents the Iok8sapicorev1ConfigMapList schema from the OpenAPI specification
type Iok8sapicorev1ConfigMapList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicorev1ConfigMap `json:"items"` // Items is the list of ConfigMaps.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapicorev1PersistentVolume represents the Iok8sapicorev1PersistentVolume schema from the OpenAPI specification
type Iok8sapicorev1PersistentVolume struct {
	Status Iok8sapicorev1PersistentVolumeStatus `json:"status,omitempty"` // PersistentVolumeStatus is the current status of a persistent volume.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapicorev1PersistentVolumeSpec `json:"spec,omitempty"` // PersistentVolumeSpec is the specification of a persistent volume.
}

// Iok8sapinetworkingv1ServiceCIDRList represents the Iok8sapinetworkingv1ServiceCIDRList schema from the OpenAPI specification
type Iok8sapinetworkingv1ServiceCIDRList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapinetworkingv1ServiceCIDR `json:"items"` // items is the list of ServiceCIDRs.
}

// Iok8sapicorev1AzureFilePersistentVolumeSource represents the Iok8sapicorev1AzureFilePersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1AzureFilePersistentVolumeSource struct {
	Sharename string `json:"shareName"` // shareName is the azure Share Name
	Readonly bool `json:"readOnly,omitempty"` // readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Secretname string `json:"secretName"` // secretName is the name of secret that contains Azure Storage Account Name and Key
	Secretnamespace string `json:"secretNamespace,omitempty"` // secretNamespace is the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod
}

// Iok8sapiresourcev1beta2ResourceClaimTemplateList represents the Iok8sapiresourcev1beta2ResourceClaimTemplateList schema from the OpenAPI specification
type Iok8sapiresourcev1beta2ResourceClaimTemplateList struct {
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapiresourcev1beta2ResourceClaimTemplate `json:"items"` // Items is the list of resource claim templates.
}

// Iok8sapiautoscalingv2ObjectMetricSource represents the Iok8sapiautoscalingv2ObjectMetricSource schema from the OpenAPI specification
type Iok8sapiautoscalingv2ObjectMetricSource struct {
	Target Iok8sapiautoscalingv2MetricTarget `json:"target"` // MetricTarget defines the target value, average value, or average utilization of a specific metric
	Describedobject Iok8sapiautoscalingv2CrossVersionObjectReference `json:"describedObject"` // CrossVersionObjectReference contains enough information to let you identify the referred resource.
	Metric Iok8sapiautoscalingv2MetricIdentifier `json:"metric"` // MetricIdentifier defines the name and optionally selector for a metric
}

// Iok8sapipolicyv1PodDisruptionBudgetList represents the Iok8sapipolicyv1PodDisruptionBudgetList schema from the OpenAPI specification
type Iok8sapipolicyv1PodDisruptionBudgetList struct {
	Items []Iok8sapipolicyv1PodDisruptionBudget `json:"items"` // Items is a list of PodDisruptionBudgets
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapiresourcev1beta2DeviceAllocationResult represents the Iok8sapiresourcev1beta2DeviceAllocationResult schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceAllocationResult struct {
	Config []Iok8sapiresourcev1beta2DeviceAllocationConfiguration `json:"config,omitempty"` // This field is a combination of all the claim and class configuration parameters. Drivers can distinguish between those based on a flag. This includes configuration parameters for drivers which have no allocated devices in the result because it is up to the drivers which configuration parameters they support. They can silently ignore unknown configuration parameters.
	Results []Iok8sapiresourcev1beta2DeviceRequestAllocationResult `json:"results,omitempty"` // Results lists all allocated devices.
}

// Iok8sapiresourcev1beta2DeviceClass represents the Iok8sapiresourcev1beta2DeviceClass schema from the OpenAPI specification
type Iok8sapiresourcev1beta2DeviceClass struct {
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapiresourcev1beta2DeviceClassSpec `json:"spec"` // DeviceClassSpec is used in a [DeviceClass] to define what can be allocated and how to configure it.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
}

// Iok8sapicorev1AppArmorProfile represents the Iok8sapicorev1AppArmorProfile schema from the OpenAPI specification
type Iok8sapicorev1AppArmorProfile struct {
	TypeField string `json:"type"` // type indicates which kind of AppArmor profile will be applied. Valid options are: Localhost - a profile pre-loaded on the node. RuntimeDefault - the container runtime's default profile. Unconfined - no AppArmor enforcement.
	Localhostprofile string `json:"localhostProfile,omitempty"` // localhostProfile indicates a profile loaded on the node that should be used. The profile must be preconfigured on the node to work. Must match the loaded name of the profile. Must be set if and only if type is "Localhost".
}

// Iok8sapidiscoveryv1ForZone represents the Iok8sapidiscoveryv1ForZone schema from the OpenAPI specification
type Iok8sapidiscoveryv1ForZone struct {
	Name string `json:"name"` // name represents the name of the zone.
}

// Iok8sapicorev1CephFSPersistentVolumeSource represents the Iok8sapicorev1CephFSPersistentVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1CephFSPersistentVolumeSource struct {
	Monitors []string `json:"monitors"` // monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	Path string `json:"path,omitempty"` // path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	Readonly bool `json:"readOnly,omitempty"` // readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	Secretfile string `json:"secretFile,omitempty"` // secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	Secretref Iok8sapicorev1SecretReference `json:"secretRef,omitempty"` // SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	User string `json:"user,omitempty"` // user is Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
}

// Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicy represents the Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicy schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicy struct {
	Spec Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicySpec `json:"spec,omitempty"` // ValidatingAdmissionPolicySpec is the specification of the desired behavior of the AdmissionPolicy.
	Status Iok8sapiadmissionregistrationv1ValidatingAdmissionPolicyStatus `json:"status,omitempty"` // ValidatingAdmissionPolicyStatus represents the status of an admission validation policy.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapiautoscalingv2MetricTarget represents the Iok8sapiautoscalingv2MetricTarget schema from the OpenAPI specification
type Iok8sapiautoscalingv2MetricTarget struct {
	Value string `json:"value,omitempty"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
	Averageutilization int `json:"averageUtilization,omitempty"` // averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type
	Averagevalue string `json:"averageValue,omitempty"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
	TypeField string `json:"type"` // type represents whether the metric type is Utilization, Value, or AverageValue
}

// Iok8sapicorev1NodeFeatures represents the Iok8sapicorev1NodeFeatures schema from the OpenAPI specification
type Iok8sapicorev1NodeFeatures struct {
	Supplementalgroupspolicy bool `json:"supplementalGroupsPolicy,omitempty"` // SupplementalGroupsPolicy is set to true if the runtime supports SupplementalGroupsPolicy and ContainerUser.
}

// Iok8sapicorev1SecretEnvSource represents the Iok8sapicorev1SecretEnvSource schema from the OpenAPI specification
type Iok8sapicorev1SecretEnvSource struct {
	Name string `json:"name,omitempty"` // Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Optional bool `json:"optional,omitempty"` // Specify whether the Secret must be defined
}

// Iok8sapieventsv1EventSeries represents the Iok8sapieventsv1EventSeries schema from the OpenAPI specification
type Iok8sapieventsv1EventSeries struct {
	Count int `json:"count"` // count is the number of occurrences in this series up to the last heartbeat time.
	Lastobservedtime string `json:"lastObservedTime"` // MicroTime is version of Time with microsecond level precision.
}

// Iok8sapiadmissionregistrationv1alpha1NamedRuleWithOperations represents the Iok8sapiadmissionregistrationv1alpha1NamedRuleWithOperations schema from the OpenAPI specification
type Iok8sapiadmissionregistrationv1alpha1NamedRuleWithOperations struct {
	Apigroups []string `json:"apiGroups,omitempty"` // APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.
	Apiversions []string `json:"apiVersions,omitempty"` // APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.
	Operations []string `json:"operations,omitempty"` // Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If '*' is present, the length of the slice must be one. Required.
	Resourcenames []string `json:"resourceNames,omitempty"` // ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.
	Resources []string `json:"resources,omitempty"` // Resources is a list of resources this rule applies to. For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources. If wildcard is present, the validation rule will ensure resources do not overlap with each other. Depending on the enclosing object, subresources might not be allowed. Required.
	Scope string `json:"scope,omitempty"` // scope specifies the scope of this rule. Valid values are "Cluster", "Namespaced", and "*" "Cluster" means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. "Namespaced" means that only namespaced resources will match this rule. "*" means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is "*".
}

// Iok8sapiresourcev1alpha3DeviceSelector represents the Iok8sapiresourcev1alpha3DeviceSelector schema from the OpenAPI specification
type Iok8sapiresourcev1alpha3DeviceSelector struct {
	Cel Iok8sapiresourcev1alpha3CELDeviceSelector `json:"cel,omitempty"` // CELDeviceSelector contains a CEL expression for selecting a device.
}

// Iok8sapicorev1ScaleIOVolumeSource represents the Iok8sapicorev1ScaleIOVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1ScaleIOVolumeSource struct {
	System string `json:"system"` // system is the name of the storage system as configured in ScaleIO.
	Secretref Iok8sapicorev1LocalObjectReference `json:"secretRef"` // LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	Sslenabled bool `json:"sslEnabled,omitempty"` // sslEnabled Flag enable/disable SSL communication with Gateway, default false
	Readonly bool `json:"readOnly,omitempty"` // readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	Storagemode string `json:"storageMode,omitempty"` // storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.
	Storagepool string `json:"storagePool,omitempty"` // storagePool is the ScaleIO Storage Pool associated with the protection domain.
	Gateway string `json:"gateway"` // gateway is the host address of the ScaleIO API Gateway.
	Protectiondomain string `json:"protectionDomain,omitempty"` // protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.
	Volumename string `json:"volumeName,omitempty"` // volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.
	Fstype string `json:"fsType,omitempty"` // fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs".
}

// Iok8sapistoragemigrationv1alpha1StorageVersionMigrationStatus represents the Iok8sapistoragemigrationv1alpha1StorageVersionMigrationStatus schema from the OpenAPI specification
type Iok8sapistoragemigrationv1alpha1StorageVersionMigrationStatus struct {
	Conditions []Iok8sapistoragemigrationv1alpha1MigrationCondition `json:"conditions,omitempty"` // The latest available observations of the migration's current state.
	Resourceversion string `json:"resourceVersion,omitempty"` // ResourceVersion to compare with the GC cache for performing the migration. This is the current resource version of given group, version and resource when kube-controller-manager first observes this StorageVersionMigration resource.
}

// Iok8sapicorev1EndpointPort represents the Iok8sapicorev1EndpointPort schema from the OpenAPI specification
type Iok8sapicorev1EndpointPort struct {
	Appprotocol string `json:"appProtocol,omitempty"` // The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either: * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior- * 'kubernetes.io/ws' - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455 * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.
	Name string `json:"name,omitempty"` // The name of this port. This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
	Port int `json:"port"` // The port number of the endpoint.
	Protocol string `json:"protocol,omitempty"` // The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Conditions []GeneratedType `json:"conditions,omitempty"` // Current service state of apiService.
}

// Iok8sapicorev1ScopedResourceSelectorRequirement represents the Iok8sapicorev1ScopedResourceSelectorRequirement schema from the OpenAPI specification
type Iok8sapicorev1ScopedResourceSelectorRequirement struct {
	Scopename string `json:"scopeName"` // The name of the scope that the selector applies to.
	Values []string `json:"values,omitempty"` // An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	Operator string `json:"operator"` // Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.
}

// Iok8sapicorev1PodAntiAffinity represents the Iok8sapicorev1PodAntiAffinity schema from the OpenAPI specification
type Iok8sapicorev1PodAntiAffinity struct {
	Requiredduringschedulingignoredduringexecution []Iok8sapicorev1PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"` // If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
	Preferredduringschedulingignoredduringexecution []Iok8sapicorev1WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"` // The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and subtracting "weight" from the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
}

// Iok8sapimachinerypkgapismetav1LabelSelectorRequirement represents the Iok8sapimachinerypkgapismetav1LabelSelectorRequirement schema from the OpenAPI specification
type Iok8sapimachinerypkgapismetav1LabelSelectorRequirement struct {
	Operator string `json:"operator"` // operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
	Values []string `json:"values,omitempty"` // values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	Key string `json:"key"` // key is the label key that the selector applies to.
}

// Iok8sapicorev1ServiceAccount represents the Iok8sapicorev1ServiceAccount schema from the OpenAPI specification
type Iok8sapicorev1ServiceAccount struct {
	Automountserviceaccounttoken bool `json:"automountServiceAccountToken,omitempty"` // AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
	Imagepullsecrets []Iok8sapicorev1LocalObjectReference `json:"imagePullSecrets,omitempty"` // ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Secrets []Iok8sapicorev1ObjectReference `json:"secrets,omitempty"` // Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a "kubernetes.io/enforce-mountable-secrets" annotation set to "true". The "kubernetes.io/enforce-mountable-secrets" annotation is deprecated since v1.32. Prefer separate namespaces to isolate access to mounted secrets. This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
}

// Iok8sapiresourcev1beta1DeviceAllocationConfiguration represents the Iok8sapiresourcev1beta1DeviceAllocationConfiguration schema from the OpenAPI specification
type Iok8sapiresourcev1beta1DeviceAllocationConfiguration struct {
	Requests []string `json:"requests,omitempty"` // Requests lists the names of requests where the configuration applies. If empty, its applies to all requests. References to subrequests must include the name of the main request and may include the subrequest using the format <main request>[/<subrequest>]. If just the main request is given, the configuration applies to all subrequests.
	Source string `json:"source"` // Source records whether the configuration comes from a class and thus is not something that a normal user would have been able to set or from a claim.
	Opaque Iok8sapiresourcev1beta1OpaqueDeviceConfiguration `json:"opaque,omitempty"` // OpaqueDeviceConfiguration contains configuration parameters for a driver in a format defined by the driver vendor.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

// Iok8sapirbacv1RoleBinding represents the Iok8sapirbacv1RoleBinding schema from the OpenAPI specification
type Iok8sapirbacv1RoleBinding struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Roleref Iok8sapirbacv1RoleRef `json:"roleRef"` // RoleRef contains information that points to the role being used
	Subjects []Iok8sapirbacv1Subject `json:"subjects,omitempty"` // Subjects holds references to the objects the role applies to.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref string `json:"$ref,omitempty"`
	Items GeneratedType `json:"items,omitempty"` // JSONSchemaPropsOrArray represents a value that can either be a JSONSchemaProps or an array of JSONSchemaProps. Mainly here for serialization purposes.
	Minitems int64 `json:"minItems,omitempty"`
	X_kubernetes_preserve_unknown_fields bool `json:"x-kubernetes-preserve-unknown-fields,omitempty"` // x-kubernetes-preserve-unknown-fields stops the API server decoding step from pruning fields which are not specified in the validation schema. This affects fields recursively, but switches back to normal pruning behaviour if nested properties or additionalProperties are specified in the schema. This can either be true or undefined. False is forbidden.
	Additionalproperties GeneratedType `json:"additionalProperties,omitempty"` // JSONSchemaPropsOrBool represents JSONSchemaProps or a boolean value. Defaults to true for the boolean property.
	Maxitems int64 `json:"maxItems,omitempty"`
	Enum []GeneratedType `json:"enum,omitempty"`
	Format string `json:"format,omitempty"` // format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated: - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like "0321751043" or "978-0321751041" - isbn10: an ISBN10 number string like "0321751043" - isbn13: an ISBN13 number string like "978-0321751041" - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^\\d{3}[- ]?\\d{2}[- ]?\\d{4}$ - hexcolor: an hexadecimal color code like "#FFFFFF: following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like "rgb(255,255,2559" - byte: base64 encoded binary data - password: any kind of string - date: a date string like "2006-01-02" as defined by full-date in RFC3339 - duration: a duration string like "22 ns" as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like "2014-12-15T19:30:20.000Z" as defined by date-time in RFC3339.
	X_kubernetes_list_map_keys []string `json:"x-kubernetes-list-map-keys,omitempty"` // x-kubernetes-list-map-keys annotates an array with the x-kubernetes-list-type `map` by specifying the keys used as the index of the map. This tag MUST only be used on lists that have the "x-kubernetes-list-type" extension set to "map". Also, the values specified for this attribute must be a scalar typed field of the child structure (no nesting is supported). The properties specified must either be required or have a default value, to ensure those properties are present for all list items.
	Maximum float64 `json:"maximum,omitempty"`
	Additionalitems GeneratedType `json:"additionalItems,omitempty"` // JSONSchemaPropsOrBool represents JSONSchemaProps or a boolean value. Defaults to true for the boolean property.
	Example GeneratedType `json:"example,omitempty"` // JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.
	Minimum float64 `json:"minimum,omitempty"`
	Definitions map[string]interface{} `json:"definitions,omitempty"`
	Uniqueitems bool `json:"uniqueItems,omitempty"`
	Not GeneratedType `json:"not,omitempty"` // JSONSchemaProps is a JSON-Schema following Specification Draft 4 (http://json-schema.org/).
	Oneof []GeneratedType `json:"oneOf,omitempty"`
	Exclusiveminimum bool `json:"exclusiveMinimum,omitempty"`
	Id string `json:"id,omitempty"`
	Anyof []GeneratedType `json:"anyOf,omitempty"`
	Maxlength int64 `json:"maxLength,omitempty"`
	Minproperties int64 `json:"minProperties,omitempty"`
	Required []string `json:"required,omitempty"`
	Dependencies map[string]interface{} `json:"dependencies,omitempty"`
	Title string `json:"title,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	DefaultField GeneratedType `json:"default,omitempty"` // JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.
	Multipleof float64 `json:"multipleOf,omitempty"`
	X_kubernetes_map_type string `json:"x-kubernetes-map-type,omitempty"` // x-kubernetes-map-type annotates an object to further describe its topology. This extension must only be used when type is object and may have 2 possible values: 1) `granular`: These maps are actual maps (key-value pairs) and each fields are independent from each other (they can each be manipulated by separate actors). This is the default behaviour for all maps. 2) `atomic`: the list is treated as a single entity, like a scalar. Atomic maps will be entirely replaced when updated.
	Pattern string `json:"pattern,omitempty"`
	Patternproperties map[string]interface{} `json:"patternProperties,omitempty"`
	Exclusivemaximum bool `json:"exclusiveMaximum,omitempty"`
	X_kubernetes_validations []GeneratedType `json:"x-kubernetes-validations,omitempty"` // x-kubernetes-validations describes a list of validation rules written in the CEL expression language.
	X_kubernetes_list_type string `json:"x-kubernetes-list-type,omitempty"` // x-kubernetes-list-type annotates an array to further describe its topology. This extension must only be used on lists and may have 3 possible values: 1) `atomic`: the list is treated as a single entity, like a scalar. Atomic lists will be entirely replaced when updated. This extension may be used on any type of list (struct, scalar, ...). 2) `set`: Sets are lists that must not have multiple items with the same value. Each value must be a scalar, an object with x-kubernetes-map-type `atomic` or an array with x-kubernetes-list-type `atomic`. 3) `map`: These lists are like maps in that their elements have a non-index key used to identify them. Order is preserved upon merge. The map tag must only be used on a list with elements of type object. Defaults to atomic for arrays.
	Description string `json:"description,omitempty"`
	Externaldocs GeneratedType `json:"externalDocs,omitempty"` // ExternalDocumentation allows referencing an external resource for extended documentation.
	Allof []GeneratedType `json:"allOf,omitempty"`
	Schema string `json:"$schema,omitempty"`
	Maxproperties int64 `json:"maxProperties,omitempty"`
	X_kubernetes_int_or_string bool `json:"x-kubernetes-int-or-string,omitempty"` // x-kubernetes-int-or-string specifies that this value is either an integer or a string. If this is true, an empty type is allowed and type as child of anyOf is permitted if following one of the following patterns: 1) anyOf: - type: integer - type: string 2) allOf: - anyOf: - type: integer - type: string - ... zero or more
	X_kubernetes_embedded_resource bool `json:"x-kubernetes-embedded-resource,omitempty"` // x-kubernetes-embedded-resource defines that the value is an embedded Kubernetes runtime.Object, with TypeMeta and ObjectMeta. The type must be object. It is allowed to further restrict the embedded object. kind, apiVersion and metadata are validated automatically. x-kubernetes-preserve-unknown-fields is allowed to be true, but does not have to be if the object is fully specified (up to kind, apiVersion, metadata).
	Nullable bool `json:"nullable,omitempty"`
	Minlength int64 `json:"minLength,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// Iok8sapiautoscalingv2ResourceMetricSource represents the Iok8sapiautoscalingv2ResourceMetricSource schema from the OpenAPI specification
type Iok8sapiautoscalingv2ResourceMetricSource struct {
	Name string `json:"name"` // name is the name of the resource in question.
	Target Iok8sapiautoscalingv2MetricTarget `json:"target"` // MetricTarget defines the target value, average value, or average utilization of a specific metric
}

// Iok8sapinetworkingv1ParentReference represents the Iok8sapinetworkingv1ParentReference schema from the OpenAPI specification
type Iok8sapinetworkingv1ParentReference struct {
	Namespace string `json:"namespace,omitempty"` // Namespace is the namespace of the object being referenced.
	Resource string `json:"resource"` // Resource is the resource of the object being referenced.
	Group string `json:"group,omitempty"` // Group is the group of the object being referenced.
	Name string `json:"name"` // Name is the name of the object being referenced.
}

// Iok8sapicoordinationv1alpha2LeaseCandidateList represents the Iok8sapicoordinationv1alpha2LeaseCandidateList schema from the OpenAPI specification
type Iok8sapicoordinationv1alpha2LeaseCandidateList struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items []Iok8sapicoordinationv1alpha2LeaseCandidate `json:"items"` // items is a list of schema objects.
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ListMeta `json:"metadata,omitempty"` // ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
}

// Iok8sapidiscoveryv1EndpointPort represents the Iok8sapidiscoveryv1EndpointPort schema from the OpenAPI specification
type Iok8sapidiscoveryv1EndpointPort struct {
	Port int `json:"port,omitempty"` // port represents the port number of the endpoint. If the EndpointSlice is derived from a Kubernetes service, this must be set to the service's target port. EndpointSlices used for other purposes may have a nil port.
	Protocol string `json:"protocol,omitempty"` // protocol represents the IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
	Appprotocol string `json:"appProtocol,omitempty"` // The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either: * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior- * 'kubernetes.io/ws' - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455 * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.
	Name string `json:"name,omitempty"` // name represents the name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is derived from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.
}

// Iok8sapistoragev1CSIDriver represents the Iok8sapistoragev1CSIDriver schema from the OpenAPI specification
type Iok8sapistoragev1CSIDriver struct {
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Spec Iok8sapistoragev1CSIDriverSpec `json:"spec"` // CSIDriverSpec is the specification of a CSIDriver.
}

// Iok8sapiappsv1StatefulSet represents the Iok8sapiappsv1StatefulSet schema from the OpenAPI specification
type Iok8sapiappsv1StatefulSet struct {
	Spec Iok8sapiappsv1StatefulSetSpec `json:"spec,omitempty"` // A StatefulSetSpec is the specification of a StatefulSet.
	Status Iok8sapiappsv1StatefulSetStatus `json:"status,omitempty"` // StatefulSetStatus represents the current state of a StatefulSet.
	Apiversion string `json:"apiVersion,omitempty"` // APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind string `json:"kind,omitempty"` // Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata Iok8sapimachinerypkgapismetav1ObjectMeta `json:"metadata,omitempty"` // ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
}

// Iok8sapicorev1EmptyDirVolumeSource represents the Iok8sapicorev1EmptyDirVolumeSource schema from the OpenAPI specification
type Iok8sapicorev1EmptyDirVolumeSource struct {
	Medium string `json:"medium,omitempty"` // medium represents what type of storage medium should back this directory. The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	Sizelimit string `json:"sizeLimit,omitempty"` // Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ``` <quantity> ::= <signedNumber><suffix> 	(Note that <suffix> may be empty, from the "" case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= "+" | "-" <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei 	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | "" | k | M | G | T | P | E 	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber> ``` No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in "canonical form". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as "1500m" - 1.5Gi will be serialized as "1536Mi" Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
}
