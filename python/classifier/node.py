from classifier.dist import Dist


class Node(object):
    def __init__(self, attr, size):
        self.name = attr
        self.size = size
        self.branches = {}
        self.attr_values = {}
        self.total = 0
        self.attr_class_values = {}

    def test(self, dataset):
        hits = failures = 0

        for sample in dataset:
            predicted = self.predict(sample)
            if predicted is None:
                failures += 1

                continue

            value = getattr(sample, "status")
            if predicted.best() == value:
                hits += 1
            else:
                failures += 1

        return hits, failures

    def set_leaf(self, val, dist):
        if val not in self.attr_values:
            self.attr_values[val] = 0

        self.attr_values[val] += 1
        self.total += 1

        for value in dist.counter.keys():
            if val not in self.attr_class_values:
                self.attr_class_values[val] = {}

            if value not in self.attr_class_values[val]:
                self.attr_class_values[val][value] = 0

            self.attr_class_values[val][value] += dist.counter[value]

    def create_branch(self, val, subtree):
        self.branches[val] = subtree

    def get_value_dist(self, attr_value):
        dist = Dist(attr_value)
        counter = self.attr_class_values[attr_value]
        for key in counter:
            dist.add(key, counter=counter[key])
        return dist

    def get_values(self):
        values = set()
        for k in self.attr_values.keys():
            values.add(k)
        for k in self.branches.keys():
            values.add(k)
        return values

    def get_attr_value_from_sample(self, sample):
        value = getattr(sample, self.name)
        values = self.get_values()
        if value in values:
            return value

        n = 1e999999
        nearest = None
        for v in values:
            if n > abs(v - value):
                n = abs(v - value)
                nearest = v

        return nearest

    def predict(self, sample, depth=0):
        value = self.get_attr_value_from_sample(sample)
        if value is None:
            return None

        if value in self.branches:
            child = self.branches[value]
            return child.predict(sample, depth=depth + 1)

        return self.get_value_dist(value)
