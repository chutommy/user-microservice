package controllers

import (
	"github.com/chutified/appointments/accounts/models"
	"github.com/gin-gonic/gin"
)

// AddAccount adds a new account to database.
func (h *Handler) AddAccount(c *gin.Context) {

	// get the account
	var a models.Account
	err := c.BindJSON(&a)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// add to database
	id, err := h.ds.AddAccount(c, &a)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success (" + id + ")",
	})
}

func (h *Handler) GetAccountsAll(c *gin.Context) {
}
func (h *Handler) GetAccountByID(c *gin.Context) {
}
func (h *Handler) GetAccountByParams(c *gin.Context) {
}
func (h *Handler) LoginAccount(c *gin.Context) {
}
func (h *Handler) EditAccountByID(c *gin.Context) {
}
func (h *Handler) DeleteAccountByID(c *gin.Context) {
}
