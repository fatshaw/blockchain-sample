package blockchain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("block hash test", func() {

	Context("block hash test", func() {
		It("block hash test should work", func() {
			b := Block{0, "0", 0, "", "", difficulty, 2740103009342231109}
			Expect(b.CalculateHash()).To(Equal("037a114c7a8a2aa7e79151eaef59ec71bd6917073ce3e7a3dfb1b788f3a6618a"))
		})
	})

	Context("first block hash", func() {
		It("block hash should meets the Pow", func() {
			b := Block{0, "0", 0, "", "", difficulty, 0}
			b.TryGenerateHash()
			Expect(strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty))).To(Equal(true))
		})
	})

})
