docker_compose("./docker-compose.yml")

#
#  Common Go projects
#
#  We're using a custom Dockerfile since the staged dockerfile doesn't have
#  the go development environment
#
go_common = ['./pkg', './twpb', './services', './cmd', './go.sum', './go.mod', './tilt-scripts']

docker_build('public-auth', 'src/go', 
  dockerfile='./src/go/cmd/compose/publicapi/auth/Dockerfile.tilt',
  only=go_common + [
    './cmd/compose/publicapi/auth',
    './services/api/auth',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('public-todo', 'src/go', 
  dockerfile='./src/go/cmd/compose/publicapi/todo/Dockerfile.tilt',
  only=go_common + [
    './cmd/compose/publicapi/todo',
    './services/api/todo',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('public-websocket', 'src/go', 
  dockerfile='./src/go/cmd/compose/publicapi/websocket/Dockerfile.tilt',
  only=go_common + [
    './cmd/compose/publicapi/websocket',
    './services/websocket/todo',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)


docker_build('core-user', 'src/go', 
  dockerfile='./src/go/cmd/compose/core/user/Dockerfile.tilt',
  only=go_common + [
    './cmd/compose/core/user',
    './services/core/user',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('core-todo', 'src/go', 
  dockerfile='./src/go/cmd/compose/core/todo/Dockerfile.tilt',
  only=go_common + [
    './cmd/compose/core/todo',
    './services/core/todo',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('core-oauth-user', 'src/go', 
  dockerfile='./src/go/cmd/compose/core/oauth_user/Dockerfile.tilt',
  only=go_common + [
    './cmd/compose/core/oauth_user',
    './services/core/oauth_user',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('core-send-email', 'src/go', 
  dockerfile='./src/go/cmd/compose/core/send_email/Dockerfile.tilt',
  only=go_common + [
    './cmd/compose/core/send_email',
    './services/core/send_email',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('core-workers', 'src/go',
  dockerfile='./src/go/cmd/compose/core/workers/Dockerfile.tilt',
  only=go_common + [
    './cmd/compose/core/workers',
    './services/core/workers',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('websocket-todo', 'src/go',
  dockerfile='./src/go/cmd/compose/websocket/todo/Dockerfile.tilt',
  only=go_common + [
    './cmd/compose/websocket/todo',
  ],
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

