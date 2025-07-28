package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Names struct {
	FirstNames struct {
		Male   []string `json:"male"`
		Female []string `json:"female"`
	} `json:"firstName"`
	LastNames []string `json:"lastNames"`
}

type PersonalityData struct {
	PersonalityTypes []struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"personalityType"`
	Alignment []string `json:"alignment"`
}

//go:embed names.json
var namesFile []byte

//go:embed randomizer.json
var randomizerFile []byte

var namesData Names
var personalityData PersonalityData

func init() {
	rand.Seed(time.Now().UnixNano()) // Initialize random number generator

	// Load names from embedded names.json
	err := json.Unmarshal(namesFile, &namesData)
	if err != nil {
		fmt.Printf("Error unmarshaling names.json: %v\n", err)
		return
	}

	// Load personality data from embedded randomizer.json
	err = json.Unmarshal(randomizerFile, &personalityData)
	if err != nil {
		fmt.Printf("Error unmarshaling randomizer.json: %v\n", err)
		return
	}
}

func main() {
	// Setup CORS for all handlers
	http.HandleFunc("/", setupCORS(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	}))
	http.HandleFunc("/name", setupCORS(nameHandler))
	http.HandleFunc("/personality", setupCORS(personalityHandler))
	http.HandleFunc("/character", setupCORS(characterHandler))

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// setupCORS adds CORS headers to the response
func setupCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler(w, r)
	}
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	if len(namesData.FirstNames.Male) == 0 && len(namesData.FirstNames.Female) == 0 && len(namesData.LastNames) == 0 {
		http.Error(w, "Names data not loaded or empty", http.StatusInternalServerError)
		return
	}

	// Generate random sex
	sexes := []string{"Male", "Female"}
	randomSex := sexes[rand.Intn(len(sexes))]

	// Generate first name based on sex
	var firstName string
	if randomSex == "Male" {
		if len(namesData.FirstNames.Male) == 0 {
			http.Error(w, "No Male first names available", http.StatusInternalServerError)
			return
		}
		firstName = namesData.FirstNames.Male[rand.Intn(len(namesData.FirstNames.Male))]
	} else {
		if len(namesData.FirstNames.Female) == 0 {
			http.Error(w, "No Female first names available", http.StatusInternalServerError)
			return
		}
		firstName = namesData.FirstNames.Female[rand.Intn(len(namesData.FirstNames.Female))]
	}

	if len(namesData.LastNames) == 0 {
		http.Error(w, "No last names available", http.StatusInternalServerError)
		return
	}
	lastName := namesData.LastNames[rand.Intn(len(namesData.LastNames))]

	response := map[string]string{
		"firstName": firstName,
		"lastName":  lastName,
		"sex":       randomSex,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func personalityHandler(w http.ResponseWriter, r *http.Request) {
	if len(personalityData.PersonalityTypes) == 0 && len(personalityData.Alignment) == 0 {
		http.Error(w, "Personality data not loaded or empty", http.StatusInternalServerError)
		return
	}

	randomPersonalityType := personalityData.PersonalityTypes[rand.Intn(len(personalityData.PersonalityTypes))]
	randomAlignment := personalityData.Alignment[rand.Intn(len(personalityData.Alignment))]

	response := map[string]interface{}{
		"personalityType": randomPersonalityType,
		"alignment":       randomAlignment,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func characterHandler(w http.ResponseWriter, r *http.Request) {
	// Generate random sex
	sexes := []string{"Male", "Female"}
	randomSex := sexes[rand.Intn(len(sexes))]

	// Generate first name based on sex
	var firstName string
	if randomSex == "Male" {
		if len(namesData.FirstNames.Male) == 0 {
			http.Error(w, "No Male first names available", http.StatusInternalServerError)
			return
		}
		firstName = namesData.FirstNames.Male[rand.Intn(len(namesData.FirstNames.Male))]
	} else {
		if len(namesData.FirstNames.Female) == 0 {
			http.Error(w, "No Female first names available", http.StatusInternalServerError)
			return
		}
		firstName = namesData.FirstNames.Female[rand.Intn(len(namesData.FirstNames.Female))]
	}

	// Generate last name
	if len(namesData.LastNames) == 0 {
		http.Error(w, "No last names available", http.StatusInternalServerError)
		return
	}
	lastName := namesData.LastNames[rand.Intn(len(namesData.LastNames))]

	// Generate personality
	if len(personalityData.PersonalityTypes) == 0 || len(personalityData.Alignment) == 0 {
		http.Error(w, "Personality data not loaded or empty", http.StatusInternalServerError)
		return
	}
	randomPersonalityType := personalityData.PersonalityTypes[rand.Intn(len(personalityData.PersonalityTypes))]
	randomAlignment := personalityData.Alignment[rand.Intn(len(personalityData.Alignment))]

	response := map[string]interface{}{
		"sex":             randomSex,
		"firstName":       firstName,
		"lastName":        lastName,
		"personalityType": randomPersonalityType,
		"alignment":       randomAlignment,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
