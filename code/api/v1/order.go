package v1

import (
	"app/config"
	"app/pkg/e"
	"app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetOrderList 获取订单列表
// @Summary 获取订单列表
// @Description Order 服务中提供的获取订单列表服务
// @Tags Order 服务
// @Security ApiKeyAuth
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Param user_id query int false "用户ID"
// @Param offset query int false "偏移量"
// @Param limit query int false "限制数量"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /orders [get]
func GetOrderList(ginCtx *gin.Context) {
	offset, _ := strconv.Atoi(ginCtx.DefaultQuery("offset", config.AppSetting.DefaultOffset))
	limit, _ := strconv.Atoi(ginCtx.DefaultQuery("limit", config.AppSetting.DefaultLimit))
	userID, _ := strconv.Atoi(ginCtx.DefaultQuery("user_id", "0"))

	count, data, code := service.GetOrderList(offset, limit, userID)
	if code != e.SUCCESS {
		ginCtx.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code)})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"data": data,
		"page_info": gin.H{
			"total":  count,
			"offset": offset,
			"limit":  limit,
		},
	})
}

// CreateOrder 创建订单
// @Summary 创建订单
// @Description Order 服务中提供的创建订单服务
// @Tags Order 服务
// @Security ApiKeyAuth
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Param order body schema.OrderCreateReq true "订单"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /orders [post]
func CreateOrder(ginCtx *gin.Context) {
	data, code := service.CreateOrder(ginCtx)
	if code != e.SUCCESS {
		ginCtx.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code)})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "data": data})
}
