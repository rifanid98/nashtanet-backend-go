package database

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string
}

func NewConfigPostgres() *config  {
	return &config{
		host:     "localhost",
		database: "okedok_golang",
		port:     "5432",
		driver:   "postgres",
		user:     "postgres",
		password: "postgres",
	}
}