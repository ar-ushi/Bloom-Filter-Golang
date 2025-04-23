package bloom

import (
    "github.com/spaolacci/murmur3"
)
func (bf *BloomFilter) add(item interface{}) {
	var digests []uint32
	for i := 0; i < bf.hashCount; i++ {
		index := createHashValue(i, item) % uint32(bf.size) 
		digests = digests.append(digests, index)
		bf.bitArray[index] = true
	}
}

func (bf *BloomFilter) check(item interface{}) bool{
	for i := 0; i < bf.hashCount; i++ {
		index := createHashValue(i, item) % uint32(bf.size) 
		if (bf.bitArray[index]){
		return true
		}
	}
return false
}

func createHashValue (index int,item interface{}) uint32{
	digest := murmur3.new32WithSeed(uint32 (index)) // returns a hash
	digest.Write([]byte(fmt.Sprintf("%v", item)))
	hashValue := digest.Sum32()  // get a hash value -- what does sum do??
	return hashValue
}