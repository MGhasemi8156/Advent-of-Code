package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
) 

func SolvePart2() {
    
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    lineRe := regexp.MustCompile("^Game (\\d+): (.*)$")
    redRe := regexp.MustCompile("(\\d+) red")
    greenRe := regexp.MustCompile("(\\d+) green")
    blueRe := regexp.MustCompile("(\\d+) blue")

    sum := 0
    outer:
    for scanner.Scan() {
        line := scanner.Text()
    
        matches := lineRe.FindStringSubmatch(line)
        game, _ := strconv.Atoi(matches[1])
        data := matches[2]

        redc := redRe.FindAllStringSubmatch(data, -1)
        for _, v := range redc {
            red, _ := strconv.Atoi(v[1])
            if red > 12 {
                continue outer
            }
        }

        greenc := greenRe.FindAllStringSubmatch(data, -1)
        for _, v := range greenc {
            green, _ := strconv.Atoi(v[1])
            if green> 13 {
                continue outer
            }    
        }

        bluec := blueRe.FindAllStringSubmatch(data, -1)
        for _, v := range bluec {
            blue, _ := strconv.Atoi(v[1])
            if blue> 14 {
                continue outer
            }
        }
        sum += game
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(sum)
}
