// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package codebuild

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	codebuild_sdkv1 "github.com/aws/aws-sdk-go/service/codebuild"
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
			Factory:  ResourceProject,
			TypeName: "aws_codebuild_project",
			Name:     "Project",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceReportGroup,
			TypeName: "aws_codebuild_report_group",
			Name:     "Report Group",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceResourcePolicy,
			TypeName: "aws_codebuild_resource_policy",
		},
		{
			Factory:  ResourceSourceCredential,
			TypeName: "aws_codebuild_source_credential",
		},
		{
			Factory:  ResourceWebhook,
			TypeName: "aws_codebuild_webhook",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CodeBuild
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) (*codebuild_sdkv1.CodeBuild, error) {
	return codebuild_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)})), nil
}

var ServicePackage = &servicePackage{}
