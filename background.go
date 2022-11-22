package backgroundtask

import (
	"go.k6.io/k6/js/modules"

	"github.com/pedroyremolo/xk6-background-task/cron"
)

// Register the extensions on module initialization.
func init() {
	modules.Register("k6/x/background_task/cron", cron.New())
}
