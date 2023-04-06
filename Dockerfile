FROM golang:1.20.3-bullseye

ENV APP_HOME /app
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"


# Install packages
RUN apt-get update -y && \
    apt-get install -y wget unzip && \
    rm -rf /var/lib/apt/lists/*

# Install nomad
RUN wget https://releases.hashicorp.com/nomad/1.4.2/nomad_1.4.2_linux_amd64.zip && \
    unzip nomad_1.4.2_linux_amd64.zip -d /usr/local/bin && \
    rm nomad_1.4.2_linux_amd64.zip

COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

RUN go build -o server

EXPOSE 8000

CMD ["/app/server"]

