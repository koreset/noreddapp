#!/usr/bin/env bash

# build views into binary and then deploy
echo "===== Generating assets file ======="
# go-assets-builder views -o assets.go

env GOOS=linux GOARCH=amd64 go build -tags 'bindatafs' -o noredd-app

rsync -azP public/ root@104.248.255.136:/home/noredd-app/public/
rsync -azP assets/ root@104.248.255.136:/home/noredd-app/assets/
rsync -azP templates/ root@104.248.255.136:/home/noredd-app/templates/

ssh -l root 104.248.255.136 "systemctl stop noredd-app.service; systemctl status noredd-app.service; rm /home/noredd-app/noredd-app"
scp noredd-app root@104.248.255.136:/home/noredd-app/

ssh -l root 104.248.255.136 "systemctl start noredd-app.service; systemctl status noredd-app.service;"

echo "Cleaning Up"
rm noredd-app

echo "Finshed build/deploy"
