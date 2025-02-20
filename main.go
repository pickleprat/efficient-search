package main 

import (
	"math/rand" 
	"time" 
	"log" 
	"os" 
	"fmt" 
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
    root := htree.Data[hashIndex]
    newNode := &TreeNode{
        Value: number,
        Left:  nil,
        Right: nil,
        Index: index,
    }

    if root == nil {
        htree.Data[hashIndex] = newNode
        return
    }
    
    currNode := root
    for currNode != nil {
        if newNode.Value < currNode.Value {
            if currNode.Left == nil {
                currNode.Left = newNode
                return
            } else {
                currNode = currNode.Left
            }
        } else {
            if currNode.Right == nil {
                currNode.Right = newNode
                return
            } else {
                currNode = currNode.Right
            }
        }
    }
}

func (htree *HashTree) Hash(number int) int {
	return number % htree.Size 
} 

func (htree *HashTree) Search(number int) int {
    hashIndex := htree.Hash(number)
    rootNode := htree.Data[hashIndex]
    if rootNode == nil {
        return -1
    }
    
    currNode := rootNode
    for currNode != nil {
        if number == currNode.Value {
            return currNode.Index
        } else if number < currNode.Value {
            currNode = currNode.Left
        } else {
            currNode = currNode.Right
        }
    }
    
    return -1
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
	fresult, err := os.Create("algorithm-race.csv") 
	if err != nil {
		panic(err) 
	} 

	defer func() {
		if err := fresult.Close() ; err != nil {
			panic(err) 
		} 
	}()  

	_, err = fresult.Write([] byte("SearchArraySize,SearchNumber,HtreeResultIndex,LSResultIndex,HtreeTime,LSTime")); 
	if err != nil {
		panic(err); 
	} 

	for _, dbLen := range [] int {1000, 10000, 100000, 1000000} {
		db := NewDB(dbLen); 
		db.Fill(); 

		ls := NewLinearSearcher(&db); 
		htree := NewHashTree(dbLen / 10); 
		htree.InsertDB(db); 


		for i := 0; i < 100 ; i++ {
			number := rand.Intn(dbLen); 

			startTime := time.Now(); 
			htreeIndex := htree.Search(number); 
			elapsedHtree := time.Since(startTime); 

			startTime = time.Now(); 
			lsIndex := ls.Search(number); 
			elapsedLinearSearch := time.Since(startTime) 

			_, err = fresult.Write([] byte(
				fmt.Sprintf("%d,%d,%d,%d,%+v,%+v\n", 
					dbLen, number, htreeIndex, lsIndex, int64(elapsedHtree), int64(elapsedLinearSearch),  
				), 
			)); 

			if err != nil {
				log.Printf("Error occured at size: %d number %d", dbLen, number);  
			} else {
				log.Printf("HT:		%+v	LS: 	%+v\n", elapsedHtree, elapsedLinearSearch); 	
			} 
		} 
	} 
} 
