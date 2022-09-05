package van

import (
	"github.com/recsys-infra/ps-lite/utils"
	"math"
	"os"
	"strconv"
	"sync"

	"github.com/recsys-infra/ps-lite/api"
	"github.com/recsys-infra/ps-lite/pkg/postoffice"

	"github.com/sirupsen/logrus"
)

type Van struct {
	scheduler   *api.Node
	myNode      *api.Node
	isScheduler bool

	startMux sync.Mutex

	initStage int
}

func (v *Van) Start(customerId int) error {
	v.startMux.Lock()
	defer v.startMux.Unlock()

	if v.initStage == 0 {
		v.scheduler.Hostname = os.Getenv("DMLC_PS_ROOT_URI")
		port, err := strconv.ParseInt(os.Getenv("DMLC_PS_ROOT_PORT"), 10, 32)
		if err != nil {
			logrus.WithError(err).WithField("port", os.Getenv("DMLC_PS_ROOT_PORT")).Info("failed to parse port")
			return err
		}
		v.scheduler.Port = int32(port)
		v.scheduler.Role = api.Role_SCHEDULER
		v.isScheduler = postoffice.Singleton.IsScheduler()

		if v.isScheduler {
			v.myNode = v.scheduler
		} else {
			role := func() api.Role {
				if postoffice.Singleton.IsWorker() {
					return api.Role_WORKER
				} else {
					return api.Role_SERVER
				}
			}()
			ip, err := utils.GetAvailableIpV4()
			if err != nil {
				logrus.WithError(err).Error("failed to get avaliable ip address")
				return err
			}
			v.myNode.Hostname = ip
			v.myNode.Role = role
			v.myNode.Port = 8080
			v.myNode.Id = math.MaxInt32
			v.myNode.CustomId = int32(customerId)
		}
	}
	return nil
}
