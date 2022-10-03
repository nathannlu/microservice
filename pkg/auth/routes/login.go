package login

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathannlu/microservice/pkg/auth/pb"
)


type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
/*
type AuthServiceClient interface {
	Register() (*RegisterResponse, error)
	Login(
		ctx context.Context,
		in *LoginRequest,
		opts ...grpc.CallOption
	) (*LoginResponse, error)
	Validate() (*ValidateResponse, error)
}
*/
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	err := ctx.BindJSON();
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Run protobuf function
	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email: body.Email,
		Password: body.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	
	// Return
	ctx.JSON(int(res.Status), &res)
}

