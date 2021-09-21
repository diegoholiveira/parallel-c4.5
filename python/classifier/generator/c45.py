from classifier.dist import Dist
from classifier.node import Node


def c45(choose_attribute_func, dataset, attributes, class_attr):
    best, values = choose_attribute_func(dataset, attributes, class_attr)

    if best is None:
        return None

    node = Node(best, len(dataset))

    for value in values:
        sub_dataset = create_subdataset(dataset, best, value)

        dist = Dist(class_attr, sub_dataset)

        if is_leaf(sub_dataset, attributes, dist):
            node.set_leaf(value, dist)

            continue  # go to the next value

        sub_attributes = [attr for attr in attributes if attr != best]

        child = c45(choose_attribute_func, sub_dataset, sub_attributes, class_attr)
        if child:
            node.create_branch(value, child)

    return node


def create_subdataset(dataset, best, value):
    subdataset = []
    for sample in dataset:
        if getattr(sample, best) == value:
            subdataset.append(sample)
    return subdataset


def is_leaf(subdataset, attributes, dist):
    return not subdataset or (len(attributes) - 1) <= 0 or len(dist.counter) <= 1
