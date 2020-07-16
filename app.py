import logging
import pprint
from flask import Flask
from flask import jsonify
from flask import request
from pprint import pformat


app = Flask(__name__)
logging.basicConfig(level=logging.DEBUG)

planets = [
    {"name": "Mercury", "position": 1, "moons":0},
    {"name": "Venus", "position": 2, "moons":0},
    {"name": "Earth", "position": 3, "moons":1},
    {"name": "Mars", "position": 4, "moons":2},
    {"name": "Jupiter", "position": 5, "moons":79},
    {"name": "Saturn", "position": 6, "moons":62},
    {"name": "Uranus", "position": 7, "moons":27},
    {"name": "Neptune", "position": 8, "moons":14},
]


@app.route('/planets/<id>')
def one_planet(id):
    return jsonify(planets[int(id)-1])

@app.route('/planets')
def list_planets():
    return jsonify(planets)

@app.route('/webhook', methods=['GET','POST'])
def webhook():
    logging.info("Request received")
    logging.info(request.args)
    logging.info(request.form)
    return "OK"
