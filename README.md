# ci-server
> Simple continuous integration server for automatically deploying
> [jadeocr-next](https://next.jadeocr.com)

This is a simple server that builds [jadeocr-next](https://next.jadeocr.com) in
Docker upon pushes to the jadeocr-next
[repository](https://github.com/jadeocr/jadeocr-next).

Using GitHub webhooks on the jadeocr-next repository, this server triggers
shell commands to fetch the changes from GitHub, rebuild the docker-compose
images, and start the new containers.

# Usage
1. Get a secret [token](https://docs.github.com/en/developers/webhooks-and-events/securing-your-webhooks)
  from GitHub.
1. Then, using the value of that token as an environment variable, run `yarn
  start`
1. Send a POST request with body `token: MY_TOKEN_VALUE` to
  `http://my_server_ip:4000/github`

