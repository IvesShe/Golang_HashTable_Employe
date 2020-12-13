package main

import (
	"fmt"
	"os"
)

// 定義Employe
type Employe struct {
	ID   int
	Name string
	Next *Employe
}

func (el *Employe) Show() {
	fmt.Printf("鏈表%d 找到該雇員 %s\n", el.ID%7, el.Name)
}

// 定義EmployeLink
// 這裡的EmployeLink不帶表頭，即第一個節點就存放雇員
type EmployeLink struct {
	Head *Employe
}

// 添加雇員的方法，保證添加時，編號從小到大
func (el *EmployeLink) Insert(emp *Employe) {
	// 這是輔助指針
	cur := el.Head

	// 這是輔助指針，pre在cur的前面
	var pre *Employe = nil

	// 如果當前的EmpLink是空鏈表
	if cur == nil {
		el.Head = emp
		return
	}

	// 如果不是空鏈表，給emp找到對應的位置並插入
	// 思路是讓cur與emp比較，然後讓pre保持在cur前面
	for {
		if cur != nil {
			// 比較
			if cur.ID > emp.ID {
				//找到位置
				break
			}
			// 保持同步
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}

	// 退出時，將emp添加到鏈表最後
	pre.Next = emp
	emp.Next = cur
}

// 顯示鏈表的信息
func (el *EmployeLink) ShowLink(no int) {
	if el.Head == nil {
		fmt.Printf("鏈表%d: 為空\n", no)
		return
	}

	// 遍歷當前鏈表，並顯示數據
	// 輔助的指針
	fmt.Printf("鏈表%d: ", no)
	cur := el.Head
	for {
		if cur != nil {
			//fmt.Printf("鏈表%d 雇員ID=%d 名字=%s -> ", no, cur.ID, cur.Name)
			fmt.Printf("雇員ID=%d,名字=%s -> ", cur.ID, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}

	// 換行
	fmt.Println()
}

// 根據id查找對應的雇員
func (el *EmployeLink) FindByID(id int) *Employe {
	cur := el.Head
	for {
		if cur != nil && cur.ID == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

// 定義HashTable，含有三個鏈表數組
type HashTable struct {
	LinkArr [7]EmployeLink
}

// 給HashTable 編寫 Insert 雇員的方法
func (ht *HashTable) Insert(emp *Employe) {
	// 使用散列方法，決定要將該雇員添加到哪個鏈表
	linkNo := ht.HashFun(emp.ID)
	// 使用對應的鏈表添加
	ht.LinkArr[linkNo].Insert(emp)
}

// 顯示HashTable的所有雇員
func (ht *HashTable) ShowAll() {
	fmt.Println()
	for i := 0; i < len(ht.LinkArr); i++ {
		ht.LinkArr[i].ShowLink(i)
	}
}

// 編寫一個散列方法
func (ht *HashTable) HashFun(id int) int {
	// 獲取對應鏈表的索引值
	return id % 7
}

// 以ID查找
func (ht *HashTable) FindByID(id int) *Employe {
	// 使用函列方法，決定在哪個鏈表找雇員
	linkNo := ht.HashFun(id)
	return ht.LinkArr[linkNo].FindByID(id)
}

func main() {

	key := ""
	id := 0
	name := ""
	var hashTable HashTable
	for {
		fmt.Println("===================雇員系統功能===================")
		fmt.Println("input 表示添加雇員")
		fmt.Println("show 表示顯示雇員")
		fmt.Println("find 表示查找雇員")
		fmt.Println("exit 表示退出雇員")
		fmt.Println("請輸入您的選擇")
		fmt.Scanln(&key)

		switch key {
		case "input":
			fmt.Println("請輸入雇員ID")
			fmt.Scanln(&id)
			fmt.Println("請輸入雇員Name")
			fmt.Scanln(&name)
			emp := &Employe{
				ID:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.ShowAll()
		case "find":
			fmt.Println("請輸入ID:")
			fmt.Scanln(&id)
			emp := hashTable.FindByID(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇員不存在\n", id)
			} else {
				emp.Show()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("輸入錯誤，請重新輸入")
		}

	}
}
