# structparser
Parse golang code using ast to get structs information. Useful to use in code generation.

# example

Given a `example/simple_struct.go`. 

```
package example

import "time"

// Simple structure is a simple struct
// Represents a user record in the database
type SimpleStruct struct {
	// Id is the user's id
	ID             int        `db:"id"`
	Name           string     `db:"name"`
	FavoriteColors []string   `db:"favorite_colors"`
	DateUpdated    *time.Time `db:"date_updated"` // only if the record has been updated
}

```

Using this library:

```
package main
import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/wricardo/structparser"
)

func main() {
	parsed, err := structparser.ParseDirectoryWithFilter("./example/simple_struct.go", nil)
	if err != nil {
		log.Fatal(err)
	}

	pretty, _ := json.MarshalIndent(parsed, "", "\t")
	fmt.Println("parsed", string(pretty))
}
```

Will produce this output:
```
parsed [
        {
                "Name": "SimpleStruct",
                "Fields": [
                        {
                                "Name": "ID",
                                "Type": "int",
                                "Tag": "`db:\"id\"`",
                                "Pointer": false,
                                "Slice": false,
                                "Docs": [
                                        "Id is the user's id"
                                ],
                                "Comment": ""
                        },
                        {
                                "Name": "Name",
                                "Type": "string",
                                "Tag": "`db:\"name\"`",
                                "Pointer": false,
                                "Slice": false,
                                "Docs": null,
                                "Comment": ""
                        },
                        {
                                "Name": "FavoriteColors",
                                "Type": "[]string",
                                "Tag": "`db:\"favorite_colors\"`",
                                "Pointer": false,
                                "Slice": true,
                                "Docs": null,
                                "Comment": ""
                        },
                        {
                                "Name": "DateUpdated",
                                "Type": "*time.Time",
                                "Tag": "`db:\"date_updated\"`",
                                "Pointer": true,
                                "Slice": false,
                                "Docs": null,
                                "Comment": "only if the record has been updated"
                        }
                ],
                "Docs": [
                        "Simple structure is a simple struct",
                        "Represents a user record in the database"
                ]
        }
]
```
