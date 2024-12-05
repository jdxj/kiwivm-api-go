package kiwi

import "context"

type MigrateGetLocationsRsp struct {
	Status

	// ID of current location
	CurrentLocation string `json:"currentLocation"`

	// IDs of locations available for migration into
	Locations []string `json:"locations"`

	// Friendly descriptions of available locations
	Descriptions map[string]string `json:"descriptions"`

	// Some locations may offer more expensive bandwidth
	// where monthly allowance will be lower.
	// This array contains monthly data transfer allowance
	// multipliers for each location.
	DataTransferMultipliers map[string]int `json:"dataTransferMultipliers"`
}

// MigrateGetLocations Return all possible migration locations.
func (c *Client) MigrateGetLocations(ctx context.Context) (*MigrateGetLocationsRsp, error) {
	call := "/migrate/getLocations"
	req := c.auth
	return doHTTP[*Auth, *MigrateGetLocationsRsp](ctx, c, call, req)
}

type MigrateStartReq struct {
	*Auth
	Location string `json:"location"`
}

type MigrateStartRsp struct{}

// MigrateStart Start VPS migration to new location.
// Takes new location ID as input.
// Note that this will result in all IPv4 addresses
// to be replaced with new ones, and all IPv6 addresses
// will be released.
// todo: test
func (c *Client) MigrateStart(ctx context.Context, req *MigrateStartReq) (*MigrateStartRsp, error) {
	call := "/migrate/start"
	req.Auth = c.auth
	return doHTTP[*MigrateStartReq, *MigrateStartRsp](ctx, c, call, req)
}

type CloneFromExternalServerReq struct {
	*Auth
	ExternalServerIP           string `json:"externalServerIP"`
	ExternalServerSSHPort      string `json:"externalServerSSHport"`
	ExternalServerRootPassword string `json:"externalServerRootPassword"`
}

type CloneFromExternalServerRsp struct{}

// CloneFromExternalServer (OVZ only) Clone a remote server or VPS.
// See Migrate from another server for example on how this works.
// todo: test
func (c *Client) CloneFromExternalServer(ctx context.Context, req *CloneFromExternalServerReq) (*CloneFromExternalServerRsp, error) {
	call := "/cloneFromExternalServer"
	req.Auth = c.auth
	return doHTTP[*CloneFromExternalServerReq, *CloneFromExternalServerRsp](ctx, c, call, req)
}
