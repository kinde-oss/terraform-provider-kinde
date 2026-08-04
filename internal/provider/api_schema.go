// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nxt-fwd/kinde-go/api/apis"
)

type APIResourceModel struct {
	ID              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Audience        types.String `tfsdk:"audience"`
	IsManagementAPI types.Bool   `tfsdk:"is_management_api"`
}

func expandAPIResourceModel(model APIResourceModel) *apis.API {
	api := &apis.API{
		ID:       model.ID.ValueString(),
		Name:     model.Name.ValueString(),
		Audience: model.Audience.ValueString(),
	}

	if !model.IsManagementAPI.IsNull() && !model.IsManagementAPI.IsUnknown() {
		api.IsManagementAPI = model.IsManagementAPI.ValueBool()
	}

	return api
}

func flattenAPIResource(resource *apis.API) APIResourceModel {
	model := APIResourceModel{
		ID:       types.StringValue(resource.ID),
		Name:     types.StringValue(resource.Name),
		Audience: types.StringValue(resource.Audience),
	}

	model.IsManagementAPI = types.BoolValue(resource.IsManagementAPI)

	return model
}

type APIDataSourceModel struct {
	ID       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Audience types.String `tfsdk:"audience"`
}

func expandAPIDataSourceModel(model APIDataSourceModel) *apis.API {
	return &apis.API{
		ID:       model.ID.ValueString(),
		Name:     model.Name.ValueString(),
		Audience: model.Audience.ValueString(),
	}
}

func flattenAPIDataSource(resource *apis.API) APIDataSourceModel {
	return APIDataSourceModel{
		ID:       types.StringValue(resource.ID),
		Name:     types.StringValue(resource.Name),
		Audience: types.StringValue(resource.Audience),
	}
}
