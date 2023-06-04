// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	lakeformation_sdkv1 "github.com/aws/aws-sdk-go/service/lakeformation"
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
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceDataLakeSettings,
			TypeName: "aws_lakeformation_data_lake_settings",
		},
		{
			Factory:  DataSourcePermissions,
			TypeName: "aws_lakeformation_permissions",
		},
		{
			Factory:  DataSourceResource,
			TypeName: "aws_lakeformation_resource",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDataLakeSettings,
			TypeName: "aws_lakeformation_data_lake_settings",
		},
		{
			Factory:  ResourceLFTag,
			TypeName: "aws_lakeformation_lf_tag",
		},
		{
			Factory:  ResourcePermissions,
			TypeName: "aws_lakeformation_permissions",
		},
		{
			Factory:  ResourceResource,
			TypeName: "aws_lakeformation_resource",
		},
		{
			Factory:  ResourceResourceLFTags,
			TypeName: "aws_lakeformation_resource_lf_tags",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.LakeFormation
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) (*lakeformation_sdkv1.LakeFormation, error) {
	return lakeformation_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)})), nil
}

var ServicePackage = &servicePackage{}
