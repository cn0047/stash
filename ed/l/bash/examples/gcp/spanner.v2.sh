#!/bin/bash

SCRIPT=`realpath $0`
SCRIPT_DIR=`dirname $SCRIPT`
MIGRATIONS_DIR=$SCRIPT_DIR/../db/migrations

gCloudInit() {
  cfg=$1
  prj=$2

  echo "==> Init Google Cloud:"
  gcloud config set disable_prompts true
  gcloud config configurations activate $cfg || gcloud config configurations create $cfg
  gcloud config set project $prj
  gcloud config set auth/disable_credentials true
}

startSpanner() {
  cfg=$1
  prj=$2
  dbi=$3
  db=$4

  echo "==> Init Google Cloud Spanner:"
  gcloud config set api_endpoint_overrides/spanner "http://localhost:9020/"

  echo "==> Start Google Cloud Spanner emulator:"
  gcloud -q emulators spanner start &
  gcloud -q spanner instances create $dbi --config=$cfg --description="TestEmulatorInstance" --nodes=1
  gcloud -q spanner databases create $db --instance=$dbi
}

downloadMigrationTool() {
  os=$1
  arch=$2

  echo "==> Init migration tool:"
  echo "OS: {$os}, Arch: {$arch}"
  rm -rf migrate
  mkdir migrate
  curl -sL \
    "https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.${os}-${arch}.tar.gz" \
    -o migrate/migrate.tar.gz
  tar -xf migrate/migrate.tar.gz -C ./migrate
}

loadMigrations() {
  cfg=$1
  prj=$2
  dbi=$3
  db=$4

  conn="projects/${prj}/instances/${dbi}/databases/${db}"

  echo "==> Load DB migrations:"
  ./migrate/migrate -database "spanner://${conn}" -path "$MIGRATIONS_DIR" up
  echo "==> Current DB migration version:"
  ./migrate/migrate -database "spanner://${conn}" -path "$MIGRATIONS_DIR" version
}

stopSpanner() {
  echo "==> Stop Google Cloud Spanner emulator:"
  docker stop `docker ps | grep spanner-emulator | awk '{print $1}'`

  echo "==> Unset Google Cloud Spanner emulator configs:"
  gcloud config unset auth/disable_credentials
  gcloud config unset api_endpoint_overrides/spanner
}

case $1 in
gcloud-init)
  cfg=$2
  prj=$3
  gCloudInit $cfg $prj
  ;;

spanner-start)
  cfg=$2
  prj=$3
  dbi=$4
  db=$5
  startSpanner $cfg $prj $dbi $db
  ;;

migrate-init)
  os=$2
  arch=$3
  downloadMigrationTool $os $arch
  ;;

spanner-init)
  cfg=$2
  prj=$3
  dbi=$4
  db=$5
  loadMigrations $cfg $prj $dbi $db
  ;;

spanner-stop)
  stopSpanner
  ;;

*)
  echo "Script to work wit spanner."
  ;;
esac
