package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Yui/app/ent"
	"github.com/Yui/app/ent/drug"
	"github.com/Yui/app/ent/registerstore"
	"github.com/Yui/app/ent/store"
	"github.com/Yui/app/ent/user"
	"github.com/gin-gonic/gin"
)

//RegisterstoreController defines the struct for the registerstore controller
type RegisterstoreController struct {
	client *ent.Client
	router gin.IRouter
}
type Registerstore struct {
	User   int
	Drug   int
	Store  int
	Amount int
}

// CreateRegisterstore handles POST requests for adding registerstore entities
// @Summary Create registerstore
// @Description Create registerstore
// @ID create-registerstore
// @Accept   json
// @Produce  json
// @Param registerstore body ent.Registerstore true "Registerstore entity"
// @Success 200 {object} ent.Register_store
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /registerstores [post]
func (ctl *RegisterstoreController) CreateRegisterstore(c *gin.Context) {
	obj := Registerstore{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "registerstore binding failed",
		})
		return
	}

	s, err := ctl.client.Store.
		Query().
		Where(store.IDEQ(int(obj.Store))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Store not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": "User not found",
		})
		return
	}

	d, err := ctl.client.Drug.
		Query().
		Where(drug.IDEQ(int(obj.Drug))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Drug not found",
		})
		return
	}
	
	r, err := ctl.client.Registerstore.
		Create().
		SetUser(u).
		SetDrug(d).
		SetStore(s).
		SetAmount(obj.Amount).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": r,
	})
}

// GetRegisterstore handles GET requests to retrieve a registerstore entity
// @Summary Get a registerstore entity by ID
// @Description get registerstore by ID
// @ID get-registerstore
// @Produce  json
// @Param id path int true "Registerstore ID"
// @Success 200 {object} ent.Registerstore
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /registerstores/{id} [get]
func (ctl *RegisterstoreController) GetRegisterstore(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	r, err := ctl.client.Registerstore.
		Query().
		Where(registerstore.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// ListRegisterstore handles request to get a list of registerstore entities
// @Summary List registerstore entities
// @Description list registerstore entities
// @ID list-registerstore
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Register_store
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /registerstores [get]
func (ctl *RegisterstoreController) ListRegisterstore(c *gin.Context) {
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

	registerstores, err := ctl.client.Registerstore.
		Query().
		WithDrug().
		WithStore().
		WithUser().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, registerstores)
}


// DeleteRegisterstore handles DELETE requests to delete a registerstore entity
// @Summary Delete a registerstore entity by ID
// @Description get registerstore by ID
// @ID delete-registerstore
// @Produce  json
// @Param id path int true "Registerstore ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /registerstores/{id} [delete]
func (ctl *RegisterstoreController) DeleteRegisterstore(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Registerstore.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
 }
  
 

// NewRegisterstoreController creates and registers handles for the registerstore controller
func NewRegisterstoreController(router gin.IRouter, client *ent.Client) *RegisterstoreController {
	rc := &RegisterstoreController{
		client: client,
		router: router,
	}
	rc.register()
	return rc
}

// InitRegisterstoreController registers routes to the main engine
func (ctl *RegisterstoreController) register() {
	registerstores := ctl.router.Group("/registerstores")

	// CRUD
	registerstores.POST("", ctl.CreateRegisterstore)
	registerstores.GET("", ctl.ListRegisterstore)
	registerstores.DELETE(":id", ctl.DeleteRegisterstore)
}
