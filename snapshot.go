package kiwi

import (
	"context"
	"encoding/json"
)

type SnapshotCreateReq struct {
	*Auth
	// optional
	Description string `json:"description"`
}

type SnapshotCreateRsp struct {
	Status
	// E-mail address on file where notification will be sent once task is completed.
	NotificationEmail string `json:"notificationEmail"`
}

// SnapshotCreate Create snapshot
func (c *Client) SnapshotCreate(ctx context.Context, req *SnapshotCreateReq) (*SnapshotCreateRsp, error) {
	call := "/snapshot/create"
	req.Auth = c.auth
	return doHTTP[*SnapshotCreateReq, *SnapshotCreateRsp](ctx, c.hc, call, req)
}

type Snapshot struct {
	FileName     string `json:"fileName"`
	Os           string `json:"os"`
	Description  string `json:"description"`
	Size         string `json:"size"`
	Md5          string `json:"md5"`
	Sticky       bool   `json:"sticky"`
	Uncompressed int64  `json:"uncompressed"`
	PurgesIn     int    `json:"purgesIn"`
	DownloadLink string `json:"downloadLink"`
}

type SnapshotListRsp struct {
	Status
	// Array of snapshots (fileName, os, description, size, md5, sticky, purgesIn, downloadLink).
	Snapshots []Snapshot `json:"snapshots"`
}

// SnapshotList Get list of snapshots.
func (c *Client) SnapshotList(ctx context.Context) (*SnapshotListRsp, error) {
	call := "/snapshot/list"
	req := c.auth
	return doHTTP[*Auth, *SnapshotListRsp](ctx, c.hc, call, req)
}

type SnapshotDeleteReq struct {
	*Auth
	Snapshot string `json:"snapshot"`
}

type SnapshotDeleteRsp struct {
	Status
	NotificationEmail string `json:"notificationEmail"`
}

// SnapshotDelete Delete snapshot by fileName (can be retrieved with snapshot/list call).
func (c *Client) SnapshotDelete(ctx context.Context, req *SnapshotDeleteReq) (*SnapshotDeleteRsp, error) {
	call := "/snapshot/delete"
	req.Auth = c.auth
	return doHTTP[*SnapshotDeleteReq, *SnapshotDeleteRsp](ctx, c.hc, call, req)
}

type SnapshotRestoreReq struct {
	*Auth
	Snapshot string `json:"snapshot"`
}

type SnapshotRestoreRsp struct {
	Status
}

// SnapshotRestore Restores snapshot by fileName (can be retrieved with snapshot/list call).
// This will overwrite all data on the VPS.
// todo: 测试
func (c *Client) SnapshotRestore(ctx context.Context, req *SnapshotRestoreReq) (*SnapshotRestoreRsp, error) {
	call := "/snapshot/restore"
	req.Auth = c.auth
	return doHTTP[*SnapshotRestoreReq, *SnapshotRestoreRsp](ctx, c.hc, call, req)
}

const (
	SetSticky    = 1
	RemoveSticky = 0
)

type SnapshotToggleStickyReq struct {
	*Auth
	Snapshot string `json:"snapshot"`
	Sticky   int    `json:"sticky"`
}

type SnapshotToggleStickyRsp struct {
	Error               int             `json:"error"`
	Message             string          `json:"message"`
	AdditionalErrorCode json.RawMessage `json:"additionalErrorCode"`
	AdditionalErrorInfo string          `json:"additionalErrorInfo"`
}

// SnapshotToggleSticky Set or remove sticky attribute ("sticky" snapshots are never purged).
// Name of snapshot can be retrieved with snapshot/list call – look for fileName variable.
func (c *Client) SnapshotToggleSticky(ctx context.Context, req *SnapshotToggleStickyReq) (*SnapshotToggleStickyRsp, error) {
	call := "/snapshot/toggleSticky"
	req.Auth = c.auth
	return doHTTP[*SnapshotToggleStickyReq, *SnapshotToggleStickyRsp](ctx, c.hc, call, req)
}

type SnapshotExportReq struct {
	*Auth
	Snapshot string `json:"snapshot"`
}

type SnapshotExportRsp struct {
	Status
	Token string `json:"token"`
}

// SnapshotExport Generates a token with which the snapshot can be transferred to another instance.
func (c *Client) SnapshotExport(ctx context.Context, req *SnapshotExportReq) (*SnapshotExportRsp, error) {
	call := "/snapshot/export"
	req.Auth = c.auth
	return doHTTP[*SnapshotExportReq, *SnapshotExportRsp](ctx, c.hc, call, req)
}

type SnapshotImportReq struct {
	*Auth
	SourceVeID  string `json:"sourceVeid"`
	SourceToken string `json:"sourceToken"`
}

type SnapshotImportRsp struct {
	Status
}

// SnapshotImport Imports a snapshot from another instance identified by VEID and Token.
// Both VEID and Token must be obtained from another instance beforehand with a snapshot/export call.
// todo: test
func (c *Client) SnapshotImport(ctx context.Context, req *SnapshotImportReq) (*SnapshotImportRsp, error) {
	call := "/snapshot/import"
	req.Auth = c.auth
	return doHTTP[*SnapshotImportReq, *SnapshotImportRsp](ctx, c.hc, call, req)
}
