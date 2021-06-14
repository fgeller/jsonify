package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `
Usage:

    jsonify [[-|=]name value]...

    Converts arguments into JSON output.

Details:

    -name causes the value to be interpreted as a string.
    =name causes the value to be interpreted as a JSON value.

    If the value is a valid file path, it's contents are used as the value.

Examples:

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

More info:

    https://github.com/fgeller/jsonify

`)
}

func main() {

	if len(os.Args) == 1 {
		printUsage()
		os.Exit(1)
	}

	if len(os.Args)%2 == 0 {
		fmt.Fprintf(os.Stderr, "Expecting even number of arguments.")
		printUsage()
		os.Exit(1)
	}

	var data map[string]interface{} = map[string]interface{}{}
	for idx := 1; idx < len(os.Args); idx += 2 {
		name := os.Args[idx][1:]
		value := os.Args[idx+1]
		_, err := os.Stat(value)
		if !os.IsNotExist(err) {
			bs, err := os.ReadFile(value)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not read contents of file %s. err=%s", value, err)
				os.Exit(1)
			}
			value = string(bs)
		}

		if os.Args[idx][0:1] == "=" {
			var o interface{}
			err = json.Unmarshal([]byte(value), &o)
			data[name] = o
		} else {
			data[name] = value
		}
	}

	bs, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal to json. err=%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", bs)
}
