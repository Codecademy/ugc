---
Title: "How To Use GraphQL With Django"
Description: "A guide on how to use GraphQL with Django."
DatePublished: "2022-06-05"
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
_**Versions:** Python 3.8.10, Django 4.0.4_

## Introduction

[GraphQL](https://www.codecademy.com/resources/docs/general/graphql) is a query language for [APIs](https://www.codecademy.com/resources/docs/general/api) and a runtime for fulfilling those queries with your existing data. Unlike a REST API, GraphQL APIs do not require verbs (`PUT`, `POST`, `GET`, `PATCH`,  and `DELETE`) for requests, nor do they need multiple endpoints. They have just one endpoint and making a query to that endpoint is all that's needed. 

This tutorial will cover the creation of a CRUD (create, read, update, and delete) GraphQL API for a restaurant with Django.

## Properties of GraphQL

The following terms are often used when interacting with GraphQL. Knowing them can be helpful, though we won't be covering all of them in this tutorial.

- **Schema**: Describes the functionality available to the client applications that connect to it.
- **Query**: A schema type that represents the `GET` request and defines the operations that can be used for reading or fetching data.
- **Nesting**: Queries can be nested inside of other queries.
- **Mutation**: A schema type that defines the kind of operations that can be done to modify data.
- **Subscription**: Notifies the client server in real time about updates to the data.
- **Resolver**: Functions that return values for fields associated with existing schema types.

## Setting up GraphQL with Django

First, we need to install Django and create a project. Run  ```virtualenv venv```  to setup a virtual environment. Then, run ```source venv/bin/activate``` on Linux/MacOs or ```source/scripts/activate``` on windows to activate the virtual environment. 

Next,install django with ```pip install django``` and run ```django-admin startproject project .``` to create a new project.

Create an app in the project with ```python manage.py startapp myapp``` . Add the app to the project's ```INSTALLED_APPS``` like this
```py
INSTALLED_APPS = [
  'myapp',

  # ...
]
```

### Using `graphene-django`

To use GraphQL with Django, you will need to install the [`graphene-django`](https://docs.graphene-python.org/projects/django/en/latest/) package. This will help us with adding GraphQL functionality to our restaurant Django app:

`pip install graphene-django`

Add it to your `settings.py` file:

```py
INSTALLED_APPS = [
  'graphene_django',

  # ...
]
```

Create a `Models.py` file in your project that contains the models you want to use:

```py
from django.db import models

class Restaurant(models.Model):
  name = models.CharField(max_length=100)
  address = models.CharField(max_length=200)

  def __str__(self):
      return self.name
```
Register the  models in the admin.py file:

```py
from . import models

admin.site.register(models.Restaurant)
```

Make migrations with ```python manage.py makemigrations``` . Then, run ```python manage.py migrate``` to apply the migrations.

If you run the server with ```python manage.py runserver```, you will see this error in your Command Prompt:

`ImportError: cannot import name 'force_text' from 'django.utils.encoding`

This error comes up because django no longer supports the ```force_text``` variable which the graphene package uses.
Add the following to the top of your `settings.py` file to resolve it:

```py
import django
from django.utils.encoding import force_str

django.utils.encoding.force_text = force_str

```
or downgrade your django version.

Create a `urls.py` file in your `myapp` directory and add a GraphQL route for Django version 2.0 and above:

```py
from graphene_django.views import GraphQLView
from django.views.decorators.csrf import csrf_exempt
from django.urls import path

urlpatterns = [
  path("graphql", csrf_exempt(GraphQLView.as_view(graphiql=True))),
]
```
Add the app urls to your `urls.py` file in your project folder:

```py
from django.urls import path, include

urlpatterns = [
    path("admin/", admin.site.urls),
    path("", include("myapp.urls")),
]
```

GraphQL comes with an API browser, [GraphiQL](https://graphiql-test.netlify.app/typedoc/), that is similar to Django's browsable API where you can use to test your queries and mutations. This is done with the `graphiql` parameter of the `.as_view()` method. However, if you do not want to use it, you can set `graphiql` to `False`.
Django's `csrf_exempt` decorator is used to allow API clients to POST to the graphql endpoint we have created.

Create a Schema.py file in the `myapp` directory and then create a type for your models in your `schema.py` file as shown below:
```py
import graphene
from graphene_django import DjangoObjectType
from myapp.models import Restaurant
class RestaurantType(DjangoObjectType):
    class Meta:
        model = Restaurant
        fields = ("id", "name", "address")

```

Create a Query type for the Restaurant model on your `schema.py` file as shown below:

```py
class Query(graphene.ObjectType):
    """
    Queries for the Restaurant model
    """
    restaurants = graphene.List(RestaurantType)

    def resolve_restaurants(self, info, **kwargs):
      return Restaurant.objects.all()

```
The Query type contains a resolver for the `restaurants` field. This resolver returns all the restaurants in the database.
Add to the end of your `schema.py` file which will create the schema for your GraphQL API:
 
```py
schema = graphene.Schema(query=Query)
``` 
Import the schema variable into the `urls.py` file and pass it to the graphql view as seen below:

```py
from myapp.schema import schema

url_patterns = [
  path("graphql", csrf_exempt(GraphQLView.as_view(graphiql=True, schema=schema))),
]
```
Th entire schema.py file should look like this:

```py
import graphene
from graphene_django import DjangoObjectType
from myapp.models import Restaurant

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


Start the django server with `python manage.py runsrver` then visit the `/graphql` route to see the api browser, it should look like this

![API browser](https://raw.githubusercontent.com/Smyja/ugc/grahql-with-django/content/smyja/api-browser.png)

Create a Superuser account,login by visiting the admin url and add restaurants to the database with the following command:

```py
python manage.py createsuperuser
```

To get the list of restaurants with specific data like name and address, run a query with this:

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

To modify any data in our database we would need to create a mutation.
Below is the `CreateRestaurant` mutation, which we will add to the `schema.py` file.

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

The `CreateRestaurant` mutation takes in the model fields as arguments while the `mutate()` function is where the db change happens using Django's object-relational mapper (ORM).

Create a `Mutation` class then initialize the mutation with the schema.

```py
class Mutation(graphene.ObjectType):
    create_restaurant = CreateRestaurant.Field()
```

After adding the mutation, pass the mutation to the schema at the end of the `schema.py` file.

```py
schema = graphene.Schema(query=Query, mutation=Mutation)
```

Start the server and a run a mutation with the GraphQL API browser using this:

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

To delete a restaurant, we would need to create a mutation. Below is the `DeleteRestaurant` mutation. Let's add it to our `schema.py` file.

```python
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
delete_restaurant = DeleteRestaurant.Field()
```

Run the mutation to delete a restaurant with the GraphQL API browser using this.

```graphql  
    mutation {
        deleteRestaurant(id: 1) {
            ok
        }
    }
```

Pass the restaurant id as an argument to the mutation as shown above. Output should look like this:

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

To update a restaurant, we would need to create a mutation. Below is the `UpdateRestaurant` mutation. We will add it to the `schema.py` file:

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
update_restaurant = UpdateRestaurant.Field()
```

Run the mutation with the GraphQL API browser using this:

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

GraphQL lets you request for what you want from your database without creating separate endpoints for each request. In this article, we built a CRUD application with Django using GraphQL queries and mutations.

Github repository: https://github.com/Smyja/codecademy