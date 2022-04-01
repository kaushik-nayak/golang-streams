package main

import (
	"log"
	"net/http"
	"proto/client/handlers"
	pb "proto/proto"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}
	defer conn.Close()

	r := gin.Default()

	// Routes

	r.GET("/bookstore", func(c *gin.Context) {
		//page := c.Param("page")

		//page1, _ := strconv.ParseInt(page, 10, 64)
		//if err != nil {
		//	panic(err)
		//}

		//page, _ := strconv.ParseInt(page1, 10, 0)
		var page int64 = 3
		if res, err := handlers.GetAllBooks(conn, page); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"books": res.Books})
		}
	})

	r.GET("/bookstore/:id", func(c *gin.Context) {
		id := c.Param("id")
		if res, err := handlers.GetBookById(conn, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"book": res.Book})
		}
	})

	r.POST("/bookstore/insert", func(c *gin.Context) {
		release_year, err := strconv.Atoi(c.PostForm("release_year"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		book := &pb.Book{
			Title:       c.PostForm("title"),
			Description: c.PostForm("description"),
			Author: &pb.BookAuthor{
				Firstname: c.PostForm("firstname"),
				Lastname:  c.PostForm("lastname"),
			},
			ReleaseYear: int32(release_year),
		}

		if res, err := handlers.InsertBook(conn, &pb.BookReq{Book: book}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"book": res.Book})
		}
	})

	r.POST("/bookstore/update/:id", func(c *gin.Context) {
		id := c.Param("id")
		release_year, _ := strconv.Atoi(c.PostForm("release_year"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		book := &pb.Book{
			Title:       c.PostForm("title"),
			Description: c.PostForm("description"),
			Author: &pb.BookAuthor{
				Firstname: c.PostForm("firstname"),
				Lastname:  c.PostForm("lastname"),
			},
			ReleaseYear: int32(release_year),
		}

		if res, err := handlers.UpdateBook(conn, &pb.UpdateBookReq{Id: id, Book: book}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"book": res.Book})
		}
	})

	r.POST("/bookstore/delete/:id", func(c *gin.Context) {
		id := c.Param("id")

		if res, err := handlers.DeleteBook(conn, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": res.Success})
		}
	})

	// Run server

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
