package users

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService *service.UserService) UserController {
	return UserController{UserService: *UserService}
}

// Me 			 godoc
// @Summary      Get current user profile
// @Description  get profile data from JWT token
// @Tags         me
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Insert your JWT token"
// @Success      200  {object}  model.User
// @Failure      401  {object}  model.Error
// @Failure      500  {object}  model.Error
// @Router       /me/ [get]
func (controller UserController) Me(c *gin.Context) {
	authorization := c.GetHeader("Authorization")

	userID, err := helper.GetUserID(authorization, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.UserService.FindUserByID").
			Int64("user_id", userID).
			Msg("Failed to fetch user request")

		c.JSON(http.StatusUnauthorized, model.Error{Error: "invalid_token", Message: "Unauthorized"})
		return
	}

	user, err := controller.UserService.GetUserByID(&entity.User{UserId: userID})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.UserService.FindUserByID").
			Int64("user_id", userID).
			Msg("Failed to fetch user request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "find_user_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Int64("user_id", userID).
		Msg("Fetch user request processed successfully")

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

// FetchUserByID godoc
// @Summary      Get user profile by userID
// @Description  get profile data from userID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Insert your admin JWT token"
// @Success      200  {object}  model.User
// @Failure      400  {object}  model.Error
// @Failure      401  {object}  model.Error
// @Failure      500  {object}  model.Error
// @Router       /users/:id [get]
func (controller UserController) FetchUserByID(c *gin.Context) {
	param := c.Param("id")

	userID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "strconv.ParseInt").
			Str("user_id", param).
			Msg("Failed to fetch user request")

		c.JSON(http.StatusBadRequest, model.Error{Error: "invalid_user_id", Message: "BadRequest"})
		return
	}

	user, err := controller.UserService.GetUserByID(&entity.User{UserId: userID})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.UserService.FindUserByID").
			Int64("user_id", userID).
			Msg("Failed to fetch user request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "find_user_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "200").
		Int64("user_id", userID).
		Msg("Fetch user request processed successfully")

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

// DiscardUserByID godoc
// @Summary        Discard user profile by userID
// @Description    delete user data from userID
// @Tags           users
// @Accept         json
// @Produce        json
// @Param          Authorization  header  string  true  "Insert your admin JWT token"
// @Success        204
// @Failure        400  {object}  model.Error
// @Failure        500  {object}  model.Error
// @Router         /users/:id [delete]
func (controller UserController) DiscardUserByID(c *gin.Context) {
	param := c.Param("id")

	userID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "strconv.ParseInt").
			Str("user_id", param).
			Msg("Failed to discard user request")

		c.JSON(http.StatusBadRequest, model.Error{Error: "invalid_user_id", Message: "BadRequest"})
		return
	}

	_, err = controller.UserService.RemoveUserByID(&entity.User{UserId: userID})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.UserService.RemoveUserByID").
			Int64("user_id", userID).
			Msg("Failed to discard user request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "discard_user_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Int64("user_id", userID).
		Msg("Discard user request processed successfully")

	c.Status(http.StatusNoContent)
}

// DiscardUserByID godoc
// @Summary        Edit user profile by userID
// @Description    delete user data from userID
// @Tags           users
// @Accept         json
// @Produce        json
// @Param          Authorization  header  string  true  "Insert your admin JWT token"
// @Success        200  {object}  model.UserResponse
// @Failure        400  {object}  model.Error
// @Failure        500  {object}  model.Error
// @Router         /users/:id [put]
func (controller UserController) EditUserByID(c *gin.Context) {
	param := c.Param("id")

	userID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "strconv.ParseInt").
			Str("user_id", param).
			Msg("Failed to discard user request")

		c.JSON(http.StatusBadRequest, model.Error{Error: "invalid_user_id", Message: "BadRequest"})
		return
	}

	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil || userID == 0 {
		c.JSON(http.StatusBadRequest, model.Error{Error: "invalid_request", Message: "StatusBadRequest"})
		return
	}

	user, err := controller.UserService.ModifyUserByID(&entity.User{
		UserId:      userID,
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Username:    req.Username,
		Balance:     req.Balance,
		RecoveryKey: req.RecoveryKey,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "controller.UserService.ModifyUserByID").
			Int64("user_id", req.UserId).
			Msg("Failed to discard user request")

		c.JSON(http.StatusInternalServerError, model.Error{Error: "edit_user_failed", Message: "InternalServerError"})
		return
	}

	log.Info().
		Str("status_code", "204").
		Int64("user_id", req.UserId).
		Msg("Edit user request processed successfully")

	c.JSON(http.StatusOK, model.UserResponse{Message: "success", Data: *user})
}
