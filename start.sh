

# Run Go server-side Application
echo "Running Storage Service..."
go run storage/main.go & P1=$!
wait $P1