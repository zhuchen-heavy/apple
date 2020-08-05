package main

import "fmt"

/**
 * <p>
 * Map集合
 * Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
 * Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
 *
 * ----------------------------
 * Map语法
 * var map_variable map[key_data_type]value_data_type： 声明变量，默认 map 是 nil
 * map_variable := make(map[key_data_type]value_data_type)：使用 make 函数
 * </p>
 * @author: zhu.chen
 * @date: 2020/8/5
 * @version: v1.0.0
 */

func main() {
	var countryCapitalMap map[string]string /*创建集合 */
	// 对Map进行切片后，容易进行相关切片Api的操作
	countryCapitalMap = make(map[string]string)
	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
	fmt.Println(countryCapitalMap)
	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	fmt.Println(capital)
	fmt.Println(ok)
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	fmt.Println("--------------------")
	mapTest()
}

func mapTest() {
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println(countryCapitalMap)
	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}
