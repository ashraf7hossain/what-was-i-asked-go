package user

import (
	"net/http"
	"rest-in-go/initializers"
	"rest-in-go/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input InputRegisterUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	// Validate fields
	if errs := validateInput(input); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors":  errs,
			"message": "Validation failed",
		})
		// c.Error(errs).SetMeta(http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	input.Password = string(hash)
	user, err := h.service.CreateUser(input)
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var input InputLoginUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByEmail(input.Email)

	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.Error(err).SetMeta(http.StatusUnauthorized)
		return
	}

	token, err := initializers.GenerateJWT(user.ID)
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	type processedUser struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	finalUser := processedUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": finalUser})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		c.Error(utils.NewError("Unauthorized")).SetMeta(http.StatusUnauthorized)
		return
	}

	user, err := h.service.GetProfile(userID)
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	type processedUser struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	finalUser := processedUser{Name: user.Name, Email: user.Email}

	c.JSON(http.StatusOK, gin.H{"user": finalUser})
}
