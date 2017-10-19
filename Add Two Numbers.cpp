// Add Two Numbers.cpp : �������̨Ӧ�ó������ڵ㡣
//

#include "stdafx.h"
#include <iostream>
#include <vector>
#include <string>
#include <algorithm>
#include <sstream>
using namespace std;

struct ListNode
{
	int val;
	ListNode *next;
	ListNode(int x) : val(x), next(NULL) {}
};

class Solution
{
  public:
	ListNode *addTwoNumbers(ListNode *l1, ListNode *l2)
	{
		ListNode *_l1 = l1;
		ListNode *_l2 = l2;
		int flag = 0;
		while (_l1 != NULL || _l2 != NULL)
		{
			if (_l1 == NULL)
			{
				_l2->val = _l2->val + flag;
				flag = 0;
				if (_l2->val >= 10)
				{
					_l2->val = _l2->val % 10;
					flag = 1;
					if (_l2->next == NULL)
					{
						_l2->next = new ListNode(1);
						return l2;
					}
				}
				if (_l2->next == NULL)
					return l2;
			}
			else if (_l2 == NULL)
			{
				_l1->val = _l1->val + flag;
				flag = 0;
				if (_l1->val >= 10)
				{
					_l1->val = _l1->val % 10;
					flag = 1;
					if (_l1->next == NULL)
					{
						_l1->next = new ListNode(1);
						return l1;
					}
				}
				if (_l1->next == NULL)
					return l1;
			}
			else
			{
				_l1->val = _l2->val = _l1->val + _l2->val + flag;
				flag = 0;
				if (_l1->val >= 10)
				{
					_l1->val = _l2->val = _l1->val % 10;
					flag = 1;
					if (_l1->next == NULL && _l2->next == NULL)
					{
						_l1->next = new ListNode(1);
						return l1;
					}
				}
				if (_l1->next == NULL && _l2->next == NULL)
					return l1;
			}
			if (_l1 != NULL)
				_l1 = _l1->next;
			if (_l2 != NULL)
				_l2 = _l2->next;
		}
	}
};

void trimLeftTrailingSpaces(string &input)
{
	input.erase(input.begin(), find_if(input.begin(), input.end(), [](int ch) {
					return !isspace(ch);
				}));
}

void trimRightTrailingSpaces(string &input)
{
	input.erase(find_if(input.rbegin(), input.rend(), [](int ch) {
					return !isspace(ch);
				}).base(),
				input.end());
}

vector<int> stringToIntegerVector(string input)
{
	vector<int> output;
	trimLeftTrailingSpaces(input);
	trimRightTrailingSpaces(input);
	input = input.substr(1, input.length() - 2);
	stringstream ss;
	ss.str(input);
	string item;
	char delim = ',';
	while (getline(ss, item, delim))
	{
		output.push_back(stoi(item));
	}
	return output;
}

ListNode *stringToListNode(string input)
{
	// Generate list from the input
	vector<int> list = stringToIntegerVector(input);

	// Now convert that list into linked list
	ListNode *dummyRoot = new ListNode(0);
	ListNode *ptr = dummyRoot;
	for (int item : list)
	{
		ptr->next = new ListNode(item);
		ptr = ptr->next;
	}
	ptr = dummyRoot->next;
	delete dummyRoot;
	return ptr;
}

string listNodeToString(ListNode *node)
{
	string result;
	while (node)
	{
		result += to_string(node->val) + ", ";
		node = node->next;
	}
	return result.substr(0, result.length() - 2);
}

int main()
{
	string line;
	while (getline(cin, line))
	{
		ListNode *l1 = stringToListNode(line);
		getline(cin, line);
		ListNode *l2 = stringToListNode(line);

		ListNode *ret = Solution().addTwoNumbers(l1, l2);

		string out = listNodeToString(ret);
		cout << out << endl;
	}
	return 0;
}
