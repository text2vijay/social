# social

golang social

> go mod init github.com/text2vijay/social

> go build -o outputname filename.go - specific build name then use this

================================================================
mkdir my-web-project
cd my-web-project
go mod init github.com/username/my-web-project

my-web-project/
├── main.go
├── go.mod
├── go.sum
├── templates/
│ └── index.html
├── static/
│ └── style.css

If you're using any third-party Go packages (e.g., a web framework like gin or mux), you’ll need to install them using go get. For example, to install the mux router:
go get -u github.com/gorilla/mux

build project:
go build

./my-web-project

go build -o my-web-server

to run the project
./my-web-server
================================================================
create a docker image for a web project

# Use the official Go image as the base image

FROM golang:1.20-alpine

# Set the Current Working Directory inside the container

WORKDIR /app

# Copy the Go Modules and download dependencies

COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container

COPY . .

# Build the Go app

RUN go build -o my-web-server .

# Expose port 8080 to the outside world

EXPOSE 8080

# Command to run the executable

CMD ["./my-web-server"]

=>after creating a docker file build a docker container.
docker build -t my-web-server .
docker run -p 8080:8080 my-web-server

================================================================

=> setup go modules
go mod init github.com/username/my-web-project

=> install dependencis
go get -u github.com/gorilla/mux

=> build the project
go build

=> run project
./my-web-project

=> Build Docker image for the project:

> docker build -t my-web-server .
> docker run -p 8080:8080 my-web-server
> ===================================================
