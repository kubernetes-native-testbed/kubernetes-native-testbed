package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

type deliveryStatusRepository interface {
	findByUUID(string) (*DeliveryStatus, error)
	store(*DeliveryStatus) (string, error)
	update(*DeliveryStatus) error
	deleteByUUID(string) error

	initDB() error
}

type deliveryStatusRepositoryCassandra struct {
	session  *gocql.Session
	keyspace string
	table    string
}

func (dr *deliveryStatusRepositoryCassandra) findByUUID(uuid string) (*DeliveryStatus, error) {
	var orderUUID, inquiryNumber string
	var status int
	var createdAt, updatedAt, deletedAt time.Time
	stmt := fmt.Sprintf(`SELECT order_uuid, status, inquiry_number, created_at, updated_at, deleted_at FROM %s.%s WHERE order_uuid=?`, dr.keyspace, dr.table)
	if err := dr.session.Query(stmt, uuid).RetryPolicy(&gocql.SimpleRetryPolicy{NumRetries: 3}).Scan(&orderUUID, &status, &inquiryNumber, &createdAt, &updatedAt, &deletedAt); err != nil {
		return nil, err
	}

	if updatedAt.Before(deletedAt) {
		return nil, fmt.Errorf("OrderUUID's delivery status is not found (OrderUUID=%s)", uuid)
	}

	return &DeliveryStatus{
		OrderUUID:     orderUUID,
		Status:        Status(status),
		InquiryNumber: inquiryNumber,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     &deletedAt,
	}, nil
}

func (dr *deliveryStatusRepositoryCassandra) store(r *DeliveryStatus) (string, error) {
	stmt := fmt.Sprintf(`INSERT INTO %s.%s(order_uuid, status, inquiry_number, created_at, updated_at) values (?, ?, ?, ?, ?)`, dr.keyspace, dr.table)
	if err := dr.session.Query(stmt, r.OrderUUID, r.Status, r.InquiryNumber, time.Now(), time.Now()).RetryPolicy(&gocql.SimpleRetryPolicy{NumRetries: 3}).Exec(); err != nil {
		return "", err
	}
	return r.OrderUUID, nil
}

func (dr *deliveryStatusRepositoryCassandra) update(r *DeliveryStatus) error {
	stmt := fmt.Sprintf(`UPDATE %s.%s SET status=?, inquiry_number=?, updated_at=? WHERE order_uuid=?`, dr.keyspace, dr.table)
	if err := dr.session.Query(stmt, r.Status, r.InquiryNumber, time.Now(), r.OrderUUID).RetryPolicy(&gocql.SimpleRetryPolicy{NumRetries: 3}).Exec(); err != nil {
		return err
	}
	return nil
}

func (dr *deliveryStatusRepositoryCassandra) deleteByUUID(uuid string) error {
	// logical delete
	stmt := fmt.Sprintf(`UPDATE %s.%s SET deleted_at=? WHERE order_uuid=?`, dr.keyspace, dr.table)
	if err := dr.session.Query(stmt, time.Now(), uuid).RetryPolicy(&gocql.SimpleRetryPolicy{NumRetries: 3}).Exec(); err != nil {
		return err
	}
	return nil
}

func (dr *deliveryStatusRepositoryCassandra) initDB() error {
	m, err := dr.session.KeyspaceMetadata(dr.keyspace)
	if err != nil {
		log.Printf("keyspace %s is not found", dr.keyspace)
		log.Printf("create keyspace %s", dr.keyspace)
		keyspaceOption := "replication = {'class': 'SimpleStrategy', 'replication_factor': 2}"
		if err := createKeyspace(dr.session, dr.keyspace, keyspaceOption); err != nil {
			return err
		}
		if m, err = dr.session.KeyspaceMetadata(dr.keyspace); err != nil {
			return err
		}
	}
	log.Printf("keyspace %s is ready", dr.keyspace)

	if _, ok := m.Tables[dr.table]; !ok {
		log.Printf("table %s is not found", dr.table)
		log.Printf("create table %s", dr.table)
		tableColumns := []string{
			"order_uuid text PRIMARY KEY",
			"status int",
			"inquiry_number text",
			"created_at timestamp",
			"updated_at timestamp",
			"deleted_at timestamp",
		}
		if err := createTable(dr.session, dr.keyspace, dr.table, tableColumns, ""); err != nil {
			return err
		}
		if m, err = dr.session.KeyspaceMetadata(dr.keyspace); err != nil {
			return err
		}
		if _, ok := m.Tables[dr.table]; !ok {
			return fmt.Errorf("failed to create table %s", dr.table)

		}
	}
	log.Printf("table %s is ready", dr.table)
	return nil
}

func createKeyspace(session *gocql.Session, keyspaceName, option string) error {
	var optstr string
	if option != "" {
		optstr = "WITH " + option
	}
	stmt := fmt.Sprintf(`CREATE KEYSPACE %s %s`, keyspaceName, optstr)
	if err := session.Query(stmt).RetryPolicy(&gocql.SimpleRetryPolicy{NumRetries: 3}).Exec(); err != nil {
		return err
	}
	return nil
}

func createTable(session *gocql.Session, keyspaceName, tableName string, columns []string, option string) error {
	colstr := strings.Join(columns, ",")
	var optstr string
	if option != "" {
		optstr = "WITH " + option
	}
	stmt := fmt.Sprintf(`CREATE TABLE %s.%s (%s) %s`, keyspaceName, tableName, colstr, optstr)
	if err := session.Query(stmt).RetryPolicy(&gocql.SimpleRetryPolicy{NumRetries: 3}).Exec(); err != nil {
		return err
	}
	return nil
}

type deliveryStatusRepositoryCassandraConfig struct {
	host     string
	port     int
	username string
	password string
	keyspace string
}

func (c *deliveryStatusRepositoryCassandraConfig) connect() (deliveryStatusRepository, func(), error) {
	cluster := gocql.NewCluster(fmt.Sprintf("%s:%d", c.host, c.port))
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.username,
		Password: c.password,
	}
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, nil, err
	}

	return &deliveryStatusRepositoryCassandra{session: session, keyspace: c.keyspace, table: c.keyspace}, session.Close, nil
}
