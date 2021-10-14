FROM golang

# Environments & Variables
ARG ssh_prv_key
ARG ssh_pub_key
ENV GO111MODULE=on

WORKDIR /app

# Update Software repository
RUN apt-get update && \
    apt-get upgrade -y
# Download libraries in linux
RUN apt-get install -y \
        git \
        openssh-server
# Authorize SSH Host
# RUN mkdir -p /root/.ssh && \
#     chmod 0700 /root/.ssh && \
#     ssh-keyscan github.com > /root/.ssh/known_hosts
# # Add the keys and set permissions
# RUN echo "$ssh_prv_key" > /root/.ssh/id_rsa && \
#     echo "$ssh_pub_key" > /root/.ssh/id_rsa.pub && \
#     chmod 600 /root/.ssh/id_rsa && \
#     chmod 600 /root/.ssh/id_rsa.pub

RUN git config --add --global url."git@github.com:".insteadOf https://github.com
RUN go env -w GOPROXY=direct

# Add the rest of the files
COPY go.mod .
#COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# Remove SSH keys
RUN rm -rf /root/.ssh/

EXPOSE 8080
#ENTRYPOINT /app/main

RUN go build ./api/main.go
RUN ls -l
ENTRYPOINT /app/main
# CMD ["./app/main"]

#docker build -t example2 --build-arg ssh_prv_key="$(cat ~/.ssh/docker_deploy)" --build-arg ssh_pub_key="$(cat ~/.ssh/docker_deploy.pub)" --squash .


#docker build -t cachy-api --squash . 
#docker run -d --rm -p 8080:8080 --name cachy-api-1 cachy-api