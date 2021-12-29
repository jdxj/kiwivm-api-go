package kiwivm_sdk_go

type PrivateIPGetAvailableIPsRsp struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

// PrivateIPGetAvailableIPs Returns all available (free) IPv4 addresses which you can activate on VM
// todo: test
func (c *Client) PrivateIPGetAvailableIPs() (*PrivateIPGetAvailableIPsRsp, error) {
	call := "/privateIp/getAvailableIps"
	req := c.auth
	rsp := &PrivateIPGetAvailableIPsRsp{}
	return rsp, c.do(call, req, rsp)
}

type PrivateIpAssignReq struct {
	*Auth
	// optional
	IP string `json:"ip"`
}

type PrivateIpAssignRsp struct {
}

// PrivateIpAssign Assign private IP address.
// If IP address not specified, a random address will be assigned.
// todo: test
func (c *Client) PrivateIpAssign(req *PrivateIpAssignReq) (*PrivateIpAssignRsp, error) {
	call := "/privateIp/assign"
	req.Auth = c.auth
	rsp := &PrivateIpAssignRsp{}
	return rsp, c.do(call, req, rsp)
}

type PrivateIpDeleteReq struct {
	*Auth
	IP string `json:"ip"`
}

type PrivateIpDeleteRsp struct {
}

// PrivateIpDelete Delete private IP address.
// todo: test
func (c *Client) PrivateIpDelete(req *PrivateIpDeleteReq) (*PrivateIpDeleteRsp, error) {
	call := "/privateIp/delete"
	req.Auth = c.auth
	rsp := &PrivateIpDeleteRsp{}
	return rsp, c.do(call, req, rsp)
}
