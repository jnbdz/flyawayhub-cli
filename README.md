# flyawayhub-cli
Universal Command Line Interface for Flyawayhub (not official). This is for the one that want to interact with Flyawayhub without the need of using a browser.

For those that want to avoid the bloated web interfaces and are into doing as many things in the terminal as possible.

## IMPORTANT!
This project is in no way related to the company or website of [Flyawayhub](https://app.prod.flyawayhub.com/).

**WARNING!** It is still under construction!

## What it does
It bassically takes your credentials and connects to Flyawayhub and replaces your browser.

The website is basically a SPA (Single-page application) that uses a RESTFul API that returns JSON data.

Some of the JSON data is then made human readble.

You can always change it to return the raw data.

## How to build
You can build it the way you do for any other Go code or you can use a [podman](https://podman.io/) container to do it for you without the need to install Go on your system.

So if you want to use podman: 
```bash
./buildgo.sh
```
It should grab automatically whatever is in this root directory (of the project) and then build it. Nothing else is needed.

## Run it
After it is built you can easily just run the output of the build: `./flyawayhub-cli`

You can always rename it: `mv flyawayhub-cli flyawayhub`.

Have fun in the terminal!
