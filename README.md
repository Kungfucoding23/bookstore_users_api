# bookstore_users_api
Users API

## Basic aplication structure: MVC Pattern

The controllers main responsability is to take the request, validate if we have all the parameters that we need to handle and send this handling to the service

The entire business logic will be in our services

The only point where we are going to interact with the http framework we are using is in the app folder where we start the application and in the contrllers

