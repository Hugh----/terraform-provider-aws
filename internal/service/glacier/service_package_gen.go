// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package glacier

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	glacier_sdkv1 "github.com/aws/aws-sdk-go/service/glacier"
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
			Factory:  ResourceVault,
			TypeName: "aws_glacier_vault",
			Name:     "Vault",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceVaultLock,
			TypeName: "aws_glacier_vault_lock",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Glacier
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) (*glacier_sdkv1.Glacier, error) {
	return glacier_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)})), nil
}

var ServicePackage = &servicePackage{}
