clean:
	docker rm -f $(docker ps -a -q)

up:
	docker compose -f ./docker-compose.yml up -d --no-deps --build

ps:
	docker compose -f ./docker-compose.yml ps -a

logs:
	docker compose -f ./docker-compose.yml logs

stop:
	docker compose -f ./docker-compose.yml stop

update:
	git pull && make stop && make up