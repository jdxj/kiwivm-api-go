package kiwivm_sdk_go

type ISOMountReq struct {
	*Auth
	ISO string `json:"iso"`
}

type ISOMountRsp struct {
}

// ISOMount Sets ISO image to boot from.
// VM must be completely shut down and restarted after this API call.
// todo: test
func (c *Client) ISOMount(req *ISOMountReq) (*ISOMountRsp, error) {
	call := "/iso/mount"
	req.Auth = c.auth
	rsp := &ISOMountRsp{}
	return rsp, c.do(call, req, rsp)
}

type ISOUnmountRsp struct {
}

// ISOUnmount Removes ISO image and configures VM to boot from primary storage.
// VM must be completely shut down and restarted after this API call.
// todo: test
func (c *Client) ISOUnmount() (*ISOUnmountRsp, error) {
	call := "/iso/unmount"
	req := c.auth
	rsp := &ISOUnmountRsp{}
	return rsp, c.do(call, req, rsp)
}
