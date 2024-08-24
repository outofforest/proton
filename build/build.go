package build

import (
	"context"

	"github.com/outofforest/build"
	"github.com/outofforest/buildgo"
)

// Setup configures the environment.
func Setup(_ context.Context, deps build.DepsFunc) error {
	deps(buildgo.InstallAll)
	return nil
}
