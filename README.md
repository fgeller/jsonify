# jsonify

Some reasons why you might be interested:

* Quickly produce JSON output based on command line arguments
* Simple syntax to interpret values as strings or arbitrary JSON values
* Supports reading file contents for easy escaping

## Installation

* Downloads are available from the [Releases](https://github.com/fgeller/jsonify/releases) section.
* `go get github.com/fgeller/jsonify && go install github.com/fgeller/jsonify`
* Via homebrew on OSX: `brew tap fgeller/tap && brew install jsonify`

## Usage

    jsonify [[-|=]name value]...

Converts arguments into JSON output.

## Details

* `-name` causes the value to be interpreted as a string.
* `=name` causes the value to be interpreted as a JSON value.
* If the value is a valid file path, it's contents are used as the value.

## Examples

    $ # basic value types, ie - vs =
    $ jsonify -name hans =age 23 =subscribed true =address null | jq
    {
      "address": null,
      "age": 23,
      "name": "hans",
      "subscribed": true
    }

    $ # nested objects via command substitution
    $ jsonify =a `jsonify -name hans` =b `jsonify -name peter` | tee outfile | jq
    {
      "a": {
        "name": "hans"
      },
      "b": {
        "name": "peter"
      }
    }

    $ # subshell output as a value to get current date
    $ # reading contents of "outfile" from previous invocation
    $ jsonify -date "$(date)" =content outfile | jq
    {
      "content": {
        "a": {
          "name": "hans"
        },
        "b": {
          "name": "peter"
        }
      },
      "date": "Thu Mar 17 19:10:04 NZDT 2016"
    }

