from collections import Counter
from math import log


def _frequency(dataset, attr):
    counter = Counter()
    for sample in dataset:
        k = getattr(sample, attr)
        counter[k] += 1.0
    return counter


def _entropy(data, attr):
    frequency = _frequency(data, attr)
    s = float(sum(f for f in frequency.values()))
    n = max(2, len(frequency))
    return -sum((f / s) * log(f / s, n) for f in frequency.values() if f)


def gain(dataset, attr, class_attr):
    subset_entropy = 0.0
    frequency = _frequency(dataset, attr)

    for value in frequency.keys():
        subset = []
        for sample in dataset:
            if getattr(sample, attr) == value:
                subset.append(sample)

        e = _entropy(subset, class_attr)
        prob = frequency[value] / sum(frequency.values())
        subset_entropy += prob * e

    main_entropy = _entropy(dataset, class_attr)
    return main_entropy - subset_entropy
