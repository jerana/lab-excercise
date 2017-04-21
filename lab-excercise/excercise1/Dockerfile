#Use official  as base image
FROM golang
#set working directory to excercise
WORKDIR /excercise
#copy source code into workdir 
ADD . /excercise
#build seqReadfile inside the container
RUN go build -o excercise1
#run time context when conatiner start
CMD ["/excercise/excercise1"] 
