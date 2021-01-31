docker run -d \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -e DRONE_RPC_PROTO=http \
  -e DRONE_RPC_HOST=116.63.129.162:3000 \
  -e DRONE_RPC_SECRET=ed231c82ff72a158a6390c8b6a6c5adc \
  -e DRONE_RUNNER_CAPACITY=2 \
  -e DRONE_RUNNER_NAME=living \
  -p 3000:3000 \
  --restart always \
  --name runner \
  drone/drone-runner-docker:1
