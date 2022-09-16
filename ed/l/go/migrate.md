migrate
-

[github](https://github.com/golang-migrate/migrate)

````sh
# download
v="v4.15.1"
os="darwin"
arch="amd64"
echo "==> Init migration tool:"
rm -rf migrate
mkdir migrate
curl -sL \
  "https://github.com/golang-migrate/migrate/releases/download/${v}/migrate.${os}-${arch}.tar.gz" \
  -o migrate/migrate.tar.gz
tar -xf migrate/migrate.tar.gz -C ./migrate



conn="projects/test-project/instances/test-instance/databases/test-db"
p="./ops/db/migrations" # path to migrations dir
./migrate/migrate -database "spanner://${conn}" -path $p up
./migrate/migrate -database "spanner://${conn}" -path $p version
./migrate/migrate -database "spanner://${conn}" -path $p down

````
