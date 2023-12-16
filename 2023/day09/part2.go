package day09

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SolvePart2(filesrc string) {
    file, err := os.Open(filesrc)
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    history := make([]int, 0)
    sum := 0
    for scanner.Scan() {
        for _, v := range strings.Fields(scanner.Text())  {
            tmpNum, err := strconv.Atoi(v)
            if err == nil {
                history = append(history, tmpNum) 
            }
        }
        sum += findNext(history)
        history = history[:0]
    }

    fmt.Println(sum)
}
