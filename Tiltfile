# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')

default_registry(
  'localhost:5005',
  host_from_cluster='ctlptl-registry:5005'
)

compile_bakery = 'cd bakery && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/server main.go'

local_resource(
  'bakery-compile',
  compile_bakery,
  deps=['bakery/bakery', 'bakery/proto', 'bakery/main.go']
  )

docker_build_with_restart(
  'bakery-service:v1',
  'bakery', # build context
  entrypoint='/app/build/server',
  dockerfile='bakery/Dockerfile',
  live_update=[
    sync('bakery/', '/app/')
  ],
)

k8s_yaml(['deployment/bakery-deployment.yaml'])
k8s_resource('bakery', port_forwards=9090,
             resource_deps=['bakery-compile'])