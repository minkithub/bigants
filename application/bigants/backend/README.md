# 개발 및 빌드용 docker-compose

## Commands
* `docker-compose up [name]` (실행하기)
* `docker-compose build [name]` (빌드하기)
* `docker-compose up --build [name]` (빌드 먼저 하고 실행하기)
* `docker-compose push [name[` (클라우드 컨테이너 레이지스트리에 올리기)

## 컨테이너 내에서 실행?
```sh
# docker-compose가 실행중일 때...
docker-compose exec <name> [명령어...]
ex) docker-compose exec <name> python3 manage.py migrate
```

## Credential
개발자들은 각각 keyfile.json을 발급받으며 레포에는 포함하지 않는다.
