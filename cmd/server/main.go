package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/quic-go/quic-go"
)

const (
	Cert = "/home/om26er/scm/xconnio/quic-vpn/cert.pem"
	Key  = "/home/om26er/scm/xconnio/quic-vpn/key.pem"
)

func main() {
	cert, err := tls.LoadX509KeyPair(Cert, Key)
	if err != nil {
		panic(err)
	}

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}
	quicConf := &quic.Config{}

	ln, err := quic.ListenAddr("localhost:1234", tlsConf, quicConf)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		go func(quicConn quic.Connection) {
			stream, err := conn.AcceptStream(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("new stream", stream.StreamID())
		}(conn)
	}
}
