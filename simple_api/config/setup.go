package config

import (
	"os"

	"context"
	"example/simple_api/models/entity"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var BucketHandle *storage.BucketHandle

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@dm1n1994@tcp(localhost:3306)/cource_net"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&entity.Comment{}, &entity.Profile{})
	DB = database
}

func ConnectFirebaseStorage() {
	opt := option.WithCredentialsFile("config/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)

	}

	client, err := app.Storage(context.TODO())
	if err != nil {
		panic(err)

	}

	bucketHandle, err := client.Bucket(os.Getenv("BUCKET_NAME"))
	if err != nil {
		panic(err)

	}
	BucketHandle = bucketHandle
}
