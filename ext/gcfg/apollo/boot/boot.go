package boot

import (
	"fmt"
	"github.com/CharLemAznable/gfx/ext/gcfg/apollo"
	"github.com/CharLemAznable/gfx/os/gcfgx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	ctx := gctx.GetInitCtx()
	adapter, err := apollo.LoadAdapter(ctx)
	if err == nil {
		origin := g.Cfg().GetAdapter()
		comb := gcfgx.NewAdapterDefault(adapter, origin)
		g.Cfg().SetAdapter(comb)
	} else {
		fmt.Printf("load apollo config: %+v\n", err)
	}
}
