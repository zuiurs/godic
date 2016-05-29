# godic

godic is a simple English-Japanese dictionary on command-line.

## Usage

An usage is shown below.

	# godic apple
	apple:  『リンゴ』;リンゴの木

Case of plural arguments.

	# godic gopher gopher gopher
	gopher:  地リス,地ネズミ(北米の草原にすむ穴居性リス)
	gopher:  地リス,地ネズミ(北米の草原にすむ穴居性リス)
	gopher:  地リス,地ネズミ(北米の草原にすむ穴居性リス)

The data of dictionary based on ejdic-hand which is public-domain.

For more details, please refer to following URL: http://kujirahand.com/web-tools/EJDictFreeDL.php

## Memo

The data of dictionary is converted to binary for using [jteeuwen/gobindata](https://github.com/jteeuwen/go-bindata).

###### TODO

- Do something about generating the map  struct every time.
