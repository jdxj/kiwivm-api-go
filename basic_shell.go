package kiwi

import "context"

type BasicShellCDReq struct {
	*Auth
	CurrentDir string `json:"currentDir"`
	NewDir     string `json:"newDir"`
}

type BasicShellCDRsp struct {
	Status
	// Result of the "pwd" command after the change.
	PWD string `json:"pwd"`
}

// BasicShellCD Simulate change of directory inside of the VPS.
// Can be used to build a shell like Basic shell.
func (c *Client) BasicShellCD(ctx context.Context, req *BasicShellCDReq) (*BasicShellCDRsp, error) {
	call := "/basicShell/cd"
	req.Auth = c.auth
	return doHTTP[*BasicShellCDReq, *BasicShellCDRsp](ctx, c.hc, call, req)
}

type BasicShellExecReq struct {
	*Auth
	Command string `json:"command"`
}

type BasicShellExecRsp struct {
	// no need Status
	// Exit status code of the executed command
	Error int `json:"error"`
	// Console output of the executed command
	Message string `json:"message"`
}

// BasicShellExec Execute a shell command on the VPS (synchronously).
func (c *Client) BasicShellExec(ctx context.Context, req *BasicShellExecReq) (*BasicShellExecRsp, error) {
	call := "/basicShell/exec"
	req.Auth = c.auth
	return doHTTP[*BasicShellExecReq, *BasicShellExecRsp](ctx, c.hc, call, req)
}
