#!/usr/bin/env bash

# build views into binary and then deploy
echo "===== Generating assets file ======="
# go-assets-builder views -o assets.go

env GOOS=linux GOARCH=amd64 go build -o noredd-app

rsync -azP public/ root@104.248.255.136:/home/noredd-app/public/
rsync -azP assets/ root@104.248.255.136:/home/noredd-app/assets/
rsync -azP templates/ root@104.248.255.136:/home/noredd-app/templates/


#rsync -azP views/ root@homefbase:/home/apps/homef/views/
#rsync -azP vendor/ root@homefbase:/home/apps/homef/vendor/

# ssh -l root homefbase "systemctl stop homef.service; systemctl status homef.service; rm /home/apps/homef/homef-gin"
scp noredd-app root@104.248.255.136:/home/noredd-app/

# ssh -l root homefbase "systemctl start homef.service; systemctl status homef.service;"

echo "Cleaning Up"
rm noredd-app

echo "Finshed build/deploy"
