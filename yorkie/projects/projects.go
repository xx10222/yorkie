/*
 * Copyright 2022 The Yorkie Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package projects

import (
	"context"

	"github.com/yorkie-team/yorkie/api/types"
	"github.com/yorkie-team/yorkie/yorkie/backend"
	"github.com/yorkie-team/yorkie/yorkie/backend/db"
)

// FindProjectByPublicKey finds the project by public key.
func FindProjectByPublicKey(
	ctx context.Context,
	be *backend.Backend,
	publicKey string,
) (*types.Project, error) {
	info, err := be.DB.FindProjectInfoByPublicKey(ctx, publicKey)
	if err != nil {
		return nil, err
	}
	return info.ToProject(), nil
}

// CreateProject creates a project.
func CreateProject(
	ctx context.Context,
	be *backend.Backend,
	name string,
) (*types.Project, error) {
	info, err := be.DB.CreateProjectInfo(ctx, name)
	if err != nil {
		return nil, err
	}

	return info.ToProject(), nil
}

// UpdateProject updates a project.
func UpdateProject(
	ctx context.Context,
	be *backend.Backend,
	project *types.Project,
) error {
	// TODO(hackerwins): If updates are executed concurrently, only one remains
	// and the rest may be deleted. Consider to update the project with CAS or update
	// the fields in the project separately.
	return be.DB.UpdateProjectInfo(ctx, db.ToProjectInfo(project))
}