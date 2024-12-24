
# Encore Go Benchmarks

This repository contains a comprehensive set of benchmarks and performance tests for Encore, Go, TypeScript, and other technologies used in modern backend development. It is designed to help developers assess and compare the performance of their backend services in terms of response time, latency, and requests per second (RPS). The benchmarks focus on several implementations, including Encore, Go-kit, and Encore-TS, providing valuable insights for optimizing backend architectures and improving scalability.

## Features

- **Benchmarking Encore with Go and TypeScript**: Includes performance tests for both Encore-Go and Encore-TS implementations.
- **Request Performance Metrics**: RPS (Requests per Second), latency, and response time analysis for different backend configurations.
- **Cold Start Performance**: Tests to analyze how quickly services start and handle traffic after deployment.
- **Visualizations**: Comprehensive charts such as response time box plots, latency percentiles, and RPS percentiles to visualize benchmarking results.
- **Easy Integration**: Setup scripts and configuration files that allow easy integration with existing Encore projects and other backend technologies.

## How to Run the Benchmarks

1. Clone the repository:

   ```bash
   git clone https://github.com/ImBIOS/encore-go-benchmarks.git
   cd encore-go-benchmarks
   ```

2. Install dependencies:

   ```bash
   pipenv install
   ```

3. Run the benchmark tests:

   ```bash
   pipenv run ./benchmark.sh
   ```

## Benchmark Types

- **Go**: Performance tests for the Go-based implementation using Encore and Go-kit.
- **TypeScript (Encore-TS)**: Tests for the TypeScript implementation using Encore and other backend frameworks.
- **Cold Starts**: Evaluating the startup time of Encore services under different conditions.
- **Request Handling**: Detailed benchmarking of how various services handle concurrent requests, focusing on response times and throughput.

## Contributing

Contributions are welcome! If you find any issues or would like to improve the benchmarks, feel free to fork the repository and submit a pull request.

## License

See the [LICENSE](LICENSE) file for details.
