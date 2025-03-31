package mail

import (
	"testing"

	"github.com/insta-app/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	// in Makefile we set the short flag to skip this test,
	// this will skip tests which take longer time to run (like this one)
	// when all tests ran (avoid spamming emails)
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName,
		config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "a test email"
	content := `
	<h1>Hello from Leo Messi</h1>
	<p>this is a message from <a href="leomessi.com"> MESSIIIISSII </a></p>
	`
	to := []string{"moazzankamran110@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
