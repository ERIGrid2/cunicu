package crypto_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"riasc.eu/wice/internal/test"
	"riasc.eu/wice/pkg/crypto"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crypto Suite")
}

var _ = test.SetupLogging()

var _ = Describe("nonce", func() {
	It("can generate a valid nonce", func() {
		nonce, err := crypto.GetNonce(100)
		Expect(err).To(Succeed())
		Expect(nonce).To(HaveLen(100))
		Expect([]byte(nonce)).To(test.BeRandom())
	})
})
