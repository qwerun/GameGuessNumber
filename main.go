package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	fmt.Println("Я загадал число от одного до 100, угадай его")

	targer := rand.Intn(100) + 1
	totalAttempts := 8 // КОЛИЧЕСТВО ПОПЫТОК МОЖНО РЕГУЛИРОВАТЬ
	myCounter := totalAttempts - 1
	numGuesses := totalAttempts - 1
	for guesses := 0; guesses < numGuesses+1; guesses++ {

		fmt.Println("Выберите число: ")
		reader := bufio.NewReader(os.Stdin)
		num, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		num = strings.Trim(num, "\n\r")
		old, errorr := strconv.Atoi(num)
		if errorr != nil {
			log.Fatal(errorr)
		}

		if old > targer {
			fmt.Println("Осталось попыток: " + strconv.Itoa(myCounter))
			myCounter--
			fmt.Println("Твоё число больше, чем нужно")
		} else if old < targer {
			fmt.Println("Осталось попыток: " + strconv.Itoa(myCounter))
			myCounter--
			fmt.Println("Твоё число меньше, чем нужно")
		} else if old == targer {
			fmt.Println("Ты победил")
			break
		}

		if guesses == numGuesses {
			fmt.Println("Ты проиграл, правильным числом было: " + strconv.Itoa(targer))
			break
		}
	}
}
