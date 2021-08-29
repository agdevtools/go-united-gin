package controllers
import (
"net/http"
"github.com/gin-gonic/gin"
"log"
"strconv"
"time"
configs "go-pg-gin/configs"
)
type Todo struct {
ID        string    `json:"id"`
Title     string    `json:"title"`
Body      string    `json:"body"`
Completed string      `json:"completed"`
CreatedAt time.Time `json:"created_at"`
UpdatedAt time.Time `json:"updated_at"`
}

type Player struct {
    ID       int64  `json:"id"`
    PLAYER_ID      string  `json:"player_id"`
    Player_name     string `json:"player_name"`
    Player_position string `json:"player_position"`
}

func GetAllPlayers(c *gin.Context) {
var team []Player
db := configs.Connect()
err := db.Model(&team).Select()
if err != nil {
log.Printf("Error while getting all players, Reason: %v\n", err)
c.JSON(http.StatusInternalServerError, gin.H{
"status":  http.StatusInternalServerError,
"message": "Something went wrong",
})
return
}
c.JSON(http.StatusOK, gin.H{
"status":  http.StatusOK,
"message": "All Players",
"data": team,
})
return
}

func GetSinglePlayer(c *gin.Context) {
db := configs.Connect()

// Id1  := c.Param("Id")
Id2, _ := strconv.Atoi(c.Param("Id"))
Id  := int64(Id2) // safe
player := &Player{ID: Id}
err := db.Select(player)

if err != nil {
log.Printf("Error while getting a single todo, Reason: %v\n", err)
c.JSON(http.StatusNotFound, gin.H{
"status":  http.StatusNotFound,
"message": "Player not found",
})
return
}
c.JSON(http.StatusOK, gin.H{
"status":  http.StatusOK,
"message": "Single Player",
"data": player,
})
return
}


func CreatePlayer(c *gin.Context) {
var player Player
c.BindJSON(&player)
player_id := player.PLAYER_ID
player_name := player.Player_name
player_position := player.Player_position

db := configs.Connect()

insertError := db.Insert(&Player{
PLAYER_ID:       player_id,
Player_name:      player_name,
Player_position:     player_position,
})
if insertError != nil {
log.Printf("Error while inserting new player into db, Reason: %v\n", insertError)
c.JSON(http.StatusInternalServerError, gin.H{
"status":  http.StatusInternalServerError,
"message": "Something went wrong",
})
return
}
c.JSON(http.StatusCreated, gin.H{
"status":  http.StatusCreated,
"message": "Player created Successfully",
})
return
}