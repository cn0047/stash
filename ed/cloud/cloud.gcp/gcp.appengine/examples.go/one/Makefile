export GOPATH=${CURDIR}

env:
	echo GOPATH=$(GOPATH)

dev:
	~/.google-cloud-sdk/bin/dev_appserver.py \
		--skip_sdk_update_check=false \
		--log_level=debug \
		--port=8080 --admin_port=8000 \
		--storage_path=$(GOPATH)/.data --support_datastore_emulator=false \
		--go_debugging=true \
		$(GOPATH)/src/go-app/app.yaml
