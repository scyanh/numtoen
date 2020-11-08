# Welcome to NumToEn!

This is an API written in golang to translate a number into english words.

### For running locally
$ go run go.main

### For running tests
$ go test -v

### For Deployment on Heroku
Create an App on Heroku, in this example I created one called "numtoen"

Install the Heroku CLI
Download and install the Heroku CLI.

If you haven't already, log in to your Heroku account and follow the prompts to create a new SSH public key.

$ heroku login
Create a new Git repository
Initialize a git repository in a new or existing directory

$ cd my-project/
$ git init
$ heroku git:remote -a numtoen

Deploy your application committing your code to the repository and deploy it to Heroku using Git.

$ git add .
$ git commit -am "make it better"
$ git push heroku master

Existing Git repository
For existing repositories, simply add the heroku remote

$ heroku git:remote -a numtoen

### Swagger API documentation 
/swagger/index.html

