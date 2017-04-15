package main

import (
  "strings"
  "fmt"
  "sort"
  "strconv"
  "flag"
  "time"
  "log"
  "io/ioutil"
)

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func readLines2(filename string) ([]int, error) {
    defer timeTrack(time.Now(), "readLines2")
    var lines []int

    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    for _, line := range strings.Split(string(data), "\n") {
        if i, err := strconv.Atoi(line); err == nil {
          lines = append(lines, i)
        }
    }

    return lines, nil
}

func rank(whiteList []int, key int) int {
  lo := 0
  hi := len(whiteList) - 1

  for lo <= hi {
    mid := lo + (hi - lo) / 2
    if key < whiteList[mid] {
      hi = mid - 1
    } else if key > whiteList[mid] {
      lo = mid + 1
    } else {
      return mid
    }
  }

  return -1;
}

func main() {
  whiteListFileNamePtr := flag.String("wFile", "white-list.txt", "White list file name")
  transactioFileNamePtr := flag.String("tFile", "transactions.txt", "Transaction list file name")
  flag.Parse()
  whiteListFileName := *whiteListFileNamePtr
  transactionFileName := *transactioFileNamePtr
  whiteList, err := readLines2(whiteListFileName)
  if err != nil {
    fmt.Println(err)
    return
  }
  transactions, err := readLines2(transactionFileName)
  if err != nil {
    fmt.Println(err)
    return
  }

  defer timeTrack(time.Now(), "sorting and searching")
  var notPermitted []int
  sort.Ints(whiteList)
  for _,value := range transactions {
    if rank(whiteList, value) == -1 {
      notPermitted = append(notPermitted, value)
    }
  }
  fmt.Println(len(notPermitted))
  fmt.Println(notPermitted[0:2])
}
