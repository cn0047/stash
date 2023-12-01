package main

type MinHeap struct {
  items []int
}

// HeapifyUp performs heapify up.
// Plan:
// get last element, in loop get parent,
// if parent has greater value - swap, continue loop.
func (m *MinHeap) HeapifyUp() {
  index := len(m.items) - 1
  for m.HasParent(index) && m.GetParent(index) > m.items[index] {
    parentIndex := m.GetParentIndex(index)
    m.Swap(parentIndex, index)
    index = parentIndex
  }
}

// HeapifyDown performs heapify down.
// Plan:
// get root element, in loop get smaller child,
// if child has smaller value - swap, continue loop.
func (m *MinHeap) HeapifyDown() {
  index := 0
  for m.HasLeftChild(index) {
    smallerChildIndex := m.GetLeftChildIndex(index)
    if m.HasRightChild(index) && m.RightChild(index) < m.LeftChild(index) {
      smallerChildIndex = m.GetRightChildIndex(index)
    }
    if m.items[index] > m.items[smallerChildIndex] {
      m.Swap(index, smallerChildIndex)
    } else {
      break
    }
    index = smallerChildIndex
  }
}

func (m *MinHeap) Add(item int) {
  m.items = append(m.items, item)
  m.HeapifyUp()
}

func (m *MinHeap) DeleteRoot() int {
  item := m.items[0]
  m.items[0] = m.items[len(m.items)-1] // last element
  m.items = m.items[:len(m.items)-1] // resize
  m.HeapifyDown()
  return item
}

func (m *MinHeap) Delete(itemToDelete int) int {
  index := 0
  for i := 0; i < len(m.items); i++ {
    if m.items[i] == itemToDelete {
      index = i
    }
  }

  m.items[index] = m.items[len(m.items)-1] // last element
  m.items = m.items[:len(m.items)-1] // resize
  m.HeapifyDown()
  return itemToDelete
}

func (m *MinHeap) GetLeftChildIndex(parentIndex int) int {
  return 2*parentIndex + 1
}

func (m *MinHeap) GetRightChildIndex(parentIndex int) int {
  return 2*parentIndex + 2
}

func (m *MinHeap) HasLeftChild(index int) bool {
  return m.GetLeftChildIndex(index) < len(m.items)
}

func (m *MinHeap) HasRightChild(index int) bool {
  return m.GetRightChildIndex(index) < len(m.items)
}

func (m *MinHeap) LeftChild(index int) int {
  return m.items[m.GetLeftChildIndex(index)]
}

func (m *MinHeap) RightChild(index int) int {
  return m.items[m.GetRightChildIndex(index)]
}

func (m *MinHeap) GetParentIndex(childIndex int) int {
  return (childIndex - 1) / 2
}

func (m *MinHeap) HasParent(index int) bool {
  return m.GetParentIndex(index) >= 0
}

func (m *MinHeap) GetParent(index int) int {
  return m.items[m.GetParentIndex(index)]
}

func (m *MinHeap) Swap(index1 int, index2 int) {
  m.items[index1], m.items[index2] = m.items[index2], m.items[index1]
}

func (m *MinHeap) Peak() int {
  return m.items[0]
}

func main() {
  case1()
  case2()
  case3()
  case4()
}

func case4() {
  h := MinHeap{}
  h.Add(10)
  h.Add(9)
  h.Add(3)
  h.Delete(9)
  h.Add(11)
  h.Delete(3)
  h.Add(1)
  h.DeleteRoot()
  h.Add(2)
  h.Add(4)
  h.Delete(2)
  h.Add(5)
  h.Delete(4)
  println(h.Peak() == 5)
}

func case3() {
  h := MinHeap{}
  h.Add(10)
  h.Add(9)
  h.Add(3)
  h.Delete(9)
  h.Delete(3)
  h.Add(11)
  h.Add(1)
  println(h.Peak() == 1)
}

func case2() {
  h := MinHeap{}
  h.Add(10)
  h.Add(9)
  h.Add(3)
  h.Delete(9)
  h.Delete(3)
  println(h.Peak() == 10)
}

func case1() {
  h := MinHeap{}
  h.Add(5)
  h.Add(3)
  h.Add(6)
  h.Add(9)
  h.Add(8)
  h.Add(1)
  h.DeleteRoot()
  h.Delete(3)
  println(h.Peak() == 5)
}
