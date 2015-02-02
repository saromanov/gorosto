package gorosto


func main() {
	testfunc:= func() {
		go func(){
			result := 0
			for i := 0; i < 1000; i++ {
				result += i
			}
		}()
		fmt.Println("This thing")
	}
	store := GoroutineStorage(0)
	store.Set("first", testfunc)
	result, status := store.Get("first")
	if(status){
		result.(gorout)()
	}
}

