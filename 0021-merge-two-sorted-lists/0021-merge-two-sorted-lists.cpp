/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        
        if (list1 == nullptr){         // base case: if list 1 is empty, return list 2
            return list2;              
        }
        if (list2 == nullptr){         // base case: if list 2 is empty, return list 1
            return list1;
        }

        if (list1->val <= list2->val){  // choose list1 node if its value is smaller or equal, else choose list2 node
            list1->next = mergeTwoLists(list1->next, list2); // call recursively
            return list1;
        } else {
            list2->next = mergeTwoLists(list1, list2->next); 
            return list2;   
        }
    }
};