package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"QLToolsV2/internal/model"
	"QLToolsV2/internal/service"
	res "QLToolsV2/pkg/response"
	val "QLToolsV2/utils/validator"
)

type PanelController struct{}

// Router 注册路由
func (c *PanelController) Router(r *gin.RouterGroup) {
	// 分页查询
	r.GET("/panel/list", c.List)
	// 全部查询
	r.GET("/panel/all/list", c.AllList)

	// 添加
	r.POST("/panel/add", c.Add)
	// 批量操作[启用/禁用]
	r.PUT("/panel/batch/operation", c.BatchOperation)
	// 修改
	r.PUT("/panel/update", c.Update)
	// 删除
	r.DELETE("/panel/delete", c.Delete)

	// 刷新Token
	r.POST("/panel/refresh", c.Refresh)
}

// List 分页查询
func (c *PanelController) List(ctx *gin.Context) {
	// 获取参数
	p := new(model.Pagination)
	if err := ctx.ShouldBindQuery(&p); err != nil {
		// 判断err是不是validator.ValidationErrors类型
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			res.ResError(ctx, res.CodeInvalidParam)
			return
		}

		// 翻译错误
		res.ResErrorWithMsg(ctx, res.CodeInvalidParam, val.RemoveTopStruct(errs.Translate(val.Trans)))
		return
	}

	// 业务处理
	resCode, msg := service.PanelList(p)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // 成功
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // 失败
	}
}

// AllList 全部查询
func (c *PanelController) AllList(ctx *gin.Context) {
	// 业务处理
	resCode, msg := service.PanelAllList()
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // 成功
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // 失败
	}
}

// Add 添加
func (c *PanelController) Add(ctx *gin.Context) {
	// 获取参数
	p := new(model.AddPanel)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 判断err是不是validator.ValidationErrors类型
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			res.ResError(ctx, res.CodeInvalidParam)
			return
		}

		// 翻译错误
		res.ResErrorWithMsg(ctx, res.CodeInvalidParam, val.RemoveTopStruct(errs.Translate(val.Trans)))
		return
	}

	// 业务处理
	resCode, msg := service.AddPanel(p)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // 成功
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // 失败
	}
}

// BatchOperation 批量操作[启用/禁用]
func (c *PanelController) BatchOperation(ctx *gin.Context) {
	// 获取参数
	p := new(model.BatchOperationPanel)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 判断err是不是validator.ValidationErrors类型
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			res.ResError(ctx, res.CodeInvalidParam)
			return
		}

		// 翻译错误
		res.ResErrorWithMsg(ctx, res.CodeInvalidParam, val.RemoveTopStruct(errs.Translate(val.Trans)))
		return
	}

	// 业务处理
	resCode, msg := service.BatchOperationPanel(p)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // 成功
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // 失败
	}
}

// Update 修改
func (c *PanelController) Update(ctx *gin.Context) {
	// 获取参数
	p := new(model.UpdatePanel)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 判断err是不是validator.ValidationErrors类型
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			res.ResError(ctx, res.CodeInvalidParam)
			return
		}

		// 翻译错误
		res.ResErrorWithMsg(ctx, res.CodeInvalidParam, val.RemoveTopStruct(errs.Translate(val.Trans)))
		return
	}

	// 业务处理
	resCode, msg := service.UpdatePanel(p)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // 成功
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // 失败
	}
}

// Delete 删除
func (c *PanelController) Delete(ctx *gin.Context) {
	// 获取参数
	p := new(model.DeletePanel)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		// 判断err是不是validator.ValidationErrors类型
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			res.ResError(ctx, res.CodeInvalidParam)
			return
		}

		// 翻译错误
		res.ResErrorWithMsg(ctx, res.CodeInvalidParam, val.RemoveTopStruct(errs.Translate(val.Trans)))
		return
	}

	// 业务处理
	resCode, msg := service.DeletePanel(p)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // 成功
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // 失败
	}
}

// Refresh 刷新
func (c *PanelController) Refresh(ctx *gin.Context) {
	// 业务处理
	resCode, msg := service.RefreshPanel()
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // 成功
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // 失败
	}
}
