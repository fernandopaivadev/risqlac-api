echo "==START=="
echo "git reset --hard"
git reset --hard HEAD~10
echo "git pull"
git pull
echo "go mod tidy"
go mod tidy
echo "go build ."
go build .
echo "sudo service risqlac-api restart"
sudo service risqlac-api restart
echo "==DONE=="
