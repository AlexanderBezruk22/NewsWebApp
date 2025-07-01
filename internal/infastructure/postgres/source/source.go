package source

import "fmt"

type Source struct {
	host     string
	port     string
	dbname   string
	user     string
	password string
}

func New(host, port, dbname, user, password string) *Source {
	return &Source{
		host:     host,
		port:     port,
		dbname:   dbname,
		user:     user,
		password: password,
	}
}

func (s Source) Connection() string {
	format := "host=%s port=%s dbname=%s user=%s password=%s sslmode=disable"

	return fmt.Sprintf(format, s.host, s.port, s.dbname, s.user, s.password)
}
