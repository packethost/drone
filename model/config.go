// Copyright 2018 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

// ConfigStore persists pipeline configuration to storage.
type ConfigStore interface {
	ConfigLoad(int64) (*Config, error)
	ConfigFind(*Repo, string) (*Config, error)
	ConfigFindApproved(*Config) (bool, error)
	ConfigCreate(*Config) error
}

// Config represents a pipeline configuration.
type Config struct {
	ID     int64  `json:"-"    meddler:"config_id,pk"`
	RepoID int64  `json:"-"    meddler:"config_repo_id"`
	Data   string `json:"data" meddler:"config_data"`
	Hash   string `json:"hash" meddler:"config_hash"`
}