package agollox_test

import (
	"context"
	"fmt"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
	"time"
)

var (
	ctx = context.TODO()
)

func Test_Default(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_APOLLO_APPID", "test")
		defer func() { _ = genv.Remove("GF_APOLLO_APPID") }()

		mockFileName := "testdata/mockdata.def.yaml"
		_ = gfile.PutContents(mockFileName, `application:
  key: "value"`)

		mockConfig := agollox.DefaultConfig()
		mockConfig.AppID = "test"
		mockServer, _ := agollox.MockServer(mockConfig, mockFileName)
		defer func() { _ = mockServer.Shutdown() }()
		mockIP := fmt.Sprintf(`http://127.0.0.1:%d`, mockServer.GetListenedPort())

		config := agollox.DefaultConfig()
		config.Cluster = ""
		config.NamespaceName = ""
		config.BackupConfigPath = ".apollo.bk"
		err := agollox.LoadConfig(ctx, config, "testdata/config.def", g.Map{
			agollox.ConfigAppIdKey:     nil,
			agollox.ConfigClusterKey:   nil,
			agollox.ConfigNamespaceKey: nil,
			agollox.ConfigIPKey:        mockIP,
		})
		t.AssertNil(err)

		client, err := agollox.NewClient(ctx, config)
		t.AssertNil(err)
		t.Assert(client.Contains("key"), true)
		t.Assert(client.Get("key"), "value")
		t.Assert(client.Map()["key"], "value")

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

func Test_Error(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		config := agollox.DefaultConfig()
		config.AppID = "test"
		_ = agollox.LoadConfig(ctx, config, "testdata/config.none")
		_, err := agollox.NewClient(ctx, config)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "Apollo IP field is required")

		config = agollox.DefaultConfig()
		err = agollox.LoadConfig(ctx, config, "testdata/config.error")
		t.AssertNE(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		config := agollox.DefaultConfig()
		config.AppID = "test"
		config.MustStart = true
		config.IP = "http://localhost:8888"
		_, err := agollox.NewClient(ctx, config)
		t.AssertNE(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_GCFG_PATH", "errorpath")
		defer func() { _ = genv.Remove("GF_GCFG_PATH") }()
		config := agollox.DefaultConfig()
		err := agollox.LoadConfig(ctx, config, "errorconfig")
		t.AssertNE(err, nil)
		_, err = agollox.MockServer(config, "errormock")
		t.AssertNE(err, nil)
	})
}
