
go build -trimpath -gcflags=-trimpath=%CD% -asmflags=-trimpath=%CD% -ldflags "-s -w"

// Push to docker hub

docker build -t vizitiuroman/cronjobs-service .

docker tag vizitiuroman/cronjobs-service vizitiuroman/cronjobs-service:1.3.2

docker push vizitiuroman/cronjobs-service:1.3.2