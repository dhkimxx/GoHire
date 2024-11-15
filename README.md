```sh
docker run -d \
  --name mariadb \
  -e MARIADB_ROOT_PASSWORD=1234 \
  -e MARIADB_DATABASE=lark-gitlab-bridge \
  -p 3306:3306 \
  mariadb:10.11
```

```
go test -v ./...
```