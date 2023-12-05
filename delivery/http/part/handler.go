package part

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kingkong-be/common"
	"kingkong-be/config/validator"
	"kingkong-be/delivery/http/part/model"
	"kingkong-be/domain/part"
	"net/http"
	"strconv"
)

func (c *controller) Add(ctx *gin.Context) {
	bodyRequest := new(model.Part)
	if err := ctx.BindJSON(bodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error()))
		return
	}

	if err := validator.ValidateStruct(bodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BadRequestResponse(err))
		return
	}

	p := new(part.Part)
	mapRequestAddPart(bodyRequest, p)

	if err := c.partService.AddPart(p); err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, common.SuccessResponseWithData(bodyRequest, "success"))
	return
}

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

	datas, counts, err := c.partService.GetList(query.Limit, query.Offset)
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

	data, err := c.partService.Get(convID)
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

func (c *controller) Update(ctx *gin.Context) {
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

	bodyRequest := new(model.Part)
	if err := ctx.BindJSON(bodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error()))
		return
	}

	if err := validator.ValidateStruct(bodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BadRequestResponse(err))
		return
	}

	p := new(part.Part)
	mapRequestAddPart(bodyRequest, p)

	if err := c.partService.Update(convID, p); err != nil && errors.Is(err, common.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, common.ErrorResponse(err.Error()))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.SuccessResponseWithData(bodyRequest, "success"))
	return
}

func (c *controller) Delete(ctx *gin.Context) {
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

	if err := c.partService.Delete(convID); err != nil && errors.Is(err, common.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, common.ErrorResponse(err.Error()))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.SuccessResponseNoData("success"))
	return
}
