PROJECT=myapp
NOW=$(shell date +'%Y%m%d_%H%M')
FOLDER_TMP := $(shell mktemp -d)

#SHELL=/bin/zsh
SHELL=/bin/bash
.SHELLFLAGS = -e -o pipefail -c
GET_IP_LOCAL=eval hostname -I | grep -oE '\b([0-9]{1,3}\.){3}[0-9]{1,3}\b' | head -1

################################################################
# https://askubuntu.com/questions/1410401/arguments-for-make
COMMAND := $(firstword $(MAKECMDGOALS))
ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

%::
	@true

################################################################

include ./envs/.backups.env
include ./envs/.postgres.env
include ./envs/.minio.env

POSTGRES_DB_USER=$(POSTGRES_USER)
POSTGRES_DB_PASSWORD=$(POSTGRES_PASSWORD)
POSTGRES_DB_NAME=$(DB_NAME)
POSTGRES_FILE_NAME=$(NOW)_$(DB_NAME).dump
POSTGRES_FILE_NAME_RESTORE=$(DUMP_FILE_NAME)
KC_REALM_FILE_NAME_RESTORE=$(KC_FILE_NAME)
KC_REALM_FILE_NAME=$(NOW)_realm.json

################################################################

up: infra-up

down: infra-down

logs: infra-logs

restart: infra-restart

infra-up:
	docker compose -f ./docker-compose.yaml -p $(PROJECT) up -d

infra-build:
	docker compose -f ./docker-compose.yaml -p $(PROJECT) build

infra-down:
	docker compose -f ./docker-compose.yaml -p $(PROJECT) down --remove-orphans $(ARGS)

infra-logs:
	docker compose -f ./docker-compose.yaml -p $(PROJECT) logs -t -f

infra-restart:
	docker compose -f ./docker-compose.yaml -p $(PROJECT) restart $(ARGS)

now:
	echo ${NOW}

set-ip:
	echo -e ""export IP_ADDRESS_LOCAL=eval hostname -I | grep -oE '\b([0-9]{1,3}\.){3}[0-9]{1,3}\b' | head -1"" >> ip.txt

cleanup:
	find ../ -name 'node_modules' -type d -prune -exec rm -rf '{}' +
	rm -rf ../dist
	find ../ -name '.next' -type d -prune -exec rm -rf '{}' +


###################################### NGINX #######################################

mk-install:
  mkdir -p ./certificates && \
	cd ./certificates && \
	curl -JLO "https://dl.filippo.io/mkcert/latest?for=linux/amd64" && \
	chmod +x mkcert-v*-linux-amd64 && \
	sudo mv mkcert-v*-linux-amd64 /usr/local/bin/mkcert && \
	sudo apt install libnss3-tools && \
	mkcert -install

mk-certs:
	mkcert -cert-file certificates/myapp-com-cert.pem -key-file certificates/myapp-com-key.pem "myapp.com" "*.myapp.com" && \
	cp $(shell eval mkcert -CAROOT)/rootCA.pem ./certificates && \
	cat ./certificates/myapp-com-cert.pem ./certificates/rootCA.pem > ./certificates/myapp-com.crt

mk-certs-deploy:
	sudo mkdir -p /certificates && \
	sudo cp ./certificates/myapp-com.crt /certificates/myapp-com.crt && \
	sudo cp ./certificates/myapp-com-cert.pem /certificates/myapp-com-cert.pem && \
	sudo cp ./certificates/myapp-com-key.pem /certificates/myapp-com-key.pem


mk-certs-hosts:
	./manage-etc-hosts.sh add 127.0.0.1 portal.myapp.com
	./manage-etc-hosts.sh add 127.0.0.1 api.myapp.com
	./manage-etc-hosts.sh add 127.0.0.1 api-go.myapp.com
	./manage-etc-hosts.sh add 127.0.0.1 id.myapp.com
	./manage-etc-hosts.sh add 127.0.0.1 files.myapp.com
	./manage-etc-hosts.sh add 127.0.0.1 cms.myapp.com
	./manage-etc-hosts.sh add 127.0.0.1 ws.myapp.com
	./manage-etc-hosts.sh add 127.0.0.1 mail.myapp.com

mk-ca-trust:
	sudo cp ./certificates/myapp-com.crt /usr/local/share/ca-certificates/myapp-com.crt && \
	sudo update-ca-certificates

###################################### KEYCLOAK #######################################
kc-backup: kc-export kc-copy-from

ks-restore: kc-copy-to kc-import

kc-export:
	docker exec -it ${PROJECT}-keycloak-1 /bin/bash -c "cd && bin/kc.sh export --file /tmp/realm.json"

kc-import:
	 docker exec -it ${PROJECT}-keycloak-1 /bin/bash -c "cd && bin/kc.sh import --file /tmp/${KC_REALM_FILE_NAME_RESTORE}"

kc-copy-from:
	docker cp ${PROJECT}-keycloak-1:/tmp/realm.json ./backups/keycloak/${KC_REALM_FILE_NAME}

kc-copy-to:
	docker cp ./backups/keycloak/${KC_REALM_FILE_NAME_RESTORE} ${PROJECT}-keycloak-1:/tmp/

api-token:
	curl --request POST \
  --url http://localhost:25006/realms/myapp/protocol/openid-connect/token \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data username=api \
  --data password=Api@temp123 \
  --data grant_type=password \
  --data client_id=portal-app


###################################### POSTGRES #######################################
pg-backup:
	docker exec -it ${PROJECT}-postgres-1 /bin/bash -c "pg_dump -U $(POSTGRES_DB_USER) -d $(POSTGRES_DB_NAME) > /tmp/$(POSTGRES_FILE_NAME)"
	mkdir -p ./backups/postgres
	docker cp ${PROJECT}-postgres-1:/tmp/$(POSTGRES_FILE_NAME) ./backups/postgres
	@printf "%s\n" "============================================="
	@printf "%s\n" "Nome do banco inserido: $(POSTGRES_DB_NAME)"
	@printf "%s\n" "Nome do arquivo: $(POSTGRES_FILE_NAME)"
	@printf "%s\n" "Status Postgres: Dump gerado com sucesso!"
	@printf "%s\n" "============================================="

pg-schema:
	docker exec -it ${PROJECT}-postgres-1 /bin/bash -c "pg_dump -U $(POSTGRES_DB_USER) -d $(POSTGRES_DB_NAME) --schema-only> /tmp/schema.sql"
	mkdir -p ./backups/postgres
	docker cp ${PROJECT}-postgres-1:/tmp/schema.sql ./backups/postgres
	@printf "%s\n" "============================================="
	@printf "%s\n" "Nome do banco inserido: $(POSTGRES_DB_NAME)"
	@printf "%s\n" "Nome do arquivo: schema.sql"
	@printf "%s\n" "Status Postgres: Dump gerado com sucesso!"
	@printf "%s\n" "============================================="

pg-restore:
	docker container stop ${PROJECT}-cms-1
	docker cp ./backups/postgres/$(POSTGRES_FILE_NAME_RESTORE) ${PROJECT}-postgres-1:/tmp
	docker exec -it ${PROJECT}-postgres-1  psql -U $(POSTGRES_DB_USER) -c "DROP DATABASE IF EXISTS $(POSTGRES_DB_NAME);"
	docker exec -it ${PROJECT}-postgres-1  psql -U $(POSTGRES_DB_USER) -c "CREATE DATABASE $(POSTGRES_DB_NAME);"
	docker exec -it ${PROJECT}-postgres-1 /bin/bash -c "psql -U $(POSTGRES_DB_USER) -d $(POSTGRES_DB_NAME) < /tmp/$(POSTGRES_FILE_NAME_RESTORE)"
	docker container start ${PROJECT}-cms-1
	@printf "%s\n" "============================================="
	@printf "%s\n" "Nome do banco inserido: $(POSTGRES_DB_NAME)"
	@printf "%s\n" "Nome do arquivo: $(POSTGRES_FILE_NAME_RESTORE)"
	@printf "%s\n" "Status: Restauração do banco gerada com sucesso!"
	@printf "%s\n" "============================================="


###################################### DBMATE #######################################
db-install:
	sudo curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64
	sudo chmod +x /usr/local/bin/dbmate

db-up:
	cd ../database && dbmate up

db-down:
	cd ../database/ && dbmate down

db-status:
	cd ../database/ && dbmate status


###################################### MINIO #######################################

mc-install:
	wget https://dl.min.io/client/mc/release/linux-amd64/mc
	chmod +x mc
	sudo mv ./mc /usr/local/bin

mc-config:
	mc config host add myapp http://127.0.0.1:25004 $(MINIO_ROOT_USER) $(MINIO_ROOT_PASSWORD)

mc-backup:
	mc cp -r  myapp/$(MINIO_BUCKET)/ ./backups/minio/$(MINIO_BUCKET)

mc-restore:
	mc cp -r  ./backups/minio/$(MINIO_BUCKET)/ myapp/$(MINIO_BUCKET)

mc-clear-temp:
	rm -rf ./backups/minio/*

###################################### ENVS ##########################################
envs-create:
	cp ./envs/.backups.env.example ./envs/.backups.env && \
	cp ./envs/.postgres.env.example ./envs/.postgres.env && \
	cp ./envs/.cms.env.example ./envs/.cms.env && \
	cp ./envs/.minio.env.example ./envs/.minio.env && \
	cp ./envs/.keycloak.env.example ./envs/.keycloak.env && \
	cp ../database/.env.example  ../database/.env
	@printf "%s\n" "================================================================================="
	@printf "%s\n" "Envs criados."
