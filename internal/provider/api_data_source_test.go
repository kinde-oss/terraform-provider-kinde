// Copyright (c) Kinde Australia Pty Ltd
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAPIDataSource(t *testing.T) {
	// Audiences must be unique within a Kinde business.
	testID := acctest.RandomWithPrefix("tfacc-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAPIDataSourceConfig(testID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.kinde_api.test", "name", testID),
					resource.TestCheckResourceAttr("data.kinde_api.test", "audience", testID),
				),
			},
		},
	})
}

func testAccAPIDataSourceConfig(testID string) string {
	return fmt.Sprintf(`
resource "kinde_api" "test" {
	name     = %[1]q
	audience = %[1]q
}

data "kinde_api" "test" {
	id = kinde_api.test.id
}
`, testID)
}
