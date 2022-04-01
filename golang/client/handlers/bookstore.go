package handlers

import (
	"context"
	pb "proto/proto"

	"google.golang.org/grpc"
)

func InsertBook(conn *grpc.ClientConn, book *pb.BookReq) (*pb.BookRes, error) {
	client := pb.NewBookstoreClient(conn)
	res, err := client.CreateBook(context.Background(), book)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetBookById(conn *grpc.ClientConn, id string) (*pb.BookRes, error) {
	client := pb.NewBookstoreClient(conn)
	res, err := client.GetBook(context.Background(), &pb.BookQuery{Id: id})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetAllBooks(conn *grpc.ClientConn, page int64) (*pb.BookList, error) {
	client := pb.NewBookstoreClient(conn)
	res, err := client.GetAllBooks(context.Background(), &pb.EmptyQuery{Page:page})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateBook(conn *grpc.ClientConn, in *pb.UpdateBookReq) (*pb.BookRes, error) {
	client := pb.NewBookstoreClient(conn)
	res, err := client.UpdateBook(context.Background(), in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeleteBook(conn *grpc.ClientConn, id string) (*pb.DeleteBookRes, error) {
	client := pb.NewBookstoreClient(conn)
	res, err := client.DeleteBook(context.Background(), &pb.BookQuery{Id: id})
	if err != nil {
		return nil, err
	}

	return res, nil
}
