import requests
import json
import subprocess

with open('config.json') as config:
	CONFIG_DATA = json.load(config)
	URL = CONFIG_DATA['TARGET']

	try:
		res = requests.post(url = URL, data = CONFIG_DATA)
	except:
		# DO NOTHING
		print("err: error sending request")

	print(CONFIG_DATA)

subprocess.run([f"python3", f"modules/read-{CONFIG_DATA['SENSOR_MODEL']}.py"])
