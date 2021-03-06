/*
Copyright 2017 Kinvolk GmbH

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

package main

import (
	"fmt"
	"os"

	"github.com/kinvolk/kube-spawn/pkg/cnispawn"
)

func main() {
	var background bool

	switch os.Args[1] {
	case "-i":
		background = false
	case "-d":
		background = true
	default:
		fmt.Println("Usage: cnispawn [-i|-d] <nspawn arguments>\n\n-i\tinteractive; container will be spawned with tty attached\n-d\tdetached; container will be run in the background")
		os.Exit(0)
	}

	if err := cnispawn.Spawn(background, os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
