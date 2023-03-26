package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	godatabase "go_database"
	"go_database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(godatabase.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(godatabase.GetConnection())
	comment, err := commentRepository.FindById(context.Background(), 18)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(godatabase.GetConnection())
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
