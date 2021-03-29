package main

import (
	"context"
	"log"
	"net"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"

	pb "./protomail"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMailServer

	clientCount uint64 // Used for uid generation
	clients map[uint64]*client.Client
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
		s.clientCount++
		s.clients[s.clientCount] = c
	}

	log.Println("Connected to", server)
	return &pb.ConnectResponce{ClientPointer: s.clientCount}, err
}

func (s *server) LoginClient(ctx context.Context, in *pb.LoginRequest) (*pb.Empty, error) {
	cid, username, password := in.GetClientPointer(), in.GetUsername(), in.GetPassword()
	log.Println("Logging in to", username)

	err := s.clients[cid].Login(username, password)

	log.Println("Logged in to", username)
	return &pb.Empty{}, err
}

func (s *server) LogoutClient(ctx context.Context, in *pb.LogoutRequest) (*pb.Empty, error) {
	cid := in.GetClientPointer()
	log.Printf("Logging out (client %v)", cid)

	err := s.clients[cid].Logout()

	return &pb.Empty{}, err
}

func (s *server) DeleteClient(ctx context.Context, in *pb.DeleteRequest) (*pb.Empty, error) {
	cid := in.GetClientPointer()

	delete(s.clients, cid)
	log.Println("Destroyed client", cid)

	return &pb.Empty{}, nil
}

func (s *server) ListMailboxes(in *pb.ListMailboxesRequest, stream pb.Mail_ListMailboxesServer) error {
	cid := in.GetClientPointer()

	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func () {
		done <- s.clients[cid].List("", "*", mailboxes)
	}()

	for m := range mailboxes {
		stream.Send(&pb.ListMailboxesResponce{MailboxName: m.Name})
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *server) SelectMailbox(ctx context.Context, in *pb.SelectMailboxRequest) (*pb.Empty, error) {
	cid, name, readonly := in.GetClientPointer(), in.GetMailboxName(), in.GetReadonly()

	_, err := s.clients[cid].Select(name, readonly)

	log.Println("Selected ", name)

	return &pb.Empty{}, err
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	s := &server{}
	s.clients = make(map[uint64]*client.Client)

	grpcServer := grpc.NewServer()
	pb.RegisterMailServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
