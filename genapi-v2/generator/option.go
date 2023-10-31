package generator

type OptionFunc func(*Generator)

func ConfigParser(parser func(text string) map[string]interface{}) OptionFunc {
	return func(g *Generator) {
		g.Parser = parser
	}
}

func ConfigExecutor(e ExecutorFunc) OptionFunc {
	return func(g *Generator) {
		g.Executor = e
	}
}
