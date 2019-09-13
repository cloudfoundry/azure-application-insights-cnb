uri() {
  echo "https://github.com/Microsoft/ApplicationInsights-Java/releases/download/$(version)/$(basename ../${DEPENDENCY}/applicationinsights-agent-*.jar)"
}

sha256() {
  shasum -a 256 "../${DEPENDENCY}/applicationinsights-agent-*.jar" | cut -f 1 -d ' '
}
