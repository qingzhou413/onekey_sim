package byteutil

func PutByteToBuffer(bs []byte, b byte, index int) {
	bs[index] = b
}

func PutByteArrayToBuf(bs []byte, toAddArray []byte, index int) {
	for i, v := range toAddArray {
		bs[index+i] = v
	}
}

func PutShortToBuffer(bs []byte, s uint16, index int) {
	PutByteArrayToBuf(bs, ShortToByteArr(s), index);
}

func ShortToByteArr(u uint16) []byte {
	var bs = make([]byte, 2)
	bs[0] = byte(u >> 8)
	bs[1] = byte(u)
	return bs
}

func PutIntToBuffer(bs []byte, i int, index int) {
	PutByteArrayToBuf(bs, IntToByteArr(i), index);
}

func IntToByteArr(i int) []byte {
	var bs = make([]byte, 4)
	bs[0] = byte(i >> 24)
	bs[1] = byte(i >> 16)
	bs[2] = byte(i >> 8)
	bs[3] = byte(i)
	return bs
}

func Put8LongToBuffer(bs []byte, l uint64, index int) {
	PutByteArrayToBuf(bs, LongTo8LenByteArr(l), index);
}

func LongTo8LenByteArr(i uint64) []byte {
	var bs = make([]byte, 8)
	bs[0] = byte(i >> 56)
	bs[1] = byte(i >> 48)
	bs[2] = byte(i >> 40)
	bs[3] = byte(i >> 32)
	bs[4] = byte(i >> 24)
	bs[5] = byte(i >> 16)
	bs[6] = byte(i >> 8)
	bs[7] = byte(i)
	return bs
}

func Put6LongToBuffer(bs []byte, l uint64, index int) {
	PutByteArrayToBuf(bs, LongTo6LenByteArr(l), index);
}

func LongTo6LenByteArr(i uint64) []byte {
	var bs = make([]byte, 6)
	bs[0] = byte(i >> 40)
	bs[1] = byte(i >> 32)
	bs[2] = byte(i >> 24)
	bs[3] = byte(i >> 16)
	bs[4] = byte(i >> 8)
	bs[5] = byte(i)
	return bs
}
