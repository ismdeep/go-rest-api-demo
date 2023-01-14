package api

import (
	"github.com/gin-gonic/gin"

	"github.com/ismdeep/go-rest-api-demo/app/server/handler"
	"github.com/ismdeep/go-rest-api-demo/internal/request"
)

// SignUp sign up
// @Summary sign up
// @Author l.jiang.1024@gmail.com
// @Description sign up
// @Tags Auth
// @Param Authorization	header string true "Bearer 31a165ba1be6dec616b1f8f3207b4273"
// @Param req body	request. true "JSON数据"
// @Success 200 {object} response.
// @Router /api/v1/auth/sign-up [post]
func SignUp(c *gin.Context) {
	var req request.SignUp
	if err := c.ShouldBindJSON(&req); err != nil {
		failed(c, err)
		return
	}

	if err := handler.Auth.SignUp(c, req); err != nil {
		failed(c, err)
		return
	}

	success(c, "", nil)
}

// SignIn sign in
// @Summary sign in
// @Author l.jiang.1024@gmail.com
// @Description sign in
// @Tags Auth
// @Param Authorization	header string true "Bearer 31a165ba1be6dec616b1f8f3207b4273"
// @Param req body	request. true "JSON数据"
// @Success 200 {object} response.
// @Router /api/v1/auth/sign-in [post]
func SignIn(c *gin.Context) {

}

// MyProfile my profile
// @Summary my profile
// @Author l.jiang.1024@gmail.com
// @Description my profile
// @Tags Auth
// @Param Authorization	header string true "Bearer 31a165ba1be6dec616b1f8f3207b4273"
// @Param req body	request. true "JSON数据"
// @Success 200 {object} response.
// @Router /api/v1/auth/profile [get]
func MyProfile(c *gin.Context) {

}
