public class Node {
    private Node[] children;
    private Boolean isTerminal;

    public Node() {
        children = new Node[26];
        isTerminal = false;
    }

    public Node[] getChildren() { return children; }

    public boolean getIsTerminal() { return isTerminal; }
    public void setIsTerminal(boolean isTerminal) { this.isTerminal = isTerminal; }

    public String toString() {
        StringBuilder str = new StringBuilder();

        str.append("[ ");
        for (int i = 0; i < children.length; i++) {
            if (children[i] != null) {
                str.append((char)(i + 'a') + " ");
            }
        }
        str.append("]");
        
        return str.toString();
    }
}