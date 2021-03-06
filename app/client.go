package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	db "github.com/hugomcfonseca/mysql-tasker/app/databases"
)

var (
	action = flag.String("action", "", "Action")
	dbName = flag.String("database", "", "Database name")
	dbHost = flag.String("host", "localhost", "Hostname, or IP, of target database server")
	dbPort = flag.Int("port", 3306, "Database port")
	dbUser = flag.String("user", "my_user", "Database user")
	dbPass = flag.String("password", "", "Database user's password")
)

func main() {
	var err error

	flag.Parse()

	if *action == "" {
		os.Exit(1)
	}

	dsnConn := buildDatasourceName()

	switch *action {
	case "createDB":
		err = db.NewDB(*dbName, dsnConn)
	case "deleteDB":
		err = db.RemoveDB(*dbName, dsnConn)
	default:
		return
	}

	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

func buildDatasourceName() string {
	if os.Getenv("DB_HOST") != "" {
		*dbHost = os.Getenv("DB_HOST")
	}

	if os.Getenv("DB_PORT") != "" {
		*dbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	}

	if os.Getenv("DB_NAME") != "" {
		*dbName = os.Getenv("DB_NAME")
	}

	if os.Getenv("DB_USER") != "" {
		*dbUser = os.Getenv("DB_USER")
	}

	if os.Getenv("DB_PASS") != "" {
		*dbPass = os.Getenv("DB_PASS")
	}

	if *dbHost == "" || (*dbPort <= 1024 || *dbPort >= 65535) || *dbName == "" || *dbUser == "" {
		return ""
	}

	dsnConn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", *dbUser, *dbPass, *dbHost, *dbPort)

	return dsnConn
}
