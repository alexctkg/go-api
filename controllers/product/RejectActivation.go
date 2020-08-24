package product

import (
	"tdez/database"
	"tdez/models"
	"tdez/requests"
	"tdez/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// RejectActivation godoc
// @Tags Product
// @Summary Update product - rejected
// @Description
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Request body requests.ResProductResponse true "Request body"
// @Success 200 {object} models.DefaultSuccess
// @Failure 400 {object} models.DefaultError
// @Router /admin/reject [put]
func RejectActivation(c *gin.Context) {
	var request requests.ResProductResponse

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}

	if err := utils.Valid(request); err != nil {
		c.JSON(400, gin.H{"errors": err})
		return
	}

	useCode := c.MustGet("use_code").(int) //superuser code

	db, err := database.SetupDB()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}
	tx := db.Begin()

	var product models.ResProduct
	//select product to reject
	queryProduct := tx.
		Where("pro_date_del is null").
		Where("pro_code = ?", request.Code).
		First(&product)
	if queryProduct.RowsAffected == 0 {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{"Não foi possível encontrar esse produto"}})
		c.Abort()
		return
	}

	timeNow := time.Now()
	product.CodeIntUse = &useCode
	product.DateUpdt = &timeNow
	product.Status = 2 //reject
	product.StatusReason = request.Reason

	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return

	}

	//SEND EMAIL TO COSTUMER...

	tx.Commit()
	c.JSON(200, gin.H{"messages": []string{"O produto foi rejeitado com sucesso"}})
	c.Abort()
	return

}
