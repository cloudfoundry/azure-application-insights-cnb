#!/usr/bin/env bash

set -euo pipefail

if [[ -d $PWD/go-module-cache && ! -d ${GOPATH}/pkg/mod ]]; then
  mkdir -p ${GOPATH}/pkg
  ln -s $PWD/go-module-cache ${GOPATH}/pkg/mod
fi

commit() {
  git commit -a -m "Dependency Upgrade: $1 $2" || true
}

sha256() {
  shasum -a 256 $1 | cut -f 1 -d ' '
}

uri() {
  echo "https://github.com/Microsoft/ApplicationInsights-Java/releases/download/$(cat ../azure-application-insights-java/version)/$(basename $1)"
}

cd "$(dirname "${BASH_SOURCE[0]}")/.."
git config --local user.name "$GIT_USER_NAME"
git config --local user.email ${GIT_USER_EMAIL}

go build -ldflags='-s -w' -o bin/dependency github.com/cloudfoundry/libcfbuildpack/dependency

bin/dependency azure-application-insights-java "[\d]+\.[\d]+\.[\d]+" $(cat ../azure-application-insights-java/version) $(uri ../azure-application-insights-java/applicationinsights-agent-*.jar) $(sha256 ../azure-application-insights-java/applicationinsights-agent-*.jar)
commit azure-application-insights-java $(cat ../azure-application-insights-java/version)
