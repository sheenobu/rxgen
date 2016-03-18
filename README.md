# rxgen

golang library for generating reactive code

## example

//go:generate rxgen -type uint -name RxUint .

r := NewRxUint(1)

v := r.Get() // 1

c := r.Subscribe()
	
go func() {
	val := <-c.C // 1234
	fmt.Printf("%s\n", val)
}()

r.Set(1234)	

