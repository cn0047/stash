SonarQube
_

sonar-project.properties:

````sh
sonar.projectKey=org_repo
sonar.projectName=org_repo

sonar.c.file.suffixes=-
sonar.cpp.file.suffixes=-
sonar.objc.file.suffixes=-

sonar.sources=.
sonar.tests=.
sonar.test.inclusions=**/*_test.go

sonar.go.coverage.reportPaths=**/coverage.out
sonar.go.tests.reportPaths=**/tests_run.json
sonar.coverage.exclusions=**/*_test.go,cmd/**

````
