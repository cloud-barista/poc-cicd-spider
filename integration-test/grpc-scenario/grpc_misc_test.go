package grpcscenario

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGrpcMisc(t *testing.T) {
	t.Run("grpc api misc test", func(t *testing.T) {
		SetUpForGrpc()

		// CimApi Set/Get Testing
		CimApi.SetServerAddr("127.0.0.1")
		res, _ := CimApi.GetServerAddr()
		assert.True(t, "127.0.0.1" == res)

		CimApi.SetTLSCA("./tls.ca")
		res, _ = CimApi.GetTLSCA()
		assert.True(t, "./tls.ca" == res)

		sec, _ := time.ParseDuration("10s")
		CimApi.SetTimeout(sec)
		tt, _ := CimApi.GetTimeout()
		assert.True(t, sec == tt)

		CimApi.SetJWTToken("abcdefg")
		res, _ = CimApi.GetJWTToken()
		assert.True(t, "abcdefg" == res)

		CimApi.SetInType("json")
		res, _ = CimApi.GetInType()
		assert.True(t, "json" == res)

		CimApi.SetInType("yaml")
		res, _ = CimApi.GetInType()
		assert.True(t, "yaml" == res)

		err := CimApi.SetInType("text")
		assert.True(t, err != nil)

		CimApi.SetOutType("json")
		res, _ = CimApi.GetOutType()
		assert.True(t, "json" == res)

		CimApi.SetOutType("yaml")
		res, _ = CimApi.GetOutType()
		assert.True(t, "yaml" == res)

		err = CimApi.SetOutType("text")
		assert.True(t, err != nil)

		// CcmApi Set/Get Testing
		CcmApi.SetServerAddr("127.0.0.1")
		res, _ = CcmApi.GetServerAddr()
		assert.True(t, "127.0.0.1" == res)

		CcmApi.SetTLSCA("./tls.ca")
		res, _ = CcmApi.GetTLSCA()
		assert.True(t, "./tls.ca" == res)

		sec, _ = time.ParseDuration("10s")
		CcmApi.SetTimeout(sec)
		tt, _ = CcmApi.GetTimeout()
		assert.True(t, sec == tt)

		CcmApi.SetJWTToken("abcdefg")
		res, _ = CcmApi.GetJWTToken()
		assert.True(t, "abcdefg" == res)

		CcmApi.SetInType("json")
		res, _ = CcmApi.GetInType()
		assert.True(t, "json" == res)

		CcmApi.SetInType("yaml")
		res, _ = CcmApi.GetInType()
		assert.True(t, "yaml" == res)

		err = CcmApi.SetInType("text")
		assert.True(t, err != nil)

		CcmApi.SetOutType("json")
		res, _ = CcmApi.GetOutType()
		assert.True(t, "json" == res)

		CcmApi.SetOutType("yaml")
		res, _ = CcmApi.GetOutType()
		assert.True(t, "yaml" == res)

		err = CcmApi.SetOutType("text")
		assert.True(t, err != nil)

		TearDownForGrpc()
	})
}
