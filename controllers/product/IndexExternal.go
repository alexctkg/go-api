package product

import (
	"strconv"
	"tdez/database.go"
	"tdez/models"
	"tdez/resources"

	"github.com/gin-gonic/gin"
)

func IndexExternal(c *gin.Context) {

	var products []models.ResProduct

	db, err := database.SetupDB()
	if err != nil {
		c.JSON(400, gin.H{"errors": []string{"Error on database connection"}})
		c.Abort()
		return
	}
	defer db.Close()
	tx := db.Begin()

	useCode := c.MustGet("use_code").(int) //externalapp code
	filterStatus := c.DefaultQuery("status", "")

	var status int
	if filterStatus != "" {
		status, err = strconv.Atoi(c.DefaultQuery("status", ""))
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"errors": []string{"Filtro inexistente"}})
			c.Abort()
			return
		}
	}

	tx = tx.Where("pro_date_del is null")

	//pending
	if status == 0 {
		tx = tx.Where("pro_status = 0")
	} else if status == 1 { //accepted
		tx = tx.Where("pro_status = 1")
	} else if status == 2 { //rejected
		tx = tx.Where("pro_status = 2")
	}

	err = tx.
		Where("use_code_ext = ?", useCode).
		Find(&products).Error
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