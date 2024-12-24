# Encore Go Benchmarks

## Note Terminal 1

```bash

```

## Note Terminal 2

```bash
   main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   8s
 ➜  oha -c 50 -z 10s -m GET "http://localhost:4000/hello" --json > benchmark.json
^C%

󰀵  󱑍 07:56  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   8s
 ➜  oha -c 150 -z 10s -m GET "http://localhost:4000/hello" --json > benchmark.json
^C%

󰀵  󱑍 07:57  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   9s
 ➜  oha -c 150 -z 10s -m GET "http://localhost:4000/hello" --json > benchmark.json
^C%

󰀵  󱑍 07:57  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   5s
 ➜  oha -c 150 -z 10s -m GET "http://127.0.0.1:4000/hello" --json > benchmark.json
^C%

󰀵  󱑍 07:57  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   5s
 ➜  oha -c 150 -z 3s -m GET "http://127.0.0.1:4000/hello" --json --latency-correction --disable-keepalive > benchmark.json
^C%

󰀵  󱑍 07:58  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)
 ➜  oha -c 150 -z 3s -m GET "http://127.0.0.1:4000/hello" --json --latency-correction --disable-keepalive > benchmark.json

󰀵  󱑍 07:58  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   3s
 ➜  pipenv run python visualize_benchmark.py

󰀵  󱑍 07:59  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)
 ➜  oha -c 150 -z 10s -m GET "http://127.0.0.1:4000/hello" --json --latency-correction --disable-keepalive > benchmark.json
^C%

󰀵  󱑍 07:59  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   5s
 ➜  oha -c 150 -z 5s -m GET "http://127.0.0.1:4000/hello" --json --latency-correction --disable-keepalive > benchmark.json

󰀵  󱑍 07:59  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   5s
 ➜  oha -c 150 -z 3s -m GET "http://127.0.0.1:4000/hello" --json --latency-correction --disable-keepalive > benchmark.json

󰀵  󱑍 07:59  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   3s
 ➜  pipenv run python visualize_benchmark.py

󰀵  󱑍 07:59  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)
 ➜  oha -c 150 -z 10s -m GET "http://127.0.0.1:4000/hello" --json --latency-correction --disable-keepalive > benchmark.json
^C%

󰀵  󱑍 08:00  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   4s
 ➜  oha -c 150 -z 4s -m GET "http://127.0.0.1:4000/hello" --json --latency-correction --disable-keepalive > benchmark.json

󰀵  󱑍 08:02  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   4s
 ➜  oha -c 150 -z 3s -m GET "http://127.0.0.1:4000/hello" --json --latency-correction --disable-keepalive > benchmark.json

󰀵  󱑍 08:03  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   3s
 ➜  oha -c 150 -z 3s -m GET "http://127.0.0.1:4000/hello" --json --latency-correction --disable-keepalive > benchmark.json

󰀵  󱑍 08:03  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   3s
 ➜  oha -c 150 -z 3s -m GET "http://127.0.0.1:4000/hello" --json --latency-correction --disable-keepalive > benchmark.json

󰀵  󱑍 08:03  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)   3s
 ➜  pipenv run python visualize_benchmark.py

󰀵  󱑍 08:03  ﱮ encore-go-benchmarks/requests/encore    main ✘!?+7 via 🐹 v1.23.1 via 🐍 v3.12.3 (encore)
 ➜  oha -c 150 -z 3s -m GET "http://127.0.0.1:4000/hello" --json --latency-correction --disable-keepalive > benchmark.json
```
