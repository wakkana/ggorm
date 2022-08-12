How to use
==
### Declare your model
```Golang
type XXX struct {   
	Id   *int64  `gorm:"primary_key,column:id"`
	Name *string `gorm:"column:name"`
}
/* need by gorm*/
func (XXX) TableName() string {
	return "XXX"
}
```
 all attrs should be ptr 
 ### Declare your search pattern
   You can declare as much as you want, since the 
   [search pattern : model] = [N : 1]
```Golang
type XXXQuery struct {
	Id   []any
	Name []any 
	// implement interface 
	// ggorm.Conditions[XXX] 
	
	// Find() XXX 
	// Kvs() map[string][]any 
	// MustFill() error

}

/*
    what model this pattern belongs to.
    never use, will generate an error if not exist
    when you write a new search pattern.
*/
func (*XXXQuery) Find() XXX {
	return XXX{}
}

// looks like reflection but not
func (xq *XXXQuery) Kvs() map[string][]any {
	src := make(map[string][]any, 0)
	src["id"] = xq.Id
	src["name"] = xq.Name
	return src
}

// check must-fill attrs
func (xq *XXXQuery) MustFill() error {
	if xq.Id == nil {
		return errors.New("id is nil")
	}
	return nil
}
```
### Add CommonConfig
```Go
type CommonConfig struct {
	Orders map[string]string
	Offset *int
	Limit  *int
}
```
### Then Use
```Go
db := ...// gorm db
repository := ggorm.NewGormRepository[XXX](db)
pattern := &XXXQuery{
	Id:   []any{">", 2},
	Name: []any{">", "0"},
}

limit := 10
res, err := repository.Query(ctx, pattern, 
	&ggorm.CommonConfig{
	        Orders: map[string]string{
		        "id": "desc",
	        }, 
	        Limit: &limit,
	}
)
```

### Check
```SQL
SELECT * FROM `XXX` WHERE id > 2 AND name > '0' ORDER BY id desc LIMIT 10
```