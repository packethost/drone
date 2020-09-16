![](https://img.shields.io/badge/Stability-Experimental-red.svg)

Drone is a Continuous Delivery system built on container technology. Drone uses a simple YAML configuration file, a superset of docker-compose, to define and execute Pipelines inside Docker containers. 

This repository is [Experimental](https://github.com/packethost/standards/blob/master/experimental-statement.md) meaning that it's based on untested ideas or techniques and not yet established or finalized or involves a radically new and innovative style! This means that support is best effort (at best!) and we strongly encourage you to NOT use this in production.

<br/>

<img src="https://github.com/drone/brand/blob/master/screenshots/screenshot_build_success.png" style="max-width:100px;" />

Sample Pipeline Configuration:

```yaml
pipeline:
  backend:
    image: golang
    commands:
      - go get
      - go build
      - go test

  frontend:
    image: node:6
    commands:
      - npm install
      - npm test

  publish:
    image: plugins/docker
    repo: octocat/hello-world
    tags: [ 1, 1.1, latest ]
    registry: index.docker.io

  notify:
    image: plugins/slack
    channel: developers
    username: drone
```

Documentation and Other Links:

* Setup Documentation [docs.drone.io/installation](http://docs.drone.io/installation/)
* Usage Documentation [docs.drone.io/getting-started](http://docs.drone.io/getting-started/)
* Plugin Index [plugins.drone.io](http://plugins.drone.io/)
* Getting Help [docs.drone.io/getting-help](http://docs.drone.io/getting-help/)
