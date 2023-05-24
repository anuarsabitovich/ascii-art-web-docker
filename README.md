# ascii-art-web-dockerize
## Authors: anuarsabitovich, Bakytzhan16
### Objectives


Objectives

You must follow the same principles as the first subject.

    For this project you must create at least :
        one Dockerfile
        one image
        one container

    You must apply metadata to Docker objects.

    You have to take caution of unused object (often referred to as "garbage collection").

Instructions

    The web server must be created in Go.
    The code must respect the good practices.
    You must use Docker.
    The project must have a Dockerfile.
    Your Dockerfile must respect the Dockerfile good practices.


## Usage

```
After cloning the repository, 
open console and follow instructions below : 
 
type in console "make create" - to build docker image
type in console "make run" - to run docker container on a port 8080
type in console "make start" - to start docker container  
open your browser paste http://localhost:8080/ to check if the container is running and working
type in console "make stop"  - to stop docker container
type in console "make prune" - to clear all stopped containers 
```

### Implementation details: algorithm

- After creating the server, each connection gets handled by the handlers that are stored in the handlers folder.
- Each request is being checked for possible errors and if none are found, the parsed template with necessary information will be shows in the browser.
- When a client enters text inside the form and submits it, the POST form with an */ascii* url is being passed as a request to the server, which extracts the input values and puts them on a separate rendered template with the Ascii art logic.
- In order to check for form submission accuracy, ascii art has been modified to return a string with an error that will handle any ASCII ART related mistakes.