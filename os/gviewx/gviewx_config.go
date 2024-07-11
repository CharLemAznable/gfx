package gviewx

import (
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

type Config struct {
	Data        map[string]interface{} `json:"data"`
	Delimiters  []string               `json:"delimiters"`
	AutoEncode  bool                   `json:"autoEncode"`
	I18nManager *gi18n.Manager         `json:"-"`
}

func (view *View) SetConfig(config Config) *View {
	if len(config.Data) > 0 {
		view.Assigns(config.Data)
	}
	if len(config.Delimiters) > 1 {
		view.SetDelimiters(config.Delimiters[0], config.Delimiters[1])
	}
	view.SetAutoEncode(config.AutoEncode)
	if config.I18nManager != nil {
		view.SetI18n(config.I18nManager)
	}
	return view
}

func (view *View) SetConfigWithMap(m map[string]interface{}) error {
	m = gutil.MapCopy(m)
	config := &Config{}
	err := gconv.Struct(m, config)
	if err != nil {
		return err
	}
	view.SetConfig(*config)
	return nil
}

func (view *View) Assigns(data gview.Params) {
	view.view.Assigns(data)
}

func (view *View) SetDelimiters(left, right string) {
	view.view.SetDelimiters(left, right)
}

func (view *View) SetAutoEncode(enable bool) {
	view.view.SetAutoEncode(enable)
}

func (view *View) SetI18n(manager *gi18n.Manager) {
	view.view.SetI18n(manager)
}
