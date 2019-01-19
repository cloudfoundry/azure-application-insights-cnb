/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package java

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cloudfoundry/libcfbuildpack/build"
	"github.com/cloudfoundry/libcfbuildpack/buildpack"
	"github.com/cloudfoundry/libcfbuildpack/helper"
	"github.com/cloudfoundry/libcfbuildpack/layers"
)

// Properties represents the azure-application-insights-properties helper application.
type Properties struct {
	buildpack buildpack.Buildpack
	layer     layers.Layer
}

// Contributes makes the contribution to launch
func (p Properties) Contribute() error {
	return p.layer.Contribute(marker{p.buildpack.Info}, func(layer layers.Layer) error {
		if err := os.RemoveAll(layer.Root); err != nil {
			return err
		}

		if err := helper.CopyFile(filepath.Join(p.buildpack.Root, "bin", "azure-application-insights-properties"), filepath.Join(layer.Root, "bin", "azure-application-insights-properties")); err != nil {
			return err
		}

		return layer.WriteProfile("azure-application-insights-properties", `printf "Configuring Azure Application Insights Properties"

export JAVA_OPTS="${JAVA_OPTS} $(azure-application-insights-properties)"
`)
	}, layers.Launch)
}

// String makes Properties satisfy the Stringer interface.
func (p Properties) String() string {
	return fmt.Sprintf("Properties{ buildpack: %s layer: %s }", p.buildpack, p.layer)
}

// NewProperties creates a new Properties instance.
func NewProperties(build build.Build) Properties {
	return Properties{build.Buildpack, build.Layers.Layer("azure-application-insights-properties")}
}

type marker struct {
	buildpack.Info
}

func (m marker) Identity() (string, string) {
	return "Azure Application Insights Properties", m.Version
}

