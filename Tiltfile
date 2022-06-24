local_resource(
  'build', './hack/build.sh')

docker_build(
  'cloud-tilt-dev', '.')

k8s_yaml(
  helm('./chart',
       name='cloud-tilt-dev',
       set=[
         'imageName=cloud-tilt-dev',
         'numReplicas=1',
       ]))

k8s_resource(
  'snapshot-frontend',
  port_forwards='10450:10450')
