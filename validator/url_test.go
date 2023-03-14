package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidation(t *testing.T) {
	name := "URLTestValidation"
	validUrls := []string{
		"http://k8s-vrl-local.test/api/v1/resultaten/57d60435-1c82-469f-ace6-934963338192",
		"https://k8s-vrl-local.test/api/v1/resultaten/57d60435-1c82-469f-ace6-934963338192",
		"http://referentielijsten-api.vng.cloud//api/v1/resultaten/57d60435-1c82-469f-ace6-934963338192",
		"http://referentielijsten-api.test.vng.cloud//api/v1/resultaten/57d60435-1c82-469f-ace6-934963338192",
		"https://referentielijsten-api.test.vng.cloud//api/v1/resultaten/57d60435-1c82-469f-ace6-934963338192",
		"https://referentielijsten-api.vng.cloud//api/v1/resultaten/57d60435-1c82-469f-ace6-934963338192",
	}

	badUrls := []string{
		"esjtfhnwskedjfn",
		"",
		"//-10:10/foobar",
		"//..",
		"//lll...lll",
		"//foobar....com",
		"\\\\123.4.5.6",
		"123.4.5.6",
		"123.34/16",
		"ftp://ftp.is.co.za/rfc/rfc1808.txt",
		"ldap://[2001:db8::7]/c=GB?objectClass?one",
		"mailto:John.Doe@example.com",
		"news:comp.infosystems.www.servers.unix",
		"tel:+1-816-555-1212",
		"telnet://192.0.2.16:80/",
		"urn:oasis:names:specification:docbook:dtd:xml:4.1.2",
	}

	t.Run("ValidUrl", func(t *testing.T) {
		for _, url := range validUrls {
			sut := Url(url, name)
			assert.Nil(t, sut)
		}
	})

	t.Run("KubeUrl", func(t *testing.T) {
		kubeUrl := "http://vrl:8000/api/v1/resultaten/57d60435-1c82-469f-ace6-934963338192"
		sut := Url(kubeUrl, name)
		assert.Nil(t, sut)
	})

	t.Run("Localhost", func(t *testing.T) {
		local := "http://localhost:8000/api/v1/resultaten/57d60435-1c82-469f-ace6-934963338192"
		sut := Url(local, name)
		assert.Nil(t, sut)
	})

	t.Run("BadUrl", func(t *testing.T) {
		for _, url := range badUrls {
			sut := Url(url, name)
			assert.NotNil(t, sut)
			assert.Equal(t, name, sut.Name)
		}
	})
}
