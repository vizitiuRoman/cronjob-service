go build -trimpath -gcflags=-trimpath=%CD% -asmflags=-trimpath=%CD% -ldflags "-s -w"

// Push to docker hub

docker build -t vizitiuroman/cronjob-service .

docker tag vizitiuroman/cronjob-service vizitiuroman/cronjob-service:1.3.2

docker push vizitiuroman/cronjob-service:1.3.2