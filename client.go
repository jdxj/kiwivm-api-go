package kiwivm_sdk_go

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

const (
	host    = "https://api.64clouds.com"
	version = "/v1"
)

func Encode(i interface{}) string {
	if i == nil {
		return ""
	}
	rt := reflect.TypeOf(i)
	if rt.Kind() != reflect.Ptr || rt.Elem().Kind() != reflect.Struct {
		panic("not point to struct of pointer")
	}

	var (
		v  = url.Values{}
		rv = reflect.ValueOf(i).Elem()
	)
	for i := 0; i < rv.NumField(); i++ {
		key := rv.Type().Field(i).Tag.Get("json")
		if key == "" {
			panic("json tag not found")
		}
		value := fmt.Sprintf("%v", rv.Field(i))
		v.Add(key, value)
	}
	return v.Encode()
}

func NewClient(veID, apiKey string) *Client {
	c := &Client{
		veID:   veID,
		apiKey: apiKey,
		hc:     &http.Client{},
	}
	return c
}

type Client struct {
	veID   string
	apiKey string

	hc *http.Client
}

type auth struct {
	VeID   string `json:"veid"`
	APIKey string `json:"api_key"`
}

func (c *Client) getAuth() *auth {
	return &auth{
		VeID:   c.veID,
		APIKey: c.apiKey,
	}
}

func (c *Client) do(call string, req, rsp interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	api := host + version + call + "?" + Encode(req)
	hReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, api, nil)
	hRsp, err := c.hc.Do(hReq)
	if err != nil {
		return err
	}
	defer hRsp.Body.Close()

	//decoder := json.NewDecoder(hRsp.Body)
	//return decoder.Decode(rsp)

	// debug
	d, err := ioutil.ReadAll(hRsp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("debug body: %s\n", d)
	return json.Unmarshal(d, rsp)
}

type StartRsp struct {
	Error     int `json:"error"`
	IsMounted int `json:"is_mounted"`
}

func (c *Client) Start() (*StartRsp, error) {
	call := "/start"
	req := c.getAuth()
	rsp := &StartRsp{}
	return rsp, c.do(call, req, rsp)
}

type StopRsp struct {
	Error int `json:"error"`
}

func (c *Client) Stop() (*StopRsp, error) {
	call := "/stop"
	req := c.getAuth()
	rsp := &StopRsp{}
	return rsp, c.do(call, req, rsp)
}

type RestartRsp struct {
	Error int `json:"error"`
}

func (c *Client) Restart() (*RestartRsp, error) {
	call := "/restart"
	req := c.getAuth()
	rsp := &RestartRsp{}
	return rsp, c.do(call, req, rsp)
}

func ()

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
	req := c.getAuth()
	rsp := &GetServiceInfoRsp{}
	return rsp, c.do(call, req, rsp)
}

type CreateSnapshotReq struct {
	auth
}

type CreateSnapshotRsp struct {
}

func (c *Client) CreateSnapshot() {

}
