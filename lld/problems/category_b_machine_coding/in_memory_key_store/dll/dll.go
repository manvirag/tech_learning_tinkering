package dll

type DoublyLinkedList struct {
	Key      string
	Value    interface{}
	Previous *DoublyLinkedList
	Next     *DoublyLinkedList
}
