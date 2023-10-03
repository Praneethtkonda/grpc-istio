package main

import (
	"fmt"
	"go_client_istio/pb"
	"go_client_istio/upload"

	"context"
	// "flag"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName       = "world"
	defaultMethod     = "registration"
	defaultKey        = "vmware_esx67"
	defaultInputUrl   = "/tmp/one.txt"
	defaultOutputPath = "/tmp/out/out.txt"
	version           = "2.0.0"
)

var (
	// url = "192.168.49.2:80"
	// url = "grpc.example.com:443"
	// url = "xds://192.168.49.2:15021"
	url = "127.0.0.1:80"
	// addr = flag.String("addr", url, "the address to connect to")
	registration string = defaultMethod
	keyid        string = defaultKey
	inputurl     string = defaultInputUrl
	outputurl    string = defaultOutputPath
)

func Client(args []string) {
	fmt.Println(registration, keyid, inputurl, outputurl)
	// Set up a connection to the grpc server.

	// New code
	// creds, err := credentials.NewClientTLSFromFile("tls.crt", "")
	// credentials.New
	// if err != nil {
	// 	log.Fatalf("could not load tls cert: %s", err)
	// }
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Connection to the server failed: %v", err)
	}
	defer conn.Close()

	// Contact the server and print out its response.
	c := pb.NewClientRequestClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendRequest(ctx, &pb.Request{Registration: &registration, Id: &keyid, InputUrl: &inputurl, OutputPath: &outputurl})
	if err != nil {
		log.Fatalf("Could not send client request: %v", err)
	}
	log.Printf("Output url: %s", r.GetOutputUrl())
	log.Printf("Status of method: %s", r.GetStatus())

	name, err := upload.Upload(c, context.Background(), inputurl)
	if err != nil {
		log.Printf("Failed uploading the file")
		log.Fatalln(err)
	}
	log.Printf("File name that got upload %s", name)
}

func main() {
	var rootCmd = &cobra.Command{
		Use:     "The client to send request to the service",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			Client(args)
		},
	}

	var Verbose bool
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&registration, "registration", "r", "", "Registration")
	rootCmd.PersistentFlags().StringVarP(&keyid, "id", "k", "", "id")
	rootCmd.PersistentFlags().StringVarP(&inputurl, "inputurl", "i", "", "Input url")
	rootCmd.PersistentFlags().StringVarP(&outputurl, "outputurl", "o", "", "Output url")
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Error occured while running client %s", err)
		os.Exit(1)
	}
}
