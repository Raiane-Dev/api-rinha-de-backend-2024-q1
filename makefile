all:
	make prelease

## dev
drelease:
	docker-compose -f docker-compose.dev.yml up -d --build

ddown:
	docker-compose -f docker-compose.dev.yml down
	docker system prune -a -f
	docker volume prune -a -f

## hml
hrelease:
	rm -rf database/data/*
	docker-compose -f docker-compose.build.yml up -d --build
	docker-compose -f docker-compose.build.yml push

htest:
	rm -rf database/data/*
	docker-compose -f docker-compose.build.yml up -d --build

hdown:
	rm -rf ./database/rinha_api.sqlite*
	docker-compose -f docker-compose.build.yml down
	docker system prune -a -f
	docker volume prune -a -f


## prod
prelease:
	docker-compose up -d --build

pdown:
	rm -rf ./database/rinha_api.sqlite*
	docker-compose down
	docker system prune -a -f
	docker volume prune -a -f