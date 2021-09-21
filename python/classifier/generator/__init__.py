from .c45 import c45
from .parallel import choose_attribute_async
from .sequential import choose_attribute

_attributes = [
    "torque",
    "pcut_speed",
    "psvol_speed",
    "vax_speed",
    "mode",
    "status",
]

_class_attr = "status"


def sequential(dataset):
    return c45(choose_attribute, dataset, _attributes, _class_attr)


def parallel(processes):
    f = choose_attribute_async(processes)

    def wrapper(dataset):
        return c45(
            f,
            dataset,
            _attributes,
            _class_attr,
        )

    return wrapper
