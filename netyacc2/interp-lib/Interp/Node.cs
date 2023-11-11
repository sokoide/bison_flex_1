namespace interp_lib.Interp
{
    public class Node
    {
        public Token Token;
        public int? I;
        public string? S;
        public Node? Left;
        public Node? Right;

        public Node(Token t)
        {
            this.Token = t;
        }

        public Node(Token t, int? i)
        {
            this.Token = t;
            this.I = i;
        }

        public Node(Token t, string? s)
        {
            this.Token = t;
            this.S = s;
        }

        public Node(Token t, Node? l, Node? r)
        {
            this.Token = t;
            this.Left = l;
            this.Right = r;
        }

        public override string ToString()
        {
            switch (this.Token)
            {
                case Token.IDENT:
                    return $"Node: {Token}, {S}";
                case Token.NUMBER:
                    return $"Node: {Token}, {I}";
                default:
                    return $"Node: {Token}, {Left}, {Right}";
            }
        }
    }
}