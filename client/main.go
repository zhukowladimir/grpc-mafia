package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/zhukowladimir/grpc-mafia/pkg/proto/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	serverAddr = flag.String("addr", "localhost:8888", "The server address in the format of host:port")
)

type mafiaClient struct {
	service service.MafiaUsersClient
	name    string
}

func (mc *mafiaClient) Login(reader *bufio.Reader) {
	for ok, text := false, "Enter your username: "; !ok; {
		fmt.Print(text)
		name, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
			continue
		}

		name = name[:len(name)-1]
		mc.name = name

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := mc.service.SetUser(ctx, &service.StringRequest{Msg: name})
		if err != nil {
			log.Fatalln(err)
		}
		ok = response.Ack
		text = response.Msg

		log.Println(fmt.Sprintf("SetUser: ACK: %v, %s", ok, text))
	}
}

func (mc *mafiaClient) Listen() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := mc.service.ListenNotifications(ctx, &service.StringRequest{Msg: mc.name})
	if err != nil {
		log.Fatalf("problems with start notification listening: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return
		} else if err != nil {
			log.Fatalf("problems with receiveng notifications: %v\n", err)
			return
		}
		fmt.Printf("SYSTEM: %s\n", msg.Msg)
	}
}

func (mc *mafiaClient) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := mc.service.Disconnect(ctx, &service.StringRequest{Msg: mc.name})
	if err != nil {
		log.Fatalf("problems with disconnecting: %v\n", err)
	}
}

func (mc *mafiaClient) CmdPlayers() *[]string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	msgList, err := mc.service.GetUsersList(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("problems with getting users list: %v\n", err)
		return nil
	}
	return &msgList.List
}

func main() {
	flag.Parse()

	log.Println("Client running...")
	var client = mafiaClient{}

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client.service = service.NewMafiaUsersClient(conn)

	reader := bufio.NewReader(os.Stdin)

	client.Login(reader)
	defer client.Disconnect()

	go client.Listen()

	for {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
			continue
		}

		cmd = cmd[:len(cmd)-1]
		if cmd == "!exit" {
			break
		} else if cmd == "!players" {
			list := client.CmdPlayers()
			if list != nil {
				for _, name := range *list {
					fmt.Println(name)
				}
				fmt.Println()
			}
		}
	}
}
