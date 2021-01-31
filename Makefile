PYGRPC=python -m grpc_tools.protoc


basic:
	${PYGRPC} -I ./api/basic/ --python_out=./crawlerpy/web/ --grpc_python_out=./crawlerpy/web/ basic.proto 


protoc  -I . --go_out=plugins=grpc,paths=source_relative:.  ./api/base.proto


docker run -d \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -e DRONE_RPC_PROTO=https \
  -e DRONE_RPC_HOST=116.63.129.162:3000 \
  -e DRONE_RPC_SECRET=super-duper-secret \
  -e DRONE_RUNNER_CAPACITY=2 \
  -e DRONE_RUNNER_NAME=living \
  -p 3000:3000 \
  --restart always \
  --name runner \
  drone/drone-runner-docker:1
