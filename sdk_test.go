package kiwivm_sdk_go

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jdxj/kiwivm-sdk-go/conf"
)

var (
	cc = NewClient(conf.VeID, conf.APIKey, WithDebug(true))
)

func TestEncode(t *testing.T) {
	req := &Auth{
		VeID:   conf.VeID,
		APIKey: conf.APIKey,
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

func TestEncode_SnapshotToggleStickyReq(t *testing.T) {
	tsr := &SnapshotToggleStickyReq{
		Auth: &Auth{
			VeID:   "abc",
			APIKey: "def",
		},
		Snapshot: "ghi",
		Sticky:   13,
	}
	fmt.Printf("%s\n", Encode(tsr))
}

func TestClient_GetServiceInfo(t *testing.T) {
	c := NewClient(conf.VeID, conf.APIKey, WithDebug(true))
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
		ip, err := v.IPV4()
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
		IP:  conf.IP,
		PTR: conf.PTR,
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
		Description: "test sticky",
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
		Snapshot: "snapshot-1670298-1641526373-2022-01-06-058c1a7336b49aa3b6ab94f49cdac42787abc15f.tar.gz",
		Sticky:   RemoveSticky,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_SnapshotExport(t *testing.T) {
	rsp, err := cc.SnapshotExport(&SnapshotExportReq{
		Snapshot: "snapshot-1670298-1640655921-2021-12-27-3af2c5902ed7f23ae32e6f1ff4cd71aa50e5919d.tar.gz",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_BackupList(t *testing.T) {
	rsp, err := cc.BackupList()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_BackupCopyToSnapshot(t *testing.T) {
	rsp, err := cc.BackupCopyToSnapshot(&BackupCopyToSnapshotReq{
		BackupToken: "52afc2c8204d2ea2f164a46c3bf06e6a6d05644e",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_MigrateGetLocations(t *testing.T) {
	rsp, err := cc.MigrateGetLocations()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_GetSuspensionDetails(t *testing.T) {
	rsp, err := cc.GetSuspensionDetails()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_GetPolicyViolations(t *testing.T) {
	rsp, err := cc.GetPolicyViolations()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_GetRateLimitStatus(t *testing.T) {
	rsp, err := cc.GetRateLimitStatus()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_PrivateIPGetAvailableIPs(t *testing.T) {
	rsp, err := cc.PrivateIPGetAvailableIPs()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_GetSSHKeys(t *testing.T) {
	rsp, err := cc.GetSSHKeys()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}
