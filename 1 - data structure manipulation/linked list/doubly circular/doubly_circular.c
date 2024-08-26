#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>


struct node {
    int data;
    struct node* next;
    struct node* prev;
};

struct node* head = NULL;


int initial_data[] = { 1, 2, 3, 4, 5 };


void print_list() {
    if (head == NULL) {
        printf("List is empty.\n");
        return;
    }

    struct node* current = head;

    printf("%d ", current->data);

    current = current->next;

    while (current != head) {
        printf("%d ", current->data);
        current = current->next;
    }

    printf("\n\n");
}


void add_node(int data, bool to_start) {
    struct node* n = (struct node*)malloc(sizeof(struct node));

    if (n == NULL) exit(EXIT_FAILURE);

    n->data = data;

    if (head == NULL) {
        n->next = n;
        n->prev = n;
        head = n;
        return;
    }

    n->next = head;
    n->prev = head->prev;

    head->prev->next = n;
    head->prev = n;

    if (to_start) {
        head = n;
    }
}


void remove_node(bool from_start) {
    // Empty list
    if (head == NULL) {
        printf("List is empty.\n");
        return;
    }

    struct node* n = (from_start) ? head : head->prev;

    struct node* next = n->next;

    // Last node
    if (next == n) {
        head = NULL;
        free(head);
        return;
    }

    n->next->prev = n->prev;
    n->prev->next = n->next;

    free(n);

    head = next;
}

void remove_node_start() {
    remove_node(true);
}
void remove_node_end() {
    remove_node(false);
}


void setup() {
    int len = sizeof(initial_data) / sizeof(initial_data[0]);

    for (int i = 0; i < len; i++) {
        add_node(initial_data[i], false);
    }
}


void command_interface() {
    char command[4];
    int data;

    printf("Add to start:      adds\n");
    printf("Add to end:        adde\n");
    printf("Remove from start: rems\n");
    printf("Remove from end:   reme\n");
    printf("Quit:              exit\n");
    printf("\nCommand: ");

    scanf("%s", command);

    if (strcmp(command, "adds") == 0) {
        printf("Data:    ");
        scanf("%d", &data);
        add_node(data, true);
    }
    else if (strcmp(command, "adde") == 0) {
        printf("Data:    ");
        scanf("%d", &data);
        add_node(data, false);
    }
    else if (strcmp(command, "rems") == 0) remove_node_start();
    else if (strcmp(command, "reme") == 0) remove_node_end();
    else if (strcmp(command, "exit") == 0) return;
    else {
        printf("Invalid command.\n");
        return;
    }

    printf("\n\n\n");
    print_list();
    command_interface();
}


int main(int argc, char* argv[]) {
    setup();
    print_list();

    command_interface();

    free(head);
    printf("Quitting.\n");
    return 0;
}

