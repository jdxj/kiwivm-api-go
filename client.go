package kiwi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"

	"github.com/go-resty/resty/v2"
)

const (
	host    = "https://api.64clouds.com"
	version = "/v1"
)

// Encode encodes the field into the form of key1=value1&key2=value2
// according to the go struct tag.
func Encode(i any) string {
	return EncodeValues(i).Encode()
}

func EncodeValues(i any) url.Values {
	v := url.Values{}
	encode(i, v)
	return v
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
	debug  bool
	logger resty.Logger
}

type OptFunc func(*Option)

func WithDebug(debug bool) OptFunc {
	return func(o *Option) {
		o.debug = debug
	}
}

func WithLogger(logger resty.Logger) OptFunc {
	return func(o *Option) {
		o.logger = logger
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
		hc:     resty.New(),
	}
	return c
}

type Client struct {
	auth   *Auth
	option *Option

	hc *resty.Client
}

type Auth struct {
	VeID   string `json:"veid"`
	APIKey string `json:"api_key"`
}

func doHTTP[R, S any](ctx context.Context, client *Client, path string, req R) (rsp S, err error) {
	api := host + version + path
	hs, err := client.hc.R().
		SetDebug(client.option.debug).
		SetContext(ctx).
		SetQueryParamsFromValues(EncodeValues(req)).
		Get(api)
	if err != nil {
		return
	}
	// rsp 的 Content-Type 是 text，所以这里手动解析 body
	err = json.Unmarshal(hs.Body(), &rsp)
	return
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
func (c *Client) Start(ctx context.Context) (*StartRsp, error) {
	call := "/start"
	req := c.auth
	return doHTTP[*Auth, *StartRsp](ctx, c, call, req)
}

type StopRsp struct {
	Status
}

// Stop the VPS
func (c *Client) Stop(ctx context.Context) (*StopRsp, error) {
	call := "/stop"
	req := c.auth
	return doHTTP[*Auth, *StopRsp](ctx, c, call, req)
}

type RestartRsp struct {
	Status
}

// Restart Reboots the VPS
func (c *Client) Restart(ctx context.Context) (*RestartRsp, error) {
	call := "/restart"
	req := c.auth
	return doHTTP[*Auth, *RestartRsp](ctx, c, call, req)
}

type KillRsp struct {
	Status
}

// Kill Allows to forcibly stop a VPS that is stuck and cannot be stopped by normal means.
// Please use this feature with great care as any unsaved data will be lost.
// todo: test
func (c *Client) Kill(ctx context.Context) (*KillRsp, error) {
	call := "/kill"
	req := c.auth
	return doHTTP[*Auth, *KillRsp](ctx, c, call, req)
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
func (c *Client) ReinstallOS(ctx context.Context, req *ReinstallOSReq) (*ReinstallOSRsp, error) {
	call := "/reinstallOS"
	req.Auth = c.auth
	return doHTTP[*ReinstallOSReq, *ReinstallOSRsp](ctx, c, call, req)
}

type ResetRootPasswordRsp struct {
	Status
}

// ResetRootPassword Generates and sets a new root password.
// todo: test
func (c *Client) ResetRootPassword(ctx context.Context) (*ResetRootPasswordRsp, error) {
	call := "/resetRootPassword"
	req := c.auth
	return doHTTP[*Auth, *ResetRootPasswordRsp](ctx, c, call, req)
}
