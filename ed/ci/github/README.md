GitHub Actions
-

[context](https://docs.github.com/en/actions/learn-github-actions/contexts#github-context)
[env](https://docs.github.com/en/actions/learn-github-actions/environment-variables#default-environment-variables)
[workflow](https://docs.github.com/en/actions/using-workflows)
[custom action](https://docs.github.com/en/actions/creating-actions/about-custom-actions)
[toolkit](https://github.com/actions/toolkit)
[docker example](https://github.com/cn007b/docker-ubuntu/blob/master/.github/workflows/docker-image.yml)

Workflow - configurable automated process that will run one or more jobs.
Event - specific activity in a repo that triggers workflow run.
Job - set of steps in a workflow.
Action - custom application for the GitHub actions platform that performs a complex task.
Runner - server that runs workflow.

````sh
# set secret in repo with workflow for more debug logs:
ACTIONS_STEP_DEBUG=true

${{secrets.MY_KEY}}
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
  working-directory: "${{ github.workspace }}"
  env-name: ${{ (github.event.inputs && github.event.inputs.environment) || 'stg' }}

defaults:
  run:
    working-directory: .

jobs:
  _not_push_to_gcr_1:
    runs-on: ubuntu-latest
    env:
      CODE: 200
    steps:
    - uses: actions/checkout@v2
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
