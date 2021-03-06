# Dev Ops

Ops are maintained in
[![opspec 0.1.5](https://img.shields.io/badge/opspec-0.1.5-brightgreen.svg?colorA=6b6b6b&colorB=fc16be)](https://opspec.io/0.1.5/packages.html)
format.

They can be consumed via tools like [opctl](https://opctl.io).

# Acceptance criteria

Contributions are subject to:

- accepted review by one or more
  [maintainers](https://github.com/orgs/opctl/teams/maintainers/members)
- the [build](.opspec/build) op continuing to run with a successful
  outcome
- adherence to
  [go code review comments](https://github.com/golang/go/wiki/CodeReviewComments)


# Repo organization

## /cli

CLI, distributed w/ the opctl binary

The CLI is built using [mow](https://github.com/jawher/mow.cli)

## /docs

docs, hosted at [https://opctl.io/docs](https://opctl.io/docs)

## /node

daemon, distributed w/ the opctl binary

hosts the opctl web app & an opspec node

## /webapp

web app, distributed w/ the opctl binary & hosted by the opctl daemon.

It is a static web app built using
[react](https://facebook.github.io/react/) & was bootstrapped with
[Create React App](https://github.com/facebookincubator/create-react-app).

## /website

opctl website, hosted at [https://opctl.io](https://opctl.io)

It is a static website built using
[metalsmith](https://github.com/metalsmith/metalsmith)


# Testing

`opctl run test` runs all unit tests inclusive of code coverage.

## Fakes

To streamline unit test related maintenance, [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) is used to auto-generate fake implementations of interfaces.

The fakes are then used to assert on & stub the object under tests interactions w/ its dependencies. 