package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var sentences = []string{
	"The cat sat on the mat.",
	"A big dog barked loudly.",
	"I can jump very high.",
	"The sun is very bright.",
	"We play ball in the park.",
	"My fish swims in a tank.",
	"I love to draw and paint.",
	"The cake was very yummy.",
	"Ducks quack and swim well.",
	"Look at the moon and stars.",
	"I have a red and blue coat.",
	"We go to school by bus.",
	"My best friend is very kind.",
	"Birds fly high in the sky.",
	"I like to eat sweet grapes.",
	"Rabbits have long, soft ears.",
	"A frog jumps into the pond.",
	"The train goes choo choo.",
	"She has a pink and white hat.",
	"We saw a big elephant at the zoo.",
}

func sentenceHandler(w http.ResponseWriter, r *http.Request) {
	index, err := strconv.Atoi(r.URL.Query().Get("index"))
	if err != nil || index < 0 || index >= len(sentences) {
		index = 0
	}
	nextIndex := (index + 1) % len(sentences)
	response := map[string]string{
		"sentence":  sentences[index],
		"nextIndex": strconv.Itoa(nextIndex),
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func clipsHandler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./static/clips")
	if err != nil {
		http.Error(w, "Error reading clips directory", http.StatusInternalServerError)
		return
	}

	clips := []string{}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".mp4" {
			clips = append(clips, file.Name())
		}
	}

	jsonResponse, err := json.Marshal(clips)
	if err != nil {
		http.Error(w, "Error generating clips response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/sentence", sentenceHandler)
	http.HandleFunc("/clips", clipsHandler)
	log.Println("Addy Gaming → http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
