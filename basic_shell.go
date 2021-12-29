package kiwivm_sdk_go

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
func (c *Client) BasicShellCD(req *BasicShellCDReq) (*BasicShellCDRsp, error) {
	call := "/basicShell/cd"
	req.Auth = c.auth
	rsp := &BasicShellCDRsp{}
	return rsp, c.do(call, req, rsp)
}

type BasicShellExecReq struct {
	*Auth
	Command string `json:"command"`
}

type BasicShellExecRsp struct {
	// Exit status code of the executed command
	Error int `json:"error"`
	// Console output of the executed command
	Message string `json:"message"`
}

// BasicShellExec Execute a shell command on the VPS (synchronously).
func (c *Client) BasicShellExec(req *BasicShellExecReq) (*BasicShellExecRsp, error) {
	call := "/basicShell/exec"
	req.Auth = c.auth
	rsp := &BasicShellExecRsp{}
	return rsp, c.do(call, req, rsp)
}
