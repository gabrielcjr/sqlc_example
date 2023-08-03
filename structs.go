package main

type Category struct {
    ID   int32
    Name string
    Produtos []Produto
}

type Produto struct {
    ID         int32
    Name       string
    Price      float64
    CategoryID int32
}