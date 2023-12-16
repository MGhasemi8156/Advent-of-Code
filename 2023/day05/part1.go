package day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func createMapping(n int) []int {
    mappings := make([]int, n)
    for i, _ := range mappings {
        mappings[i] = -1
    }
    return mappings
}

func SolvePart1(filesrc string) {
    file, err := os.Open(filesrc)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    
    seedsRe := regexp.MustCompile("^seeds: (.*)$")
    seeds := make([]int, 0)
    
    scanner := bufio.NewScanner(file)
    if scanner.Scan() {
        parts := seedsRe.FindStringSubmatch(scanner.Text())
        fmt.Println(parts)

        for _, v := range strings.Split(parts[1], " ") {
            num, err := strconv.Atoi(v) 
            if err == nil {
                seeds = append(seeds, num)
            }
        }
    }
   
    var mappings []int
    copy(mappings, seeds)

    mapRe := regexp.MustCompile("^.* map.*$")
    rangeRe := regexp.MustCompile("^(\\d+) (\\d+) (\\d+)$")
    for scanner.Scan() {
        line := scanner.Text()
        switch {
        case mapRe.FindString(line) != "": {
            for i, v := range mappings {
                if v == -1 {
                    mappings[i] = seeds[i]
                }
            }
            copy(seeds, mappings) 
            mappings = createMapping(len(seeds))

            fmt.Println(seeds)
        }
        case rangeRe.FindString(line) != "": {
            parts := rangeRe.FindStringSubmatch(line)
            dst, _ := strconv.Atoi(parts[1])
            src, _ := strconv.Atoi(parts[2])
            rng, _ := strconv.Atoi(parts[3])
            
            for i, v := range seeds {
                if v >= src && v < src + rng {
                    mappings[i] = v - src + dst
                }
            }
        }
        }
    }

    for i, v := range mappings {
        if v == -1 {
            mappings[i] = seeds[i]
        }
    }

    fmt.Println(mappings)
    minVal := 2147483647
    for _, v := range mappings {
        if v < minVal {
            minVal = v
        }
    }
    fmt.Println(minVal)
}
