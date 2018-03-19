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

	Context("check block validate", func() {
		It("correct block should validate successfully", func() {
			firstBlock := Block{0, "0", 0, "0a1e2340446d3e7dd298d5ed01c102811d4f7da34255f287b242d0eb7471e352", "", difficulty, 0}
			newBlock, _ := BlockchainInstance.GenerateBlock(1)
			Expect(newBlock.IsValidBlock(firstBlock)).To(Equal(true))
		})
	})

})
