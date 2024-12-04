package repository

import "log"

func (rep *Repository) PrepareTable() {
	_, err := rep.DB.Exec(`CREATE TABLE IF NOT EXISTS wallets (
	uuid INTEGER UNIQUE,
	balance INTEGER
	)`)
	if err != nil {
		log.Fatal(err)
	}
}
