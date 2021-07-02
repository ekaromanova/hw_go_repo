port: 8080
db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
jaeger_url: http://jaeger:16686
sentry_url: http://sentry:9000
deploy:
	name: kafka_broker
	value: broker1:8080
some_app_id: testid
some_app_key: testkey