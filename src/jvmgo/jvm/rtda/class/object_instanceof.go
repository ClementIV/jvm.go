package class

// jvms8-6.5.instanceof
func (self *Obj) IsInstanceOf(class *Class) (bool) {
    s := self.class
    t := class

    if s == t {
        return true
    }

    if !s.IsArray() {
        if !s.IsInterface() {
            if !t.IsInterface() {
                return s.isSubClassOf(t)
            } else {
                return s.isImplements(t)
            }
        } else {
            if !t.IsInterface() {
                return t.isObject()
            } else {
                return t.isSuperInterfaceOf(t)
            }
        }
    } else { // s is array
        if !t.IsArray() {
            if !t.IsInterface() {
                return t.isObject()
            } else {
                return t.IsCloneable() || t.IsSerializable()
            }
        } else { // t is array
            // todo
        }
    }

    return false
}