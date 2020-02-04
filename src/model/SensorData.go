package model

// SensorData represents a data package for sensor data linked to a Node.
type SensorData struct {
	NODE_ID    	string 	`json:"_id" xml:"_id" bson:"_id"`
	TIME    	string 	`json:"time" xml:"time" bson:"time"`
	TEMP 		string 	`json:"temp" xml:"temp" bson:"temp"`
	PRESSURE   	string 	`json:"pressure" xml:"pressure" bson:"pressure"`
	HUMIDITY   	string	`json:"humidity" xml:"humidity" bson:"humidity"`
}
