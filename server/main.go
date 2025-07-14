package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Player holds the data for a player's quiz result.
type Player struct {
	Name  string `bson:"name"`
	Score int    `bson:"score"`
}

func mongoConnect() *mongo.Client {
	uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	return client
}

var correctAnswers = []int{1, 2, 1, 2, 1, 1, 2, 2, 1, 2, 1, 1, 2, 1, 1, 1, 1, 1, 0, 1}

type ScoreRequest struct {
	Name    string `json:"name"`
	Answers []int  `json:"answers"`
}

type ScoreResponse struct {
	Score int `json:"score"`
}

// App holds application-wide dependencies.
type App struct {
	resultsCollection *mongo.Collection
}

func (app *App) scoreHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req ScoreRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received score request for user: %s", req.Name)
	log.Println(req)

	score := 0
	for i := 0; i < len(req.Answers) && i < len(correctAnswers); i++ {
		if req.Answers[i] == correctAnswers[i] {
			score++
		}
	}

	playerResult := Player{Name: req.Name, Score: score}

	insertResult, err := app.resultsCollection.InsertOne(context.TODO(), playerResult)
	if err != nil {
		log.Printf("Error inserting score into database: %v", err)
		http.Error(w, "Failed to save score", http.StatusInternalServerError)
		return
	}

	log.Printf("Inserted score for %s with ID: %v", req.Name, insertResult.InsertedID)

	resp := ScoreResponse{Score: score}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	client := mongoConnect()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		log.Println("Disconnected from MongoDB.")
	}()

	app := &App{
		resultsCollection: client.Database("goquiz").Collection("results"),
	}

	http.HandleFunc("/score", app.scoreHandler)
	log.Println("Backend server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
