package post_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sharing-vision-2021/common"
	"sharing-vision-2021/domain/post"
	mock_repository "sharing-vision-2021/mocks/repository"
	"time"
)

var errSomething = errors.New("something error")

var _ = Describe("Post Service", func() {
	var (
		mockCtrl  *gomock.Controller
		postUC    post.Service
		repo      *mock_repository.MockRepository
		mockPost  post.Post
		mockPosts []post.Post
		now       time.Time
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		repo = mock_repository.NewMockRepository(mockCtrl)
		postUC = post.NewPostImplementation(repo)
		now = time.Now()

		mockPost = post.Post{
			ID:          1,
			Title:       "Test title",
			Content:     "Test Content",
			Category:    "mock",
			CreatedDate: &now,
			UpdatedDate: &now,
			Status:      "Publish",
		}

		mockPosts = []post.Post{
			{
				ID:          1,
				Title:       "Test title",
				Content:     "Test Content",
				Category:    "mock",
				CreatedDate: &now,
				UpdatedDate: &now,
				Status:      "Publish",
			}, {
				ID:          2,
				Title:       "Test title",
				Content:     "Test Content",
				Category:    "mock",
				CreatedDate: &now,
				UpdatedDate: &now,
				Status:      "Publish",
			},
		}

	})

	Describe("AddPost", func() {
		mockRequest := post.Post{
			Title:    "Test title",
			Content:  "Test Content",
			Category: "mock",
			Status:   "Publish",
		}
		It("return success", func() {
			repo.EXPECT().AddPost(gomock.Any()).Return(nil)
			err := postUC.AddPost(&mockRequest)
			Expect(err).Should(Succeed())
		})

		It("return error", func() {
			repo.EXPECT().AddPost(gomock.Any()).Return(errSomething)
			err := postUC.AddPost(&mockRequest)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("GetList", func() {
		mockLimit := 10
		mockOffset := 0
		It("return success", func() {
			repo.EXPECT().GetList(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockPosts, int64(2), nil)
			posts, counts, err := postUC.GetList(mockLimit, mockOffset, "")
			Expect(err).Should(Succeed())
			Expect(posts).Should(Equal(mockPosts))
			Expect(counts).Should(Equal(int64(2)))
		})

		It("return error", func() {
			repo.EXPECT().GetList(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, int64(0), errSomething)
			_, _, err := postUC.GetList(mockLimit, mockOffset, "")
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("Get", func() {
		mockID := 1
		It("return success", func() {
			repo.EXPECT().Get(gomock.Any()).Return(&mockPost, nil)
			post, err := postUC.Get(mockID)
			Expect(err).Should(Succeed())
			Expect(post).Should(Equal(&mockPost))
			Expect(post.ID).Should(Equal(mockID))
		})

		It("return error", func() {
			repo.EXPECT().Get(gomock.Any()).Return(nil, errSomething)
			_, err := postUC.Get(mockID)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("Update", func() {
		mockRequest := post.Post{
			Title:    "Test title",
			Content:  "Test Content",
			Category: "mock",
			Status:   "Publish",
		}
		mockID := 1
		It("return success", func() {
			repo.EXPECT().Get(gomock.Any()).Return(&mockPost, nil).AnyTimes()
			repo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			_, err := postUC.Get(mockID)
			err = postUC.Update(mockID, &mockRequest)
			Expect(err).Should(Succeed())
		})

		It("return not found", func() {
			repo.EXPECT().Get(gomock.Any()).Return(nil, common.ErrRecordNotFound).AnyTimes()
			repo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errSomething)
			_, err := postUC.Get(mockID)
			err = postUC.Update(mockID, &mockRequest)
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(Equal(common.ErrRecordNotFound))
		})

		It("return error", func() {
			repo.EXPECT().Get(gomock.Any()).Return(&mockPost, nil).AnyTimes()
			repo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errSomething)
			_, err := postUC.Get(mockID)
			err = postUC.Update(mockID, &mockRequest)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("Delete", func() {
		mockID := 1
		It("return success", func() {
			repo.EXPECT().Get(gomock.Any()).Return(&mockPost, nil).AnyTimes()
			repo.EXPECT().Delete(gomock.Any()).Return(nil)
			_, err := postUC.Get(mockID)
			err = postUC.Delete(mockID)
			Expect(err).Should(Succeed())
		})

		It("return not found", func() {
			repo.EXPECT().Get(gomock.Any()).Return(nil, common.ErrRecordNotFound).AnyTimes()
			repo.EXPECT().Delete(gomock.Any()).Return(errSomething)
			err := postUC.Delete(mockID)
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(Equal(common.ErrRecordNotFound))
		})

		It("return error", func() {
			repo.EXPECT().Get(gomock.Any()).Return(&mockPost, nil).AnyTimes()
			repo.EXPECT().Delete(gomock.Any()).Return(errSomething)
			_, err := postUC.Get(mockID)
			err = postUC.Delete(mockID)
			Expect(err).Should(HaveOccurred())
		})
	})
})
