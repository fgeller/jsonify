package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `
Usage:

    jsonify [[-|=]name value]...

    Converts arguments into JSON output.

    jsonify -validate data.json -schema data_schema.json

    Validates contents of the file data.json agains json schema in data_schema.json

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

func validate(schema string, json string) error {
	schemaLoader := gojsonschema.NewStringLoader(schema)
	documentLoader := gojsonschema.NewStringLoader(json)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}
	if result.Valid() {
		return nil
	} else {
		descs := make([]string, len(result.Errors()))
		for _, desc := range result.Errors() {
			descs = append(descs, desc.Description())
		}
		return errors.New(strings.Join(descs, "\n"))
	}
}

func runValidate() error {
	schemaArg := flag.String("schema", "", "file with json schema")
	validateArg := flag.String("validate", "", "file with json for validation")
	flag.Parse()

	schemaFile := *schemaArg
	jsonFile := *validateArg

	jsonText, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return fmt.Errorf("Could not read file [%s]: %v", jsonFile, err)
	}
	schema, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("Could not read file [%s]: %v", schemaFile, err)
	}

	return validate(string(schema), string(jsonText))
}

func main() {

	if len(os.Args) == 1 {
		printUsage()
		os.Exit(1)
	}

	if os.Args[1] == "-validate" || os.Args[1] == "-schema" {
		err := runValidate()
		if err != nil {
			log.Fatal(err)
		}
		return
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
			bs, err := ioutil.ReadFile(value)
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
