package controllers

import (
	"context"
	"strconv"

	"github.com/Yui/app/ent"
	"github.com/Yui/app/ent/store"
	"github.com/gin-gonic/gin"
)

//StoreController defines the struct for the store controller
type StoreController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateStore handles POST requests for adding store entities
// @Summary Create store
// @Description Create store
// @ID create-store
// @Accept   json
// @Produce  json
// @Param Store body ent.Store true "Store entity"
// @Success 200 {object} ent.Store
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /stores [post]
func (ctl *StoreController) CreateStore(c *gin.Context) {
	obj := ent.Store{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Store binding failed",
		})
		return
	}

	s, err := ctl.client.Store.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, s)
}

// GetStore handles GET requests to retrieve a store entity
// @Summary Get a store entity by ID
// @Description get store by ID
// @ID get-store
// @Produce  json
// @Param id path int true "Store ID"
// @Success 200 {object} ent.Store
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /stores/{id} [get]
func (ctl *StoreController) GetStore(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.Store.
		Query().
		Where(store.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListStore handles request to get a list of store entities
// @Summary List store entities
// @Description list store entities
// @ID list-store
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Store
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /stores [get]
func (ctl *StoreController) ListStore(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	stores, err := ctl.client.Store.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, stores)
}

// NewStoreController creates and store handles for the store controller
func NewStoreController(router gin.IRouter, client *ent.Client) *StoreController {
	sc := &StoreController{
		client: client,
		router: router,
	}
	sc.store()
	return sc
}

// InitStoreController stores routes to the main engine
func (ctl *StoreController) store() {
	stores := ctl.router.Group("/stores")

	// CRUD
	stores.POST("", ctl.CreateStore)
	stores.GET(":id", ctl.GetStore)
	stores.GET("", ctl.ListStore)
}
