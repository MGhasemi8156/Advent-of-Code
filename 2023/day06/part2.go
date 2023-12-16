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
func SolvePart2(filesrc string) {
    file, err := os.Open(filesrc)
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    scanner.Scan()
    timeRe := regexp.MustCompile("^Time: (.*)$")
    tparts := timeRe.FindStringSubmatch(scanner.Text())
    var tstring string
    for _, v := range strings.Fields(tparts[1]) {
        _, err := strconv.Atoi(v)
        if err == nil {
            tstring = tstring + v
        }
    }

    scanner.Scan()
    distRe := regexp.MustCompile("^Distance: (.*)$")
    dparts := distRe.FindStringSubmatch(scanner.Text())
    var dstring string
    for _, v := range strings.Fields(dparts[1]) {
        _, err := strconv.Atoi(v)
        if err == nil {
            dstring = dstring + v
        }
    }

    time, _ := strconv.Atoi(tstring)
    distance, _ := strconv.Atoi(dstring)
    
    res := 0
    for i := 0; i < time; i++ {
        if (time - i) * i > distance {
            res++
        }
    }

    fmt.Println(res)
}
