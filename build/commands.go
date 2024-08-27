package build

import "github.com/outofforest/build/v2/pkg/types"

// Commands is a list of commands.
var Commands = map[string]types.Command{
	"generate": {Fn: generate, Description: "Generates proton code for tests"},
}
