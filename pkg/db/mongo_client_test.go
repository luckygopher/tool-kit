package db

import "testing"

func TestInitMongoClient(t *testing.T) {
	if err := InitMongoClient(MongoConfig{
		UserName:       "",
		PassWord:       "",
		Host:           "127.0.0.1:27017",
		AuthSource:     "",
		DefaultTimeout: "20s",
		MaxPoolSize:    10,
	}); err != nil {
		t.Fatal(err)
	}
}
