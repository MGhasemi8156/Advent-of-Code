package day03

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func isNearGear(i int, lineNumber int, linebuf [3]string) (bool, int, int) {
    switch {
    case linebuf[0] != "" && rune(linebuf[0][i]) == '*': return true, lineNumber-1, i
    case linebuf[0] != "" && i > 0 && rune(linebuf[0][i-1]) == '*': return true, lineNumber-1, i-1
    case linebuf[0] != "" && i < len(linebuf[0])-1 && rune(linebuf[0][i+1]) == '*': return true, lineNumber-1, i+1
    case linebuf[2] != "" && rune(linebuf[2][i]) == '*': return true, lineNumber+1, i
    case linebuf[2] != "" && i > 0 && rune(linebuf[2][i-1]) == '*': return true, lineNumber+1, i-1
    case linebuf[2] != "" && i < len(linebuf[2])-1 && rune(linebuf[2][i+1]) == '*': return true, lineNumber+1, i+1
    case i > 0 && rune(linebuf[1][i-1]) == '*': return true, lineNumber, i-1
    case i < len(linebuf[1])-1 && rune(linebuf[1][i+1]) == '*': return true, lineNumber, i+1
    }

    return false, 0, 0
}

func SolvePart2(filesrc string) {
    file, err := os.Open(filesrc)

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var linebuf [3]string
    sum, num := 0, 0
    lineNumber := 0
    
    allGears := make(map[string][]int)

    for scanner.Scan() || linebuf[1] != ""{
        linebuf[0] = linebuf[1]
        linebuf[1] = linebuf[2]
        linebuf[2] = scanner.Text()

        if linebuf[1] != "" {
            lineNumber++
            num = 0
            nearGears := make(map[string]bool, 0)
            for i, v := range linebuf[1] {
                if unicode.IsDigit(v) {
                    digit, _ := strconv.Atoi(string(v))
                    num = num * 10 + digit
                    if near, y, x := isNearGear(i, lineNumber, linebuf); near {
                        gear := strconv.Itoa(y) + "," + strconv.Itoa(x)
                        nearGears[gear] = true
                    }
                } else {
                    for k := range nearGears {
                        allGears[k] = append(allGears[k], num)
                    }
                    nearGears = make(map[string]bool, 0)
                    num = 0
                }
                if unicode.IsDigit(v) && i == len(linebuf[1]) - 1 { 
                    for k := range nearGears {
                        allGears[k] = append(allGears[k], num)
                    }
                }
            }
        }
    }

    for _, v := range allGears {
        if len(v) == 2 {
            sum += v[0] * v[1]
        }
    }

    fmt.Println(sum)
}
