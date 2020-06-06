# Omoku API example application

This is a small example application demonstrating the usage of the omoku public API (https://api-gateway.omoku.io/docs/).

## Prerequisites

In order to run this application you need the following tools installed on your device:

1. Download Go from the [download page](https://golang.org/dl/) and follow instructions
2. Install fyne.io using [this instructions](https://fyne.io/develop/index)

That's all you need!

## Build and run the demo application

To run this example application just execute the command

    go run main.go

To build a execution file simply use the following code

    go build main.go
    ./main.go

## Packaging and Distribution

Packaging for multiple operating systems can be a complex task. Graphical applications typically have icons and metadata associated with them as well as specific formats required to integrate with each environment.

They fyne command provides support for preparing applications to be distributed across all the platforms the toolkit supports. Running “fyne package” will create an application ready to be installed on your computer and to be distributed to other computers by simply copying the created files from the current directory.

For Windows it will create a .exe file with icons embedded. For a macOS computer it will create a .app bundle and for Linux it will generate a .tar.xz file that can be installed in the usual manner.

Of course you can still run your applications using the standard Go tools if you prefer.

    go get fyne.io/fyne/cmd/fyne

    go build
    fyne package -icon mylogo.png