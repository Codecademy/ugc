---
Title: "Building a ToDo List with Ruby on Rails"
Description: "Learn how to build a ToDo List with Ruby on Rails"
DatePublished: "2023-07-07"
Categories:
  - "computer-science"
  - "web-development"
Tags:
  - "Ruby on Rails"
  - "Ruby"
  - "PostgreSQL"
  - "Database"
CatalogContent:
  - "learn-ruby"
  - "learn-rails"
---

_**Prerequisites:** Ruby, Ruby on Rails, PostgreSQL_

_**Versions:** Ruby 3.2.0, Ruby on Rails 7.0.4, PostgreSQL 15_

## Introduction

Have you ever felt lost while learning something new? I know exactly how you feel. When I started my journey with Ruby on Rails, I found myself constantly searching for information in the official Rails guide, only to get overwhelmed and confused. But fear not! This article is here to guide you through the process of building a full-stack app using Ruby on Rails, with a specific focus on creating a To-do list.

As someone who has experienced the frustration of trying to piece together information from various sources, I understand the importance of having a clear and beginner-friendly resource that can help you get started with building your own Rails app. That's why I was inspired to write this article: to give you a detailed guide that will save you time and enable you confidently start your own Rails journey.

Whether you're a complete beginner or have some prior programming experience, this article is designed to be accessible and informative, making it easier for you to grasp the concepts and techniques involved.

If you're ready to dive into the exciting world of Ruby on Rails and build your own To-do list app, let's get started! By the end of this article, you'll have a solid foundation in Rails development and the confidence to continue exploring and creating your own amazing applications.

## What I’ll talk about:

In this article, we will take a look at building a full-stack app using Ruby on Rails. Throughout the course of the article, we'll cover the following main points:

1. Setting up a new Rails app.
2. Understanding the MVC design pattern.
3. Creating a new schema in your local database using migrations.
4. Handling requests within controller methods.
5. Using ERB to write logic and display content in HTML.
6. Understanding the Rails routing system.

## Brief overview of Ruby on Rails

Rails is a web application development framework written in the Ruby programming language. It is an easy framework to get started with because it allows you to write less code while accomplishing more than many other languages and frameworks.

The Rails framework is guided by two major principles:

1. Don’t Repeat Yourself: The [DRY principle](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself) is very common in software development. The idea is that developers should spend less time writing code by avoiding repetition. This is done by making sure every element of your code can stand alone and be called anywhere, and changes to that element don't have to be repeated across several points.
2. Convention over Configuration: This is one of Rails’ superpowers. Rails has a set of guidelines and format for doing things. From naming methods, classes, files, and building a database schema. These conventions help to keep the whole application in sync and help you spend less time debugging errors that result from the wrong configuration.

## Setting up your dev environment: Installing Ruby, Rails, and PostgreSQL

[GoRails](https://goRails.com/) Has in-depth guides on setting up Rails on various machines. Here's a link to all of them. Since we're using PostgreSQL for this project, skip any step that mentions setting up SQLite or MySQL.

- [For Ubuntu](https://goRails.com/setup/ubuntu/22.04)
- [Mac OS](https://goRails.com/setup/macos/13-ventura)
- [For Windows](https://www.hanselman.com/blog/ruby-on-Rails-on-windows-is-not-just-possible-its-fabulous-using-wsl2-and-vs-code)

## Setting up a new Rails app

To create a new Rails app, simply open up your terminal, navigate to where you want this app to stay, and run the following commands.

```bash
rails new todolist -d postgresql

cd todolist/

bundle add tailwindcss-rails

rails tailwindcss:install
```

The commands after the app name (**todolist**) are optional. The `-d postgresql` flag tells rails to use postgreSQL as the database for this app. 
`bundle add tailwindcss-rails` adds tailwindcss to our Gemfile. Tailwind is a utility-first CSS framework and a great tool to help you build a responsive and beautiful UI. 
`rails tailwindcss:install` installs tailwindcss in our app and sets up a Procfile for running our app in development.

After running these commands, you’ll see a bunch of files and folders created for us. This is the default structure of a rails app. You can read more about it [here](https://guides.rubyonrails.org/getting_started.html#creating-the-blog-application).

To get our app running, we need to make sure it is connected to a database. The database config lives in the `config/database.yml` file. There we can see the provided database name and some other info. This database doesn't exist yet and needs to be created. For most people, you need to supply your PostgreSQL username and password that was set during installation to this file. The default username is `postgres` and this applies to most Linux distros.

Most times we would use a **.env** file to manage this but we can keep it there for the scope of this project.

```yml
...
default: &default
  adapter: postgresql
  encoding: unicode
  username: postgres
  password: yourpassword
...
```

Make sure your terminals' current directory is the root of your app before running the following commands.

If you are using a Linux machine, you need to start your postgreSQL service before proceeding. Run the following command to start the service:

```bash
sudo service postgresql start
```

To create the local database for our app, run: 

```bash 
rails db:create
```

Then run `rails s` to start the Rails server. If everything went well, you should see a URL to visit (`http://127.0.0.1:3000/`) in your terminal. Open that URL in your browser and you should see the default rails page.

Since we have TailwindCSS in our project, starting our server requires an additional step. We need to run `rails tailwindcss:watch` in a separate terminal window pointing at our app directory. This will compile our tailwind styles and make them available to our app.

To make it easier, Rails ties both the server and the tailwind compiler together into a single command. To run both at the same time, run `bin/dev`.

## Understanding the MVC design pattern

MVC stands for Model, View, Controller. It is a design pattern that helps to organize the code in a Rails app. It is a common pattern in web development and is used in many other frameworks. The MVC pattern is used to separate the code into three different layers. Each layer has a specific responsibility and is responsible for handling a specific part of the app. We'll look at the three layers closely as we build our app.

## Creating a new schema in the local database using migrations.

Migrations help us build/alter our database schema in a consistent way. It uses the Ruby DSL (domain-specific language), and the dedicated rails ORM (Object-relational mapping) called Active Record so we don't have to write any SQL by hand.

Our app requires a table and a column to store our todo items. We'll call that column `description`.
To create this table with the column, we'll run:

```bash
rails generate model Todo description
```

There are a few things to unpack about this:
1. `rails generate`: Compulsory prefix for generating various resources like models, controllers, etc.
2. `model`: This tells rails that we intend to create a model. Creating a model will create a migration by default. Models represent the M in MVC and are part of the application that handles the business logic. How data is represented, accessed, and modified can be specified in the model. The model interacts with the database to carry out several operations.
3. `Todo`: This is the name of the model we want to create. Rails will automatically pluralize this to create a table named `todos`.
4. `description`: This is the column name we want our table to have. Columns are usually created with a default string datatype. We could do `clicks:integer` to add a column with an integer datatype instead.

After running this command, you'll see a new file in the `db/migrate` folder. This file contains the migration code that will be used to create the table and the column in our database. The migration file name is in the format `yyyymmddhhmmss_create_todos.rb`. The first part of the file name is the timestamp of when the migration was created. This is used to keep track of the order in which migrations were created. The second part is the name is the kind of migration we want to create.

To apply the schema to our local database, we'll run:

```bash
rails db:migrate
```

A model file will also be created in the `app/models` folder. This file contains the model class and is used to define the model's attributes and relationships with other models.

## Handling requests within controller methods.

Controllers are responsible for handling requests and sending responses. They are the middleman between the user and the model. They receive requests from the user and send responses back to the user. They also interact with the model to get the data it needs to perform its operations.

Let's create a controller for our todo items. In the terminal, run:

```bash 
rails generate controller todos
```

This will create a controller file in the `app/controllers` folder. It will also create a folder named `todos` in the `app/views` folder. This folder will contain all the views for our todo items.

To display a simple `'Hello World'` message, we'll create a new method in the todos controller, and specify routes for it. Open the `app/controllers/todos_controller.rb` file and add the following code.

```ruby
class TodosController < ApplicationController
  def index
  end
end
```

This is a simple controller method that will render the index view when a request is made to the `/todos` route. The index view is the default view for a controller. Rails will look for a view with the same name as the controller method. In this case, it will look for a view named `index.html.erb` in the `app/views/todos` folder.

To create this view, create a new file named `index.html.erb` in the `app/views/todos` folder. Add the following code to the file.

```html
<h1>Hello World</h1>
```

Then add the following code to the `config/routes.rb` file.

```ruby
Rails.application.routes.draw do
  get '/todos', to: 'todos#index'
end
```

This will create a route that maps the `/todos` URL to the `todos#index` controller method. The `get` method specifies the HTTP verb to use. The `to` option specifies the controller and method to use. The `todos#index` syntax is a shorthand for `controller: 'todos', action: 'index'`.

Now run `rails s` or `bin/dev` to start the server, and visit `http://localhost:3000/todos` in your browser. You should see the '**Hello World**' message.

## Using ERB to write logic and display content in HTML.

ERB stands for Embedded Ruby. It is a templating system that allows us to write ruby code within HTML files. This is useful because it allows us to write logic and display content in HTML. We can also use it to write HTML in a more concise way. 
This aspect of Rails is the V in MVC. The view layer is responsible for displaying the content to the user. It is the part of the application that handles the presentation logic. The view layer interacts with the controller to get the data it needs to display. Without controllers, views would not be able to display any data.

On our webpage, we want to have a simple form for creating new todo items, and a table to list all the todo items. We'll use ERB to write the logic for this.

Open the `app/views/todos/index.html.erb` file and replace its contents with the following code.

```erb
<h1 class="text-2xl">Todo List</h1>

<%= form_with(model: Todo.new, class: "my-10") do |form| %>
  <div class="my-5">
    <%= form.label :description %>
    <%= form.text_field :description, placeholder: "new todo", class: "block shadow rounded-md border border-gray-200 outline-none px-3 py-2 mt-2 w-2/4" %>
  </div>

  <div class="inline">
    <%= form.submit class: "rounded-lg py-3 px-5 bg-blue-600 text-white inline-block font-medium cursor-pointer" %>
  </div>
<% end %>
```

This code will render a form with a text field and a submit button. The `form_with` method is a helper method that creates a form for a given model. The `model` option specifies the model we are creating a record for.
The `class` option specifies the CSS class to use for the form. The `do |form|` block specifies the content of the form. The `form.label` method creates a label for the text field. The `form.text_field` method creates a text field for the form. The `form.submit` method creates a submit button for the form.

To view this on our webpage. Run `rails s` or `bin/dev` to start the server, and visit `http://localhost:3000/todos` in your browser. You should see the form we just created.

We can't create todo items with this form yet. We need to add a route to handle the form submission. Open the `config/routes.rb` file and replace what you have with the following code.

```ruby
Rails.application.routes.draw do
  resources :todos, only: [:index, :create]
end
```

## Understanding the Rails routing system

The `resources` method in the code block above creates a set of routes for a given resource. The `only` option specifies the controller methods to create routes for. In our case, we only want to create routes for the `index` and `create` methods. The `resources` method creates the following routes:

| HTTP Verb | Path   | Controller#Action | Used for                    |
|-----------|--------|-------------------|-----------------------------|
| GET       | /todos | todos#index       | display a list of all todos |
| POST      | /todos | todos#create      | create a new todo           |

Without passing the `only` option, the `resources` method would create routes for all the controller methods. The `resources` method creates the following routes.

| HTTP Verb | Path   | Controller#Action | Used for                    |
|-----------|--------|-------------------|-----------------------------|
| GET       | /todos | todos#index       | display a list of all todos |
| GET       | /todos/new | todos#new | return an HTML form for creating a new todo |
| POST      | /todos | todos#create      | create a new todo           |
| GET       | /todos/:id | todos#show | display a specific todo |
| GET       | /todos/:id/edit | todos#edit | return an HTML form for editing a todo |
| PATCH/PUT | /todos/:id | todos#update | update a specific todo |
| DELETE    | /todos/:id | todos#destroy | delete a specific todo |

## Creating a new todo item

Now that we have our routes setup, we need add the corresponding method in our controller with some logic to handle the form submission. Open the `app/controllers/todos_controller.rb` file and replace what you have with the following code.

```ruby
class TodosController < ApplicationController
  def index
  end

  def create
    @todo = Todo.create(description: params[:todo][:description])
    if @todo.valid?
      redirect_to todos_path 
    end
  end
end
```

The `create` method creates a new instance of the Todo model and saves it to the database. The `description` attribute is set to the value of the `description` field in the form. The `params` method returns a hash of the form data. The hash looks like this:

```ruby
{
  todo: {
    description: "new todo"
  }
}
```

The `params[:todo][:description]` syntax gets the value of the hash. The `valid?` method checks if the model is valid. If the model is valid it renders the `index` view and show the list of todos. If the model is not valid, it just renders the `index` view. In a real application, we would want to display the errors to the user.

You can submit the form multiple time to have enough data to test the table. We'll add the table in the next section. You can verify that your todos are being saved to the database by opening your terminal to your apps directory and running `rails c` to open the rails console. Then run `Todo.all` to see the list of todos.

## Displaying the list of todos

In order to display the list of todos, we need to get the list of todos from the database. Open the `app/controllers/todos_controller.rb` file and add the following code to the `index` method.

```ruby
...
def index
  @todos = Todo.all
end
...
```

The `Todo.all` method returns an array of all the Todo models in the database. The `@todos` variable is now available in the view.

In our view, we need to add a table to display the list of todos. Open the `app/views/todos/index.html.erb` file and append the following code to it.

```erb
...
<table class="w-3/5 mt-10 border-separate table-auto text-slate-500">
  <thead>
    <tr>
      <th class="p-3 text-sm font-semibold text-left text-gray-700 bg-gray-200">ID</th>
      <th class="p-3 text-sm font-semibold text-left text-gray-700 bg-gray-200">Description</th>
    </tr>
  </thead>
  <tbody>
    <% @todos.each do |todo| %>
      <tr>
        <td class="px-3 py-4 text-sm font-normal text-left text-gray-600 bg-gray-50">
          <%= todo.id %>
        </td>
        <td class="px-3 py-4 text-sm font-normal text-left text-gray-600 bg-gray-50">
          <%= todo.description %>
        </td>
      </tr>
    <% end %>
  </tbody>
</table>
```

This code will render a table with the list of todos. The `@todos` variable is an array of Todo models. The `each` method iterates over the array and renders a row for each todo. The `todo.id` and `todo.description` syntax gets the id and description of the todo.


Once you have the code above, you should see a table with the list of todos previous created on the browser.

## Adding a delete button

First we need to add a `destroy` method will handle the delete request. 

Open the `app/controllers/todos_controller.rb` file and add the following code to create `destroy` method.

```ruby
...
  def destroy
    @todo = Todo.find(params[:id])
    if @todo.destroy
      redirect_to todos_path
    end
  end
```

The `Todo.find` method finds a todo by the id (Every todo gets a unique id when it's created). The `destroy` method deletes the todo from the database. The `redirect_to` method redirects the user to the `todos_path` which is the index page. This will reload the page and show the updated list of todos.


In our view, we'll add a delete button to each row in the table. Open the `app/views/todos/index.html.erb` and add an extra column to the table head and table body with the following code.

```erb
...
  <thead>
    <tr>
      <th class="p-3 text-sm font-semibold text-left text-gray-700 bg-gray-200">ID</th>
      <th class="p-3 text-sm font-semibold text-left text-gray-700 bg-gray-200">Description</th>
      <th class="p-3 text-sm font-semibold text-left text-gray-700 bg-gray-200">Action</th>
    </tr>
  </thead>

  
  <tbody>
    <% @todos.each do |todo| %>
      <tr>
        <td class="px-3 py-4 text-sm font-normal text-left text-gray-600 bg-gray-50">
          <%= todo.id %>
        </td>
        <td class="px-3 py-4 text-sm font-normal text-left text-gray-600 bg-gray-50">
          <%= todo.description %>
        </td>
        <td class="px-3 py-4 text-sm font-normal text-left text-red-600 underline break-words bg-light">
          <%= button_to "Delete", todo_path(todo), method: :delete %>
        </td>
      </tr>
    <% end %>
  </tbody>
```

The `button_to` method creates a form with a button. The `method: :delete` option specifies that the form should be submitted using the `DELETE` HTTP verb. The `todo_path(todo)` syntax generates the path for the todo. The `todo_path` method takes the todo model as an argument.

Finally, we need to add a route to handle the delete request. Open the `config/routes.rb` file and update the todos resource with the following code.

```ruby
...
  resources :todos, only: [:index, :create, :destroy]
```
Now, you should be able to delete a todo item from the list.

That is all. We’ve built a working todo list with READ, CREATE and DESTROY methods.

You can find the working code on this [Github repo](https://github.com/deyemiobaa/Rails_todolist).
