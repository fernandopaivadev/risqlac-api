echo "git pull"
git pull
echo "go mod tidy"
go mod tidy
echo "go build ."
go build .
echo "sudo service risqlac-api restart"
sudo service risqlac-api restart
