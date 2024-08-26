import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Trie {
    private Node root;

    public Trie() {
        root = new Node();
    }

    public void insertKey(String keyString) {
        char[] key = keyString.toCharArray();
        int charIndex;
        
        Node node = root;
        Node[] children;
        
        for (int i = 0; i < key.length; i++) {
            children = node.getChildren();
            charIndex = key[i] - 'a';
            
            // If there isn't already a node for the character,
            // make it and attach it to the tree
            if (children[charIndex] == null) {
                Node n = new Node();
                children[charIndex] = n;
            }
            
            // Go to next node
            node = children[charIndex];
        }
        
        // Indicate end at last node
        node.setIsTerminal(true);
    }

    public boolean keyExists(String keyString) {
        char[] key = keyString.toCharArray();
        int charIndex;

        Node node = root;
        Node[] children;

        for (int i = 0; i < key.length; i++) {
            children = node.getChildren();
            charIndex = key[i] - 'a';

            // If the node for the current character doesn't exist
            if (children[charIndex] == null) {
                return false;
            }

            node = children[charIndex];
        }

        return node.getIsTerminal();
    }

    public void deleteKey(String keyString) {
        char[] key = keyString.toCharArray();
        deleteKeyRecurse(root, key);
    }
    public Node deleteKeyRecurse(Node node, char[] key) {
        Node[] children = node.getChildren();

        // If we are at the bottom of the node structure
        if (key.length == 0) {
            // Signify that this last node is to be removed
            if (node.getIsTerminal()) {
                node.setIsTerminal(false);
            }

            // Check if there are any other active nodes as to not delete
            // other entries
            for (int i = 0; i < children.length; i++) {
                if (children[i] != null) {
                    return node;
                }
            }

            // Delete node if it was the only one
            return null;
        }

        // Go to next node
        children[key[0] - 'a'] = deleteKeyRecurse(children[key[0] - 'a'], Arrays.copyOfRange(key, 1, key.length));
        return node;
    }

    public void printTrie() {
        recursePrintTrie(root, new StringBuilder(), 0);
    }
    private void recursePrintTrie(Node node, StringBuilder str, int level) {
        if (node == null) return;
        if (node.getIsTerminal()) {
            System.out.println(str.toString());
        }

        Node[] children;

        for (int i = 0; i < 26; i++) {
            children = node.getChildren();

            if (children[i] != null) {
                str.append((char)(i + 'a'));
                recursePrintTrie(children[i], str, level + 1);
                str.setLength(level);
            }
        }
    }
}