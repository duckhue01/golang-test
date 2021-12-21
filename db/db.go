package db

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/duckhue01/golang_test/proto/v1"
)

type DatabaseStore struct {
	db *sql.DB
}

func NewDatabaseStore() *DatabaseStore {
	cfg := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		DBName:               os.Getenv("MYSQL_DATABASE"),
		Net:                  "tcp",
		Addr:                 "db:3306",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err, "ping")
	}

	return &DatabaseStore{
		db: db,
	}
}

func (d *DatabaseStore) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := d.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

func (d *DatabaseStore) Add(ctx context.Context, req *proto.AddRequest) error {
	// get SQL connection from pool
	c, err := d.connect(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	insertAt, err := ptypes.Timestamp(req.Todo.CreateAt)
	if err != nil {
		return err
	}

	updateAt, err := ptypes.Timestamp(req.Todo.UpdateAt)
	if err != nil {
		return err
	}

	// insert Todo entity data
	_, err = c.ExecContext(ctx, "INSERT INTO Todo(`Title`, `Description`, `InsertAt`, `UpdateAt`, `IsDone`) VALUES(?, ?, ?, ?, ?)",
		req.Todo.Title, req.Todo.Description, insertAt, updateAt, req.Todo.IsDone)
	if err != nil {
		return err
	}

	return nil
}
func (d *DatabaseStore) GetOne(ctx context.Context, req *proto.AddRequest) error {



	return nil
}