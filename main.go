package main

import (
    "encoding/csv"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
)

// ANSI color codes for console text colors
const (
    ColorRed    = "\033[31m"
    ColorGreen  = "\033[32m"
    ColorYellow = "\033[33m"
    ColorReset  = "\033[0m"
)

// readTableData reads data from the CSV file and returns a map of countries and their capitals
func readTableData() (map[string]string, error) {
    file, err := os.Open("CountryCapital.csv")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    countries := make(map[string]string)

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    // Parsing CSV records to create a map of countries and their capitals
    for _, record := range records {
        countries[record[1]] = record[0]
    }

    return countries, nil
}

// generateOptions generates multiple choice options for the quiz
func generateOptions(correctAnswer string, capitals []string) []string {
    options := make([]string, 4)
    optionsIndex := make([]int, 4)
    correctIndex := rand.Intn(4)
    optionsIndex[0] = correctIndex
    options[correctIndex] = correctAnswer

    // Filling the options with random capitals
    for i := 0; i < 4; i++ {
        if i != correctIndex {
            randomIdx := rand.Intn(len(capitals))
            options[i] = capitals[randomIdx]
        }
    }

    return options
}

// getCapitals extracts all the capitals from the countries map
func getCapitals(countries map[string]string) []string {
    capitals := []string{}
    for capital := range countries {
        capitals = append(capitals, capital)
    }
    return capitals
}

// printUsage prints the usage instructions for the quiz game
func printUsage() {
    fmt.Println(`Welcome to the Capital Quiz Game!
Usage:
  quizgame [options]
Options:
  -h\t\tDisplay this help message
  -q <length>\tSet the number of questions to <length>
Examples:
  To start the game with default settings:
    quizgame
  To set the number of questions to 10:
    quizgame -q 10`)
}

// RatePlayer rates the player's performance based on correct answers
func RatePlayer(questionLength, correctAnswers int) {
    switch {
    case questionLength == correctAnswers:
        fmt.Printf("\nThanks for playing!\n"+ ColorGreen +":) You answered %d/%d questions correctly\n"+ColorReset, correctAnswers, questionLength)
    case correctAnswers >= questionLength / 2:
        fmt.Printf("\nThanks for playing!\n"+ ColorYellow +":) You answered %d/%d questions correctly\n"+ColorReset, correctAnswers, questionLength)
    default:
        fmt.Printf("\nThanks for playing!\n"+ ColorRed +":( You answered %d/%d questions correctly\n"+ColorReset, correctAnswers, questionLength)
    }
}

func main() {
    args := os.Args[1:]
    questionLength := -1
    lengthFlag := false

    // Parsing command line arguments
    for i := 0; i < len(args); i++ {
        if args[i][0] == '-' {
            switch args[i][1] {
            case 'h':
                printUsage()
                return
            case 'q':
                if len(args) > i+1 {
                    questionLength, _ = strconv.Atoi(args[i+1])
                    lengthFlag = true
                    i++
                } else {
                    fmt.Println("Question length is missing")
                    return
                }
            }
        } else {
            fmt.Println("Invalid option")
            return
        }
    }

    if questionLength < 1 && lengthFlag {
        fmt.Println("Invalid option")
        return
    }

    countries, err := readTableData()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    capitals := getCapitals(countries)
    correctAnswers := 0
    questions := 0

    // Main game loop
    for {
        capital := capitals[rand.Intn(len(capitals))]
        country := countries[capital]
        options := generateOptions(capital, capitals)

        // Check if question limit is reached
        if questionLength == questions {
            RatePlayer(questionLength, correctAnswers)
            return
        }

        fmt.Printf("\nCountry: %s\nChoose the capital:\n", country)
        for i, option := range options {
            fmt.Printf("%d. %s\n", i+1, option)
        }

    start:
        fmt.Print("Your choice: ")

        var input string
        fmt.Scanf("%s", &input)
        fmt.Scanln()

        input = strings.TrimSpace(strings.ToLower(input))

        // Exiting the game if requested by the user
        if input == "exit" || input == "q" || input == "quit" {
            RatePlayer(questionLength, correctAnswers)
            return
        }

        choice, err := strconv.Atoi(input)
        if err != nil || choice < 1 || choice > len(options) {
            fmt.Println("Invalid choice. Please select a valid option.")
            goto start
        }

        questions++
        if options[choice-1] == capital {
            fmt.Println(ColorGreen + "Correct!" + ColorReset)
            correctAnswers++
        } else {
            fmt.Printf(ColorRed + "Incorrect. The capital of %s is %s.\n" + ColorReset, country, capital)
        }
    }
}
