// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// Rules is a list of rules generated from Magic Modules
var Rules = []tflint.Rule{
	NewGoogleAccessContextManagerServicePerimeterInvalidPerimeterTypeRule(),
	NewGoogleActiveDirectoryDomainTrustInvalidTrustDirectionRule(),
	NewGoogleActiveDirectoryDomainTrustInvalidTrustTypeRule(),
	NewGoogleApigeeInstanceInvalidPeeringCidrRangeRule(),
	NewGoogleApigeeOrganizationInvalidRuntimeTypeRule(),
	NewGoogleAppEngineDomainMappingInvalidOverrideStrategyRule(),
	NewGoogleAppEngineFirewallRuleInvalidActionRule(),
	NewGoogleAppEngineFlexibleAppVersionInvalidServingStatusRule(),
	NewGoogleBigqueryRoutineInvalidDeterminismLevelRule(),
	NewGoogleBigqueryRoutineInvalidLanguageRule(),
	NewGoogleBigqueryRoutineInvalidRoutineTypeRule(),
	NewGoogleBinaryAuthorizationPolicyInvalidGlobalPolicyEvaluationModeRule(),
	NewGoogleCloudAssetFolderFeedInvalidContentTypeRule(),
	NewGoogleCloudAssetOrganizationFeedInvalidContentTypeRule(),
	NewGoogleCloudAssetProjectFeedInvalidContentTypeRule(),
	NewGoogleCloudIdentityGroupInvalidInitialGroupConfigRule(),
	NewGoogleCloudiotDeviceInvalidLogLevelRule(),
	NewGoogleCloudiotRegistryInvalidLogLevelRule(),
	NewGoogleComputeAddressInvalidAddressTypeRule(),
	NewGoogleComputeAddressInvalidNameRule(),
	NewGoogleComputeAddressInvalidNetworkTierRule(),
	NewGoogleComputeBackendBucketInvalidNameRule(),
	NewGoogleComputeBackendBucketSignedUrlKeyInvalidNameRule(),
	NewGoogleComputeBackendServiceInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeBackendServiceInvalidLocalityLbPolicyRule(),
	NewGoogleComputeBackendServiceInvalidProtocolRule(),
	NewGoogleComputeBackendServiceInvalidSessionAffinityRule(),
	NewGoogleComputeBackendServiceSignedUrlKeyInvalidNameRule(),
	NewGoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule(),
	NewGoogleComputeFirewallInvalidDirectionRule(),
	NewGoogleComputeForwardingRuleInvalidIpProtocolRule(),
	NewGoogleComputeForwardingRuleInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeForwardingRuleInvalidNetworkTierRule(),
	NewGoogleComputeGlobalAddressInvalidAddressTypeRule(),
	NewGoogleComputeGlobalAddressInvalidIpVersionRule(),
	NewGoogleComputeGlobalForwardingRuleInvalidIpProtocolRule(),
	NewGoogleComputeGlobalForwardingRuleInvalidIpVersionRule(),
	NewGoogleComputeGlobalForwardingRuleInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeGlobalNetworkEndpointGroupInvalidNetworkEndpointTypeRule(),
	NewGoogleComputeInterconnectAttachmentInvalidBandwidthRule(),
	NewGoogleComputeInterconnectAttachmentInvalidEncryptionRule(),
	NewGoogleComputeInterconnectAttachmentInvalidNameRule(),
	NewGoogleComputeInterconnectAttachmentInvalidTypeRule(),
	NewGoogleComputeManagedSslCertificateInvalidTypeRule(),
	NewGoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule(),
	NewGoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule(),
	NewGoogleComputeRegionBackendServiceInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule(),
	NewGoogleComputeRegionBackendServiceInvalidProtocolRule(),
	NewGoogleComputeRegionBackendServiceInvalidSessionAffinityRule(),
	NewGoogleComputeRegionNetworkEndpointGroupInvalidNetworkEndpointTypeRule(),
	NewGoogleComputeRouteInvalidNameRule(),
	NewGoogleComputeRouterNatInvalidNatIpAllocateOptionRule(),
	NewGoogleComputeRouterNatInvalidSourceSubnetworkIpRangesToNatRule(),
	NewGoogleComputeRouterPeerInvalidAdvertiseModeRule(),
	NewGoogleComputeSslPolicyInvalidMinTlsVersionRule(),
	NewGoogleComputeSslPolicyInvalidProfileRule(),
	NewGoogleComputeSubnetworkInvalidRoleRule(),
	NewGoogleComputeTargetHttpsProxyInvalidQuicOverrideRule(),
	NewGoogleComputeTargetInstanceInvalidNatPolicyRule(),
	NewGoogleComputeTargetSslProxyInvalidProxyHeaderRule(),
	NewGoogleComputeTargetTcpProxyInvalidProxyHeaderRule(),
	NewGoogleDataCatalogEntryGroupInvalidEntryGroupIdRule(),
	NewGoogleDataCatalogEntryInvalidTypeRule(),
	NewGoogleDataCatalogEntryInvalidUserSpecifiedSystemRule(),
	NewGoogleDataCatalogEntryInvalidUserSpecifiedTypeRule(),
	NewGoogleDataCatalogTagTemplateInvalidTagTemplateIdRule(),
	NewGoogleDataLossPreventionJobTriggerInvalidStatusRule(),
	NewGoogleDatastoreIndexInvalidAncestorRule(),
	NewGoogleDeploymentManagerDeploymentInvalidCreatePolicyRule(),
	NewGoogleDeploymentManagerDeploymentInvalidDeletePolicyRule(),
	NewGoogleDialogflowAgentInvalidApiVersionRule(),
	NewGoogleDialogflowAgentInvalidMatchModeRule(),
	NewGoogleDialogflowAgentInvalidTierRule(),
	NewGoogleDialogflowCxEntityTypeInvalidAutoExpansionModeRule(),
	NewGoogleDialogflowCxEntityTypeInvalidKindRule(),
	NewGoogleDialogflowEntityTypeInvalidKindRule(),
	NewGoogleDialogflowIntentInvalidWebhookStateRule(),
	NewGoogleDnsManagedZoneInvalidVisibilityRule(),
	NewGoogleDnsRecordSetInvalidTypeRule(),
	NewGoogleFilestoreInstanceInvalidTierRule(),
	NewGoogleFirestoreIndexInvalidQueryScopeRule(),
	NewGoogleHealthcareFhirStoreInvalidVersionRule(),
	NewGoogleKmsCryptoKeyInvalidPurposeRule(),
	NewGoogleKmsKeyRingImportJobInvalidImportMethodRule(),
	NewGoogleKmsKeyRingImportJobInvalidProtectionLevelRule(),
	NewGoogleMemcacheInstanceInvalidMemcacheVersionRule(),
	NewGoogleMonitoringAlertPolicyInvalidCombinerRule(),
	NewGoogleMonitoringCustomServiceInvalidServiceIdRule(),
	NewGoogleMonitoringMetricDescriptorInvalidLaunchStageRule(),
	NewGoogleMonitoringMetricDescriptorInvalidMetricKindRule(),
	NewGoogleMonitoringMetricDescriptorInvalidValueTypeRule(),
	NewGoogleMonitoringSloInvalidCalendarPeriodRule(),
	NewGoogleMonitoringSloInvalidSloIdRule(),
	NewGoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule(),
	NewGoogleNotebooksInstanceInvalidBootDiskTypeRule(),
	NewGoogleNotebooksInstanceInvalidDataDiskTypeRule(),
	NewGoogleNotebooksInstanceInvalidDiskEncryptionRule(),
	NewGoogleOsConfigPatchDeploymentInvalidPatchDeploymentIdRule(),
	NewGooglePrivatecaCaPoolInvalidTierRule(),
	NewGooglePrivatecaCertificateAuthorityInvalidTypeRule(),
	NewGooglePubsubSchemaInvalidTypeRule(),
	NewGoogleRedisInstanceInvalidConnectModeRule(),
	NewGoogleRedisInstanceInvalidNameRule(),
	NewGoogleRedisInstanceInvalidTierRule(),
	NewGoogleRedisInstanceInvalidTransitEncryptionModeRule(),
	NewGoogleSccSourceInvalidDisplayNameRule(),
	NewGoogleSpannerDatabaseInvalidNameRule(),
	NewGoogleSpannerInstanceInvalidNameRule(),
	NewGoogleSqlSourceRepresentationInstanceInvalidDatabaseVersionRule(),
	NewGoogleStorageBucketAccessControlInvalidRoleRule(),
	NewGoogleStorageDefaultObjectAccessControlInvalidRoleRule(),
	NewGoogleStorageHmacKeyInvalidStateRule(),
	NewGoogleStorageObjectAccessControlInvalidRoleRule(),
}
