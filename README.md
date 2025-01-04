# Go-Mail

A learning project demonstrating how to send emails using Go's `net/smtp` package. This repository serves as an educational resource for understanding email functionality implementation in Go.

## 📚 What I Learned

- Basic email sending using Go's `net/smtp` package
- Working with environment variables in Go
- Structuring a Go project
- Handling email templates
- SMTP configuration and authentication

## 📁 Project Structure

```
go-mail/
├── messages/       # Email message definitions
├── service/        # Email service implementation
├── templates/      # Email templates
├── .env.example    # Environment variables template
├── .gitignore
├── go.mod
├── go.sum
└── main.go
```

## 🚀 Getting Started

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-mail.git
cd go-mail
```

2. Set up your environment variables:
```bash
cp .env.example .env
```

3. Update `.env` with your email settings:
```env
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SENDER_EMAIL=your_email
SENDER_PASSWORD=your_password
```

### Gmail Setup

For Gmail users:
1. Enable 2-Step Verification in your Google Account
2. Generate an App Password:
   - Go to Google Account Settings > Security
   - Under "Signing in to Google," select App Passwords
   - Generate a new app password for "Mail"
3. Use the app password in your `.env` file


## 📝 Notes

This is a learning project and may not be suitable for production use. Feel free to use it as a reference for understanding how to implement email functionality in Go.
