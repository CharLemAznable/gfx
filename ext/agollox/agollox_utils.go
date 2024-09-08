package agollox

import "github.com/gogf/gf/v2/util/gutil"

func MapOmitNil(data map[string]interface{}) map[string]interface{} {
	if len(data) == 0 {
		return data
	}
	for k, v := range data {
		if v == nil {
			delete(data, k)
		}
	}
	return data
}

func MapFill(dst map[string]interface{}, src map[string]interface{}) {
	for key, value := range src {
		if !gutil.MapContainsPossibleKey(dst, key) {
			dst[key] = value // fill dst map with src map
		}
	}
}
