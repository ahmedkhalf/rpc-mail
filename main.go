package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/emersion/go-imap/client"

	pb "./protomail"
	"google.golang.org/grpc"
)

var clientCount uint64 // Used for uid generation
var clients = make(map[uint64]*client.Client)

type server struct {
	pb.UnimplementedMailServer
}

func (s *server) ConnectClient(ctx context.Context, in *pb.ConnectRequest) (*pb.ConnectResponce, error) {
	server, useTLS := in.GetServer(), in.GetUseTLS()
	log.Println("Connecting to", server, "TLS:", useTLS)

	var c *client.Client
	var err error

	if useTLS {
		c, err = client.DialTLS(server, nil)
	} else {
		c, err = client.Dial(server)
	}

	if err == nil {
		clientCount++
		clients[clientCount] = c
	}

	log.Println("Connected to", server)
	return &pb.ConnectResponce{ClientPointer: clientCount}, err
}

func (s *server) LoginClient(ctx context.Context, in *pb.LoginRequest) (*pb.Empty, error) {
	cid, username, password := in.GetClientPointer(), in.GetUsername(), in.GetPassword()
	log.Println("Logging in to", username)

	err := clients[cid].Login(username, password)

	log.Println("Logged in to", username)
	return &pb.Empty{}, err
}

func (s *server) SelectMailbox(ctx context.Context, in *pb.SelectMailboxRequest) (*pb.Empty, error) {
	cid, name, readonly := in.GetClientPointer(), in.GetName(), in.GetReadonly()

	var err error
	if val, ok := clients[cid]; ok {
		_, err = val.Select(name, readonly)
	} else {
		err = errors.New("cannot call SelectMailbox: client pointer previously destroyed")
	}

	log.Println("Selected ", name)

	return &pb.Empty{}, err
}

func (s *server) DestroyClient(ctx context.Context, in *pb.DestroyRequest) (*pb.Empty, error) {
	cid := in.GetClientPointer()

	delete(clients, cid)
	log.Println("Destroyed client", cid)

	return &pb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	s := &server{}

	grpcServer := grpc.NewServer()
	pb.RegisterMailServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
