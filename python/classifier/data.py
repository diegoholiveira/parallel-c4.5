import csv


class Sample:
    torque: float
    torque1: float
    torque10: float
    torque100: float
    pcut_speed: float
    psvol_speed: float
    vax_speed: float
    vax_speed1: float
    vax_speed10: float
    vax_speed100: float
    mode: int
    status: str
    lag_error1: float
    lag_error10: float
    lag_error100: float


def _parse_row(row):
    sample = Sample()
    sample.torque = float(row["pCut::Motor_Torque"])
    sample.torque1 = float(row["pCut::Motor_Torque1"])
    sample.torque10 = float(row["pCut::Motor_Torque10"])
    sample.torque100 = float(row["pCut::Motor_Torque100"])
    sample.pcut_speed = float(row["pCut::CTRL_Position_controller::Actual_speed"])
    sample.psvol_speed = float(row["pSvolFilm::CTRL_Position_controller::Actual_speed"])
    sample.vax_speed = float(row["pSpintor::VAX_speed"])
    sample.vax_speed1 = float(row["pSpintor::VAX_speed1"])
    sample.vax_speed10 = float(row["pSpintor::VAX_speed10"])
    sample.vax_speed100 = float(row["pSpintor::VAX_speed100"])
    sample.lag_error1 = float(row["pCut::CTRL_Position_controller::Lag_error1"])
    sample.lag_error10 = float(row["pCut::CTRL_Position_controller::Lag_error10"])
    sample.lag_error100 = float(row["pCut::CTRL_Position_controller::Lag_error100"])
    sample.mode = int(row["Mode"])
    sample.status = row["Status"]

    return sample


def from_csv(filename):
    with open(filename, newline="") as csvfile:
        reader = csv.DictReader(csvfile)
        return [_parse_row(row) for row in reader]
