package day08

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)


type lr struct {
    left string
    right string
}

func SolvePart1(filesrc string) {
    file, err := os.Open(filesrc) 
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    scanner.Scan()
    instructions := scanner.Text()

    LRRE := regexp.MustCompile("^(\\w+) = \\((\\w+), (\\w+)\\)$")
   
    dirs := make(map[string]lr)
    var node, l, r string
    scanner.Scan() // empty line
    for scanner.Scan() {
        parts := LRRE.FindStringSubmatch(scanner.Text())
        node = parts[1]
        l = parts[2]
        r = parts[3]
        dirs[node] = lr{l, r}
    }
    
    tmp := "AAA"
    steps := 0

    outer:
    for tmp != "ZZZ" {
        for _, v := range instructions {
            switch v {
            case 'L': tmp = dirs[tmp].left
            case 'R': tmp = dirs[tmp].right
            }
            steps++
            if tmp == "ZZZ" {
                break outer
            }
        }
    }

    fmt.Println(steps)

}
