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

package util

import (
	"deployment-manager/manager/itf"
	"fmt"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func GenEnv(ev map[string]string) (env []string) {
	if len(ev) > 0 {
		for key, val := range ev {
			env = append(env, fmt.Sprintf("%s=%s", key, val))
		}
	}
	return
}

func GenStopTimeout(d *itf.Duration) *int {
	if d != nil {
		t := int(d.Seconds())
		return &t
	}
	return nil
}

func GenPortMap(ports []itf.Port) (nat.PortMap, error) {
	pm := make(nat.PortMap)
	set := make(map[string]struct{})
	for _, p := range ports {
		if _, ok := set[p.KeyStr()]; ok {
			return pm, fmt.Errorf("port duplicate '%s'", p.KeyStr())
		}
		set[p.KeyStr()] = struct{}{}
		port, err := nat.NewPort(PortTypeRMap[p.Protocol], strconv.FormatInt(int64(p.Number), 10))
		if err != nil {
			return pm, err
		}
		var bindings []nat.PortBinding
		for _, binding := range p.Bindings {
			bindings = append(bindings, nat.PortBinding{
				HostIP:   binding.Interface.String(),
				HostPort: strconv.FormatInt(int64(binding.Number), 10),
			})
		}
		pm[port] = bindings
	}
	return pm, nil
}

func GenMounts(mounts []itf.Mount) ([]mount.Mount, error) {
	var msl []mount.Mount
	set := make(map[string]struct{})
	for i := 0; i < len(mounts); i++ {
		m := mounts[i]
		if _, ok := set[m.KeyStr()]; ok {
			return msl, fmt.Errorf("mount duplicate '%s'", m.KeyStr())
		}
		set[m.KeyStr()] = struct{}{}
		mnt := mount.Mount{
			Type:     MountTypeRMap[m.Type],
			Source:   m.Source,
			Target:   m.Target,
			ReadOnly: m.ReadOnly,
		}
		switch m.Type {
		case itf.VolumeMount:
			mnt.VolumeOptions = &mount.VolumeOptions{Labels: m.Labels}
		case itf.TmpfsMount:
			mnt.TmpfsOptions = &mount.TmpfsOptions{
				SizeBytes: m.Size,
				Mode:      m.Mode,
			}
		}
		msl = append(msl, mnt)
	}
	return msl, nil
}

func GenFilterArgs(filter [][2]string) (f filters.Args) {
	if filter != nil && len(filter) > 0 {
		f = filters.NewArgs()
		for _, i := range filter {
			f.Add(i[0], i[1])
		}
	}
	return
}

func GenNetIPAMConfig(n itf.Network) (c []network.IPAMConfig) {
	c = append(c, network.IPAMConfig{
		Subnet:  n.Subnet.KeyStr(),
		Gateway: n.Gateway.String(),
	})
	return
}

func GenTimestamp(t time.Time) string {
	tp := strings.Split(t.Format(time.RFC3339Nano), ":")
	s := strings.TrimSuffix(tp[2], "Z")
	var ns string
	if strings.Contains(s, ".") {
		sp := strings.Split(s, ".")
		s = sp[0]
		ns = sp[1]
	}
	nsLen := utf8.RuneCountInString(ns)
	if nsLen < 9 {
		ns += strings.Repeat("0", 9-nsLen)
	}
	return fmt.Sprintf("%s:%s:%s.%sZ", tp[0], tp[1], s, ns)
}

func CheckNetworks(n []itf.ContainerNet) error {
	set := make(map[string]struct{})
	for _, net := range n {
		if _, ok := set[net.Name]; ok {
			return fmt.Errorf("network duplicate '%s'", net.Name)
		}
		set[net.Name] = struct{}{}
	}
	return nil
}
