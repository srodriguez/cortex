/*
Copyright 2020 Cortex Labs, Inc.

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

package local

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/cortexlabs/cortex/pkg/lib/errors"
)

const (
	ErrConnectToDockerDaemon = "cli.connect_to_docker_daemon"
	ErrDockerPermissions     = "cli.docker_permissions"
)

func ErrorConnectToDockerDaemon() error {
	installMsg := "install it by following the instructions for your operating system: https://docs.docker.com/install"
	if strings.HasPrefix(runtime.GOOS, "darwin") {
		installMsg = "install it here: https://docs.docker.com/docker-for-mac/install"
	}

	return errors.WithStack(&errors.Error{
		Kind:    ErrConnectToDockerDaemon,
		Message: fmt.Sprintf("unable to connect to the Docker daemon\n\nplease confirm Docker is running, or if Docker is not installed, %s", installMsg),
	})
}

func ErrorDockerPermissions(err error) error {
	errStr := errors.Message(err)

	var groupAddStr string
	if strings.HasPrefix(runtime.GOOS, "linux") {
		groupAddStr = " (e.g. by running `sudo groupadd docker && sudo gpasswd -a $USER docker`)"
	}

	return errors.WithStack(&errors.Error{
		Kind:    ErrDockerPermissions,
		Message: errStr + "\n\nyou can re-run this command with `sudo`, or grant your current user access to docker" + groupAddStr,
	})
}
