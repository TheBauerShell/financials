type Project struct {
    ID          int     `json:"id" gorm:"primaryKey"`
    Name        string  `json:"name" gorm:"not null"`
    Description string  `json:"description"`
    Owner       string  `json:"owner" gorm:"not null"`
    Status      string  `json:"status" gorm:"not null"` // planned, active, archived
    Department  string  `json:"department" gorm:"not null"`
    Components  []Component `json:"components" gorm:"foreignKey:ProjectID"`
}

type Component struct {
    ID          int     `json:"id" gorm:"primaryKey"`
    ProjectID   int     `json:"project_id" gorm:"not null"`
    Type        string  `json:"type" gorm:"not null"` // e.g., IaaS, PaaS, etc.
    Details     string  `json:"details"` // Additional details about the component
    Cost        float64 `json:"cost" gorm:"not null"`
}