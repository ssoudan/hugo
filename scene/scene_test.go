package scene_test

import (
	. "github.com/ssoudan/hugo/scene"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scene", func() {

	DescribeTable("converts RGB to HSL",
		func(r, g, b, hE, sE, lE float32) {
			h, s, l := RGBtoHSL(r, g, b)
			Expect(h).To(Equal(hE))
			Expect(s).To(Equal(sE))
			Expect(l).To(Equal(lE))
		},
		Entry("#FFFFFF", float32(1.), float32(1.), float32(1.), float32(0.), float32(0.), float32(1.)),
		Entry("#000000", float32(0.), float32(0.), float32(0.), float32(0.), float32(0.), float32(0.)),
		Entry("#FF0000", float32(1.), float32(0.), float32(0.), float32(0.), float32(1.), float32(0.5)),
		Entry("#00FF00", float32(0.), float32(1.), float32(0.), float32(1/3.), float32(1.), float32(0.5)),
		Entry("#0000FF", float32(0.), float32(0.), float32(1.), float32(2/3.), float32(1.), float32(0.5)),
	)

})
