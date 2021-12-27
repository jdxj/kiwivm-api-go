package kiwivm_sdk_go

type ShellScriptExecReq struct {
	*Auth
	Script string `json:"script"`
}

type ShellScriptExecRsp struct {
	Error          int    `json:"error"`
	NodeIP         string `json:"node_ip"`
	Log            string `json:"log"`
	OutputStreamID string `json:"output_stream_id"`
}

// ShellScriptExec Execute a shell script on the VPS (asynchronously).
func (c *Client) ShellScriptExec(req *ShellScriptExecReq) (*ShellScriptExecRsp, error) {
	call := "/shellScript/exec"
	req.Auth = c.auth
	rsp := &ShellScriptExecRsp{}
	return rsp, c.do(call, req, rsp)
}
