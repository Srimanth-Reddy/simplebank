package api

import (
	"database/sql"
	"errors"
	"github.com/Srimanth-Reddy/simplebank/db/sqlc"
	"github.com/Srimanth-Reddy/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user sqlc.User) userResponse {
	return userResponse{
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := sqlc.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Name:     req.Name,
		Email:    req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case pgerrcode.UniqueViolation:
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := newUserResponse(user)
	ctx.JSON(http.StatusOK, resp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := loginUserResponse{
		User:        newUserResponse(user),
		AccessToken: accessToken,
	}

	ctx.JSON(http.StatusOK, resp)
}
