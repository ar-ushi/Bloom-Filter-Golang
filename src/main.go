package main

import (
    "fmt"
    "log"
    "flag"
    "bloomfilter/bloom"
)

func main() {
    itemsCount := flag.Int("items", 1000, "Number of items to store in the Bloom filter")
    fpProb := flag.Float64("fp", 0.01, "False positive probability")
    flag.Parse() 

    if *itemsCount <= 0 || *fpProb <= 0 || *fpProb >= 1 {
        log.Fatalf("Invalid input: itemsCount and fpProb must be positive and fpProb < 1")
    }

    filter := bloom.New(uint(*itemsCount), *fpProb)

    filter.Add("item1")
    filter.Add("item2")

    fmt.Println(filter.Check("item1")) 
    fmt.Println(filter.Check("item3")) 
    fmt.Println(filter.CheckRetrievalTime("item1")) 
}
