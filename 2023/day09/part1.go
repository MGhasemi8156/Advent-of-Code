package day09

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findNext(history []int) int {
    tmp := make([]int, 0)
    l := len(history)
    
    zero := true
    for _, v := range history {
        if v != 0 {
            zero = false
            break
        }
    }
    if zero {
        return 0
    }
    
    for i := 1; i < l; i++ {
        tmp = append(tmp, history[i] - history[i-1])
    }

    return history[l-1] + findNext(tmp)
}


func SolvePart1(filesrc string) {
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
