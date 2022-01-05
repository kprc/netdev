package webserver

import (
	"context"
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/api"
	"net"
	"net/http"
	"regexp"
	"time"
)

type NetDevWebServer struct {
	listenAddr string
	quit       chan struct{}
	server     *http.Server
	db         *mysqlconn.NetDevDbConn
}

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHandler struct {
	routes []*route
}

func (rh *RegexpHandler) Handle(pattern string, handler http.Handler) {
	rh.routes = append(rh.routes, &route{pattern: regexp.MustCompilePOSIX(pattern), handler: handler})
}

func (rh *RegexpHandler) HandleFunc(pattern string, handleFunc func(http.ResponseWriter, *http.Request)) {
	rh.routes = append(rh.routes, &route{pattern: regexp.MustCompilePOSIX(pattern), handler: http.HandlerFunc(handleFunc)})
}

func (rh *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range rh.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

func NewNetDevWebServer(networkAddr string, db *mysqlconn.NetDevDbConn) *NetDevWebServer {
	ws := &NetDevWebServer{
		listenAddr: networkAddr,
		quit:       make(chan struct{}, 8),
		db:         db,
	}

	return ws.init()
}

func (ws *NetDevWebServer) init() *NetDevWebServer {
	rh := &RegexpHandler{
		routes: make([]*route, 0),
	}

	wapi := api.NewWebApi(ws.db)

	rh.HandleFunc(api.NetDevPathStr(api.FoodTowerPath), wapi.FoodTower)
	rh.HandleFunc(api.NetDevPathStr(api.WaterPath), wapi.Water)
	rh.HandleFunc(api.NetDevPathStr(api.WeighPath), wapi.Weigh)
	rh.HandleFunc(api.NetDevPathStr(api.UniphasePath), wapi.UniPhase)
	rh.HandleFunc(api.NetDevPathStr(api.TriphasePath), wapi.Triphase)
	rh.HandleFunc(api.NetDevPathStr(api.UploadFile),wapi.UploadFile)
	rh.HandleFunc(api.SummaryPathStr(api.IndexSource),wapi.IndexSource)
	rh.HandleFunc(api.ProxyAPIPath(api.PigPirce),api.Proxy)

	server := &http.Server{
		Handler: rh,
	}

	ws.server = server

	return ws
}

func (ws *NetDevWebServer) Start() error {
	if l, err := net.Listen("tcp4", ws.listenAddr); err != nil {
		return err
	} else {
		fmt.Println("web server start at :", ws.listenAddr)
		go ws.server.Serve(l)
		return nil
	}
}

func (ws *NetDevWebServer) ShutDown() error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ws.server.Shutdown(ctx)
}
