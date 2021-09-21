import csv


class Sample:
    torque: float
    pcut_speed: float
    psvol_speed: float
    vax_speed: float
    mode: int
    status: str


def _parse_row(row):
    sample = Sample()
    sample.torque = float(row["pCut::Motor_Torque"])
    sample.pcut_speed = float(row["pCut::CTRL_Position_controller::Actual_speed"])
    sample.psvol_speed = float(row["pSvolFilm::CTRL_Position_controller::Actual_speed"])
    sample.vax_speed = float(row["pSpintor::VAX_speed"])
    sample.mode = int(row["Mode"])
    sample.status = row["Status"]

    return sample


def from_csv(filename):
    with open(filename, newline="") as csvfile:
        reader = csv.DictReader(csvfile)
        return [_parse_row(row) for row in reader]
