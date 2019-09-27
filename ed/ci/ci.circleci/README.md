CircleCI
-

[doc](https://circleci.com/docs/)
[conf](https://circleci.com/docs/2.0/configuration-reference/#section=configuration)
[env vars](https://circleci.com/docs/2.0/env-vars)
[fan-out](https://circleci.com/docs/2.0/workflows/#fan-outfan-in-workflow-example)
[parallelism](https://circleci.com/docs/2.0/parallelism-faster-jobs/#section=projects)
parallelism looks like [this](https://monosnap.com/file/5nrtmQDEyLbK3eff7Wa8BMpoMRn4kQ)

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
