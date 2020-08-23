package product

import (
	"strconv"
	"tdez/database.go"
	"tdez/models"
	"tdez/resources"

	"github.com/gin-gonic/gin"
)

func ProductIndex(c *gin.Context) {

	var products []models.ResProduct

	db, err := database.SetupDB()
	if err != nil {
		c.JSON(400, gin.H{"errors": []string{"Error on database connection"}})
		c.Abort()
		return
	}
	defer db.Close()
	tx := db.Begin()

	filterStatus, err := strconv.Atoi(c.DefaultQuery("status", ""))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{"Filtro inexistente"}})
		c.Abort()
		return
	}

	tx = tx.Where("pro_date_del is null")

	//pending
	if filterStatus == 0 {
		tx = tx.Where("pro_status == 0")
	} else if filterStatus == 1 { //accepted
		tx = tx.Where("pro_status == 1")
	} else if filterStatus == 2 { //rejected
		tx = tx.Where("pro_status == 2")
	}

	err = tx.Find(&products).Error
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}

	var productsResource []resources.ResProduct

	for _, worker := range products {
		var productResource resources.ResProduct
		productResource.ResProductResource(worker)
		productsResource = append(productsResource, productResource)
	}

	tx.Commit()
	c.JSON(200, gin.H{"data": productsResource})

	c.Abort()
	return

}
