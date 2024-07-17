package apollo_test

import (
	"context"
	"fmt"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/CharLemAznable/gfx/ext/gviewx/apollo"
	"github.com/CharLemAznable/gfx/frame/gx"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

var (
	ctx = context.TODO()
)

func Test_Normal(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mockConfig := agollox.DefaultConfig()
		mockConfig.AppID = "test"
		mockServer, _ := agollox.MockServer(mockConfig, "testdata/mockdata")
		mockIP := fmt.Sprintf(`http://127.0.0.1:%d`, mockServer.GetListenedPort())

		_ = gfile.PutContents("testdata/apollo.yaml", `appId: "test"
ip: "`+mockIP+`"
backupConfigPath: ".apollo.bk"`)

		adapter, err := apollo.LoadAdapter(ctx, "testdata/apollo")
		t.AssertNE(adapter, nil)
		t.AssertNil(err)

		content, err := adapter.GetContent("tmpl")
		t.Assert(content, "Hello, {{.Name}}!")
		t.AssertNil(err)

		content, err = adapter.GetContent("tmpl_none")
		t.Assert(content, "")
		t.AssertNE(err, nil)

		gx.ViewX().SetAdapter(adapter)
		parsed, err := gx.ViewX().Parse(ctx, "tmpl", map[string]interface{}{"Name": "gf"})
		t.Assert(parsed, "Hello, gf!")
		t.AssertNil(err)

		_ = gfile.PutContents("testdata/apollo.yaml", "")
	})
}

func Test_Error(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_GCFG_PATH", "errorpath")
		defer func() { _ = genv.Remove("GF_GCFG_PATH") }()

		_, err := apollo.LoadAdapter(ctx, "testdata/apollo")
		t.AssertNE(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_GVIEWX_APOLLO_APPID", "AppId")
		defer func() { _ = genv.Remove("GF_GVIEWX_APOLLO_APPID") }()

		adapter, err := apollo.LoadAdapter(ctx, "testdata/apollo_error")
		t.AssertNil(adapter)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "gviewx: Apollo IP field is required")

		adapter, err = apollo.LoadAdapter(ctx, "testdata/apollo_none")
		t.AssertNil(adapter)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "gviewx: Apollo IP field is required")
	})
	gtest.C(t, func(t *gtest.T) {
		adapter, err := apollo.LoadAdapter(ctx, "testdata/apollo_local")
		t.AssertNil(adapter)
		t.AssertNE(err, nil)
	})
}
