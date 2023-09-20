package foods

import (
	"foods-spider-server/com"
	"foods-spider-server/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IFoodsApi interface {
	api.IBase
	GetCommonFoods() func(ctx *gin.Context)
	SearchFood() func(ctx *gin.Context)
}

type FoodApi struct {
}

func (f FoodApi) SetPrefix() string {
	return "/api/v1/foods"
}

// GetCommonFoods 获取常见的食物类型列表
func (f FoodApi) GetCommonFoods() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		cf, err := GetCommonFoodsLogic()
		if err != nil {
			com.ResponseData(ctx, http.StatusOK, cf, "获取食物类型失败")
			return
		}
		com.ResponseData(ctx, com.CodeSuccess, cf)
	}
}

// SearchFood 查询食物
func (f FoodApi) SearchFood() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		foodName := ctx.Query("foodName")
		foodSearchRes, err := SearchFoodLogic(foodName)
		if err != nil {
			com.ResponseData(ctx, http.StatusOK, foodSearchRes, "搜索食物失败")
			return
		}
		com.ResponseData(ctx, com.CodeSuccess, foodSearchRes)
	}

}
func NewFoodApi() IFoodsApi {
	return &FoodApi{}
}
