---
Title: "Build a Discord Bot with Node.js"
Description: "Step-by-step tutorial about how to build a Discord Bot with Node.js."
DatePublished: "2022-02-28"
Categories:
  - "machine-learning"
  - "developer-tools"
  - "bash"
  - "javascript"
Tags:
  - "Chatbots"
  - "Node"
CatalogContent:
  - "learn-node-js"
  - "paths/build-chatbots-with-python"
---

_**Prerequisites:** Basic understanding of Discord, JavaScript, Node.js, Visual Studio Code_  
_**Versions:** Node.js 12.20.2_

## Introduction

[Discord](https://discord.com/) is a popular instant messaging application consisting of servers and channels. Servers are synonymous with group chats. Inside of servers, users may text, voice, or video chat. Channels belong to servers and are typically named according to their purpose. For example, you may have a server named "Gaming Group" with an "#announcements" channel to post announcements for users on that server.

If you’re familiar with Discord, you may have noticed the presence of a Bot. Bots can help automate tasks such as playing music or moderating chats. 

In this article, we will set up a Discord Bot using [Node.js](https://nodejs.org/en/). Node.js will allow us to write JavaScript outside of the browser.

![Gif of Bot replying "Hello" to user](https://github.com/Codecademy/articles/blob/main/build-a-discord-bot-with-node-js/discord_bot_reply.gif?raw=true)

## Create our Discord Application

Let’s first head over to Discord [Developer Portal](https://discord.com/developers/applications). This is where we will create a new application.

![Discord Developer Portal to create new applications](https://github.com/Codecademy/articles/blob/main/build-a-discord-bot-with-node-js/discord_developer_portal.jpg?raw=true)

When we click the top right button labeled "New Application," a modal form will prompt us to create an application by first entering a name.
  
![Form modal to create an application with a name input](https://github.com/Codecademy/articles/blob/main/build-a-discord-bot-with-node-js/create_discord_app_modal.png?raw=true)

After creating our application, we’ll be brought over to the general information tab where we can customize our bot’s profile icon and description.

![Application’s general information page with a customizable icon image and an About Me section](https://github.com/Codecademy/articles/blob/main/build-a-discord-bot-with-node-js/general_info_bot.png?raw=true)

In the bot tab, we will add our bot user to the application with a click of the "Add Bot" button.  

[Application’s bot page with an add bot button](https://github.com/Codecademy/articles/blob/main/build-a-discord-bot-with-node-js/discord_dev_bot.png?raw=true)

We should see a message that says, “A wild bot has appeared!”
Our bot has a _secret token_ to share with us. Let’s copy and save it for a later step.
 
![After adding a bot to the application, a hidden authorization token is generated](https://github.com/Codecademy/articles/blob/main/build-a-discord-bot-with-node-js/a_wild_bot.png?raw=true)

Over on the settings tab, we can select the [OAuth2](https://discord.com/developers/docs/topics/oauth2) tab. This is where we can obtain the client ID and client secret to authenticate our application. 

In the scopes section near the bottom of the page, we can generate a URL to authorize our application. Let’s check off the bot box.

![OAuth2 page with the bot option checked and an auto-generated URL bar](https://github.com/Codecademy/articles/blob/main/build-a-discord-bot-with-node-js/oauth2_scopes.png?raw=true)

When we choose the bot scope, we are then prompted to check off any permissions we wish to give our bot.

![OAuth2 page bot permissions selected for a messaging bot](https://github.com/Codecademy/articles/blob/main/build-a-discord-bot-with-node-js/oauth2_bot_permissions.png?raw=true)

After selecting the desired permissions, we will copy and paste the URL into a new window or tab.

### Add Bot to Server

The URL should take us to a private Discord page where we can add our bot to an existing server. 

After selecting a server, we will follow the prompts. Once our bot is authorized and we are ready to close the window/tab, we'll hop over to the Discord server to confirm the action was a success.

![Discord message letting us know a bot has joined the server](https://github.com/Codecademy/articles/blob/main/build-a-discord-bot-with-node-js/bot_hops_into_server.png?raw=true)

## Build Discord Bot

Now that we have created our Discord Bot Application and added it to a server, we can start building out our bot’s functionality. You may use a text editor of choice; in this tutorial, we will be using [Visual Studio Code](https://code.visualstudio.com/).

### Step 1: Create Project Directory

Let’s open our terminal to where we wish our project to live and run the following commands to create our project directory and files:

```bash
$ mkdir discord-bot
$ cd discord-bot
$ touch discordbot.js .env
```

### Step 2: Add Auth Token and Node Packages

The **discordbot.js** file will hold the code for our bot’s functionality and the **.env** file will securely store the _secret token_ copied over from the previous section.

```pseudo
// .env

CLIENT_TOKEN=PasteYourTokenHere
```

Node allows us to incorporate open-source code packages in our projects via [npm](https://www.codecademy.com/resources/docs/javascript/npm). There are tons of great npm packages.

We will install two packages: [Discord.js](https://www.npmjs.com/package/discord.js) is what allows us to interact with the Discord API and [dotenv](https://www.npmjs.com/package/dotenv) allows us to load environment variables from the **.env** file we created. It’s better to use an **.env** file in this situation because we want to keep our token secure.

```bash
$ npm install discord.js dotenv
```

Our project should now have the two files we originally created in addition to the three folder/file(s) generated from the node package manager installation:

- **node_modules** folder
- **package.json** file
- **package-lock.json** file

### Step 3: Log In Bot and Add Functionality

We will now create some simple functionality for our bot. In order to do so, we need to first require and initialize the modules we installed via npm.

In **discordbot.js**:

```js
// Initialize dotenv
require('dotenv').config();

// Discord.js versions ^13.0 require us to explicitly define client intents
const { Client, Intents } = require('discord.js');
const client = new Client({ intents: [Intents.FLAGS.GUILDS, Intents.FLAGS.GUILD_MESSAGES] });

// The console log string will appear in our terminal when we first run this file
client.on('ready', () => {
 console.log(`Logged in as ${client.user.tag}!`);
});

// Log In our bot
client.login(process.env.CLIENT_TOKEN);
```

If we run **discordbot.js** in our terminal, our bot should come online in the discord server, and we should see the following message logged to the console (bot number will vary):

```bash
$ node discordbot.js
Logged in as discordBot#0000!
```

### Add Functionality to Discord Bot

Let’s set up a simple bot reply for whenever a user types "Hello". 

```js
client.login(process.env.CLIENT_TOKEN);

client.on('messageCreate', msg => {
// You can view the msg object here with console.log(msg)
 if (msg.content === 'Hello') {
   msg.reply(`Hello ${msg.author.username}`);
 }
});
```

In the newly added lines of code, the bot is listening for a message on the server. If the content of that message equals the string "Hello," our bot will reply "Hello" back with the author’s username.

Let’s re-run the file and type "Hello" into the Discord chat.

![Gif of Bot replying "Hello" to user](https://github.com/Codecademy/articles/blob/main/build-a-discord-bot-with-node-js/discord_bot_reply.gif?raw=true)

## Conclusion

We created a Discord Bot using Discord’s Developer Portal and Node.js. We used the discord.js module to interact with the Discord API and used dotenv to read **.env** files. Other [node packages/modules] can be utilized to upgrade the functionality of our bot. While the purpose of the bot we created today may be simple, the possibilities are endless! 

Here is the source code:

* https://github.com/Codecademy/articles/tree/main/build-a-discord-bot-with-node-js/discord-bot
