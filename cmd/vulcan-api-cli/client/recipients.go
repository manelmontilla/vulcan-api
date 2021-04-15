// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "Vulcan-API": recipients Resource Client
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

// ListRecipientsPath computes a request path to the list action of recipients.
func ListRecipientsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s/recipients", param0)
}

// List all recipients from a team.
func (c *Client) ListRecipients(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListRecipientsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListRecipientsRequest create the request corresponding to the list action endpoint of the recipients resource.
func (c *Client) NewListRecipientsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateRecipientsPath computes a request path to the update action of recipients.
func UpdateRecipientsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s/recipients", param0)
}

// Update team recipients.
func (c *Client) UpdateRecipients(ctx context.Context, path string, payload *RecipientsPayload) (*http.Response, error) {
	req, err := c.NewUpdateRecipientsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateRecipientsRequest create the request corresponding to the update action endpoint of the recipients resource.
func (c *Client) NewUpdateRecipientsRequest(ctx context.Context, path string, payload *RecipientsPayload) (*http.Request, error) {
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
	req, err := http.NewRequestWithContext(ctx, "PUT", u.String(), &body)
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