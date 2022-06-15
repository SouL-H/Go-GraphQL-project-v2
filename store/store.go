package store

import (
	"context"
	"graphQLv2/graph/model"
	"net/http"
)


type Store struct{
	Todos []*model.Todo
}
func (s *Store)AddTodo(t *model.NewTodo)error{
	s.Todos=append(s.Todos,&model.Todo{
		ID:"1",
		Text:t.Text,
		Done:false,
		User:&model.User{
			ID:t.UserID,
			Name:"test",
		},
	})
	return nil

}
type StoreKeyType string
var StoreKey StoreKeyType="STORE"

func NewStore() *Store{
	todo:=make([]*model.Todo,0)
	return &Store{
		Todos:todo,
	}
}

//HTTP ara katman

func WithStore(store *Store,next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		reqWithStore:=r.WithContext(context.WithValue(r.Context(),StoreKey,store)) 
		next.ServeHTTP(w,reqWithStore)
	})
}
func GetStoreFromContext(ctx context.Context) *Store{
	store,ok:= ctx.Value(StoreKey).(*Store)
	if !ok{ 
		panic("Error")
	}
	return store
}