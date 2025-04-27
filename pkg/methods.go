package bloom

import (
    "github.com/spaolacci/murmur3"
	"fmt"
	"time"
)
func (bf *BloomFilter) Add(item interface{}) {
	var digests []uint32
	var i uint
	for i = 0; i < bf.hashCount; i++ {
		index := createHashValue(i, item) % uint32(bf.size) 
		digests = append(digests, index)
		bf.bitArray[index] = true
	}
}

func (bf *BloomFilter) Check(item interface{}) bool{
	var i uint
	for i = 0; i < bf.hashCount; i++ {
		index := createHashValue(i, item) % uint32(bf.size) 
		if (bf.bitArray[index]){
		return true
		}
	}
return false
}

func createHashValue (index uint,item interface{}) uint32{
	digest := murmur3.New32WithSeed(uint32 (index)) // returns a hash
	digest.Write([]byte(fmt.Sprintf("%v", item)))
	hashValue := digest.Sum32()  // get a hash value -- what does sum do??
	return hashValue
}

func (bf *BloomFilter) CheckRetrievalTime(item interface{}) string {
    start := time.Now()

    bf.Check(item)

    duration := time.Since(start)

	if duration >= time.Second {
        return fmt.Sprintf("Operation took %.2f seconds", duration.Seconds())
    } else if duration >= time.Millisecond {
        return fmt.Sprintf("Operation took %.2f milliseconds", float64(duration.Milliseconds())) 
    } else if duration >= time.Microsecond {
        return fmt.Sprintf("Operation took %.2f microseconds", float64(duration.Microseconds())) 
    } else {
        return fmt.Sprintf("Operation took %.2f nanoseconds", float64(duration.Nanoseconds()))
    }
}
