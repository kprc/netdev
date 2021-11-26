package udpserver

import (
	"errors"
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"net"
	"strconv"
	"strings"
)

const (
	udpBuffer = 8192
)

type NetDevUdpServer struct {
	listenAddr string
	quit       chan struct{}
	udpServer  *net.UDPConn
	db         *mysqlconn.NetDevDbConn
}

func NewNetDevUdpServer(listenAddr string, db *mysqlconn.NetDevDbConn) *NetDevUdpServer {
	return &NetDevUdpServer{
		listenAddr: listenAddr,
		quit:       make(chan struct{}, 8),
		db:         db,
	}
}

func (us *NetDevUdpServer) Start() error {

	arr := strings.Split(us.listenAddr, ":")
	if len(arr) != 2 {
		return errors.New("address error")
	}

	port, err := strconv.Atoi(arr[1])
	if err != nil {
		return err
	}

	us.udpServer, err = net.ListenUDP("udp4", &net.UDPAddr{
		Port: port,
	})
	if err != nil {
		return err
	}

	fmt.Println("udp server start at:", port)

	go us.serve()

	return nil
}

func (us *NetDevUdpServer) serve() error {
	for {
		buf := make([]byte, udpBuffer)

		nr, addr, err := us.udpServer.ReadFrom(buf)
		if err != nil {
			continue
		}

		go RfidUdpMsg(us.db, buf[:nr], addr)

	}
}

func (us *NetDevUdpServer) ShutDown() error {
	var err error

	if us.udpServer != nil {
		err = us.udpServer.Close()
		us.udpServer = nil
	}

	return err
}
