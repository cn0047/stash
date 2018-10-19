export GOPATH=${CURDIR}/go
export NODEJSPATH=${CURDIR}/nodejs
export PHPPATH=${CURDIR}/php
export GOOGLE_APPLICATION_CREDENTIALS=${CURDIR}/serviceAccount.json

env:
	echo GOPATH=$(GOPATH)
	echo NODEJSPATH=$(NODEJSPATH)
	echo PHPPATH=$(PHPPATH)

initgo:
	go get -u google.golang.org/appengine/...
	go get -u cloud.google.com/go/storage

devgo:
	~/.google-cloud-sdk/bin/dev_appserver.py \
		--skip_sdk_update_check=false \
		--log_level=debug \
		--port=8080 --admin_port=8000 \
		--storage_path=$(GOPATH)/.data --support_datastore_emulator=false \
		--go_debugging=true \
		$(GOPATH)/src/go-app/app.yaml

initnodejs:
	cd $(NODEJSPATH) && npm i && cd -

devnodejs:
	cd $(NODEJSPATH) && npm start && cd -

initphp:
	docker run -it --rm -v $(PHPPATH):/app -w /app cn007b/php composer install

devphp:
	~/.google-cloud-sdk/bin/dev_appserver.py \
		--skip_sdk_update_check=false \
		--log_level=debug \
		--port=8080 --admin_port=8000 \
		--storage_path=$(GOPATH)/.data --support_datastore_emulator=false \
		--go_debugging=true \
		$(PHPPATH)/app.yaml

deploy:
	# gcloud auth login
	# gcloud auth activate-service-account --key-file=$(GOOGLE_APPLICATION_CREDENTIALS)
	gcloud config set project itisgnp
	gcloud config list

deploygo:
	gcloud app deploy -q $(GOPATH)/src/go-app/app.yaml

deploynodejs:
	gcloud app deploy -q $(NODEJSPATH)/app.yaml

deployphp:
	gcloud app deploy -q $(PHPPATH)/app.yaml
