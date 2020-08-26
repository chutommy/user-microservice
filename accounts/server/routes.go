package server

import (
	"github.com/chutified/appointments/accounts/controllers"
	"github.com/gin-gonic/gin"
)

// SetRoutes returns the gin.Engine with the routings set up.
func SetRoutes(h *controllers.Handler) *gin.Engine {

	// initializa the router
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	api := r.Group("/i")
	{
		api.GET("/auth", h.LoginAccount)

		api.GET("/accounts/len", h.GetAccountsLen)
		api.GET("/accounts", h.GetAccountsAll)

		api.POST("/account", h.AddAccount)
		api.GET("/account/:id", h.GetAccountByID)
		api.GET("/account/params", h.GetAccountByParams)
		api.PUT("/account/:id", h.EditAccountByID)
		api.DELETE("/account/:id", h.DeleteAccountByID)
	}

	return r
}
