#!/bin/sh
TZ=US/Eastern ${GOPATH}/bin/clock -port 8010 &
TZ=Asia/Tokyo ${GOPATH}/bin/clock -port 8020 &
TZ=Europe/London ${GOPATH}/bin/clock -port 8030 &
${GOPATH}/bin/clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
killall clock
