// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package api

import (
	"fmt"
	"net/http"

	"github.com/go-vela/types"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/worker/router/middleware/executor"
)

// GetBuild represents the API handler to capture the
// build currently running on an executor.
func GetBuild(c *gin.Context) {
	e := executor.Retrieve(c)

	build, err := e.GetBuild()
	if err != nil {
		msg := fmt.Errorf("unable to read build: %w", err).Error()

		c.AbortWithStatusJSON(http.StatusInternalServerError, types.Error{Message: &msg})

		return
	}

	c.JSON(http.StatusOK, build)
}

// KillBuild represents the API handler to kill a
// build currently running on an executor.
//
// This function performs a hard cancellation of a build on worker.
// Any build running during this time will immediately be stopped.
func KillBuild(c *gin.Context) {
	e := executor.Retrieve(c)

	repo, err := e.GetRepo()
	if err != nil {
		msg := fmt.Errorf("unable to repo build: %w", err).Error()

		c.AbortWithStatusJSON(http.StatusInternalServerError, types.Error{Message: &msg})

		return
	}

	build, err := e.KillBuild()
	if err != nil {
		msg := fmt.Errorf("unable to kill build: %w", err).Error()

		c.AbortWithStatusJSON(http.StatusInternalServerError, types.Error{Message: &msg})

		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("killing build %s/%d", repo.GetFullName(), build.GetNumber()))
}
