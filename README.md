# VueJS Portfolio Website

This is my personal portfolio website built with **VueJS**. It showcases my projects, skills, and experience in a clean, modern design.

### Status:
[![Deploy](https://github.com/ivantomic77/VuePortfolioSite/actions/workflows/main.yml/badge.svg?branch=master)](https://github.com/ivantomic77/VuePortfolioSite/actions/workflows/main.yml)

## Components

- **Frontend:** The entire website is built using **VueJS**

The website is deployed on **Kubernetes** and the container images are stored in the **GitHub Container Registry**. Rolling updates are used to keep the application available during updates. All of this is achieved through Github Actions to speed up the process.

## Project Setup

To get started with the development environment, run the following command:

```sh
npm install
```

Run application in developent mode:
```sh
npm run dev
```

Build the application using docker (for local testing purposes):
```sh
npm run build
docker build . -t ivantomic77/portfolio-website:latest
```
