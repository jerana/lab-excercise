# LabExcercise
There are some exercise Go project code which are as follow
Project goal
------------
Count word frequency of Moby Dick (aka, from a text corpus from web). There
are three parts to this project. In the first part, write a program to count
the words from a single text file. In the second part, write a program to
count the words in multiple text files in parallel. In the third part, write a
program that has master-slave configuration.

For all parts, build one or more container(s) containing your app and run them
using docker tools.

Part I - sequential
-------------------

Write a golang program to
1. Read the text file from http://www.gutenberg.org/files/15/text/moby-000.txt
2. Generate word count for the whole file. It should be a map of kind
   map[string]int where key would be word and value would be frequency of that
   word in the text file.
3. Save results to a text file.

Build a docker container containing above app
1. Update the Dockerfile to add your application
2. Use docker build commands to build and run the container

Deliverables:
1. One .go file containing all the code.
2. One Dockerfile to build this code as container.
	
Part II - parallel
------------------

Write a golang program that 
1. Uses multiple goroutines - one for each text file to generate the
   map. The goroutines should read all text files from
   http://www.gutenberg.org/files/15/text/
   [NOTE] - gutenberg.org imposes an access/rate limit. You can either use
   a mirror site, or download http://www.gutenberg.org/files/15/15-text.zip,
   unzip it, and use local file I/O for rest of the project.
2. Merge results from all these goroutines in main
3. Save results to a text file.

Build a docker container containing above app
1. Update the Dockerfile to add your application
2. Use docker build commands to build and run the container

Deliverables: 
1. one .go file containing all the code.
2. One Dockerfile to build this code as container. 

Part III - distributed across multiple containers (optional)
-----------------------------------------------------------
NOTE: Send the solution for Part I & II as you work on Part III, which may
take longer timer.

1. Setup multiple containers - one master and N slaves, where N would
   be decided at runtime.

2. Slaves run a json-rpc server. They implement a single API, called
   word count. It's signature is:

	func WordCount (text []string) []string {

	}

   One limitation of json-rpc is that it can't serialize/de-serialize a
   map, so you need to do it yourself. On slave, the map of Part1, step 2 should
   get serialized into a slice of strings. On master you need to deserialize back
   to original map structure.

3. Master divides all files across different slaves, makes the rpc
   calls, and then aggregates the overall response.

4. Save results to a text file.

Build a docker container containing above app
1. Update the Dockerfile to build master and slaves
2. Use docker build commands to build and run the containers
3. Update the sample docker-compose.yml file to run the containers
4. Run all the containers using docker-compose command (hint: use
   https://docs.docker.com/compose/reference/scale/ to scale the slaves)

Deliverables: 
	1. One .go file containing code for master
	2. One .go file containing code for slave
	3. One Dockerfile to build master container
	4. One Dockerfile to build slave container
	5. One docker-compose.yml file to run these various containers.



