---
Title: "How to Use GraphQL With Django"
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

_**Prerequisites:** Understanding of Python and Django._  
_**Versions:** Python 3.8.10, Django 4.0.4_

## Introduction

[GraphQL](https://www.codecademy.com/resources/docs/general/graphql) is a query language for [APIs](https://www.codecademy.com/resources/docs/general/api) and a runtime for fulfilling those queries with your existing data. Unlike a REST API, GraphQL APIs do not require verbs (`PUT`, `POST`, `GET`, etc.) for requests, nor do they need multiple endpoints. They have just one endpoint and making a query to that endpoint is all that's needed. 

This tutorial will cover the creation of a CRUD (create, read, update, delete) GraphQL API for a restaurant with Django.

## Properties of GraphQL

The following terms would be used a lot when interacting with graphql. It is important that you know them, though we wouldn't cover all of them in this tutorial.

- **Schema**: Describes the functionality available to the client applications that connect to it.
- **Query**: A schema type that represents the `GET` request and defines the operations that can be used for reading or fetching data.
- **Nesting**: Queries can be nested inside of other queries.
- **Mutation**: A schema type that defines the kind of operations that can be done to modify data.
- **Subscription**: Notifies the client server in real time about updates to the data.
- **Resolvers**: Functions that return values for fields associated with existing schema types.

## Setting up GraphQL with Django

First, we need to install Django and create a project. I will be using a [bash script](https://github.com/Smyja/Django-bash) I created to do this.


To use GraphQL with Django, you will need to install the `graphene-django` package:

`pip install graphene-django`

Add it to your `settings.py` file:

```py
INSTALLED_APPS = [
    'graphene_django',

    # ...
]
```

If you run the server, you will see this error:

`ImportError: cannot import name 'force_text' from 'django.utils.encoding`

Add the following to the top of your `settings.py` file:

```py
import django
from django.utils.encoding import force_str

django.utils.encoding.force_text = force_str

```
or downgrade your django version.

Create a `Models.py` file in your project that contains the models you want to use:
```py
from django.db import models

class Restaurant(models.Model):
    name = models.CharField(max_length=100)
    address = models.CharField(max_length=200)

    def __str__(self):
        return self.name
```

Add a GraphQL route to your `urls.py` file for Django version 2.0 and above:

```py
from graphene_django.views import GraphQLView
from django.views.decorators.csrf import csrf_exempt
from djql.schema import schema #change djql to your app name

urlpatterns = [
    path("graphql", csrf_exempt(GraphQLView.as_view(graphiql=True, schema=schema))),
]
```

GraphQL comes with an API browser, [GraphiQL](https://graphiql-test.netlify.app/typedoc/), that is similar to Django's browsable API where you can use to test your queries and mutations. This is done with the `graphiql` parameter of the `.as_view()` method. However, if you do not want to use it, you can set `graphiql` to `False`.
The third import statement ```from djql.schema import schema``` is the schema that we will use to create our queries. Create a `schema.py` file in your project directory or your app directory.
Django's csrf_exempt decorator is used to allow API clients to POST to the graphql endpoint we have created.

Create a GraphQL type for your models on your `schema.py` file as shown below:

```py
import graphene
from graphene_django import DjangoObjectType
from djql.models import Restaurant

class Query(graphene.ObjectType):
    """
    Queries for the Restaurant model
    """
    restaurants = graphene.List(RestaurantType)

    def resolve_restaurants(self, info, **kwargs):
      return Restaurant.objects.all()

```
Start the django server with `python manage.py runsrver` then visit the `/graphql` route to see the api browser, it should look like this

![Api browser](https://raw.githubusercontent.com/Smyja/ugc/grahql-with-django/content/smyja/api-browser.png)


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

After adding the mutation and query, define the schema at the end of the `schema.py` file.

```schema = graphene.Schema(query=Query, mutation=Mutation)```

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

```delete_restaurant = DeleteRestaurant.Field()```

Run the mutation with the GraphQL API browser using this.

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
