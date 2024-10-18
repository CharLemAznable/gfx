package gviewx

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
)

type AdapterFile struct {
	searchPaths *garray.StrArray
}

const (
	DefaultTemplatePath = "template"
)

func NewAdapterFile(templatePath ...string) *AdapterFile {
	tmplPath := DefaultTemplatePath
	if len(templatePath) > 0 && templatePath[0] != "" {
		tmplPath = templatePath[0]
	}
	searchPaths := garray.NewStrArraySize(0, 3, true)
	addSearchPath(searchPaths, gfile.Pwd(), tmplPath)
	addSearchPath(searchPaths, gfile.MainPkgPath(), tmplPath)
	addSearchPath(searchPaths, gfile.SelfDir(), tmplPath)
	return &AdapterFile{
		searchPaths: searchPaths,
	}
}

func (a *AdapterFile) GetContent(key string) (content string, err error) {
	a.searchPaths.RLockFunc(func(array []string) {
		for _, searchPath := range array {
			filePath := gfile.Join(searchPath, key)
			if gfile.IsFile(filePath) {
				content = gfile.GetContentsWithCache(filePath)
				return
			}
		}
		err = gerror.NewCodef(gcode.CodeInvalidParameter, `template file "%s" not found`, key)
	})
	return
}

func addSearchPath(paths *garray.StrArray, rootPath string, tmplPath string) {
	if rootPath != "" && gfile.Exists(rootPath) {
		joinPath := gfile.Join(rootPath, tmplPath)
		if gfile.IsDir(joinPath) {
			paths.Append(joinPath)
		}
	}
}
