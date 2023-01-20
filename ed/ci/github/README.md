GitHub Actions
-

[context](https://docs.github.com/en/actions/learn-github-actions/contexts#github-context)
[env](https://docs.github.com/en/actions/learn-github-actions/environment-variables#default-environment-variables)
[workflow](https://docs.github.com/en/actions/using-workflows)
[custom action](https://docs.github.com/en/actions/creating-actions/about-custom-actions)
[toolkit](https://github.com/actions/toolkit)
[playground](https://github-actions-hero.vercel.app/)
[docker example](https://github.com/cn007b/docker-ubuntu/blob/master/.github/workflows/docker-image.yml)

Workflow - configurable automated process that will run one or more jobs.
Event - specific activity in a repo that triggers workflow run.
Job - set of steps in a workflow.
Action - custom application for the GitHub actions platform that performs a complex task.
Runner - server that runs workflow.

````sh
# set secret in repo with workflow for more debug logs:
ACTIONS_STEP_DEBUG=true

${{ secrets.MY_KEY }}
${{ secrets.GITHUB_TOKEN }}

github.ref == 'refs/heads/master' # example: refs/pull/13/merge
github.ref == 'refs/pull/13/merge'
github.event_name == 'pull_request'

${{ github.event.number }}
${{ github.HEAD_REF }} # example: refs/heads/master
${{ github.BASE_REF }} # example: master
${{ github.repository }}
${{ github.event.push.ref }} # branch name

format('Out: {0} {1} {2}', 'a', 'b', 'c')
join(github.event.issue.labels.*.name, ', ')
toJSON(value)
fromJSON(value)
contains(fromJson('["bar", "foo"]'), github.event.action)

::set-env name=DBG::1
::set-output name=CODE::200
::add-path::/path/to/dir
::add-mask::msg
::debug::msg
::debug file=name,line=1,col=1::msg
::warning file=name,line=1,col=1::msg
::error file=name,line=1,col=1::msg
::stop-comands::token # stop
::token:: # start
````

````yaml
# .github/workflows/x.yaml

on:
  push:
  pull_request:
    branches:
      - master
    types: [closed]
  workflow_dispatch: # ‼️ works only for workflow_dispatch event
    inputs:
      env: # @use: ${{ github.event.inputs.env || 'test' }}
        description: "env name"
        type: string
        required: true
        default: "test"

env:
  HASH: $(git rev-parse --short "$GITHUB_SHA")
  BRANCH: ${GITHUB_REF##*/}
  working-directory: "${{ github.workspace }}" # ${{ env.working-directory }}
  env-name: ${{ (github.event.inputs && github.event.inputs.environment) || 'stg' }}

defaults:
  run:
    working-directory: .

jobs:
  _not_push_to_gcr_1:
    runs-on: ubuntu-latest
    env:
      CODE: 200
    outputs:
      project-id: ${{ steps.setvars.outputs.project-id }}
    steps:
    - uses: actions/checkout@v2
    - name: "Set vars"
      id: setvars
      run: |
        PROJECT_ID="test-prj"
        echo PROJECT_ID=$PROJECT_ID >> $GITHUB_ENV
        # use it with: ${{ env.PROJECT_ID }}
        echo "::set-output name=project-id::$PROJECT_ID"
        # use it with: ${{ steps.setvars.outputs.project-id }}
        #
        echo "{name}={value}" >> $GITHUB_STATE
        echo "{name}={value}" >> $GITHUB_OUTPUT
    - name: Setup gcloud
      uses: GoogleCloudPlatform/github-actions/setup-gcloud@master
      with:
        version: '290.0.1'
        service_account_key: ${{ secrets.GCP_SERVICE_ACCOUNT_KEY }}
        project_id: ${{ secrets.GCP_PROJECT_ID }}
    - name: Configure docker for GCP
      run: gcloud auth configure-docker
    - name: Build docker image
      run: docker build . -f ./docker/1.17-alpine/Dockerfile -t gcr.io/${{ secrets.GCP_PROJECT_ID }}/go:1.17-alpine
    - name: Push to Google Container Registry
      run: docker push gcr.io/${{ secrets.GCP_PROJECT_ID }}/go:1.17-alpine

  _not_push_to_gcr_2:
    name: Push to GCR
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2
      - uses: GoogleCloudPlatform/github-actions/setup-gcloud@master
        with:
          service_account_key: ${{ secrets.GCP_SERVICE_ACCOUNT_KEY }}
          project_id: ${{ secrets.GCP_PROJECT_ID }}
          export_default_credentials: true
      - name: Build
        run: |-
          docker build . --file ./docker/1.17-alpine/Dockerfile -t gcr.io/${{ secrets.GCP_PROJECT_ID }}/go:1.17-alpine
      - run: |
          gcloud auth configure-docker -q
      - name: Push
        run: |-
          docker push eu.gcr.io/${{ secrets.GCP_PROJECT_ID }}/go:1.17-alpine

````

````yaml
# .github/actions/x.yaml

name: "X-Action."
description: "X-Action."

inputs:
  name:
    required: true
    description: "Name"

runs:
  using: "composite"
  steps:
    - run: echo Hello ${{ inputs.name }}.
      shell: bash
````
