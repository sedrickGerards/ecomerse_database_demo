package api

import (
	"net/http"
	"os"


	campay "github.com/Iknite-Space/sqlc-example-api/campay"
	"github.com/Iknite-Space/sqlc-example-api/db/repo"
	"github.com/gin-gonic/gin"
)

type ShopHandler struct {
	querier repo.Querier
}

func NewOrdersHandler(querier repo.Querier) *ShopHandler {
	return &ShopHandler{
		querier: querier,
	}
}

func (h *ShopHandler) WireHttpHandler() http.Handler {

	r := gin.Default()
	r.Use(gin.CustomRecovery(func(c *gin.Context, _ any) {
		c.String(http.StatusInternalServerError, "Internal Server Error: panic")
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	// r.POST("/message", h.handleCreateMessage)
	// r.GET("/message/:id", h.handleGetMessage)
	// r.POST("/thread", h.handleCreateThread)
	// r.POST("/thread/:id/messages", h.handleGetThreadMessages)
	// r.PATCH("/message/:id", h.handleEditMessage)
	//r.DELETE("/message/:id", h.handleDeleteMessage)

	r.POST("/customer", h.handleCreateCustomer)
	r.POST("/orders", h.handleCreateOrders)
	r.POST("/product", h.handleCreateProduct)
	r.GET("/product/:id", h.handleGetProductByID)

	return r
}

func (h *ShopHandler) handleCreateOrders(c *gin.Context) {
	var req repo.CreateOrdersParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.querier.GetCustomerByID(c, req.CustomerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orders, err := h.querier.CreateOrders(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	des := "making order payment"
	ref := orders.ID

	apikey := os.Getenv("API_KEY")
	// totalAmountStr := orders.TotalAmount.String()

	pay := campay.MakePayment(apikey, customer.Contact, orders.TotalAmount, des, ref)

	check := campay.CheckStatus(apikey, pay.Reference)

	// makePayment(amount, from, description, external_reference)

	c.JSON(http.StatusOK, check)
}

func (h *ShopHandler) handleCreateCustomer(c *gin.Context) {
	var req repo.CreateCustomerParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.querier.CreateCustomer(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *ShopHandler) handleCreateProduct(c *gin.Context) {
	var req repo.CreateProductParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.querier.CreateProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ShopHandler) handleGetProductByID(c *gin.Context) {

	//we set a string "id" that stores the string id that we will use to pass it to the header
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id not set"})
		return
	}

	// we set the id string params  here also since we passed the id string along the other arguments in the message.sql.go
	product, err := h.querier.GetProductByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// func (h *MessageHandler) handleGetMessage(c *gin.Context) {
// 	id := c.Param("id")
// 	if id == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
// 		return
// 	}

// 	message, err := h.querier.GetMessageByID(c, id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, message)
// }

// func (h *MessageHandler) handleGetThreadMessages(c *gin.Context) {
// 	id := c.Param("id")
// 	if id == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
// 		return
// 	}

// 	messages, err := h.querier.GetMessagesByThread(c, id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"thread":   id,
// 		"topic":    "example",
// 		"messages": messages,
// 	})
// }

// func (h *MessageHandler) handleEditMessage(c *gin.Context) {

// 	//we set a string "id" that stores the string id that we will use to pass it to the header
// 	id := c.Param("id")
// 	if id == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Id not set"})
// 		return
// 	}

// 	// creating a struct var that holds the editMessageNyIdParams that was created in the message.sql.go
// 	var req repo.EditMessageByIDParams
// 	err := c.ShouldBindBodyWithJSON(&req)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// we set the id string params  here also since we passed the id string along the other arguments in the message.sql.go
// 	message, err := h.querier.EditMessageByID(c, req)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, message)
// }

// func (h *MessageHandler) handleCreateThread(c *gin.Context) {
// 	var req repo.CreateThreadParams
// 	err := c.ShouldBindBodyWithJSON(&req)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	thread, err := h.querier.CreateThread(c, req)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, thread)
// }
