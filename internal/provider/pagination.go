// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"
	"net/url"
)

// kindeRequester is the subset of the kinde-go client interface needed to
// walk a paginated Kinde Management API endpoint.
type kindeRequester interface {
	NewRequest(ctx context.Context, method, path string, query url.Values, payload any) (*http.Request, error)
	DoRequest(req *http.Request, result any) error
}

// kindePage describes a single page of a paginated Kinde Management API list
// response.
type kindePage[T any] interface {
	getData() []T
	getNextToken() string
}

// getAllPages walks every page of a paginated Kinde Management API endpoint,
// following next_token until the API stops returning one. The Kinde
// Management API silently returns only the first page when next_token is
// not followed, so callers must not assume a single request is sufficient
// for any listing endpoint.
func getAllPages[T any, P kindePage[T]](ctx context.Context, client kindeRequester, endpoint string, query url.Values) ([]T, error) {
	// Copy the caller's query so we don't mutate a map they might reuse.
	localQuery := url.Values{}
	for key, values := range query {
		localQuery[key] = append([]string(nil), values...)
	}
	if localQuery.Get("page_size") == "" {
		localQuery.Set("page_size", "100")
	}

	var results []T
	for {
		req, err := client.NewRequest(ctx, http.MethodGet, endpoint, localQuery, nil)
		if err != nil {
			return nil, err
		}

		var response P
		if err := client.DoRequest(req, &response); err != nil {
			return nil, err
		}

		results = append(results, response.getData()...)

		nextToken := response.getNextToken()
		if nextToken == "" {
			break
		}
		localQuery.Set("next_token", nextToken)
	}

	return results, nil
}
