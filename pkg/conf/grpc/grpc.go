package grpc

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
	resolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	Conf Conf
	Conn *grpc.ClientConn
}

func NewGrpcClient(conf Conf) *GrpcClient {
	etcdAddress := strings.ReplaceAll(conf.EtcdAddress, "http", "etcd")
	rawURL := fmt.Sprintf("%s/%s/services/%s", etcdAddress, conf.AppName, conf.ServiceName)
	u, err := url.Parse(rawURL)
	if err != nil {
		log.Panicln("url parse error:", err)
	}
	cli, cerr := clientv3.NewFromURL(u.Host)
	if cerr != nil {
		log.Panicln("init etcd client failed", cerr)
	}
	etcdResolver, err := resolver.NewBuilder(cli)
	if err != nil {
		log.Panicln("init etcd resolver failed", cerr)
	}
	conn, err := grpc.Dial(rawURL,
		grpc.WithResolvers(etcdResolver),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [ { "round_robin": {} } ]}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicln("dial failed", err)
	}
	log.Println("connected to ", rawURL)
	return &GrpcClient{
		Conf: conf,
		Conn: conn,
	}
}
