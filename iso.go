package kiwi

import "context"

type ISOMountReq struct {
	*Auth
	ISO string `json:"iso"`
}

type ISOMountRsp struct{}

// ISOMount Sets ISO image to boot from.
// VM must be completely shut down and restarted after this API call.
// todo: test
func (c *Client) ISOMount(ctx context.Context, req *ISOMountReq) (*ISOMountRsp, error) {
	call := "/iso/mount"
	req.Auth = c.auth
	return doHTTP[*ISOMountReq, *ISOMountRsp](ctx, c.hc, call, req)
}

type ISOUnmountRsp struct{}

// ISOUnmount Removes ISO image and configures VM to boot from primary storage.
// VM must be completely shut down and restarted after this API call.
// todo: test
func (c *Client) ISOUnmount(ctx context.Context) (*ISOUnmountRsp, error) {
	call := "/iso/unmount"
	req := c.auth
	return doHTTP[*Auth, *ISOUnmountRsp](ctx, c.hc, call, req)
}
