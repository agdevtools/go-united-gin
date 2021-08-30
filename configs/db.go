package config
import (
"log"
"os"
"github.com/go-pg/pg/v9"
_ "github.com/lib/pq"      // postgres golang driver
)

//Connecting to db
func Connect() *pg.DB {
opts := &pg.Options{
User: "postgres",
Password: "pgpass",
Addr: "localhost:55001",
Database: "postgres",
}
var db *pg.DB = pg.Connect(opts)
if db == nil {
log.Printf("Failed to connect")
os.Exit(100)
}
log.Printf("Connected to elephant db")
return db
}

