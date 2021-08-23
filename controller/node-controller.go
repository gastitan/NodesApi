package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gotest.com/nodes-api/model"
	"gotest.com/nodes-api/service"
)

type NodeController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Nearest(ctx *gin.Context)
}

// FindAll godoc
// @Summary List all existing nodes
// @Description List all existing nodes
// @Accept json
// @Produce json
// @Success 200 {array} model.Node
// @Router /nodes [get]
func FindAll(ctx *gin.Context) {
	nodes, err := service.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nodes)
}

// FindById godoc
// @Summary Get existing node by id
// @Description Get existing node by id
// @Accept json
// @Produce json
// @Param id path int true "Node ID"
// @Success 200 {object} model.Node
// @Router /nodes/{id} [get]
func FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	node, err := service.FindById(id)
	if node.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "NodeId not found"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, node)
}

// Update godoc
// @Summary Update existing node by id
// @Description Update existing node by id
// @Accept json
// @Produce json
// @Param id path int true "Node ID"
// @Success 200 {object} model.Node
// @Router /nodes/{id} [put]
func Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	nodeToUpdate, _ := service.FindById(id)
	if nodeToUpdate.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "NodeId not found"})
		return
	}

	err = ctx.BindJSON(&nodeToUpdate)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = validateRequest(nodeToUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updatedNode, err := service.Update(nodeToUpdate)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	log.Printf("Node with id %d succesfully updated", id)
	ctx.JSON(http.StatusOK, updatedNode)
}

// Save godoc
// @Summary Create a node
// @Description Create a node
// @Accept json
// @Produce json
// @Param node body model.Node true "Node to create"
// @Success 200 {object} model.Node
// @Router /nodes [post]
func Save(ctx *gin.Context) {
	var node model.Node
	err := ctx.BindJSON(&node)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = validateRequest(node)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	node, err = service.Save(node)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Node create with id %d", node.ID)
	ctx.JSON(http.StatusOK, node)
}

// Nearest godoc
// @Summary Get the nearest node from a param Location
// @Description Get the nearest node
// @Accept json
// @Produce json
// @Param lat query number true "Latitude"
// @Param lng query number true "Longitud"
// @Success 200 {object} model.Node
// @Router /nodes/nearest [get]
func Nearest(ctx *gin.Context) {

	lat, err := strconv.ParseFloat(ctx.Query("lat"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	lng, err := strconv.ParseFloat(ctx.Query("lng"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	nearest, err := service.Nearest(lat, lng)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, nearest)
}

func validateRequest(node model.Node) error {

	validate := validator.New()
	err := validate.Struct(node)
	if err != nil {
		log.Printf("Could not validate Node Struct: %v", err.Error())
		return err
	}

	if err := node.NodeType.IsValid(); err != nil {
		return err
	}
	return nil
}
