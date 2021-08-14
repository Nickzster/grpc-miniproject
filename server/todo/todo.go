package todo

// protoc --go_out=. --go-grpc_out=. todo.proto

import (
	"fmt"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedTodoServer
}

var todos []*TodoItem

func (s *Server) CreateTodo(ctx context.Context, in *TodoItem) (*TodoItem, error) {
	fmt.Printf("Receive message body from client: %s\n", in.Text)
	todos = append(todos, in)
	return in, nil
}
func (s *Server) ReadTodos(ctx context.Context, in *Empty) (*TodoItems, error) {
	fmt.Println("Requested to read entire todo list!")
	foobar := &todos
	bar := TodoItems{Items: *foobar}
	return &bar, nil
}

