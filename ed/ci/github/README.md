GitHub Actions
-

````
${{secrets.MY_KEY}}
````

````yaml
env:
  HASH: $(git rev-parse --short "$GITHUB_SHA")
  BRANCH: ${GITHUB_REF##*/}

jobs:
  _not_push_to_gcr_1:
    runs-on: ubuntu-latest
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
