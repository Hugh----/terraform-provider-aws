// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package cloudsearch

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	cloudsearch_sdkv1 "github.com/aws/aws-sdk-go/service/cloudsearch"
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
			Factory:  ResourceDomain,
			TypeName: "aws_cloudsearch_domain",
		},
		{
			Factory:  ResourceDomainServiceAccessPolicy,
			TypeName: "aws_cloudsearch_domain_service_access_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CloudSearch
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) (*cloudsearch_sdkv1.CloudSearch, error) {
	return cloudsearch_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)})), nil
}

var ServicePackage = &servicePackage{}
