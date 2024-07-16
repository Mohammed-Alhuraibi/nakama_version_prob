FROM registry.heroiclabs.com/heroiclabs/nakama-pluginbuilder:3.22.0  AS builder

ENV GO111MODULE on
ENV CGO_ENABLED 1
ENV GOPRIVATE "github.com/FrostlineGamesRepos/FrostlineGames-Nakama"

WORKDIR /backend

RUN apt-get update && \
    apt-get -y upgrade && \
    apt-get install -y --no-install-recommends gcc libc6-dev

RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY . .

RUN go mod tidy
RUN go mod vendor       
# RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./frostline.so

RUN go build --trimpath --gcflags "all=-N -l" --mod=vendor --buildmode=plugin -o ./frostline.so


FROM registry.heroiclabs.com/heroiclabs/nakama-dsym:3.22.0

COPY --from=builder /go/bin/dlv /nakama
COPY --from=builder /backend/frostline.so /nakama/data/modules
COPY --from=builder /backend/local.yml /nakama/data/

ENTRYPOINT [ "/bin/bash" ]
# POSTGRES_DB
# POSTGRES_USER
# POSTGRES_PASSWORD

# CMD [ "/bin/sh", "-ecx", "/nakama/nakama migrate up --database.address $POSTGRES_USER:$POSTGRES_PASSWORD@postgres:5432/$POSTGRES_DB && exec /nakama/nakama --config /nakama/data/local.yml --database.address $POSTGRES_USER:$POSTGRES_PASSWORD@postgres:5432/$POSTGRES_DB" ]
