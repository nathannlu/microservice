package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathannlu/microservice/pkg/auth/pb"
)

type RegisterRequestBody struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}
	err := ctx.BindJSON();
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// run protobuf function
	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email: body.Email,
		Password: body.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	// return
	ctx.JSON(int(res.Status), &res)
}
