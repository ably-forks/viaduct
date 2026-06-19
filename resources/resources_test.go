package resources

import "github.com/ably-forks/viaduct"

var testLogger *viaduct.Logger

func init() {
	viaduct.Cli.SetSilent()
	testLogger = viaduct.NewLogger("Test", "Testing")
}
