package kiwivm_sdk_go

import (
	"fmt"
	"testing"
)

var (
	cc = NewClient(veID, apiKey)
)

func TestEncode(t *testing.T) {
	req := &auth{
		VeID:   veID,
		APIKey: apiKey,
	}
	_ = req
	s := Encode(nil)
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", veID)
}

func TestNewClient(t *testing.T) {
	c := NewClient(veID, apiKey)
	rsp, err := c.GetServiceInfo()
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_Start(t *testing.T) {
	rsp, err := cc.Start()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_Stop(t *testing.T) {
	rsp, err := cc.Stop()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_Restart(t *testing.T) {
	rsp, err := cc.Restart()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}
