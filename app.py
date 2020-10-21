import logging
import pprint
import json
import sqlite3
from flask import Flask
from flask import request
from pprint import pformat

app = Flask(__name__)
logging.basicConfig(level=logging.DEBUG)

@app.route('/planets/<id>')
def one_planet(id):
    conn = sqlite3.connect('planets.db')
    conn.row_factory = sqlite3.Row

    c = conn.cursor()
    t = (id,)
    c.execute("select * from planets where position=?", t)
    row = c.fetchone()

    planet = {}
    for k in row.keys():
        planet[k] = row[k]

    return json.dumps(planet), 200, {'Content-Type': 'application/json'}

@app.route('/planets')
def list_planets():
    conn = sqlite3.connect('planets.db')
    conn.row_factory = sqlite3.Row

    c = conn.cursor()
    planets = []

    for row in c.execute("select * from planets"):
        planet = {}
        for k in row.keys():
            planet[k] = row[k]

        planets.append(planet)

    return json.dumps(planets), 200, {'Content-Type': 'application/json'}

@app.route('/webhook', methods=['GET','POST'])
def webhook():
    logging.info("Request received")
    logging.info(request.args)
    logging.info(request.form)
    return "OK"
