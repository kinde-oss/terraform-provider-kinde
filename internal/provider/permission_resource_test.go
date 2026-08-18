// Copyright (c) Kinde Australia Pty Ltd
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPermissionResource(t *testing.T) {
	// Permission keys must be unique within a Kinde business.
	suffix := acctest.RandInt()
	name := fmt.Sprintf("tfacc-permission-%d", suffix)
	key := fmt.Sprintf("tfacc_permission_%d", suffix)
	updatedName := fmt.Sprintf("tfacc-permission-updated-%d", suffix)
	updatedKey := fmt.Sprintf("tfacc_permission_updated_%d", suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccPermissionResourceConfig(name, key, "Test permission description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("kinde_permission.test", "name", name),
					resource.TestCheckResourceAttr("kinde_permission.test", "key", key),
					resource.TestCheckResourceAttr("kinde_permission.test", "description", "Test permission description"),
					resource.TestCheckResourceAttrSet("kinde_permission.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "kinde_permission.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccPermissionResourceConfig(updatedName, updatedKey, "Updated test permission description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("kinde_permission.test", "name", updatedName),
					resource.TestCheckResourceAttr("kinde_permission.test", "key", updatedKey),
					resource.TestCheckResourceAttr("kinde_permission.test", "description", "Updated test permission description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccPermissionResourceConfig(name string, key string, description string) string {
	return fmt.Sprintf(`
resource "kinde_permission" "test" {
  name        = %[1]q
  key         = %[2]q
  description = %[3]q
}
`, name, key, description)
}
