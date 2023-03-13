package manifestutils

import (
	"github.com/ViaQ/logerr/v2/kverrors"
	"github.com/imdario/mergo"
	corev1 "k8s.io/api/core/v1"

	"github.com/os-observability/tempo-operator/apis/tempo/v1alpha1"
)

func getS3Storage(tempo *v1alpha1.TempoStack) ([]corev1.EnvVar, []string) {
	var environment []corev1.EnvVar = []corev1.EnvVar{
		{
			Name: "S3_SECRET_KEY",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					Key: "access_key_secret",
					LocalObjectReference: corev1.LocalObjectReference{
						Name: tempo.Spec.Storage.Secret.Name,
					},
				},
			},
		},
		{
			Name: "S3_ACCESS_KEY",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					Key: "access_key_id",
					LocalObjectReference: corev1.LocalObjectReference{
						Name: tempo.Spec.Storage.Secret.Name,
					},
				},
			},
		},
	}
	args := []string{
		"--storage.trace.s3.secret_key=$(S3_SECRET_KEY)",
		"--storage.trace.s3.access_key=$(S3_ACCESS_KEY)",
	}
	return environment, args
}

// ConfigureStorage configures storage.
func ConfigureStorage(tempo v1alpha1.TempoStack, pod *corev1.PodSpec) error {
	if tempo.Spec.Storage.Secret.Name != "" {
		ingesterContainer := pod.Containers[0].DeepCopy()
		envVars, args := getS3Storage(&tempo)
		ingesterContainer.Env = append(ingesterContainer.Env, envVars...)
		ingesterContainer.Args = append(ingesterContainer.Args, args...)

		if err := mergo.Merge(&pod.Containers[0], ingesterContainer, mergo.WithOverride); err != nil {
			return kverrors.Wrap(err, "failed to merge ingester container spec")
		}
	}
	return nil
}
