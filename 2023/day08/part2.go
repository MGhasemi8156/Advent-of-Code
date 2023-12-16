package day08

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func gcd(a int, b int) int {
    if a < b {
        return gcd(b, a)
    }
    if a%b == 0 {
        return b
    }
    return gcd(b, a%b)
}

func SolvePart2(filesrc string) {
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
    starts := make([]string, 0)
    scanner.Scan() // empty line
    for scanner.Scan() {
        parts := LRRE.FindStringSubmatch(scanner.Text())
        node = parts[1]
        l = parts[2]
        r = parts[3]
        dirs[node] = lr{l, r}
        if rune(node[2]) == 'A' {
            starts = append(starts, node)
        }
    }
    
    var tmp string 
    var steps int
    allSteps := make([]int, 0)
    for _, start := range starts {
        tmp = start 
        steps = 0
        outer:
        for rune(tmp[2]) != 'Z' {
            for _, v := range instructions {
                switch v {
                case 'L': tmp = dirs[tmp].left
                case 'R': tmp = dirs[tmp].right
                }
                steps++
                if rune(tmp[2]) == 'Z' {
                    break outer
                }
            }
        }
        allSteps = append(allSteps, steps)
    }
    
    for len(allSteps) > 1 {
        allSteps[1] = allSteps[0] * allSteps[1] / gcd(allSteps[0], allSteps[1])
        allSteps = allSteps[1:]
    }

    fmt.Println(allSteps[0])

}
