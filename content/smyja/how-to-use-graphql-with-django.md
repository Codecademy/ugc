---
Title: "How To Use GraphQL With Django"
Description: "A guide on how to use GraphQL with Django."
DatePublished: "2022-07-18"
Categories:
  - "python"
  - "web-development"
  - "computer-science"
Tags:
  - "APIs"
  - "Django"
  - "Queries"
CatalogContent:
  - "python-for-programmers"
  - "paths/full-stack-engineer-career-path"
---

_**Prerequisites:** Understanding of Python, the Command Line, and Django._  
_**Versions:** Django 4.0.4, Python 3.8.10, virtualenv 20.15.1_

### Introduction

[GraphQL](https://www.codecademy.com/resources/docs/general/graphql) is a query language for [APIs](https://www.codecademy.com/resources/docs/general/api) and a runtime for fulfilling those queries with existing data. Unlike a REST API, GraphQL APIs do not require verbs (`PUT`, `POST`, `GET`, `PATCH`, and `DELETE`) for requests, nor do they need multiple endpoints. They have just one endpoint and making a query to that endpoint is all that's needed. 

This tutorial will cover the creation of a CRUD (create, read, update, and delete) GraphQL API with Django providing a list of restaurants.

## Properties of GraphQL

The following terms are often used when interacting with GraphQL. Knowing them can be helpful, though we won't be covering all of them in this tutorial.

- **Schema**: Describes the functionality available to the client applications that connect to it.
- **Query**: A schema type that represents the `GET` request and defines the operations that can be used for reading or fetching data.
- **Nesting**: Queries can be nested inside of other queries.
- **Mutation**: A schema type that defines the kind of operations that can be done to modify data.
- **Subscription**: Notifies the client server in real time about updates to the data.
- **Resolver**: Functions that return values for fields associated with existing schema types.

## Step 1: setting up our virtual environment

First, we are going to create and launch a virtual environment for our project with the `virtualenv` package (which can be [installed via `pip`](https://virtualenv.pypa.io/en/latest/installation.html#via-pip). While not necessary for starting a new Django project, working in separate environments is generally a best practice that mitigates conflicts between sites. Let's open a [terminal](https://www.codecademy.com/resources/docs/general/terminal) and create a new environment named `my_env` by running the following:

```bash
virtualenv my_env
```

Next, we will activate our new environment `my_env` with either of the following commands:

```bash
# Linux/macOS
source my_env/bin/activate

# Windows
source my_env/scripts/activate
```

Let's go to the next step.

## Step 2: creating our Django project

Next, if we haven't already, let's [install](https://docs.djangoproject.com/en/4.0/topics/install/#how-to-install-django) the `Django` package. 

Once we've done that, let's create a new project called `restaurant_graphql_api` and change into it:

```bash
django-admin startproject restaurant_graphql_api
cd restaurant_graphql_api
```

Next, we're going to create a new application within our project called `my_app` by running the following:

```bash
python manage.py startapp my_app
```

Then, we'll add `my_app` to our list of `INSTALLED_APPS` in our `settings.py` file under the `restaurant-graphql_api/` directory:

```py
INSTALLED_APPS = [
  'my_app',
  'django.contrib.admin',
  'django.contrib.auth',
  # ...
]
```

## Step 3: using GraphQL with `graphene-django`

To use GraphQL with Django, we will need to install the [`graphene-django`](https://docs.graphene-python.org/projects/django/en/latest/) package.

```bash
pip install graphene-django
```

This will add GraphQL functionality to our restaurant Django app such as resolvers and mutations. Next, let's add `'graphene_django'` to the list of `INSTALLED_APPS` in our `settings.py` file:

```py
INSTALLED_APPS = [
  'graphene_django',
  'my_app',
  'django.contrib.admin',
  # ...
]
```

Now, let's go to the `models.py` file in our project and then define a new `Restaurant` class:

```py
from django.db import models

class Restaurant(models.Model):
  name = models.CharField(max_length=100)
  address = models.CharField(max_length=200)

  def __str__(self):
    return self.name
```

Inside the `Restaurant` class model above, we've defined a few fields, `name` and `address`, along with a `__str__()` [dunder method](https://www.codecademy.com/resources/docs/python/dunder-methods) that returns the `name` of the restaurant.

Next, let's register our new `Restaurant` model in the `admin.py` file of our application:

```py
from django.contrib import admin
from . import models

admin.site.register(models.Restaurant)
```

It is now time to create and perform a migration for this new data. This will allow our `Restaurant` model to be referenced in a GraphQL schema (which we will define later on). To make the migration, we can run `python manage.py makemigrations`; to apply the migrations, let's run `python manage.py migrate`.

By now, we may encounter the following error:

```shell
ImportError: cannot import name 'force_text' from 'django.utils.encoding'
```

The `ImportError` is due to Django 4.0 not supporting the `force_text` variable (which the `graphene` package uses with earlier versions of Django). To resolve this, we can add the following to our `settings.py` file:

```py
import django
from django.utils.encoding import force_str

django.utils.encoding.force_text = force_str
```

Alternatively, we can downgrade our Django version to 3.2.x. 

After this, it would be good to run `python manage.py runserver` and check `http://127.0.0.1:8000` on a browser to ensure our application starts properly.

Let's now create a `urls.py` in the `my_app` directory (for our _application_, not our overall Django project) and add the following:

```py
from graphene_django.views import GraphQLView
from django.views.decorators.csrf import csrf_exempt
from django.urls import path

urlpatterns = [
  path("graphql", csrf_exempt(GraphQLView.as_view(graphiql=True))),
]
```

With the help of `import` statements, we added a `"graphql"` route to our list of `urlpatterns` that will automatically open the [GraphiQL](https://graphiql-test.netlify.app/typedoc/) API browser for testing our queries and mutations. This is done with the `graphiql` parameter of the `GraphQLView.as_view()` method. However, it can be switched off by setting `graphiql` to `False`. Django's `csrf_exempt` decorator is used to allow API clients to POST to the graphql endpoint we have created.

Next, let's import the `include()` function to add the app urls to our `restaurants_graphql_api/urls.py` file (for our entire Django _project_):

```py
from django.urls import path, include

urlpatterns = [
    path("admin/", admin.site.urls),
    path("", include("my_app.urls")),
]
```

## Step 4: building a GraphQL schema

Let's create a new file in our `my_app` directory called `schema.py`. Inside, we'll define a new type for the `Restaurant` model we previously created:

```py
import graphene
from graphene_django import DjangoObjectType
from my_app.models import Restaurant

class RestaurantType(DjangoObjectType):
  class Meta:
    model = Restaurant
    fields = ("id", "name", "address")
```

Our `RestaurantType` class borrows from the `DjangoObjectType` class. The inner-`Meta` class is where general type attributes like `model` and `fields` are defined.

Next, let's create a `Query` type class for the `Restaurant` model:

```py
class Query(graphene.ObjectType):
  """
  Queries for the Restaurant model
  """
  restaurants = graphene.List(RestaurantType)

  def resolve_restaurants(self, info, **kwargs):
    return Restaurant.objects.all()
```

The `Query` type contains a resolver function for the `restaurants` field (e.g., `resolve_restaurants()`). This resolver returns all the restaurants in the database.

Next, at the end of our `schema.py` file, we will pass in our `Query` type into the `graphene.Schema()` function. This will allow our schema to be exportable to other files:
 
```py
schema = graphene.Schema(query=Query)
```

The entire `schema.py` file should look like this:

```py
import graphene
from graphene_django import DjangoObjectType
from my_app.models import Restaurant

class RestaurantType(DjangoObjectType):
  class Meta:
      model = Restaurant
      fields = ("id", "name", "address")

class Query(graphene.ObjectType):
  """
  Queries for the Restaurant model
  """
  restaurants = graphene.List(RestaurantType)

  def resolve_restaurants(self, info, **kwargs):
    return Restaurant.objects.all()


schema = graphene.Schema(query=Query)
```

Let's now import the `schema` variable into the `my_app/urls.py` file and pass it to the Graphql view as seen below:

```py
from my_app.schema import schema

url_patterns = [
  path("graphql", csrf_exempt(GraphQLView.as_view(graphiql=True, schema=schema))),
]
```

Let's run the Django server with `python manage.py runserver` then visit the `/graphql` route to see the GraphiQL browser, which should look like this:

![API browser](https://raw.githubusercontent.com/Smyja/ugc/grahql-with-django/content/smyja/api-browser.png)

Let's quickly test our query by doing the following:

1. Create a superuser account by running `python manage.py createsuperuser` in the terminal window, and following the prompts to create a username and password.
2. Log into our application as an admin by visiting the `"/admin"` URL in the browser.
3. Add restaurants to the database by interacting with the admin dashboard.

To get the list of restaurants with specific data like `name` and `address`, we can type and run the following query on the browser:

```graphql
query {
  restaurants {
    id
    name
    address
  }
}
```

The output should look like this:

![list of restaurants](https://raw.githubusercontent.com/Smyja/ugc/grahql-with-django/content/smyja/restaurant-list.png)

## Step 5: mutating the database

To modify any data in our GraphQL database we need to create a mutation. In this step, we're going to build three mutations for creating, updating, and deleting data in our database.

Below is the `CreateRestaurant` mutation, which we will add to the `schema.py` file:

```py
class CreateRestaurant(graphene.Mutation):
  class Arguments:
    name = graphene.String()
    address = graphene.String()

  ok = graphene.Boolean() 
  restaurant = graphene.Field(RestaurantType)

  def mutate(self, info, name, address):
    restaurant = Restaurant(name=name, address=address)
    restaurant.save()
    return CreateRestaurant(ok=True, restaurant=restaurant)
```

The `CreateRestaurant` mutation takes in the model fields as arguments within the inner-`Argument` class. The `mutate()` function is where the database change happens using Django's object-relational mapper (ORM).

Next, let's create a `Mutation` class and initialize it with the schema at the end of the file:

```py
class Mutation(graphene.ObjectType):
  create_restaurant = CreateRestaurant.Field()

```

After adding the mutation, let's pass the mutation to the schema at the end of the `schema.py` file.

```py
schema = graphene.Schema(query=Query, mutation=Mutation)
```

Start the server and run a mutation with the GraphQL API browser by using this:

```graphql
mutation {
  createRestaurant(name: "Kada Plaza", address: "Lekki GARDENS") {
    ok
    restaurant {
        id
        name
        address
    }
  }
}
```

The mutation returns a restaurant object with the fields that were passed in.

Let's now define a `DeleteRestaurant` mutation that removes a single restaurant from our database. We'll add it to our `schema.py` file between our `CreateRestaurant` and `Mutation` classes:

```py
class DeleteRestaurant(graphene.Mutation):
  class Arguments:
    id = graphene.Int()

  ok = graphene.Boolean()

  def mutate(self, info, id):
    restaurant = Restaurant.objects.get(id=id)
    restaurant.delete()
    return DeleteRestaurant(ok=True)
```

Next, we'll add the `DeleteRestaurant` mutation to the `Mutation` class:

```py
class Mutation(graphene.ObjectType):
  create_restaurant = CreateRestaurant.Field()
  delete_restaurant = DeleteRestaurant.Field()
```

Next, let's run the mutation on the browser to delete a restaurant from our GraphQL database:

```graphql
mutation {
  deleteRestaurant(id: 1) {
    ok
  }
}
```

We pass the restaurant id as an argument to the mutation as shown above. The output should look like this:

```json
{
  "data": {
    "deleteRestaurant": {
      "ok": true
    }
  }
}
```

**Note**: We should run a query to get the list of restaurants again to see the change.

Lastly, let's make an `UpdateRestaurant` mutation that modifies data for a single restaurant. This will be added to our `schema.py` file, above our `Mutation` class:

```py
class UpdateRestaurant(graphene.Mutation):
  class Arguments:
    id = graphene.Int()
    name = graphene.String()
    address = graphene.String()

  ok = graphene.Boolean()
  restaurant = graphene.Field(RestaurantType)

  def mutate(self, info, id, name, address):
    restaurant = Restaurant.objects.get(id=id)
    restaurant.name = name
    restaurant.address = address
    restaurant.save()
    return UpdateRestaurant(ok=True, restaurant=restaurant)
```

Let's add the `UpdateRestaurant` mutation to the `Mutation` class:

```py
class Mutation(graphene.ObjectType):
  create_restaurant = CreateRestaurant.Field()
  delete_restaurant = DeleteRestaurant.Field()
  update_restaurant = UpdateRestaurant.Field()
```

We'll now run the mutation on the browser like so:

``` graphql
mutation {
  updateRestaurant(id: 2, name: "Kada Plaza Ltd", address: "Lekki Gardens") {
    ok
    restaurant {
      id
      name
      address
    }
  }
}
``` 

The output should look like this:

```json
{
  "data": {
    "updateRestaurant": {
      "ok": true,
      "restaurant": {
        "id": 2,
        "name": "Kada Plaza Ltd",
        "address": "Lekki Gardens"
      }
    }
  }
}
```

### Conclusion

GraphQL allows us to make requests from our database without creating separate endpoints for each request. In this article, we built a CRUD application with Django using GraphQL queries and mutations. 

Source code for this article: [https://github.com/Smyja/codecademy](https://github.com/Smyja/codecademy)
