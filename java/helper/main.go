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
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/cloudfoundry/libcfbuildpack/helper"
)

func main() {
	properties, code, err := p()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(code)
	}

	fmt.Println(strings.Join(properties, " "))
	os.Exit(code)
}

func p() ([]string, int, error) {
	credentials, ok, err := helper.FindServiceCredentials("azure-application-insights", "instrumentation_key")
	if err != nil {
		return nil, 1, err
	}

	if !ok {
		return nil, 0, nil
	}

	var properties []string

	if i, ok := credentials["instrumentation_key"]; ok {
		properties = append(properties, fmt.Sprintf("-DAPPLICATION_INSIGHTS_IKEY=%s", i))
	}

	sort.Strings(properties)
	return properties, 0, nil
}
