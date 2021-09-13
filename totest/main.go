package main

import (
	"context"
	"google.golang.org/grpc"
	v1 "grpc-url-shortener/pkg/api/v1"
	"log"
	"time"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8083", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewUrlShortnerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	// Call Create
	req1 := v1.CreateRequest{
		Api: "v1",
		Url: "https://www.google.com/search?q=make+a+post+request+to+grpc+endpoint&client=ubuntu&hs=0lj&channel=fs&ei=5ME9YYCaBq2C9u8Pyv-KkAE&oq=make+a+post+request+to+grpc+endpoint&gs_lcp=Cgdnd3Mtd2l6EAMyBQghEKABOgQIABATOgYIABAKEBM6CAgAEA0QHhATOggIABAWEB4QEzoGCAAQFhAeOggIIRAWEB0QHjoHCCEQChCgAUoECEEYAFCOF1ifSGDnSWgBcAJ4AIABkAGIAZkfkgEEMS4zNZgBAKABAbABAMABAQ&sclient=gws-wiz&ved=0ahUKEwiAnKu-ifnyAhUtgf0HHcq_AhIQ4dUDCA0&uact=5",
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)


	getRes, err := c.Get(ctx, &v1.GetRequest{Api: "v1", ShortURL: res1.ShortUrl})
	if err!=nil{
		log.Fatalf("Retrieve failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", getRes)
}
