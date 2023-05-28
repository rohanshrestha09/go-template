package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rohanshrestha09/go-template/enums"
	"github.com/rohanshrestha09/go-template/service/database"

	"github.com/rohanshrestha09/go-template/dto"
	"github.com/rohanshrestha09/go-template/models"
	"github.com/rohanshrestha09/go-template/utils"
)

// Auth godoc
//
//	@Summary	Get auth profile
//	@Tags		Auth
//	@Produce	json
//	@Success	200		{object}	database.GetResponse[models.User]
//	@Router		/auth 	[get]
//	@Security	Bearer
func Get(ctx *gin.Context) {
	authUser := ctx.MustGet("auth").(models.User)

	ctx.JSON(
		http.StatusOK,
		database.GetResponse[models.User]{
			Message: "Authorised",
			Data:    authUser,
		})

}

// Update Profile godoc
//
//	@Summary	Update profile
//	@Tags		Auth
//	@Accept		mpfd
//	@Produce	json
//	@Param		name	formData	string	false	"Name"
//	@Param		bio		formData	string	false	"Description"
//	@Param		image	formData	file	false	"File to upload"
//	@Success	201		{object}	database.Response
//	@Router		/auth 	[patch]
//	@Security	Bearer
func Update(ctx *gin.Context) {

	authUser := ctx.MustGet("auth").(models.User)

	var profileUpdateDto dto.ProfileUpdateDTO

	if err := ctx.ShouldBindWith(&profileUpdateDto, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	var imageUrl, imageName string

	if file, err := ctx.FormFile("image"); err == nil {
		imageUrl, imageName, err = utils.UploadFile(file, enums.USER_DIR, enums.IMAGE)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		if err := utils.DeleteFile(string(enums.USER_DIR) + authUser.ImageName); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	data := models.User{
		Name:      profileUpdateDto.Name,
		Bio:       profileUpdateDto.Bio,
		Image:     imageUrl,
		ImageName: imageName,
	}

	response, err := database.Update(authUser, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Project Updated"))

}
