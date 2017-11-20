package config

import (
  "encoding/json"
  "io/ioutil"
)

// Data is a struct that stores information regarding a user's setup.
type Data struct {
  Dimensions map[string]Dimension
  // CurrentDimension indicates the current default Dimension.
  CurrentDimension string
}

// HasDimensionNamed checks whether there is a Dimension named name.
func (d Data) HasDimensionNamed(name string) bool {
  _, ok := d.Dimensions[name]
  return ok
}

// DataFromJSON converts contents to data.
func DataFromJSON(contents []byte, data *Data)  error {
  err := json.Unmarshal(contents, data)
  return err
}

// DataToGlobalDataJSONFile write d to global data JSON file.
func (d Data) DataToGlobalDataJSONFile() error {
  dataJSON, err := d.toJSON()
  if err != nil {
    return err
  }

  err = ioutil.WriteFile(GetGlobalDataFile(), dataJSON, 0644)
  return err
}

func (d Data) toJSON() ([]byte, error) {
  return json.MarshalIndent(d, "", "  ") // 2-spaces indentation
}