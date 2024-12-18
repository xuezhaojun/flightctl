package tasks_test

import (
	"context"

	"github.com/flightctl/flightctl/internal/tasks"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FleetSelector", func() {
	var (
		ctx           context.Context
		orgId         uuid.UUID
		configStorage tasks.ConfigStorage
	)

	BeforeEach(func() {
		ctx = context.Background()
		orgId, _ = uuid.NewUUID()
		var err error
		configStorage, err = tasks.NewConfigStorage("localhost", 6379)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		configStorage.DeleteAllKeys(ctx)
		configStorage.Close()
	})

	When("fetching a git revision", func() {
		It("returns what is stored if the key exists", func() {
			key := tasks.ConfigStorageGitRevisionKey{
				OrgID:           orgId,
				Fleet:           "myfleet",
				TemplateVersion: "mytv",
				Repository:      "myrepo",
				TargetRevision:  "main",
			}

			updated, err := configStorage.SetNX(ctx, key.ComposeKey(), []byte("abc123"))
			Expect(err).ToNot(HaveOccurred())
			Expect(updated).To(BeTrue())

			updated, err = configStorage.SetNX(ctx, key.ComposeKey(), []byte("def456"))
			Expect(err).ToNot(HaveOccurred())
			Expect(updated).To(BeFalse())

			hash, err := configStorage.Get(ctx, key.ComposeKey())
			Expect(err).ToNot(HaveOccurred())
			Expect(hash).To(Equal([]byte("abc123")))
		})

		It("returns an empty string if the key doesn't exist", func() {
			key := tasks.ConfigStorageGitRevisionKey{
				OrgID:           orgId,
				Fleet:           "myfleet",
				TemplateVersion: "mytv",
				Repository:      "myrepo",
				TargetRevision:  "main",
			}

			hash, err := configStorage.Get(ctx, key.ComposeKey())
			Expect(err).ToNot(HaveOccurred())
			Expect(hash).To(HaveLen(0))
		})
	})

	When("setting a repo URL", func() {
		It("stores what is passed if the key doesn't exist", func() {
			key := tasks.ConfigStorageRepositoryUrlKey{
				OrgID:           orgId,
				Fleet:           "myfleet",
				TemplateVersion: "mytv",
				Repository:      "myrepo",
			}

			url, err := configStorage.GetOrSetNX(ctx, key.ComposeKey(), []byte("https://myurl"))
			Expect(err).ToNot(HaveOccurred())
			Expect(url).To(Equal([]byte("https://myurl")))
		})

		It("returns what is stored if the key exists", func() {
			key := tasks.ConfigStorageRepositoryUrlKey{
				OrgID:           orgId,
				Fleet:           "myfleet",
				TemplateVersion: "mytv",
				Repository:      "myrepo",
			}

			url, err := configStorage.GetOrSetNX(ctx, key.ComposeKey(), []byte("https://myurl"))
			Expect(err).ToNot(HaveOccurred())
			Expect(url).To(Equal([]byte("https://myurl")))

			url, err = configStorage.GetOrSetNX(ctx, key.ComposeKey(), []byte("https://otherurl"))
			Expect(err).ToNot(HaveOccurred())
			Expect(url).To(Equal([]byte("https://myurl")))
		})
	})

	When("deleting a TemplateVersion", func() {
		It("deletes all its related keys", func() {
			key := tasks.ConfigStorageRepositoryUrlKey{
				OrgID:           orgId,
				Fleet:           "myfleet",
				TemplateVersion: "mytv",
				Repository:      "myrepo",
			}

			updated, err := configStorage.SetNX(ctx, key.ComposeKey(), []byte("https://myurl"))
			Expect(err).ToNot(HaveOccurred())
			Expect(updated).To(BeTrue())
			key.TemplateVersion = "othertv"
			updated, err = configStorage.SetNX(ctx, key.ComposeKey(), []byte("https://otherurl"))
			Expect(err).ToNot(HaveOccurred())
			Expect(updated).To(BeTrue())

			key2 := tasks.ConfigStorageGitRevisionKey{
				OrgID:           orgId,
				Fleet:           "myfleet",
				TemplateVersion: "mytv",
				Repository:      "myrepo",
				TargetRevision:  "main",
			}

			updated, err = configStorage.SetNX(ctx, key2.ComposeKey(), []byte("abc123"))
			Expect(err).ToNot(HaveOccurred())
			Expect(updated).To(BeTrue())
			key2.TemplateVersion = "othertv"
			updated, err = configStorage.SetNX(ctx, key2.ComposeKey(), []byte("def456"))
			Expect(err).ToNot(HaveOccurred())
			Expect(updated).To(BeTrue())

			tvkey := tasks.TemplateVersionKey{OrgID: orgId, Fleet: "myfleet", TemplateVersion: "mytv"}
			err = configStorage.DeleteKeysForTemplateVersion(ctx, tvkey.ComposeKey())
			Expect(err).ToNot(HaveOccurred())

			ret, err := configStorage.Get(ctx, key.ComposeKey())
			Expect(err).ToNot(HaveOccurred())
			Expect(ret).To(Equal([]byte("https://otherurl")))
			ret, err = configStorage.Get(ctx, key2.ComposeKey())
			Expect(err).ToNot(HaveOccurred())
			Expect(ret).To(Equal([]byte("def456")))

			key.TemplateVersion = "mytv"
			key2.TemplateVersion = "mytv"
			ret, err = configStorage.Get(ctx, key.ComposeKey())
			Expect(err).ToNot(HaveOccurred())
			Expect(ret).To(BeEmpty())
			ret, err = configStorage.Get(ctx, key2.ComposeKey())
			Expect(err).ToNot(HaveOccurred())
			Expect(ret).To(BeEmpty())
		})
	})
})
