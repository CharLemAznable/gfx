package apollo_test

import (
	"context"
	"fmt"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/CharLemAznable/gfx/ext/gcfg/apollo"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
	"time"
)

var (
	ctx = context.TODO()
)

func Test_ApolloAdapter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_GCFG_PATH", "errorpath")
		defer func() { _ = genv.Remove("GF_GCFG_PATH") }()
		_, err := apollo.LoadAdapter(ctx)
		t.AssertNE(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_APOLLO_CONFIG_FILE", "testdata/apollo_none")
		_ = genv.Set("GF_APOLLO_APPID", "test")
		defer func() { _ = genv.Remove("GF_APOLLO_CONFIG_FILE", "GF_APOLLO_APPID") }()
		_, err := apollo.LoadAdapter(ctx)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "create agollo client failed: Apollo IP field is required")
	})
	gtest.C(t, func(t *gtest.T) {
		_ = gfile.PutContents("testdata/mockdata.yaml", `application:
  config: "server:\n  address: \":8080\""
  lazy: "{{"`)

		mockConfig := agollox.DefaultConfig()
		mockConfig.AppID = "test"
		mockServer, _ := agollox.MockServer(mockConfig, "testdata/mockdata")
		defer func() { _ = mockServer.Shutdown() }()
		mockIP := fmt.Sprintf(`http://127.0.0.1:%d`, mockServer.GetListenedPort())

		_ = genv.Set("GF_APOLLO_APPID", "test")
		_ = genv.Set("GF_APOLLO_IP", mockIP)
		_ = genv.Set("GF_GCFG_APOLLO_KEY", "")
		defer func() { _ = genv.Remove("GF_APOLLO_APPID", "GF_APOLLO_IP", "GF_GCFG_APOLLO_KEY") }()

		adapter, err := apollo.LoadAdapter(ctx)
		t.AssertNE(adapter, nil)
		t.AssertNil(err)
		t.Assert(adapter.Available(ctx), true)

		cfgVal, err := adapter.Get(ctx, "server.address")
		t.Assert(cfgVal, ":8080")
		t.AssertNil(err)
		_ = gfile.PutContents("testdata/mockdata.yaml", `application:
  config: "server:\n  address: \":8081\""`)
		time.Sleep(time.Second * 3)
		cfgMap, err := adapter.Data(ctx)
		t.Assert(cfgMap["server"].(map[string]interface{})["address"], ":8081")
		t.AssertNil(err)

		_ = genv.Set("GF_GCFG_APOLLO_KEY", "notfound")

		adapter, err = apollo.LoadAdapter(ctx)
		t.AssertNE(adapter, nil)
		t.AssertNil(err)
		t.Assert(adapter.Available(ctx), false)

		cfgVal, err = adapter.Get(ctx, "server.address")
		t.AssertNil(cfgVal)
		t.AssertNil(err)
		cfgMap, err = adapter.Data(ctx)
		t.AssertNil(cfgMap)
		t.AssertNil(err)

		_ = gfile.PutContents("testdata/mockdata.yaml", `application:
  lazy: "{{"`)
		time.Sleep(time.Second * 3)
		_ = genv.Set("GF_GCFG_APOLLO_KEY", "lazy")
		_ = genv.Set("GF_GCFG_APOLLO_WATCH", "false")
		defer func() { _ = genv.Remove("GF_GCFG_APOLLO_WATCH") }()

		adapter, _ = apollo.LoadAdapter(ctx)
		_, err = adapter.Get(ctx, "server.address")
		t.AssertNE(err, nil)

		adapter, _ = apollo.LoadAdapter(ctx)
		_, err = adapter.Data(ctx)
		t.AssertNE(err, nil)

		_ = gfile.PutContents("testdata/mockdata.yaml", "")
	})
}

//func Test_Error(t *testing.T) {
//	gtest.C(t, func(t *gtest.T) {
//		_ = genv.Set("GF_GCFG_APOLLO_APPID", "AppId")
//		defer func() { _ = genv.Remove("GF_GCFG_APOLLO_APPID") }()
//
//		adapter, err := apollo.LoadAdapter(ctx, "testdata/apollo_error")
//		t.AssertNil(adapter)
//		t.AssertNE(err, nil)
//		t.Assert(err.Error(), "Apollo Key field is required")
//
//		adapter, err = apollo.LoadAdapter(ctx, "testdata/apollo_none")
//		t.AssertNil(adapter)
//		t.AssertNE(err, nil)
//		t.Assert(err.Error(), "Apollo IP field is required")
//	})
//	gtest.C(t, func(t *gtest.T) {
//		adapter, err := apollo.LoadAdapter(ctx, "testdata/apollo_local")
//		t.AssertNil(adapter)
//		t.AssertNE(err, nil)
//	})
//}
