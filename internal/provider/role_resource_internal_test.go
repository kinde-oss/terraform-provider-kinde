// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/nxt-fwd/kinde-go"
)

// TestGetRoleUsesAllPermissionPages proves that RoleResource.getRole follows
// next_token when reading a role's permissions, instead of relying on the
// kinde-go SDK's roles.Client.Get, which only reads a single, non-paginated
// page of /api/v1/roles/{id}/permissions and silently truncates roles with
// more permissions than fit on one page.
func TestGetRoleUsesAllPermissionPages(t *testing.T) {
	ctx := context.Background()
	var permissionRequests []string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.URL.Path {
		case "/oauth2/token":
			_ = json.NewEncoder(w).Encode(map[string]any{
				"access_token": "test-token",
				"expires_in":   3600,
				"token_type":   "bearer",
			})
		case "/api/v1/roles/role_1":
			_ = json.NewEncoder(w).Encode(map[string]any{
				"role": map[string]any{
					"id":          "role_1",
					"name":        "Role",
					"key":         "role",
					"description": "Role description",
				},
			})
		case "/api/v1/roles/role_1/permissions":
			permissionRequests = append(permissionRequests, r.URL.RawQuery)
			if r.URL.Query().Get("next_token") == "" {
				_ = json.NewEncoder(w).Encode(map[string]any{
					"next_token": "page-2",
					"permissions": []map[string]string{
						{"id": "perm_01"},
						{"id": "perm_02"},
						{"id": "perm_03"},
						{"id": "perm_04"},
						{"id": "perm_05"},
						{"id": "perm_06"},
						{"id": "perm_07"},
						{"id": "perm_08"},
						{"id": "perm_09"},
						{"id": "perm_10"},
					},
				})
				return
			}

			_ = json.NewEncoder(w).Encode(map[string]any{
				"permissions": []map[string]string{
					{"id": "perm_11"},
				},
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := kinde.New(ctx, kinde.NewClientOptions().
		WithDomain(server.URL).
		WithAudience("audience").
		WithClientID("client-id").
		WithClientSecret("client-secret"))

	roleResource := RoleResource{client: client.Roles}
	role, err := roleResource.getRole(ctx, "role_1")
	if err != nil {
		t.Fatalf("get role: %s", err)
	}

	if got := len(role.Permissions); got != 11 {
		t.Fatalf("expected 11 permissions across pages, got %d: %#v", got, role.Permissions)
	}

	if len(permissionRequests) != 2 {
		t.Fatalf("expected 2 permission requests, got %d: %#v", len(permissionRequests), permissionRequests)
	}

	wantQueries := []url.Values{
		{"page_size": []string{"100"}},
		{"page_size": []string{"100"}, "next_token": []string{"page-2"}},
	}
	for i, wantQuery := range wantQueries {
		gotQuery, err := url.ParseQuery(permissionRequests[i])
		if err != nil {
			t.Fatalf("parse permission request query %q: %s", permissionRequests[i], err)
		}
		if !reflect.DeepEqual(gotQuery, wantQuery) {
			t.Fatalf("unexpected permission request query %d\nwant: %#v\n got: %#v", i, wantQuery, gotQuery)
		}
	}
}
