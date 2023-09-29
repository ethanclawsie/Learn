package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/urfave/cli"
)

type Term struct {
	Term       string
	Definition string
}

type AppConfig struct {
	LastFilePath string `json:"lastFilePath"`
}

func main() {
	app := &cli.App{
		Name:  "Learn CLI",
		Usage: "A simple Learn CLI app",
		Action: func(c *cli.Context) error {
			// Load the last used file path from the configuration file, if available.
			config := loadConfig()
			fileName := config.LastFilePath

			// Ask the user if they want to change the file location.
			changeLocation := false
			if fileName != "" {
				fmt.Printf("Current file location: %s\n", fileName)
				fmt.Print("Do you want to change the file location? (yes/no): ")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				answer := strings.ToLower(scanner.Text())
				changeLocation = answer == "yes"
			}

			// If the user wants to change the file location or there is no saved location, ask for a new one.
			if changeLocation || fileName == "" {
				fmt.Print("Enter the full path to the text file (e.g., /path/to/terms.txt): ")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				fileName = scanner.Text()
				// Save the file path to the configuration file for future use.
				config.LastFilePath = fileName
				saveConfig(config)
			}

			// Ask the user whether to learn terms in random order or not.
			fmt.Print("Do you want to learn terms in random order? (yes/no): ")
			scanner := bufio.NewScanner(os.Stdin) // Declare scanner here
			scanner.Scan() // Scan for input
			randomOrder := strings.ToLower(scanner.Text()) == "yes"

			// Load terms and definitions from the file.
			terms, err := loadTermsFromFile(fileName)
			if err != nil {
				log.Fatalf("Error loading terms from file: %v", err)
			}

			// Shuffle terms if in random order.
			if randomOrder {
				shuffle(terms)
			}

			// Divide terms into chunks of 6 (or fewer if the total number is not divisible by 6).
			chunkSize := 6
			chunks := makeChunks(terms, chunkSize)

			// Main learning loop.
			for _, chunk := range chunks {
    incorrectTerms := []Term{}
    for _, term := range chunk {
        fmt.Printf("Term: %s\nYour Answer: ", term.Term)
        scanner.Scan()  // Add this line to capture user input
        answer := scanner.Text() // Capture the user's answer
        if answer == term.Definition {
            fmt.Println("Correct!\n")
        } else {
            fmt.Printf("Wrong! The correct answer is: %s\n\n", term.Definition)
            incorrectTerms = append(incorrectTerms, term)
        }
    }

    // Re-ask only the incorrect terms.
    for len(incorrectTerms) > 0 {
        term := incorrectTerms[0]
        fmt.Printf("Term: %s\nYour Answer: ", term.Term)
        scanner.Scan()  // Add this line to capture user input
        answer := scanner.Text() // Capture the user's answer
        if answer == term.Definition {
            fmt.Println("Correct!\n")
            incorrectTerms = incorrectTerms[1:]
        } else {
            fmt.Printf("Wrong! The correct answer is: %s\n\n", term.Definition)
        }
    }
}

			fmt.Println("Congratulations! You've completed all the terms.")

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// Load terms and definitions from a file.
func loadTermsFromFile(fileName string) ([]Term, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var terms []Term
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			term := Term{
				Term:       strings.TrimSpace(parts[0]),
				Definition: strings.TrimSpace(parts[1]),
			}
			terms = append(terms, term)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return terms, nil
}

// Shuffle the order of terms randomly.
func shuffle(terms []Term) {
	for i := range terms {
		j := rand.Intn(i + 1)
		terms[i], terms[j] = terms[j], terms[i]
	}
}

// Create chunks of terms with a specified size.
func makeChunks(terms []Term, chunkSize int) [][]Term {
	numChunks := (len(terms) + chunkSize - 1) / chunkSize
	chunks := make([][]Term, numChunks)

	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > len(terms) {
			end = len(terms)
		}
		chunks[i] = terms[start:end]
	}

	return chunks
}

func loadConfig() AppConfig {
	// Load the configuration from a file (e.g., config.json).
	configFileName := "learnconfig.json"
	configPath := filepath.Join(getUserHomeDir(), configFileName)
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return AppConfig{}
	}

	var config AppConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return AppConfig{}
	}

	return config
}

func saveConfig(config AppConfig) {
	// Save the configuration to a file (e.g., config.json).
	configFileName := "learnconfig.json"
	configPath := filepath.Join(getUserHomeDir(), configFileName)
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Printf("Error saving configuration: %v", err)
		return
	}

	if err := ioutil.WriteFile(configPath, data, 0644); err != nil {
		log.Printf("Error saving configuration: %v", err)
	}
}

func getUserHomeDir() string {
	// Get the user's home directory.
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting user home directory: %v", err)
	}
	return userHome
}
