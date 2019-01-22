package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-dedup/simhash"
	"github.com/tjarratt/babble"
	//_ "github.com/lib/pq"
)

var (
	dbhost, dbpass, dbname, dbport, dbuser string
	count                                  uint64
)

func init() {
	flag.Uint64Var(&count, "c", 0, "number of simhashes to create")

	dbhost = os.Getenv("SHE_dbhost")
	dbpass = os.Getenv("SHE_dbpass")
	dbname = os.Getenv("SHE_dbname")
	dbport = os.Getenv("SHE_dbport")
	dbuser = os.Getenv("SHE_dbuser")
}

func main() {
	flag.Parse()

	if count == 0 {
		log.Fatal("count must be >= 1")
	}

	fmt.Printf("Generating %d hashes\n", count)

	sh := simhash.NewSimhash()
	start_db_load := time.Now()

	babbler := babble.NewBabbler()
	babbler.Count = 10
	babbler.Separator = " "

	var i uint64
	for i = 1; i <= count; i++ {
		words := babbler.Babble()
		h := sh.GetSimhash(sh.NewWordFeatureSet([]byte(words)))
		fmt.Printf("Simhash %d: %s %x\n", i, words, h)
	}

	elapsed_db_load := time.Since(start_db_load)
	fmt.Printf("load took %s\n", elapsed_db_load)

	// connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require",
	// 	dbuser, dbpass, dbhost, dbport, dbname)
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal("open: " + err.Error())
	// }
	// defer db.Close()

	// q := fmt.Sprintf("SELECT sim_hash FROM content_msm WHERE created_at >= (now() - interval '%d days')", since)
	// rows, err := db.Query(q)
	// if err != nil {
	// 	log.Fatal("query: " + err.Error())
	// }
	// defer rows.Close()
}
