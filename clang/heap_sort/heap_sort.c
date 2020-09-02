#include <stdio.h>

void swap(int tree[], int i, int j) {
    int temp = tree[i];
    tree[i] = tree[j];
    tree[j] =  temp;
}

// heapify the node of tree
// len is the length of tree
// node is the node of tree
void heapify(int tree[], int len, int node) {
    if (node >= len) {
        return ;
    }

    int child_1 = node * 2 + 1;
    int child_2 = node * 2 + 2;
    int max = node;

    if (child_1 < len && tree[child_1] > tree[max]) {
        max = child_1;
    }
    if (child_2 < len && tree[child_2] > tree[max]) {
        max = child_2;
    }
    if (max != node) {
        swap(tree, max, node);
        heapify(tree, len, max);
    }
}

void build_heap(int tree[], int len) {
    int last_node = len - 1;
    int parent = (last_node - 1) / 2;
    int i;
    for (i = parent;i >= 0; i--) {
        heapify(tree, len, i);
    }
}

void heap_sort(int tree[], int len) {
    build_heap(tree, len);
    int i;
    for (i = len - 1; i >= 0; i--) {
        swap(tree, i, 0);
        heapify(tree, i, 0);
    }
}

int main() {
    int tree[] = {2, 5, 3, 1, 10, 4};
    int len = 6;
    heap_sort(tree, len);
    int i;
    for (i = 0; i < len; i++) {
        printf("%d\n", tree[i]);
    }
    return 0;
}
