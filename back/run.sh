source env.sh
echo -n "Type in your graph hopper API Key: "
read api_key
export GRAPHHOPPER_API_KEY="$api_key"
go run ./cmd/studi-guide
