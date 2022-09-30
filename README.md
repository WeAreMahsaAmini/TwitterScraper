<p align="center">     
    <a href="https://github.com/WeAreMahsaAmini/TwitterScraper/tree/main">
        <img src="https://img.shields.io/circleci/project/github/badges/shields/master?style=plastic" alt="build status">
    </a>
    <a href="https://github.com/WeAreMahsaAmini/TwitterScraper/tree/main">
        <img src="https://img.shields.io/badge/Go-00ADD8?style=plastic&logo=go&logoColor=white" alt="Go Lang">
    </a href="https://go.dev">
    <a href="https://vuejs.org">
        <img src="https://img.shields.io/badge/Vue.js-1a1a1a?style=plastic&logo=Vue.js&logoColor=4FC08D" alt="Vue js">
    </a>
    <a href="https://v3.nuxtjs.org">
        <img src="https://img.shields.io/badge/Nuxt.js-black?style=plastic&logo=Nuxt.js&logoColor=00DC82" alt="Vue js">
    </a>
    <a href="https://github.com/WeAreMahsaAmini/TwitterScraper/blob/main/LICENSE">
        <img src="https://img.shields.io/github/license/WeAreMahsaAmini/TwitterScraper?style=plastic" alt="License Apache 2">
    </a>
    <br>
    <a href="/">
        <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/WeAreMahsaAmini/TwitterScraper?style=social">
    </a>
    <a href="https://discord.gg/JJsKx9de2a">
        <img src="https://img.shields.io/discord/1024073094366769283?logo=discord&&style=plastic" alt="chat on Discord">
    </a>
    <a href="https://github.com/WeAreMahsaAmini">
        <img alt="GitHub followers" src="https://img.shields.io/github/followers/WeAreMahsaAmini?color=black&style=social">
    </a>
</p>

# “For Woman, freedom, life”
## Live version
Go to  <a href="https://forliberty.info" target="_blank">`ForLiberty.info`</a>
> Still under developement but you can see updates here
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
