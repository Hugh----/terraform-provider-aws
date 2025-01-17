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
)

func testAccVoiceConnectorLogging_basic(t *testing.T) {
	ctx := acctest.Context(t)
	name := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_chime_voice_connector_logging.test"

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
				Config: testAccVoiceConnectorLoggingConfig_basic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVoiceConnectorLoggingExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "enable_sip_logs", "true"),
					resource.TestCheckResourceAttr(resourceName, "enable_media_metric_logs", "true"),
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

func testAccVoiceConnectorLogging_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	name := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_chime_voice_connector_logging.test"

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
				Config: testAccVoiceConnectorLoggingConfig_basic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVoiceConnectorLoggingExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfchime.ResourceVoiceConnectorLogging(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccVoiceConnectorLogging_update(t *testing.T) {
	ctx := acctest.Context(t)
	name := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_chime_voice_connector_logging.test"

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
				Config: testAccVoiceConnectorLoggingConfig_basic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVoiceConnectorLoggingExists(ctx, resourceName),
				),
			},
			{
				Config: testAccVoiceConnectorLoggingConfig_updated(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVoiceConnectorLoggingExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "enable_sip_logs", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_media_metric_logs", "false"),
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

func testAccVoiceConnectorLoggingConfig_basic(name string) string {
	return fmt.Sprintf(`
resource "aws_chime_voice_connector" "chime" {
  name               = "vc-%[1]s"
  require_encryption = true
}

resource "aws_chime_voice_connector_logging" "test" {
  voice_connector_id       = aws_chime_voice_connector.chime.id
  enable_sip_logs          = true
  enable_media_metric_logs = true
}
`, name)
}

func testAccVoiceConnectorLoggingConfig_updated(name string) string {
	return fmt.Sprintf(`
resource "aws_chime_voice_connector" "chime" {
  name               = "vc-%[1]s"
  require_encryption = true
}

resource "aws_chime_voice_connector_logging" "test" {
  voice_connector_id       = aws_chime_voice_connector.chime.id
  enable_sip_logs          = false
  enable_media_metric_logs = false
}
`, name)
}

func testAccCheckVoiceConnectorLoggingExists(ctx context.Context, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Chime Voice Connector logging ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).ChimeSDKVoiceConn(ctx)

		_, err := tfchime.FindVoiceConnectorResourceWithRetry(ctx, false, func() (*chimesdkvoice.LoggingConfiguration, error) {
			return tfchime.FindVoiceConnectorLoggingByID(ctx, conn, rs.Primary.ID)
		})

		return err
	}
}
