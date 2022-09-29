# “For Woman, freedom, life”

## What is our goal?
Our goal in creating this project is to translate the collection of Farsi “for…” tweets published with the #mahsaAmini hashtags into English and make it known to the world.
This movement is a part of the protests formed in Iran, which were formed by the people to stand for the killing of Mahsa Amini by the morality police and dissatisfaction with the government of the Islamic Republic due to the numerous political, social, and economic problems such as violating human rights.

## What are the series of “for…” tweets, how was it formed, and what is the purpose of writing them?

The tweets of “for…” tell all the regrets, desires, anger, etc… of the Iranian people in the last 44 years (current regime).
These tweets start with the word “for” and then people express their desires, regrets, and wishes that they could not fulfill due to the rules and conditions of the dictatorship in the society (e.g., “For Woman, freedom, life”).
Hoping for a day when all people can live freely and fulfill their dreams.

## Executing Twitter Scraper
```bash
go mod tidy
go get main.go
```

## Setup Web UI

Change working directory to ./web-ui
```bash
cd web-ui
```

Make sure to install the dependencies:

```bash
# yarn
yarn install

# npm
npm install

# pnpm
pnpm install --shamefully-hoist
```

## Development Server

Start the development server on http://localhost:3000

```bash
npm run dev
```

## Production

Build the application for production:

```bash
npm run build
```

Locally preview production build:

```bash
npm run preview
```
