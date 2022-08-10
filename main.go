package main

import (
	"flag"
	"fmt"
	"net/url"
)

var (
	username      string
	password      string
	cluster       string
	authSource    string
	authMechanism string
)

func main() {
	flag.StringVar(&username, "username", "nicholas.makes", "Okta Username")
	flag.StringVar(&password, "password", "", "Okta Password")
	flag.StringVar(&cluster, "cluster", "baas-local.rvhot.mongodb.net/test", "MongoDB Atlas Cluster")
	flag.StringVar(&authSource, "authsource", "$external", "Auth Source")
	flag.StringVar(&authMechanism, "authmech", "PLAIN", "Auth Mechanism")

	flag.Parse()

	uri := "mongodb+srv://" + url.QueryEscape(username) + "@mongodb.com:" +
		url.QueryEscape(password) + "@" + cluster +
		"/?authSource=" + authSource +
		"&authMechanism=" + authMechanism

	fmt.Println(uri)

	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	// if err != nil {
	// 	panic(err)
	// }
	// defer client.Disconnect(context.TODO())
	//
	// collection := client.Database("app").Collection("apps")
	//
	// cursor, err := collection.Find(context.TODO(), bson.D{})
	// if err != nil {
	// 	panic(err)
	// }
	//
	// var results []bson.D
	// if err = cursor.All(context.TODO(), &results); err != nil {
	// 	panic(err)
	// }
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
}
