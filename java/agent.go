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
	"path/filepath"

	"github.com/cloudfoundry/libcfbuildpack/build"
	"github.com/cloudfoundry/libcfbuildpack/buildpack"
	"github.com/cloudfoundry/libcfbuildpack/helper"
	"github.com/cloudfoundry/libcfbuildpack/layers"
)

// Dependency indicates that a JVM application should be run with Azure Application Insights enabled.
const Dependency = "azure-application-insights-java"

// Agent represents an agent contribution by the buildpack.
type Agent struct {
	buildpack buildpack.Buildpack
	layer     layers.DependencyLayer
}

// Contribute makes the contribution to launch.
func (a Agent) Contribute() error {
	return a.layer.Contribute(func(artifact string, layer layers.DependencyLayer) error {
		layer.Logger.SubsequentLine("Copying to %s", layer.Root)

		destination := filepath.Join(layer.Root, layer.ArtifactName())

		if err := helper.CopyFile(artifact, destination); err != nil {
			return err
		}

		if err := helper.CopyFile(filepath.Join(a.buildpack.Root, "AI-Agent.xml"), filepath.Join(layer.Root, "AI-Agent.xml")); err != nil {
			return err
		}

		return layer.AppendLaunchEnv("JAVA_OPTS", " -javaagent:%s", destination)
	}, layers.Launch)
}

// NewAgent creates a new Agent instance.
func NewAgent(build build.Build) (Agent, bool, error) {
	bp, ok := build.BuildPlan[Dependency]
	if !ok {
		return Agent{}, false, nil
	}

	deps, err := build.Buildpack.Dependencies()
	if err != nil {
		return Agent{}, false, err
	}

	dep, err := deps.Best(Dependency, bp.Version, build.Stack)
	if err != nil {
		return Agent{}, false, err
	}

	return Agent{build.Buildpack, build.Layers.DependencyLayer(dep)}, true, nil
}
