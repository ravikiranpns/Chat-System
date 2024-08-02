package db

import (
	"fmt"
	"github.com/gocql/gocql"
	"strconv"
)

var session *gocql.Session

func InitCassandra(host, port string) error {
	cluster := gocql.NewCluster(host)
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("invalid port number: %v", err)
	}
	cluster.Port = portInt
	cluster.Keyspace = "chat"
	cluster.Consistency = gocql.Quorum

	session, err = cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("failed to connect to Cassandra: %v", err)
	}

	return nil
}

func GetSession() *gocql.Session {
	return session
}

func CreateKeyspaceAndTables() error {
	err := session.Query(`CREATE KEYSPACE IF NOT EXISTS chat WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`).Exec()
	if err != nil {
		return err
	}

	err = session.Query(`CREATE TABLE IF NOT EXISTS chat.users (username text PRIMARY KEY, password text)`).Exec()
	if err != nil {
		return err
	}

	err = session.Query(`CREATE TABLE IF NOT EXISTS chat.messages (id UUID PRIMARY KEY, sender text, recipient text, content text, timestamp timestamp)`).Exec()
	if err != nil {
		return err
	}

	return nil
}
