package main

// 值类型变量可以直接赋值或修改变量对应的值
// 引用类型变量,不仅要声明变量,更重要的是需要对该引用类型分配内存空间
// new() 该方法要求传入一个类型,它会申请该类型大小的内存空间,并会初始化值,返回指向该内存空间的地址(指针)
// make() 它只是用于slice,map,channel的内存创建,返回的类型就是类型本身,而不是他们的指针
