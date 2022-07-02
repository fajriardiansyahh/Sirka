# Specify the base image for the go app.
FROM golang:1.18

# Add Maintainer info
LABEL maintainer="Fajri Ardiansyah"

# Specify that we now need to execute any commands in this directory.
WORKDIR /go/src/github.com/postgres-go

# Copy everything from this project into the filesystem of the container.
COPY . .

# Obtain the package needed to run code. Alternatively use GO Modules. 
RUN go get -u github.com/lib/pq
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/joho/godotenv

# Download & Install all the dependencies
# RUN go get -d -v ./...
# RUN go install -v ./...

# Compile the binary exe for our app.
RUN go build -o main .

# Start the application.
CMD [ "/build" ]