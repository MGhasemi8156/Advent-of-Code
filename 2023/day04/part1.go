package day04

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
) 

func SolvePart1(filesrc string) {
    file, err := os.Open(filesrc)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


    re := regexp.MustCompile("^Card\\s+\\d+: (.*) \\| (.*)$")
    scanner := bufio.NewScanner(file)
    winning := make(map[int]bool)
    numbers := make([]int, 0)

    sum, score := 0, 0

    for scanner.Scan() {
        clear(winning)
        numbers = numbers[:0]
        score = 0
        
        parts := re.FindStringSubmatch(scanner.Text())
       
        for _, v := range strings.Split(parts[1], " ") {
            num, err := strconv.Atoi(v)
            if err == nil {
                winning[num] = true
            }
        }

        for _, v := range strings.Split(parts[2], " ") {
            num, err := strconv.Atoi(v)

            if winning[num] && err == nil {
                if score == 0 {
                    score = 1
                } else {
                    score *= 2
                }
            }
        }
        sum += score
    }

    fmt.Println(sum)
}
