package skeleton_test

import (
	// . "github.com/kieron-pivotal/baby-app/system-tests/skeleton"

	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("Skeleton", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("can see the home page", func() {
		url := os.Getenv("APP_URL")
		Expect(url).ToNot(BeEmpty())

		Expect(page.Navigate(url)).To(Succeed())
		Eventually(page.FirstByClass("App")).Should(MatchText("hello, world"))
	})
})
