docker_compose("./docker-compose.yml")

#
#  Common Go projects
#
#  We're using a custom Dockerfile since the staged dockerfile doesn't have
#  the go development environment
#
go_common = ['./pkg', './genpb', './go.sum', './go.mod', './tilt-scripts']

docker_build('public-auth', 'src/go', 
  dockerfile='./src/go/cmd/publicapi/auth/Dockerfile.tilt',
  only=go_common + [
    './cmd/publicapi/auth',
    './services/api/auth',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('public-todo', 'src/go', 
  dockerfile='./src/go/cmd/publicapi/todo/Dockerfile.tilt',
  only=go_common + [
    './cmd/publicapi/todo',
    './services/api/todo',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)


docker_build('api-extauth', 'src/go', 
  dockerfile='./src/go/cmd/middleware/extauth/Dockerfile.tilt',
  only=go_common + [
    './cmd/middleware/extauth',
    './services/api/extauth',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('core-user', 'src/go', 
  dockerfile='./src/go/cmd/core/user/Dockerfile.tilt',
  only=go_common + [
    './cmd/core/user',
    './services/core/user',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('core-oauth-user', 'src/go', 
  dockerfile='./src/go/cmd/core/oauth_user/Dockerfile.tilt',
  only=go_common + [
    './cmd/core/oauth_user',
    './services/core/oauth_user',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('core-send-email', 'src/go', 
  dockerfile='./src/go/cmd/core/send_email/Dockerfile.tilt',
  only=go_common + [
    './cmd/core/send_email',
    './services/core/send_email',
  ],
  live_update=[
    sync('./src/go', '/build'),
    run('./tilt-scripts/restart.sh'),
  ],
)

docker_build('core-workers', 'src/go',
  dockerfile='./src/go/cmd/core/workers/Dockerfile.tilt',
  only=go_common + [
    './cmd/core/workers',
    './services/core/workers',
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

