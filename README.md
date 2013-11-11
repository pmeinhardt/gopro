# gopro.go

## Overview

The package contains 2 sets of APIs (soon available, work in progress):

1. **control your [GoPro](http://gopro.com/) camera**, this includes
  - switching camera modes (video, photo, burst, timelapse)
  - starting and stopping capture
  - setting the camera orientation
  - video resolution, field of view
  - photo resolution, timer
  - deleting your last or all files
2. **inspect and download captured photos and videos**
  - bye, bye USB adapter

## Installation

Make sure you have the a working Go environment. See the
[install instructions](http://golang.org/doc/install.html).

To install gopro.go, simply run:

    go get github.com/pmeinhardt/gopro

## Usage

    package main

    import (
        "github.com/pmeinhardt/gopro"
        "time"
    )

    func main() {
        cam := gopro.DefaultCamera("password")
        err := cam.StartCapture()

        time.Sleep(4 * time.Second)

        cam.StopCapture()
    }

To run the application, put the code in a file called capture.go,
connect to your GoPro's Wi-Fi network and run:

    go run capture.go

## References

1. http://gopro.com/support/open-source
2. http://forums.openpilot.org/topic/15545-gcs-go-pro-wifi-widget/
3. http://goprouser.freeforums.org/howto-livestream-to-pc-and-view-files-on-pc-smartphone-t9393-170.html
