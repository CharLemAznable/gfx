package gviewx

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/util/gutil"
)

const ConfigNodeNameViewer = "viewer"

var localInstances = gmap.NewStrAnyMap(true)

func Instance(name ...string) *View {
	instanceName := gview.DefaultName
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	return localInstances.GetOrSetFuncLock(instanceName, func() interface{} {
		view := New()
		tryConfigView(context.Background(), view, instanceName)
		return view
	}).(*View)
}

func tryConfigView(ctx context.Context, view *View, instanceName string) {
	if !g.Config().Available(ctx) {
		return
	}
	configMap, err := g.Config().Data(ctx)
	if err != nil {
		g.Log().Errorf(ctx, `retrieve config data map failed: %+v`, err)
	}
	configNodeName := ConfigNodeNameViewer
	if len(configMap) > 0 {
		if v, _ := gutil.MapPossibleItemByKey(configMap, ConfigNodeNameViewer); v != "" {
			configNodeName = v
		}
	}
	configMap = g.Config().MustGet(ctx, fmt.Sprintf(`%s.%s`, configNodeName, instanceName)).Map()
	if len(configMap) == 0 {
		configMap = g.Config().MustGet(ctx, configNodeName).Map()
	}
	if len(configMap) == 0 {
		return
	}
	if err = view.SetConfigWithMap(configMap); err != nil {
		panic(err)
	}
}
