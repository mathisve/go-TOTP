package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

const (
	SECRET_LENGTH   = 256
	TIME_BLOCK_SIZE = 10
)

func main() {
	secret, err := generateSecret()
	if err != nil {
		log.Panic(err)
	}

	// log.Println(secret)

	// for i := 0; i < 30; i++ {
	// 	log.Println(getTimeBlock())
	// 	time.Sleep(1 * time.Second)
	// }

	for i := 0; i < 30; i++ {
		log.Println(hash(secret, getTimeBlock()))
		time.Sleep(1 * time.Second)
	}

}

func generateSecret() (b []byte, err error) {
	randomBytes := make([]byte, SECRET_LENGTH)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return []byte(randomBytes), err
	}
	return randomBytes, nil
}

func getTimeBlock() (t int64) {
	return time.Now().Unix() / TIME_BLOCK_SIZE
}

func hash(secret []byte, t int64) (s [32]byte) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(t))
	timedSecret := append(secret, b...)

	return sha256.Sum256(timedSecret)
}
