package kiwivm_sdk_go

type Backup struct {
	Size      int    `json:"size"`
	OS        string `json:"os"`
	MD5       string `json:"md5"`
	Timestamp int    `json:"timestamp"`
}

type BackupListRsp struct {
	Error int `json:"error"`
	// Array of backups (backup_token, size, os, md5, timestamp).
	Backups map[string]Backup `json:"backups"`
}

// BackupList Get list of automatic backups.
func (c *Client) BackupList() (*BackupListRsp, error) {
	call := "/backup/list"
	req := c.auth
	rsp := &BackupListRsp{}
	return rsp, c.do(call, req, rsp)
}

type BackupCopyToSnapshotReq struct {
	*Auth
	BackupToken string `json:"backup_token"`
}

type BackupCopyToSnapshotRsp struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

// BackupCopyToSnapshot Copies a backup identified by backup_token (returned by backup/list) into a restorable Snapshot.
// todo: 测试报错, error: 756130
func (c *Client) BackupCopyToSnapshot(req *BackupCopyToSnapshotReq) (*BackupCopyToSnapshotRsp, error) {
	call := "/backup/copyToSnapshot"
	req.Auth = c.auth
	rsp := &BackupCopyToSnapshotRsp{}
	return rsp, c.do(call, req, rsp)
}
