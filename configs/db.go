package config
import (
"log"
"os"
"github.com/go-pg/pg/v9"
_ "github.com/lib/pq"      // postgres golang driver
)

// Connecting to db
func Connect() *pg.DB {
opts := &pg.Options{
User: "izwkmryg",
Password: "nE7cHfwNbpL4T2zv_3wI1u_5JxDHB0Cr",
Addr: "tai.db.elephantsql.com:5432",
Database: "izwkmryg",
}
var db *pg.DB = pg.Connect(opts)
if db == nil {
log.Printf("Failed to connect")
os.Exit(100)
}
log.Printf("Connected to elephant db")
return db
}

