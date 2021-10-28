package service_register

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/outreach-golang/etcd"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
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

		serviceName = serviceName + "." + getRandomString(12)

		accessAddress := s.nodeIP + ":" + port

		if err := sr.Register(ctx, serviceName, accessAddress, s.lease); err != nil {
			s.err = err
		}
	})

	return s.err
}

func getRandomString(n int) string {
	s := fmt.Sprintf("%x", sha256.Sum256([]byte(uuid.NewV4().String()+strconv.FormatInt(time.Now().UnixNano(), 10))))

	randBytes := make([]byte, len(s)/2)
	rand.Read(randBytes)
	s1 := fmt.Sprintf("%x", randBytes)

	return s[:n-3] + s1[15:18]
}
