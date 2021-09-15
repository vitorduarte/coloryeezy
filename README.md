# coloryeezy 🎨👟

![](https://media-exp1.licdn.com/dms/image/C4E22AQGTmGycnag__A/feedshare-shrink_2048_1536/0/1629156813171?e=1634774400&v=beta&t=f69aHRAW071bI1OhMFc8WieavOm1aqQrxNWlHuRRa-s)

The ColorYeezy project is a bot that generate random colorways for Mr. West's sneakers. We post the results every hour on [Twitter](https://twitter.com/coloryeezy). It's the exact opposite of what sneaker-related bots do. We don't take from the community. We're making more colorways for people to discuss and to enjoy seeing them.

It's like having 24 drops a day!

# Pre requisites

- [go](https://golang.org/doc/install)
- [docker](https://docs.docker.com/engine/install/) (optional)
- [docker-compose](https://docs.docker.com/compose/install/) (optional)

## Setup

Create a `.env` file with the following parameters

| Parameter             | Description                        |
| --------------------- | ---------------------------------- |
| `API_KEY`             | API Key of Twitter API             |
| `API_SECRET_KEY`      | API Secret Key of Twitter API      |
| `ACCESS_TOKEN`        | Access Token of Twitter API        |
| `ACCESS_TOKEN_SECRET` | Access Token Secret of Twitter API |

## Config

This project could be used to paint other images too. You just need to modify the configuration file `./config.json`. Its structure has the following properties:

- template: The path of the template file that will be painted.
- masks: List of masks that will be used to guide the paint process. It has to be a transparent png image with the area that will be painted filled with some color.
  - path: The path of mask image.
  - color: The color that will be used to paint the image on mask area. It should be a hex string (#xxxxxx or #xxx) or "random" which means that the color will be randomly generated.

## Running the project

To run the project in your machine, first run the following comand to get the project dependencies:

```
go mod download

```

After that, run the project:

```
go run .
```

### Docker

You can run the project using docker, just run the following command:

```
docker-compose up
```
