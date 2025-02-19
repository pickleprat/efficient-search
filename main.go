package main 

import (
	"math/rand" 
	"fmt" 
) 


type Database [] int ; 

func (db *Database) FillWithRandom(minimum, maximum int) {
	for i := 0; i < len(*db); i++ {
		(*db)[i] = minimum + rand.Intn(maximum - minimum)  
	} 
} 


func NewDB(sz int) Database {
	return make(Database, sz);  
} 

type TreeNode struct {
	Value 	int  
	Left  	*TreeNode  
	Right 	*TreeNode 
	Index 	int 
} 

type HashTree struct {
	Data 	[] *TreeNode 
	Size 	int 
} 

func NewHashTree(size int)  *HashTree {
	return &HashTree { Size: size, Data: make([] * TreeNode, size) } 
} 

func (htree *HashTree) InsertDB(db Database) {
	for idx, number := range db {
		htree.Insert(number, idx); 
	} 
} 

func (htree *HashTree) Insert(number, index int) {
	hashIndex := htree.Hash(number) 
	root := htree.Data[hashIndex]; 
	
	newNode := &TreeNode {
		Value : number, 
		Left  : nil, 
		Right : nil, 
		Index : index,  
	} 

	if root == nil {
		root = newNode; 
		return 
	} 

	currNode := root; 
	for  currNode != nil  {
		if newNode.Value < currNode.Value {
			if currNode.Right == nil {
				currNode.Right = newNode; 
				return 
			} else {
				currNode = currNode.Right; 
			} 
		} else {
			if currNode.Left == nil {
				currNode.Left = newNode; 
				return 
			} else {
				currNode = currNode.Left; 
			} 
		} 
	} 

	htree.Data[hashIndex] = root; 
} 

func (htree *HashTree) Hash(number int) int {
	return number % htree.Size 
} 

func (htree *HashTree) Search(number int) int {
	hashIndex := htree.Hash(number); 
	rootNode := htree.Data[hashIndex]; 
	if rootNode == nil {
		return -1; 
	} 
	
	if rootNode.Value == number {
		return rootNode.Index; 
	} 

	currNode := rootNode; 
	for currNode != nil {
		if number == currNode.Value {
			return currNode.Index;  

		} else if number > currNode.Value {
			currNode = currNode.Right; 
		} else {
			currNode = currNode.Left;  
		} 
	} 

	if currNode == nil {
		return -1; 
	} else {
		return currNode.Index; 
	} 
} 


func main() {
	var size int = 100;  
	db := NewDB(size); 
	db.FillWithRandom(0, 100); 

	htree := NewHashTree(size / 10); 
	htree.InsertDB(db); 

	searchNumber := 9; 
	hashIndex := htree.Hash(searchNumber); 
	node := htree.Data[hashIndex];  
	fmt.Printf("Hash Index: %+v\n", *node); 


} 
