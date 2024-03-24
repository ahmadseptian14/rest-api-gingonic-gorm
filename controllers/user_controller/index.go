package usercontroller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"gin-gonic-gorm/requests"
	"gin-gonic-gorm/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var users = []models.User{}

func RegisterUser(ctx *gin.Context) {
	var user models.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        ctx.JSON(500, gin.H{"error": "Failed to hash password"})
        return
    }

    // Update user's password with hashed password
    user.Password = string(hashedPassword)

    if err := controller.DB.Create(&user).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(200, gin.H{"message": "User registered successfully"})
}

func GetAllUser(ctx *gin.Context)  {
	users := new([]models.User)
	err := database.DB.Table("users").Find(&users).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"Message": "Internal server error"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success get all user",
		"data" : users,
	})
}

func GetById(ctx *gin.Context)  {
	id := ctx.Param("id")
	user := new(responses.UserResponse)

	err := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if err != nil{
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success get user",
		"data" : user,
	})
}

func Store(ctx *gin.Context)  {
	userReq := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil{
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	userEmailExists := new(models.User)
	database.DB.Table("users").Where("email = ?", userReq.Email).First(&userEmailExists)

	if userEmailExists.Email != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already used",
		})
		return
	}

	user := new(models.User)

	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	err := database.DB.Table("users").Create(&user).Error

	if err != nil{
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
	}

	ctx.JSON(200, gin.H{
		"message": "Success create user",
		"data": user,
	})
}

func UpdateById(ctx *gin.Context)  {
	id := ctx.Param("id")
	user := new(models.User)
	userEmailExists := new(models.User)
	userReq := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "data not found",
		})
		return
	}

	errUserEmailExists := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExists).Error
	if errUserEmailExists != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	if userEmailExists.Email != nil && *user.ID != *userEmailExists.ID {
		ctx.JSON(400, gin.H{
			"message": "Email already used",
		})
		return
	}

	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		ctx.JSON(500, gin.H{
			"message": "cant update data",
		})
		return
	}

	userResponse := responses.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
	}

	ctx.JSON(200, gin.H{
		"message": "Data updated successfully",
		"data": userResponse,
	})
}

func DeleteById(ctx *gin.Context)  {
	id := ctx.Param("id")
	user := new(models.User)

	errFind := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if errFind != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "Data not found",
		})
		return
	}

	errDb := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&user).Error

	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Data deleted successfully",
	})
}

func GetUserPaginate(ctx *gin.Context) {
	page := ctx.Query("page")
	if page == "" {
		page = "1"
	}

	perPage := ctx.Query("perPage")
	if perPage == "" {
		perPage = "10"
	}

	perPageInt, _ := strconv.Atoi(perPage)
	pageInt, _ := strconv.Atoi(page)

	if pageInt < 1 {
		pageInt = 1
	}

	users := new([]models.User)
	err := database.DB.Table("users").Offset((pageInt-1) * perPageInt).Limit(perPageInt).Find(&users).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"Message": "Internal server error"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success get user paginate",
		"data" : users,
		"page": pageInt,
		"per_page": perPageInt,
	})
}
