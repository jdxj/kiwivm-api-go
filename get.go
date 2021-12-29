package kiwivm_sdk_go

import (
	"encoding/json"
	"net"
	"strconv"
)

type GetServiceInfoRsp struct {
	// Hypervizor type (ovz or kvm)
	VmType string `json:"vm_type"`
	// Hostname of the VPS
	Hostname string `json:"hostname"`
	// IP address of the physical node
	NodeIp string `json:"node_ip"`
	// Internal nickname of the physical node
	NodeAlias string `json:"node_alias"`
	// Physical location (country, state)
	NodeLocation string `json:"node_location"`
	// Whether IPv6 is supported at the current location
	LocationIpv6Ready bool `json:"location_ipv6_ready"`
	// Name of plan
	Plan string `json:"plan"`
	// Disk quota (bytes)
	PlanDisk int64 `json:"plan_disk"`
	// RAM (bytes)
	PlanRam int `json:"plan_ram"`
	// SWAP (bytes)
	PlanSwap int `json:"plan_swap"`
	// Operating system
	OS string `json:"os"`
	// Primary e-mail address of the account
	Email string `json:"email"`
	// Allowed monthly data transfer (bytes).
	// Needs to be multiplied by monthly_data_multiplier - see below.
	PlanMonthlyData int64 `json:"plan_monthly_data"`
	// Data transfer used in the current billing month.
	// Needs to be multiplied by monthly_data_multiplier - see below.
	DataCounter int64 `json:"data_counter"`
	// Some locations offer more expensive bandwidth;
	// this variable contains the bandwidth accounting coefficient.
	MonthlyDataMultiplier int `json:"monthly_data_multiplier"`
	// Date and time of transfer counter reset (UNIX timestamp)
	DataNextReset int64 `json:"data_next_reset"`
	// IPv4 and IPv6 addresses assigned to VPS (Array)
	IpAddresses []string `json:"ip_addresses"`
	// Private IPv4 addresses assigned to VPS (Array)
	PrivateIpAddresses []string `json:"private_ip_addresses"`
	// Information on IP address nullrouting during (D)DoS attacks (Array).
	// Sample output when IP is under attack.
	IpNullRoutes json.RawMessage `json:"ip_nullroutes"`
	// Mounted image #1
	Iso1 json.RawMessage `json:"iso1"`
	// Mounted image #2 (currently unsupported)
	Iso2 json.RawMessage `json:"iso2"`
	// Array of ISO images available for use
	AvailableISOs []string `json:"available_isos"`
	// Maximum number of IPv6 addresses allowed by plan
	PlanMaxIPV6s int `json:"plan_max_ipv6s"`
	// Whether or not rDNS records can be set via API
	RDNSAPIAvailable bool `json:"rdns_api_available"`
	// Whether or not Private Network features are available on this plan
	PlanPrivateNetworkAvailable bool `json:"plan_private_network_available"`
	// Whether or not Private Network features are available at this location
	LocationPrivateNetworkAvailable bool `json:"location_private_network_available"`
	// rDNS records (Array of two-dimensional arrays: ip=>value)
	// todo: test
	Ptr map[string]json.RawMessage `json:"ptr"`
	// Whether VPS is suspended
	Suspended bool `json:"suspended"`
	// Whether there is an active policy violation that needs attention (see getPolicyViolations)
	PolicyViolation bool `json:"policy_violation"`
	// Number of times service was suspended in current calendar year
	SuspensionCount json.RawMessage `json:"suspension_count"`
	// Total abuse points accumulated in current calendar year
	TotalAbusePoints int `json:"total_abuse_points"`
	// Maximum abuse points allowed by plan in a calendar year
	MaxAbusePoints int    `json:"max_abuse_points"`
	NodeLocationId string `json:"node_location_id"`
	NodeDatacenter string `json:"node_datacenter"`
	Status
}

func (c *Client) GetServiceInfo() (*GetServiceInfoRsp, error) {
	call := "/getServiceInfo"
	req := c.auth
	rsp := &GetServiceInfoRsp{}
	return rsp, c.do(call, req, rsp)
}

type GetLiveServiceInfoRsp struct {
	GetServiceInfoRsp

	// Depending on hypervisor this call will return the following information:
	//
	// [OVZ hypervisor]
	// array containing OpenVZ beancounters, system load average,
	// number of processes, open files, sockets, memory usage etc
	VzStatus json.RawMessage `json:"vz_status"`
	// array containing OpenVZ disk size, inodes and usage info
	VzQuota json.RawMessage `json:"vz_quota"`

	// [KVM hypervisor]
	// Starting, Running or Stopped
	VeStatus string `json:"ve_status"`
	// MAC address of primary network interface
	VeMac1 string `json:"ve_mac1"`
	// Occupied (mapped) disk space in bytes
	VeUsedDiskSpaceB int64 `json:"ve_used_disk_space_b"`
	// Actual size of disk image in GB
	VeDiskQuotaGb string `json:"ve_disk_quota_gb"`
	// 0 = Disk I/O is not throttled, 1 = Disk I/O is throttled due to high usage.
	// Throttling resets automatically every 15-180 minutes depending on sustained
	// storage I/O utilization.
	// todo: test
	IsDiskThrottled string `json:"is_disk_throttled"`
	// Result of "hostname" command executed inside VPS
	LiveHostname string `json:"live_hostname"`
	// Raw load average string
	LoadAverage string `json:"load_average"`
	// Amount of available RAM in KB
	MemAvailableKB int `json:"mem_available_kb"`
	// Total amount of Swap in KB
	SwapTotalKB int `json:"swap_total_kb"`
	// Amount of available Swap in KB
	SwapAvailableKB int `json:"swap_available_kb"`
	// base64 encoded png screenshot of the VGA console
	ScreenDumpPngBase64 string `json:"screendump_png_base64"`
	VeID                int    `json:"veid"`

	// both
	// 0 = CPU is not throttled, 1 = CPU is throttled due to high usage.
	// Throttling resets automatically every 2 hours.
	// todo: test
	IsCpuThrottled string `json:"is_cpu_throttled"`
	// SSH port of the VPS (returned only if VPS is running)
	SSHPort int `json:"ssh_port"`
}

// GetLiveServiceInfo This function returns all data provided by getServiceInfo.
// In addition, it provides detailed status of the VPS.
// Please note that this call may take up to 15 seconds to complete.
func (c *Client) GetLiveServiceInfo() (*GetLiveServiceInfoRsp, error) {
	call := "/getLiveServiceInfo"
	req := c.auth
	rsp := &GetLiveServiceInfoRsp{}
	return rsp, c.do(call, req, rsp)
}

type GetAvailableOSRsp struct {
	Status
	// Currently installed Operating System
	Installed string `json:"installed"`
	// Array of available OS
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
	Status
}

// GetRawUsageStats Returns a two-dimensional array with the detailed
// usage statistics shown under Detailed Statistics in KiwiVM.
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
	Status
	LogEntries []AuditLog `json:"log_entries"`
}

// GetAuditLog Returns an array with the detailed audit
// log shown under Audit Log in KiwiVM.
func (c *Client) GetAuditLog() (*GetAuditLogRsp, error) {
	call := "/getAuditLog"
	req := c.auth
	rsp := &GetAuditLogRsp{}
	return rsp, c.do(call, req, rsp)
}

type GetSuspensionDetailsRsp struct {
}

// GetSuspensionDetails Retrieve information related to service suspensions.
// todo: test
func (c *Client) GetSuspensionDetails() (*GetSuspensionDetailsRsp, error) {
	call := "/getSuspensionDetails"
	req := c.auth
	rsp := &GetSuspensionDetailsRsp{}
	return rsp, c.do(call, req, rsp)
}

type UnsuspendReq struct {
	*Auth
	RecordID string `json:"record_id"`
}

type UnsuspendRsp struct {
	Status
}

// Unsuspend Clear abuse issue identified by record_id and unsuspend the VPS.
// Refer to getSuspensionDetails call for details.
// todo: test
func (c *Client) Unsuspend(req *UnsuspendReq) (*UnsuspendRsp, error) {
	call := "/unsuspend"
	req.Auth = c.auth
	rsp := &UnsuspendRsp{}
	return rsp, c.do(call, req, rsp)
}

type GetPolicyViolationsRsp struct {
	Status
}

// GetPolicyViolations Retrieve information related to active policy violations.
// todo: test
func (c *Client) GetPolicyViolations() (*GetPolicyViolationsRsp, error) {
	call := "/getPolicyViolations"
	req := c.auth
	rsp := &GetPolicyViolationsRsp{}
	return rsp, c.do(call, req, rsp)
}

type ResolvePolicyViolationReq struct {
	*Auth
	RecordID string `json:"record_id"`
}

type ResolvePolicyViolationRsp struct {
	Status
}

// ResolvePolicyViolation Mark policy violation as resolved.
// This is required to avoid service suspension.
// Refer to getPolicyViolations call for details.
// todo: test
func (c *Client) ResolvePolicyViolation(req *ResolvePolicyViolationReq) (*ResolvePolicyViolationRsp, error) {
	call := "/resolvePolicyViolation"
	req.Auth = c.auth
	rsp := &ResolvePolicyViolationRsp{}
	return rsp, c.do(call, req, rsp)
}

type GetRateLimitStatusRsp struct {
	Status
	RemainingPoints15Min int `json:"remaining_points_15min"`
	RemainingPoints24H   int `json:"remaining_points_24h"`
}

// GetRateLimitStatus When you perform too many API calls in a short amount of time,
// KiwiVM API may start dropping your requests for a few minutes.
// This call allows monitoring this matter.
func (c *Client) GetRateLimitStatus() (*GetRateLimitStatusRsp, error) {
	call := "/getRateLimitStatus"
	req := c.auth
	rsp := &GetRateLimitStatusRsp{}
	return rsp, c.do(call, req, rsp)
}

type GetSSHKeysRsp struct {
	Status
	// Per-VM SSH Keys stored in Hypervisor Vault
	SSHKeysVeID string `json:"ssh_keys_veid"`
	// Per-Account SSH keys stored in Billing Portal
	SshKeysUser string `json:"ssh_keys_user"`
	// SSH Keys which will be actually used during a reinstallOS call
	// (Per-VM Keys will always override Per-Account keys)
	SSHKeysPreferred string `json:"ssh_keys_preferred"`
	// Visually shortened keys
	ShortenedSshKeysVeID string `json:"shortened_ssh_keys_veid"`
	// Visually shortened keys
	ShortenedSshKeysUser string `json:"shortened_ssh_keys_user"`
	// Visually shortened keys
	ShortenedSshKeysPreferred string `json:"shortened_ssh_keys_preferred"`
}

// GetSSHKeys Get SSH keys stored in Hypervisor Vault,
// as well as the ones stored in Billing Portal.
func (c *Client) GetSSHKeys() (*GetSSHKeysRsp, error) {
	call := "/getSshKeys"
	req := c.auth
	rsp := &GetSSHKeysRsp{}
	return rsp, c.do(call, req, rsp)
}
