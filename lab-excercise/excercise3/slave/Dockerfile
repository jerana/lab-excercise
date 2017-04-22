#Use official golang  as base image
FROM golang
#set working directory to excercise
WORKDIR /slave
#copy source code into workdir 
ADD . /slave
#build seqReadfile inside the container
RUN go get github.com/jerana/lab-excercise/lab-excercise/excercise3/common
RUN go build -o slave
#run time context when conatiner start
CMD ["/slave/slave"] 
