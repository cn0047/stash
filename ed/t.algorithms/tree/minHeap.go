package main

type MinHeap struct {
  items []int64
}

func (m *MinHeap) GetLeftChildIndex(parentIndex int64) int64 {
  return 2*parentIndex + 1
}

func (m *MinHeap) GetRightChildIndex(parentIndex int64) int64 {
  return 2*parentIndex + 2
}

func (m *MinHeap) HasLeftChild(index int64) bool {
  return m.GetLeftChildIndex(index) < int64(len(m.items))
}

func (m *MinHeap) HasRightChild(index int64) bool {
  return m.GetRightChildIndex(index) < int64(len(m.items))
}

func (m *MinHeap) LeftChild(index int64) int64 {
  return m.items[m.GetLeftChildIndex(index)]
}

func (m *MinHeap) RightChild(index int64) int64 {
  return m.items[m.GetRightChildIndex(index)]
}

func (m *MinHeap) GetParentIndex(childIndex int64) int64 {
  return (childIndex - 1) / 2
}

func (m *MinHeap) HasParent(index int64) bool {
  return m.GetParentIndex(index) >= 0
}

func (m *MinHeap) Parent(index int64) int64 {
  return m.items[m.GetParentIndex(index)]
}

func (m *MinHeap) Swap(index1 int64, index2 int64) {
  tmp := m.items[index1]
  m.items[index1] = m.items[index2]
  m.items[index2] = tmp
}

func (m *MinHeap) Peak() int64 {
  return m.items[0]
}

func (m *MinHeap) DeleteRoot() int64 {
  item := m.items[0]
  m.items[0] = m.items[len(m.items)-1]
  m.items = m.items[:len(m.items)-1]
  m.HeapifyDown()
  return item
}

func (m *MinHeap) Delete(itemToDelete int64) int64 {
  index := 0
  for i := 0; i < len(m.items); i++ {
    if m.items[i] == itemToDelete {
      index = i
    }
  }

  m.items[index] = m.items[len(m.items)-1]
  m.items = m.items[:len(m.items)-1]
  m.HeapifyDown()
  return itemToDelete
}

func (m *MinHeap) Add(item int64) {
  m.items = append(m.items, item)
  m.HeapifyUp()
}

func (m *MinHeap) HeapifyUp() {
  index := int64(len(m.items) - 1)
  for m.HasParent(index) && m.Parent(index) > m.items[index] {
    m.Swap(m.GetParentIndex(index), index)
    index = m.GetParentIndex(index)
  }
}

func (m *MinHeap) HeapifyDown() {
  index := int64(0)
  for m.HasLeftChild(index) {
    smallerChildIndex := m.GetLeftChildIndex(index)
    if m.HasRightChild(index) && m.RightChild(index) < m.LeftChild(index) {
      smallerChildIndex = m.GetRightChildIndex(index)
    }
    if m.items[index] < m.items[smallerChildIndex] {
      break
    } else {
      m.Swap(index, smallerChildIndex)
    }
    index = smallerChildIndex
  }
}

func main() {
  case2()
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
