import sys
from multiprocessing import Pool

from classifier.gain import gain


def choose_attribute_async(processes):
    pool = Pool(processes=processes)

    def async_func(dataset, attributes, class_attr):
        best = None
        value = sys.float_info.min

        gains = {}

        for attr in attributes:
            if attr == class_attr:
                continue

            gains[attr] = pool.apply_async(
                gain,
                (
                    dataset,
                    attr,
                    class_attr,
                ),
            )

        for attr in gains:
            v = gains[attr].get()
            if v > value:
                value = v
                best = attr

        if best is None:
            return None, []

        return best, _unique_values(dataset, best)

    return async_func


def _unique_values(dataset, attr):
    values = set()
    for sample in dataset:
        values.add(getattr(sample, attr))
    return values
