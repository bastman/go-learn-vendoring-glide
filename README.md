# go-learn-vendoring-glide
getting started with glide dependency manager for golang vendoring

## download glide
https://glide.sh/

## define dependencies
    $ glide get github.com/ahmetb/go-linq
    $ glide get github.com/yanatan16/itertools
    $ github.com/doloopwhile/go-itertools

    $ glide get github.com/chrislusf/glow
    $ glide get github.com/chrislusf/glow/flow

    //# gleam: requires half of the world internet ;)
    //$ glide get github.com/chrislusf/gleam/flow
    // $ glide --debug install net/http

## build dependencies
    $ glide install
    $ glide update


