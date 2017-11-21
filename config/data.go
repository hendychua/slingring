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

// GetGlobalData converts contents in the global data file to data.
func GetGlobalData() (*Data, error) {
  dataJSONContents, err := ioutil.ReadFile(GetGlobalDataFile())
  if err != nil {
    return nil, err
  }

  data := Data{}
  err = json.Unmarshal(dataJSONContents, &data)
  if err != nil {
    return nil, err
  }

  return &data, nil
}

// DataToGlobalDataJSONFile write d to global data JSON file.
// This rewrites the whole file and can be problematic (slow) when the data gets huge.
// TODO: improve it by writing only the diff.
func (d Data) DataToGlobalDataJSONFile() error {
  dataJSON, err := json.MarshalIndent(d, "", "  ") // 2-spaces indentation
  if err != nil {
    return err
  }

  err = ioutil.WriteFile(GetGlobalDataFile(), dataJSON, 0644)
  return err
}