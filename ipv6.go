package kiwi

import "context"

type IPV6AddReq struct {
	*Auth
	IP string `json:"ip"`
}

type IPV6AddRsp struct {
	Status
}

// IPV6Add Assigns a new IPv6 address.
// For initial IPv6 assignment an empty IP is required (call without parameters),
// and a new IP from the available pool is assigned automatically.
// All subsequent requested IPv6 addresses must be within the /64 subnet of the
// first IPv6 address.
// todo: test
func (c *Client) IPV6Add(ctx context.Context, req *IPV6AddReq) (*IPV6AddRsp, error) {
	call := "/ipv6/add"
	req.Auth = c.auth
	return doHTTP[*IPV6AddReq, *IPV6AddRsp](ctx, c, call, req)
}

type IPV6DeleteReq struct {
	*Auth
	IP string `json:"ip"`
}

type IPV6DeleteRsp struct {
	Status
}

// IPV6Delete Releases specified IPv6 address.
// todo: test
func (c *Client) IPV6Delete(ctx context.Context, req *IPV6DeleteReq) (*IPV6DeleteRsp, error) {
	call := "/ipv6/delete"
	req.Auth = c.auth
	return doHTTP[*IPV6DeleteReq, *IPV6DeleteRsp](ctx, c, call, req)
}
