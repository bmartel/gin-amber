package ginamber

import (
	"html/template"

	"github.com/eknkc/amber"
	"github.com/gin-gonic/gin/render"
)

// NewViewRenderer ...
func NewViewRenderer(dir string, ext string, funcs *template.FuncMap) *ViewRenderer {
	options := amber.DirOptions{}
	options.Recursive = true
	options.Ext = ext

	if funcs != nil {
		tmplFuncs := *funcs
		for k, v := range amber.FuncMap {
			tmplFuncs[k] = v
		}
		amber.FuncMap = tmplFuncs
	}

	templates, err := amber.CompileDir(dir, options, amber.DefaultOptions)
	if err != nil {
		panic(err)
	}

	return &ViewRenderer{
		templates: templates,
	}
}

// ViewRenderer ...
type ViewRenderer struct {
	templates map[string]*template.Template
}

// Add ...
func (r *ViewRenderer) Add(name string, tmpl *template.Template) {
	if r.templates == nil {
		r.templates = make(map[string]*template.Template)
	}
	r.templates[name] = tmpl
}

// Instance ...
func (r *ViewRenderer) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r.templates[name],
		Data:     data,
	}
}
