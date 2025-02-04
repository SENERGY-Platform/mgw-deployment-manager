/*
 * Copyright 2025 InfAI (CC SES)
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

package wrapper

import (
	"context"
	"github.com/SENERGY-Platform/mgw-container-engine-wrapper/lib/model"
	"io"
)

type ContainerEngineHandler interface {
	ListNetworks(ctx context.Context) ([]model.Network, error)
	ListContainers(ctx context.Context, filter model.ContainerFilter) ([]model.Container, error)
	ListImages(ctx context.Context, filter model.ImageFilter) ([]model.Image, error)
	ListVolumes(ctx context.Context, filter model.VolumeFilter) ([]model.Volume, error)
	NetworkInfo(ctx context.Context, id string) (model.Network, error)
	NetworkCreate(ctx context.Context, net model.Network) (string, error)
	NetworkRemove(ctx context.Context, id string) error
	ContainerInfo(ctx context.Context, id string) (model.Container, error)
	ContainerCreate(ctx context.Context, container model.Container) (id string, err error)
	ContainerRemove(ctx context.Context, id string, force bool) error
	ContainerStart(ctx context.Context, id string) error
	ContainerStop(ctx context.Context, id string) error
	ContainerRestart(ctx context.Context, id string) error
	ContainerLog(ctx context.Context, id string, logOptions model.LogFilter) (io.ReadCloser, error)
	ContainerExec(ctx context.Context, id string, execOpt model.ExecConfig) error
	ImageInfo(ctx context.Context, id string) (model.Image, error)
	ImagePull(ctx context.Context, id string) error
	ImageRemove(ctx context.Context, id string) error
	VolumeInfo(ctx context.Context, id string) (model.Volume, error)
	VolumeCreate(ctx context.Context, vol model.Volume) (string, error)
	VolumeRemove(ctx context.Context, id string, force bool) error
}
