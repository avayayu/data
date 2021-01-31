docker run \
  --volume=/var/lib/drone:/data \
  --env=DRONE_GITHUB_CLIENT_ID=437c14237c823c91d916 \
  --env=DRONE_GITHUB_CLIENT_SECRET=6d3a0b8242d61ab860ec07a622d238104ba173ff  \
  --env=DRONE_RPC_SECRET=ed231c82ff72a158a6390c8b6a6c5adc \
  --env=DRONE_SERVER_HOST=116.63.129.162:3000 \
  --env=DRONE_SERVER_PROTO=http \
  --publish=3000:80 \
  --publish=4430:443 \
  --restart=always \
  --detach=true \
  --name=drone \
  drone/drone:1

docker run \
  --volume=/var/lib/drone:/data \
  --env=DRONE_GITHUB_CLIENT_ID=437c14237c823c91d916 \
  --env=DRONE_GITHUB_CLIENT_SECRET=6d3a0b8242d61ab860ec07a622d238104ba173ff  \
  --env=DRONE_RPC_SECRET=ed231c82ff72a158a6390c8b6a6c5adc \
  --env=DRONE_SERVER_HOST=172.27.232.73:3000 \
  --env=DRONE_SERVER_PROTO=http \
  --publish=3000:80 \
  --publish=4430:443 \
  --restart=always \
  --detach=true \
  --name=drone \
  drone/drone:1
