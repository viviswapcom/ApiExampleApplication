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

## Packaging and Distribution

Packaging for multiple operating systems can be a complex task. Graphical applications typically have icons and metadata associated with them as well as specific formats required to integrate with each environment.

They fyne command provides support for preparing applications to be distributed across all the platforms the toolkit supports. Running “fyne package” will create an application ready to be installed on your computer and to be distributed to other computers by simply copying the created files from the current directory.

For Windows it will create a .exe file with icons embedded. For a macOS computer it will create a .app bundle and for Linux it will generate a .tar.xz file that can be installed in the usual manner.

Of course you can still run your applications using the standard Go tools if you prefer.

    go get fyne.io/fyne/cmd/fyne

    go build
    fyne package -icon mylogo.png
