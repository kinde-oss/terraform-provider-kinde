// Copyright (c) Kinde Australia Pty Ltd
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OrganizationUserResourceModel struct {
	ID               types.String `tfsdk:"id"`
	OrganizationCode types.String `tfsdk:"organization_code"`
	UserID           types.String `tfsdk:"user_id"`
	Roles            types.List   `tfsdk:"roles"`
}
