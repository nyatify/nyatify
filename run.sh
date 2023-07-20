#!/bin/bash

set -o allexport
source .env
set +o allexport

if [ "$1" == "all" ]; then
    go run ./services/api/ &

    go run ./services/scheduler/

elif [ "$1" == "scheduler" ]; then
    go run ./services/scheduler/

elif [ "$1" == "api" ]; then
    go run ./services/api/

elif [ "$1" == "bot" ]; then
    cd ./services/tg_bot
    go run main.go

elif [ "$1" == "mu" ]; then
    migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=$POSTGRES_MODE" -path "pkg/storage/migrations" up

elif [ "$1" == "md" ]; then
    migrate -database postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=$POSTGRES_MODE -path pkg/storage/migrations down 1

elif [ "$1" == "mf" ]; then
    migrate -database postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=$POSTGRES_MODE -path pkg/storage/migrations force "$2"

elif [ "$1" == "mc" ]; then
    migrate create -ext sql -dir "pkg/storage/migrations" -seq "$2"

else
    echo "Invalid argument."
fi

