/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package usecase

import (
	"unicare/internal/config"
	"unicare/internal/server/repo"
)

type UseCase struct {
	app  *config.App
	repo *repo.Repo
}

func New(app *config.App, repo2 *repo.Repo) *UseCase {
	return &UseCase{app, repo2}
}
