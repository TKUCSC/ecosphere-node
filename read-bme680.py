#!/bin/bash python
import bme680

# configuration settings
try:
	sensor = bme680.BME680(bme680.I2C_ADDR_PRIMARY)
except:
	print("err")
	exit(1)

sensor.set_humidity_oversample(bme680.OS_2X)
sensor.set_pressure_oversample(bme680.OS_4X)
sensor.set_temperature_oversample(bme680.OS_8X)

# read data from sensor
try:
	if false == sensor.get_sensor_data():
		print("err")
		exit(1)
except:
	print("err")
	exit(1)

# store and print data
humidity  = sensor.data.humidity
pressure  = sensor.data.pressure
ambient_temperature = sensor.data.temperature

print(humidity, pressure, ambient_temperature, sep="-")
exit(0)
