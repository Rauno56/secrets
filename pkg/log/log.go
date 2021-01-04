// `log` is a singleton for conveniently handling debug output
package log

import (
	"jobbatical/secrets/pkg/options"
	"jobbatical/secrets/pkg/utils"
)

var PrintDebugln = utils.NoopDebugln

func init() {
	if options.Verbose {
		PrintDebugln = utils.ErrPrintln
	}
}
