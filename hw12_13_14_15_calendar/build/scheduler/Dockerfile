# Собираем в гошке
FROM golang:1.17.3 as build

ENV BIN_FILE /opt/calendar/sheduler
ENV CODE_DIR /go/src/
ENV CGO_ENABLED=0

RUN go install github.com/antelman107/net-wait-go@latest

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/scheduler/*

# На выходе тонкий образ
FROM alpine:3.15

LABEL ORGANIZATION="OTUS Online Education"
LABEL SERVICE="calendar_scheduler"
LABEL MAINTAINERS="vlyagusha@gmail.com"

ENV BIN_FILE /opt/calendar/sheduler
COPY --from=build ${BIN_FILE} ${BIN_FILE}

COPY --from=build /go/bin/net-wait-go /bin/net-wait-go

ENV CONFIG_FILE ${SCHEDULER_CONFIG_FILE}
COPY ./configs/scheduler_config.yaml ${CONFIG_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}
