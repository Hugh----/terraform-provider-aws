// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package lightsail

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	lightsail_sdkv1 "github.com/aws/aws-sdk-go/service/lightsail"
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
			Factory:  ResourceBucket,
			TypeName: "aws_lightsail_bucket",
			Name:     "Bucket",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceBucketAccessKey,
			TypeName: "aws_lightsail_bucket_access_key",
		},
		{
			Factory:  ResourceBucketResourceAccess,
			TypeName: "aws_lightsail_bucket_resource_access",
		},
		{
			Factory:  ResourceCertificate,
			TypeName: "aws_lightsail_certificate",
			Name:     "Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceContainerService,
			TypeName: "aws_lightsail_container_service",
			Name:     "Container Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceContainerServiceDeploymentVersion,
			TypeName: "aws_lightsail_container_service_deployment_version",
		},
		{
			Factory:  ResourceDatabase,
			TypeName: "aws_lightsail_database",
			Name:     "Database",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceDisk,
			TypeName: "aws_lightsail_disk",
			Name:     "Disk",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceDiskAttachment,
			TypeName: "aws_lightsail_disk_attachment",
		},
		{
			Factory:  ResourceDistribution,
			TypeName: "aws_lightsail_distribution",
			Name:     "Distribution",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceDomain,
			TypeName: "aws_lightsail_domain",
		},
		{
			Factory:  ResourceDomainEntry,
			TypeName: "aws_lightsail_domain_entry",
		},
		{
			Factory:  ResourceInstance,
			TypeName: "aws_lightsail_instance",
			Name:     "Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceInstancePublicPorts,
			TypeName: "aws_lightsail_instance_public_ports",
		},
		{
			Factory:  ResourceKeyPair,
			TypeName: "aws_lightsail_key_pair",
		},
		{
			Factory:  ResourceLoadBalancer,
			TypeName: "aws_lightsail_lb",
			Name:     "LB",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceLoadBalancerAttachment,
			TypeName: "aws_lightsail_lb_attachment",
		},
		{
			Factory:  ResourceLoadBalancerCertificate,
			TypeName: "aws_lightsail_lb_certificate",
		},
		{
			Factory:  ResourceLoadBalancerCertificateAttachment,
			TypeName: "aws_lightsail_lb_certificate_attachment",
		},
		{
			Factory:  ResourceLoadBalancerHTTPSRedirectionPolicy,
			TypeName: "aws_lightsail_lb_https_redirection_policy",
		},
		{
			Factory:  ResourceLoadBalancerStickinessPolicy,
			TypeName: "aws_lightsail_lb_stickiness_policy",
		},
		{
			Factory:  ResourceStaticIP,
			TypeName: "aws_lightsail_static_ip",
		},
		{
			Factory:  ResourceStaticIPAttachment,
			TypeName: "aws_lightsail_static_ip_attachment",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Lightsail
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session) (*lightsail_sdkv1.Lightsail, error) {
	return lightsail_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.endpoint)})), nil
}

var ServicePackage = &servicePackage{}
