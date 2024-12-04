FROM golang:1.23.2

WORKDIR /app

COPY . .

RUN CGO_enabled=1 GOOS=linux GOARCH=amd64 go build -o ./JC_project.exe ./cmd/*.go

CMD [ "./JC_project.exe" ]