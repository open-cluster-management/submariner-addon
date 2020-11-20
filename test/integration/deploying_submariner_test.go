package integration

import (
	"context"
	"fmt"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/open-cluster-management/submariner-addon/test/utils"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
)

const (
	expectedCRDsWork     = "submariner-agent-crds"
	expectedRBACWork     = "submariner-agent-rbac"
	expectedOperatorWork = "submariner-agent-operator"

	expectedSCCWork     = "submariner-scc"
	expectedSCCRBACWork = "submariner-scc-rbac"
)

var _ = ginkgo.Describe("Deploy a submariner on hub", func() {
	var managedClusterSetName string
	var managedClusterName string

	ginkgo.BeforeEach(func() {
		managedClusterSetName = fmt.Sprintf("set-%s", rand.String(6))
		managedClusterName = fmt.Sprintf("cluster-%s", rand.String(6))
	})

	ginkgo.Context("Deploy submariner agent manifestworks", func() {
		var expectedBrokerNamespace string

		ginkgo.BeforeEach(func() {
			ginkgo.By("Create a ManagedClusterSet")
			managedClusterSet := utils.NewManagedClusterSet(managedClusterSetName)
			_, err := clusterClient.ClusterV1alpha1().ManagedClusterSets().Create(context.Background(), managedClusterSet, metav1.CreateOptions{})
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Check if the submariner broker is deployed")
			expectedBrokerNamespace = fmt.Sprintf("submariner-clusterset-%s-broker", managedClusterSetName)
			gomega.Eventually(func() bool {
				return utils.FindSubmarinerBrokerResources(kubeClient, expectedBrokerNamespace)
			}, eventuallyTimeout, eventuallyInterval).Should(gomega.BeTrue())
		})

		ginkgo.It("Should deploy the submariner agent manifestworks on managed cluster namespace successfully", func() {
			ginkgo.By("Create a ManagedCluster")
			managedCluster := utils.NewManagedCluster(managedClusterName, map[string]string{
				"cluster.open-cluster-management.io/submariner-agent": "true",
				"cluster.open-cluster-management.io/clusterset":       managedClusterSetName,
			})
			_, err := clusterClient.ClusterV1().ManagedClusters().Create(context.Background(), managedCluster, metav1.CreateOptions{})
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Setup the managed cluster namespace")
			_, err = kubeClient.CoreV1().Namespaces().Create(context.Background(), utils.NewManagedClusterNamespace(managedClusterName), metav1.CreateOptions{})
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Setup the serviceaccount")
			err = utils.SetupServiceAccount(kubeClient, expectedBrokerNamespace, managedClusterName)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Check if the submariner agent manifestworks are deployed")
			gomega.Eventually(func() bool {
				return utils.FindManifestWorks(workClient, managedClusterName, expectedCRDsWork, expectedRBACWork, expectedOperatorWork)
			}, eventuallyTimeout, eventuallyInterval).Should(gomega.BeTrue())
		})

		ginkgo.It("Should deploy submariner agent manifestworks with vendor label on managed cluster namespace successfully", func() {
			ginkgo.By("Create a ManagedCluster with vendor label")
			managedCluster := utils.NewManagedCluster(managedClusterName, map[string]string{
				"cluster.open-cluster-management.io/submariner-agent": "true",
				"cluster.open-cluster-management.io/clusterset":       managedClusterSetName,
				"vendor": "OCP",
			})
			_, err := clusterClient.ClusterV1().ManagedClusters().Create(context.Background(), managedCluster, metav1.CreateOptions{})
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Create the managedcluster namespace")
			_, err = kubeClient.CoreV1().Namespaces().Create(context.Background(), utils.NewManagedClusterNamespace(managedClusterName), metav1.CreateOptions{})
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Setup the serviceaccount")
			err = utils.SetupServiceAccount(kubeClient, expectedBrokerNamespace, managedClusterName)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Check if the submariner agent manifestworks are deployed")
			gomega.Eventually(func() bool {
				return utils.FindManifestWorks(workClient, managedClusterName, expectedCRDsWork, expectedRBACWork, expectedOperatorWork, expectedSCCWork, expectedSCCRBACWork)
			}, eventuallyTimeout, eventuallyInterval).Should(gomega.BeTrue())
		})
	})

	ginkgo.Context("Remove submariner agent manifestworks", func() {
		ginkgo.BeforeEach(func() {
			ginkgo.By("Create a ManagedClusterSet")
			managedClusterSet := utils.NewManagedClusterSet(managedClusterSetName)
			_, err := clusterClient.ClusterV1alpha1().ManagedClusterSets().Create(context.Background(), managedClusterSet, metav1.CreateOptions{})
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			brokerNamespace := fmt.Sprintf("submariner-clusterset-%s-broker", managedClusterSetName)
			gomega.Eventually(func() bool {
				return utils.FindSubmarinerBrokerResources(kubeClient, brokerNamespace)
			}, eventuallyTimeout, eventuallyInterval).Should(gomega.BeTrue())

			ginkgo.By("Create a ManagedCluster")
			managedCluster := utils.NewManagedCluster(managedClusterName, map[string]string{
				"cluster.open-cluster-management.io/submariner-agent": "true",
				"cluster.open-cluster-management.io/clusterset":       managedClusterSetName,
			})
			_, err = clusterClient.ClusterV1().ManagedClusters().Create(context.Background(), managedCluster, metav1.CreateOptions{})
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Setup the managed cluster namespace")
			_, err = kubeClient.CoreV1().Namespaces().Create(context.Background(), utils.NewManagedClusterNamespace(managedClusterName), metav1.CreateOptions{})
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Setup the serviceaccount")
			err = utils.SetupServiceAccount(kubeClient, brokerNamespace, managedClusterName)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			gomega.Eventually(func() bool {
				return utils.FindManifestWorks(workClient, managedClusterName, expectedCRDsWork, expectedRBACWork, expectedOperatorWork)
			}, eventuallyTimeout, eventuallyInterval).Should(gomega.BeTrue())
		})

		ginkgo.It("Should remove the submariner agent manifestworks after the submariner label is removed from the managed cluster", func() {
			ginkgo.By("Remove the submariner label from the managed cluster")
			newLabels := map[string]string{"cluster.open-cluster-management.io/clusterset": managedClusterSetName}
			err := utils.UpdateManagedClusterLabels(clusterClient, managedClusterName, newLabels)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Check if the submariner agent manifestworks are removed")
			gomega.Eventually(func() bool {
				works, err := workClient.WorkV1().ManifestWorks(managedClusterName).List(context.Background(), metav1.ListOptions{})
				if err != nil {
					return false
				}
				return len(works.Items) == 0
			}, eventuallyTimeout, eventuallyInterval).Should(gomega.BeTrue())
		})

		ginkgo.It("Should remove the submariner agent manifestworks after the managedclusterset label is removed from the managed cluster", func() {
			ginkgo.By("Remove the managedclusterset label from the managed cluster")
			newLabels := map[string]string{"cluster.open-cluster-management.io/submariner-agent": "true"}
			err := utils.UpdateManagedClusterLabels(clusterClient, managedClusterName, newLabels)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Check if the submariner agent manifestworks are removed")
			gomega.Eventually(func() bool {
				works, err := workClient.WorkV1().ManifestWorks(managedClusterName).List(context.Background(), metav1.ListOptions{})
				if err != nil {
					return false
				}
				return len(works.Items) == 0
			}, eventuallyTimeout, eventuallyInterval).Should(gomega.BeTrue())
		})

		ginkgo.It("Should remove the submariner agent manifestworks after the managedcluster is removed", func() {
			ginkgo.By("Remove the managedcluster")
			err := clusterClient.ClusterV1().ManagedClusters().Delete(context.Background(), managedClusterName, metav1.DeleteOptions{})
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Check if the submariner agent manifestworks are removed")
			gomega.Eventually(func() bool {
				works, err := workClient.WorkV1().ManifestWorks(managedClusterName).List(context.Background(), metav1.ListOptions{})
				if err != nil {
					return false
				}
				return len(works.Items) == 0
			}, eventuallyTimeout, eventuallyInterval).Should(gomega.BeTrue())
		})

	})

	ginkgo.Context("Remove submariner broker", func() {
		ginkgo.It("Should remove the submariner broker after the managedclusterset is removed", func() {
			ginkgo.By("Create a ManagedClusterSet")
			managedClusterSet := utils.NewManagedClusterSet(managedClusterSetName)
			_, err := clusterClient.ClusterV1alpha1().ManagedClusterSets().Create(context.Background(), managedClusterSet, metav1.CreateOptions{})
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Check if the submariner broker is deployed")
			brokerNamespace := fmt.Sprintf("submariner-clusterset-%s-broker", managedClusterSetName)
			gomega.Eventually(func() bool {
				return utils.FindSubmarinerBrokerResources(kubeClient, brokerNamespace)
			}, eventuallyTimeout, eventuallyInterval).Should(gomega.BeTrue())

			ginkgo.By("Remove the managedclusterset")
			err = clusterClient.ClusterV1alpha1().ManagedClusterSets().Delete(context.Background(), managedClusterSetName, metav1.DeleteOptions{})
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			ginkgo.By("Check if the submariner broker is removed")
			gomega.Eventually(func() bool {
				ns, err := kubeClient.CoreV1().Namespaces().Get(context.Background(), brokerNamespace, metav1.GetOptions{})
				if errors.IsNotFound(err) {
					return true
				}
				if err != nil {
					return false
				}

				// the controller-runtime does not have a gc controller, so if the namespace is in terminating and
				// there is no broker finalizer on it, it will be consider as removed
				if ns.Status.Phase == corev1.NamespaceTerminating &&
					!utils.FindExpectedFinalizer(ns.Finalizers, "cluster.open-cluster-management.io/submariner-cleanup") {
					return true
				}
				return false
			}, eventuallyTimeout, eventuallyInterval).Should(gomega.BeTrue())
		})
	})
})
