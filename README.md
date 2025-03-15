<!-- new migration file command -->
migrate create -ext sql -dir db/migration -seq add_users

<!-- access postgres shell -->
docker exec -it postgres17 psql -U root

<!-- github upload steps -->
git init
git add .
git commit -m "message for commit..."
git remote add origin https://github.com/your-username/your-repo-name.git
git remote -v
git branch -M main
git push -u origin main

<!-- build docker image of project -->
docker build -t insta-app:latest .

<!-- to check ip address of running container -->
docker container inspect postgres17

<!-- remove a container -->
docker rm insta-app

<!-- start container from built image -->
docker run --name insta-app --network insta-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres17:5432/insta-app?sslmode=disable" insta-app:latest

<!-- check networks of docker containers -->
docker network ls

<!-- for more details about a docker network (bridge is name) -->
docker network inspect bridge

<!-- creating a new network so postgres and insta can have same network -->
docker network create insta-network

<!-- to connect network  -->
docker network connect insta-network postgres17