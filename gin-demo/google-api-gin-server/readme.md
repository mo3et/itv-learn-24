# Developing a RESTful API with Go and Gin
This tutorial introduces the basics of writing a RESTful web service API with Go and the Gin Web Framework (Gin).

You'll get the most out of this tutorial if you have a basic familiarity with Go and its tooling. If this is your first exposure to Go, please see Tutorial: Get started with Go for a quick introduction.

Gin simplifies many coding tasks associated with building web applications, including web services. In this tutorial, you'll use Gin to route requests, retrieve request details, and marshal JSON for responses.

In this tutorial, you will build a RESTful API server with two endpoints. Your example project will be a repository of data about vintage jazz records.

The tutorial includes the following sections:

1.Design API endpoints.
2.Create a folder for your code.
3.Create the data.
4.Write a handler to return all items.
5.Write a handler to add a new item.
6.Write a handler to return a specific item.

# Design API endpoints
You'll build an API that provides access to a store selling vintage recordings on vinyl. So you'll need to provide endpoints through which a client can get and add albums for users.

When developing an API, you typically begin by designing the endpoints. Your API's users will have more success if the endpoints are easy to understand.

Here are the endpoints you'll create in this tutorial.

/albums

- `GET` – Get a list of all albums, returned as JSON.
- `POST` – Add a new album from request data sent as JSON.
/albums/:id

- `GET` – Get an album by its ID, returning the album data as JSON.
Next, you'll create a project for your code.

> Ref: https://shell.cloud.google.com/?walkthrough_tutorial_url=https%3A%2F%2Fraw.githubusercontent.com%2Fgolang%2Ftour%2Fmaster%2Ftutorial%2Fweb-service-gin.md&pli=1&show=ide&environment_deployment=ide