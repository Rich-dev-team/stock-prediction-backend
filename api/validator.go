package api

import "github.com/go-playground/validator/v10"

var validateStockPrices validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if prices, ok := fieldlevel.Field().Interface().([]int32); ok {
		for i := 0; i < len(prices); i++ {
			if prices[i] <= 0 {
				return false
			}
		}
		return true
	}
	return false
}

var validateCompanysId validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if ids, ok := fieldlevel.Field().Interface().([]int64); ok {
		for i := 0; i < len(ids); i++ {
			if ids[i] <= 0 {
				return false
			}
		}
		return true
	}
	return false
}
