package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Quiz struct {
	question string
	answer   string
}

func OpenFile(path string) (*os.File, error) {

	f, err := os.Open(path)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo.")
	}

	return f, err

}

func ReaderData(file *os.File) ([][]string, error) {
	defer file.Close()

	csvReader := csv.NewReader(file)

	data, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Erro ao ler arquivo.")
	}

	return data, err

}

func main() {
	f, _ := OpenFile("./problems.csv")
	defer f.Close()

	data, _ := ReaderData(f)

	var score int = 0

	for _, item := range data {
		var quizz Quiz
		quizz = Quiz{
			question: item[0],
			answer:   item[1],
		}

		fmt.Printf("Qual é o resultado de: %s: \n", quizz.question)
		var input string
		fmt.Scan(&input)
		if input == quizz.answer {
			score += 1
		}
	}

	fmt.Printf("\n\nSeu score é: %d", score)
}
