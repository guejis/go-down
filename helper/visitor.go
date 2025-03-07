package helper

type VisitorData struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

type Visitor struct {
 Data []VisitorData	`json:"data"`
}

func (d *Visitor) Write(key string, value any) {
	for i, values := range d.Data {
		if values.Key == key {
			d.Data[i].Value = value
			return
		}
	}
	d.Data = append(d.Data, VisitorData{
		Key:   key,
		Value: value,
	})
}

func (d *Visitor) Read(key string) any {
	for _, value := range d.Data {
		if value.Key == key {
			return value.Value
		}
	}

	return nil
}

func (d *Visitor) ReadAll() *Visitor {
	return d
}
