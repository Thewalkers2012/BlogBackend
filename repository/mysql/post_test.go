package mysql

import (
	"testing"

	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/settings"
)

func init() {
	dbConfig := settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "root",
		DBName:       "blog_database",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&dbConfig)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := &models.Post{
		ID:          10,
		AuthorID:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}

	err := CreatePost(post)
	if err != nil {
		t.Fatalf("Create insert record into mysql failed, err: %v\n", err)
	}

	t.Logf("CreatePost insert record into mysql failed")
}
