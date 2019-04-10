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

package java_test

import (
	"path/filepath"
	"testing"

	"github.com/cloudfoundry/azure-application-insights-cnb/java"
	"github.com/cloudfoundry/libcfbuildpack/test"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func TestProperties(t *testing.T) {
	spec.Run(t, "Properties", func(t *testing.T, _ spec.G, it spec.S) {

		g := NewGomegaWithT(t)

		it("contributes new relic properties", func() {
			f := test.NewBuildFactory(t)
			test.TouchFile(t, f.Build.Buildpack.Root, "bin", "azure-application-insights-properties")

			c := java.NewProperties(f.Build)

			g.Expect(c.Contribute()).To(Succeed())

			layer := f.Build.Layers.Layer("azure-application-insights-properties")
			g.Expect(layer).To(test.HaveLayerMetadata(false, false, true))
			g.Expect(filepath.Join(layer.Root, "bin", "azure-application-insights-properties")).To(BeARegularFile())
			g.Expect(layer).To(test.HaveProfile("azure-application-insights-properties", `printf "Configuring Azure Application Insights Properties"

export JAVA_OPTS="${JAVA_OPTS} $(azure-application-insights-properties)"
`))
		})

	})
}

