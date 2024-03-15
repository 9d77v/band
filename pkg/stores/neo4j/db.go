package neo4j

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var (
	client *Neo4jClient
	once   sync.Once
)

type Neo4jClient struct {
	client neo4j.DriverWithContext
	conf   Conf
}

func (c *Neo4jClient) NewSession(ctx context.Context) neo4j.SessionWithContext {
	return c.client.NewSession(ctx, neo4j.SessionConfig{DatabaseName: c.conf.Username})
}

func NewNeo4jClient(conf Conf) (*Neo4jClient, error) {
	client, err := neo4j.NewDriverWithContext(fmt.Sprintf("neo4j://%s:%d",
		conf.Host, conf.Port),
		neo4j.BasicAuth(conf.Username, conf.Password, ""))
	if err != nil {
		panic(err)
	}
	log.Println("connected to Neo4jClient:", client)
	return &Neo4jClient{
		client: client,
		conf:   conf,
	}, nil
}

func Neo4jClientSingleton(conf Conf) (*Neo4jClient, error) {
	var err error
	once.Do(func() {
		client, err = NewNeo4jClient(conf)
	})
	return client, err
}

// CreateUniqueConstraint 创建唯一性约束
func (c *Neo4jClient) CreateUniqueConstraint(label, property string) error {
	ctx := context.Background()
	session := c.client.NewSession(ctx,
		neo4j.SessionConfig{
			DatabaseName: c.conf.Username,
			AccessMode:   neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.Run(ctx, fmt.Sprintf("CREATE CONSTRAINT ON (n:%s) ASSERT n.%s IS UNIQUE", label, property), nil)
	return err
}

// Transaction 事务
func (c *Neo4jClient) Transaction(ctx context.Context, fc func(neo4j.ExplicitTransaction) error, opts ...*sql.TxOptions) (err error) {
	session := c.NewSession(ctx)
	defer session.Close(ctx)
	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	if err = fc(tx); err == nil {
		return tx.Commit(ctx)
	}
	return
}
