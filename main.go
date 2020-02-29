package main

import (
	"context"
	"fmt"

	pb "github.com/wealdtech/walletd/pb/v1"
	"google.golang.org/grpc"
)

func main() {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial("localhost:12346", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	account := "My validators/account 1"
	passphrase := []byte("secret")

	accountClient := pb.NewAccountManagerClient(conn)
	unlockReq := &pb.UnlockAccountRequest{
		Account:    account,
		Passphrase: passphrase,
	}
	if _, err := accountClient.Unlock(context.Background(), unlockReq); err != nil {
		panic(err)
	}

	signClient := pb.NewSignerClient(conn)
	signRawReq := &pb.SignRequest{
		Data:    []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
		Account: "My validators/account 1",
	}
	resp, err := signClient.Sign(context.Background(), signRawReq)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#x\n", resp.Signature)
}
