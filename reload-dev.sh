mvn clean install -Dmaven.test.skip=true
docker compose down
docker compose up -d
