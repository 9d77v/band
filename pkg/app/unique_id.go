package app

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/bwmarrin/snowflake"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type UniqueID struct {
	NodeID int64
	node   *snowflake.Node
}

var (
	client *UniqueID
	once   sync.Once
)

const (
	maxNodes = 1024
)

func NewUniqueID(p Conf) *UniqueID {
	once.Do(func() {
		client = &UniqueID{
			NodeID: getNodeID(p),
		}
		snowflake.Epoch = 1713196800000 //2024-04-16 00:00:00 GMT+8
		node, err := snowflake.NewNode(client.NodeID)
		if err != nil {
			log.Panicln(err)
		}
		client.node = node
	})
	return client
}

func (u *UniqueID) GetID() int64 {
	return u.node.Generate().Int64()
}

func getNodeID(p Conf) int64 {
	// single node default 0
	if p.EtcdAddress == "" {
		return 0
	}
	cli, err := clientv3.NewFromURL(p.EtcdAddress)
	if err != nil {
		log.Panicln(err)
	}
	etcdKeyPrefix := p.AppName + "/nodes"
	nodesResp := checkAndGetNodes(cli, etcdKeyPrefix)
	lease, leaseID := getLeaseID(cli)
	existNodeMap := make(map[int]int)
	for _, ev := range nodesResp.Kvs {
		num, _ := strconv.Atoi(string(ev.Value))
		existNodeMap[num] = num
	}
	i := 0
	for ; i < maxNodes; i++ {
		if _, ok := existNodeMap[i]; !ok {
			key := etcdKeyPrefix + "/" + strconv.Itoa(i)
			txnResp, err := cli.Txn(context.Background()).
				If(clientv3.Compare(clientv3.LeaseValue(key), "=", clientv3.NoLease)).
				Then(clientv3.OpPut(key, strconv.Itoa(i), clientv3.WithLease(leaseID))).
				Commit()
			if err != nil {
				log.Panicln(err)
			}
			if txnResp.Succeeded {
				break
			}
		}
	}
	if i == maxNodes {
		log.Panicln("Machine ID has reached the limit")
	}
	leaseKeepAlive(lease, leaseID)
	return int64(i)
}

func checkAndGetNodes(cli *clientv3.Client, etcdKeyPrefix string) *clientv3.GetResponse {
	resp, err := cli.Get(context.Background(), etcdKeyPrefix, clientv3.WithPrefix())
	if err != nil {
		log.Panicln(err)
	}
	if len(resp.Kvs) >= maxNodes {
		log.Panicln("Machine ID has reached the limit")
	}
	return resp
}

func getLeaseID(cli *clientv3.Client) (clientv3.Lease, clientv3.LeaseID) {
	lease := clientv3.NewLease(cli)
	leaseGrantResp, err := lease.Grant(context.TODO(), _ttl)
	if err != nil {
		log.Panicln(err)
	}
	return lease, leaseGrantResp.ID
}

func leaseKeepAlive(lease clientv3.Lease, leaseID clientv3.LeaseID) {
	leaseRespChan, err := lease.KeepAlive(context.TODO(), leaseID)
	if err != nil {
		log.Panicf("lease failed:%s\n", err.Error())
	}
	go func() {
		for {
			leaseKeepResp := <-leaseRespChan
			if leaseKeepResp == nil {
				fmt.Printf("lease closed\n")
				return
			}
		}
	}()
}
