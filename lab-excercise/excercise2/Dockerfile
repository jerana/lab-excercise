#Use official golang  as base image
FROM golang
#set working directory to excercise
WORKDIR /excercise
#copy source code into workdir 
ADD . /excercise
RUN go get golang.org/x/net/html
#build seqReadfile inside the container
RUN go build -o excercise2
#run time context when conatiner start
CMD ["/excercise/excercise2"] 
