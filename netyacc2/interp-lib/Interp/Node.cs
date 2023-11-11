namespace interp_lib.Interp
{
    public enum NodeType
    {
        I,
        S,
        LR
    }

    public class Node
    {
        public NodeType NodeType;
        public Token Token;
        public int? I;
        public string? S;
        public Node? Left;
        public Node? Right;

        public Node(Token token)
        {
            this.NodeType = NodeType.I;
            this.Token = token;
        }

        public Node(Token token, int? i)
        {
            this.NodeType = NodeType.I;
            this.Token = token;
            this.I = i;
        }

        public Node(Token token, string? s)
        {
            this.NodeType = NodeType.S;
            this.Token = token;
            this.S = s;
        }
        public Node(Token token, Node? l, Node? r)
        {
            this.NodeType = NodeType.LR;
            this.Token = token;
            this.Left = l;
            this.Right = r;
        }

        public override string ToString()
        {
            switch (this.NodeType)
            {
                case NodeType.S:
                    return $"Node: {NodeType}, {Token}, {S}";
                case NodeType.I:
                    return $"Node: {NodeType}, {Token}, {I}";
                default:
                    return $"Node: {NodeType}, {Token}, {Left}, {Right}";
            }
        }
    }
}