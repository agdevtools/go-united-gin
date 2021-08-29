package main
import (
"log"
"github.com/gin-gonic/gin"
"go-pg-gin/routes"
config "go-pg-gin/configs"
)
func main() {
config.Connect()
// Init Router
router := gin.Default()
// Route Handlers / Endpoints
routes.Routes(router)
log.Fatal(router.Run(":4747"))
}
