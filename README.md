# covid19-info
this is get info covid 19 from whole world base on https://disease.sh/docs/


# How To Use Program 
because this app still using sanbox mode whatsapp, the number app must me register to sandbox mode

1. Register your number to sandbox mode in whatsapp number `+1 415 523 8886` with message `join foreign-arrangement` or click [Here](https://api.whatsapp.com/send?phone=14155238886&text=join%20foreign-arrangement) 
2. Send command to use app, the list command:

* `CASES <country>` (return total active case in specific country per today)
* `DEATHS <country>` (return total deaths case in spesific country per today)
* `CASES TOTAL` (return total active case in global wolrd)
* `DEATHS TOTAL` (return total deaths case in global world)

# How to deploy to Heroku
* Login to heroku with `heroku login`
* clone or add existing code to repo git
* create app with `heroku create`
* create Procfile to run job when code push to heroku with value file 
   web: bin/go-getting-started
* set config var with `heroku config:set VARNAME=VARVALUE`
* Config var to set is ACCOUNTID, TOKEN, SERVICE_NUMBER

* push and build to heroku `git push heroku master`
* view logs with `heroku logs --tail`

# How to use demo example
* 

# Next Feature 
* Check if country code not exist
* Migrate Whatsapp sanbox to Whatsapp Live Production