import json
import os
import time

from classifier import generator
from classifier.data import from_csv

_data_dir = os.path.join(os.getcwd(), "..", "data")


def classify():
    samples = os.path.join(_data_dir, "samples.csv")
    combined = os.path.join(_data_dir, "ts-large.csv")

    dataset = from_csv(combined)
    samples = from_csv(samples)

    stats = []
    evolution = {}

    for i in range(3):
        stats.append(
            _get_stats_from_execution(
                "sequential",
                generator.sequential,
                dataset,
                samples,
                i,
            )
        )
        for p in [2, 4, 8, 16]:
            stats.append(
                _get_stats_from_execution(
                    "parallel_" + str(p),
                    generator.parallel(p),
                    dataset,
                    samples,
                    i,
                )
            )
            evolution[p] = generator.display_timeline()

    output_stats = "stats_" + time.strftime("%Y%m%d-%H%M%S") + ".json"
    output_evolution = "evolution_" + time.strftime("%Y%m%d-%H%M%S") + ".json"

    with open(output_stats, "w", encoding="utf-8") as f:
        json.dump(stats, f, ensure_ascii=False, indent=4)
    with open(output_evolution, "w", encoding="utf-8") as f:
        json.dump(evolution, f, ensure_ascii=False, indent=4)


def _get_stats_from_execution(method, f, dataset, samples, i):
    t = time.time()
    tree = f(dataset)
    elapsed_time = time.time() - t

    hits, failures = tree.test(samples)

    stat = {
        "method": method,
        "execution": i + 1,
        "elapsed": elapsed_to_str(elapsed_time),
        "hits": hits,
        "failures": failures,
    }
    return stat


def elapsed_to_str(seconds):
    hours = 0
    minutes = 0

    if seconds > 59:
        hours = int(seconds / int(60 * 60))
        seconds = int(seconds % int(60 * 60))

        minutes = int(seconds / int(60))
        seconds = int(seconds % int(60))

    return "%02d:%02d:%02d" % (hours, minutes, seconds)
