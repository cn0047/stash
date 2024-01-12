CircleCI
-

[doc](https://circleci.com/docs/)
[conf](https://circleci.com/docs/2.0/configuration-reference/#section=configuration)
[env vars](https://circleci.com/docs/2.0/env-vars)
[fan-out](https://circleci.com/docs/2.0/workflows/#fan-outfan-in-workflow-example)
[parallelism](https://circleci.com/docs/2.0/parallelism-faster-jobs/#section=projects)
[parallelism example](https://monosnap.com/file/5nrtmQDEyLbK3eff7Wa8BMpoMRn4kQ)
[example 1](https://github.com/cn007b/api-gateway/tree/master/.circleci)
[example 2](https://github.com/cn007b/eop/tree/master/.circleci)
[example 3](https://github.com/cn007b/monitoring/tree/master/.circleci)
[example 4](https://github.com/cn007b/short-string-number/tree/master/.circleci)
[example 5](https://github.com/thepkg/awsl/tree/master/.circleci)
[example 6](https://github.com/thepkg/gcd/tree/master/.circleci)
[example 7](https://github.com/thepkg/recover/tree/master/.circleci)
[example 8](https://github.com/thepkg/rest/tree/master/.circleci)
[example 9](https://github.com/thepkg/strings/tree/master/.circleci)

To fail build - exit with non zero code.

Parallel Jobs:

````yaml
workflows:
  version: 2
  build_and_test:
    jobs: #  build and test jobs run in parallel
      - build
      - test
````
