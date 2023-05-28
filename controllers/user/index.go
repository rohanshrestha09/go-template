package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/rohanshrestha09/go-template/dto"
	"github.com/rohanshrestha09/go-template/enums"
	"github.com/rohanshrestha09/go-template/models"
	"github.com/rohanshrestha09/go-template/service/database"
	"github.com/rohanshrestha09/go-template/utils"
	"golang.org/x/crypto/bcrypt"
)

// Regsiter godoc
//
//	@Summary	Create an account
//	@Tags		User
//	@Accept		mpfd
//	@Produce	json
//	@Param		name			formData	string	true	"Name"
//	@Param		email			formData	string	true	"Email"
//	@Param		password		formData	string	true	"Password"			minlength(8)
//	@Param		confirmPassword	formData	string	true	"Confirm Password"	minlength(8)
//	@Param		image			formData	file	false	"File to upload"
//	@Success	201				{object}	user.Register.Response
//	@Router		/user/register [post]
func Register(ctx *gin.Context) {

	var registerDto dto.RegisterDTO

	if err := ctx.ShouldBindWith(&registerDto, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if err := validator.New().Struct(registerDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := models.User{
		Name:     registerDto.Name,
		Email:    registerDto.Email,
		Password: registerDto.Password,
	}

	recordExists, err := database.RecordExists(models.User{Email: user.Email})

	if recordExists && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if recordExists {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "User already exists"})
		return
	}

	if file, err := ctx.FormFile("image"); err == nil {
		user.Image, user.ImageName, err = utils.UploadFile(file, enums.USER_DIR, enums.IMAGE)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Something went wrong"})
	}

	user.Password = string(hash)

	if _, err := database.Create(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	authToken, err := utils.SignJwt(user.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	type Response struct {
		Message string `json:"message"`
		Data    struct {
			Token string `json:"token"`
		} `json:"data"`
	}

	ctx.JSON(http.StatusCreated, Response{"Register Successful", struct {
		Token string `json:"token"`
	}{
		Token: authToken,
	}})

}

// Login godoc
//
//	@Summary	Login User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		body	body		dto.LoginDTO	true	"Request Body"
//	@Success	201		{object}	user.Login.Response
//	@Router		/user/login [post]
func Login(ctx *gin.Context) {

	var loginDto dto.LoginDTO

	if err := ctx.BindJSON(&loginDto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user, err := database.Get(database.GetArgs[models.User]{
		Filter: models.User{Email: loginDto.Email},
	})

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect Password"})
		return
	}

	authToken, err := utils.SignJwt(user.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	type Response struct {
		Message string `json:"message"`
		Data    struct {
			Token string `json:"token"`
		} `json:"data"`
	}

	ctx.JSON(http.StatusCreated, Response{"Register Successful", struct {
		Token string `json:"token"`
	}{
		Token: authToken,
	}})

}

// Get User godoc
//
//	@Summary	Get a user
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"User ID"
//	@Success	200	{object}	database.GetResponse[models.User]
//	@Router		/user/{id} [get]
func Get(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	ctx.JSON(
		http.StatusOK,
		database.GetResponse[models.User]{
			Message: "User Fetched",
			Data:    user,
		})
}
