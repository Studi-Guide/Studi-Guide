source env.sh
export GRAPHHOPPER_API_KEY="$(cat ../../Deployment-Secrets/GRAPHHOPPER_API_KEY)"
export ASSET_STORAGE="$(cat ../../Deployment-Secrets/ASSET_STORAGE)"
go run ./cmd/studi-guide
