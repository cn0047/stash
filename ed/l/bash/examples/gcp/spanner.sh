#!/bin/bash

SCRIPT=`realpath $0`
SCRIPT_DIR=`dirname $SCRIPT`
MIGRATIONS_DIR=$SCRIPT_DIR/../../migration

startSpanner() {
  cfg=$1
  dbi=$2
  db=$3

  echo "==> Init Google Cloud Spanner:"
  gcloud config configurations activate $cfg || gcloud config configurations create $cfg
  gcloud config set auth/disable_credentials true
  gcloud config set api_endpoint_overrides/spanner 'http://localhost:9020/'

  echo "==> Start Google Cloud Spanner emulator:"
  gcloud emulators spanner start &
  gcloud spanner instances create $dbi --config=$cfg --description="TestEmulatorInstance" --nodes=1
  gcloud spanner databases create $db --instance=$dbi

  echo "==> Load DB migrations:"
  # Split migration file into separate queries by ; (semicolon) and execute each independently.
  for migrationFile in $(ls $MIGRATIONS_DIR); do
    query=''
    for line in $(cat $MIGRATIONS_DIR/$migrationFile); do
      query+=" $line"
      if [[ $line == *";"* ]]; then
        gcloud spanner databases ddl update $db --instance=$dbi --ddl="$query"
        query=''
      fi
    done
  done
}

stopSpanner() {
  echo "==> Stop Google Cloud Spanner emulator:"
  docker stop `docker ps | grep spanner-emulator | awk '{print $1}'`

  echo "==> Unset Google Cloud Spanner emulator configs:"
  gcloud config unset auth/disable_credentials
  gcloud config unset api_endpoint_overrides/spanner
}

case $1 in
spanner-start)
  startSpanner $2 $3 $4
  ;;

spanner-stop)
  stopSpanner
  ;;

*)
  echo "Script to work wit spanner."
  ;;
esac
