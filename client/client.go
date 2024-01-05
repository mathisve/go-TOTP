package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gototp/internal"
	p "gototp/totpProto"
)

func main() {
	conn, err := grpc.Dial("localhost:6000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	c := p.NewServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetSecret(ctx, &p.GetSecretRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("SecretID:", r.SecretId[:10])

	for i := 0; i < 100; i++ {
		totp := internal.Hash(r.Secret, internal.GetTimeBlock())

		// if i%2 == 0 {
		// 	for i2, j := 0, len(totp)-1; i2 < j; i2, j = i2+1, j-1 {
		// 		totp[i2], totp[j] = totp[j], totp[i2]
		// 	}

		// }

		a, err := c.Challenge(context.Background(), &p.ChallengeRequest{
			SecretId: r.SecretId,
			Topt:     totp,
		})

		if err != nil {
			log.Panic(err)
		}

		if a.Ok {
			log.Println("Challenge accepted", totp[:10])
		}

		time.Sleep(10 * time.Second)
	}
}
