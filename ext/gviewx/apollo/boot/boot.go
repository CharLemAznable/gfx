package boot

import (
	"github.com/CharLemAznable/gfx/ext/gviewx/apollo"
	"github.com/CharLemAznable/gfx/frame/gx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	ctx := gctx.GetInitCtx()
	adapter, err := apollo.LoadAdapter(ctx)
	if err != nil {
		g.Log().Fatalf(ctx, `%+v`, err)
	}
	gx.ViewX().SetAdapter(adapter)
}
