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

package docker

import (
	"deployment-manager/manager/itf"
	"github.com/docker/docker/api/types/mount"
)

var stateMap = map[string]itf.ContainerState{
	"created":    itf.UnknownState,
	"running":    itf.RunningState,
	"paused":     itf.UnknownState,
	"restarting": itf.UnknownState,
	"removing":   itf.UnknownState,
	"exited":     itf.StoppedState,
	"dead":       itf.UnhealthyState,
}

var restartPolicyMap = map[string]itf.RestartStrategy{
	"no":             itf.RestartNever,
	"on-failure":     itf.RestartOnFail,
	"always":         itf.RestartAlways,
	"unless-stopped": itf.RestartNotStopped,
}

var restartPolicyRMap = func() map[itf.RestartStrategy]string {
	m := make(map[itf.RestartStrategy]string)
	for k, v := range restartPolicyMap {
		m[v] = k
	}
	return m
}()

var mountTypeMap = map[mount.Type]itf.MountType{
	mount.TypeBind:   itf.BindMount,
	mount.TypeVolume: itf.VolumeMount,
	mount.TypeTmpfs:  itf.TmpfsMount,
}

var mountTypeRMap = func() map[itf.MountType]mount.Type {
	m := make(map[itf.MountType]mount.Type)
	for k, v := range mountTypeMap {
		m[v] = k
	}
	return m
}()

var portTypeMap = map[string]itf.PortType{
	"tcp":  itf.TcpPort,
	"udp":  itf.UdpPort,
	"sctp": itf.SctpPort,
}

var portTypeRMap = func() map[itf.PortType]string {
	m := make(map[itf.PortType]string)
	for k, v := range portTypeMap {
		m[v] = k
	}
	return m
}()
