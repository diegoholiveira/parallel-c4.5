import sys

from classifier.gain import gain


def choose_attribute(dataset, attributes, class_attr):
    best = None
    value = sys.float_info.min

    for attr in attributes:
        if attr == class_attr:
            continue

        v = gain(dataset, attr, class_attr)
        if v > value:
            value = v
            best = attr

    if best is None:
        return None, []

    return best, _unique_values(dataset, best)


def _unique_values(dataset, attr):
    values = set()
    for sample in dataset:
        values.add(getattr(sample, attr))
    return values
