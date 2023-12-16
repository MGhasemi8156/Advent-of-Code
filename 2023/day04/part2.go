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

func SolvePart2(filesrc string) {
    file, err := os.Open(filesrc)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    
    re := regexp.MustCompile("^Card\\s+\\d+: (.*) \\| (.*)$")
    scanner := bufio.NewScanner(file)
    winning := make(map[int]bool)
    numbers := make([]int, 0)

    match := 0
    matches := make([]int, 0)
    for scanner.Scan() {
        clear(winning)
        numbers = numbers[:0]
        match = 0
        
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
                match += 1
            }
        }
        matches = append(matches, match)
    }

    cards := make([]int, len(matches))
    for i, v := range matches {
        cards[i]++
        for j := 1; j <= v; j++ {
            cards[i+j] += cards[i]
        }
    }

    sum := 0
    for _, v := range cards {
        sum += v
    }


    fmt.Println(sum)
}
