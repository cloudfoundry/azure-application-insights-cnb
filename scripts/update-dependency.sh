uri() {
  echo "https://github.com/Microsoft/ApplicationInsights-Java/releases/download/$(version)/$(basename ../dependency/applicationinsights-agent-*.jar)"
}

sha256() {
  shasum -a 256 ../dependency/applicationinsights-agent-*.jar | cut -f 1 -d ' '
}
