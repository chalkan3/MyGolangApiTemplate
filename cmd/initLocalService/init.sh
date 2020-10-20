#!/bin/bash

echo "************** [INIT DOCKER COMPOSE] *******************"
docker-compose  up -d
echo "************** [INIT API SERVICE] *******************"
go run cmd/whats-app-api-server/main.go