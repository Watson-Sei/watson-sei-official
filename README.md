# watson-sei-official
watson-sei.tokyoがドメインのなんでもサイト

```
/docker/db/.env

HOST=tcp(db)
MYSQL_ROOT_PASSWORD=root
MYSQL_USER=docker_user
MYSQL_PASSWORD=docker_pass
MYSQL_DATABASE=watson-db
ACCESS_TOKEN_SECRETKEY=f3*pcq$#f3#79%nboqude9g8y3^hk3#+x!g3pm*r79lrcp=4-$
REFRESH_TOKEN_SECRETKEY=wh0+t&#w)tsb7ipvd0^hu@fkzml93dty!_h(o^xz7hmcnvf4#z
```

```
## /docker/front/config/.env.development
API_URL="http://api:8080"
API_URL_BROWSER='http://localhost:8080'
BASE_URL="http://localhost"
```
```
## /dockr/front/config/.env.production
API_URL="http://api:8080"
API_URL_BROWSER="http://localhost:8080"
BASE_URL="https://www.domain"
GA_ID="google analytics"
```


オレオレ証明書なしの場合
```
docker-compose -f docker-compose.development.yml up -d --build
```

オレオレ証明書ありの場合
```
docker-compose -f docker-compose.production.yml up -d --build
```

gin-adminでcreate admin user
```
docker exec -it watson-sei-official_api_1 sh
↓
cd $GOPATH/bin
↓
./gin-admin create -u <username> -p <password>
```