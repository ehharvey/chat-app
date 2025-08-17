# chat-app
A chat app!

# Development: Getting Started
## Starting Development Environment
1. Open project in VSCode dev container
2. Install dependencies (run `./install-deps.sh`)
3. Start k8s local cluter via `kind cluster create`
4. Start development environment via `skaffold dev --port-forward`. Note that ports are forwarded fromt the k8s cluster, including:
   * `8080`: This is our API
   * `5432`: This is our DB (currently unused)
   * `3000`: This is Grafana (currently unused)

## Database changes
We use `sqlc` to manage our DB schemas and queries and `goose` to manage migrations.