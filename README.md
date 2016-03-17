# jsonify

Some reasons why you might be interested:

* Produces JSON output based on command line arguments
* Simple syntax to interpret values as strings or arbitrary JSON values
* Supports reading file contents for easy escaping

## Usage

    jsonify [[-|=]name value]...

Converts arguments into JSON output.

## Details

* `-name` causes the value to be interpreted as a string.
* `=name` causes the value to be interpreted as a JSON value.
* If the value is a valid file path, it's contents are used as the value.

## Examples

    $ jsonify -first_name hans -last_name schmitt | jq
    {
      "first_name": "hans",
      "last_name": "schmitt"
    }

    $ jsonify =a `+"`"+`jsonify -name hans`+"`"+` =b `+"`"+`jsonify -name peter`+"`"+` | tee out | jq
    {
      "a": {
        "name": "hans"
      },
      "b": {
        "name": "peter"
      }
    }

    $ jsonify -date "$(date)" =content out | jq
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

