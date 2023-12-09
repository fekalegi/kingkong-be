package price_changes_log

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kingkong-be/common"
	"kingkong-be/config/validator"
	"kingkong-be/delivery/http/price_changes_log/model"
	"kingkong-be/domain/price_changes_log"
	"net/http"
	"strconv"
)

func (c *controller) GetList(ctx *gin.Context) {

	query := new(model.List)
	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error()))
		return
	}

	if err := validator.ValidateStruct(query); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BadRequestResponse(err))
		return
	}

	datas, counts, err := c.pcLogService.GetList(&price_changes_log.GetListRequest{
		Limit:     query.Limit,
		Offset:    query.Offset,
		Search:    query.Search,
		StartDate: query.StartDate,
		EndDate:   query.EndDate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.SuccessPaginationResponse(
		datas,
		"success",
		common.ResponseMeta{
			Limit:  query.Limit,
			Offset: query.Offset,
			Total:  int(counts),
		}))
	return
}

func (c *controller) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	convID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse("id is invalid"))
		return
	}

	if convID <= 0 {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse("id has to be above 0"))
		return
	}

	data, err := c.pcLogService.Get(convID)
	if err != nil && errors.Is(err, common.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, common.ErrorResponse(err.Error()))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.SuccessResponseWithData(data, "success"))
	return
}
