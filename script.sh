# Names for our image & container
IMG=my_img
CTR=my_ctr

# Remove all Docker images
docker rmi -f $(docker images -aq)

# Remove all Docker containers
docker rm -f $(docker ps -aq)

# Build Our image, (--no-cache)=> each build step will be executed without retrieving already stored data.
docker build --no-cache -t $IMG .

# Run the Docker container and link port 8080 on the container to port 8080 on your machine
# Run container -d => This option starts the container in detached mode (background)
docker run -d -p 8080:8080 --name $CTR $IMG

# this command to execute a command in our container (i)=> let the standard input open, (t)=> create a virtual terminal in Our container
docker exec -it my_ctr  /bin/bash