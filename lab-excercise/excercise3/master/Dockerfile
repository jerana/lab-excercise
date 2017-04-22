#Use official golang  as base image
FROM golang
#set working directory to excercise
WORKDIR /master
#copy source code into workdir 
ADD . /master
#build seqReadfile inside the container
RUN go get github.com/jerana/lab-excercise/lab-excercise/excercise3/common
RUN go build  -o master
#run time context when conatiner start
CMD ["/master/master"] 
