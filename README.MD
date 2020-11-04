## describe
some go example test

## dependency

- export GO111MODULE=on && export GOPROXY=https://goproxy.cn
- go mod init project_name
- go mod tidy

## docker build
`docker build . -t goweb_app`

`docker run -p 3000:3000 goweb_app`
