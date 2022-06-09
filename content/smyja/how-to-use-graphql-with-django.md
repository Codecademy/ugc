---
Title: "How to use GrahQl with Django"
Description: "An introduction to ADTs in JavaScript."
DatePublished: "2022-06-05"
Categories:
  - "web-development"
  - "computer-science"
Tags:
  - "APIs"
  - "Django"
  - "Python"
CatalogContent:
  - "how-to-use-graphql-with-django"
  - "paths/computer-science"
---
_**Prerequisites:** Understanding of Python and Django._

_**Versions:** Python 3.8.10, Django 4.0.4_

## Introduction
GraphQL is a query language for APIs and a runtime for fulfilling those queries with your existing data. Unlike a REST api, Graphql api’s do not require verbs(PUT,POST,GET) for requests, they do not need multiple Endpoints. They have just one endpoint and making a query to that endpoint is all that's needed. 

This tutorial will cover the creation of a CRUD(create,read,update,delete)graphql api for a restaurant with django.

 **Properties of Graphql**:

The following terms would be used a lot when interacting with graphql.
- schema
- query
- Nesting
- Mutation
- Subscription
- Resolvers

**Schema:** 
A schema is a collection of types. A type is a collection of fields. A field is a property of a type.

**Query:**
A Query is a type on a schema that defines the kind of operations that can be done to get data.
Creating a query involves adding fields to a query type, then creating resolvers for the fields

**Resolvers:**
They are functions that return values for fields that exist on types in a schema. 

**Mutation:**
A mutation is a type on a schema that defines the kind of operations that can be done to modify data.

**Subscription:**
A subscription in graphql i

**Nesting:**
A nested query is a query that is a child of another query.

## Setting up graphql with django
First, we need to install Django and create a project. I will be using a bash script i created to do this.
[Django-bash-script](https://github.com/Smyja/Django-bash)

To use graphql with django, you will need to install graphene-django

```pip install graphene-django```

Add it to your settings.py file:
```python
INSTALLED_APPS = [
    'graphene_django',

    # ...
]
```
If you run the server, you see this error:
```ImportError: cannot import name 'force_text' from 'django.utils.encoding```

add this to your settings.py file:

```python
import django
from django.utils.encoding import force_str

django.utils.encoding.force_text = force_str

```
or downgrade your django version.

Create a Models.py file in your project that contains tthe models you want to use.:
```python
from django.db import models
class Restaurant(models.Model):
    name = models.CharField(max_length=100)
    address = models.CharField(max_length=200)

    def __str__(self):
        return self.name
```


Add a graphql route to your urls.py file for django 2.0 and above:

```python
from graphene_django.views import GraphQLView
from django.views.decorators.csrf import csrf_exempt
from djql.schema import schema #change djql to your app name

urlpatterns = [
    path("graphql", csrf_exempt(GraphQLView.as_view(graphiql=True, schema=schema))),
]
```

Graphql comes with an api browser similar to django's browsable api that you can use to test your queries and mutations. If you do not want to use it, you can set the `graphiql` parameter to False.
The fourth import statement ```from djql.schema import schema``` is the schema that we will use to create our queries. Create a `schema.py` file in your project directory or your app directory.
Django's csrf_exempt decorator is used to allow API clients to POST to the graphql endpoint we have created.

Create a Graphql Type for your models on your schema.py file as shown below:
```python
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
Visit the `/graphql` route to see the api browser, it sh

To get the list of Restaurants run a query with this
```graphql
query {
    restaurants {
        id
        name
        address
    }
}

```

Next create a mutation to create a restaurant

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

