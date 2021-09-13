package controller

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	v1 "grpc-url-shortener/pkg/api/v1"
	"testing"
)

func TestController_Create(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	r := require.New(t)
	conn, err := createConnection()
	r.NoError(err)
	r.NotNil(conn)
	t.Cleanup(func() {
		conn.Close()
	})

	c := v1.NewUrlShortnerServiceClient(conn)

	create, err := c.Create(ctx, &v1.CreateRequest{
		Api: "v1",
		Url: "https://www.google.com/search?q=some+loooooong+loooong+loooong+url+python+golang+deveport&client=firefox-b-d&sxsrf=AOaemvIqDJ4hpOzkRC8bhNcacpIVH-tWcQ%3A1631572534560&ei=NtI_YcPKIZj87_UPv7OlmAo&oq=some+loooooong+loooong+loooong+url+python+golang+devepo&gs_lcp=Cgdnd3Mtd2l6EAMYADIHCCEQChCgATIHCCEQChCgATIHCCEQChCgAToHCAAQRxCwAzoHCCMQ6gIQJzoECCMQJzoECAAQQzoOCC4QgAQQsQMQxwEQrwE6DgguEIAEELEDEMcBEKMCOggIABCABBCxAzoGCCMQJxATOgQILhBDOg0ILhCxAxDHARDRAxBDOgsIABCABBCxAxCDAToFCC4QgAQ6CAgAELEDEIMBOgoILhDHARDRAxBDOgUIABCABDoFCAAQywE6BQguEMsBOggILhDLARCTAjoECAAQDToGCAAQDRAeOgUIIRCgAToGCAAQFhAeOggIIRAWEB0QHjoECCEQFUoECEEYAFCSH1j8cWDGe2gCcAF4AoABoQSIAaeBAZIBDTAuMTguOS45LjE0LjSYAQCgAQGwAQrIAQjAAQE&sclient=gws-wiz",
	})
	r.NoError(err)
	r.NotNil(create)
}

func TestController_Get(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	r := require.New(t)
	conn, err := createConnection()
	r.NoError(err)
	r.NotNil(conn)

	t.Cleanup(func() {
		conn.Close()
	})

	c := v1.NewUrlShortnerServiceClient(conn)

	createReq := &v1.CreateRequest{
		Api: "v1",
		Url: "https://stackoverflow.com/questions/51071020/golang-net-listen-binds-to-port-thats-already-in-use",
	}

	create, err := c.Create(ctx, createReq)
	r.NoError(err)
	r.NotNil(create)

	get, err := c.Get(ctx, &v1.GetRequest{
		Api:      "v1",
		ShortURL: create.ShortUrl,
	})
	r.NoError(err)
	r.NotNil(get)
	r.Equal(get.OriginalURL, createReq.Url)

}

func createConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:8083", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, err
}
