# milkcocoa-go-mqtt-example


## Installation

set GOPATH and GOBIN

Install glide and clone this repository.

```
curl https://glide.sh/get | sh
mkdir $GOPATH/src/github.com/kunihiko-t
cd $GOPATH/src/github.com/kunihiko-t
git clone https://github.com/kunihiko-t/milkcocoa-go-mqtt-example.git
```

Do glide Install

```
cd $GOPATH/src/github.com/kunihiko-t/milkcocoa-go-mqtt-example
glide install
```


## Run

### Subscribe

MILKCOCOA_APP_ID=YOUR_APP_ID go run main.go

### Publish

MILKCOCOA_APP_ID=YOUR_APP_ID go run pub/pub.go
