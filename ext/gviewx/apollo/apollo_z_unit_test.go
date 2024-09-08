package apollo_test

import (
	"context"
	"fmt"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/CharLemAznable/gfx/ext/gviewx/apollo"
	"github.com/CharLemAznable/gfx/frame/gx"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

var (
	ctx = context.TODO()
)

func Test_ApolloAdapter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_APOLLO_CONFIG_FILE", "testdata/config.none")
		_ = genv.Set("GF_APOLLO_APPID", "test")
		defer func() { _ = genv.Remove("GF_APOLLO_CONFIG_FILE", "GF_APOLLO_APPID") }()
		_, err := apollo.LoadAdapter(ctx)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "create agollo client failed: Apollo IP field is required")
	})
	gtest.C(t, func(t *gtest.T) {
		mockConfig := agollox.DefaultConfig()
		mockConfig.AppID = "test"
		mockServer, _ := agollox.MockServer(mockConfig, "testdata/mockdata")
		defer func() { _ = mockServer.Shutdown() }()
		mockIP := fmt.Sprintf(`http://127.0.0.1:%d`, mockServer.GetListenedPort())

		_ = genv.Set("GF_APOLLO_APPID", "test")
		_ = genv.Set("GF_APOLLO_IP", mockIP)
		defer func() { _ = genv.Remove("GF_APOLLO_APPID", "GF_APOLLO_IP") }()

		adapter, err := apollo.LoadAdapter(ctx)
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
	})
}
