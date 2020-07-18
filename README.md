# Sample Flask App for 2-way APIs

This application contains: one flask app in `app.py`. It serves up data about planets from [NASA data](https://solarsystem.nasa.gov/moons/in-depth/) at `/planets` and `/planets/<position>`. It also logs the incoming data it receives to a `/webhook` endpoint.

Run it like this:

```
FLASK_APP=app.py python -mflask run
```

This was created for a one-off talk, questions, comments and PRs all welcome but I'm not actively working on this project so please be patient!

## HTTP Tools

These are the HTTP tools I enjoy and recommend:

* Start simple with curl <https://curl.haxx.se/>, a widely-available cross-platform command line HTTP client.
* Another command line HTTP client, HttPie <https://httpie.org/> shows response status codes, headers and pretty prints JSON by default
* Graphical favourite: Postman <https://www.postman.com/> is awesome and the features would take a whole other README file! Let's go with: import OpenAPI specs, save requests, inspect responses, change parameters easily, and use environments to hold regularly-used values
* Use HTTPBin <https://httpbin.org> to supply known responses to an API call. This is brilliant for checking your code responds as expected to specific status codes, or to have HTTPBin return information about the request it received (such as `/headers` which simply prints the headers your request included)
* A brillant tool from former stellar startup Runscope, RequestBin <https://github.com/lornajane/requestbin> is a place to receive and inspect any HTTP request. I use it for checking my API is sending what I _think_ it is sending, and for receving events by webhook when I haven't implemented my own receiver yet. (The link here is my own fork, sadly the main repo is no longer maintained but I have assembled the best of what I found in the PR queue for it so I think this is the best available version today)
* Ngrok <https://ngrok.com> is a tunnel to your local development platform. Use it to test a local website project on your phone, or to receive webhooks from an external supplier into your development code.

