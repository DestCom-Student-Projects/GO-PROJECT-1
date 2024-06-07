FROM golang:1.22.3-alpine

USER root
WORKDIR /app/go-supermarket

RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.2.0/zsh-in-docker.sh)"
#RUN chsh -s /bin/zsh root

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-supermarket

CMD ["/bin/zsh"]