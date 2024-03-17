/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_model/go"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	demov1alpha1 "github.com/930c/sleep-operator/api/v1alpha1"
)

var (
	reconcileDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "operator_reconcile_duration_seconds",
		Help: "Duration of reconcile loops in seconds.",
	}, []string{"controller"})
)

var _ = Describe("Sleep Controller", func() {
	Context("When reconciling a resource", func() {
		const resourceName = "test-resource"

		ctx := context.Background()

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default", // TODO(user):Modify as needed
		}
		sleep := &demov1alpha1.Sleep{}

		BeforeEach(func() {
			By("creating the custom resource for the Kind Sleep")
			err := k8sClient.Get(ctx, typeNamespacedName, sleep)
			if err != nil && errors.IsNotFound(err) {
				resource := &demov1alpha1.Sleep{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
					// TODO(user): Specify other spec details if needed.
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
			prometheus.MustRegister(reconcileDuration)
		})

		AfterEach(func() {
			// TODO(user): Cleanup logic after each test, like removing the resource instance.
			resource := &demov1alpha1.Sleep{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			By("Cleanup the specific resource instance Sleep")
			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
		})
		It("should successfully reconcile the resource and update the reconcileDuration metric", func() {
			By("Reconciling the created resource")
			controllerReconciler := &SleepReconciler{
				Client:            k8sClient,
				Scheme:            k8sClient.Scheme(),
				ReconcileDuration: reconcileDuration,
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())
			// TODO(user): Add more specific assertions depending on your controller's reconciliation logic.
			// Example: If you expect a certain status condition after reconciliation, verify it here.

			By("Gathering all registered metrics")
			metrics, err := prometheus.DefaultGatherer.Gather()
			Expect(err).NotTo(HaveOccurred())

			By("Finding the reconcileDuration metric")
			var reconcileDurationMetric *io_prometheus_client.MetricFamily
			for _, m := range metrics {
				if m.GetName() == "operator_reconcile_duration_seconds" {
					reconcileDurationMetric = m
					break
				}
			}

			By("Checking the reconcileDuration metric")
			Expect(reconcileDurationMetric).NotTo(BeNil())
			Expect(reconcileDurationMetric.GetMetric()).ToNot(BeEmpty())
		})
	})
})
