package agollox

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/util/guid"
)

func MockServer(appConfig *Config, mockFileName string) (*ghttp.Server, error) {
	mockFile, err := gcfg.NewAdapterFile(mockFileName)
	if err != nil {
		return nil, err
	}
	server := ghttp.GetServer(guid.S())
	server.SetDumpRouterMap(false)
	server.BindHandler(fmt.Sprintf("/configfiles/json/%s/%s/:namespace",
		appConfig.AppID, appConfig.Cluster), func(r *ghttp.Request) {
		value, _ := mockFile.Get(context.Background(), r.GetRouter("namespace").String())
		r.Response.WriteJson(value)
	})
	server.BindHandler(fmt.Sprintf("/configs/%s/%s/:namespace",
		appConfig.AppID, appConfig.Cluster), func(r *ghttp.Request) {
		namespace := r.GetRouter("namespace").String()
		value, _ := mockFile.Get(context.Background(), namespace)
		r.Response.WriteJson(map[string]interface{}{
			"appId":          appConfig.AppID,
			"cluster":        appConfig.Cluster,
			"namespaceName":  namespace,
			"configurations": value,
			"releaseKey":     "",
		})
	})
	server.BindHandler("/notifications/v2", func(r *ghttp.Request) {
		mockDataMap, _ := mockFile.Data(context.Background())
		notifications := make([]map[string]interface{}, 0)
		for namespace := range mockDataMap {
			notifications = append(notifications, map[string]interface{}{
				"namespaceName":  namespace,
				"notificationId": 2,
			})
		}
		r.Response.WriteJson(notifications)
	})
	err = server.Start()
	return server, err
}
