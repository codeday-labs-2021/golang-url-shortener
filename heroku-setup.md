## Setting up Heroku

## Step 1 - Login

```
$ heroku login
```

## Step 2 - Create a heroku app

```
$ heroku create
```

## Step 3 - Tell heroku this is a docker container

```
$ heroku stack:set container
```

## Step 4 - Push changes to heroku

```
$ git push heroku main
```

## Step 5 - View the app

```
$ heroku open
```
