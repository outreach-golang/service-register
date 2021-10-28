package service_register

import (
	"context"
	"errors"
	"github.com/outreach-golang/etcd"
	"os"
	"sync"
)

var (
	ServiceRegisterHandler *ServiceRegister
)

type ServiceRegister struct {
	serviceName string
	nodeIP      string
	lease       int64
	init        sync.Once
	err         error
}

func init() {
	ServiceRegisterHandler = newServiceRegister()
}

func newServiceRegister() *ServiceRegister {
	return &ServiceRegister{
		nodeIP: "http://" + os.Getenv("NODE_IP"),
		lease:  30,
		init:   sync.Once{},
		err:    nil,
	}
}

func (s *ServiceRegister) InitServiceRegister(ctx context.Context, sr *etcd.ServiceRegister, serviceName string,
	port string,
) error {
	s.init.Do(func() {
		if s.nodeIP == "" {

			s.err = errors.New("获取 NODE_IP 失败！")

			return
		}

		if serviceName == "" {
			s.err = errors.New("注册服务时 ServiceName 不能为空！")

			return
		}

		accessAddress := s.nodeIP + ":" + port

		if err := sr.Register(ctx, serviceName, accessAddress, s.lease); err != nil {
			s.err = err
		}
	})

	return s.err
}
