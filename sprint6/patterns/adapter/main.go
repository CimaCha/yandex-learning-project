package main

// JSONData — интерфейс для декодирования JSON.
type JSONData interface {
	DecodeJSON() interface{}
}

// YAMLData — интерфейс для декодирования YAML.
type YAMLData interface {
	DecodeYAML() interface{}
}

type Client struct {
	Data interface{}
}

func (client *Client) Decode(input JSONData) {
	client.Data = input.DecodeJSON()
}

type Adapter struct {
	data YAMLData
}

func (adapter *Adapter) DecodeJSON() interface{} {
	return adapter.data.DecodeYAML()
}

func Load(client Client, input YAMLData) {
	adapter := &Adapter{
		data: input,
	}
	client.Decode(adapter)
}
