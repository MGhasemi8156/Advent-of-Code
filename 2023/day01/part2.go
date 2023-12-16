package day01

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var digits [9]string;

func checkForNum(line string, i int) (int, error) {
    outer:    
    for num, v := range digits {
        if len(v) > len(line) - i {
            continue
        }
        for j := 0; j < len(v); j++ {
            if v[j] != line[i + j] {
                continue outer
            }
        }
        return num + 1, nil
    } 
    return 0, errors.New("Ivalid")
}

func SolvePart2() {
    digits = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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
        
        for i, v := range line {
            num, numerror := strconv.Atoi(string(v))
            if numerror == nil {
                numbers = append(numbers, num)
                continue
            }
            num, numerror = checkForNum(line, i)
            if numerror == nil {
                numbers = append(numbers, num)
            }
        }

        fmt.Println(numbers)

        sum += numbers[0] * 10 + numbers[len(numbers) - 1]
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(sum)
}
