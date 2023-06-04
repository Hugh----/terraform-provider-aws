// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package transcribe

import (
	"context"
	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	transcribe_sdkv2 "github.com/aws/aws-sdk-go-v2/service/transcribe"
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
			Factory:  ResourceLanguageModel,
			TypeName: "aws_transcribe_language_model",
			Name:     "Language Model",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceMedicalVocabulary,
			TypeName: "aws_transcribe_medical_vocabulary",
			Name:     "Medical Vocabulary",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceVocabulary,
			TypeName: "aws_transcribe_vocabulary",
			Name:     "Vocabulary",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceVocabularyFilter,
			TypeName: "aws_transcribe_vocabulary_filter",
			Name:     "Vocabulary Filter",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Transcribe
}

func (p *servicePackage) SetEndpoint(endpoint string) {
	p.endpoint = endpoint
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, cfg aws_sdkv2.Config) (*transcribe_sdkv2.Client, error) {
	return transcribe_sdkv2.NewFromConfig(cfg, func(o *transcribe_sdkv2.Options) {
		if p.endpoint != "" {
			o.EndpointResolver = transcribe_sdkv2.EndpointResolverFromURL(p.endpoint)
		}
	}), nil
}

var ServicePackage = &servicePackage{}
