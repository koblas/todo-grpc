docker_compose("./docker-compose.yml")

#
#  Common Go projects
#
#  We're using a custom Dockerfile since the staged dockerfile doesn't have
#  the go development environment
#
go_common = ['./pkg', './genpb', './go.sum', './go.mod']

docker_build('public-auth', 'src/go', 
  dockerfile='./src/go/cmd/auth/Dockerfile.tilt',
  only=go_common.extend([
    './cmd/auth',
    './services/api/auth',
  ]),
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('public-todo', 'src/go', 
  dockerfile='./src/go/cmd/todo/Dockerfile.tilt',
  only=go_common.extend([
    './cmd/todo',
    './services/api/todo',
  ]),
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)


docker_build('api-extauth', 'src/go', 
  dockerfile='./src/go/cmd/extauth/Dockerfile.tilt',
  only=go_common.extend([
    './cmd/extauth',
    './services/api/extauth',
  ]),
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('core-user', 'src/go', 
  dockerfile='./src/go/cmd/user/Dockerfile.tilt',
  only=go_common.extend([
    './cmd/user',
    './services/core/user',
  ]),
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

#
#
#
docker_build('ts-client-react', 'ts-client-react', 
  dockerfile='./ts-client-react/Dockerfile.tilt',
  live_update=[
    sync('./ts-client-react', '/app'),
  ],
)

