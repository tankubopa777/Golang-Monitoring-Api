package userModel

type User struct {
    ID             uint   `json:"id" gorm:"primaryKey"`
    Name           string `json:"name"`
    Email          string `json:"email" gorm:"unique"`
    Password       string `json:"-"`
    IsEmailVerified bool   `json:"is_email_verified" gorm:"default:false"`
    EmailVerificationToken string `json:"-"`
}
