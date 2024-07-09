package service

import "testing"

func TestStartHealthCheckReport(t *testing.T) {
	StartHealthCheckReport(map[string]interface{}{
		"project": "aaa",
	})

}
