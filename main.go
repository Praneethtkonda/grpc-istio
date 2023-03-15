package main

import (
	"fmt"
	"go_signc/pb"
	"go_signc/upload"

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
	defaultName = "world"
	defaultSignMethod = "digestsign"
	defaultKey = "vmware_esx67"
	defaultInputUrl = "/tmp/one.txt"
	defaultOutputPath = "/tmp/out/out.txt"
	version = "2.0.0"
)

var (
	// url = "192.168.49.2:80"
	// url = "grpc.example.com:443"
	// url = "xds://192.168.49.2:15021"
	url = "127.0.0.1:80"
	// addr = flag.String("addr", url, "the address to connect to")
	signmethod string = defaultSignMethod
	keyid string = defaultKey
	inputurl string = defaultInputUrl
	outputurl string = defaultOutputPath
)

func SignClient(args []string) {
	fmt.Println(signmethod, keyid, inputurl, outputurl)
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
	c := pb.NewSigningRequestClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendRequest(ctx, &pb.SignRequest{Signmethod: &signmethod, Keyid: &keyid, InputUrl: &inputurl, OutputPath: &outputurl})
	if err != nil {
		log.Fatalf("Could not send signing request: %v", err)
	}
	log.Printf("Output url: %s", r.GetOutputUrl())
	log.Printf("Status of signmethod: %s", r.GetStatus())

	name, err := upload.Upload(c, context.Background(), inputurl)
	if err != nil {
		log.Printf("Failed uploading the file")
		log.Fatalln(err)
	}
	log.Printf("File name that got upload %s", name)
}

func main() {
	var rootCmd = &cobra.Command{
		Use: "The sign client to provide the signing service",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			SignClient(args)
		},
	}

	var Verbose bool
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&signmethod, "signmethod", "", "", "Sign method")
	rootCmd.PersistentFlags().StringVarP(&keyid, "keyid", "k", "", "Key id")
	rootCmd.PersistentFlags().StringVarP(&inputurl, "inputurl", "i", "", "Input url")
	rootCmd.PersistentFlags().StringVarP(&outputurl, "outputurl", "o", "", "Output url")
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Error occured while running signc %s", err)
		os.Exit(1)
	}
}
