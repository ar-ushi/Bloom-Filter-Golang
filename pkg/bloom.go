package bloom

import (
	"math"
)

type BloomFilter struct {
	fpProb     float64
	size       uint
	hashCount  uint
	bitArray   []bool
}


func New(itemsCount uint, fpProb float64) *BloomFilter {
	size := getArrSize(itemsCount, fpProb)
	hashCount := getHashCount(size, itemsCount)

	return &BloomFilter{
		fpProb:    fpProb,
		size:      size,
		hashCount: hashCount,
		bitArray:  make([]bool, size), 
	}
}

func getArrSize(itemsCount uint, fpProb float64) uint {
	arrSize := math.Ceil(-float64(itemsCount) * math.Log(fpProb) / (math.Ln2 * math.Ln2))
	return uint(arrSize);
}

func getHashCount(size, itemsCount uint) uint{
	count :=  math.Ceil((float64(size) / float64(itemsCount)) * math.Ln2)
	return uint(count)
}

