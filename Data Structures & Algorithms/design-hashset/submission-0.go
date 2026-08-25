type MyHashSet struct {
	arr []int
}

func Constructor() MyHashSet {
    return MyHashSet{arr: []int{}}
}

func (this *MyHashSet) Add(key int) {
    if !this.Contains(key){
		this.arr = append(this.arr, key)
	}
}

func (this *MyHashSet) Remove(key int) {
    for i,v:= range this.arr{
		if v==key{
			this.arr = append(this.arr[:i],this.arr[i+1:]...)
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
    for _,v:=range this.arr{
		if v==key{
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 