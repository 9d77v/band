package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
)

const _ttl = 10

type App struct {
	AppName     string
	ServiceName string
	ServerHost  string
	ServerPort  uint64
	MaxConns    int
	EtcdClient  *clientv3.Client
}

func NewApp(p Conf) App {
	cli, err := clientv3.NewFromURL(p.EtcdAddress)
	if err != nil {
		panic(err)
	}
	return App{
		AppName:     p.AppName,
		ServiceName: p.ServiceName,
		ServerPort:  p.ServerPort,
		ServerHost:  p.ServerHost,
		MaxConns:    p.MaxConns,
		EtcdClient:  cli,
	}
}

func (a *App) Register() {
	target := a.AppName + "/services/" + a.ServiceName
	em, err := endpoints.NewManager(a.EtcdClient, target)
	if err != nil {
		log.Panicln(err)
	}
	addr := fmt.Sprintf("%s:%d", a.ServerHost, a.ServerPort)
	key := target + "/" + strings.ReplaceAll(addr, ".", "-")
	lease, leaseID := getLeaseID(a.EtcdClient)
	err = em.AddEndpoint(context.TODO(), key, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(leaseID))
	if err != nil {
		log.Panicln("etce add endpoint failed")
	}
	leaseKeepAlive(lease, leaseID)
}

func (a *App) Deregister() {
	target := a.AppName + "/services/" + a.ServiceName
	em, _ := endpoints.NewManager(a.EtcdClient, target)
	addr := fmt.Sprintf("%s:%d", a.ServerHost, a.ServerPort)
	key := target + "/" + strings.ReplaceAll(addr, ".", "-")
	em.DeleteEndpoint(context.TODO(), key)
}

func (a *App) StartGrpcServer(register func(srv *grpc.Server), opt ...grpc.ServerOption) {
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.ServerPort))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer(opt...)
	register(srv)
	a.Register()
	go func() {
		errc <- srv.Serve(lis)
	}()
	log.Printf("exiting (%v)", <-errc)
	srv.GracefulStop()
	log.Println("exited")
}

func (a *App) StartHttpServer(handler http.Handler) {
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.ServerPort),
		Handler: handler,
	}
	go func() {
		errc <- srv.ListenAndServe()
		log.Printf("connect to http://localhost:%d/", a.ServerPort)
	}()
	log.Printf("exiting (%v)", <-errc)
	srvCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := srv.Shutdown(srvCtx)
	if err != nil {
		log.Println("server shut down error:", err)
	}
	log.Println("exited")
}
