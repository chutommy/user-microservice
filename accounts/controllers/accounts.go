package controllers

import (
	"strconv"

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

// GetAccountsLen returns the number of accounts in the database.
func (h *Handler) GetAccountsLen(c *gin.Context) {

	// get number of the pages
	pages, err := h.ds.AccountsPages(c, 1)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"pages": pages,
	})
}

// GetAccountsAll resturns all accounts based on the pigination.
func (h *Handler) GetAccountsAll(c *gin.Context) {

	// get query values
	pageNum, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil || pageNum < 0 {
		c.JSON(400, gin.H{
			"error": "page value must be non-negative integer",
		})
		return
	}

	pageCap, err := strconv.Atoi(c.DefaultQuery("cap", "1"))
	if err != nil || pageCap < 1 {
		c.JSON(400, gin.H{
			"error": "the page cap must be positive integer",
		})
		return
	}

	sort := c.DefaultQuery("sort", "id")
	switch sort {
	case "id", "username", "email", "phone", "hpassword", "first_name", "last_name", "birth_day", "perm_address", "mail_address", "created_at", "updated_at":
	default:
		c.JSON(400, gin.H{
			"errors": "invalid sort by query value",
		})
		return
	}

	asc := c.DefaultQuery("asc", "true") == "true"

	// get the page
	accs, err := h.ds.GetAllAccounts(c, pageCap, pageNum, sort, asc)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, accs)
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
