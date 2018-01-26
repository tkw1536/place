# place

A lightweight docker image that allows you to place your static website on the internet and update it using Web Hooks.

## How it works

TODO: Document me


## Usage

1. Create a repository containing a static website. We will assume your repository is `https://github.com/user/domain.tld`, with matching domain `domain.tld`.
2. Start up this docker container and serve it at the given domain:
```
docker run -p 80:80 -v /data -v /var/www/html -e REPO_URL=git@github.com:user/domain.tld.git -e HOOK_SECRET=super-secret-token place/place
```

3. Add the generated ssh public key as a deploy key to your repository.
4. Create a GitHub webhook to point to "https://domain.tld/hook/" and add the token `super-secret-token`
5. Push to your repository, and the deployment will magically be updated
