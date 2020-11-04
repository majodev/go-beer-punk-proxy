package mailer_test

import (
	"context"
	"testing"

	"github.com/majodev/go-beer-punk-proxy/internal/api"
	"github.com/majodev/go-beer-punk-proxy/internal/config"
	"github.com/majodev/go-beer-punk-proxy/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMailerSendPasswordReset(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fixtures := test.Fixtures()

	m := test.NewTestMailer(t)
	passwordResetLink := "http://localhost/password/reset/12345"
	err := m.SendPasswordReset(ctx, fixtures.User1.Username.String, passwordResetLink)
	require.NoError(t, err)

	mt := test.GetTestMailerMockTransport(t, m)
	mail := mt.GetLastSentMail()
	mails := mt.GetSentMails()
	require.NotNil(t, mail)
	require.Len(t, mails, 1)
	assert.Equal(t, m.Config.DefaultSender, mail.From)
	assert.Len(t, mail.To, 1)
	assert.Equal(t, fixtures.User1.Username.String, mail.To[0])
	assert.Equal(t, test.TestMailerDefaultSender, mail.From)
	assert.Equal(t, "Password reset", mail.Subject)
	assert.Contains(t, string(mail.HTML), passwordResetLink)
}

func SkipTestMailerSendPasswordResetWithMailhog(t *testing.T) {
	t.Skip()
	t.Parallel()

	ctx := context.Background()
	fixtures := test.Fixtures()

	m := test.NewSMTPMailerFromDefaultEnv(t)

	passwordResetLink := "http://localhost/password/reset/12345"
	err := m.SendPasswordReset(ctx, fixtures.User1.Username.String, passwordResetLink)
	require.NoError(t, err)
}

func SkipTestMailerSendPasswordResetWithMailhogAndServer(t *testing.T) {
	t.Skip()
	t.Parallel()

	ctx := context.Background()
	fixtures := test.Fixtures()

	defaultConfig := config.DefaultServiceConfigFromEnv()
	defaultConfig.Mailer.Transporter = config.MailerTransporterSMTP.String()
	test.WithTestServerConfigurable(t, defaultConfig, func(s *api.Server) {
		passwordResetLink := "http://localhost/password/reset/12345"
		err := s.Mailer.SendPasswordReset(ctx, fixtures.User1.Username.String, passwordResetLink)
		require.NoError(t, err)
	})
}
