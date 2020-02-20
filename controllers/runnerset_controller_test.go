package controllers

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"math/rand"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	actionsv1alpha1 "github.com/summerwind/actions-runner-controller/api/v1alpha1"
)

// SetupTest will set up a testing environment.
// This includes:
// * creating a Namespace to be used during the test
// * starting the 'RunnerReconciler'
// * stopping the 'RunnerSetReconciler" after the test ends
// Call this function at the start of each of your tests.
func SetupTest(ctx context.Context) *corev1.Namespace {
	var stopCh chan struct{}
	ns := &corev1.Namespace{}

	BeforeEach(func() {
		stopCh = make(chan struct{})
		*ns = corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{Name: "testns-" + randStringRunes(5)},
		}

		err := k8sClient.Create(ctx, ns)
		Expect(err).NotTo(HaveOccurred(), "failed to create test namespace")

		mgr, err := ctrl.NewManager(cfg, ctrl.Options{})
		Expect(err).NotTo(HaveOccurred(), "failed to create manager")

		controller := &RunnerSetReconciler{
			Client:   mgr.GetClient(),
			Scheme:   scheme.Scheme,
			Log:      logf.Log,
			Recorder: mgr.GetEventRecorderFor("runnerset-controller"),
		}
		err = controller.SetupWithManager(mgr)
		Expect(err).NotTo(HaveOccurred(), "failed to setup controller")

		go func() {
			err := mgr.Start(stopCh)
			Expect(err).NotTo(HaveOccurred(), "failed to start manager")
		}()
	})

	AfterEach(func() {
		close(stopCh)

		err := k8sClient.Delete(ctx, ns)
		Expect(err).NotTo(HaveOccurred(), "failed to delete test namespace")
	})

	return ns
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func intPtr(v int) *int {
	return &v
}

var _ = Context("Inside of a new namespace", func() {
	ctx := context.TODO()
	ns := SetupTest(ctx)

	Describe("when no existing resources exist", func() {

		It("should create a new Runner resource from the specified template and one replica if one replica is provided", func() {
			name := "example-runnerset"

			rs := &actionsv1alpha1.RunnerSet{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: ns.Name,
				},
				Spec: actionsv1alpha1.RunnerSetSpec{
					Replicas: intPtr(1),
					Template: actionsv1alpha1.RunnerSpec{
						Repository: "foo",
						Image:      "bar",
						Env:        nil,
					},
				},
			}

			err := k8sClient.Create(ctx, rs)
			Expect(err).NotTo(HaveOccurred(), "failed to create test Foo resource")

			runners := &actionsv1alpha1.RunnerList{}
			Eventually(
				getResourceFunc(ctx, client.ObjectKey{Name: name, Namespace: rs.Namespace}, runners),
				time.Second*5, time.Millisecond*500).Should(BeNil())

			Expect(len(runners.Items)).To(Equal(1))
		})
	})
})

func getResourceFunc(ctx context.Context, key client.ObjectKey, obj runtime.Object) func() error {
	return func() error {
		return k8sClient.Get(ctx, key, obj)
	}
}
