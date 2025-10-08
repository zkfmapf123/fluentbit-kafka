up:
	cd infrastructure && docker-compose up --build -d && \
	docker logs -f fluent-bit

down:
	cd infrastructure && docker-compose down

kafka:
	cd infrastructure && docker-compose -f docker-compose.kafka.yml up --build -d

kafka-down:
	cd infrastructure && docker-compose -f docker-compose.kafka.yml down