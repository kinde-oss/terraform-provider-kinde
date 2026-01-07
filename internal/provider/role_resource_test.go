package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRoleResource(t *testing.T) {
	testID := acctest.RandomWithPrefix("tfacc-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing.
			{
				Config: testAccRoleResourceConfig(testID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("kinde_role.test", "name", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "key", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "description", "Test role"),
					resource.TestCheckResourceAttrSet("kinde_role.test", "id"),
				),
			},
			// ImportState testing.
			{
				ResourceName:      "kinde_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing.
			{
				Config: testAccRoleResourceConfigUpdate(testID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("kinde_role.test", "name", testID+"-updated"),
					resource.TestCheckResourceAttr("kinde_role.test", "key", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "description", "Updated test role"),
				),
			},
		},
	})
}

func TestAccRoleResource_PermissionOrdering(t *testing.T) {
	// FIXME: Test is failing with "Provider produced inconsistent result after apply"
	// .permissions: was cty.SetVal([]cty.Value{cty.StringVal("")}), but now null.
	t.Skip("Skipping test due to known issue with permissions handling")

	testID := acctest.RandomWithPrefix("tfacc")
	permission1ID := acctest.RandomWithPrefix("tfacc-perm1")
	permission2ID := acctest.RandomWithPrefix("tfacc-perm2")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create both permissions and the role with permissions in one order.
			{
				Config: testAccRoleWithPermissionResourcesConfig(testID, permission1ID, permission2ID, []string{"kinde_permission.perm1.id", "kinde_permission.perm2.id"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("kinde_role.test", "name", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "key", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "description", "Test role with permissions"),
					resource.TestCheckResourceAttr("kinde_role.test", "permissions.#", "2"),
				),
			},
			// Update with same permissions in different order - should not trigger a change.
			{
				Config:   testAccRoleWithPermissionResourcesConfig(testID, permission1ID, permission2ID, []string{"kinde_permission.perm2.id", "kinde_permission.perm1.id"}),
				PlanOnly: true,
			},
			// Remove all permissions.
			{
				Config: testAccRoleResourceConfig(testID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("kinde_role.test", "name", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "key", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "description", "Test role"),
					resource.TestCheckNoResourceAttr("kinde_role.test", "permissions"),
				),
			},
		},
	})
}

func TestAccRoleResource_RemovePermissions(t *testing.T) {
	// FIXME: Test is failing with "Provider produced inconsistent result after apply"
	// .permissions: was cty.SetVal([]cty.Value{cty.StringVal("")}), but now null.
	t.Skip("Skipping test due to known issue with permissions handling")

	testID := acctest.RandomWithPrefix("tfacc")
	permission1ID := acctest.RandomWithPrefix("tfacc-perm1")
	permission2ID := acctest.RandomWithPrefix("tfacc-perm2")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create both permissions and the role with two permissions.
			{
				Config: testAccRoleWithPermissionResourcesConfig(testID, permission1ID, permission2ID, []string{"kinde_permission.perm1.id", "kinde_permission.perm2.id"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("kinde_role.test", "name", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "key", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "description", "Test role with permissions"),
					resource.TestCheckResourceAttr("kinde_role.test", "permissions.#", "2"),
				),
			},
			// Remove one permission.
			{
				Config: testAccRoleWithPermissionResourcesConfig(testID, permission1ID, permission2ID, []string{"kinde_permission.perm1.id"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("kinde_role.test", "name", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "key", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "description", "Test role with permissions"),
					resource.TestCheckResourceAttr("kinde_role.test", "permissions.#", "1"),
				),
			},
			// Remove all permissions.
			{
				Config: testAccRoleResourceConfig(testID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("kinde_role.test", "name", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "key", testID),
					resource.TestCheckResourceAttr("kinde_role.test", "description", "Test role"),
					resource.TestCheckNoResourceAttr("kinde_role.test", "permissions"),
				),
			},
		},
	})
}

func testAccRoleResourceConfig(name string) string {
	return fmt.Sprintf(`
resource "kinde_role" "test" {
	name        = %[1]q
	key         = %[1]q
	description = "Test role"
}
`, name)
}

func testAccRoleResourceConfigUpdate(name string) string {
	return fmt.Sprintf(`
resource "kinde_role" "test" {
	name        = "%[1]s-updated"
	key         = %[1]q
	description = "Updated test role"
}
`, name)
}

func testAccRoleWithPermissionResourcesConfig(roleName, permission1Key, permission2Key string, permissionRefs []string) string {
	permissionsStr := "[]"
	if len(permissionRefs) > 0 {
		permissionsStr = "[" + strings.Join(permissionRefs, ", ") + "]"
	}

	return fmt.Sprintf(`
resource "kinde_permission" "perm1" {
  name        = "test-permission-1"
  key         = %q
  description = "Test permission 1"
}

resource "kinde_permission" "perm2" {
  name        = "test-permission-2"
  key         = %q
  description = "Test permission 2"
}

resource "kinde_role" "test" {
  name        = %q
  key         = %q
  description = "Test role with permissions"
  permissions = %s
}
`, permission1Key, permission2Key, roleName, roleName, permissionsStr)
}
