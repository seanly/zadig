/*
Copyright 2022 The KodeRover Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/koderover/zadig/v2/pkg/microservice/aslan/core/stat/service"
	internalhandler "github.com/koderover/zadig/v2/pkg/shared/handler"
	e "github.com/koderover/zadig/v2/pkg/tool/errors"
)

type getStatReq struct {
	StartDate    int64    `json:"startDate,omitempty" form:"startDate,default=0"`
	EndDate      int64    `json:"endDate,omitempty"   form:"endDate,default=0"`
	ProductNames []string `json:"productNames"`
}

func GetAllPipelineTask(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	ctx.RespErr = service.GetAllPipelineTask(ctx.Logger)
}

func GetBuildDailyAverageMeasure(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()
	//params validate

	args := new(getStatReq)
	if err := c.BindJSON(args); err != nil {
		ctx.RespErr = e.ErrInvalidParam.AddErr(err)
		return
	}

	ctx.Resp, ctx.RespErr = service.GetBuildDailyAverageMeasure(args.StartDate, args.EndDate, args.ProductNames, ctx.Logger)
}

func GetBuildDailyMeasure(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()
	//params validate
	args := new(getStatReq)
	if err := c.BindJSON(args); err != nil {
		ctx.RespErr = e.ErrInvalidParam.AddErr(err)
		return
	}

	ctx.Resp, ctx.RespErr = service.GetBuildDailyMeasure(args.StartDate, args.EndDate, args.ProductNames, ctx.Logger)
}

func GetBuildHealthMeasure(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()
	//params validate
	args := new(getStatReq)
	if err := c.BindJSON(args); err != nil {
		ctx.RespErr = e.ErrInvalidParam.AddErr(err)
		return
	}

	ctx.Resp, ctx.RespErr = service.GetBuildHealthMeasure(args.StartDate, args.EndDate, args.ProductNames, ctx.Logger)
}

func GetLatestTenBuildMeasure(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()
	//params validate
	// TODO: ?????????????????
	args := new(getStatReq)
	if err := c.BindJSON(args); err != nil {
		ctx.RespErr = e.ErrInvalidParam.AddErr(err)
		return
	}

	// TODO:                                                    ↓ only this is used
	ctx.Resp, ctx.RespErr = service.GetLatestTenBuildMeasure(args.ProductNames, ctx.Logger)
}

func GetTenDurationMeasure(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()
	//params validate
	args := new(getStatReq)
	if err := c.BindJSON(args); err != nil {
		ctx.RespErr = e.ErrInvalidParam.AddErr(err)
		return
	}

	ctx.Resp, ctx.RespErr = service.GetTenDurationMeasure(args.StartDate, args.EndDate, args.ProductNames, ctx.Logger)
}

func GetBuildTrendMeasure(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()
	//params validate
	args := new(getStatReq)
	if err := c.BindJSON(args); err != nil {
		ctx.RespErr = e.ErrInvalidParam.AddErr(err)
		return
	}

	ctx.Resp, ctx.RespErr = service.GetBuildTrendMeasure(args.StartDate, args.EndDate, args.ProductNames, ctx.Logger)
}
