package boot

import (
	"github.com/CharLemAznable/gfx/ext/gcfg/apollo"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	ctx := gctx.GetInitCtx()
	adapter, err := apollo.LoadAdapter(ctx)
	if err != nil {
		g.Log().Fatalf(ctx, `%+v`, err)
	}
	g.Cfg().SetAdapter(adapter)
}
