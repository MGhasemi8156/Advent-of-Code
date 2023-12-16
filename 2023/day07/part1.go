package day07

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
    deck string
    bid int
}

func getTypeValue(h hand) int {
    tmp := make(map[rune]int)
    for _, v := range h.deck {
        tmp[v]++
    }
    
    decks := make([]int, 0)
    for _, v := range tmp {
        decks = append(decks, v)  
    }

    sort.Ints(decks)
    switch len(decks) {
    case 1: return 7
    case 2: {
        if decks[1] == 4 {
            return 6
        }
        return 5
    }
    case 3: {
        if decks[2] == 3 {
            return 4
        }
        return 3
    }
    case 4: return 2
    case 5: return 1
    }
    
    return 0
}

func SolvePart1(filesrc string) {
    file, err := os.Open(filesrc)
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    hands := make([]hand, 0)
    for scanner.Scan() {
        parsedLine := strings.Fields(scanner.Text())
        deck := parsedLine[0]
        bid, _ := strconv.Atoi(parsedLine[1])
        hands = append(hands, hand{deck, bid})
    }

    scores := map[string]int {"A": 13, "K":12, "Q":11, "J":10, "T":9, "9":8, "8":7, "7":6, "6":5, "5":4, "4":3, "3":2, "2":1}

    sort.Slice(hands, func(i, j int) bool {
        if getTypeValue(hands[i]) > getTypeValue(hands[j]) {
            return false
        } else if getTypeValue(hands[i]) < getTypeValue(hands[j]) {
            return true
        } else {
            a := hands[i]
            b := hands[j]
            for k := 0; k < 5; k++ {
                if scores[string(a.deck[k])] > scores[string(b.deck[k])] {
                    return false
                } else if scores[string(a.deck[k])] < scores[string(b.deck[k])] {
                    return true
                }
            }
        }
        return false
    })
  
    fmt.Println(hands)

    sum := 0
    for i, v := range hands {
        sum += (i + 1) * v.bid
    }

    fmt.Println(sum)

}
