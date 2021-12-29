package kiwivm_sdk_go

type MigrateGetLocationsRsp struct {
	Error int `json:"error"`

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
func (c *Client) MigrateGetLocations() (*MigrateGetLocationsRsp, error) {
	call := "/migrate/getLocations"
	req := c.auth
	rsp := &MigrateGetLocationsRsp{}
	return rsp, c.do(call, req, rsp)
}

type MigrateStartReq struct {
	*Auth
	Location string `json:"location"`
}

type MigrateStartRsp struct {
}

// MigrateStart Start VPS migration to new location.
// Takes new location ID as input.
// Note that this will result in all IPv4 addresses
// to be replaced with new ones, and all IPv6 addresses
// will be released.
// todo: test
func (c *Client) MigrateStart(req *MigrateStartReq) (*MigrateStartRsp, error) {
	call := "/migrate/start"
	req.Auth = c.auth
	rsp := &MigrateStartRsp{}
	return rsp, c.do(call, req, rsp)
}

type CloneFromExternalServerReq struct {
	*Auth
	ExternalServerIP           string `json:"externalServerIP"`
	ExternalServerSSHPort      string `json:"externalServerSSHport"`
	ExternalServerRootPassword string `json:"externalServerRootPassword"`
}

type CloneFromExternalServerRsp struct {
}

// CloneFromExternalServer (OVZ only) Clone a remote server or VPS.
// See Migrate from another server for example on how this works.
// todo: test
func (c *Client) CloneFromExternalServer(req *CloneFromExternalServerReq) (*CloneFromExternalServerRsp, error) {
	call := "/cloneFromExternalServer"
	req.Auth = c.auth
	rsp := &CloneFromExternalServerRsp{}
	return rsp, c.do(call, req, rsp)
}
