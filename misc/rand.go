package misc

import "crypto/rand"

const RandStringChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandBytes(amount int) []byte {
	buf := make([]byte, amount)

	rand.Read(buf)

	return buf
}

func RandString(amount int) string {
	buf := RandBytes(amount)

	for k, v := range buf {
		buf[k] = RandStringChars[int(v)%len(RandStringChars)]
	}

	return string(buf)
}
