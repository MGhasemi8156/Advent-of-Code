package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strconv"
)

func SolvePart1() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    sum := 0

    for scanner.Scan() {
        line := scanner.Text()
        numbers := make([]int, 0)
        for _, v := range line {
            num, numerror := strconv.Atoi(string(v))
            if numerror == nil {
                numbers = append(numbers, num)
            }
        }

        sum += numbers[0] * 10 + numbers[len(numbers) - 1]
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(sum)
}
