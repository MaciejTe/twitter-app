# Twitter API
Twitter is simple Go API which sends the messages and query for messages
containing a specific tag.
## Choices
1. Fiber - REST API framework, chosen beacuse of good performance benchmarks
2. MongoDB - document database, perfect for this case
3. Python/pytest - for integration/E2E tests
4. Kubernetes - perfect to deploy and scale application
5. Skaffold/Tilt - nice k8s development tools

## Endpoints

**POST /messages**

Example: `http://localhost:3000/messages`

JSON Schema:
1. `text` - message to be posted
2. `tags` - tags to be assigned to message

Sample JSON payload:
```
{
    "text": "sample message",
    "tags": ["tag1", "tag2"]
}
```

**GET /messages**

Example: `http://localhost:3000/messages?from=2020-06-13T15:00:05Z&to=2021-06-13T15:00:05Z&tags=zx,zs`

Query parameters:
1. `tags` - list of tags, comma separated
2. `from` - RFC3339-compatible start filtering time
3. `to` - RFC3339-compatible stop filtering time

## Development

1. Install [Skaffold](https://skaffold.dev/docs/install/)
2. Install [kubectl](https://kubernetes.io/docs/tasks/tools/)
3. Install [minikube](https://minikube.sigs.k8s.io/docs/start/)


All needed dependencies and quickstart [here](https://skaffold.dev/docs/quickstart/).

4. `skaffold dev`
5. Use uploaded [Postman](https://www.postman.com/) collection (`Twitter.postman_collection.json`) to call API

### Linting and testing
Prerequisites: use configuration from [Configuration](#Configuration) section.

1. Unit tests
Type following commands to execute unit tests in Twitter development container:
   - `make dev`
   - `make test`

2. Linting
Folowing command installs necessary test dependencies and launches linters(golangci-lint, golint, goimports, gofmt) and checkers (gocyclo, go vet):
    - `make dev`
    - `make lint`

3. E2E tests
To launch E2E tests enter test Docker container (assumption: Twitter app, database and Test container are in the same network):
    - `make dev` - Enter development container
    - `make build; ./twitter_app` - in dev container
    - `make dev-tests` - to enter test container
    - `pytest` - to launch tests

4. Coverage report
In order to generate coverage report, type following commands:
    - `make dev` - enter development container
    - `make coverage` - generate global coverage report
    - `make coverhtml` - generate HTML coverage report

## Configuration
Both Twitter API and Test container are configured using environment variables(it's convenient to have `.env` file in main application directory):

```
API_PORT=3000
DB_URI=mongodb://root:root@localhost:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false
DB_NAME=twitter
DB_COLLECTION_NAME=messages
```

## Deployment
Kubernetes is used to deploy solution:

1. Build prod Docker image: `make build_image`
2. Apply deployment yaml: `kubectl apply -f deployment/twitter-deployment.yaml`
3. Check your minikube IP: `minikube ip`
4. You can access Twitter API on minikube IP address, port `32000`, like `http://192.168.49.2:32000`

## Summary

### Shortcuts:
1. No authentication
2. No performance tests (Python and locust could be used to measure that)
3. No security tests
4. No monitoring solution (ELK/Prometheus)
5. No CI (private repository issues)
6. No `api/v1` used in endpoints - all in all, it's JSON API ;-)
7. No helm and other k8s improvements
8. No automated API docs like Swagger
