package kiwivm_sdk_go

type Backup struct {
	Size      int    `json:"size"`
	OS        string `json:"os"`
	MD5       string `json:"md5"`
	Timestamp int    `json:"timestamp"`
}

type BackupListRsp struct {
	Status
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
	BackupToken string `json:"backupToken"`
}

type BackupCopyToSnapshotRsp struct {
	Status
	NotificationEmail string `json:"notificationEmail"`
}

// BackupCopyToSnapshot Copies a backup identified by backup_token
// (returned by backup/list) into a restorable Snapshot.
func (c *Client) BackupCopyToSnapshot(req *BackupCopyToSnapshotReq) (*BackupCopyToSnapshotRsp, error) {
	call := "/backup/copyToSnapshot"
	req.Auth = c.auth
	rsp := &BackupCopyToSnapshotRsp{}
	return rsp, c.do(call, req, rsp)
}
