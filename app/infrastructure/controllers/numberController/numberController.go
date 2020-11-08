package numberController

import (
	"github.com/gin-gonic/gin"
	"github.com/scyanh/numtoen/app/application/numberService"
	model "github.com/scyanh/numtoen/app/domain/models"
	"net/http"
)

type GetTranslationNumberSwagger struct {
	model.Response
}

// Translate number to english godoc
// @Summary Translator response
// @Tags Translator
// @Accept  json
// @Produce  json
// @Param number query int true "number"
// @Success 200 {object} numberController.GetTranslationNumberSwagger
// @Router /num_to_english [get]
func TranslateToEnglish(c *gin.Context) {
	responseContextData := model.ResponseContext{Ctx: c}
	numberString := c.Query("number")

	translation, err := numberService.TranslateNumber(numberString)

	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(model.StatusSuccess, translation))
}
