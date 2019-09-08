package config

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	DB     *sqlx.DB
	Router *mux.Router
}
