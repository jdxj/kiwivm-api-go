package kiwivm_sdk_go

type MigrateGetLocationsRsp struct {
	Error int `json:"error"`
	// ID of current location
	CurrentLocation string `json:"currentLocation"`
	// IDs of locations available for migration into
	Locations               []string          `json:"locations"`
	Descriptions            map[string]string `json:"descriptions"`
	DataTransferMultipliers map[string]int    `json:"dataTransferMultipliers"`
}

// MigrateGetLocations Return all possible migration locations.
// todo: test
func (c *Client) MigrateGetLocations() (*MigrateGetLocationsRsp, error) {
	call := "/migrate/getLocations"
	req := c.auth
	rsp := &MigrateGetLocationsRsp{}
	return rsp, c.do(call, req, rsp)
}
