# godic

godic is a simple dictionary tool on command-line.  
It includes Japanese-English, English-Japanese and Thesaurus.

## Installation

You can use `go get`.

```
go get -u github.com/zuiurs/godic
```

## Usage

A simple usage is shown below.

```
$ godic apple
--<apple>--
リンゴ
```

Case of plural arguments.

```
$ godic apple grape melon
--<apple>--
リンゴ
--<grape>--
ブドウ
ブドウの木
--<melon>--
(マスク)メロン
ウリ科の植物
```

Above result's source is [ejje.weblio.jp](http://ejje.weblio.jp).  
You can select dictionary for search by designating option below.

```
-a, -ant
	search antonym from thesaurus.com
-s, -syn
	search synonym from thesaurus.com
-l, -local
	search from local embedded dictionary
```

The data of embedded dictionary based on ejdic-hand which is public-domain.  

For more details, please refer to following URL: http://kujirahand.com/web-tools/EJDictFreeDL.php

## Data of local dictionary

The data of dictionary is converted to binary for using [jteeuwen/gobindata](https://github.com/jteeuwen/go-bindata).  

If you want to modify the data, you have to generate godic/local/data.go again.  

The instruction.

```
go get -u github.com/jteeuwen/go-bindata/...
cd ${GOPATH}/src/github.com/zuiurs/godic
go-bindata -o local/data.go -pkg local data/dict_data.txt
```

## TODO

- standardize return value of each dictionary api.
- extend weblio api.
- add an option serializing output
- add interactive mode

## License

This software is released under the MIT License, see LICENSE.txt.
