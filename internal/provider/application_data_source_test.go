// Copyright (c) Kinde Australia Pty Ltd
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccApplicationDataSource(t *testing.T) {
	testID := acctest.RandomWithPrefix("tfacc-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "kinde_application" "test" {
					name = %[1]q
					type = "reg"
				}

				data "kinde_application" "test" {
					id = kinde_application.test.id
				}
				`, testID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.kinde_application.test", "id"),
					resource.TestCheckResourceAttr("data.kinde_application.test", "name", testID),
					resource.TestCheckResourceAttr("data.kinde_application.test", "type", "reg"),
				),
			},
		},
	})
}
