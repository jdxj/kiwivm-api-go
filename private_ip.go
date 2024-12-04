package kiwi

import "context"

type PrivateIPGetAvailableIPsRsp struct {
	Status
}

// PrivateIPGetAvailableIPs Returns all available (free) IPv4 addresses which you can activate on VM
// todo: test
func (c *Client) PrivateIPGetAvailableIPs(ctx context.Context) (*PrivateIPGetAvailableIPsRsp, error) {
	call := "/privateIp/getAvailableIps"
	req := c.auth
	return doHTTP[*Auth, *PrivateIPGetAvailableIPsRsp](ctx, c.hc, call, req)
}

type PrivateIpAssignReq struct {
	*Auth
	// optional
	IP string `json:"ip"`
}

type PrivateIpAssignRsp struct {
	Status
}

// PrivateIpAssign Assign private IP address.
// If IP address not specified, a random address will be assigned.
// todo: test
func (c *Client) PrivateIpAssign(ctx context.Context, req *PrivateIpAssignReq) (*PrivateIpAssignRsp, error) {
	call := "/privateIp/assign"
	req.Auth = c.auth
	return doHTTP[*PrivateIpAssignReq, *PrivateIpAssignRsp](ctx, c.hc, call, req)
}

type PrivateIpDeleteReq struct {
	*Auth
	IP string `json:"ip"`
}

type PrivateIpDeleteRsp struct {
	Status
}

// PrivateIpDelete Delete private IP address.
// todo: test
func (c *Client) PrivateIpDelete(ctx context.Context, req *PrivateIpDeleteReq) (*PrivateIpDeleteRsp, error) {
	call := "/privateIp/delete"
	req.Auth = c.auth
	return doHTTP[*PrivateIpDeleteReq, *PrivateIpDeleteRsp](ctx, c.hc, call, req)
}
