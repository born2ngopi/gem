package core

type Engine struct {
	Name    string
	Version string
}

func (e *Engine) String() string {
	return e.Name + " " + e.Version
}
