# Welcome to scrabble.nehpets.co.uk

A Scrabble game to keep the family entertained during lockdown,
mostly because bandwidth is too low to have so many videos going.


### Start the web server:

  revel run -a scrabble.nehpets.co.uk

### Go to http://localhost:9000/ and you'll see:

  [demo]: https://raw.githubusercontent.com/daflad/scrabble.nehpets.co.uk/master/public/img/demo.png "Demo board"

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers here
        controllers/  App models here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        img/          Image files
        node_modules/ Node based css & js files

    tests/            Test suites
