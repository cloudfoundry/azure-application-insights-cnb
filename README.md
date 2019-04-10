# `azure-application-insights-cnb`
The Cloud Foundry Azure Application Insights Buildpack is a Cloud Native Buildpack V3 that provides the [Application Insights][a] agent and configuration to applications.

This buildpack is designed to work in collaboration with bound service instances.

[a]: https://docs.microsoft.com/en-us/azure/azure-monitor/app/app-insights-overview

## Detection
The detection phase passes if

* A service is bound with a payload containing a `binding_name`, `instance_name`, `label`, or `tag` containing `azure-application-insights` as a substring and build plan contains `jvm-application`
  * Contributes `azure-application-insights-java` to the build plan

## Build
If the build plan contains

* `azure-application-insights-java`
  * Contributes an Azure Application Insights agent and `AI-Agent.xml` to a layer marked `launch`
  * Sets `-javaagent` `JAVA_OPTS`

## License
This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0

