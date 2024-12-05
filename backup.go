package kiwi

import "context"

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
func (c *Client) BackupList(ctx context.Context) (*BackupListRsp, error) {
	call := "/backup/list"
	rsp, err := doHTTP[*Auth, *BackupListRsp](ctx, c, call, c.auth)
	return rsp, err
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
func (c *Client) BackupCopyToSnapshot(ctx context.Context, req *BackupCopyToSnapshotReq) (*BackupCopyToSnapshotRsp, error) {
	call := "/backup/copyToSnapshot"
	req.Auth = c.auth
	return doHTTP[*BackupCopyToSnapshotReq, *BackupCopyToSnapshotRsp](ctx, c, call, req)
}
