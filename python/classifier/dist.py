from collections import Counter


class Dist(object):
    def __init__(self, attr, dataset=[]):
        self.counter = Counter()
        self.total = 0
        self.attr = attr

        for sample in dataset:
            self.add(getattr(sample, self.attr))

    def add(self, sample, counter=1):
        self.counter[sample] += counter
        self.total += counter

    def best(self):
        commons = self.counter.most_common(1)

        if not commons:
            return ""

        return commons[0][0]
