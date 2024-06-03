package transaction

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kingkong-be/common"
	"kingkong-be/config/validator"
	"kingkong-be/delivery/http/transaction/model"
	"kingkong-be/domain/transaction"
	"net/http"
	"strconv"
)

func (c *controller) Add(ctx *gin.Context) {
	bodyRequest := new(model.Transaction)
	if err := ctx.BindJSON(bodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error()))
		return
	}

	if err := validator.ValidateStruct(bodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BadRequestResponse(err))
		return
	}

	p := new(transaction.Transaction)
	mapRequestAddTransaction(bodyRequest, p)

	var parts []transaction.TransactionPart
	for k, _ := range bodyRequest.TransactionParts {
		tmpParts := new(transaction.TransactionPart)
		mapRequestAddTransactionPart(&bodyRequest.TransactionParts[k], tmpParts)
		parts = append(parts, *tmpParts)
	}

	req := &transaction.RequestInsertTransaction{
		Transaction:      *p,
		TransactionParts: parts,
	}

	if err := c.transactionService.AddTransaction(req); err != nil {
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

	datas, counts, err := c.transactionService.GetList(query.Limit, query.Offset, query.Type)
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

	data, err := c.transactionService.Get(convID)
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

	bodyRequest := new(model.Transaction)
	if err := ctx.BindJSON(bodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error()))
		return
	}

	if err := validator.ValidateStruct(bodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BadRequestResponse(err))
		return
	}

	p := new(transaction.Transaction)
	mapRequestAddTransaction(bodyRequest, p)

	var parts []transaction.TransactionPart
	for k, _ := range bodyRequest.TransactionParts {
		tmpParts := new(transaction.TransactionPart)
		mapRequestAddTransactionPart(&bodyRequest.TransactionParts[k], tmpParts)
		parts = append(parts, *tmpParts)
	}

	if err := c.transactionService.Update(convID, p); err != nil && errors.Is(err, common.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, common.ErrorResponse(err.Error()))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse(err.Error()))
		return
	}

	if err := c.transactionService.UpdateBatchPart(convID, parts); err != nil && errors.Is(err, common.ErrRecordNotFound) {
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

	if err := c.transactionService.Delete(convID); err != nil && errors.Is(err, common.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, common.ErrorResponse(err.Error()))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.SuccessResponseNoData("success"))
	return
}

func (c *controller) GetChart(ctx *gin.Context) {
	var resp transaction.ResponseChart
	salesWeeklyChart, err := c.transactionService.GetSumSales7DaysBefore()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	purchaseWeeklyChart, err := c.transactionService.GetSumPurchase7DaysBefore()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	salesMonthlyChart, err := c.transactionService.GetSumSalesMonthly()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	purchaseMonthlyChart, err := c.transactionService.GetSumPurchaseMonthly()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	resp.WeeklyChartSales = salesWeeklyChart
	resp.WeeklyChartPurchase = purchaseWeeklyChart
	resp.MonthlyChartSales = salesMonthlyChart
	resp.MonthlyChartPurchase = purchaseMonthlyChart
	ctx.JSON(http.StatusOK, common.SuccessResponseWithData(resp, "success"))

}
