# place

A lightweight docker image that allows you to place your static (or not-so-static) website on the internet and update it using Web Hooks.

** WORK IN PROGRESS DO NOT USE YET **

## TODO

* TODO: More testing
* TODO: `GitSSHKeyPath` should be in-lineable

## TL;DR

1. Create a repository, e.g. https://github.com/example/domain.tld
2. Start this docker container with:

```
docker run -p 80:80 -v /data -v /var/www/html place/place
```

3. Go to http://localhost/setup/ and follow the on-screen instructions. 

## How it works

This container consists of two main components:

- The hook server, simply refered to as `place`
- The setup server, found inside the `setup` folder, and called with `place-setup`

### The Hook Server

The hook server implements two behaviours:

1. It serves static content (or proxies to another server)
2. It listens for webhooks under the `/webhook/` path

Whenever a valid web hook is received (currently only GitHub and GitLab hooks are supported), the hook server updates content from a remote (git) repository and places it into the folder served by the static web server.
The server is configured using various settings, see `The Configuration File` section below.  

### The Setup Server

The setup server runs as follows:

1. Try to load the configuration file
2. If that succeeds, start the hook server with this configuration file
3. If that fails, start the setup server, wait for the user to finish setup, and go back to 1.

As such, it needs three parameters:

1. The path to the configuration file, given with the `CONFIG_PATH` variable, and defaulting to `config.json`
2. The port to listen on, given with the `BIND_ADDRESS` variable, and defaulting to `0.0.0.0:80`
3. The path to the `place` executable, given as a command-line parameter

The first parameter (`CONFIG_PATH`) is set to `/data/config.json` in the Dockerfile.
Since a `/data/` is also declared as a Docker Volume, this means that the configuration file is stored persistently.
Furthermore, it allows to skip setup entirely by pre-seeding the file into the Volume manually.

The second parameter, `BIND_ADDRESS` is left unset, and thus assumes the default.
Port `80` is `EXPOSE`d in the Dockerfile, and thus it should not be neccessary to change this.
Nonetheless, the port can be changed by passing an appropriate environment variable to the docker container.

The third parameter is hard-coded using inside the Dockerfile by giving an appropriate `ENTRYPOINT` parameter.
This should not need changing.


### The Configuration File

The configuration file is a simple JSON file, which can be used to configure the behaviour of the hook server.
It is defined in [config/config.go] and has the following configuration options:

* `BindAddress` The interface address and port the server should bind to.
* `WebhookPath` A string path to server webhook under, defaults to `/webhook/`
* `GitURL` Url to clone upstream repository from
* `GitBranch` Branch to clone and set of events to listen to, defaults to `master`
* `GitSSHKeyPath` Path to ssh key for git clone (if any)
* `GitUsername` username for git clone (if any)
* `GitPassword` password for git clone (if any)
* `GitCloneTimeout` timeout for git clone in Nanoseconds, defaults to 10 minutes
* `GitHubSecret` secret to use when listening to GitHub Events (if any)
* `GitLabSecret` secret to use when listening to GitLab Events (if any)
* `Debug` if set to `true`, trigger the webhook on any post request to webhook path (not recommended)
* `StaticPath`
    Path to place static content in, and to serve static content from.
* `ProxyURL`
    If set, instead of serving static content, proxy all content this url.
    This can be used if you want to automaticallty update content via `place`, but are not serving a static website.
    This allows to e.g. serve PHP content via an external server.
    Note that no external servers are included in the docker image by default and you would have to add a custom one.

* `BuildScript`
    By default, files will be placed from the source repository directly into the target
    directory. By specifying the `BuildScript` parameter, the files can be pre-processed before
    being injected into the target directory.

    The build script will be started inside a directory that contains a checkout of the repository
    and will receive the target directory as an argument. The target directory will contain
    the old (dirty) state of the working directory. Build scripts are run using the
    current $SHELL environment variable (falling back to /bin/sh if unspecified).

    Example value: "bundle install; bundle exec jekyll build -t"

## How about HTTPS and multiple host support?

This minimal docker image is not intended to handle HTTPS and friends.
If you desire support for multiple hosts, and support https, you should consider using [jwilder/nginx-proxy](https://github.com/jwilder/nginx-proxy) in front of this docker image.

## License

MIT
