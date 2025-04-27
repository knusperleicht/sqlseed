package sqlite

import (
	"log/slog"
	"sqlseed/lib/fixtures"
)

func init() {
	slog.Info("sqlite driver")
	fixtures.Init()
}
