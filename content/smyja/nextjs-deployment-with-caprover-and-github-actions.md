---
Title: "NextJs Deployment with Caprover and Github Actions"
Description: "A guide on how to deploy Nextjs on Caprover using Github Actions"
DatePublished: "2023-06-13"
Categories:
  - "javascript"
  - "computer-science"
Tags:
  - "javascript"
CatalogContent:
  - ""

---

_**Prerequisites:** Understanding of Javascript._  
_**Versions:** Node v18.12.1_

### Introduction
Next.js is an open-source React framework that enables you to create server-rendered, static generated, and hybrid web applications. It provides a number of features that make it easy to build high-performance web applications, including:

- Static site generation: Next.js can be used to generate static websites that are served directly from the browser. This can improve performance and SEO.
- Server-side rendering: Next.js can also be used to render pages on the server. This can improve performance for search engines and users who are not using JavaScript.
- Hybrid rendering: Next.js can also be used to combine server-side rendering and static site generation. This gives you the best of both worlds: performance and SEO.
- Automatic routing: Next.js provides automatic routing that makes it easy to create complex web applications with nested routes.
- Data fetching: Next.js provides a number of ways to fetch data from APIs, including asynchronous async/await and the fetch() API.
- Internationalization: Next.js supports internationalization, so you can easily create web applications that support multiple languages.
- Deployment: Next.js can be deployed to a variety of hosting services, including Vercel, Netlify, and AWS Amplify, DigitalOcean.

This tutorial will cover the creation of a NextJs app and Deploying it on Caprover(Opensource Platform as a service) using DigitalOcean and Github actions.

## What is Caprover?
CapRover is a free and open-source platform that simplifies the deployment and management of applications. It supports a wide range of programming languages, databases, and web servers, making it a versatile solution for developers of all levels. CapRover is also a cost-effective alternative to other popular platforms, such as Heroku and Microsoft Azure.

## What is Github actions
GitHub Actions is a tool that allows you to automate tasks and processes in your GitHub repository. You can define workflows as code, which means that you can use GitHub Actions to automate anything that you can do with code.

Workflows are triggered by events, such as code pushes, pull requests, or scheduled intervals. When a workflow is triggered, it will run a series of steps that you have defined. These steps can be used to perform any task that you need to automate, such as building your code, running tests, or deploying your application.

## Why use Next.js, Caprover, and Github Actions together?
Next.js, Caprover, and GitHub Actions are all powerful tools that can be used together to create and deploy web applications. Indiehackers often use these tools because they are:
  - Efficient: These tools can help indiehackers save time and resources by automating tasks and making it easy to deploy applications.
  - Scalable: These tools can be scaled to handle large traffic loads.
  - Cost-effective: These tools are often free or low-cost, making them a good option for indiehackers with limited budgets.

### Setting up Caprover on Digital Ocean
Sign up on Digitalocean if you don't have an account yet. Once you have signed up, create a droplet.
Choose a region, specify the CPU as 1GB RAM, you can always upgrade it to a higher RAM.
Once the droplet has been created, you can access the caprover dashboard by visiting http://YOUR_IP_ADDRESS:3000 or click the `Get started` link ,then the Quick access to Caprover console.
The default login for caprover apps is `captain42`, ensure you change it from your dashboard settings.
After logging in, create an app. Give it any name.
Next, add a domain name to your app. You should login your registrar and point the domain name to server's ip address
Now, Add a wildcard to your subdomain. The wildcard domain is needed so we can change the caprover server link from an ip address to a domain name.
Select force https and click the "save and update" button. This redirectes http traffic to https.
Enable Https for your domain. Caprover issues lets encrypt certificate for domains.
Our website is live now.

### Setting up a NextJs app

