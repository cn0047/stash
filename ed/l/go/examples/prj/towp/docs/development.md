Development
-

Ensure golang version is `>= 1.19`.

Use next command to run wp locally:

````sh
export GOOGLE_PROJECT_ID=test-project
export SPANNER_EMULATOR_HOST=localhost:9010
export PUBSUB_EMULATOR_HOST=localhost:8681

make run-local-env

make run
````

Use next link to open  swagger spec on local env: http://localhost:8080/

Use next base URL to reach wp API on local env: http://localhost:8080/v1/

Use next command to run tests:

````sh
make test
````

More commands on how to work with wp you can find in
[Makefile](https://github.com/to-com/wp/blob/master/Makefile).

For more info about local development, please take a look to
[Go handbook](https://engineering-handbook.to-com.org/docs/guilds/architecture/go/).

#### How to run contract tests locally

Some environment vars are needed for contract tests, in CI they get set in the workflow
to be able to run the tests locally you need to create `.env` file in the project dir and add the following vars there:
  `PACT_BROKER_URL`,
  `PACT_BROKER_TOKEN`.
The values should be taken from here: [pactflow](https://to.pactflow.io/settings/api-tokens)
