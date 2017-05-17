// Copyright © 2017 thingful
// +build integration

package listener

import (
	"testing"

	// TODO : move import to upstream project

	"github.com/stretchr/testify/assert"
	testing_helper "github.com/thingful/device-hub/utils/testing"
)

func TestMQTT_MultipleEndpoints(t *testing.T) {

	t.Parallel()

	environment := testing_helper.MustUp()
	defer environment.Down()

	l, err := newMQTTListener(environment.MQTTClient)
	assert.Nil(t, err)

	channel1, err := l.NewChannel("/a")
	assert.Nil(t, err)

	channel2, err := l.NewChannel("/b")
	assert.Nil(t, err)

	environment.MQTTClient.Publish("/a", 0, false, "hello")
	environment.MQTTClient.Publish("/b", 0, false, "hello")

	message := <-channel1.Out()
	assert.Equal(t, message.Payload, []byte("hello"))

	message2 := <-channel2.Out()
	assert.Equal(t, message2.Payload, []byte("hello"))

}
