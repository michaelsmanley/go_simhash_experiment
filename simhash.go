package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-dedup/simhash/sho"
	_ "github.com/lib/pq"
)

var (
	dbhost, dbpass, dbname, dbport, dbuser string
	needle                                 uint64
	distance                               uint
	since                                  int
)

func init() {
	flag.Uint64Var(&needle, "n", 0, "hash to find (needle) as integer")
	flag.UintVar(&distance, "d", 3, "hamming distance; default 3")
	flag.IntVar(&since, "s", 10000, "window in days; leave blank for all")

	dbhost = os.Getenv("SHE_dbhost")
	dbpass = os.Getenv("SHE_dbpass")
	dbname = os.Getenv("SHE_dbname")
	dbport = os.Getenv("SHE_dbport")
	dbuser = os.Getenv("SHE_dbuser")
}

func main() {
	flag.Parse()

	if needle == 0 {
		log.Fatal("needle must be non-blank")
	}

	fmt.Printf("Looking for hash %d (%x), distance %d, since %d days ago\n", needle, needle, distance, since)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require",
		dbuser, dbpass, dbhost, dbport, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("open: " + err.Error())
	}
	defer db.Close()

	q := fmt.Sprintf("SELECT sim_hash FROM content_msm WHERE created_at >= (now() - interval '%d days')", since)
	start_db_load := time.Now()
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal("query: " + err.Error())
	}
	defer rows.Close()

	var h uint64

	haystack := sho.NewOracle()

	hash_count := 0
	for rows.Next() {
		err := rows.Scan(&h)
		if err != nil {
			log.Fatal(err)
		}
		haystack.See(h)
		hash_count += 1
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	elapsed_db_load := time.Since(start_db_load)
	fmt.Printf("%d hashes returned\n", hash_count)
	fmt.Printf("load took %s\n", elapsed_db_load)

	start_search := time.Now()
	n := haystack.Search(needle, uint8(distance))
	elapsed_search := time.Since(start_search)
	if len(n) > 0 {
		fmt.Printf("%d similiar found for %x.\n", len(n), needle)
	} else {
		fmt.Println("no matches found")
	}
	fmt.Printf("hash search took %s\n", elapsed_search)

}
