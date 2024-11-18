```sh
docker run -d \
  --name mariadb \
  -e MARIADB_ROOT_PASSWORD=1234 \
  -e MARIADB_DATABASE=go-hire \
  -p 3306:3306 \
  mariadb:10.11
```

```
go test -v ./...
```

- atlas
```
curl -sSf https://atlasgo.sh | sh
go get -u ariga.io/atlas-provider-gorm
```

atlas migrate diff --dev-url "mysql://root:1234@127.0.0.1:3306/go-hire" --to "file://migrations"

go get entgo.io/ent/cmd/ent
go get entgo.io/ent
go run entgo.io/ent/cmd/ent init User
go run entgo.io/ent/cmd/ent generate ./ent/schema
go generate ./ent
