# Sample Flask App for 2-way APIs

This application contains: one flask app in `app.py`. It serves up data about planets from [NASA data](https://solarsystem.nasa.gov/moons/in-depth/) at `/planets` and `/planets/<position>`. It also logs the incoming data it receives to a `/webhook` endpoint.

Run it like this:

```
FLASK_APP=app.py python -mflask run
```

This was created for a one-off talk, questions, comments and PRs all welcome but I'm not actively working on this project so please be patient!
