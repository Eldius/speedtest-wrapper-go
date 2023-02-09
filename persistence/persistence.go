package persistence

import (
	"github.com/Eldius/speedtest-wrapper-go/speedtest"
	"github.com/asdine/storm/v3"
	_ "github.com/asdine/storm/v3"
)

type Persistence struct {
	db *storm.DB
}

func NewPersistence(f string) (*Persistence, error) {
	db, err := storm.Open(f)
	if err != nil {
		return nil, err
	}

	return &Persistence{
		db: db,
	}, err
}

func (p *Persistence) Persist(t speedtest.SpeedtestResult) error {
	return p.db.Save(&t)
}
