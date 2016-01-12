package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cactus/go-statsd-client/statsd"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	prefix := "business"
	mysqlDSN := "root:root@tcp(127.0.0.1:3306)/mysql"
	statsdURL := "127.0.0.1:8125"
	query := ""

	flag.StringVar(&prefix, "prefix", "", "prefix to send to statsd, usually server hostname")
	flag.StringVar(&mysqlDSN, "mysql", "root:root@tcp(127.0.0.1:3306)/mysql", "mysql to connect to")
	flag.StringVar(&statsdURL, "statsd", "127.0.0.1:8125", "statsd to send metric to")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: mysql-query-metric-statsd QUERY METRIC_NAME")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	query = flag.Arg(0)
	metricName := flag.Arg(1)

	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		log.Fatal("Error connecting to mysql:", err)
	}

	log.Println("Executing mysql query:", query)
	var value int64
	if err := db.QueryRow(query).Scan(&value); err != nil {
		log.Fatal("Error executing mysql query:", err)
	}

	sc, err := statsd.NewClient(statsdURL, prefix)
	if err != nil {
		log.Fatal("Error creating statsd client:", err)
	}
	defer sc.Close()

	sc.Gauge(metricName, value, 1.0)

	log.Printf("Sent metric prefix=%v name=%v value=%v", prefix, metricName, value)
}
