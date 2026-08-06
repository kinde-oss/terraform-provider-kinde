// Copyright (c) Kinde Australia Pty Ltd
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"testing"

	frameworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func TestKindeProviderSchemaClientSecretSensitive(t *testing.T) {
	p := &KindeProvider{}
	var resp frameworkprovider.SchemaResponse

	p.Schema(context.Background(), frameworkprovider.SchemaRequest{}, &resp)

	attribute, ok := resp.Schema.Attributes["client_secret"].(providerschema.StringAttribute)
	if !ok {
		t.Fatalf("expected client_secret to be a string attribute, got %T", resp.Schema.Attributes["client_secret"])
	}

	if !attribute.Sensitive {
		t.Fatal("expected provider client_secret attribute to be marked sensitive")
	}
}
