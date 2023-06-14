package handler

import (
	"Project_Title/app/features/users"
	"Project_Title/config"
	"Project_Title/helper"
	"Project_Title/middleware"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	s users.Service 
	c *config.Configuration
}

func New(h users.Service, c *config.Configuration) users.Handler {
		return &UserController{s : h, c: c}
} 

func (uc *UserController) Register() echo.HandlerFunc {
		return func(c echo.Context) error {
				input := RegisterInput{} 
				if err := c.Bind(&input); err != nil {
					c.Logger().Error(err.Error())
					return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
				}
				err := uc.s.Register(users.Core{Name: input.Name, Email: input.Email, Country: input.Country, Password: input.Password})
				if err != nil {
						return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
				}
				return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Success Created an Account", nil))
		}
}

func (uc *UserController) Login() echo.HandlerFunc {
		return func(c echo.Context) error {
				var input LoginInput
				if err := c.Bind(&input); err != nil {
						c.Logger().Error(err.Error())
						return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
				} 

				user, err := uc.s.Login(input.Email, input.Password)
				if err != nil {
						c.Logger().Error(err.Error())
						return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
				}

				token, err := middleware.CreateJWT(user.ID,user.Email, user.Name)
				if err != nil {
						c.Logger().Error(err.Error())
						return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "internal Server Error", nil ))
				} 

				return c.JSON(http.StatusOK, helper.DataResponse{
						Code: http.StatusOK,
						Message: "Successful login, please use this token for further access.",
						Data: map[string]interface{}{"token": token},
				})
		}
} 

func (uc *UserController) GetProfile() echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			claims, err := middleware.ValidateJWT2(tokenString)
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT."+err.Error(), nil))
			}

			id := claims.ID 
			data, err := uc.s.GetProfile(uint(id))
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resourse was not found. please check your email and password input.",nil))
			}
			res := UserResponse{}
			copier.Copy(&res, &data)
			return c.JSON(http.StatusOK, helper.DataResponse{
					Code	: 		http.StatusOK,
					Message	: 		"Successful opertion.",
					Data	: 		res,
			})
		}
}

func (uc *UserController) UpdateProfile() echo.HandlerFunc {
		return func(c echo.Context) error {
			var input UpdateInput 
			tokenString := c.Request().Header.Get("Authorization") 
			claims, err := middleware.ValidateJWT2(tokenString)
			if err != nil {
					c.Logger().Error(err.Error())
					return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT."+err.Error(), nil))
			} 

			id := claims.ID
			if err := c.Bind(&input); err != nil {
					c.Logger().Error(err.Error())
					return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest,"Bad Request", nil))
			} 

			// file, err := c.FormFile("image")
			// if err != nil {
			// 	c.Logger().Error(err.Error())
			// 	return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
			// } 

			// image, err := helper.UploadToS3(c, filename)
			// if err != nil {
			// 		c.Logger().Error(err.Error())
			// 		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server Error", nil))
			// }

			updatedUser := users.Core{
				ID: uint(id), 
				Name: input.Name,
				Email: input.Email,
				Country: input.Country,
				Password: input.Password,
			}

			err = uc.s.UpdateProfile(uint(id), updatedUser)
			if err != nil {
					c.Logger().Error(err.Error())
					return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
			}

			return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Success Update an Account", nil))
		}
} 

func (uc *UserController) DeleteProfile() echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			claims, err := middleware.ValidateJWT2(tokenString)
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
			} 

			id := claims.ID 
			err = uc.s.DeleteProfile(uint(id))
			if err != nil {
					c.Logger().Error("Error deleting profile", err.Error())
					return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
			}
			return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success Deleted an Account", nil))
		}
}

