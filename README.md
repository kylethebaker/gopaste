# GoPaste

_simple command line pastebin-esque site for personal use_

- - -

Using [prismjs](http://prismjs.com/) for the optional syntax highlighting
and Gorilla Toolkit [Mux](http://www.gorillatoolkit.org/pkg/mux) for routing.

## Setup and Installation

 - Pick the appropriate paths, site url, and listen port for your setup. 
 - Create a directory to store the uploaded pastes. 
 - ```go get github.com/gorilla/mux```
 - ```go build``` or install or whatever you want.

## Usage

```cat file-to-upload.txt | curl -F 'paste=<-' http://your-site```

I like to add ``` | xargs firefox``` to the end to open it in firefox

You can add ?lang=html to the url to get syntax highlighting. The langs
supported are provided by Prism. Without the lang param (or with an invalid
lang), a plaintext response is served. You can see the LANGS global for 
langauges I've included from prismjs.

## Thanks

Thanks for http://sprunge.us for the idea which I shamelessly copied. Thanks
to prismjs for the awesome syntax highlighting. Thanks to Goriall Toolkit for 
the awesome Golang http extensions.

## Disclaimer

I'm not responsible for any personal injuries (or any other damages) caused
by this software.
