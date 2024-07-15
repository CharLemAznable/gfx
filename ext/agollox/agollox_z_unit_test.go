package agollox_test

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
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

		mockData := map[string]map[string]string{
			"application": {"key": "value"},
		}
		mockConfig := agollox.DefaultConfig()
		mockConfig.AppID = "test"
		mockServer := agollox.MockServer(mockConfig, mockData)
		mockIP := mockServer.URL

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

		client, err := agollox.NewClient(config)
		t.AssertNil(err)
		t.Assert(true, client.Contains("key"))
		t.Assert("value", client.Get("key"))
		t.Assert("value", client.Map()["key"])

		client.SetOnChangeFn(func(event *agollox.ChangeEvent) {
			_, ok := event.Changes["key"]
			t.Assert(true, ok)
		})
		mockData["application"]["key"] = "new value"
		time.Sleep(time.Second * 3)
		t.Assert("new value", client.Get("key"))
	})
}

func Test_Error(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		config := agollox.DefaultConfig()
		config.MustStart = true
		_ = agollox.LoadConfig(ctx, config, "testdata/config.none")
		_, err := agollox.NewClient(config)
		t.AssertNE(err, nil)
	})
}
