package gmqconf

import (
	"errors"
	"gmq/queue"
	"net"
)

const (
	DEFAULT_LISTEN_PORT = "4884"
	DEFAULT_PROTOCOL    = "tcp"
	DEFAULT_INET        = ""
)

var server *Server

type Server struct {
	Proto, LocalInet, Port string
	Listener               net.Listener
}

func init() {
	server = new(Server)
}

func ConfigureQueue(conf *Params) (queue gmq.QueueInterface, err error) {

	if conf.Queue.MaxQueueN < 1 {
		err = errors.New("Please configure MAX_QUEUE_NUMBER with a positive number")
	}
	if conf.Queue.MaxQueueC < 1 {
		err = errors.New("Please configure MAX_QUEUE_CAPACITY with a positive number")
	}
	if conf.Queue.MaxMessageL < 1 {
		err = errors.New("Please configure MAX_MESSAGE_LENGHT with a positive number")
	}

	switch conf.Queue.QueueType {
	case USE_MEMORY:
		mq := gmq.Queue{}
		mq.Init(conf.Queue.MaxQueueC)
		queue = mq

	case USE_DATABASE:
		queue = gmq.DbQueue{}

	case USE_FILESYSTEM:
		queue = gmq.FsQueue{}

	default:
		err = errors.New("Please configure QUEUE_TYPE with 1 (memory), 2 (database) or 3 (filesystem)")
	}

	return queue, err
}

func configureServer(conf *Params) *Server {
	var inet, proto, port string
	if conf.Network.Port == "" {
		port = DEFAULT_LISTEN_PORT
	} else {
		port = conf.Network.Port
	}
	if conf.Network.Proto == "" {
		proto = DEFAULT_PROTOCOL
	} else {
		proto = conf.Network.Proto
	}
	if conf.Network.Inet == "" {
		inet = DEFAULT_INET
	} else {
		inet = conf.Network.Inet
	}
	return &Server{
		Port:      port,
		Proto:     proto,
		LocalInet: inet,
	}
}

func InitServer(params *Params) (server *Server, err error) {
	server = configureServer(params)
	server.Listener, err = net.Listen(server.Proto, server.LocalInet+":"+server.Port)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (server *Server) StopServer() {
	server.Listener.Close()
}
