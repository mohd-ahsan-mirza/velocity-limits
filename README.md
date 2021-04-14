
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

### Error 1
There is a pause of 10s in execution of the script to give the docker network to be fully built.
If you see this error `connect: connection refused`, open the makefile and increase the sleep count
### Error 2
If you get this error `no such host`, check your docker and docker-compose version and make it's upto date. 
You can also try removing all the docker images you have right now on your machine `docker rmi $(docker images -a -q)`. If you have too many images, postgres might not have enough space to initiate and run.
