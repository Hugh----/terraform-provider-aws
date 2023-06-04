// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package configservice

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	configservice_sdkv1 "github.com/aws/aws-sdk-go/service/configservice"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct {
	endpoint string
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAggregateAuthorization,
			TypeName: "aws_config_aggregate_authorization",
			Name:     "Aggregate Authorization",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceConfigRule,
			TypeName: "aws_config_config_rule",
			Name:     "Config Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceConfigurationAggregator,
			TypeName: "aws_config_configuration_aggregator",
			Name:     "Configuration Aggregator",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceConfigurationRecorder,
			TypeName: "aws_config_configuration_recorder",
		},
		{
			Factory:  ResourceConfigurationRecorderStatus,
			TypeName: "aws_config_configuration_recorder_status",
		},
		{
			Factory:  ResourceConformancePack,
			TypeName: "aws_config_conformance_pack",
		},
		{
			Factory:  ResourceDeliveryChannel,
			TypeName: "aws_config_delivery_channel",
		},
		{
			Factory:  ResourceOrganizationConformancePack,
			TypeName: "aws_config_organization_conformance_pack",
		},
		{
			Factory:  ResourceOrganizationCustomPolicyRule,
			TypeName: "aws_config_organization_custom_policy_rule",
		},
		{
			Factory:  ResourceOrganizationCustomRule,
			TypeName: "aws_config_organization_custom_rule",
		},
		{
			Factory:  ResourceOrganizationManagedRule,
			TypeName: "aws_config_organization_managed_rule",
		},
		{
			Factory:  ResourceRemediationConfiguration,
			TypeName: "aws_config_remediation_configuration",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ConfigService
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) (*configservice_sdkv1.ConfigService, error) {
	return configservice_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)})), nil
}

var ServicePackage = &servicePackage{}
