package kiwi

type SetHostnameReq struct {
	*Auth
	NewHostname string `json:"newHostname"`
}

type SetHostnameRsp struct {
	Status
}

// SetHostname Sets new hostname.
func (c *Client) SetHostname(req *SetHostnameReq) (*SetHostnameRsp, error) {
	call := "/setHostname"
	req.Auth = c.auth
	rsp := &SetHostnameRsp{}
	return rsp, c.do(call, req, rsp)
}

type SetPTRReq struct {
	*Auth
	IP  string `json:"ip"`
	PTR string `json:"ptr"`
}

type SetPTRRsp struct {
	Status
}

// SetPTR Sets new PTR (rDNS) record for IP.
func (c *Client) SetPTR(req *SetPTRReq) (*SetPTRRsp, error) {
	call := "/setPTR"
	req.Auth = c.auth
	rsp := &SetPTRRsp{}
	return rsp, c.do(call, req, rsp)
}
