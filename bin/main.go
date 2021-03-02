package main

import (

  "encoding/json"

  "strings"

  "fmt"

  "log"

  "github.com/graphql-go/graphql"
)

func main() {

    countries := make(map[string]string)
    countries["ag"] = "Argentina"
    countries["au"] = "Australia"
    countries["be"] = "Belgium"
    countries["br"] = "Brazil"
    countries["ca"] = "Canada"
    countries["mx"] = "Mexico"
    countries["cu"] = "Cuba"
    countries["nl"] = "Netherlands"
    countries["br"] = "Britian"
    countries["de"] = "Germany"

      keys := make([]string, 0, len(countries))
      values := make([]string, 0, len(countries))
      for key, val := range countries {
        keys = append(keys, key)
        values = append(values, val)
      }
    // Schema
    fields := graphql.Fields{
        "Abbr": &graphql.Field{
            Type: graphql.NewList(graphql.String),
            Args: graphql.FieldConfigArgument{
                "startWith": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
              results := make([]string, 0)
              id, ok := p.Args["startWith"].(string)
                if ok {
                    // Parse our tutorial array for the matching id
                    for _, item := range keys {
                        if strings.HasPrefix(item, id) {
                            // return our tutorial
                            results = append(results, item)
                        }
                    }
                }
                return results, nil
            },
        },
        "Name": &graphql.Field{
            Type: graphql.NewList(graphql.String),
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return values, nil
            },
        },
    }
    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
    schema, err := graphql.NewSchema(schemaConfig)
    if err != nil {
        log.Fatalf("failed to create new schema, error: %v", err)
    }

    // Query
    query := `
        {
            Abbr(startWith: "b")
            Name
        }
    `
    params := graphql.Params{Schema: schema, RequestString: query}
    r := graphql.Do(params)
    if len(r.Errors) > 0 {
        log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
    }
    rJSON, _ := json.Marshal(r)
    fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}
