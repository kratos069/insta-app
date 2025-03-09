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
