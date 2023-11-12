namespace interp_lib.Interp
{
    public partial class InterpParser
    {
        public List<Instr> Code = new List<Instr>();
        public InterpParser() : base(null) { }


        public void Reset()
        {
            Code = new List<Instr>();
        }

        public void Parse(string s)
        {
            byte[] inputBuffer = System.Text.Encoding.Default.GetBytes(s);
            MemoryStream stream = new MemoryStream(inputBuffer);
            this.Scanner = new InterpScanner(stream);
            this.Parse();
        }

        public Node MakeNode(Token t, Node? l, Node? r)
        {
            return new Node(t, l, r);
        }

        public Node MakeNode(Token t, int i)
        {
            return new Node(t, i);
        }

        public Node MakeNode(Token t, string s)
        {
            return new Node(t, s);
        }

        public void GenNode(Node n)
        {
            if (n.Left != null)
            {
                GenNode(n.Left);
            }
            if (n.Right != null)
            {
                GenNode(n.Right);
            }
            switch (n.Token)
            {
                case Token.IDENT:
                    GenCode(Op.PushI, n);
                    break;
                case Token.NUMBER:
                    GenCode(Op.PushN, n);
                    break;
                default:
                    GenCode(Op.Calc, n);
                    break;
            }
        }

        public void GenCode(Op op, Node n)
        {
            Instr instr;
            switch (op)
            {
                case Op.PushI:
                    instr = new Instr(op, IdentId(n.S));
                    break;
                case Op.Pop:
                    instr = new Instr(op, IdentId(n.S));
                    break;
                case Op.PutI:
                    instr = new Instr(op, IdentId(n.S));
                    break;
                case Op.Calc:
                    instr = new Instr(op, (int)n.Token);
                    break;
                default:
                    instr = new Instr(op, n.I);
                    break;
            }
            Code.Add(instr);
        }

        public void GenCode(Op op, int i)
        {
            Instr instr = new Instr(op, i);
            Code.Add(instr);
        }


        public const int FIRST_LABEL = 1001; // for debugging purpose only. 0 is fine, too.
        internal static int labelno = FIRST_LABEL;
        public int NewLabel()
        {
            return labelno++;
        }

        public int IdentId(string? s)
        {
            // TODO: currently, only 'a' to 'z' are supported
            if (s != null)
            {
                return s.ToLower()[0] - 'a';
            }
            else
            {
                return int.MinValue;
            }
        }
    }
}
