package internal

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"time"
)

const (
	SECRET_LENGTH   = 256
	ID_LENGTH       = 16
	TIME_BLOCK_SIZE = 10
)

// func main() {

// 	secret, err := generateSecret()
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	// log.Println(secret)

// 	// for i := 0; i < 30; i++ {
// 	// 	log.Println(getTimeBlock())
// 	// 	time.Sleep(1 * time.Second)
// 	// }

// 	for i := 0; i < 30; i++ {
// 		log.Println(hash(secret, getTimeBlock()))
// 		time.Sleep(1 * time.Second)
// 	}

// }

func GenerateSecret() (b []byte, err error) {
	randomBytes := make([]byte, SECRET_LENGTH)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return []byte(randomBytes), err
	}
	return randomBytes, nil
}

func GenerateID() (b []byte, err error) {
	randomBytes := make([]byte, ID_LENGTH)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return []byte(randomBytes), err
	}
	return randomBytes, nil
}

func GetTimeBlock() (t int64) {
	return time.Now().Unix() / TIME_BLOCK_SIZE
}

func Hash(secret []byte, t int64) (s []byte) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(t))
	timedSecret := append(secret, b...)

	sum := sha256.Sum256(timedSecret)
	return sum[:]
}

func ByteArrayToInt(b []byte) int {
	return int(binary.LittleEndian.Uint64(b))
}
