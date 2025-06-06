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

	"github.com/koderover/zadig/v2/pkg/microservice/aslan/core/stat/repository/models"
	"github.com/koderover/zadig/v2/pkg/microservice/aslan/core/stat/service"
	internalhandler "github.com/koderover/zadig/v2/pkg/shared/handler"
	e "github.com/koderover/zadig/v2/pkg/tool/errors"
)

type GetBuildStatArgs struct {
	StartDate int64 `json:"startDate"      form:"startDate,default=0"`
	EndDate   int64 `json:"endDate"        form:"endDate,default=0"`
}

func GetBuildStat(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	args := new(GetBuildStatArgs)
	if err := c.ShouldBindQuery(args); err != nil {
		ctx.RespErr = e.ErrInvalidParam.AddDesc(err.Error())
		return
	}

	ctx.Resp, ctx.RespErr = service.GetBuildTotalAndSuccess(&models.BuildStatOption{
		StartDate: args.StartDate,
		EndDate:   args.EndDate,
	}, ctx.Logger)
}

type OpenAPIGetBuildStatArgs struct {
	StartDate int64  `json:"startDate"      form:"startDate,default=0"`
	EndDate   int64  `json:"endDate"        form:"endDate,default=0"`
	Project   string `json:"project"        form:"projectKey"`
}

func GetBuildStatForOpenAPI(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	args := new(OpenAPIGetBuildStatArgs)
	if err := c.ShouldBindQuery(args); err != nil {
		ctx.RespErr = e.ErrInvalidParam.AddDesc(err.Error())
		return
	}

	queryArgs := &models.BuildStatOption{
		StartDate: args.StartDate,
		EndDate:   args.EndDate,
	}

	if args.Project != "" {
		queryArgs.ProductNames = []string{args.Project}
	}

	ctx.Resp, ctx.RespErr = service.GetBuildStats(queryArgs, ctx.Logger)
}
