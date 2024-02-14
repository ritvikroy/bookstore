package controllers

import (
	mocks "bookstore-api/mocks"
	"bookstore-api/model"
	"bookstore-api/service"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	//	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/suite"
)

type BooksControllerTestSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	service    service.BookStoreService
	recorder   *httptest.ResponseRecorder
	booksMockRepo  *mocks.MockBookStoreRepository
	context    *gin.Context
	goContext  context.Context
	controller BooksController
}

func TestBooksControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BooksControllerTestSuite))
}

func (suite *BooksControllerTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	// suite.db, _ = mocks
	suite.booksMockRepo = mocks.NewMockBookStoreRepository(suite.ctrl)
	suite.goContext = context.TODO()
	suite.service = service.NewBooksService(suite.booksMockRepo)
	suite.controller = NewGetAllBooks(suite.service)
}
func (suite *BooksControllerTestSuite) TearDowTest() {
	suite.ctrl.Finish()
}

func (suite *BooksControllerTestSuite) TestViewAllEmptyBooks() {
	expected := model.AllBooks{
		Books: []model.Books{},
	}
	exp, _ := json.Marshal(expected)
	suite.context.Request = httptest.NewRequest(http.MethodGet, "/api/books", nil)
	suite.controller.GetAllBooks(suite.context)

	suite.Equal(string(exp), suite.recorder.Body.String())
}

func (suite *BooksControllerTestSuite) TestViewAllBooks() {
	expected := model.AllBooks{
		Books: []model.Books{
			{
				Name:        "Book1",
				Price:       "100",
				Description: "new book1",
			},
		},
	}
	exp, _ := json.Marshal(expected)
	suite.context.Request = httptest.NewRequest(http.MethodGet, "/api/books", nil)
	suite.booksMockRepo.EXPECT().GetAllBooks(suite.context,"",0,0).Return(expected.Books,nil).Times(1)
	suite.controller.GetAllBooks(suite.context)

	suite.Equal(string(exp), suite.recorder.Body.String())
}

func (suite *BooksControllerTestSuite) TestViewAllBooksWithSearch() {
	expected := model.AllBooks{
		Books: []model.Books{
			{
				Name:        "Book1",
				Price:       "100",
				Description: "new book1",
			},
		},
	}
	exp, _ := json.Marshal(expected)
	suite.context.Request = httptest.NewRequest(http.MethodGet, "/api/books?searchText=Book", nil)
	suite.booksMockRepo.EXPECT().GetAllBooks(suite.context,"Book",0,0).Return(expected.Books,nil).Times(1)
	suite.controller.GetAllBooks(suite.context)

	suite.Equal(string(exp), suite.recorder.Body.String())
}

func (suite *BooksControllerTestSuite) TestViewAllBooksWithSearchWhereNoResult() {
	suite.context.Request = httptest.NewRequest(http.MethodGet, "/api/books?searchText=Test", nil)
	suite.booksMockRepo.EXPECT().GetAllBooks(suite.context,"Test",0,0).Return([]model.Books{},errors.New("some error occured")).Times(1)
	suite.controller.GetAllBooks(suite.context)
	suite.Equal(http.StatusBadRequest,suite.recorder.Result().StatusCode)
}