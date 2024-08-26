#include <stdio.h>
#include <stdlib.h>
#include <string.h>


struct node {
    int data;
    struct node* next;
};

struct node* head = NULL;


int initial_data[] = { 1, 2, 3, 4, 5 };



void print_list() {
    if (head == NULL) {
        printf("List is empty.\n");
        return;
    }

    struct node* current = head;

    while (current != NULL) {
        printf("%d ", current->data);
        current = current->next;
    }

    printf("\n\n");
}


void add_node_start(int data) {
    struct node* n = (struct node*)malloc(sizeof(struct node));

    if (n == NULL) exit(EXIT_FAILURE);

    n->data = data;
    n->next = head;

    head = n;
}


void add_node_end(int data) {
    struct node* n = (struct node*)malloc(sizeof(struct node));

    if (n == NULL) exit(EXIT_FAILURE);

    n->data = data;
    n->next = NULL;

    if (head == NULL) {
        head = n;
        return;
    }

    struct node* current = head;

    while (current->next != NULL) {
        current = current->next;
    }

    current->next = n;
}


void remove_node_start() {
    if (head == NULL) {
        printf("List is empty.\n");
        return;
    }

    struct node* next = head->next;
    
    free(head);

    head = next;
}


void remove_node_end() {
    // Empty list
    if (head == NULL) {
        printf("List is empty.\n");
        return;
    }

    // Last node
    if (head->next == NULL) {
        head = NULL;
        free(head);
        return;
    }

    struct node* current = head;

    while (current->next->next != NULL) {
        current = current->next;
    }

    current->next = NULL;
    free(current->next);
}


void setup() {
    int len = sizeof(initial_data) / sizeof(initial_data[0]);

    for (int i = 0; i < len; i++) {
        add_node_end(initial_data[i]);
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
        add_node_start(data);
    }
    else if (strcmp(command, "adde") == 0) {
        printf("Data:    ");
        scanf("%d", &data);
        add_node_end(data);
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

