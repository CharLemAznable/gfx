package boot

import (
	"fmt"
	"github.com/CharLemAznable/gfx/ext/gviewx/apollo"
	"github.com/CharLemAznable/gfx/frame/gx"
	"github.com/CharLemAznable/gfx/os/gviewx"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	ctx := gctx.GetInitCtx()
	adapter, err := apollo.LoadAdapter(ctx)
	if err == nil {
		origin := gx.ViewX().GetAdapter()
		comb := gviewx.NewAdapterDefault(adapter, origin)
		gx.ViewX().SetAdapter(comb)
	} else {
		fmt.Printf("load apollo config: %+v\n", err)
	}
}
