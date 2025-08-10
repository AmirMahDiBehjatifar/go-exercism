package twofer

// the Function checks out the name is passed or not and returns string in each way
func ShareWith(name string) string {
    if name == ""{
        return "One for you, one for me."
    }else{
         return "One for " + name +", one for me."
    }
}
