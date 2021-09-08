package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/korovkin/gotils"

	"blackbox/gen/service_v1"
)

var VERSION string = "000.001.001"

//////////////////////////////////////////
// Blackbox server handler: //////////////
//////////////////////////////////////////
type BlackboxHandler struct {
}

func NewBlackboxHandler() *BlackboxHandler {
	return &BlackboxHandler{}
}

func (p *BlackboxHandler) Ping(ctx context.Context) (err error) {
	log.Println("received ping()")
	return nil
}

func (p *BlackboxHandler) GetVersion(ctx context.Context) (ver string, err error) {
	log.Println("received GetVersion()")
	return VERSION, nil
}

func (p *BlackboxHandler) GetName(ctx context.Context) (ver string, err error) {
	log.Println("received GetName()")
	return "blackbox", nil
}

func (p *BlackboxHandler) LogLocation(ctx context.Context, loc *service_v1.Location) (_err error) {
	log.Println("received LogLocation - ", gotils.ToJSONString(loc))
	return nil
}

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket(addr)

	if err != nil {
		return err
	}
	handler := NewBlackboxHandler()
	processor := service_v1.NewBlackboxProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println(" => Running server on: ", addr)
	return server.Serve()
}

//////////////////////////////////////////
// Blackbox client: //////////////////////
//////////////////////////////////////////

func runClient(
	transportFactory thrift.TTransportFactory,
	protocolFactory thrift.TProtocolFactory,
	addr string) error {

	var transport thrift.TTransport
	var err error

	transport, err = thrift.NewTSocketConf(addr, &thrift.TConfiguration{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		}})
	gotils.CheckFatal(err)

	transport, err = transportFactory.GetTransport(transport)
	gotils.CheckFatal(err)
	err = transport.Open()
	gotils.CheckFatal(err)
	defer transport.Close()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := service_v1.NewBlackboxClient(thrift.NewTStandardClient(iprot, oprot))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
	defer cancel()

	// ping:
	log.Println("ping() - start")
	client.Ping(ctx)
	log.Println("ping() - done")

	// get version:
	var ver string
	var name string
	log.Println("get_version() - start")
	ver, err = client.GetVersion(ctx)
	name, err = client.GetName(ctx)
	log.Println("get_version() - ", ver, name)

	// log location:
	err = client.LogLocation(ctx, &service_v1.Location{
		TimestampUnixSec: float64(time.Now().Unix()),
		LongitudeDegrees: -122.0,
		LatitudeDegrees:  33.0,
	})
	log.Println("LogLocation()")

	return err
}

// Main:

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	log.Println("hello")

	version := flag.Bool(
		"version",
		false,
		"print the version and exit")

	server := flag.Bool(
		"server",
		false,
		"run the server")

	flag.Parse()

	if *version {
		log.Println(VERSION)
		return
	}

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryConf(nil)

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	if *server {
		e := runServer(transportFactory, protocolFactory, "localhost:9003")
		gotils.CheckFatal(e)
	} else {
		e := runClient(transportFactory, protocolFactory, "localhost:9003")
		gotils.CheckFatal(e)
	}

	log.Println("=> done:")
}
