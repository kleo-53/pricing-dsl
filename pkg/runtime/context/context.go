package context

type Context struct {
	User    UserContext
	Product ProductContext
	Cart    CartContext
}

type UserContext struct {
	ID   string
	Type string
	Age  float64
}

type ProductContext struct {
	ID       string
	Category string
	Price    float64
}

type CartContext struct {
	Total      float64
	ItemsCount float64
}
