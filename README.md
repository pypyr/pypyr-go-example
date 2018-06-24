# pypyr-go-example
example project for golang using a ci devops container pipeline runner for
build, test and deploy

Notice the shiny CI and coverage badges:
[![Run Status](https://api.shippable.com/projects/5b2ef39934d4ed070034de8d/badge?branch=master)](https://app.shippable.com/github/pypyr/pypyr-go-example)
[![Coverage Badge](https://api.shippable.com/projects/5b2ef39934d4ed070034de8d/coverageBadge?branch=master)](https://app.shippable.com/github/pypyr/pypyr-go-example)

## install
Get the source into your GOPATH:

`go get -u github.com/pypyr/pypyr-go-example`

## prerequisites
- go >=1.10 https://golang.org/
- go dep https://github.com/golang/dep
- docker >=18

If you want to run the pipelines directly, you'll need the following:
- &gt;= python 3.6
- `pip install --user pypyr`

Maybe do this in a virtualenv. You don't necessarily _need_ pypyr on your
dev machine, you could actually just use the _pypyr/pypyr-go_ official container
image to run all the devops and ci functions if you had wanted to keep your dev
machine clean. You'll get annoyed with this, however. It's handy to use the
container as a final clean, sanity check before you push your code as part of
your commit and pull request discipline, and run the same container on the CI
so you don't have to deal with initializing the environment on the build server
outside of the known container format.

## dependencies, build, test
You can run all the ci commands locally. In fact, it's handy for checking your
work before you push.
```bash
$ pypyr --dir ops install # install dev tooling
$ pypyr --dir ops vendor # install go vendor packages
$ pypyr --dir ops quality # run linting and code quality checks. Unit Tests.
$ pypyr --dir ops build # run quality, then build the go package to ./cmd/magritte
$ pypyr --dir ops ci # run all of the above, and compile statically linked binary to ./bin/app
```

To get started as a first install on your dev machine, run these 3:
```bash
$ pypyr --dir ops install # install dev tooling
$ pypyr --dir ops vendor # install go vendor packages
$ pypyr --dir ops build # run quality, then build the go package to ./cmd/magritte
```

In theory you'll run install and vendor only as a first setup, and then after
only when you add new dependencies or vendor packages. The build pipeline you'll
probably want to run in isolation a lot more frequently.

## run go ci build inside container
To isolate your environment and avoid having to install dependencies, do it
all inside a container:

`docker run -v ${PWD}:/go/src/github.com/pypyr/pypyr-go-example -w /go/src/github.com/pypyr/pypyr-go-example pypyr/pypyr-go --dir ops ci`

This is the same _ops/pipelines/ci.yaml_ pipeline that you can run locally with
`pypyr --dir ops ci`.

It outputs a statically linked binary to `bin/app`.

You can also run this container on your CI system the same way as you run it
locally on your dev box - that way you know you're always building in the same
sandbox without surprises. See _shippable.yaml_ for an example.
