APP_NAME="risqlac-api"
PORT=3000

echo "==> DEPLOY START <=="

echo "==> git pull"
git pull

echo "----- CLEAN UP START -----"
chmod 777 clean.sh && ./clean.sh
echo "----- CLEAN UP DONE -----"

echo "==> build image"
podman build -t $APP_NAME .

echo "==> run image"
podman run -d --name $APP_NAME -p $PORT:3000 $APP_NAME

echo "==> DEPLOY DONE <=="
