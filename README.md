greentros
==============

avatar generator web service

Contributors
=============

* [Santiago](https://github.com/santiaago)
* [Remy](https://github.com/rjourde)


Stack
======

* Go
* mongodb
* heroku

Installation
======

*   `cd $GOPATH/src`
*   `go get github.com/taironas/greentros`
*   `cd $GOPATH/src/github.com/greentros`
*   `go get`
*   `export PORT=8080`

Run App
=======

    > pwd
    $GOPATH/src/github.com/taironas/greentros
    > greentros
    2014/11/19 22:23:57 Listening on 8080

Build
======
    >`cd $GOPATH/src/github.com/greentros`
    >`go get`

Deploy
=======
    santiagos-MacBook-Pro:greentros santiaago$ git push heroku master
    Fetching repository, done.
    Counting objects: 5, done.
    Delta compression using up to 8 threads.
    Compressing objects: 100% (3/3), done.
    Writing objects: 100% (3/3), 287 bytes | 0 bytes/s, done.
    Total 3 (delta 2), reused 0 (delta 0)
    
    -----> Fetching custom git buildpack... done
    -----> Go app detected
    -----> Using go1.3
    -----> Running: go get -tags heroku ./...
    -----> Discovering process types
           Procfile declares types -> web
    
    -----> Compressing... done, 1.5MB
    -----> Launching... done, v6
           https://greentros.herokuapp.com/ deployed to Heroku
    
    To git@heroku.com:greentros.git
       56a3000..5572085  master -> master


