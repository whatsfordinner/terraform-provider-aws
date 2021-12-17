package backup_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/backup"
	"github.com/aws/aws-sdk-go/service/fsx"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)

func TestAccBackupRegionSettings_basic(t *testing.T) {
	var settings backup.DescribeRegionSettingsOutput
	resourceName := "aws_backup_region_settings.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(fsx.EndpointsID, t)
			testAccPreCheck(t)
		},
		ErrorCheck:   acctest.ErrorCheck(t, backup.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccBackupRegionSettingsConfig1(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRegionSettingsExists(&settings),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.%", "11"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.Aurora", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.DocumentDB", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.DynamoDB", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.EBS", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.EC2", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.EFS", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.FSx", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.Neptune", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.RDS", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.Storage Gateway", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.VirtualMachine", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_management_preference.%", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "resource_type_management_preference.DynamoDB"),
					resource.TestCheckResourceAttrSet(resourceName, "resource_type_management_preference.EFS"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBackupRegionSettingsConfig2(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRegionSettingsExists(&settings),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.%", "11"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.Aurora", "false"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.DocumentDB", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.DynamoDB", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.EBS", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.EC2", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.EFS", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.FSx", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.Neptune", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.RDS", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.Storage Gateway", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.VirtualMachine", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_management_preference.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_management_preference.DynamoDB", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_management_preference.EFS", "true"),
				),
			},
			{
				Config: testAccBackupRegionSettingsConfig3(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRegionSettingsExists(&settings),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.%", "11"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.Aurora", "false"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.DocumentDB", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.DynamoDB", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.EBS", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.EC2", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.EFS", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.FSx", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.Neptune", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.RDS", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.Storage Gateway", "true"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_opt_in_preference.VirtualMachine", "false"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_management_preference.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_management_preference.DynamoDB", "false"),
					resource.TestCheckResourceAttr(resourceName, "resource_type_management_preference.EFS", "true"),
				),
			},
		},
	})
}

func testAccCheckRegionSettingsExists(settings *backup.DescribeRegionSettingsOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		conn := acctest.Provider.Meta().(*conns.AWSClient).BackupConn
		resp, err := conn.DescribeRegionSettings(&backup.DescribeRegionSettingsInput{})
		if err != nil {
			return err
		}

		*settings = *resp

		return nil
	}
}

func testAccBackupRegionSettingsConfig1() string {
	return `
resource "aws_backup_region_settings" "test" {
  resource_type_opt_in_preference = {
    "Aurora"          = true
    "DocumentDB"      = true
    "DynamoDB"        = true
    "EBS"             = true
    "EC2"             = true
    "EFS"             = true
    "FSx"             = true
    "Neptune"         = true
    "RDS"             = true
    "Storage Gateway" = true
    "VirtualMachine"  = true
  }
}
`
}

func testAccBackupRegionSettingsConfig2() string {
	return `
resource "aws_backup_region_settings" "test" {
  resource_type_opt_in_preference = {
    "Aurora"          = false
    "DocumentDB"      = true
    "DynamoDB"        = true
    "EBS"             = true
    "EC2"             = true
    "EFS"             = true
    "FSx"             = true
    "Neptune"         = true
    "RDS"             = true
    "Storage Gateway" = true
    "VirtualMachine"  = true
  }

  resource_type_management_preference = {
    "DynamoDB" = true
    "EFS"      = true
  }
}
`
}

func testAccBackupRegionSettingsConfig3() string {
	return `
resource "aws_backup_region_settings" "test" {
  resource_type_opt_in_preference = {
    "Aurora"          = false
    "DocumentDB"      = true
    "DynamoDB"        = true
    "EBS"             = true
    "EC2"             = true
    "EFS"             = true
    "FSx"             = true
    "Neptune"         = true
    "RDS"             = true
    "Storage Gateway" = true
    "VirtualMachine"  = false
  }

  resource_type_management_preference = {
    "DynamoDB" = false
    "EFS"      = true
  }
}
`
}
