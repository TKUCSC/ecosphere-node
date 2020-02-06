import bme280
import smbus2

# configuration settings
port = 1
address = 0x77

try:
    bus = smbus2.SMBus(port)
    calibration_params = bme280.load_calibration_params(bus,address)
except:
    print("err: config error")
    exit(1)

# read data from sensor
try:
    bme280_data = bme280.sample(bus,address, calibration_params)
except:
    print("err: read error")
    exit(1)

# store and print data
humidity  = bme280_data.humidity
pressure  = bme280_data.pressure
ambient_temperature = bme280_data.temperature

print(humidity, pressure, ambient_temperature)
exit(0)
