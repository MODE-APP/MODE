package servers

import "go.mongodb.org/mongo-driver/mongo"

type DatabaseServer struct {
	TLSserver
	*mongo.Client
}
