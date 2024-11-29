package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/quic-go/quic-go"
	"log"
	"os"
)

func main() {
	certPool := x509.NewCertPool()
	cert, err := os.ReadFile("/home/om26er/scm/xconnio/quic-vpn/cert.pem")
	if err != nil {
		panic(err)
	}
	certPool.AppendCertsFromPEM(cert)

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		RootCAs:            certPool,
	}
	quicConf := &quic.Config{}

	quicConn, err := quic.DialAddr(context.Background(), "localhost:1234", tlsConf, quicConf)
	if err != nil {
		log.Fatal(err)
	}

	stream, err := quicConn.OpenStreamSync(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	println(stream.StreamID())
}
