## Nomad Job Runner

### Build image and push to the registry
docker build -t registry.service.simedarby.demo/nomad-jobs/nomad-jobs-runner:latest .
docker push registry.service.simedarby.demo/nomad-jobs/nomad-jobs-runner:latest
docker image rmi registry.service.simedarby.demo/nomad-jobs/nomad-jobs-runner:latest

### Run image locally
docker run -d -p 8000:8000 --env NOMAD_TOKEN=24391954-3d95-b29c-f1ff-1ea40908bec2 --env NOMAD_ADDR=http://10.35.35.10:4646 --name nomad-jobs-runner registry.service.simedarby.demo/nomad-jobs/nomad-jobs-runner

### Mount app files
/mnt/cephfs.... to /app/files

### Connection with MySQL database
Create .env file
Add string with your parameters
```MYSQL_URL=admin:password@tcp(127.0.0.1:3306)/nomad_jobs```

### Getting Started

Download ZIP from Github  
`go run main.go`

### Validate Nomad job
`http://localhost:8000/validate`

### Licensing
API phones is [MIT licensed](./LICENSE).