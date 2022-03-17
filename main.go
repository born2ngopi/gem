// this game engine build with GoLang
// this project for research can golang be used to build a game engine?
// I focusing using standard library
package gem

import "github.com/born2ngopi/gem/core"

type Config struct {
	Name    string
	Version string
}

func NewEngine(config Config) core.Engine {
	return core.Engine{
		Name:    config.Name,
		Version: config.Version,
	}
}
