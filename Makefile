up:
	cd infrastructure && docker-compose up --build -d && \
	docker logs -f fluent-bit

down:
	cd infrastructure && docker-compose down