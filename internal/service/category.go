package service

import (
	"context"

	"github.com/claudiojsantos/grpc-go/internal/database"
	"github.com/claudiojsantos/grpc-go/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Nome, in.Descricao)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var categoryList []*pb.Category

	for _, category := range categories {
		categoryList = append(categoryList, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{
		Categories: categoryList,
	}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryId) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Find(in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}
