package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"

	//to do fix error 
	"phoneNumberNormalixer/db"

	_ "github.com/lib/pq"
)

type ConnString struct {
	user     string
	password string
	port     int
	db       string
	host     string
}

func main() {
	conn := ConnString{
		user:     "postgres",
		password: "postgres",
		port:     5432,
		db:       "church",
		host:     "localhost",
	}
	DB, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", conn.host, conn.port, conn.user, conn.password))
	if err != nil {
		log.Fatal(err.Error())
	}
	db.ResetDB(DB)
}

func Normalization(phone string) string {
	var buf bytes.Buffer

	for _, ch := range phone {
		if len(buf.String()) == 0 {
			buf.Write([]byte("("))
		}
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
		if len(buf.String()) == 4 {
			buf.Write([]byte(") "))
		}
	}

	return buf.String()
}
