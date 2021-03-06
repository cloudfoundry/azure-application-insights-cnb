/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package java

import (
	"path/filepath"

	"github.com/cloudfoundry/libcfbuildpack/v2/build"
	"github.com/cloudfoundry/libcfbuildpack/v2/buildpack"
	"github.com/cloudfoundry/libcfbuildpack/v2/helper"
	"github.com/cloudfoundry/libcfbuildpack/v2/layers"
)

// Properties represents the azure-application-insights-properties helper application.
type Properties struct {
	buildpack buildpack.Buildpack
	layer     layers.HelperLayer
}

// Contributes makes the contribution to launch
func (p Properties) Contribute() error {
	return p.layer.Contribute(func(artifact string, layer layers.HelperLayer) error {
		if err := helper.CopyFile(artifact, filepath.Join(layer.Root, "bin", "azure-application-insights-properties")); err != nil {
			return err
		}

		return layer.WriteProfile("azure-application-insights-properties", `printf "Configuring Azure Application Insights Properties"

export JAVA_OPTS="${JAVA_OPTS} $(azure-application-insights-properties)"
`)
	}, layers.Launch)
}

// NewProperties creates a new Properties instance.
func NewProperties(build build.Build) Properties {
	return Properties{build.Buildpack, build.Layers.HelperLayer("azure-application-insights-properties", "Azure Application Insights Properties")}
}
