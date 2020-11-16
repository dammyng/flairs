package events

type UserCreatedEvent struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Host  string `json:"host"`
}
type WelcomeUserEvent struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Host     string `json:"host"`
	Username string `json:"username"`
}

type PasswordReset struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Host  string `json:"host"`
}

type OTPCreated struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Host  string `json:"host"`
}

type CreateDefWallet struct {
	URL  string `json:"host"`
	UserID  string `json:"userId"`
	Token  string `json:"token"`
}

type CreditWallet struct {
	URL  string `json:"host"`
	Amount  float64 `json:"userId"`
	WalletID  string `json:"token"`
}

type DebitWallet struct {
	URL  string `json:"host"`
	Amount  float64 `json:"userId"`
	WalletID  string `json:"token"`
}

func (e *UserCreatedEvent) EventName() string {
	return "user.created"
}
func (e *PasswordReset) EventName() string {
	return "user.reset_password"
}

func (e *OTPCreated) EventName() string {
	return "otp.created"
}

func (e *WelcomeUserEvent) EventName() string {
	return "user.welcome"
}

func (e *CreateDefWallet) EventName() string {
	return "user.defaultwallet"
}

func (e *CreditWallet) EventName() string {
	return "user.creditwallet"
}
func (e *DebitWallet) EventName() string {
	return "user.debitwallet"
}