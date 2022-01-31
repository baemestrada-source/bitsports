package test

import ( 
	"testing"
	"context"
    "math/rand"
    "strconv"
	"time"
    "github.com/baemestrada-source/bitsports/db"
	_ "github.com/lib/pq"
)

func Test(t *testing.T) {
    ctx := context.Background()
   
    rand.Seed(time.Now().UnixNano())
    min := 1
    max := 100

    c := db.PosgresCN
    defer c.Close()

    u, err := c.Categories.Create().
        SetName("test"+strconv.Itoa(rand.Intn(max - min + 1) + min)).
        Save(ctx)

    if err != nil {
        t.Fatal(err)
    }

    p, err := c.Products.Create().
        SetCategorieID(u.ID).
        SetName("product test").
        SetInfo("test ").
        SetPrice(10).
        Save(ctx)

    if err != nil {
        t.Fatal(err)
    }
   
    ProductsWithOwnerId := c.Products.GetX(ctx, p.ID)

    t.Log(ProductsWithOwnerId.ID)
}