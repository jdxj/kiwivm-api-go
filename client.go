package kiwi

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

// Encode encodes the field into the form of key1=value1&key2=value2
// according to the go struct tag.
func Encode(i interface{}) string {
	v := url.Values{}
	encode(i, v)
	return v.Encode()
}

func encode(d interface{}, v url.Values) {
	if d == nil {
		return
	}

	rt := reflect.TypeOf(d)
	if rt.Kind() != reflect.Ptr || rt.Elem().Kind() != reflect.Struct {
		return
	}

	rv := reflect.ValueOf(d).Elem()
	for i := 0; i < rv.NumField(); i++ {
		frt := rv.Field(i).Type()
		if reflect.Int <= frt.Kind() && frt.Kind() <= reflect.Float64 ||
			frt.Kind() == reflect.String {

			key := rv.Type().Field(i).Tag.Get("json")
			if key == "" {
				fieldName := rv.Type().Field(i).Name
				err := fmt.Sprintf("json tag not found: %s", fieldName)
				panic(err)
			}
			value := fmt.Sprintf("%v", rv.Field(i).Interface())
			v.Add(key, value)
		} else {
			encode(rv.Field(i).Interface(), v)
		}
	}
}

type Option struct {
	debug bool
}

type OptFunc func(*Option)

func WithDebug(debug bool) OptFunc {
	return func(o *Option) {
		o.debug = debug
	}
}

func NewClient(veID, apiKey string, optFunc ...OptFunc) *Client {
	o := new(Option)
	for _, f := range optFunc {
		f(o)
	}

	c := &Client{
		auth: &Auth{
			VeID:   veID,
			APIKey: apiKey,
		},
		option: o,
		hc:     &http.Client{},
	}
	return c
}

type Client struct {
	auth   *Auth
	option *Option

	hc *http.Client
}

type Auth struct {
	VeID   string `json:"veid"`
	APIKey string `json:"api_key"`
}

func (c *Client) do(call string, req, rsp interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	api := host + version + call + "?" + Encode(req)

	if c.option.debug {
		fmt.Printf("debug api: %s\n", api)
	}

	hReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, api, nil)
	hRsp, err := c.hc.Do(hReq)
	if err != nil {
		return err
	}
	defer hRsp.Body.Close()

	d, err := ioutil.ReadAll(hRsp.Body)
	if err != nil {
		return err
	}

	if c.option.debug {
		fmt.Printf("debug body: %s\n", d)
	}

	return json.Unmarshal(d, rsp)
}

type Status struct {
	Error               int    `json:"error"`
	Message             string `json:"message"`
	AdditionalErrorInfo string `json:"additionalErrorInfo"`
}

type StartRsp struct {
	Status
	IsMounted int `json:"is_mounted"`
}

// Start the VPS
func (c *Client) Start() (*StartRsp, error) {
	call := "/start"
	req := c.auth
	rsp := &StartRsp{}
	return rsp, c.do(call, req, rsp)
}

type StopRsp struct {
	Status
}

// Stop the VPS
func (c *Client) Stop() (*StopRsp, error) {
	call := "/stop"
	req := c.auth
	rsp := &StopRsp{}
	return rsp, c.do(call, req, rsp)
}

type RestartRsp struct {
	Status
}

// Restart Reboots the VPS
func (c *Client) Restart() (*RestartRsp, error) {
	call := "/restart"
	req := c.auth
	rsp := &RestartRsp{}
	return rsp, c.do(call, req, rsp)
}

type KillRsp struct {
	Status
}

// Kill Allows to forcibly stop a VPS that is stuck and cannot be stopped by normal means.
// Please use this feature with great care as any unsaved data will be lost.
// todo: test
func (c *Client) Kill() (*KillRsp, error) {
	call := "/kill"
	req := c.auth
	rsp := &KillRsp{}
	return rsp, c.do(call, req, rsp)
}

type ReinstallOSReq struct {
	*Auth
	OS string `json:"os"`
}

type ReinstallOSRsp struct {
	Status
}

// ReinstallOS Reinstall the Operating System.
// OS must be specified via "os" variable.
// Use getAvailableOS call to get list of available systems.
// todo: test
func (c *Client) ReinstallOS(req *ReinstallOSReq) (*ReinstallOSRsp, error) {
	call := "/reinstallOS"
	req.Auth = c.auth
	rsp := &ReinstallOSRsp{}
	return rsp, c.do(call, req, rsp)
}

type ResetRootPasswordRsp struct {
	Status
}

// ResetRootPassword Generates and sets a new root password.
// todo: test
func (c *Client) ResetRootPassword() (*ResetRootPasswordRsp, error) {
	call := "/resetRootPassword"
	req := c.auth
	rsp := &ResetRootPasswordRsp{}
	return rsp, c.do(call, req, rsp)
}
