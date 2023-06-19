package routes

import (
	"github.com/brnocorreia/go-movies-crud/src/controller"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.GET("/getUserByNickname/:userNickname", model.VerifyTokenMiddleware, userController.FindUserByNickname)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
	r.POST("/login/", userController.LoginUser)

}

func InitBetRoutes(r *gin.RouterGroup, betController controller.BetControllerInterface) {
	r.POST("/createBet", model.VerifyTokenMiddleware, betController.CreateBet)
	r.GET("/getBetById/:betId", model.VerifyTokenMiddleware, betController.FindBetByID)
}
