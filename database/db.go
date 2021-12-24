package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	proto "github.com/duckhue01/golang_test/proto/v2"
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

	// test if connect successfully
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
	// convert to time.Time
	createAt, err := ptypes.Timestamp(req.Todo.CreateAt)
	if err != nil {
		return err
	}
	updateAt, err := ptypes.Timestamp(req.Todo.UpdateAt)
	if err != nil {
		return err
	}

	// add todo to  Todo table
	_, err = c.ExecContext(ctx, "INSERT INTO Todo(`Id`, `Title`, `Description`, `createAt`, `UpdateAt`, `Status`) VALUES(?, ?, ?, ?, ?, ?)",
		req.Todo.GetId(), req.Todo.GetTitle(), req.Todo.GetDescription(), createAt, updateAt, req.Todo.GetStatus())
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	for _, v := range req.Todo.Tags {

		// add tag to Tag table
		_, err = c.ExecContext(ctx, "CALL AddTagIfNotExist(?)", v)
		if err != nil {
			return err
		}

		// add tag and todo to TagTodo table
		_, err = c.ExecContext(ctx, "CALL AddTagTodoIfNotExist(?, ?)", v, req.Todo.GetId())
		if err != nil {
			return err
		}
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
	if err := rows.Scan(&td.Id, &td.Title, &td.Description, &createAt, &updateAt, &td.Status, &td.Order); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Todo row-> "+err.Error())
	}

	// get tag to this todos
	c, err = d.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	tagsRows, err := c.QueryContext(ctx, "SELECT Tag FROM TagTodo WHERE `Id` = ?", td.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Todo-> "+err.Error())
	}
	defer tagsRows.Close()
	tags := []string{}

	for tagsRows.Next() {
		var temp string
		if err := tagsRows.Scan(&temp); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Todo row-> "+err.Error())
		}

		tags = append(tags, temp)
	}

	td.Tags = tags
	td.CreateAt = timestamppb.New(createAt)
	td.UpdateAt = timestamppb.New(updateAt)

	return &td, nil
}

func (d *DatabaseStore) GetAll(ctx context.Context, req *proto.GetAllRequest) ([]*proto.Todo, error) {

	c, err := d.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// default 10
	var lim int32
	if req.Pag == 0 {
		lim = 10
	} else {
		lim = req.Pag
	}

	var rows *sql.Rows

	// divide into 4 cases to avoid the complexity of concatnation
	// but it also too long
	if req.Status == nil && req.Tags == nil {
		rows, err = c.QueryContext(ctx, "SELECT * FROM Todo ORDER BY `Order` DESC LIMIT ?", lim)
	}

	if req.Status != nil && req.Tags == nil {

		temp := ""
		for _, v := range req.Status {
			switch v {
			case 0:
				temp += "0"
			case 1:
				temp += "1"
			case 2:
				temp += "2"

			}
		}
		temp = "[" + temp + "]"
		rows, err = c.QueryContext(ctx, "SELECT * FROM Todo WHERE Status REGEXP ? ORDER BY `Order` DESC LIMIT ?", temp, lim)

	}

	if req.Tags != nil && req.Status == nil {

		temp := make([]string, len(req.Tags))
		for i := range req.Tags {
			temp[i] = fmt.Sprintf("%s%s%s", "^", req.Tags[i], "$")
		}

		rows, err = c.QueryContext(ctx, "SELECT * FROM Todo WHERE Id IN (SELECT DISTINCT `Id`  FROM `TagTodo` WHERE `Tag` REGEXP ? ) ORDER BY `Order` DESC LIMIT ?", strings.Join(temp, "|"), lim)
	}

	if req.Tags != nil && req.Status != nil {
		tTemp := make([]string, len(req.Tags))
		for i := range req.Tags {
			tTemp[i] = fmt.Sprintf("%s%s%s", "^", req.Tags[i], "$")
		}

		sTemp := ""
		for _, v := range req.Status {
			switch v {
			case 0:
				sTemp += "0"
			case 1:
				sTemp += "1"
			case 2:
				sTemp += "2"

			}
		}
		sTemp = "[" + sTemp + "]"

		rows, err = c.QueryContext(ctx, "SELECT * FROM Todo WHERE Id IN (SELECT DISTINCT `Id`  FROM `TagTodo` WHERE `Tag` REGEXP ? ) AND Status REGEXP ? ORDER BY `Order` DESC LIMIT ?", strings.Join(tTemp, "|"), sTemp, lim)
	}

	// check error for all 4 query
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Todo-> "+err.Error())
	}
	defer rows.Close()

	var createAt time.Time
	var updateAt time.Time
	list := []*proto.Todo{}
	for rows.Next() {
		td := &proto.Todo{}
		if err := rows.Scan(&td.Id, &td.Title, &td.Description, &createAt, &updateAt, &td.Status, &td.Order); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Todo row-> "+err.Error())
		}

		// get tags for each id
		c, err := d.connect(ctx)
		if err != nil {
			return nil, err
		}
		defer c.Close()
		tagsRows, err := c.QueryContext(ctx, "SELECT Tag FROM TagTodo WHERE `Id` = ?", td.Id)
		if err != nil {
			return nil, status.Error(codes.Unknown, "failed to select from Todo-> "+err.Error())
		}
		defer tagsRows.Close()
		tags := []string{}

		for tagsRows.Next() {
			var temp string
			if err := tagsRows.Scan(&temp); err != nil {
				return nil, status.Error(codes.Unknown, "failed to retrieve field values from Todo row-> "+err.Error())
			}

			tags = append(tags, temp)
		}
		td.Tags = tags
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

	res, err := c.ExecContext(ctx, "UPDATE Todo SET `Title`=?, `Description`=?, `UpdateAt`=?, `Status`=? WHERE `ID`=?",
		req.Todo.Title, req.Todo.Description, updateAt, req.Todo.Status, req.Todo.Id)
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

func (d *DatabaseStore) Reorder(ctx context.Context, req *proto.ReorderRequest) error {

	c, err := d.connect(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	// get the order of Todo
	res, err := c.QueryContext(ctx, "SELECT `Order` FROM `Todo` WHERE `Id` = ? LIMIT 1", req.Id)
	if err != nil {
		return status.Error(codes.Unknown, "failed to reorder Todo-> "+err.Error())
	}
	defer res.Close()
	if !res.Next() {
		if err := res.Err(); err != nil {
			return status.Error(codes.Unknown, "failed to retrieve data from Todo-> "+err.Error())
		}
		return status.Error(codes.NotFound, fmt.Sprintf("Todo with ID='%d' is not found", req.Id))
	}
	var start int32

	if err := res.Scan(&start); err != nil {
		return status.Error(codes.Unknown, "failed to retrieve field values from Todo row-> "+err.Error())
	}

	c, err = d.connect(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	// change order of Todo
	res1, err := c.ExecContext(ctx, "CALL ReorderTodo(?, ?)", start, req.Pos)
	if err != nil {
		return status.Error(codes.Unknown, "failed to reorder Todo-> "+err.Error())
	}
	_, err = res1.RowsAffected()
	if err != nil {
		return status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	return nil
}
