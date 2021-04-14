
# velocity-limits

## Prerequisite
* Docker version 20.10.5
* Docker-compose version 1.28.5

## Usage
* CD in `cmd`
* Run `make`

## Results
* [Screen Capture](https://github.com/mohd-ahsan-mirza/velocity-limits/blob/master/result/screen_recording.mov)
* [Result Output](https://github.com/mohd-ahsan-mirza/velocity-limits/blob/master/result/result.txt)

## Troubleshooting
There is a pause of 10s in execution of the script to give the docker network to be fully built.
If you see this error `connect: connection refused`, open the makefile and increase the sleep count

