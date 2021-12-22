package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

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
		ParseTime:            true,
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
	c, err := d.connect(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	createAt, err := ptypes.Timestamp(req.Todo.CreateAt)
	if err != nil {
		return err
	}

	updateAt, err := ptypes.Timestamp(req.Todo.UpdateAt)
	if err != nil {
		return err
	}

	_, err = c.ExecContext(ctx, "INSERT INTO Todo(`Title`, `Description`, `createAt`, `UpdateAt`, `IsDone`) VALUES(?, ?, ?, ?, ?)",
		req.Todo.GetTitle(), req.Todo.Description, createAt, updateAt, req.Todo.IsDone)
	if err != nil {
		return err
	}

	return nil
}
func (d *DatabaseStore) GetOne(ctx context.Context, id int32) (*proto.Todo, error) {

	c, err := d.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	rows, err := c.QueryContext(ctx, "SELECT * FROM Todo WHERE `ID`=?", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Todo-> "+err.Error())
	}

	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Todo-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Todo with ID='%d' is not found", id))
	}
	var td proto.Todo
	var createAt time.Time
	var updateAt time.Time
	if err := rows.Scan(&td.Id, &td.Title, &td.Description, &createAt, &updateAt, &td.IsDone); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Todo row-> "+err.Error())
	}
	td.CreateAt = timestamppb.New(createAt)

	td.UpdateAt = timestamppb.New(updateAt)

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Todo rows with ID='%d'", id))
	}

	return &td, nil
}

func (d *DatabaseStore) GetAll(ctx context.Context) ([]*proto.Todo, error) {

	c, err := d.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()


	rows, err := c.QueryContext(ctx, "SELECT * FROM Todo")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Todo-> "+err.Error())
	}
	defer rows.Close()

	var createAt time.Time
	var updateAt time.Time
	list := []*proto.Todo{}
	for rows.Next() {
		td := &proto.Todo{}
		if err := rows.Scan(&td.Id, &td.Title, &td.Description, &createAt, &updateAt, &td.IsDone); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Todo row-> "+err.Error())
		}
		td.CreateAt = timestamppb.New(createAt)
		td.UpdateAt = timestamppb.New(updateAt)
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Todo-> "+err.Error())
	}

	return list, nil
}

func (d *DatabaseStore) Update(ctx context.Context, req *proto.UpdateRequest) error {
	c, err := d.connect(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	updateAt, err := ptypes.Timestamp(req.Todo.UpdateAt)
	if err != nil {
		return status.Error(codes.InvalidArgument, "updateAt field has invalid format-> "+err.Error())
	}

	res, err := c.ExecContext(ctx, "UPDATE Todo SET `Title`=?, `Description`=?, `UpdateAt`=?, `IsDone`=? WHERE `ID`=?",
		req.Todo.Title, req.Todo.Description, updateAt, req.Todo.IsDone, req.Todo.Id)
	if err != nil {
		return status.Error(codes.Unknown, "failed to update Todo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return status.Error(codes.NotFound, fmt.Sprintf("Todo with ID='%d' is not found",
			req.Todo.Id))
	}

	return nil
}

func (d *DatabaseStore) Delete(ctx context.Context, id int32) error {

	c, err := d.connect(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx, "DELETE FROM Todo WHERE `ID`=?", id)
	if err != nil {
		return status.Error(codes.Unknown, "failed to delete Todo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return status.Error(codes.NotFound, fmt.Sprintf("Todo with ID='%d' is not found", id))
	}

	return nil
}
