package services_test

import (
	"Project_Title/app/features/books"
	"Project_Title/app/features/books/services"
	 mocks "Project_Title/app/features/books/mocks"
	"errors" 
	"testing"

	"github.com/stretchr/testify/mock" 
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services Suite")
}

var _ = Describe("book", func() {
	var Mock *mocks.Repository 
	var BooksService books.Service
	var payload = books.Core{
		ID: 1,
		Title: "Belajar Golang",
		PublishedYear: "2020",
		ISBN: "1232111122",
	}
	BeforeEach(func() {
		Mock = mocks.NewRepository(GinkgoT()) 
		BooksService = services.New(Mock)
	})
	Context("User Register", func() {
		When("Request Body Salah", func() {
			BeforeEach(func() {
				Mock.On("CreateBook", mock.Anything).Return(errors.New("Invalid Request Body")).Once()				
			})
			It("Akan Mengembalikan Error", func (){
				err := BooksService.CreateBook(payload, 2) 
				Expect(err).ShouldNot(BeNil())
				Expect(err.Error()).To(Equal("Invalid Request"))
			})
		}) 

		When("Berhasil menambahkan buku", func () {
			BeforeEach(func () {
				Mock.On("CreateBook", mock.Anything, mock.Anything).Return(nil).Once()				
			})
			It("Akan Mengembalikan Erorr", func() {
				err := BooksService.CreateBook(payload, 2)
				Expect(err).ShouldNot(BeNil())
			})
		})
	})
})