# Questionnaire

This is a Go-based command-line application that fetches multiple-choice quiz questions from the Open Trivia Database API and presents them to the user. The application evaluates the user's answers and provides feedback on their performance.

## Features

-   Fetches quiz questions from the Open Trivia Database API.
-   Presents multiple-choice questions to the user.
-   Evaluates user answers and provides immediate feedback.
-   Displays the final score at the end of the quiz.
-   Allows the user to specify the number of questions via a command-line flag.

## Prerequisites

-   Go 1.16 or later

## Installation

1. Clone the repository:

```sh
git clone https://github.com/etharrra/questionnaire.git
cd questionnaire
```

2. Install dependencies:

```sh
go mod tidy
```

## Usage

To build the application, use the following command:

```sh
go build .
```

To install the application, use the following command:

```sh
go install .
```

To run the application

```sh
questionnaire --limit <number_of_questions>
```

Replace `<number_of_questions>` with the desired number of quiz questions (default is 5).

Example:

```sh
questionnaire --limit 10
```

## Explanation

This application retrieves quiz questions from the Open Trivia Database API, presents them as multiple-choice questions, and evaluates the user's answers. The user can specify the number of questions using the `--limit` flag. The application provides immediate feedback for each answer and displays the final score at the end.
