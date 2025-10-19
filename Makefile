up:
	cd infrastructure && docker-compose up --build -d && \
	docker logs -f fluent-bit

down:
	cd infrastructure && docker-compose down

kafka:
	cd infrastructure && docker-compose -f docker-compose.local.kafka.yml up --build -d

kafka-down:
	cd infrastructure && docker-compose -f docker-compose.local.kafka.yml down

all-up:
	cd infrastructure && docker-compose \
	-f docker-compose.yml \
	-f docker-compose.local.kafka.yml \
	-f docker-compose.connector.yml \
	up --build -d && \
	sleep 10 && \
	curl -X DELETE http://localhost:8083/connectors/mysql-source-connector && \
	curl -X POST http://localhost:8083/connectors \
  		-H "Content-Type: application/json" \
  		-d @$(PWD)/configs//debizium/source.json

all-down:
	cd infrastructure && docker-compose \
	-f docker-compose.yml \
	-f docker-compose.local.kafka.yml \
	-f docker-compose.connector.yml \
	down 