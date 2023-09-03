---
Title: "NextJs Deployment with Caprover and Github Actions"
Description: "A guide on how to deploy Nextjs on Caprover using Github Actions"
DatePublished: "2023-06-13"
Categories:
  - "javascript"
  - "computer-science"
Tags:
  - "GitHub"
  - "JavaScript"
CatalogContent:
  - "introduction-to-javascript"
  - "paths/front-end-engineer-career-path"
---

_**Prerequisites:** Understanding of JavaScript._
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

Caprover has a marketplace of apps that can be created with one click including postgres, supabase and more.
![caprover app marketplace](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/app-market.png)

Next, add a domain name to your app. You should login to your domain registrar and point the domain name to droplet/server's ip address. I am using a subdomain on Namecheap, here's what that looks like.
![point domain to ip](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/set-domain.png)

Now, Add a wildcard to your domain. The wildcard domain is needed so we can change the caprover server link from an ip address to a domain name.
![domain wildcard](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/root-domain.png)

Select force https and click the "save and update" button. This redirectes http traffic to https.
![force https](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/force-https.png)

Enable Https for your domain. Caprover issues lets encrypt certificate for domains.
![enable https certificate](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/enable-https.png)

Our website is live now, this is default page for caprover.
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

Navigate to the deployments tab and enable app token, the token generated will be needed for this deployment.

![github actions settings](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/enable-token.png)

Caprover uses Docker to create apps, to deploy our Nextjs app we will create a Dockerfile

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

The first stage of the build uses the node:16-alpine image as a base. This image is a lightweight version of Node.js that is optimized for production use. The `ENV NODE_ENV=production` line sets the environment variable `NODE_ENV` to production. This tells Next.js to use its production build configuration.

The second stage of the build copies the application files into the image. The `COPY package*.json ./` line copies the package.json and package-lock.json files into the image. These files are used to install the application's dependencies. The COPY ./.next ./.next line copies the built application files into the image. These files are the static files that are served by the Next.js server.

The `EXPOSE 3000` line exposes port 3000 on the image. This is the port that the Next.js server will listen on.

The `CMD ["npm", "run", "start"]` line tells Docker to run the npm start command when the container is started. This command will start the Next.js server.

Now, we need to specify the `PORT` on Caprover as 3000
![container port](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/container-port.png)

Caprover also uses a `captain-definition` file which specifies the path tothe Dockerfile. For our Nextjs app, the `captain-definition` file and the Dockerfile are to be placed at the root of our app along with the `package.json` file.
```json
{
    "schemaVersion": 2,
    "dockerfilePath": "Dockerfile"
}
```

To ensure every change we make to our nextjs app through commits automatically shows on the live website, we will need a workflow file with github actions. Create a` .github` folder and a subfolder `workflows` with `release.yaml` file in this subfolder.
The yaml file should contain the below code.

```yaml
name: Deploy to caprover instance.

env:
    CONTEXT_DIR: './'
    IMAGE_NAME: ${{ github.repository }}
    DOCKERFILE: ./Dockerfile
    # CAPROVER_APP: myapp-staging
    DOCKER_REGISTRY: ghcr.io

on:
    push:
        branches:
            - main


jobs:
    build-and-publish:
        runs-on: ubuntu-latest
        permissions:
          contents: read
          packages: write
        steps:
            - uses: actions/checkout@v1
            - name: Cache 
              uses: actions/cache@v2
              with:
                  # See here for caching with `yarn` https://github.com/actions/cache/blob/main/examples.md#node---yarn or you can leverage caching with actions/setup-node https://github.com/actions/setup-node
                  path: |
                    ~/.npm
                    ${{ github.workspace }}/.next/cache
                  # Generate a new cache whenever packages or source files change.
                  key: ${{ runner.os }}-nextjs-${{ hashFiles('**/package-lock.json') }}-${{ hashFiles('**/*.[jt]s', '**/*.[jt]sx') }}
                  # If source files changed but packages didn't, rebuild from a prior cache.
                  restore-keys: |
                    ${{ runner.os }}-nextjs-${{ hashFiles('**/package-lock.json') }}-
            - name: Use Node.js ${{ matrix.node-version }}
              uses: actions/setup-node@v3
              with:
                node-version: ${{ matrix.node-version }}
                cache: "npm"
            - run: npm ci
            - run: npm run build --if-present
            - run: npm run test --if-present
           
            - name: Log in to the Container registry
              uses: docker/login-action@f054a8b539a109f9f41c372932f1ae047eff08c9
              with:
                  registry: ${{ env.DOCKER_REGISTRY }}
                  username: ${{ github.actor }}
                  password: ${{ secrets.GITHUB_TOKEN }}

            - name: Extract metadata (tags, labels) for Docker
              id: meta
              uses: docker/metadata-action@v4
              with:
                images: ${{ env.DOCKER_REGISTRY }}/${{ env.IMAGE_NAME }}

            - name: Build and push Docker image
              uses: docker/build-push-action@v3
              with:
                context: .
                push: true
                tags: ${{ steps.meta.outputs.tags }}
                labels: ${{ steps.meta.outputs.labels }}


            - name: Deploy to CapRover
              uses: caprover/deploy-from-github@d76580d79952f6841c453bb3ed37ef452b19752c
              with:
                  server: ${{ secrets.CAPROVER_SERVER }}
                  app: ${{ secrets.APP_NAME }}
                  token: '${{ secrets.APP_TOKEN }}'
                  image: ${{ steps.meta.outputs.tags }}
```

The code provided is a GitHub Actions workflow file for deploying a Docker image to a CapRover instance. CapRover is a multi-purpose deployment tool that simplifies the process of deploying applications to your own servers.

Let's break down the workflow file:

- `env`: This section defines environment variables used in the workflow.

- `on`: This section specifies the trigger for the workflow. In this case, it runs when a push event occurs on the `main` branch.

- `jobs`: This section contains the main sequence of steps to be executed as part of the workflow.

  - `build-and-publish`: This job is responsible for building the application, publishing the Docker image, and deploying it to CapRover.

    - `runs-on`: Specifies that this job runs on the latest Ubuntu environment.

    - `permissions`: Grants necessary permissions to the job.

    - `steps`: Contains a sequence of steps to be executed.

      - `uses: actions/checkout@v1`: Checks out the repository code.

      - `name: Cache`: Configures caching to speed up subsequent builds.

      - `uses: actions/cache@v2`: Caches specified directories and files to avoid redundant operations.

      - `name: Use Node.js ${{ matrix.node-version }}`: Sets up the Node.js environment.

      - `run: npm ci`: Installs the project dependencies.

      - `run: npm run build --if-present`: Builds the project (if there is a build script defined in the `package.json` file).

      - `run: npm run test --if-present`: Runs tests (if there is a test script defined in the `package.json` file).

      - `name: Log in to the Container registry`: Logs in to the specified container registry using the provided credentials.

      - `name: Extract metadata (tags, labels) for Docker`: Retrieves metadata for the Docker image using the `docker/metadata-action` action.

      - `name: Build and push Docker image`: Builds the Docker image and pushes it to the container registry using the `docker/build-push-action` action.

      - `name: Deploy to CapRover`: Deploys the Docker image to the CapRover instance using the `caprover/deploy-from-github` action.

        - `server`: Specifies the CapRover server URL (configured as a secret).

        - `app`: Specifies the CapRover application name (configured as a secret).

        - `token`: Specifies the CapRover application token (configured as a secret).

        - `image`: Specifies the Docker image to be deployed (using the image tag from the previous step).

This workflow will trigger a build, test, and deployment whenever a push event occurs on the `main` branch. The Docker image will be built and pushed to the specified container registry, and then it will be deployed to the CapRover instance.
You will need to set up the necessary secrets in your GitHub repository to provide the CapRover server URL, application name, and application token. Make sure you have the CapRover server up and running and the required secrets configured correctly.

To configure your secrets, navigate to your Github repository's settings and click the Secrets and variables

![github actions settings](https://raw.githubusercontent.com/smyja/ugc/nextjs/content/smyja/gh-actions-settings.png)

Your `CAPROVER_SERVER` should be similar to this` https://captain.example.scrapeweb.page`
APP_NAME is `server1`, the name you specified when creating the app.
`APP_TOKEN` is the Token generated when we enabled app token on the dashboard.

## Conclusion

We’ve now learned how to deploy a NextJS app with Caprover and connect a domain to it.

Source code: https://github.com/smyja/nextapp
