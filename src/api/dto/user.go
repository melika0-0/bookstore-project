package dto

type Loginreq struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}





// check to map concept 
func (r *Loginreq) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":     r.Name,
		"password": r.Password,
	}
}

type Loginres struct {
	Token string `json:"token"`
}

func (r *Loginres) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"token": r.Token,
	}
}

