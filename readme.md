# Welcome to NumToEn!

This is an API written in golang to translate a number into english words.
The endpoint accepts numbers between -9223372036854775808 and 9223372036854775807

## For running locally with Docker container

Download from DockerHub
```shell script
docker pull scyanh/numtoen:latest
```

Run container
```shell script
docker run -d -p 5000:5000 scyanh/numtoen:latest
```

#### Endpoint to translate numbers...
http://localhost:5000/num_to_english?number=123

#### Endpoint to documentation...
http://localhost:5000/swagger/index.html


### Run the tests
Open Docker cli from your container and run
```shell script
go test -v
```

## For Deployment the container on Heroku

Clone repository and navigate to the app’s directory
```shell script
git clone https://github.com/scyanh/numtoen
cd numtoen
```

Download and install the Heroku CLI
```shell script
brew tap heroku/brew && brew install heroku
```

Log in to Heroku
```shell script
heroku login
```

Log in to Container Registry:
```shell script
heroku container:login
```

Create a Heroku app (it will create a random appname)
```shell script
heroku create
```

Build the image and push to Container Registry:
```shell script
heroku container:push web -a yourappname
```

Then release the image to your app:
```shell script
heroku container:release web -a yourappname
```

Now open the app in your browser:
```shell script
heroku open -a yourappname
```


#### Live API Documentation...
https://infinite-falls-90558.herokuapp.com/swagger/index.html

#### Live API Endpoint to translate numbers...
https://infinite-falls-90558.herokuapp.com/num_to_english?number=123