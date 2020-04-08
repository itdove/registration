package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_BuildControllerConfig = map[string]string{
	"additionalTrustedCA": "additionalTrustedCA is a path to a pem bundle file containing additional CAs that should be trusted for image pushes and pulls during builds.",
}

func (BuildControllerConfig) SwaggerDoc() map[string]string {
	return map_BuildControllerConfig
}

var map_BuildDefaultsConfig = map[string]string{
	"":                       "BuildDefaultsConfig controls the default information for Builds",
	"gitHTTPProxy":           "gitHTTPProxy is the location of the HTTPProxy for Git source",
	"gitHTTPSProxy":          "gitHTTPSProxy is the location of the HTTPSProxy for Git source",
	"gitNoProxy":             "gitNoProxy is the list of domains for which the proxy should not be used",
	"env":                    "env is a set of default environment variables that will be applied to the build if the specified variables do not exist on the build",
	"sourceStrategyDefaults": "sourceStrategyDefaults are default values that apply to builds using the source strategy.",
	"imageLabels":            "imageLabels is a list of labels that are applied to the resulting image. User can override a default label by providing a label with the same name in their Build/BuildConfig.",
	"nodeSelector":           "nodeSelector is a selector which must be true for the build pod to fit on a node",
	"annotations":            "annotations are annotations that will be added to the build pod",
	"resources":              "resources defines resource requirements to execute the build.",
}

func (BuildDefaultsConfig) SwaggerDoc() map[string]string {
	return map_BuildDefaultsConfig
}

var map_BuildOverridesConfig = map[string]string{
	"":             "BuildOverridesConfig controls override settings for builds",
	"forcePull":    "forcePull indicates whether the build strategy should always be set to ForcePull=true",
	"imageLabels":  "imageLabels is a list of labels that are applied to the resulting image. If user provided a label in their Build/BuildConfig with the same name as one in this list, the user's label will be overwritten.",
	"nodeSelector": "nodeSelector is a selector which must be true for the build pod to fit on a node",
	"annotations":  "annotations are annotations that will be added to the build pod",
	"tolerations":  "tolerations is a list of Tolerations that will override any existing tolerations set on a build pod.",
}

func (BuildOverridesConfig) SwaggerDoc() map[string]string {
	return map_BuildOverridesConfig
}

var map_ClusterNetworkEntry = map[string]string{
	"":                 "ClusterNetworkEntry defines an individual cluster network. The CIDRs cannot overlap with other cluster network CIDRs, CIDRs reserved for external ips, CIDRs reserved for service networks, and CIDRs reserved for ingress ips.",
	"cidr":             "CIDR defines the total range of a cluster networks address space.",
	"hostSubnetLength": "HostSubnetLength is the number of bits of the accompanying CIDR address to allocate to each node. eg, 8 would mean that each node would have a /24 slice of the overlay network for its pod.",
}

func (ClusterNetworkEntry) SwaggerDoc() map[string]string {
	return map_ClusterNetworkEntry
}

var map_DockerPullSecretControllerConfig = map[string]string{
	"registryURLs":             "registryURLs is a list of urls that the docker pull secrets should be valid for.",
	"internalRegistryHostname": "internalRegistryHostname is the hostname for the default internal image registry. The value must be in \"hostname[:port]\" format.  Docker pull secrets will be generated for this registry.",
}

func (DockerPullSecretControllerConfig) SwaggerDoc() map[string]string {
	return map_DockerPullSecretControllerConfig
}

var map_FrontProxyConfig = map[string]string{
	"clientCA":            "clientCA is a path to the CA bundle to use to verify the common name of the front proxy's client cert",
	"allowedNames":        "allowedNames is an optional list of common names to require a match from.",
	"usernameHeaders":     "usernameHeaders is the set of headers to check for the username",
	"groupHeaders":        "groupHeaders is the set of headers to check for groups",
	"extraHeaderPrefixes": "extraHeaderPrefixes is the set of header prefixes to check for user extra",
}

func (FrontProxyConfig) SwaggerDoc() map[string]string {
	return map_FrontProxyConfig
}

var map_ImageConfig = map[string]string{
	"":       "ImageConfig holds the necessary configuration options for building image names for system components",
	"format": "Format is the format of the name to be built for the system component",
	"latest": "Latest determines if the latest tag will be pulled from the registry",
}

func (ImageConfig) SwaggerDoc() map[string]string {
	return map_ImageConfig
}

var map_ImageImportControllerConfig = map[string]string{
	"maxScheduledImageImportsPerMinute":          "maxScheduledImageImportsPerMinute is the maximum number of image streams that will be imported in the background per minute. The default value is 60. Set to -1 for unlimited.",
	"disableScheduledImport":                     "disableScheduledImport allows scheduled background import of images to be disabled.",
	"scheduledImageImportMinimumIntervalSeconds": "scheduledImageImportMinimumIntervalSeconds is the minimum number of seconds that can elapse between when image streams scheduled for background import are checked against the upstream repository. The default value is 15 minutes.",
}

func (ImageImportControllerConfig) SwaggerDoc() map[string]string {
	return map_ImageImportControllerConfig
}

var map_ImagePolicyConfig = map[string]string{
	"maxImagesBulkImportedPerRepository": "maxImagesBulkImportedPerRepository controls the number of images that are imported when a user does a bulk import of a container repository. This number is set low to prevent users from importing large numbers of images accidentally. Set -1 for no limit.",
	"allowedRegistriesForImport":         "allowedRegistriesForImport limits the container image registries that normal users may import images from. Set this list to the registries that you trust to contain valid Docker images and that you want applications to be able to import from. Users with permission to create Images or ImageStreamMappings via the API are not affected by this policy - typically only administrators or system integrations will have those permissions.",
	"internalRegistryHostname":           "internalRegistryHostname sets the hostname for the default internal image registry. The value must be in \"hostname[:port]\" format. For backward compatibility, users can still use OPENSHIFT_DEFAULT_REGISTRY environment variable but this setting overrides the environment variable.",
	"externalRegistryHostnames":          "externalRegistryHostnames provides the hostnames for the default external image registry. The external hostname should be set only when the image registry is exposed externally. The first value is used in 'publicDockerImageRepository' field in ImageStreams. The value must be in \"hostname[:port]\" format.",
	"additionalTrustedCA":                "additionalTrustedCA is a path to a pem bundle file containing additional CAs that should be trusted during imagestream import.",
}

func (ImagePolicyConfig) SwaggerDoc() map[string]string {
	return map_ImagePolicyConfig
}

var map_IngressControllerConfig = map[string]string{
	"ingressIPNetworkCIDR": "ingressIPNetworkCIDR controls the range to assign ingress ips from for services of type LoadBalancer on bare metal. If empty, ingress ips will not be assigned. It may contain a single CIDR that will be allocated from. For security reasons, you should ensure that this range does not overlap with the CIDRs reserved for external ips, nodes, pods, or services.",
}

func (IngressControllerConfig) SwaggerDoc() map[string]string {
	return map_IngressControllerConfig
}

var map_JenkinsPipelineConfig = map[string]string{
	"":                     "JenkinsPipelineConfig holds configuration for the Jenkins pipeline strategy",
	"autoProvisionEnabled": "autoProvisionEnabled determines whether a Jenkins server will be spawned from the provided template when the first build config in the project with type JenkinsPipeline is created. When not specified this option defaults to true.",
	"templateNamespace":    "templateNamespace contains the namespace name where the Jenkins template is stored",
	"templateName":         "templateName is the name of the default Jenkins template",
	"serviceName":          "serviceName is the name of the Jenkins service OpenShift uses to detect whether a Jenkins pipeline handler has already been installed in a project. This value *must* match a service name in the provided template.",
	"parameters":           "parameters specifies a set of optional parameters to the Jenkins template.",
}

func (JenkinsPipelineConfig) SwaggerDoc() map[string]string {
	return map_JenkinsPipelineConfig
}

var map_NetworkControllerConfig = map[string]string{
	"":                "MasterNetworkConfig to be passed to the compiled in network plugin",
	"clusterNetworks": "clusterNetworks contains a list of cluster networks that defines the global overlay networks L3 space.",
}

func (NetworkControllerConfig) SwaggerDoc() map[string]string {
	return map_NetworkControllerConfig
}

var map_OpenShiftAPIServerConfig = map[string]string{
	"aggregatorConfig":               "aggregatorConfig contains information about how to verify the aggregator front proxy",
	"imagePolicyConfig":              "imagePolicyConfig feeds the image policy admission plugin",
	"projectConfig":                  "projectConfig feeds an admission plugin",
	"routingConfig":                  "routingConfig holds information about routing and route generation",
	"serviceAccountOAuthGrantMethod": "serviceAccountOAuthGrantMethod is used for determining client authorization for service account oauth client. It must be either: deny, prompt, or \"\"",
	"jenkinsPipelineConfig":          "jenkinsPipelineConfig holds information about the default Jenkins template used for JenkinsPipeline build strategy.",
	"cloudProviderFile":              "cloudProviderFile points to the cloud config file",
}

func (OpenShiftAPIServerConfig) SwaggerDoc() map[string]string {
	return map_OpenShiftAPIServerConfig
}

var map_OpenShiftControllerManagerConfig = map[string]string{
	"servingInfo":    "servingInfo describes how to start serving",
	"leaderElection": "leaderElection defines the configuration for electing a controller instance to make changes to the cluster. If unspecified, the ControllerTTL value is checked to determine whether the legacy direct etcd election code will be used.",
	"controllers":    "controllers is a list of controllers to enable.  '*' enables all on-by-default controllers, 'foo' enables the controller \"+ named 'foo', '-foo' disables the controller named 'foo'. Defaults to \"*\".",
}

func (OpenShiftControllerManagerConfig) SwaggerDoc() map[string]string {
	return map_OpenShiftControllerManagerConfig
}

var map_ProjectConfig = map[string]string{
	"defaultNodeSelector":    "defaultNodeSelector holds default project node label selector",
	"projectRequestMessage":  "projectRequestMessage is the string presented to a user if they are unable to request a project via the projectrequest api endpoint",
	"projectRequestTemplate": "projectRequestTemplate is the template to use for creating projects in response to projectrequest. It is in the format namespace/template and it is optional. If it is not specified, a default template is used.",
}

func (ProjectConfig) SwaggerDoc() map[string]string {
	return map_ProjectConfig
}

var map_RegistryLocation = map[string]string{
	"":           "RegistryLocation contains a location of the registry specified by the registry domain name. The domain name might include wildcards, like '*' or '??'.",
	"domainName": "DomainName specifies a domain name for the registry In case the registry use non-standard (80 or 443) port, the port should be included in the domain name as well.",
	"insecure":   "Insecure indicates whether the registry is secure (https) or insecure (http) By default (if not specified) the registry is assumed as secure.",
}

func (RegistryLocation) SwaggerDoc() map[string]string {
	return map_RegistryLocation
}

var map_RoutingConfig = map[string]string{
	"":          "RoutingConfig holds the necessary configuration options for routing to subdomains",
	"subdomain": "subdomain is the suffix appended to $service.$namespace. to form the default route hostname DEPRECATED: This field is being replaced by routers setting their own defaults. This is the \"default\" route.",
}

func (RoutingConfig) SwaggerDoc() map[string]string {
	return map_RoutingConfig
}

var map_SecurityAllocator = map[string]string{
	"":                    "SecurityAllocator controls the automatic allocation of UIDs and MCS labels to a project. If nil, allocation is disabled.",
	"uidAllocatorRange":   "UIDAllocatorRange defines the total set of Unix user IDs (UIDs) that will be allocated to projects automatically, and the size of the block each namespace gets. For example, 1000-1999/10 will allocate ten UIDs per namespace, and will be able to allocate up to 100 blocks before running out of space. The default is to allocate from 1 billion to 2 billion in 10k blocks (which is the expected size of the ranges container images will use once user namespaces are started).",
	"mcsAllocatorRange":   "MCSAllocatorRange defines the range of MCS categories that will be assigned to namespaces. The format is \"<prefix>/<numberOfLabels>[,<maxCategory>]\". The default is \"s0/2\" and will allocate from c0 -> c1023, which means a total of 535k labels are available (1024 choose 2 ~ 535k). If this value is changed after startup, new projects may receive labels that are already allocated to other projects. Prefix may be any valid SELinux set of terms (including user, role, and type), although leaving them as the default will allow the server to set them automatically.\n\nExamples: * s0:/2     - Allocate labels from s0:c0,c0 to s0:c511,c511 * s0:/2,512 - Allocate labels from s0:c0,c0,c0 to s0:c511,c511,511",
	"mcsLabelsPerProject": "MCSLabelsPerProject defines the number of labels that should be reserved per project. The default is 5 to match the default UID and MCS ranges (100k namespaces, 535k/5 labels).",
}

func (SecurityAllocator) SwaggerDoc() map[string]string {
	return map_SecurityAllocator
}

var map_ServiceAccountControllerConfig = map[string]string{
	"managedNames": "managedNames is a list of service account names that will be auto-created in every namespace. If no names are specified, the ServiceAccountsController will not be started.",
}

func (ServiceAccountControllerConfig) SwaggerDoc() map[string]string {
	return map_ServiceAccountControllerConfig
}

var map_ServiceServingCert = map[string]string{
	"":       "ServiceServingCert holds configuration for service serving cert signer which creates cert/key pairs for pods fulfilling a service to serve with.",
	"signer": "Signer holds the signing information used to automatically sign serving certificates. If this value is nil, then certs are not signed automatically.",
}

func (ServiceServingCert) SwaggerDoc() map[string]string {
	return map_ServiceServingCert
}

var map_SourceStrategyDefaultsConfig = map[string]string{
	"":            "SourceStrategyDefaultsConfig contains values that apply to builds using the source strategy.",
	"incremental": "incremental indicates if s2i build strategies should perform an incremental build or not",
}

func (SourceStrategyDefaultsConfig) SwaggerDoc() map[string]string {
	return map_SourceStrategyDefaultsConfig
}

// AUTO-GENERATED FUNCTIONS END HERE
