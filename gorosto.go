package gorosto

import
(
	"sync"
	"time"
	"fmt"
)

type gorout func()

//Basic elements for store
type Items struct {
	elems [] interface{}
}

type Storage struct {
	title string
	//limit to number of keys
	limit uint
	//Actual number of keys
	current uint
	storage map[string]Items
	mutex sync.Mutex
}


func GoroutineStorage(limit uint) (*Storage){
	store := new(Storage)
	store.title = "first"
	if(limit <= 0){
		store.limit = 0
	} else{
		store.limit = limit
	}
	store.storage =  make(map[string]Items)
	return store
}

/*
	Return target goroutine by key
	In the case, if one key contain several goroutines, return list of it
*/
func (st*Storage) Get(name string)(interface{}, bool){
	st.mutex.Lock()
	defer st.mutex.Unlock()

	result, ok := st.storage[name]
	if(ok){
		elems := result.elems
		if len(elems) == 1{
			return elems[0], ok
		}
		return result.elems, ok
	}
	return " ", false
}

func(st*Storage) Set(name string, value gorout){
	st.mutex.Lock()
	defer st.mutex.Unlock()
	if st.current == st.limit{
		return
	}
	ze := make([]interface{},1)
	ze[0] = value
	newelem := Items {ze }
	st.current += 1
	fmt.Println(time.Now(), "was append new item", name)
	st.storage[name] = newelem

}

func (st*Storage) SetItems(name string, value [] gorout){
	st.mutex.Lock()
	defer st.mutex.Unlock()
	fmt.Println(time.Now(), "was append new item", name)
	newdata := make([]interface{}, len(value))
	for i := range value{
		newdata[i] = value[i]
	}
	newelem := Items {newdata }
	st.storage[name] = newelem
}


