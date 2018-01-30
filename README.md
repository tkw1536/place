# place

A lightweight docker image that allows you to place your static website on the internet and update it using Web Hooks.

** WORK IN PROGRESS DO NOT USE YET **

## TODO

* TODO: `bin/place` argument passing
* TODO: Dockerfile argument testing
* TODO: tl;dr documentation
* TODO: `bin/place-server` testing
* TODO: `bin/place testing`
* TODO: Rewrite tests in go

## TL;DR

1. Create a repository, e.g. `https://github.com/user/domain.tld`, and setup `domain.tld` DNS to point to your docker instance. 
2. Start this docker container

```
docker run -p 80:80 -v /data -v /var/www/html -e place/place
```

3. Add the generated ssh public key as a deploy key to your repository.
4. Create a GitHub webhook to point to "https://domain.tld/hook/" and add the token `super-secret-token`
5. Push to your repository, and the deployment will magically be updated

### bin/place
The main place executable, which serves as the entry point for the container. It reads in all parsed parameters, generates an ssh key if it does not exist, and then delegates to `bin/place-server` (see below). 

```
Usage: bin/place
Arguments:
    SSH_KEY_PATH
        If set, enable support for ssh.
        First, try to load the passwordless SSH Key at the given path. If it does not exist, generate a new one. 
    
    BIND_ADDRESS
        Address to bind to (see `place-server --bind`)

    WEBHOOK_PATH
        Path that should respond to webhooks, defaults to `/webook/` (see `place-server --webhook`)

    GIT_URL
        Repository to clone (see `place-git-update --from`)

    GIT_BRANCH
        Branch of the Git Repository to clone
    
    GIT_CLONE_TIMEOUT
        Timeout to give to webhook (see `place-server --timeout`)
    
    GITHUB_SECRET
        If non-empty, listen to GitHub Webhooks with the given secret (see `place-server --github`)
    
    GITLAB_SECRET
        If non-empty, listen to GitLab Webhooks with the given secret (see `place-server --gitlab`)
    
    DEBUG
        If set to 1, listen to Debug Webhooks (see `place-server --debug`)
    
    STATIC_PATH
        Path to place static content in
    
    BUILD_SCRIPT
        If given, build script to inject into static process (see `place-git-update --build`)

    PROXY_URL
        When set proxy to the given url instead of servering static content (see `place-server --proxy`)

```

Instead of calling `bin/place-server` as an external program, the code directly links into the binary code below. 

### bin/place-server

The web server that servers the static directory and listens to webhooks. 

```
Usage: bin/place-server \
    --bind address --webhook path \
    [--github token[,ref[,events...]]] [--gitlab token[,ref[,events...]]] [--debug] \
    --script script [--timeout timeout] --static dir|--proxy url

    --bind address
        The interface address and port the server should bind to.
    --webhook path
        The path that should respond to webhooks (e.g. /webhook/).

    --github token[,ref[,events...]]
        Run the webhook whenever a GitHub web hook request is received.
        token
            The WebHook Token
        ref
            The Ref to run WebHooks on. Defaults to refs/heads/master.
        events
            Comma-seperated list of events to run WebHooks on.
            Defaults to "Push".
    --gitlab token[,branch[,events...]]
        Run the webhook whenever a GitLab web hook request is received.
        token
            The WebHook Token
        ref
            The Ref to run WebHooks on. Defaults to refs/heads/master.
        events
            Comma-seperated list of events to run Webhook on.
            Defaults to "Push Hook".
    --debug
        Run in debug mode and run the webhook whenever any POST request is received.

    --script script
        Script to run whenever a webhook is received. 
    --timeout timeout
        Timeout for script in seconds. Defaults to 600.

    --static dir
        Serve static content from directory dir.
    --proxy url
        Instead of serving static content, proxy all requests to url
```

### bin/place-git-update

Executable that updates a directory with static content from a git repository.  

```
Usage: bin/place-git-update --from url [--ssh-key path] --to path [--ref ref] [--build script]

    --from url
        The url the remote repository is located at
    --ssh-key path
        If set, load a passwordless ssh key from the given path and use it to clone the repository. 
    --to path
        The (local) path the static files should be placed at
    --ref ref
        The ref the repository that should be checked out (e.g. refs/heads/master)
    --build script
        By default, files will be placed from the source repository directly into the target
        directory. By specifying the `build` parameter, the files can be pre-processed before
        being injected into the target directory.

        The build script will be started inside a directory that contains a checkout of the repository
        and will receive the target directory as an argument. The target directory will contain
        the old (dirty) state of the working directory. Build scripts are run using the
        current $SHELL environment variable (falling back to /bin/sh if unspecified).

        Example value: "bundle install; bundle exec jekyll build -t"
```
