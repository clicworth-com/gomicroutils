package grpcclient

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"

	"github.com/clicworth-com/gomicroutils/v3/genproto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GrpcClient struct {
	conn   *grpc.ClientConn
	client auth.AuthServiceClient
}

var grpcClient *GrpcClient

func GetAuthClient() *GrpcClient {
	return grpcClient
}

func StartAuthClient() {

	// read ca's cert
	caCert, err := os.ReadFile("/app/cert/ca-cert.pem")
	if err != nil {
		log.Fatal(caCert)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal(err)
	}

	//read client cert
	clientCert, err := tls.LoadX509KeyPair("/app/cert/client-cert.pem", "/app/cert/client-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	tlsCredential := credentials.NewTLS(config)

	// create client connection
	conn, err := grpc.Dial(
		"auth-srv:50051",
		grpc.WithTransportCredentials(tlsCredential),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := auth.NewAuthServiceClient(conn)

	grpcClient = &GrpcClient{
		conn:   conn,
		client: client,
	}
}

func StopGrpcClient() {
	grpcClient.conn.Close()
}
