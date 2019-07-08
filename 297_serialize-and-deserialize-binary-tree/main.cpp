#include <iostream>
#include <string>
#include <stack>
#include <vector>
using namespace std;
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Codec
{
public:
    // Encodes a tree to a single string.
    string serialize(TreeNode *root)
    {
        stack<TreeNode *> s;
        s.push(root);
        string res = "";
        while (!s.empty())
        {
            auto el = s.top();
            s.pop();
            if (el == NULL)
            {
                res += "#,";
            }
            else
            {
                res += to_string(el->val) + ",";
                s.push(el->right);
                s.push(el->left);
            }
        }
        return res;
    }

    // Decodes your encoded data to tree.
    TreeNode *deserialize(string data)
    {
        TreeNode *root = nullptr;
        stack<TreeNode **> s;
        s.push(&root);
        string::iterator first = data.begin();
        while (first != data.end())
        {
            auto el = s.top();
            s.pop();
            if (*first == '#')
            {
                advance(first, 2);
            }
            else
            {
                string::iterator last = find(first, data.end(), ',');
                auto node = new TreeNode(stoi(string(first, last)));
                *el = node;
                s.push(&(node->right));
                s.push(&(node->left));
                first = next(last);
            }
        }
        return root;
    }
};

// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.deserialize(codec.serialize(root));