# Fishbot 🐟

A clone look-a-like from the original bot of Quakenet, but made in Go.

### 🔧 How to use

1. Rename the file __config.yml.example__ to __config.yml__
2. Change the `Hostname`, `Port` and other parameters as desired
3. Compile and run using `go build .` and `./fishbot`
   - Or run it directly with `go run .`

If you are using Docker to deploy the bot you can use environment variables to
override config.yml parameters.

### 🎣 How to invite bot to my channel

You have to set the config parameter `allow_invites` to `true`. Then everybody
will be able to add the bot to its channel using `/INVITE <botnick> <#channel>`.
