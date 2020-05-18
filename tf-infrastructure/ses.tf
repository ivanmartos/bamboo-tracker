resource "aws_ses_email_identity" "dev_email" {
  email = var.dev_ses_sender_email
}
