package generator

import "text/template"

type OptionFunc func(*Generator)

func ConfigParser(parser func(text string) map[string]interface{}) OptionFunc {
	return func(g *Generator) {
		g.Parser = parser
	}
}

func ConfigTemplate(tmpl *template.Template) OptionFunc {
	return func(g *Generator) {
		g.Template = tmpl
	}
}
