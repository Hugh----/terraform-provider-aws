// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chime_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/chimesdkvoice"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfchime "github.com/hashicorp/terraform-provider-aws/internal/service/chime"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func testAccVoiceConnector_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var voiceConnector *chimesdkvoice.VoiceConnector

	vcName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_chime_voice_connector.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, chimesdkvoice.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVoiceConnectorDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVoiceConnectorConfig_basic(vcName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVoiceConnectorExists(ctx, resourceName, voiceConnector),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("vc-%s", vcName)),
					resource.TestCheckResourceAttrSet(resourceName, "aws_region"),
					resource.TestCheckResourceAttr(resourceName, "require_encryption", "true"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVoiceConnector_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var voiceConnector *chimesdkvoice.VoiceConnector

	vcName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_chime_voice_connector.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, chimesdkvoice.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVoiceConnectorDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVoiceConnectorConfig_basic(vcName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVoiceConnectorExists(ctx, resourceName, voiceConnector),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfchime.ResourceVoiceConnector(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccVoiceConnector_update(t *testing.T) {
	ctx := acctest.Context(t)
	var voiceConnector *chimesdkvoice.VoiceConnector

	vcName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_chime_voice_connector.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, chimesdkvoice.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVoiceConnectorDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVoiceConnectorConfig_basic(vcName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVoiceConnectorExists(ctx, resourceName, voiceConnector),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("vc-%s", vcName)),
					resource.TestCheckResourceAttrSet(resourceName, "aws_region"),
					resource.TestCheckResourceAttr(resourceName, "require_encryption", "true"),
				),
			},
			{
				Config: testAccVoiceConnectorConfig_updated(vcName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "require_encryption", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVoiceConnector_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var voiceConnector *chimesdkvoice.VoiceConnector

	vcName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_chime_voice_connector.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, chimesdkvoice.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVoiceConnectorDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVoiceConnectorConfig_tags1(vcName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVoiceConnectorExists(ctx, resourceName, voiceConnector),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("vc-%s", vcName)),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVoiceConnectorConfig_tags2(vcName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVoiceConnectorExists(ctx, resourceName, voiceConnector),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("vc-%s", vcName)),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccVoiceConnectorConfig_tags1(vcName, "key2", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVoiceConnectorExists(ctx, resourceName, voiceConnector),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("vc-%s", vcName)),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func testAccVoiceConnectorConfig_basic(name string) string {
	return fmt.Sprintf(`
resource "aws_chime_voice_connector" "test" {
  name               = "vc-%s"
  require_encryption = true
}
`, name)
}

func testAccVoiceConnectorConfig_updated(name string) string {
	return fmt.Sprintf(`
resource "aws_chime_voice_connector" "test" {
  name               = "vc-%s"
  require_encryption = false
}
`, name)
}

func testAccVoiceConnectorConfig_tags1(name, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_chime_voice_connector" "test" {
  name               = "vc-%s"
  require_encryption = true

  tags = {
    %[2]q = %[3]q
  }
}
`, name, tagKey1, tagValue1)
}

func testAccVoiceConnectorConfig_tags2(name, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_chime_voice_connector" "test" {
  name               = "vc-%s"
  require_encryption = true

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, name, tagKey1, tagValue1, tagKey2, tagValue2)
}

func testAccCheckVoiceConnectorExists(ctx context.Context, name string, vc *chimesdkvoice.VoiceConnector) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Chime voice connector ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).ChimeSDKVoiceConn(ctx)

		resp, err := tfchime.FindVoiceConnectorResourceWithRetry(ctx, false, func() (*chimesdkvoice.VoiceConnector, error) {
			return tfchime.FindVoiceConnectorByID(ctx, conn, rs.Primary.ID)
		})

		if err != nil {
			return err
		}

		vc = resp

		return nil
	}
}

func testAccCheckVoiceConnectorDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_chime_voice_connector" {
				continue
			}
			conn := acctest.Provider.Meta().(*conns.AWSClient).ChimeSDKVoiceConn(ctx)

			_, err := tfchime.FindVoiceConnectorResourceWithRetry(ctx, false, func() (*chimesdkvoice.VoiceConnector, error) {
				return tfchime.FindVoiceConnectorByID(ctx, conn, rs.Primary.ID)
			})

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("voice connector still exists: (%s)", rs.Primary.ID)
		}

		return nil
	}
}

func testAccPreCheck(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).ChimeSDKVoiceConn(ctx)

	input := &chimesdkvoice.ListVoiceConnectorsInput{}

	_, err := conn.ListVoiceConnectorsWithContext(ctx, input)

	if acctest.PreCheckSkipError(err) {
		t.Skipf("skipping acceptance testing: %s", err)
	}

	if err != nil {
		t.Fatalf("unexpected PreCheck error: %s", err)
	}
}
