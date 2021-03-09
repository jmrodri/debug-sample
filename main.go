/*


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

package main

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap/zapcore"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
	crconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
	cfg      *rest.Config
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	// +kubebuilder:scaffold:scheme
}

func main() {
	ctrl.SetLogger(zap.New(zap.UseDevMode(true), zap.Level(zapcore.DebugLevel)))

	cfg, err := crconfig.GetConfig()
	if err != nil {
		setupLog.Error(err, "getconfig failed: %v\n")
		os.Exit(-3)
	}

	opts := crclient.Options{}
	c, err := crclient.New(cfg, opts)
	if err != nil {
		setupLog.Error(err, "getconfig failed: %v\n")
		os.Exit(-4)
	}

	namespaceList := &corev1.NamespaceList{}
	options := crclient.ListOptions{}
	if err := c.List(context.Background(), namespaceList, &options); err != nil {
		setupLog.Error(err, "ERROR! %v\n")
		os.Exit(-1)
	}

	setupLog.V(1).Info("level 1 debug")
	setupLog.V(2).Info("level 2 debug")   // never prints out :(
	setupLog.V(3).Info("level 3 debug")   // never prints out :(
	setupLog.V(4).Info("level 4 debug")   // never prints out :(
	setupLog.V(5).Info("level 5 debug")   // never prints out :(
	setupLog.V(6).Info("level 6 debug")   // never prints out :(
	setupLog.V(7).Info("level 7 debug")   // never prints out :(
	setupLog.V(8).Info("level 8 debug")   // never prints out :(
	setupLog.V(9).Info("level 9 debug")   // never prints out :(
	setupLog.V(10).Info("level 10 debug") // never prints out :(
	for i, ns := range namespaceList.Items {
		setupLog.Info(fmt.Sprintf("Namespace:: %d is %v", i, ns.Name))
	}
}
