package kiwivm_sdk_go

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	cc = NewClient(veID, apiKey, WithDebug(true))
)

func TestEncode(t *testing.T) {
	req := &Auth{
		VeID:   veID,
		APIKey: apiKey,
	}
	s := Encode(req)
	fmt.Printf("req: %s\n", s)

	req2 := &SetHostnameReq{
		Auth:        req,
		NewHostname: "",
	}
	s = Encode(req2)
	fmt.Printf("req2: %s\n", s)
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

func TestClient_GetLiveServiceInfo(t *testing.T) {
	rsp, err := cc.GetLiveServiceInfo()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_GetAvailableOS(t *testing.T) {
	rsp, err := cc.GetAvailableOS()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_GetRawUsageStats(t *testing.T) {
	rsp, err := cc.GetRawUsageStats()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_GetAuditLog(t *testing.T) {
	rsp, err := cc.GetAuditLog()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, v := range rsp.LogEntries {
		ip, err := v.GetIP()
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		fmt.Printf("%s\n", ip)
	}
}

func TestJsonParseString(t *testing.T) {
	abc := "\"123\""
	def := ""
	err := json.Unmarshal([]byte(abc), &def)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("res: %s\n", def)
}

func TestClient_SetHostname(t *testing.T) {
	rsp, err := cc.SetHostname(&SetHostnameReq{
		NewHostname: "jxdj-jp",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_SetPTR(t *testing.T) {
	rsp, err := cc.SetPTR(&SetPTRReq{
		IP:  ip,
		PTR: ptr,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_BasicShellCD(t *testing.T) {
	rsp, err := cc.BasicShellCD(&BasicShellCDReq{
		CurrentDir: "/root",
		NewDir:     "download",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_BasicShellExec(t *testing.T) {
	rsp, err := cc.BasicShellExec(&BasicShellExecReq{
		Command: "ls",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_ShellScriptExec(t *testing.T) {
	rsp, err := cc.ShellScriptExec(&ShellScriptExecReq{
		Script: "ls",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_SnapshotCreate(t *testing.T) {
	rsp, err := cc.SnapshotCreate(&SnapshotCreateReq{
		Description: "test",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_SnapshotList(t *testing.T) {
	rsp, err := cc.SnapshotList()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_SnapshotDelete(t *testing.T) {
	rsp, err := cc.SnapshotDelete(&SnapshotDeleteReq{
		Snapshot: "snapshot-1670298-1640598255-2021-12-27-44b654ec70ebb997f3019903743f7d30363f2f51.tar.gz",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_SnapshotToggleSticky(t *testing.T) {
	rsp, err := cc.SnapshotToggleSticky(&SnapshotToggleStickyReq{
		Snapshot: "snapshot-1670298-1640598255-2021-12-27-44b654ec70ebb997f3019903743f7d30363f2f51.tar.gz",
		Sticky:   SetSticky,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}
