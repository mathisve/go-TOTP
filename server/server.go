package main

import (
	"context"
	"errors"
	"log"
	"net"

	"gototp/internal"
	p "gototp/totpProto"

	"google.golang.org/grpc"
)

var (
	m = map[int][]byte{}
)

type server struct {
	p.UnimplementedServerServer
}

func (s *server) GetSecret(ctx context.Context, in *p.GetSecretRequest) (*p.GetSecretResponse, error) {
	log.Println("GetSecret request received")

	id, err := internal.GenerateID()
	if err != nil {
		return nil, err
	}

	secret, err := internal.GenerateSecret()
	if err != nil {
		return nil, err
	}

	m[internal.ByteArrayToInt(id)] = secret
	defer log.Println("Issued new secret ", id[:10])

	return &p.GetSecretResponse{
		SecretId: id,
		Secret:   secret,
	}, nil
}

func (s *server) Challenge(ctx context.Context, in *p.ChallengeRequest) (*p.ChallengeResponse, error) {
	log.Println("Challenge request received ", in.SecretId[:10], in.Topt[:10])

	secret := m[internal.ByteArrayToInt(in.SecretId)]
	if secret == nil {
		return &p.ChallengeResponse{Ok: false}, errors.New("unknown secretId")
	}

	hashed := internal.Hash(secret, internal.GetTimeBlock())

	if len(in.Topt) != len(hashed) {
		return &p.ChallengeResponse{Ok: false}, errors.New("totp length mismatch")
	}

	for i := range in.Topt {
		if hashed[i] != in.Topt[i] {
			return &p.ChallengeResponse{Ok: false}, errors.New("totp content mismatch")
		}
	}

	log.Print("Challenge request OK")

	return &p.ChallengeResponse{Ok: true}, nil
}

func main() {

	m = make(map[int][]byte)

	lis, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Panic(err)
	}

	s := grpc.NewServer()
	p.RegisterServerServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
