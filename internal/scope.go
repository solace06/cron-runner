package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/solace06/cron-runner/database"
	"github.com/solace06/cron-runner/job/config"
)

type Scope struct {
	DB  *database.DB
	Cfg *config.Config
	Log *log.Logger
}

func NewScope() (*Scope, error) {
	//load the config file
	cfg := config.MustLoad()

	//set up the database
	db, err := database.NewDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	//set up the logger
	logger:=log.New(os.Stdout, "[APP] ", log.LstdFlags)

	return &Scope{
		DB: db,
		Cfg: cfg,
		Log: logger,
	}, nil
}
