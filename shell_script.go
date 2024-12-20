package kiwi

import "context"

type ShellScriptExecReq struct {
	*Auth
	Script string `json:"script"`
}

type ShellScriptExecRsp struct {
	Status
	NodeIP string `json:"node_ip"`
	// Name of the output log file.
	Log            string `json:"log"`
	OutputStreamID string `json:"output_stream_id"`
}

// ShellScriptExec Execute a shell script on the VPS (asynchronously).
func (c *Client) ShellScriptExec(ctx context.Context, req *ShellScriptExecReq) (*ShellScriptExecRsp, error) {
	call := "/shellScript/exec"
	req.Auth = c.auth
	return doHTTP[*ShellScriptExecReq, *ShellScriptExecRsp](ctx, c, call, req)
}
