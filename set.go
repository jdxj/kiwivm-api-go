package kiwi

import "context"

type SetHostnameReq struct {
	*Auth
	NewHostname string `json:"newHostname"`
}

type SetHostnameRsp struct {
	Status
}

// SetHostname Sets new hostname.
func (c *Client) SetHostname(ctx context.Context, req *SetHostnameReq) (*SetHostnameRsp, error) {
	call := "/setHostname"
	req.Auth = c.auth
	return doHTTP[*SetHostnameReq, *SetHostnameRsp](ctx, c.hc, call, req)
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
func (c *Client) SetPTR(ctx context.Context, req *SetPTRReq) (*SetPTRRsp, error) {
	call := "/setPTR"
	req.Auth = c.auth
	return doHTTP[*SetPTRReq, *SetPTRRsp](ctx, c.hc, call, req)
}
