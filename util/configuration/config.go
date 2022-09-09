/*
 * Copyright 2022 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package configuration

import (
	"deployment-manager/manager/api/engine"
	"deployment-manager/util/logger"
	"encoding/json"
	envldr "github.com/y-du/go-env-loader"
	"os"
)

type Config struct {
	SocketPath string        `json:"socket_path" env_var:"SOCKET_PATH"`
	Logger     logger.Config `json:"logger" env_var:"LOGGER_CONFIG"`
	ApiEngine  engine.Config `json:"api_engine" env_var:"API_ENGINE_CONFIG"`
}

func NewConfig(path *string) (cfg *Config, err error) {
	cfg = &Config{
		SocketPath: "/opt/deployment-manager/manager.sock",
		Logger: logger.Config{
			Level:  logger.WarningLvl,
			Utc:    true,
			Prefix: "[DM] ",
		},
	}
	if path != nil {
		var file *os.File
		if file, err = os.Open(*path); err != nil {
			return
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(cfg); err != nil {
			return
		}
	}
	err = envldr.LoadEnvUserParser(cfg, typeParsers, nil)
	return cfg, err
}
