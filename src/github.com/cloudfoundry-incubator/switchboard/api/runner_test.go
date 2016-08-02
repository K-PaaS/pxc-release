package api_test

import (
	"fmt"
	"net"
	"os"

	"github.com/cloudfoundry-incubator/switchboard/api"
	"github.com/cloudfoundry-incubator/switchboard/config"
	"github.com/pivotal-golang/lager/lagertest"

	"github.com/cloudfoundry-incubator/switchboard/api/apifakes"
	"github.com/cloudfoundry-incubator/switchboard/domain/domainfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/ifrit"
)

var _ = Describe("APIRunner", func() {
	It("shuts down gracefully when signalled", func() {
		apiPort := 10000 + GinkgoParallelNode()
		backends := new(domainfakes.FakeBackends)
		logger := lagertest.NewTestLogger("APIRunner Test")
		config := config.API{}
		staticDir := ""
		cluster := new(apifakes.FakeClusterManager)
		handler := api.NewHandler(cluster, backends, logger, config, staticDir)
		apiRunner := api.NewRunner(uint(apiPort), handler, logger)
		apiProcess := ifrit.Invoke(apiRunner)
		apiProcess.Signal(os.Kill)
		Eventually(apiProcess.Wait()).Should(Receive())

		_, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", apiPort))
		Expect(err).To(HaveOccurred())
	})
})
