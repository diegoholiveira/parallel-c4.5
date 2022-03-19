from .c45 import c45, reset, timeline
from .parallel import choose_attribute_async
from .sequential import choose_attribute

_attributes = [
    "torque",
    "torque1",
    "torque10",
    "torque100",
    "pcut_speed",
    "psvol_speed",
    "vax_speed",
    "vax_speed1",
    "vax_speed10",
    "vax_speed100",
    "mode",
    "status",
    "lag_error1",
    "lag_error10",
    "lag_error100",
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


def display_timeline():
    t = timeline()
    reset()
    return t
