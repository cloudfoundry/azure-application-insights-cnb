/*
 * Copyright 2018-2019 the original author or authors.
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

package java_test

import (
	"path/filepath"
	"testing"

	"github.com/cloudfoundry/azure-application-insights-cnb/java"
	"github.com/cloudfoundry/libcfbuildpack/buildpackplan"
	"github.com/cloudfoundry/libcfbuildpack/test"
	"github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestAgent(t *testing.T) {
	spec.Run(t, "Agent", func(t *testing.T, _ spec.G, it spec.S) {

		g := gomega.NewWithT(t)

		var f *test.BuildFactory

		it.Before(func() {
			f = test.NewBuildFactory(t)
		})

		it("returns true if build plan does exist", func() {
			f.AddPlan(buildpackplan.Plan{Name: java.Dependency})
			f.AddDependency(java.Dependency, filepath.Join("testdata", "stub-azure-application-insights-agent.jar"))

			_, ok, err := java.NewAgent(f.Build)
			g.Expect(err).NotTo(gomega.HaveOccurred())
			g.Expect(ok).To(gomega.BeTrue())
		})

		it("returns false if build plan does not exist", func() {
			_, ok, err := java.NewAgent(f.Build)
			g.Expect(err).NotTo(gomega.HaveOccurred())
			g.Expect(ok).To(gomega.BeFalse())
		})

		it("contributes agent", func() {
			f.AddPlan(buildpackplan.Plan{Name: java.Dependency})
			f.AddDependency(java.Dependency, filepath.Join("testdata", "stub-azure-application-insights-agent.jar"))
			test.TouchFile(t, filepath.Join(f.Build.Buildpack.Root, "AI-Agent.xml"))

			a, ok, err := java.NewAgent(f.Build)
			g.Expect(err).NotTo(gomega.HaveOccurred())
			g.Expect(ok).To(gomega.BeTrue())

			g.Expect(a.Contribute()).To(gomega.Succeed())

			layer := f.Build.Layers.Layer("azure-application-insights-java")
			g.Expect(layer).To(test.HaveLayerMetadata(false, false, true))
			g.Expect(filepath.Join(layer.Root, "stub-azure-application-insights-agent.jar")).To(gomega.BeARegularFile())
			g.Expect(filepath.Join(layer.Root, "AI-Agent.xml")).To(gomega.BeARegularFile())
			g.Expect(layer).To(test.HaveAppendLaunchEnvironment("JAVA_OPTS", " -javaagent:%s",
				filepath.Join(layer.Root, "stub-azure-application-insights-agent.jar")))
		})

	}, spec.Report(report.Terminal{}))
}
