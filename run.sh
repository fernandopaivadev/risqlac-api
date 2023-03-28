APP_NAME="risqlac-api"

echo "==START=="

echo "git pull"
git pull

echo "go mod tidy"
go mod tidy

echo "build project"
go build -o $APP_NAME .

echo "service restart"
sudo service $APP_NAME restart

echo "==DONE=="
