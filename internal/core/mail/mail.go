package mail

/*Service is a mail service that provides you
to send mails with smtp protocol.
*/
type Service interface {
	SendTextMail(to, subject, content string) error
}
