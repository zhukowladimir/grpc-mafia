package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/zhukowladimir/grpc-mafia/pkg/proto/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	port = flag.Int("port", 8888, "The server port")
)

type server struct {
	service.UnimplementedMafiaServer

	curSession Session
	mutex      sync.Mutex
}

func (s *server) SetUser(_ context.Context, request *service.StringRequest) (*service.ACK, error) {
	name := request.Msg
	log.Printf("SetUser: StringRequest: %s", name)

	s.mutex.Lock()
	ack, err := s.curSession.IsValidUsername(name)
	if ack {
		s.curSession.NotifyPlayers(&Notification{ntype: NOTIFICATION_USER_JOIN, msg: name})
		s.curSession.players[name] = *CreatePlayer(name)
	}
	s.mutex.Unlock()

	err_str := ""
	if err != nil {
		err_str = err.Error()
	}

	return &service.ACK{Ack: ack, Msg: err_str}, nil
}

func (s *server) GetUsersList(_ context.Context, _ *emptypb.Empty) (*service.UsersListResponse, error) {
	return &service.UsersListResponse{List: s.curSession.GetPlayersList()}, nil
}

func (s *server) ListenNotifications(request *service.StringRequest, stream service.Mafia_ListenNotificationsServer) error {
	username := request.Msg
	for {
		msg, err := s.curSession.TakePlayerNotification(username)
		if err != nil {
			log.Fatalf("Client %s has problems with notifications: %v\n", username, err)
			return nil
		}
		notificationMsg := ""
		switch msg.ntype {
		case NOTIFICATION_USER_JOIN:
			notificationMsg += "Player " + msg.msg + " joined!"
		case NOTIFICATION_USER_LEAVE:
			notificationMsg += "Player " + msg.msg + " leaved!"
		}
		if err := stream.Send(&service.Notification{Msg: notificationMsg}); err != nil {
			return err
		}
	}
}

func (s *server) Disconnect(_ context.Context, request *service.StringRequest) (*emptypb.Empty, error) {
	name := request.Msg
	log.Println(fmt.Sprintf("Disconnect: StringRequest: %s", name))

	delete(s.curSession.players, name)
	s.curSession.NotifyPlayers(&Notification{ntype: NOTIFICATION_USER_LEAVE, msg: name})

	out := new(emptypb.Empty)
	return out, nil
}

func (s *server) StartGame(_ context.Context, request *service.StringRequest) (*emptypb.Empty, error) {
	s.curSession.ready++

	if s.curSession.ready == len(s.curSession.players) {
		go s.curSession.StartGame()
	}

	out := new(emptypb.Empty)
	return out, nil
}

func (s *server) LynchingVote(_ context.Context, vote *service.VoteRequest) (*service.ACK, error) {
	if _, ok := s.curSession.game.vote[vote.To]; !ok {
		s.curSession.game.vote[vote.To] = 0
	}
	s.curSession.game.vote[vote.To] += 1
	return &service.ACK{Ack: true, Msg: ""}, nil
}

func (s *server) GoSleep(_ context.Context, request *service.StringRequest) (*emptypb.Empty, error) {
	// TODO
	out := new(emptypb.Empty)
	return out, nil
}

func (s *server) NightMurder(_ context.Context, vote *service.VoteRequest) (*service.ACK, error) {
	// TODO
}

func (s *server) SneakPeek(_ context.Context, vote *service.VoteRequest) (*service.ACK, error) {
	// TODO
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	thisServer := server{curSession: *createSession("The only one server", 1)}

	srv := grpc.NewServer()
	service.RegisterMafiaServer(srv, &thisServer)

	log.Printf("Server listening at %v\n", lis.Addr().String())

	if err := srv.Serve(lis); err != nil {
		log.Fatalln("gg")
		log.Fatalln(err)
	}
}
