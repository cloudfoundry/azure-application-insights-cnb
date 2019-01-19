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

package main

import (
	"testing"

	"github.com/cloudfoundry/libcfbuildpack/test"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestAzureApplicationInsightsProperties(t *testing.T) {
	spec.Run(t, "Azure Application Insights Properties", func(t *testing.T, when spec.G, it spec.S) {

		g := NewGomegaWithT(t)

		when("service is bound", func() {

			it("returns instrumentation key", func() {
				defer test.ReplaceEnv(t, "CNB_SERVICES", `{
  "azure-application-insights": [
    {
      "credentials": {
		"instrumentation_key": "test-instrumentation-key"
      },
      "label": "azure-application-insights"
    }
  ]
}
`)()

				p, code, err := p()

				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(code).To(Equal(0))
				g.Expect(p).To(ContainElement("-DAPPLICATION_INSIGHTS_IKEY=test-instrumentation-key"))
			})
		})

		when("service is not bound", func() {

			it("returns empty", func() {
				p, code, err := p()

				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(code).To(Equal(0))
				g.Expect(p).To(BeEmpty())
			})
		})
	}, spec.Report(report.Terminal{}))
}

