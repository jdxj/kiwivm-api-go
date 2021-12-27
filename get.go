package kiwivm_sdk_go

import (
	"encoding/json"
	"net"
	"strconv"
)

type GetServiceInfoRsp struct {
	VmType                          string                     `json:"vm_type"`
	Hostname                        string                     `json:"hostname"`
	NodeIp                          string                     `json:"node_ip"`
	NodeAlias                       string                     `json:"node_alias"`
	NodeLocation                    string                     `json:"node_location"`
	NodeLocationId                  string                     `json:"node_location_id"`
	NodeDatacenter                  string                     `json:"node_datacenter"`
	LocationIpv6Ready               bool                       `json:"location_ipv6_ready"`
	Plan                            string                     `json:"plan"`
	PlanMonthlyData                 int64                      `json:"plan_monthly_data"`
	MonthlyDataMultiplier           int                        `json:"monthly_data_multiplier"`
	PlanDisk                        int64                      `json:"plan_disk"`
	PlanRam                         int                        `json:"plan_ram"`
	PlanSwap                        int                        `json:"plan_swap"`
	PlanMaxIpv6S                    int                        `json:"plan_max_ipv6s"`
	Os                              string                     `json:"os"`
	Email                           string                     `json:"email"`
	DataCounter                     int64                      `json:"data_counter"`
	DataNextReset                   int                        `json:"data_next_reset"`
	IpAddresses                     []string                   `json:"ip_addresses"`
	PrivateIpAddresses              []string                   `json:"private_ip_addresses"`
	IpNullRoutes                    json.RawMessage            `json:"ip_nullroutes"`
	Iso1                            json.RawMessage            `json:"iso1"`
	Iso2                            json.RawMessage            `json:"iso2"`
	AvailableISOs                   []string                   `json:"available_isos"`
	PlanPrivateNetworkAvailable     bool                       `json:"plan_private_network_available"`
	LocationPrivateNetworkAvailable bool                       `json:"location_private_network_available"`
	RDNSAPIAvailable                bool                       `json:"rdns_api_available"`
	Ptr                             map[string]json.RawMessage `json:"ptr"`
	Suspended                       bool                       `json:"suspended"`
	PolicyViolation                 bool                       `json:"policy_violation"`
	SuspensionCount                 json.RawMessage            `json:"suspension_count"`
	TotalAbusePoints                int                        `json:"total_abuse_points"`
	MaxAbusePoints                  int                        `json:"max_abuse_points"`
	Error                           int                        `json:"error"`
}

func (c *Client) GetServiceInfo() (*GetServiceInfoRsp, error) {
	call := "/getServiceInfo"
	req := c.auth
	rsp := &GetServiceInfoRsp{}
	return rsp, c.do(call, req, rsp)
}

type GetLiveServiceInfoRsp struct {
	VmType                          string                     `json:"vm_type"`
	VeStatus                        string                     `json:"ve_status"`
	VeMac1                          string                     `json:"ve_mac1"`
	VeUsedDiskSpaceB                int64                      `json:"ve_used_disk_space_b"`
	VeDiskQuotaGb                   string                     `json:"ve_disk_quota_gb"`
	IsCpuThrottled                  string                     `json:"is_cpu_throttled"`
	IsDiskThrottled                 string                     `json:"is_disk_throttled"`
	SshPort                         int                        `json:"ssh_port"`
	LiveHostname                    string                     `json:"live_hostname"`
	LoadAverage                     string                     `json:"load_average"`
	MemAvailableKb                  int                        `json:"mem_available_kb"`
	SwapTotalKb                     int                        `json:"swap_total_kb"`
	SwapAvailableKb                 int                        `json:"swap_available_kb"`
	ScreenDumpPngBase64             string                     `json:"screendump_png_base64"`
	Hostname                        string                     `json:"hostname"`
	NodeIp                          string                     `json:"node_ip"`
	NodeAlias                       string                     `json:"node_alias"`
	NodeLocation                    string                     `json:"node_location"`
	NodeLocationId                  string                     `json:"node_location_id"`
	NodeDatacenter                  string                     `json:"node_datacenter"`
	LocationIpv6Ready               bool                       `json:"location_ipv6_ready"`
	Plan                            string                     `json:"plan"`
	PlanMonthlyData                 int64                      `json:"plan_monthly_data"`
	MonthlyDataMultiplier           int                        `json:"monthly_data_multiplier"`
	PlanDisk                        int64                      `json:"plan_disk"`
	PlanRam                         int                        `json:"plan_ram"`
	PlanSwap                        int                        `json:"plan_swap"`
	PlanMaxIpv6S                    int                        `json:"plan_max_ipv6s"`
	Os                              string                     `json:"os"`
	Email                           string                     `json:"email"`
	DataCounter                     int64                      `json:"data_counter"`
	DataNextReset                   int                        `json:"data_next_reset"`
	IpAddresses                     []string                   `json:"ip_addresses"`
	PrivateIpAddresses              []string                   `json:"private_ip_addresses"`
	IpNullRoutes                    json.RawMessage            `json:"ip_nullroutes"`
	Iso1                            json.RawMessage            `json:"iso1"`
	Iso2                            json.RawMessage            `json:"iso2"`
	AvailableISOs                   []string                   `json:"available_isos"`
	PlanPrivateNetworkAvailable     bool                       `json:"plan_private_network_available"`
	LocationPrivateNetworkAvailable bool                       `json:"location_private_network_available"`
	RDNSAPIAvailable                bool                       `json:"rdns_api_available"`
	Ptr                             map[string]json.RawMessage `json:"ptr"`
	Suspended                       bool                       `json:"suspended"`
	PolicyViolation                 bool                       `json:"policy_violation"`
	SuspensionCount                 json.RawMessage            `json:"suspension_count"`
	TotalAbusePoints                int                        `json:"total_abuse_points"`
	MaxAbusePoints                  int                        `json:"max_abuse_points"`
	Error                           int                        `json:"error"`
	VeID                            int                        `json:"veid"`
}

func (c *Client) GetLiveServiceInfo() (*GetLiveServiceInfoRsp, error) {
	call := "/getLiveServiceInfo"
	req := c.auth
	rsp := &GetLiveServiceInfoRsp{}
	return rsp, c.do(call, req, rsp)
}

type GetAvailableOSRsp struct {
	Error     int      `json:"error"`
	Installed string   `json:"installed"`
	Templates []string `json:"templates"`
}

func (c *Client) GetAvailableOS() (*GetAvailableOSRsp, error) {
	call := "/getAvailableOS"
	req := c.auth
	rsp := &GetAvailableOSRsp{}
	return rsp, c.do(call, req, rsp)
}

type RawUsage struct {
	Timestamp       int64 `json:"timestamp,string"`
	NetworkInBytes  int64 `json:"network_in_bytes,string"`
	NetworkOutBytes int64 `json:"network_out_bytes,string"`
	DiskReadBytes   int64 `json:"disk_read_bytes,string"`
	DiskWriteBytes  int64 `json:"disk_write_bytes,string"`
	CpuUsage        int64 `json:"cpu_usage,string"`
}

type GetRawUsageStatsRsp struct {
	Data   []RawUsage `json:"data"`
	VmType string     `json:"vm_type"`
	Error  int        `json:"error"`
}

func (c *Client) GetRawUsageStats() (*GetRawUsageStatsRsp, error) {
	call := "/getRawUsageStats"
	req := c.auth
	rsp := &GetRawUsageStatsRsp{}
	return rsp, c.do(call, req, rsp)
}

type AuditLog struct {
	Timestamp     int64           `json:"timestamp,string"`
	RequestorIpv4 json.RawMessage `json:"requestor_ipv4"`
	Type          int             `json:"type,string"`
	Summary       string          `json:"summary"`
}

// todo: refactor
func (al AuditLog) GetIP() (net.IP, error) {
	num := len(al.RequestorIpv4)
	// 空值
	if num <= 1 {
		return net.IP{}, nil
	}

	ipUint32Str := ""
	err := json.Unmarshal(al.RequestorIpv4, &ipUint32Str)
	if err != nil {
		return nil, err
	}
	tmp, err := strconv.ParseUint(ipUint32Str, 10, 32)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 4)
	for i := 3; i >= 0; i-- {
		buf[i] = byte(tmp)
		tmp = tmp >> 8
	}
	return buf, nil
}

type GetAuditLogRsp struct {
	Error      int        `json:"error"`
	LogEntries []AuditLog `json:"log_entries"`
}

func (c *Client) GetAuditLog() (*GetAuditLogRsp, error) {
	call := "/getAuditLog"
	req := c.auth
	rsp := &GetAuditLogRsp{}
	return rsp, c.do(call, req, rsp)
}
