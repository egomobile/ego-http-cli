# ego-http-cli
A command line interface for @egomobile/http-server framework

## Installation
ego-http-cli is a binary that can be downloaded from Github.

Execute the following command to download and extract the binary. Please remember to replace the url with latest version.

`curl -sL https://github.com/egomobile/ego-http-cli/releases/download/v0.1.0/ego-http-v0.1.0-darwin-amd64.tar.gz | tar xz`

After that move the binary to `usr/local/bin` directory by executing the following command

`mv ego-http /usr/local/bin`

Execute the following command to test if installation is successful

`ego-http help`

## Execute
Run the following command from the root directory of the project.

`ego-http generate controller -n Foo -p v1/foo`

This command will generate a controller called `FooController` in the controllers directory inside `v1/foo` path.