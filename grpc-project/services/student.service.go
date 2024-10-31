package services

import (
	"context"
	database "grpc-project/database"
	"grpc-project/models"
	pb "grpc-project/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StudentService struct {
	pb.UnimplementedStudentServiceServer
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (service *StudentService) CreateNewStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentResponse, error) {
	// Kiểm tra xem tất cả thông tin đã được cung cấp
	if req.Name == "" {
		return &pb.CreateStudentResponse{Message: "Name is required"}, nil
	}
	if req.DateOfBirth == "" {
		return &pb.CreateStudentResponse{Message: "Date of birth is required"}, nil
	}
	if req.Gender == "" {
		return &pb.CreateStudentResponse{Message: "Gender is required"}, nil
	}

	student := models.Student{
		Name:        req.Name,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
	}

	// Chèn học sinh vào collection
	_, err := database.GetStudentCollection().InsertOne(context.TODO(), student)
	if err != nil {
		return &pb.CreateStudentResponse{Message: "Failed to create student"}, err
	}
	return &pb.CreateStudentResponse{Message: "Them student thanh cong"}, nil
}

func (service *StudentService) GetAllStudent(ctx context.Context, emptypb *emptypb.Empty) (*pb.GetAllStudentResponse, error) {
	cursors, err := database.GetStudentCollection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var studentList []*pb.Student
	for cursors.Next(context.TODO()) {
		var student models.Student
		cursors.Decode(&student)

		studentList = append(studentList, &pb.Student{
			Id:          student.ID.Hex(),
			Name:        student.Name,
			DateOfBirth: student.DateOfBirth,
			Gender:      student.Gender,
			ClassId:     student.ClassID.Hex(),
		})
	}
	if len(studentList) == 0 {
		return &pb.GetAllStudentResponse{Message: "Khong tim thay data nao", Students: nil}, nil
	}

	return &pb.GetAllStudentResponse{Message: "Tim thay tat ca data", Students: studentList}, nil
}

func (service *StudentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentResponse, error) {
	var UpdateStudent models.Student
	UpdateStudent.Name = req.Name
	UpdateStudent.DateOfBirth = req.DateOfBirth

	// Tìm kiếm học sinh theo ID
	ObjectId, err := primitive.ObjectIDFromHex(req.Id) // Chuyển đổi từ string sang ObjectID
	if err != nil {
		log.Printf("Invalid student ID: %v", err)
		return &pb.UpdateStudentResponse{Message: "Invalid student ID"}, err
	}

	result, err := database.GetStudentCollection().UpdateOne(context.TODO(), bson.M{"_id": ObjectId}, bson.M{"$set": UpdateStudent})
	if err != nil {
		log.Printf("Error updating student: %v", err)
		return &pb.UpdateStudentResponse{Message: "Update that bai"}, err
	}

	if result.MatchedCount == 0 {
		return &pb.UpdateStudentResponse{Message: "Khong tim thay id hoc sinh"}, nil
	}

	return &pb.UpdateStudentResponse{Message: "Update thanh cong"}, nil
}

func (service *StudentService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentResponse, error) {
	ObjectId, err := primitive.ObjectIDFromHex(req.Id) // Chuyển đổi từ string sang ObjectID
	if err != nil {
		log.Printf("Invalid student ID: %v", err)
		return &pb.DeleteStudentResponse{Message: "Invalid student ID"}, err
	}
	result, err := database.GetStudentCollection().DeleteOne(context.TODO(), bson.M{"_id": ObjectId})
	if err != nil {
		log.Printf("Error deleting student: %v", err)
		return &pb.DeleteStudentResponse{Message: "Failed to delete student"}, err
	}

	if result.DeletedCount == 0 {
		return &pb.DeleteStudentResponse{Message: "Student not found"}, nil
	}

	return &pb.DeleteStudentResponse{Message: "Student deleted successfully"}, nil
}
