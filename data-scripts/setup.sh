#!/usr/bin/env bash

docker run --name bookings_dev \
    -p 5432:5432 \
    -e 'POSTGRES_PASSWORD=P@ssw0rd' \
    -d postgres:14.5
