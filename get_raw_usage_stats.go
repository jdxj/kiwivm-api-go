package kiwi

import (
	"context"
	"sort"
	"time"
)

type RawUsage struct {
	Timestamp       int64 `json:"timestamp"`
	NetworkInBytes  int64 `json:"network_in_bytes"`
	NetworkOutBytes int64 `json:"network_out_bytes"`
	DiskReadBytes   int64 `json:"disk_read_bytes"`
	DiskWriteBytes  int64 `json:"disk_write_bytes"`
	CpuUsage        int64 `json:"cpu_usage"`
}

func (ru RawUsage) Datetime() string {
	return time.Unix(ru.Timestamp, 0).Format(time.DateTime)
}

type GetRawUsageStatsRsp struct {
	Data   []RawUsage `json:"data"`
	VmType string     `json:"vm_type"`
	Status
}

func (s GetRawUsageStatsRsp) Traffic(beginUnix, endUnix int64) int64 {
	begin, end := 0, len(s.Data)
	if beginUnix > 0 {
		begin, _ = sort.Find(len(s.Data), func(i int) int {
			return int(beginUnix - s.Data[i].Timestamp)
		})
	}
	if endUnix > 0 {
		end, _ = sort.Find(len(s.Data), func(i int) int {
			return int(endUnix - s.Data[i].Timestamp)
		})
	}

	var sum int64
	for i := begin; i < end; i++ {
		sum += s.Data[i].NetworkInBytes + s.Data[i].NetworkOutBytes
	}
	return sum
}

// GetRawUsageStats Returns a two-dimensional array with the detailed
// usage statistics shown under Detailed Statistics in KiwiVM.
func (c *Client) GetRawUsageStats(ctx context.Context) (*GetRawUsageStatsRsp, error) {
	call := "/getRawUsageStats"
	req := c.auth
	return doHTTP[*Auth, *GetRawUsageStatsRsp](ctx, c, call, req)
}
