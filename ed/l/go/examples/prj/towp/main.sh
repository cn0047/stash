#!/bin/bash

hash realpath || echo "Please install: realpath"
hash liquibase || echo "Please install: liquibase"

SCRIPT=`realpath $0`
SCRIPT_DIR=`dirname $SCRIPT`
LIQUIBASE_CHANGELOG=database/changelog.xml
LIQUIBASE_CLASSPATH=database/liquibase-spanner-4.5.0-all.jar

gCloudInit() {
  cfg=$1
  prj=$2

  echo "==> Init Google Cloud:"
  echo "CFG: ${cfg}, PRJ: ${prj}"
  gcloud config set disable_prompts true
  gcloud config configurations activate $cfg || gcloud config configurations create $cfg
  gcloud config set project $prj
  gcloud config set auth/disable_credentials true
}

spannerInit() {
  cfg=$1
  prj=$2
  dbi=$3
  db=$4

  echo "==> Init Google Cloud Spanner:"
  echo "CFG: ${cfg}, PRJ: ${prj}, DBI: ${dbi}, DB: ${db}"
  gcloud config set api_endpoint_overrides/spanner "http://localhost:9020/"
  gcloud -q spanner instances create $dbi --config=$cfg --description="TestEmulatorInstance" --nodes=1
  gcloud -q spanner databases create $db --instance=$dbi
}

startSpanner() {
  echo "==> Start Google Cloud Spanner emulator:"
  gcloud -q emulators spanner start &
}

loadMigrations() {
  echo "==> Load DB migrations:"
  execMigrationsCommand "$@" updateSQL
  execMigrationsCommand "$@" update
}

upMigration() {
  echo "==> Up DB migration:"
  execMigrationsCommand "$@" updateSQL
  execMigrationsCommand "$@" update
}

downMigration() {
  echo "==> Down DB migration:"
  execMigrationsCommand "$@" rollbackCountSQL 1
  execMigrationsCommand "$@" rollbackCount 1
}

execMigrationsCommand() {
  cfg=$1
  prj=$2
  dbh=$3
  dbi=$4
  db=$5
  cmd=$6
  n=$7

  conn="jdbc:cloudspanner://${dbh}/projects/${prj}/instances/${dbi}/databases/${db};usePlainText=true"
  liquibase --url=$conn --changelog-file=$LIQUIBASE_CHANGELOG --classpath=$LIQUIBASE_CLASSPATH \
    --hub-mode=off $cmd $n
}

stopSpanner() {
  echo "==> Stop Google Cloud Spanner emulator:"
  docker stop `docker ps | grep spanner-emulator | awk '{print $1}'`
}

case $1 in
spanner-start)
  startSpanner
  ;;

spanner-init)
  cfg=$2
  prj=$3
  dbi=$4
  db=$5
  spannerInit $cfg $prj $dbi $db
  ;;

gcloud-init)
  cfg=$2
  prj=$3
  gCloudInit $cfg $prj
  ;;

load-migrations)
  cfg=$2
  prj=$3
  dbh=$4
  dbi=$5
  db=$6
  loadMigrations $cfg $prj $dbh $dbi $db
  ;;

up-migration)
  cfg=$2
  prj=$3
  dbh=$4
  dbi=$5
  db=$6
  upMigration $cfg $prj $dbh $dbi $db
  ;;

down-migration)
  cfg=$2
  prj=$3
  dbh=$4
  dbi=$5
  db=$6
  downMigration $cfg $prj $dbh $dbi $db
  ;;

spanner-stop)
  stopSpanner
  ;;

*)
  echo "Script to work wit app."
  ;;
esac
