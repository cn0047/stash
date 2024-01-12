GitLab Pipelines
-

[syntax](https://docs.gitlab.com/ee/ci/yaml/index.html)
[predefined env vars](https://docs.gitlab.com/ee/ci/variables/predefined_variables.html)
[example 1](https://gitlab.com/cn007b/test/-/blob/main/.gitlab-ci.yml)

````sh
.gitlab-ci.yml
````

````yaml
stages:
  - cleanup_test

cleanup_test_1:
  stage: cleanup_test
  script:
    - 'echo "==> Pipeline failed, Running additional commands..."'
  # skip when not failed, see: https://gitlab.com/cn007b/test/-/pipelines/1126850785
  # executed when failed, see: https://gitlab.com/cn007b/test/-/pipelines/1126851796
  when: on_failure # on_success|on_failure|always
````
