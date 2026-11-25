class Solution {
public:
ListNode* swapPairs(ListNode* head) {
ListNode* dummy = new ListNode(0);
dummy->next = head;
ListNode* curr = dummy;

while (curr->next != nullptr && curr->next->next != nullptr) {
ListNode* first = curr->next;
ListNode* second = curr->next->next;

curr->next = second;
first->next = second->next;
second->next = first;

curr = first;
}

ListNode* newHead = dummy->next;
delete dummy;
return newHead;
}
};
