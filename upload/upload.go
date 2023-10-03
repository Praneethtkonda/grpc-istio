package upload

import (
	"go_client_istio/pb"

	"context"
	"fmt"
	"io"
	"os"
	"time"
)

func Upload(client pb.ClientRequestClient, ctx context.Context, file string) (string, error) {
	// Context is used for holding up the connection only till 30 seconds here for instance
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(20*time.Second))
	defer cancel()

	stream, err := client.UploadToServer(ctx)
	if err != nil {
		return "", err
	}

	fil, err := os.Open(file)
	fmt.Println(fil)
	if err != nil {
		return "", err
	}

	// Maximum 1KB size per stream.
	buf := make([]byte, 1024)

	for {
		num, err := fil.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		fmt.Println("Sending stream")
		if err := stream.Send(&pb.UploadRequest{Chunk: buf[:num]}); err != nil {
			return "", err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return "", err
	}

	return res.GetName(), nil
}
