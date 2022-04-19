package domain

type User struct {
	Models
	FirstName string `json:"first_name" gorm:"not null" binding:"required" form:"first_name"`
	LastName  string `json:"last_name" gorm:"not null" binding:"required" form:"last_name"`
	Phone1    string `json:"phone" gorm:"not null" binding:"required" form:"phone1"`
	Phone2    string `json:"phone_2,omitempty" form:"phone2"`
	Email     string `json:"email" gorm:"unique" binding:"required,email" form:"email"`
	Address   string `json:"address,omitempty" gorm:"not null" form:"address"`
	//HashedPassword  string `json:"-,omitempty" gorm:"not null"`
	Password        string `json:"-" gorm:"-" binding:"required" form:"password"`
	ConfirmPassword string `json:"confirm_password" gorm:"-" form:"confirm_password"`
	BookmarkedItems []Item `gorm:"many2many:bookmarked_items" json:"bookmarked_items,omitempty"`
	Image           string `json:"image,omitempty"`
	RoleID          string `json:"role_id"`
	Role            Role
	IsActive        bool   `json:"is_active" gorm:"default:false"`
	Token           string `json:"-,omitempty"`
}

type UpdateUser struct {
	FirstName string `json:"first_name" binding:"required" form:"first_name"`
	LastName  string `json:"last_name" binding:"required" form:"last_name"`
	Phone1    string `json:"phone" binding:"required" form:"phone1"`
	Phone2    string `json:"phone_2" form:"phone2"`
	Email     string `json:"email" binding:"required,email" form:"email"`
	Address   string `json:"address"  form:"address"`
}
