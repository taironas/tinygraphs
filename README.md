[tinygraphs](http://tinygraphs.com)
==============

**tinygraphs** is an avatar generator web service. Checkout [tinygraphs.com](http://www.tinygraphs.com/) to try it.

[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

#### Contributors

* [Santiago](https://github.com/santiaago)
* [Remy](https://github.com/rjourde)
* [Carmen](https://plus.google.com/+CarmenRebolledo)

#### Blog post:

<a href="http://www.sanarias.com/blog/415BuildingtinygraphsanavatarwebserviceinGo" target="_blank">Building tinygraphs an avatar web service in Go</a>

How to use:
======

* You can set the HTML source of the image to point directly to [tinygraphs.com](http://tinygraphs.com)
~~~html
<img src="http://tinygraphs.com/squares/helloworld">
~~~
* You can save the image and use it directly on your site
* You can `go get` this repo and use it.

*just remember to give us credit with a link to tinygraphs.com ;)*

## Supported routes:

`http://tinygraphs.com/squares/anything`

![squares](http://tinygraphs.com/squares/anything?size=120&theme=frogideas&numcolors=2&fmt=svg) ![squares](http://tinygraphs.com/squares/anything?size=120&theme=frogideas&numcolors=3&fmt=svg) ![squares](http://tinygraphs.com/squares/anything?size=120&theme=frogideas&numcolors=4&fmt=svg)

`http://tinygraphs.com/isogrids/helloworld`

![squares](http://tinygraphs.com/isogrids/helloworld?size=120&theme=heatwave) ![squares](http://tinygraphs.com/isogrids/helloworld?size=120&theme=heatwave&numcolors=3) ![squares](http://tinygraphs.com/isogrids/helloworld?size=120&theme=heatwave&numcolors=4)

`http://tinygraphs.com/spaceinvaders/helloworld`

![squares](http://tinygraphs.com/spaceinvaders/holamundo?size=120&theme=frogideas)
![squares](http://tinygraphs.com/spaceinvaders/hello?size=120&theme=berrypie) ![squares](http://tinygraphs.com/spaceinvaders/world?size=120&theme=bythepool)

`http://tinygraphs.com/squares/banner/random?h=50&xs=100`

![square random banner](http://tinygraphs.com/squares/banner/random?h=50&xs=100&theme=frogideas&fmt=svg)

![square random banner](http://tinygraphs.com/squares/banner/random?h=50&xs=100&theme=seascape&fmt=svg)

![square random banner](http://tinygraphs.com/squares/banner/random?h=50&xs=100&theme=heatwave&fmt=svg)

`http://tinygraphs.com/squares/banner/random/gradient?theme=frogideas&xs=100`

![square random banner](http://tinygraphs.com/squares/banner/random/gradient?theme=frogideas&h=50&xs=100)

![square random banner](http://tinygraphs.com/squares/banner/random/gradient?theme=seascape&h=50&xs=100)

![square random banner](http://tinygraphs.com/squares/banner/random/gradient?theme=duskfalling&h=50&xs=100)

`http://tinygraphs.com/isogrids/banner/random?h=50&xt=100`

![isogrids random banner](http://tinygraphs.com/isogrids/banner/random?h=50&xt=100&theme=frogideas)

![isogrids random banner](http://tinygraphs.com/isogrids/banner/random?h=50&xt=100&theme=seascape)

![isogrids random banner](http://tinygraphs.com/isogrids/banner/random?h=50&xt=100&theme=heatwave)

`http://tinygraphs.com/isogrids/banner/random/gradient?theme=frogideas&h=50&xt=100`

![square random banner](http://tinygraphs.com/isogrids/banner/random/gradient?theme=frogideas&h=50&xt=100)

![square random banner](http://tinygraphs.com/isogrids/banner/random/gradient?theme=seascape&h=50&xt=100)

![square random banner](http://tinygraphs.com/isogrids/banner/random/gradient?theme=heatwave&h=50&xt=100)

##Lab routes:

`http://tinygraphs.com/labs/checkerboard`

![checkerboard](http://tinygraphs.com/labs/checkerboard?size=120)

`http://tinygraphs.com/labs/squares/random`

![random](http://tinygraphs.com/labs/squares/random?size=120&theme=summerwarmth&numcolors=4) ![random](http://tinygraphs.com/labs/squares/random?size=120&theme=daisygarden&numcolors=4)
![random](http://tinygraphs.com/labs/squares/random?size=120&theme=duskfalling&numcolors=4)

`http://tinygraphs.com/labs/isogrids/random`

![random](http://tinygraphs.com/labs/isogrids/random?size=120&theme=seascape) ![random](http://tinygraphs.com/labs/isogrids/random?size=120&theme=seascape&numcolors=3) ![random](http://tinygraphs.com/labs/isogrids/random?size=120&theme=seascape&numcolors=4)

`http://tinygraphs.com/labs/isogrids/hexa`

![hexa](http://tinygraphs.com/labs/isogrids/hexa/hello?size=120&theme=frogideas&numcolors=4) ![hexa](http://tinygraphs.com/labs/isogrids/hexa/hello?size=120&theme=bythepool&numcolors=4) ![hexa](http://tinygraphs.com/labs/isogrids/hexa/hello?size=120&theme=berrypie&numcolors=4)

`http://tinygraphs.com/labs/isogrids/hexa16`

![hexa16](http://tinygraphs.com/labs/isogrids/hexa16/hello?size=120&theme=frogideas&numcolors=2) ![hexa16](http://tinygraphs.com/labs/isogrids/hexa16/hello?size=120&theme=bythepool&numcolors=3) ![hexa16](http://tinygraphs.com/labs/isogrids/hexa16/hello?size=120&theme=berrypie&numcolors=4)

`http://tinygraphs.com/labs/squares/banner/gradient`

![gradient colors squares](http://www.tinygraphs.com/labs/squares/banner/gradient?h=100&xs=50&theme=frogideas&xs=25)

![gradient colors squares](http://www.tinygraphs.com/labs/squares/banner/gradient?h=100&xs=50&theme=frogideas&xs=25&numcolors=2&gx2=40&gx1=0&gy1=50&gy2=100)

`http://tinygraphs.com/labs/isogrids/banner/gradient`

![gradient colors isogrids](http://www.tinygraphs.com/labs/isogrids/banner/gradient?h=100&xs=50&theme=frogideas&xs=25)

![gradient colors isogrids](http://www.tinygraphs.com/labs/isogrids/banner/gradient?h=100&xs=50&theme=frogideas&xs=25&numcolors=2&gx2=40&gx1=0&gy1=50&gy2=100)

## Parameters:

* **size**: `tinygraphs.com/squares/hello?size=60`
* **formats**: `tinygraphs.com/squares/hello?fmt=svg`

    The default format is `SVG`.

    All routes support SVG format, except `Square` routes who also support JPEG.

* **background and foreground**: `tinygraphs.com/squares/hello?bg=ff4008&fg=04d6f2`

    You can specify the color of the background or foreground by using parameters `bg` and `fg` and passing an hexadecimal value of the color:

* **theme**: `tinygraphs.com/labs/squares/random?theme=frogideas`

    You can specify the theme you want to take into account in the image.

    ![theme](http://tinygraphs.com/labs/squares/random?theme=frogideas&size=120&fmt=svg)

    Here is the list of existing themes:
    * ![theme](http://tinygraphs.com/themes/base?fmt=svg&h=20) base
    * ![theme](http://tinygraphs.com/themes/sugarsweets?fmt=svg&h=20) sugarsweets
    * ![theme](http://tinygraphs.com/themes/heatwave?fmt=svg&h=20) heatwave
    * ![theme](http://tinygraphs.com/themes/daisygarden?fmt=svg&h=20)daisygarden
    * ![theme](http://tinygraphs.com/themes/seascape?fmt=svg&h=20) seascape
    * ![theme](http://tinygraphs.com/themes/summerwarmth?fmt=svg&h=20) summerwarmth
    * ![theme](http://tinygraphs.com/themes/bythepool?fmt=svg&h=20) bythepool
    * ![theme](http://tinygraphs.com/themes/duskfalling?fmt=svg&h=20) duskfalling
    * ![theme](http://tinygraphs.com/themes/frogideas?fmt=svg&h=20) frogideas
    * ![theme](http://tinygraphs.com/themes/berrypie?fmt=svg&h=20) berrypie


* **numcolors**: `tinygraphs.com/labs/squares/random?theme=summerwarmth&numcolors=4`

    You can specify the number of colors that you want to render the image.
    Default value is **2** and can be extended to **4**.

    ![theme](http://tinygraphs.com/labs/squares/random?theme=frogideas&size=120&fmt=svg&numcolors=2) ![theme](http://tinygraphs.com/labs/squares/random?theme=frogideas&size=120&fmt=svg&numcolors=3) ![theme](http://tinygraphs.com/labs/squares/random?theme=frogideas&size=120&fmt=svg&numcolors=4)

* **inv**:
`tinygraphs.com/squares/hello?theme=frogideas&numcolors=2&inv=1`

    You can specify if you want to see the colors inverted.
    Default value is **false**. `inv` parameter works with theme colors or default (black and white) colors. The number of colors has to be equal to **2**.

    ![normal](http://tinygraphs.com/squares/hello?theme=frogideas&size=120&fmt=svg&numcolors=2&inv=0) ![inverse](http://tinygraphs.com/squares/hello?theme=frogideas&size=120&fmt=svg&numcolors=2&inv=1)

* **order**:
`tinygraphs.com/squares/hello?theme=frogideas&numcolors=4&order=3&order=2&order=1&order=0`

    You can specify the order in which you want to see the colors by using the `order` parameter. Just add the indexes in which you wish to see the colors to the order array (indexes are zero based)
    Like so: `order=3&order=2&order=1&order=0`

    ![normal](http://tinygraphs.com/squares/hello?theme=seascape&size=120&fmt=svg&numcolors=4)
    ![reorder](http://tinygraphs.com/squares/hello?theme=seascape&size=120&fmt=svg&numcolors=4&order=3&order=2&order=1&order=0)

* **lines**:
`tinygraphs.com/isogrids/hello?lines=4`

 You can specify the number of lines that an isogrid can have using the `lines`parameter. **Default** parameter is **6**. Value has to be greater or equal to 4.

 ![number of lines in isogrid image.](http://tinygraphs.com/isogrids/hello?lines=4&size=120&fmt=svg)

* **colors**:

 `tinygraphs.com/isogrids/hello?lines=4&colors=43e0e8&colors=e84b43`
 You can specify use you own colors by using the `colors` parameter and passing hexa colors.

    ![isogrid with custom colors](http://tinygraphs.com/isogrids/hello?lines=4&size=120&colors=43e0e8&colors=e84b43)

* **banner parameters:**
    * `h`: height parameter.
    * `w`: width parameter.
    * `xt`: number of triangles in `isogrid` banner in x axis.
    * `xs`: number of squares in `squares` banner in x axis.
    * `theme`: tinygraphs theme to use in banner.
    * `numcolors`: number of colors to take into account when rendering banner.

   `tinygraphs.com/isogrids/banner/random/gradient?theme=bythepool&numcolors=4`

 ![isogrid gradient banner.](http://tinygraphs.com/isogrids/banner/random/gradient?h=75&xt=80&theme=bythepool&numcolors=4)

* **random banner parameters:**
    * `p`: probability of the main color in the banner. Default value is `0.5`. Values should be between `0` and `1`

    `tinygraphs.com/labs/isogrids/banner/gradient?theme=bythepool&p=0.1`

    ![isogrid gradient banner.](http://tinygraphs.com/labs/isogrids/banner/gradient?h=75&xt=80&theme=bythepool&numcolors=2&p=0.1)

* **isogrid color gradient parameters**

    You can manipulate the gradient vector using the following parameter.
    * `gx1`: first gradient coordinate along x axis.
    * `gy1`: first gradient coordinate along y axis.
    * `gx2`: second gradient coordinate along x axis.
    * `gy2`: second gradient coordinate along y axis.

    `tinygraphs.com/labs/isogrids/banner/gradient?theme=bythepool&xt=4&h=120&w=120&gx1=0&gy1=0&gy2=60&gx2=60`

    ![isogrid color gradient](http://www.tinygraphs.com/labs/isogrids/banner/gradient?theme=bythepool&xt=4&h=120&w=120&gx1=0&gy1=0&gy2=60&gx2=60)

Stack
======

* Go
* Heroku

Third parties
=====

* [Bootstrap](http://getbootstrap.com/)
* [AngularJS](http://angularjs.org/)
* [route](http://github.com/taironas/route/)
* [svgo](https://github.com/ajstarks/svgo)

Installation
======

    cd $GOPATH/src
    go get github.com/ajstarks/svgo
    go get github.com/taironas/route
    go get github.com/taironas/tinygraphs
    cd $GOPATH/src/github.com/tinygraphs
    go get ./app-backend
    export PORT=8080

Run App
=======

    > pwd
    $GOPATH/src/github.com/taironas/tinygraphs
    > app-backend
    2014/11/19 22:23:57 Listening on 8080

Build
======
    >cd $GOPATH/src/github.com/tinygraphs
    >go get ./app-backend

Test locally
=============
**option 1:**

    > app-backend
    2014/12/07 00:35:02 Listening on 8080

**option 2:**

If you have heroku install you should be able to run

    > foreman start
    00:37:38 web.1  | started with pid 5762
    00:37:38 web.1  | 2014/12/07 00:37:38 Listening on 8080

Deploy
=======

Easy

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

Manual

**Note:** heroku is now configured to build and deploy any `git push`to `master`. If you still want to manual deploy the app follow the steps below.

Before you start be sure to have the proper rsa key. [See Managing Your SSH Keys](https://devcenter.heroku.com/articles/keys) for more details and that. Also be sure to be logged in with heroku.

    > heroku login
    Enter your Heroku credentials.
    Email: ga@tinygraphs.com
    Password:

After that you can deploy as follows:

    > git push heroku master
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
           https://tinygraphs.herokuapp.com/ deployed to Heroku

    To git@heroku.com:tinygraphs.git
       56a3000..5572085  master -> master
