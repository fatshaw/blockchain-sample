package blockchain

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
	. "github.com/fatshaw/blockchain-sample/util"
)

var _ = Describe("block hash test", func() {

	Context("block hash test", func() {
		It("block hash test should work", func() {
			b := Block{0, "0", 0, "", ""}
			Logger.Info(b.CalculateHash())
		})
	})

})
