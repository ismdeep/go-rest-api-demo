package api

import (
	"github.com/gin-gonic/gin"

	"github.com/ismdeep/go-rest-api-demo/app/server/store"
)

// GetUserList get user list
// @Summary get user list
// @Author jianglinwei@uniontech.com
// @Author l.jiang.1024@gmail.com
// @Description get user list
// @Tags User
// @Param Authorization	header string true "Bearer 31a165ba1be6dec616b1f8f3207b4273"
// @Param req body	request. true "JSON数据"
// @Success 200 {object} response.
// @Router /api/v1/users [get]
func GetUserList(c *gin.Context) {
	users, err := store.User.GetAll()
	if err != nil {
		failed(c, err)
		return
	}

	success(c, "", users)
}
