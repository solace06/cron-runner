package job

import (
	"fmt"

	"github.com/solace06/cron-runner/database"
	"github.com/solace06/cron-runner/job/config"
)

type Scope struct {
	db  *database.DB
	Cfg *config.Config
}

func NewScope() (*Scope, error) {
	//load the config file
	cfg := config.MustLoad()

	//set up the database
	db, err := database.NewDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	return &Scope{
		db: db,
		Cfg: cfg,
	}, nil
}

func (s *Scope) Migrate(){
	s.db.Migrate()
}

func (s *Scope) Close() error{
	return s.db.Close()
}