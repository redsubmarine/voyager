# VOYAGER 프로젝트
Golang, postgres db 를 사용한 기본적인 auth api 서버 프로젝트입니다. 추 후 개발될 프로젝트들의 템플릿역할을 하는 프로젝트입니다.

## 기술 스택
- asdf: [asdf](https://asdf-vm.com) 를 사용하여, 언어의 버전을 명시합니다. .tool-versions 에 현재 프로젝트에 사용할 언어의 버전을 확인할수 있습니다.
- golang: 특이사항이 없다면 최신 버전 사용을 권장합니다.
- postgres: 특이사항이 없다면 최신 버전 사용을 권장합니다. 다만, docker 배포를 위해 alpine 을 권장하며, 개발 단계에서 충분히 테스트가 되어야 합니다.

## 서버 아키텍쳐
- DB ORM: [sqlc](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html) 를 사용합니다.
- DB 마이그레이션: [golang-migrate](https://github.com/golang-migrate/migrate)를 사용합니다.

* golang-migrate 설치
```
brew install golang-migrate
migrate create Name # up, down 파일 생성.
migrate goto V # V 버전으로 적용
migrate up/down # 버전 업 다운 적용.
```
migrate 사용
```
migrate create -ext sql -dir db/migration -seq "init_scheme"
# db/migration 폴더에 000001_init-scheme.up/down.sql 이 생성됨.
```

## sqlc 사용
- [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html) 설치
```
sqlc init
sqlc generate
```
## 
```
go get github.com/lib/pq
```
## testfy
```
go get github.com/stretchr/testify
```

## gin
```
go get -u github.com/gin-gonic/gin
```

## viper(환경변수)
```
go get github.com/spf13/viper
```

## gomock
- [gomock](https://github.com/uber-go/mock)
```
go install go.uber.org/mock/mockgen@latest
```

## api 목록
### role: Admin
- signIn
    - signUp 은 보안을 위해 제공하지 않는다. 필요시, 직접 db 에 추가.
- blockUser
- deleteUser
- listUser

### role: User
- signUp
- signIn