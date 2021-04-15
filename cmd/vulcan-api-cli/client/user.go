// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "Vulcan-API": user Resource Client
//
// Command:
// $ main

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateUserPath computes a request path to the create action of user.
func CreateUserPath() string {

	return fmt.Sprintf("/api/v1/users")
}

// Create user
func (c *Client) CreateUser(ctx context.Context, path string, payload *UserPayload) (*http.Response, error) {
	req, err := c.NewCreateUserRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateUserRequest create the request corresponding to the create action endpoint of the user resource.
func (c *Client) NewCreateUserRequest(ctx context.Context, path string, payload *UserPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// DeleteUserPath computes a request path to the delete action of user.
func DeleteUserPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/api/v1/users/%s", param0)
}

// Remove user
func (c *Client) DeleteUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteUserRequest create the request corresponding to the delete action endpoint of the user resource.
func (c *Client) NewDeleteUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ListUserPath computes a request path to the list action of user.
func ListUserPath() string {

	return fmt.Sprintf("/api/v1/users")
}

// List all users
func (c *Client) ListUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListUserRequest create the request corresponding to the list action endpoint of the user resource.
func (c *Client) NewListUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ListTeamsUserPath computes a request path to the list-teams action of user.
func ListTeamsUserPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/api/v1/users/%s/teams", param0)
}

// List all teams for an user.
func (c *Client) ListTeamsUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListTeamsUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListTeamsUserRequest create the request corresponding to the list-teams action endpoint of the user resource.
func (c *Client) NewListTeamsUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ProfileUserPath computes a request path to the profile action of user.
func ProfileUserPath() string {

	return fmt.Sprintf("/api/v1/profile")
}

// Show profile information for the current authenticated user based on the key used to make the request.
func (c *Client) ProfileUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewProfileUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewProfileUserRequest create the request corresponding to the profile action endpoint of the user resource.
func (c *Client) NewProfileUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ShowUserPath computes a request path to the show action of user.
func ShowUserPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/api/v1/users/%s", param0)
}

// Describe user
func (c *Client) ShowUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowUserRequest create the request corresponding to the show action endpoint of the user resource.
func (c *Client) NewShowUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// UpdateUserPath computes a request path to the update action of user.
func UpdateUserPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/api/v1/users/%s", param0)
}

// Update user
func (c *Client) UpdateUser(ctx context.Context, path string, payload *UserUpdatePayload) (*http.Response, error) {
	req, err := c.NewUpdateUserRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateUserRequest create the request corresponding to the update action endpoint of the user resource.
func (c *Client) NewUpdateUserRequest(ctx context.Context, path string, payload *UserUpdatePayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}