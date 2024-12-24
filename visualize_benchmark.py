import json
import sys

import matplotlib.pyplot as plt
import matplotlib.ticker as ticker
import pandas as pd
import seaborn as sns

# Get arguments from command line
benchmark1_file = sys.argv[1]
benchmark2_file = sys.argv[2]
label1 = sys.argv[3]
label2 = sys.argv[4]

# Load JSON data
with open(benchmark1_file, "r") as f:
    data1 = json.load(f)

with open(benchmark2_file, "r") as f:
    data2 = json.load(f)


# Function to plot response time box plot
def plot_response_time_boxplot(histogram1, histogram2):
    # Prepare DataFrame for comparison
    hist_df1 = pd.DataFrame(list(histogram1.items()), columns=["Response Time", label1])
    hist_df2 = pd.DataFrame(list(histogram2.items()), columns=["Response Time", label2])

    # Convert response times to float and scale them (e.g., milliseconds)
    hist_df1["Response Time"] = (
        hist_df1["Response Time"].astype(float) * 1000
    )  # Convert to ms
    hist_df2["Response Time"] = (
        hist_df2["Response Time"].astype(float) * 1000
    )  # Convert to ms

    # Combine both dataframes into one for plotting
    combined_df = pd.concat([hist_df1, hist_df2], axis=1)

    # Check the number of columns in the combined DataFrame
    print(f"Columns in combined_df: {combined_df.columns.tolist()}")  # Debugging line

    # Only keep the necessary columns
    combined_df = combined_df[[label1, label2]]  # Select the relevant columns

    # Plot the box plot
    plt.figure(figsize=(12, 6))
    sns.boxplot(data=combined_df[[label1, label2]], orient="h", palette="Set2")
    plt.xlabel("Response Time (ms)")
    plt.ylabel("Benchmark")
    plt.title("Response Time Box Plot Comparison (Lower is Better)")
    plt.savefig("requests-results/response_time_boxplot_comparison.png")
    plt.close()


# Function to plot latency percentiles
def plot_latency_percentiles(percentiles1, percentiles2):
    # Prepare DataFrame for comparison
    perc_df1 = pd.DataFrame(list(percentiles1.items()), columns=["Percentile", label1])
    perc_df2 = pd.DataFrame(list(percentiles2.items()), columns=["Percentile", label2])
    combined_df = pd.merge(perc_df1, perc_df2, on="Percentile", how="outer").fillna(0)
    combined_df["Percentile"] = pd.Categorical(
        combined_df["Percentile"],
        categories=[
            "p10",
            "p25",
            "p50",
            "p75",
            "p90",
            "p95",
            "p99",
            "p99.9",
            "p99.99",
        ],
        ordered=True,
    )

    # Plot the latency percentiles
    plt.figure(figsize=(12, 6))
    sns.lineplot(
        x="Percentile",
        y=label1,
        data=combined_df,
        marker="o",
        label=label1,
        color="blue",
    )
    sns.lineplot(
        x="Percentile",
        y=label2,
        data=combined_df,
        marker="o",
        label=label2,
        color="orange",
    )
    plt.xlabel("Percentile")
    plt.ylabel("Latency (s)")
    plt.title("Latency Percentiles Comparison (Lower is Better)")
    plt.legend()
    plt.savefig("requests-results/latency_percentiles_comparison.png")
    plt.close()


# Function to plot requests per second (RPS) percentiles
def plot_rps_percentiles(rps_percentiles1, rps_percentiles2):
    # Prepare DataFrame for comparison
    rps_df1 = pd.DataFrame(
        list(rps_percentiles1.items()), columns=["Percentile", label1]
    )
    rps_df2 = pd.DataFrame(
        list(rps_percentiles2.items()), columns=["Percentile", label2]
    )
    combined_df = pd.merge(rps_df1, rps_df2, on="Percentile", how="outer").fillna(0)
    combined_df["Percentile"] = pd.Categorical(
        combined_df["Percentile"],
        categories=[
            "p10",
            "p25",
            "p50",
            "p75",
            "p90",
            "p95",
            "p99",
            "p99.9",
            "p99.99",
        ],
        ordered=True,
    )

    # Plot the RPS percentiles
    plt.figure(figsize=(12, 6))
    sns.lineplot(
        x="Percentile",
        y=label1,
        data=combined_df,
        marker="o",
        label=label1,
        color="blue",
    )
    sns.lineplot(
        x="Percentile",
        y=label2,
        data=combined_df,
        marker="o",
        label=label2,
        color="orange",
    )
    plt.xlabel("Percentile")
    plt.ylabel("Requests Per Second")
    plt.title("RPS Percentiles Comparison (Higher is Better)")
    plt.legend()
    plt.savefig("requests-results/rps_percentiles_comparison.png")
    plt.close()


# Visualizations
plot_response_time_boxplot(
    data1["responseTimeHistogram"], data2["responseTimeHistogram"]
)
plot_latency_percentiles(data1["latencyPercentiles"], data2["latencyPercentiles"])
plot_rps_percentiles(data1["rps"]["percentiles"], data2["rps"]["percentiles"])
