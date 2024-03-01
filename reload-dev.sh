cd ./geniusecommerce 
mvn compile war:exploded
cd ..
docker compose down
docker compose up -d
