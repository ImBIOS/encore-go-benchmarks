#!/bin/bash

# Paths
ENCORE_DIR="requests/encore"
ENCORE_TS_DIR="requests/encore-ts"
PYTHON_SCRIPT="./visualize_benchmark.py"

# Benchmark parameters
CONCURRENT_USERS=150
DURATION="3s"

# Function to extract the API server port from Encore logs
get_api_server_port() {
  local log_output=$1
  echo "$log_output" | grep "api server listening for incoming requests" | grep -oP "127\.0\.0\.1:\K[0-9]+"
}

# Function to run benchmark
run_benchmark() {
  local dir=$1
  local name=$2
  echo "Starting benchmark for $name..."

  # Navigate to the directory
  cd "$dir" || exit 1

  # Start the Encore server and capture logs
  LOG_OUTPUT=$(ENCORE_LOG=off ENCORE_NOTRACE=1 ENCORE_RUNTIME_LOG=debug encore run 2>&1) &
  SERVER_PID=$!
  echo "Server started (PID: $SERVER_PID)."

  # Extract the API server port from the logs
  API_SERVER_PORT=$(get_api_server_port "$LOG_OUTPUT" || echo "4000")
  if [ -z "$API_SERVER_PORT" ]; then
    echo "Failed to extract API server port from logs. Aborting."
    kill "$SERVER_PID"
    wait "$SERVER_PID" 2>/dev/null
    exit 1
  fi
  BENCHMARK_URL="http://127.0.0.1:$API_SERVER_PORT/hello"
  echo "Benchmarking against URL: $BENCHMARK_URL"

  # Wait for server to be ready
  sleep 5

  # Run oha benchmark
  echo "Running benchmark for $name..."
  oha -c "$CONCURRENT_USERS" -z "$DURATION" -m GET "$BENCHMARK_URL" --json --latency-correction --disable-keepalive >benchmark.json

  # Stop the server
  echo "Stopping server for $name..."
  kill "$SERVER_PID"
  wait "$SERVER_PID" 2>/dev/null
  echo "Benchmark for $name completed."

  # Navigate back
  cd - || exit 1
}

# Run benchmarks for encore and encore-ts
run_benchmark "$ENCORE_DIR" "Encore-Go"
run_benchmark "$ENCORE_TS_DIR" "Encore-TS"

# Visualize benchmark results
echo "Visualizing results..."
python3 visualize_benchmark.py \
  "$ENCORE_DIR/benchmark.json" \
  "$ENCORE_TS_DIR/benchmark.json" \
  "Encore-Go" \
  "Encore-TS"

echo "All benchmarks completed."
