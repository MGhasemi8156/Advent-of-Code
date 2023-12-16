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

// Closed interval
type interval struct {
    start int
    end int
}

func intersection(a interval, b interval) (interval, bool) {
    if a.start > b.end || a.end < b.start {
        return interval{}, false
    } 

    res := interval{}
    res.start = max(a.start, b.start)
    res.end = min(a.end, b.end)

    return res, true
}

func difference(first interval, second interval) []interval {
    res := make([]interval, 0)
    if first.start < second.start {
        res = append(res, interval{first.start, second.start - 1}) 
    }
    if second.end < first.end && first.start != second.end {
        res = append(res, interval{second.end + 1, first.end})
    }

    return res
}

func SolvePart2(filesrc string) {
    file, err := os.Open(filesrc)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    
    seedsRe := regexp.MustCompile("^seeds: (.*)$")
    seeds := make([]interval, 0)
    
    scanner := bufio.NewScanner(file)
    count := 0 
    var tmpInterval interval
    if scanner.Scan() {
        parts := seedsRe.FindStringSubmatch(scanner.Text())

        for _, v := range strings.Split(parts[1], " ") {
            num, err := strconv.Atoi(v) 
            if err != nil {
                continue
            }
            if count%2 == 0 {
                tmpInterval = interval{start: num} 
            } else {
                tmpInterval.end = tmpInterval.start + num - 1
                seeds = append(seeds, tmpInterval)
            }
            count++
        }
    }
   
    var mappings []interval
    copy(mappings, seeds)

    mapRe := regexp.MustCompile("^.* map.*$")
    rangeRe := regexp.MustCompile("^(\\d+) (\\d+) (\\d+)$")
    for scanner.Scan() {

        line := scanner.Text()
        switch {
        case mapRe.FindString(line) != "": {
            for _, v := range seeds {
                mappings = append(mappings, v)
            }

            seeds = make([]interval, len(mappings))
            copy(seeds, mappings)
            mappings = mappings[:0]
        }
        case rangeRe.FindString(line) != "": {
            parts := rangeRe.FindStringSubmatch(line)
            dst, _ := strconv.Atoi(parts[1])
            src, _ := strconv.Atoi(parts[2])
            rng, _ := strconv.Atoi(parts[3])
            cnt := dst - src

            for i := 0; i < len(seeds); {
                matchedInterval, isIntersected := intersection(seeds[i], interval{src, src + rng - 1})
                if isIntersected {
                    mappings = append(mappings, interval{matchedInterval.start + cnt , matchedInterval.end + cnt})
                    
                    seeds = append(seeds, difference(seeds[i], interval{src, src + rng - 1})...) 
                    seeds = append(seeds[:i], seeds[i+1:]...)

                } else {
                    i++
                }
            }
        }
        }
    }

    for _, v := range seeds {
        mappings = append(mappings, v)
    }


    minVal := 2147483647
    for _, v := range mappings {
        minVal = min(minVal, v.start)
    }
    fmt.Println(minVal)
}
