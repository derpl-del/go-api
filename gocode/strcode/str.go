package strcode

//RqHelloWorld struct
type RqHelloWorld struct {
	Name string `json:"name"`
}

//Response struct
type Response struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

//ProductInfo struct
type ProductInfo struct {
	ProductName   string `json:"name"`
	ProductAmount int    `json:"amount"`
	ProductPrice  int    `json:"price"`
}

//ProductList struct
type ProductList struct {
	ProductList []ProductInfo `json:"product_data"`
}

//UserInfo struct
type UserInfo struct {
	UserName string `json:"user_name"`
	Wallet   int    `json:"wallet"`
	Mail     string `json:"mail"`
}

//BuyProduct struct
type BuyProduct struct {
	UserName    string `json:"user_name"`
	ProductName string `json:"product_name"`
	Amount      int    `json:"amount"`
	Price       int    `json:"price"`
}

//SenderMail struct
type SenderMail struct {
	Sender   string `json:"sender"`
	Password string `json:"password"`
}
