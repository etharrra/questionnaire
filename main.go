package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"io"
	"net/http"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

type Quiz struct {
	Results []struct {
		Question         string   `json:"question"`
		CorrectAnswer    string   `json:"correct_answer"`
		IncorrectAnswers []string `json:"incorrect_answers"`
	} `json:"results"`
}

func main() {
	limit := flag.Int("limit", 5, "the questions limit for the quiz")
	flag.Parse()

	url := fmt.Sprintf("https://opentdb.com/api.php?amount=%d&category=18&difficulty=easy&type=multiple", *limit)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("API is not avaliable")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// Convert the response body to a string
	bodyString := string(body)

	// Unescape HTML entities in the string
	unescapedBody := html.UnescapeString(bodyString)

	var quiz Quiz
	err = json.Unmarshal([]byte(unescapedBody), &quiz)
	if err != nil {
		panic(err)
	}

	correct := 0
	for _, q := range quiz.Results {
		choices := append(q.IncorrectAnswers, q.CorrectAnswer)
		answerCh := make(chan string)
		go func() {
			prompt := promptui.Select{
				Label: q.Question,
				Items: choices,
			}
			_, result, err := prompt.Run()
			if err != nil {
				panic(fmt.Sprintf("Prompt failed %v\n", err))
			}
			answerCh <- result
		}()
		
		answer := <-answerCh
		if answer == q.CorrectAnswer {
			correct++
			color.Green("Correct Answer")
		} else {
			color.Red("Incorrect Answer")
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(quiz.Results))
}
