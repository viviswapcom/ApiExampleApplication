# Omoku API example application

This is a small example application demonstrating the usage of the omoku public API (https://api-gateway.omoku.io/docs/).
This application is not meant to be used for production environments but only for testing and local development purposes!

## Prerequisites

In order to run this application you need the following tools installed on your device:

1. Download Go from the [download page](https://golang.org/dl/) and follow instructions
2. Install fyne.io using [this instructions](https://fyne.io/develop/index)

That's all you need!

## Configuration

In order to configure this application simply duplicate the `config-example.json` file and rename it to `config.json`.
It is important that you at least adjust the `apiKey` variable to your own api-key.
But of course you can also customize all the other values as you like ;-)

## Build and run the demo application

To run the application you can run the command

    go run main.go

In order to compile this application to run multiple times on your operating system execute

    go build main.go
