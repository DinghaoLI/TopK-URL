package utils

// BKDR Hash Function
func BKDRHash(str string) uint32 {
	list := []byte(str)
	var seed uint32 = 131 // 31 131 1313 13131 131313 etc..
	var hash uint32 = 0
	for i := 0; i < len(list); i++ {
		hash = hash*seed + uint32(list[i])
	}
	return (hash & 0x7FFFFFFF)
}

// BKDR Hash Function 64
func BKDRHash64(str string) uint64 {
	list := []byte(str)
	var seed uint64 = 131 // 31 131 1313 13131 131313 etc..
	var hash uint64 = 0
	for i := 0; i < len(list); i++ {
		hash = hash*seed + uint64(list[i])
	}
	return (hash & 0x7FFFFFFFFFFFFFFF)
}
