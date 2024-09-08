package agollox_test

import (
	"context"
	"fmt"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
	"time"
)

var (
	ctx = context.TODO()
)

func Test_Agollox(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_APOLLO_CONFIG_FILE", "testdata/config.none")
		_ = genv.Set("GF_APOLLO_APPID", "test")
		defer func() { _ = genv.Remove("GF_APOLLO_CONFIG_FILE", "GF_APOLLO_APPID") }()
		_, err := agollox.NewClient(ctx)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "create agollo client failed: Apollo IP field is required")
	})
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_APOLLO_CONFIG_FILE", "testdata/config.error")
		defer func() { _ = genv.Remove("GF_APOLLO_CONFIG_FILE") }()
		_, err := agollox.NewClient(ctx)
		t.AssertNE(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_APOLLO_CONFIG_FILE", "testdata/config.local")
		defer func() { _ = genv.Remove("GF_APOLLO_CONFIG_FILE") }()
		_, err := agollox.NewClient(ctx)
		t.AssertNE(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_APOLLO_CONFIG_FILE", "testdata/config.empty")
		defer func() { _ = genv.Remove("GF_APOLLO_CONFIG_FILE") }()
		_, err := agollox.NewClient(ctx)
		t.AssertNE(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_GCFG_PATH", "errorpath")
		_ = genv.Set("GF_APOLLO_CONFIG_FILE", "errorconfig")
		defer func() { _ = genv.Remove("GF_GCFG_PATH", "GF_APOLLO_CONFIG_FILE") }()
		config, err := agollox.LoadConfig(ctx)
		t.AssertNE(err, nil)
		_, err = agollox.MockServer(config, "errormock")
		t.AssertNE(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		mockFileName := "testdata/mockdata.def.yaml"
		_ = gfile.PutContents(mockFileName, `application:
  key: "value"`)

		mockConfig := agollox.DefaultConfig()
		mockConfig.AppID = "test"
		mockServer, _ := agollox.MockServer(mockConfig, mockFileName)
		defer func() { _ = mockServer.Shutdown() }()
		mockIP := fmt.Sprintf(`http://127.0.0.1:%d`, mockServer.GetListenedPort())

		_ = genv.Set("GF_APOLLO_CONFIG_FILE", "testdata/config.def")
		_ = genv.Set("GF_APOLLO_APPID", "test")
		_ = genv.Set("GF_APOLLO_CLUSTER", "")
		_ = genv.Set("GF_APOLLO_NAMESPACE", "")
		_ = genv.Set("GF_APOLLO_IP", mockIP)
		defer func() {
			_ = genv.Remove("GF_APOLLO_CONFIG_FILE",
				"GF_APOLLO_APPID", "GF_APOLLO_CLUSTER", "GF_APOLLO_NAMESPACE", "GF_APOLLO_IP")
		}()

		client, err := agollox.NewClient(ctx)
		t.AssertNil(err)
		t.Assert(client.Contains("key"), true)
		t.Assert(client.Get("key"), "value")
		t.Assert(client.Map()["key"], "value")

		client2, err := agollox.NewClient(ctx)
		t.AssertNil(err)
		t.Assert(client2.Contains("key"), true)
		t.Assert(client2.Get("key"), "value")
		t.Assert(client2.Map()["key"], "value")

		client.SetChangeListener(agollox.ChangeListenerFunc(func(event *agollox.ChangeEvent) {
			_, ok := event.Changes["key"]
			t.Assert(ok, true)
		}))
		_ = gfile.PutContents(mockFileName, `application:
  key: "new value"`)
		time.Sleep(time.Second * 3)
		t.Assert(client.Get("key"), "new value")

		client.SetChangeListener(agollox.ChangeListenerFunc(func(event *agollox.ChangeEvent) {
			panic("ignored")
		}))
		_ = gfile.PutContents(mockFileName, `application:
  key: "value"`)
		time.Sleep(time.Second * 3)
		t.Assert(client.Get("key"), "value")
	})
}
