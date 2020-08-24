package product

import (
	"tdez/database"
	"tdez/models"
	"tdez/requests"
	"tdez/utils"

	"github.com/gin-gonic/gin"

	cnpjValid "github.com/Nhanderu/brdoc"
)

// IssueActivation godoc
// @Tags product
// @Summary Store a product issue
// @Description create a new request
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Request body requests.ResProductResponse true "Request body"
// @Success 200 {object} models.DefaultSuccess
// @Failure 400 {object} models.DefaultError
// @Router /reject [put]
func IssueActivation(c *gin.Context) {
	var request requests.ResProduct

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}

	if err := utils.Valid(request); err != nil {
		c.JSON(400, gin.H{"errors": err})
		return
	}

	if request.CostumerType == 0 {
		if cnpjValid.IsCPF(request.CostumerCpfCnpj) {
			c.JSON(200, gin.H{"messages": []string{"CNPJ/CPF inválido"}})
			c.Abort()
		}
	} else {
		if cnpjValid.IsCNPJ(request.CostumerCpfCnpj) != true {
			c.JSON(200, gin.H{"messages": []string{"CNPJ/CPF inválido"}})
			c.Abort()
		}
	}

	useCode := c.MustGet("use_code").(int) //externalapp user code

	db, err := database.SetupDB()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}
	tx := db.Begin()

	var product models.ResProduct

	product.ResProductFill(request)
	product.CodeExtUse = useCode //externalapp user code

	//select open issues
	var productDB models.ResProduct
	query := tx.
		Where("pro_costumermid_cpf_cnpj = ?", product.CostumerCpfCnpj).
		Where("pro_date_updt is null").
		Where("pro_date_del is not null").
		Where("use_code_ext = ?", product.CodeExtUse).
		Find(&productDB)
	if query.RowsAffected > 0 {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{"Já existe um produto pendente, aguarde até a resolução"}})
		c.Abort()
		return
	}

	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}

	tx.Commit()

	c.JSON(200, gin.H{"messages": []string{"O produto foi salvo com sucesso"}})
	c.Abort()
	return

}
