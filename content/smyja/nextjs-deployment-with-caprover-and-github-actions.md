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
Sign up on [Digitalocean](https://digitalocean.com) if you don't have an account yet. 
Once you have signed up, create a droplet.

![created caprover droplet](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/droplets.png)

Next, select caprover from the marketplace.

![created caprover droplet](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/caprover-marketplace.png)


Choose a region, specify the CPU as 1GB RAM, you can always upgrade it to a higher RAM.
![Specify CPU](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/cpu-choice.png)

Once the droplet has been created, you can access the caprover dashboard by visiting http://YOUR_IP_ADDRESS:3000 or click the `Get started` link ,then the Quick access to Caprover console.

![created caprover droplet](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/created-caprover-droplet.png)

The quick access button can be found here.
![caprover quick access](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/caprover-quick-access.png)

The default login for caprover apps is `captain42`, ensure you change it from your dashboard settings.

After logging in, create an app. Give it any name.
![create caprover app](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/caprover-app.png)

Next, add a domain name to your app. You should login to your domain registrar and point the domain name to droplet/server's ip address. I am using a subdomain on Namecheap, here's what that looks like.
![point domain to ip](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/set-domain.png)

Now, Add a wildcard to your domain. The wildcard domain is needed so we can change the caprover server link from an ip address to a domain name.
![domain wildcard](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/root-domain.png)

Select force https and click the "save and update" button. This redirectes http traffic to https.
![force https](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/force-https.png)

Enable Https for your domain. Caprover issues lets encrypt certificate for domains.
![enable https certificate](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/enable-https.png)

Our website is live now.
![created caprover droplet](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/website.png)

### Setting up a NextJs app
Follow the guide on [Nextjs Docs](https://nextjs.org/docs/getting-started/installation) to create a NextJs app. 
Once that's setup, we will create a cluster for the app. 
![remote registry](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/remote-registry.png)

To add a private Docker registry to CapRover, you will need to provide your username, personal access token(begins with ghp_), domain, and image prefix. We will be using the GitHub Container Registry (ghcr.io), your username would be your GitHub username, your password would be a personal token that you create with read package access, your domain would be ghcr.io, and your image prefix would be your GitHub username.

If your Docker images are stored in the format `your-username/your-image`, then you should use your GitHub username as your image prefix. Otherwise, if your images are stored in the format `my-org/my-image`, where `my-org` is your GitHub organization, then you should use `my-org` as your image prefix.

Once you have provided these credentials, CapRover will be able to pull images from your private Docker registry.
Your created registry would show
![created caprover droplet](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/docker-registeries.png)


```docker 
FROM node:16-alpine

# Set working directory

ENV NODE_ENV=production
# Copy package.json and package-lock.json (if available)
COPY package*.json ./


# Copy the built application files

COPY ./.next ./.next
COPY ./next.config.js ./next.config.js
COPY ./public ./public
COPY ./.next/static ./_next/static
COPY ./node_modules ./node_modules
# Expose the desired port (e.g., 3000)

EXPOSE 3000

# Start the Node.js server
CMD ["npm", "run", "start"]
```

`CAPROVER_SERVER` is https://captain.example.scrapeweb.page