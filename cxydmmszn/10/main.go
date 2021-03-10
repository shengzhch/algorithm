package main

import (
	"fmt"
	"algorithm/basic"
	"strings"
)

type Pet interface {
	PetName() string
}

type Cat struct {
	name string
}

func (c *Cat) PetName() string {
	return "cat_" + c.name
}

type Dog struct {
	name string
}

func (c *Dog) PetName() string {
	return "dog_" + c.name
}

//构造新的结构
type PetEnqueueEle struct {
	pet   Pet
	count int64
}

func (c *PetEnqueueEle) GetPet() Pet {
	return c.pet
}
func (c *PetEnqueueEle) Count() int64 {
	return c.count
}
func (c *PetEnqueueEle) GetPetName() string {
	return c.pet.PetName()
}

type PetEnqueue struct {
	dogQ *basic.Queue
	catQ *basic.Queue
	//编号
	count int64
}

func (pq *PetEnqueue) Add(pet Pet) {
	if strings.HasPrefix(pet.PetName(), "dog") {
		pq.dogQ.EnQueue(&PetEnqueueEle{pet, pq.count})
	} else if strings.HasPrefix(pet.PetName(), "cat") {
		pq.catQ.EnQueue(&PetEnqueueEle{pet, pq.count})
	}
	pq.count++
}

func (pq *PetEnqueue) Pull() Pet {
	if pq.dogQ.Size() > 0 && pq.catQ.Size() > 0 {
		if pq.dogQ.Head().(*PetEnqueueEle).Count() < pq.catQ.Head().(*PetEnqueueEle).Count() {
			return pq.dogQ.DeQueueWithValue().(*PetEnqueueEle).pet
		} else {
			return pq.catQ.DeQueueWithValue().(*PetEnqueueEle).pet
		}
	} else if pq.dogQ.Size() > 0 {
		return pq.dogQ.DeQueueWithValue().(*PetEnqueueEle).pet
	} else if pq.catQ.Size() > 0 {
		return pq.catQ.DeQueueWithValue().(*PetEnqueueEle).pet
	} else {
		//队列为空
		return nil
	}
}

func (pq *PetEnqueue) PullDog() Pet {
	if pq.dogQ.Size() > 0 {
		return pq.dogQ.DeQueueWithValue().(*PetEnqueueEle).pet
	}
	return nil
}
func (pq *PetEnqueue) PullCat() Pet {
	if pq.catQ.Size() > 0 {
		return pq.catQ.DeQueueWithValue().(*PetEnqueueEle).pet
	}
	return nil
}

func main() {
	pq := &PetEnqueue{&basic.Queue{}, &basic.Queue{}, 0}

	pq.Add(&Cat{"1"})
	pq.Add(&Dog{"1"})
	pq.Add(&Cat{"2"})
	pq.Add(&Dog{"2"})

	c := pq.Pull()
	fmt.Println(c.PetName())

	c = pq.Pull()
	fmt.Println(c.PetName())

	c = pq.PullDog()
	fmt.Println(c.PetName())

	c = pq.PullCat()
	fmt.Println(c.PetName())
}
