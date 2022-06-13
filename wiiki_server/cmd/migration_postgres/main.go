package main

import (
	"flag"
	"log"
	"os"
	_ "wiiki_server/infra/postgres/migrations"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", ".", "directory with migration files")
)

func main() {

	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 2 {
		flags.Usage()
		return
	}

	dbName, command := args[1], args[2]

	db, err := goose.OpenDBWithDriver("postgres", dbName)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	var arguments []string
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	err = goose.Run(command, db, *dir, arguments...)
	if err != nil {
		log.Fatalf("goose %v: %v\n", command, err)
	}
}
