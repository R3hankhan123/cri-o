package server_test

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

// The actual test suite.
var _ = t.Describe("ImageFsInfo", func() {
	// Prepare the sut
	BeforeEach(func() {
		beforeEach()
		setupSUT()
	})
	AfterEach(afterEach)

	t.Describe("ImageFsInfo", func() {
		It("should succeed", func() {
			// Given
			gomock.InOrder(
				imageServerMock.EXPECT().GetStore().Return(storeMock),
				storeMock.EXPECT().GraphRoot().Return(""),
				storeMock.EXPECT().ImageStore().Return(""),
				storeMock.EXPECT().GraphDriverName().Return("test"),
			)
			testImageDir := "test-images"
			Expect(os.MkdirAll(testImageDir, 0o755)).To(Succeed())
			defer os.RemoveAll(testImageDir)

			// When
			response, err := sut.ImageFsInfo(context.Background(), nil)

			// Then
			Expect(err).ToNot(HaveOccurred())
			Expect(response).NotTo(BeNil())
			Expect(len(response.GetImageFilesystems())).To(BeEquivalentTo(1))
			Expect(len(response.GetContainerFilesystems())).To(BeEquivalentTo(1))
		})

		It("should fail on invalid image dir", func() {
			// Given
			gomock.InOrder(
				imageServerMock.EXPECT().GetStore().Return(storeMock),
				storeMock.EXPECT().GraphRoot().Return(""),
				storeMock.EXPECT().ImageStore().Return(""),
				storeMock.EXPECT().GraphDriverName().Return(""),
			)

			// When
			response, err := sut.ImageFsInfo(context.Background(), nil)

			// Then
			Expect(err).To(HaveOccurred())
			Expect(response).To(BeNil())
		})
	})
})
