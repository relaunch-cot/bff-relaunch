package resource

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/relaunch-cot/bff/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func openGrpcClientConn[V any](url string, f func(conn grpc.ClientConnInterface) V) V {
	var conn *grpc.ClientConn
	var dial grpc.DialOption

	if config.IS_INSECURE == "true" {
		dial = grpc.WithTransportCredentials(insecure.NewCredentials())
	} else {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			panic("Cannot load root CA certs")
		}

		creds := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})

		dial = grpc.WithTransportCredentials(creds)
	}

	conn, err := grpc.NewClient(url, dial)
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	c := f(conn)

	return c
}
