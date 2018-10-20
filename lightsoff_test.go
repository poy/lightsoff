package lightsoff_test

import (
	"sync"

	"github.com/poy/lightsoff"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lightsoff", func() {
	var (
		lightsOff  *lightsoff.LightsOff
		funcCalled chan struct{}
	)

	BeforeEach(func() {
		funcCalled = make(chan struct{})
		lightsOff = lightsoff.New(5, func() {
			close(funcCalled)
		})
	})

	It("calls the callback function once everyone is done", func() {
		for i := 0; i < 5; i++ {
			go func() {
				lightsOff.TurnOff()
			}()
		}

		Eventually(funcCalled).Should(BeClosed())
	})

	It("calls the callback only once", func() {
		var wg sync.WaitGroup
		wg.Add(50)
		defer wg.Wait()

		for i := 0; i < 50; i++ {
			go func() {
				defer wg.Done()
				lightsOff.TurnOff()
			}()
		}

		Eventually(funcCalled).Should(BeClosed())
	})

	Context("Incorrectly initialized", func() {
		It("panics if count is less than 0", func() {
			f := func() {
				lightsoff.New(-5, func() {})
			}
			Expect(f).To(Panic())
		})

		It("panics if the callback is nil", func() {
			f := func() {
				lightsoff.New(5, nil)
			}
			Expect(f).To(Panic())
		})

	})

})
