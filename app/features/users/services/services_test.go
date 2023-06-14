package services_test

import (
	"errors"
	"testing"

	
	 mocks "Project_Title/app/features/users/mocks"
	"Project_Title/app/features/users"
	"Project_Title/app/features/users/services"
	
	."github.com/onsi/ginkgo/v2"
	."github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail) 
	RunSpecs(t, "Services Suite")

} 

var _ = Describe("user", func(){
	var Mock *mocks.Repository 
	var UserService users.Service 
	var payload = users.Core{
		ID: 1, 
		Name: "wanta", 
		Country: "Indonesia", 
		Email: "wanta12@gmail.com", 
		Password: "wanta123",
	}
	BeforeEach(func(){
		Mock = mocks.NewRepository(GinkgoT()) 
		UserService = services.New(Mock)
	})
	Context("User Register", func() {
		When("Jumlah inputan melebihi 72 karakter", func() {
			BeforeEach(func() {
				Mock.On("Register", mock.Anything).Return(users.Core{}, errors.New("Error When Hashing Password")).Once()
			}) 
			It("Akan Mengembalikan Erorr", func () {
				err := UserService.Register(payload) 
				Expect(err).ShouldNot(BeNil())
				Expect(err.Error()).To(Equal("Failed to register user"))
			})  
		}) 

		When("Email User Sudah Tredaftar", func(){
			BeforeEach(func() {
				Mock.On("Register", mock.Anything).Return(users.Core{}, errors.New("Email has been register")).Once()
			}) 
			It("Akan Mengembalikan Error", func ()  {
				err := UserService.Register(payload)
				Expect(err).ShouldNot(BeNil())
				Expect(err.Error()).To(Equal("Failed to register user"))
			})
		})

		When("Berhasil Mendaftarkan User", func ()  {
			BeforeEach(func ()  {
				Mock.On("Register", mock.Anything).Return(payload, nil).Once()
			}) 
			It("Akan Mengembalikan Error", func ()  {
				err := UserService.Register(payload)
				Expect(err).Should(BeNil())
			})
		})
	}) 
	Context("User Login", func() {
		When("Email Belum Terdaftar", func ()  {
			BeforeEach(func()  {
				Mock.On("Login", payload.Email, payload.Password).Return(users.Core{}, errors.New("Email Not Registered")).Once()
			}) 
			It("Akan Mengembalikan Error", func ()  {
				_, err := UserService.Login(payload.Email, payload.Password)
				Expect(err).ShouldNot(BeNil()) 
				Expect(err.Error()).To(Equal("Email Not Registered"))
			})
		}) 
		When("Salah Memasukan Password", func ()  {
			BeforeEach(func()  {
				Mock.On("Login", payload.Email, payload.Password).Return(users.Core{}, errors.New("Wrong Password")).Once()
			})
			It("Akan Mengembalikan Error", func ()  {
				_, err := UserService.Login(payload.Email, payload.Password)
				Expect(err).ShouldNot(BeNil())
				Expect(err.Error()).To(Equal("Wrong Password"))
			})
		})

		When("Berhasil Login", func() {
			BeforeEach(func ()  {
				Mock.On("Login", payload.Email, payload.Password).Return(payload, nil).Once()
			})
			It("Akan Mengembalikan Erorr", func () {
				data, err := UserService.Login(payload.Email, payload.Password)
				Expect(err).Should(BeNil())
				Expect(data.Name).To(Equal("wanta"))
			})
		})
	})
	Context("Profile", func () {
		When("Tidak terdapat data user", func () {
			BeforeEach(func() {
				Mock.On("GetProfile", payload.ID).Return(users.Core{},errors.New("Data Not Found")).Once()
			})
			It("Akan Mengembalikan Erorr", func() {
				_, err := UserService.GetProfile(payload.ID)
				Expect(err).ShouldNot(BeNil()) 
				Expect(err.Error()).To(Equal("Data Not Found"))
			})
		})
		When("Terdapat data user", func() {
			BeforeEach(func() {
				Mock.On("GetProfile", payload.ID).Return(payload, nil).Once()
			}) 
			It("Akan Mengembalikan Erorr", func () { 
				data, err := UserService.GetProfile(payload.ID)
				Expect(err).Should(BeNil())
				Expect(data.Name).To(Equal("wanta"))
			})
		})
    })
     
	Context("Update Profile", func() {
		When("Password diganti dan password baru melebihi 70 karakter", func() {
			It("Akan Mengembalikan Erorr", func() {
				err := UserService.UpdateProfile(payload.ID, users.Core{Password: "22222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222"})
				Expect(err).ShouldNot(BeNil())
			})
		})

		When("Id tidak ditemukan", func() {
			BeforeEach(func() {
				Mock.On("UpdateProfile", payload.ID, mock.Anything).Return(errors.New("Id not found")).Once()
			})
			It("Akan Mengembalikan Erorr", func() {
				err := UserService.UpdateProfile(payload.ID, payload)
				Expect(err).ShouldNot(BeNil())
			})
		})

		When("Berhasil memperbaharui data", func() {
			BeforeEach(func() {
				Mock.On("UpdateProfile", payload.ID, mock.Anything).Return(nil).Once()
			})
			It("Akan Mengembalikan Erorr", func() {
				err := UserService.UpdateProfile(payload.ID, payload)
				Expect(err).Should(BeNil())
			})
		})

	})

	Context("Delete Profile", func() {

		When("Id tidak ditemukan", func() {
			BeforeEach(func() {
				Mock.On("DeleteProfile", uint(10)).Return(gorm.ErrRecordNotFound).Once()
			})
			It("Akan Mengembalikan Erorr", func() {
				err := UserService.DeleteProfile(uint(10))
				Expect(err).ShouldNot(BeNil())
				Expect(err.Error()).To(Equal("Id: 10 ,not found"))
			})
		})
		When("Kesalahan Query Database", func() {
			BeforeEach(func() {
				Mock.On("DeleteProfile", payload.ID).Return(errors.New("Terjadi masalah pada server")).Once()
			})
			It("Akan Mengembalikan Erorr", func() {
				err := UserService.DeleteProfile(payload.ID)
				Expect(err).ShouldNot(BeNil())
				Expect(err.Error()).To(Equal("Terjadi masalah pada server"))
			})
		})

		When("Berhasil mengahapus profile", func() {
			BeforeEach(func() {
				Mock.On("DeleteProfile", payload.ID).Return(nil).Once()
			})
			It("Akan Mengembalikan Erorr", func() {
				err := UserService.DeleteProfile(payload.ID)
				Expect(err).Should(BeNil())
			})
		})

	})

})