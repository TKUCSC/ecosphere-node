import bme680

# configuration settings
try:
    try:
        sensor = bme680.BME680(bme680.I2C_ADDR_PRIMARY)
    except IOError:
        sensor = bme680.BME680(bme680.I2C_ADDR_SECONDARY)
except:
    print("err: config error")
    exit(1)

sensor.set_humidity_oversample(bme680.OS_2X)
sensor.set_pressure_oversample(bme680.OS_4X)
sensor.set_temperature_oversample(bme680.OS_8X)
sensor.set_filter(bme680.FILTER_SIZE_3)

# read data from sensor
try:
    if sensor.get_sensor_data():
        humidity  = sensor.data.humidity
        pressure  = sensor.data.pressure
        ambient_temperature = sensor.data.temperature
        print(humidity, pressure, ambient_temperature)
except:
    print("err: read error")
    exit(1)

exit(0)
