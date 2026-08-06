// Copyright (c) Kinde Australia Pty Ltd
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"
	"testing"
)

func TestDiffPermissions(t *testing.T) {
	toRemove, toAdd := diffPermissions(
		[]string{"perm-b", "", "perm-a", "perm-b"},
		[]string{"perm-c", "", "perm-b", "perm-c"},
	)

	if !reflect.DeepEqual(toRemove, []string{"perm-c"}) {
		t.Fatalf("expected permissions to remove %v, got %v", []string{"perm-c"}, toRemove)
	}

	if !reflect.DeepEqual(toAdd, []string{"perm-a"}) {
		t.Fatalf("expected permissions to add %v, got %v", []string{"perm-a"}, toAdd)
	}
}

func TestDiffPermissionsNoChanges(t *testing.T) {
	toRemove, toAdd := diffPermissions(
		[]string{"perm-b", "perm-a"},
		[]string{"perm-a", "perm-b"},
	)

	if len(toRemove) != 0 {
		t.Fatalf("expected no permissions to remove, got %v", toRemove)
	}

	if len(toAdd) != 0 {
		t.Fatalf("expected no permissions to add, got %v", toAdd)
	}
}
