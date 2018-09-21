CircleCI
-

[doc](https://circleci.com/docs/)

Parallel Jobs:

````yaml
workflows:
  version: 2
  build_and_test:
    jobs: #  build and test jobs run in parallel
      - build
      - test
````
