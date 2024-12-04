package kiwi

type UpdateSSHKeysReq struct {
	*Auth
	SSHKeys string `json:"ssh_keys"`
}

type UpdateSSHKeysRsp struct{}

// UpdateSSHKeys Update per-VM SSH keys in Hypervisor Vault.
// Keys will be written to /root/.ssh/authorized_keys during a reinstallOS call.
// These keys will override any keys set in Billing Portal.
// todo: test
func (c *Client) UpdateSSHKeys(req *UpdateSSHKeysReq) (*UpdateSSHKeysRsp, error) {
	call := "/updateSshKeys"
	req.Auth = c.auth
	rsp := &UpdateSSHKeysRsp{}
	return rsp, c.do(call, req, rsp)
}
