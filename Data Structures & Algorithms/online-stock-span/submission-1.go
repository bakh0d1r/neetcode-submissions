type StockSpanner struct {
    stock []Stock
}
type Stock struct {
    price int
    span int
}

func Constructor() StockSpanner {
    return StockSpanner{
        stock: []Stock{},
    }
}

func (this *StockSpanner) Next(price int) int {
    span:=1
    for len(this.stock) > 0 && this.stock[len(this.stock)-1].price <= price {
        topStock:=this.stock[len(this.stock)-1]
        this.stock=this.stock[:len(this.stock)-1]
        span+=topStock.span
    }
    this.stock=append(this.stock,Stock{
        price:price,
        span:span,
    })
    return span
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

