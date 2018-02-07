package blockchain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func doprint(p Printer) string {
	return p.Print()
}


var _ = Describe("lock string function", func() {

	Context("block string function", func() {
		It("block string function should work", func() {
			b := Block{1, "1", 1, "1", "1"}
			Expect(doprint(&b)).To(Equal("index=1,timestamp=1,bpm=1,hash=1,prehash=1"))
		})
	})

	Context("bigblock string function", func() {
		It("bigblock string function should work", func() {
			b := BigBlock{Block{1, "1", 1, "1", "1"}, 1}
			Expect(doprint(&b)).To(Equal("index=1,timestamp=1,bpm=1,hash=1,prehash=1,big=1"))
		})
	})

})
