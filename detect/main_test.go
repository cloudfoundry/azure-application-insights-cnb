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

package main

import (
	"testing"

	"github.com/buildpacks/libbuildpack/v2/buildplan"
	"github.com/cloudfoundry/azure-application-insights-cnb/java"
	"github.com/cloudfoundry/libcfbuildpack/v2/detect"
	"github.com/cloudfoundry/libcfbuildpack/v2/services"
	"github.com/cloudfoundry/libcfbuildpack/v2/test"
	"github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestDetect(t *testing.T) {
	spec.Run(t, "Detect", func(t *testing.T, _ spec.G, it spec.S) {

		g := gomega.NewWithT(t)

		var f *test.DetectFactory

		it.Before(func() {
			f = test.NewDetectFactory(t)
		})

		it("fails without service", func() {
			g.Expect(d(f.Detect)).To(gomega.Equal(detect.FailStatusCode))
		})

		it("passes with service and jvm-application", func() {
			f.AddService("azure-application-insights", services.Credentials{"instrumentation_key": "instrumentation_value"})

			g.Expect(d(f.Detect)).To(gomega.Equal(detect.PassStatusCode))
			g.Expect(f.Plans).To(test.HavePlans(buildplan.Plan{
				Provides: []buildplan.Provided{
					{Name: java.Dependency},
				},
				Requires: []buildplan.Required{
					{Name: java.Dependency},
					{Name: "jvm-application"},
				},
			}))
		})
	}, spec.Report(report.Terminal{}))
}
