package postgres_test

import (
	"App/internal/config"
	"App/pkg/postgres"
	"testing"
)

func TestDb(t *testing.T) {

	cfg := config.NewConfig()
	_, err := postgres.NewPosgrs(cfg.DbCfg)

	if err != nil {
		t.Error("КАЛЛЛЛ", err)
	}

	// если есть принт то всё робит

}
