package agollox

import (
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"net/http"
	"net/http/httptest"
	"strings"
)

func MockServer(appConfig *Config, configMap map[string]*gmap.StrStrMap) *httptest.Server {
	uriHandlerMap := make(map[string]http.HandlerFunc, 0)
	notifications := make([]map[string]interface{}, 0)
	for namespace, keyValueMap := range configMap {
		uriHandlerMap[fmt.Sprintf("/configfiles/json/%s/%s/%s",
			appConfig.AppID, appConfig.Cluster, namespace)] =
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = fmt.Fprintf(w, gjson.New(keyValueMap).MustToJsonString())
			}
		uriHandlerMap[fmt.Sprintf("/configs/%s/%s/%s",
			appConfig.AppID, appConfig.Cluster, namespace)] =
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				m := map[string]interface{}{
					"appId":          appConfig.AppID,
					"cluster":        appConfig.Cluster,
					"namespaceName":  namespace,
					"configurations": keyValueMap,
					"releaseKey":     "",
				}
				_, _ = fmt.Fprintf(w, gjson.New(m).MustToJsonString())
			}
		notifications = append(notifications, map[string]interface{}{
			"namespaceName":  namespace,
			"notificationId": 3,
		})
	}
	uriHandlerMap["/notifications/v2"] = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, gjson.New(notifications).MustToJsonString())
	}

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for path, handler := range uriHandlerMap {
			if strings.HasPrefix(r.RequestURI, path) {
				handler(w, r)
				break
			}
		}
	}))
}
