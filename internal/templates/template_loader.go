package templates

import (
	"errors"
	"html/template"
	"sync"

	"github.com/glebateee/basic/internal/config"
)

var once = sync.Once{}

func LoadTemplates(cfg config.Config) error {
	var err error
	path, ok := cfg.GetString("templates:path")
	if !ok {
		return errors.New("cannot load template config")
	}
	reload := cfg.GetBoolDefault("templates:reload", false)
	once.Do(func() {
		doLoad := func() *template.Template {
			t := template.New("localt")
			t.Funcs(template.FuncMap{
				"body":   func() string { return "" },
				"layout": func() string { return "" },
			})
			t, err = t.ParseGlob(path)
			return t
		}
		if reload {
			getTemplates = doLoad
		} else {
			templates := doLoad()
			getTemplates = func() *template.Template {
				t, _ := templates.Clone()
				return t
			}

		}
	})
	return err
}
