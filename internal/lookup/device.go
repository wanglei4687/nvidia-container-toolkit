/**
# Copyright (c) 2021, NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package lookup

import (
	"fmt"
	"os"
)

const (
	devRoot = "/dev"
)

// NewCharDeviceLocator creates a Locator that can be used to find char devices at the specified root. A logger is
// also specified.
func NewCharDeviceLocator(opts ...Option) Locator {
	filter := assertCharDevice
	// TODO: We should have a better way to inject this logic than this envvar.
	if os.Getenv("__NVCT_TESTING_DEVICES_ARE_FILES") == "true" {
		filter = assertFile
	}

	opts = append(opts,
		WithSearchPaths("", devRoot),
		WithFilter(filter),
	)
	return NewFileLocator(
		opts...,
	)
}

// assertCharDevice checks whether the specified path is a char device and returns an error if this is not the case.
func assertCharDevice(filename string) error {
	info, err := os.Lstat(filename)
	if err != nil {
		return fmt.Errorf("error getting info: %v", err)
	}
	if info.Mode()&os.ModeCharDevice == 0 {
		return fmt.Errorf("%v is not a char device", filename)
	}
	return nil
}
