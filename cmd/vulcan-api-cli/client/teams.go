// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "Vulcan-API": teams Resource Client
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

// CreateTeamsPath computes a request path to the create action of teams.
func CreateTeamsPath() string {

	return fmt.Sprintf("/api/v1/teams")
}

// Create a new team.
func (c *Client) CreateTeams(ctx context.Context, path string, payload *TeamPayload) (*http.Response, error) {
	req, err := c.NewCreateTeamsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateTeamsRequest create the request corresponding to the create action endpoint of the teams resource.
func (c *Client) NewCreateTeamsRequest(ctx context.Context, path string, payload *TeamPayload) (*http.Request, error) {
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

// DeleteTeamsPath computes a request path to the delete action of teams.
func DeleteTeamsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s", param0)
}

// Delete a team.
func (c *Client) DeleteTeams(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteTeamsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteTeamsRequest create the request corresponding to the delete action endpoint of the teams resource.
func (c *Client) NewDeleteTeamsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListTeamsPath computes a request path to the list action of teams.
func ListTeamsPath() string {

	return fmt.Sprintf("/api/v1/teams")
}

// List all teams in Vulcan.
func (c *Client) ListTeams(ctx context.Context, path string, tag *string) (*http.Response, error) {
	req, err := c.NewListTeamsRequest(ctx, path, tag)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListTeamsRequest create the request corresponding to the list action endpoint of the teams resource.
func (c *Client) NewListTeamsRequest(ctx context.Context, path string, tag *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if tag != nil {
		values.Set("tag", *tag)
	}
	u.RawQuery = values.Encode()
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

// ShowTeamsPath computes a request path to the show action of teams.
func ShowTeamsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s", param0)
}

// Show information about a team.
func (c *Client) ShowTeams(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowTeamsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowTeamsRequest create the request corresponding to the show action endpoint of the teams resource.
func (c *Client) NewShowTeamsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateTeamsPath computes a request path to the update action of teams.
func UpdateTeamsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s", param0)
}

// Update information about a team.
func (c *Client) UpdateTeams(ctx context.Context, path string, payload *TeamUpdatePayload) (*http.Response, error) {
	req, err := c.NewUpdateTeamsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateTeamsRequest create the request corresponding to the update action endpoint of the teams resource.
func (c *Client) NewUpdateTeamsRequest(ctx context.Context, path string, payload *TeamUpdatePayload) (*http.Request, error) {
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