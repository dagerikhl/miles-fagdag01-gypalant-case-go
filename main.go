package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// character represents data about a character.
type character struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	InPlay       bool   `json:"inPlay"`
	Strength     int    `json:"strength"`
	Intelligence int    `json:"intelligence"`
	Charisma     int    `json:"charisma"`
}

// characters slice to seed character data.
var characters = []character{
	{ID: "1", Name: "Kelsier", InPlay: true, Strength: 16, Intelligence: 9, Charisma: 16},
	{ID: "2", Name: "Shallan Devar", InPlay: true, Strength: 15, Intelligence: 17, Charisma: 9},
	{ID: "3", Name: "Jarl Berserkarson", InPlay: false, Strength: 18, Intelligence: 3, Charisma: 18},
}

// To run this Go application, run:
// 1. `go get .` (fetched dependencies).
// 2. `go run .` (starts the server).
// 3. Test the server with one of the examples below.
func main() {
	router := gin.Default()
	// Test with: curl http://localhost:8080/characters
	router.GET("/characters", getCharacters)
	// Test with: curl http://localhost:8080/characters --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4", "name": "Kaladin Stormblessed", "inPlay": true, "strength": 16, "intelligence": 14, "charisma": 15}'
	router.POST("/characters", postCharacters)
	// Test with: curl http://localhost:8080/characters/3
	router.GET("/characters/:id", getCharacterByID)

	router.Run("localhost:8080")
}

// getCharacters responds with the list of all characters as JSON.
func getCharacters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, characters)
}

// postCharacters adds an character from JSON received in the request body.
func postCharacters(c *gin.Context) {
	var newCharacter character

	// Call BindJSON to bind the received JSON to newCharacter.
	if err := c.BindJSON(&newCharacter); err != nil {
		return
	}

	// Add the new character to the slice.
	characters = append(characters, newCharacter)
	c.IndentedJSON(http.StatusCreated, newCharacter)
}

// getCharacterByID locates the character whose ID value matches the id parameter sent by the client, then returns that character as a response.
func getCharacterByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of characters, looking for an character whose ID value matches the parameter.
	for _, character := range characters {
		if character.ID == id {
			c.IndentedJSON(http.StatusOK, character)

			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Character not found."})
}
