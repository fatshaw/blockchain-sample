package ws

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/fatshaw/blockchain-sample/blockchain"
)

var _ = Describe("response processor", func() {

	Context("append block", func() {
		It("should append block", func() {
			newBlock, _ := BlockchainInstance.GenerateBlock(1)
			message := &BlockchainWsResponse{Blocks: []Block{newBlock}}
			response := ResponseProcessor{}
			response.OnMessage(message)
			Expect(len(BlockchainInstance.Blocks)).To(Equal(2))
		})
	})
})
