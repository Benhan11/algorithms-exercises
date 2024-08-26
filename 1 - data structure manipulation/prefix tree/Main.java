import java.util.Scanner;

public class Main {
    private static Trie trie;

    public static void main(String[] args) {
        trie = new Trie();
        addData();

        Scanner scanner = new Scanner(System.in);
        ui(scanner);
    }

    private static void ui(Scanner scanner) {
        System.out.println("> Insert: i <string>");
        System.out.println("> Search: s <string>");
        System.out.println("> Delete: d <string>");
        System.out.println("> Words:  p");
        System.out.println("> Exit:   e");
        System.out.print("> ");

        String input = scanner.nextLine();

        switch(input.charAt(0)) {
            case 'i':
                trie.insertKey(input.substring(2));
                break;
            case 's':
                System.out.println("Exists: " + trie.keyExists(input.substring(2)));
                break;
            case 'd':
                trie.deleteKey(input.substring(2));
                break;
            case 'p':
                trie.printTrie();
                break;
            case 'e':
                System.out.println("Quitting.");
                return;
            default:
                System.out.println("Invalid command.");
                break;
        }

        System.out.println();
        ui(scanner);
    }

    private static void addData() {
        // trie.insertKey("banana");
        // trie.insertKey("grape");
        // trie.insertKey("orange");
        // trie.insertKey("melon");
        // trie.insertKey("mango");
        // trie.insertKey("cherry");
        // trie.insertKey("blueberry");
        // trie.insertKey("strawberry");
        // trie.insertKey("pineapple");
        // trie.insertKey("kiwi");
        // trie.insertKey("peach");
        // trie.insertKey("plum");
        // trie.insertKey("pear");
        // trie.insertKey("lemon");
        // trie.insertKey("lime");
        // trie.insertKey("watermelon");
        // trie.insertKey("raspberry");
        // trie.insertKey("blackberry");
        // trie.insertKey("apricot");
        // trie.insertKey("nectarine");
        // trie.insertKey("grapefruit");
        // trie.insertKey("pomegranate");
        // trie.insertKey("tangerine");
        // trie.insertKey("coconut");
        // trie.insertKey("papaya");
        // trie.insertKey("guava");
        // trie.insertKey("passionfruit");
        // trie.insertKey("dragonfruit");
        // trie.insertKey("fig");
        // trie.insertKey("date");
        // trie.insertKey("lychee");
        // trie.insertKey("persimmon");
        // trie.insertKey("cranberry");
        // trie.insertKey("gooseberry");
        // trie.insertKey("mulberry");
        // trie.insertKey("cantaloupe");
        // trie.insertKey("honeydew");
        // trie.insertKey("jackfruit");
        // trie.insertKey("durian");
        
        trie.insertKey("duriant");
        trie.insertKey("durian");
        trie.insertKey("dumbo");
        trie.insertKey("dont");
        trie.insertKey("banana");
    }
}
