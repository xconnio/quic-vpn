package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/quic-go/quic-go"
)

const (
	Cert = "cert.pem"
	Key  = "key.pem"
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
	quicConf := &quic.Config{
		EnableDatagrams: true,
	}

	//config := water.Config{
	//	DeviceType: water.TUN,
	//}

	//iface, err := water.New(config)
	//if err != nil {
	//	log.Fatal(err)
	//}

	ln, err := quic.ListenAddr("localhost:1234", tlsConf, quicConf)
	if err != nil {
		log.Fatal(err)
	}

	for {
		quicConn, err := ln.Accept(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		go func(quicConn quic.Connection) {
			for {
				data, err := quicConn.ReceiveDatagram(context.Background())
				if err != nil {
					log.Fatal(err)
				}

				p := gopacket.NewPacket(data, layers.LayerTypeIPv6, gopacket.DecodeStreamsAsDatagrams)
				fmt.Println(p)

				//if _, err = iface.Write(data); err != nil {
				//	log.Fatal(err)
				//}
			}
		}(quicConn)

		//go func() {
		//	// read IP packets from the TUN interface and forward to the server
		//	data := make([]byte, 1500)
		//	for {
		//		count, err := iface.Read(data)
		//		if err != nil {
		//			log.Fatal(err)
		//		}
		//
		//		fmt.Println(count, string(data[:count]))
		//
		//		if err = quicConn.SendDatagram(data[:count]); err != nil {
		//			log.Fatal(err)
		//		}
		//	}
		//}()
	}
}
