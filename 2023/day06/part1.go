package day06

import (
    "fmt"
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// inequlity has better perfomance but is not necessary considering the input size
func SolvePart1(filesrc string) {
    file, err := os.Open(filesrc)
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    scanner.Scan()
    timeRe := regexp.MustCompile("^Time: (.*)$")
    tparts := timeRe.FindStringSubmatch(scanner.Text())
    times := make([]int, 0)
    for _, v := range strings.Fields(tparts[1]) {
        num, err := strconv.Atoi(v)
        if err == nil {
            times = append(times, num)
        }
    }

    scanner.Scan()
    distRe := regexp.MustCompile("^Distance: (.*)$")
    dparts := distRe.FindStringSubmatch(scanner.Text())
    dists := make([]int, 0)
    for _, v := range strings.Fields(dparts[1]) {
        num, err := strconv.Atoi(v)
        if err == nil {
            dists = append(dists, num)
        }
    }

    fmt.Println(times)
    fmt.Println(dists)

    res := 1
    for i := 0; i < len(times); i++ {
        t, d := times[i], dists[i]
        c := 0
        for j := 1; j < t; j++ {
            if (t - j) * j > d {
                c++
            }
        }
        res *= c
    }

    fmt.Println(res)

}
