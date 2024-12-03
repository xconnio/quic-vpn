package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/quic-go/quic-go"
	"github.com/songgao/water"
)

func main() {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
	}
	quicConf := &quic.Config{
		EnableDatagrams: true,
	}

	quicConn, err := quic.DialAddr(context.Background(), "localhost:1234", tlsConf, quicConf)
	if err != nil {
		log.Fatal(err)
	}

	config := water.Config{
		DeviceType: water.TUN,
	}

	iface, err := water.New(config)
	if err != nil {
		log.Fatal(err)
	}

	// read IP packets from the other end and send to the TUN interface
	go func() {
		data, err := quicConn.ReceiveDatagram(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		_, err = iface.Write(data)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// read IP packets from the TUN interface and forward to the server
	data := make([]byte, 1500)
	for {
		count, err := iface.Read(data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(count, string(data[:count]))

		if err = quicConn.SendDatagram(data[:count]); err != nil {
			log.Fatal(err)
		}
	}
}
