package methods

import (
	"context"
	"strconv"

	"github.com/ajalck/service_1/pkg/pb"
	"github.com/gin-gonic/gin"
)

type Client struct {
	C pb.UsersClient
}

func (c *Client) ClientMethod(ctx *gin.Context) {
	methodNo, _ := strconv.Atoi(ctx.Query("method"))
	waitTime, _ := strconv.Atoi(ctx.Query("waitTime"))

	if methodNo == 1 {
		res, err := c.C.ListUsers(context.Background(), &pb.RequestParams{
			Method:   int32(methodNo),
			WaitTime: int32(waitTime),
		})
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		ctx.JSON(200, res)
	} else if methodNo == 2 {
	} else {
		ctx.JSON(400, "Method not found")
		return
	}

}
