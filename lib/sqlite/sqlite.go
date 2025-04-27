package sqlite

import (
	"github.com/knusperleicht/sqlseed/lib/fixtures"
	"log/slog"
)

func Init() {
	slog.Info("sqlite driver")
	fixtures.Init()
}
