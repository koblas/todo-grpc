go_common = ['./pkg', './gen', './cmd/compose/shared_config', './go.sum', './go.mod', './tilt-scripts']

def go_docker(name, path):
  docker_build('todo-grpc/' + name, './src/go', 
    dockerfile='src/go/cmd/compose/'+path+'/Dockerfile.tilt',
    only=go_common + [
      './cmd/compose/'+path,
      './services/'+path,
    ],
    live_update=[
      sync('./src/go', '/build'),
      run('./tilt-scripts/restart.sh'),
    ],
  )

# The front end client
docker_build('todo-grpc/ts-client-react', './ts-client-react', 
  dockerfile='./ts-client-react/Dockerfile.tilt',
  live_update=[
    sync('./ts-client-react', '/app'),
  ],
)
k8s_yaml(['./infra/client-deployment.yaml', './infra/client-service.yaml'])
k8s_resource('client', labels=['frontend'])

# All of the GO backend microservices
go_docker('publicapi-auth', 'publicapi/auth')
k8s_yaml(['./infra/publicapi-auth-deployment.yaml', './infra/publicapi-auth-service.yaml'])
k8s_resource('publicapi-auth', labels=['public'])

go_docker('publicapi-todo', 'publicapi/todo')
k8s_yaml(['./infra/publicapi-todo-deployment.yaml', './infra/publicapi-todo-service.yaml'])
k8s_resource('publicapi-todo', labels=['public'])

go_docker('publicapi-file', 'publicapi/file')
k8s_yaml(['./infra/publicapi-file-deployment.yaml', './infra/publicapi-file-service.yaml'])
k8s_resource('publicapi-file', labels=['public'])

go_docker('publicapi-fileput', 'publicapi/fileput')
k8s_yaml(['./infra/publicapi-fileput-deployment.yaml', './infra/publicapi-fileput-service.yaml'])
k8s_resource('publicapi-fileput', labels=['public'])

go_docker('publicapi-user', 'publicapi/user')
k8s_yaml(['./infra/publicapi-user-deployment.yaml', './infra/publicapi-user-service.yaml'])
k8s_resource('publicapi-user', labels=['public'])

go_docker('publicapi-websocket', 'publicapi/websocket')
k8s_yaml(['./infra/publicapi-websocket-deployment.yaml', './infra/publicapi-websocket-service.yaml'])
k8s_resource('publicapi-websocket', labels=['public'])

go_docker('core-oauth-user', 'core/oauth_user')
k8s_yaml(['./infra/core-oauth-user-deployment.yaml', './infra/core-oauth-user-service.yaml'])
k8s_resource('core-oauth-user', labels=['backend'])

go_docker('core-send-email', 'core/send_email')
k8s_yaml(['./infra/core-send-email-deployment.yaml'])
k8s_resource('core-send-email', labels=['backend'])

go_docker('core-todo', 'core/todo')
k8s_yaml(['./infra/core-todo-deployment.yaml', './infra/core-todo-service.yaml'])
k8s_resource('core-todo', labels=['backend'])

go_docker('core-user', 'core/user')
k8s_yaml(['./infra/core-user-deployment.yaml', './infra/core-user-service.yaml'])
k8s_resource('core-user', labels=['backend'])

go_docker('core-file', 'core/file')
k8s_yaml(['./infra/core-file-deployment.yaml', './infra/core-file-service.yaml'])
k8s_resource('core-file', labels=['backend'])

go_docker('websocket-todo', 'websocket/todo')
k8s_yaml(['./infra/core-websocket-todo-deployment.yaml'])
k8s_resource('websocket-todo', labels=['event'])

go_docker('websocket-user', 'websocket/user')
k8s_yaml(['./infra/core-websocket-user-deployment.yaml'])
k8s_resource('websocket-user', labels=['event'])

go_docker('workers-user', 'workers/workers_user')
k8s_yaml(['./infra/workers-user-deployment.yaml'])
k8s_resource('workers-user', labels=['event'])

go_docker('workers-file', 'workers/workers_file')
k8s_yaml(['./infra/workers-file-deployment.yaml'])
k8s_resource('workers-file', labels=['event'])

# Infrastructure
k8s_yaml(['./infra/ingress.yaml'])
k8s_yaml(['./infra/dynamodb-deployment.yaml', './infra/dynamodb-service.yaml'])
k8s_resource('dynamodb', labels=['infrastructure'])
k8s_yaml(['./infra/redis-deployment.yaml', './infra/redis-service.yaml'])
k8s_resource('redis', labels=['infrastructure'])
