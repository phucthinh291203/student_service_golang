package database

import (
	"context"
	"grpc-project/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	StudentData models.StudentData
)

func ConnectToDatabase() {
	godotenv.Load()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Fatal("Kết nối đến database thất bại:", err)
		return
	}

	// Kiểm tra kết nối
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Kết nối đến database thất bại:", err)
		return
	}

	mongoClient = client
	StudentData = models.StudentData{
		StudentCollection: mongoClient.Database(GetDBName()).Collection("studentCollection"),
	}

	if mongoClient == nil {
		log.Fatal("mongoClient chưa được khởi tạo")
	} else {
		log.Println("mongoClient đã khởi tạo thành công")
	}
	log.Print("Kết nối đến database thành công!")
}

func GetDBName() string {
	return os.Getenv("dbName")
}

// Hàm trả về collection của student
func GetStudentCollection() *mongo.Collection {
	return StudentData.StudentCollection
}

