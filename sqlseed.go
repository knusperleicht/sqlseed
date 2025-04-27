package sqlseed

import (
	"github.com/knusperleicht/sqlseed/lib/sqlite"
	"log/slog"
)

func Seed() {
	slog.Info("sqlseed")
	sqlite.Init()
}
