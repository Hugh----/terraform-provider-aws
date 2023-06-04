// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package scheduler

import (
	"context"
	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	scheduler_sdkv2 "github.com/aws/aws-sdk-go-v2/service/scheduler"
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
			Factory:  resourceSchedule,
			TypeName: "aws_scheduler_schedule",
		},
		{
			Factory:  ResourceScheduleGroup,
			TypeName: "aws_scheduler_schedule_group",
			Name:     "Schedule Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Scheduler
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, cfg aws_sdkv2.Config) (*scheduler_sdkv2.Client, error) {
	return scheduler_sdkv2.NewFromConfig(cfg, func(o *scheduler_sdkv2.Options) {
		if p.endpoint != "" {
			o.EndpointResolver = scheduler_sdkv2.EndpointResolverFromURL(p.endpoint)
		}
	}), nil
}

var ServicePackage = &servicePackage{}
