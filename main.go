package main 

import (
	"math/rand" 
	"time" 
	"log" 
) 


type Database [] int; 

type Searcher interface {
	Search( num int ) int 
} 

func (db *Database) Fill() {
	for i, _ := range *db {
		j := rand.Intn(i + 1); 
		(*db)[j] = i; 
		(*db)[i] = j; 
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
		htree.Data[hashIndex] = root; 
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
			currNode = currNode.Left; 
		} else {
			currNode = currNode.Right;  
		} 
	} 

	if currNode == nil {
		return -1; 
	} else {
		return currNode.Index; 
	} 
} 

type LinearSearch struct {
	Db *Database 
} 

func (ls *LinearSearch) Search(number int) int {
	for idx, num := range (*ls.Db) {
		if num == number {
			return idx
		} 
	} 

	return -1; 
} 

func NewLinearSearcher( db * Database)  * LinearSearch {
	return &LinearSearch {
		Db : db, 
	} 
} 


func main() {
	for _, dbLen := range [] int {1000, 10000, 100000, 1000000} {
		db := NewDB(dbLen); 
		db.Fill(); 

		ls := NewLinearSearcher(&db); 
		htree := NewHashTree(dbLen / 10); 

		for i := 0; i < 100 ; i++ {
			startTime := time.Now(); 
			htree.Search(rand.Intn(dbLen)); 
			elapsedHtree := time.Since(startTime); 

			startTime = time.Now(); 
			ls.Search(rand.Intn(dbLen)); 
			elapsedLinearSearch := time.Since(startTime) 

			log.Printf("LS: %+v 	HT: %+v", elapsedLinearSearch, elapsedHtree); 
		} 
	} 
} 
