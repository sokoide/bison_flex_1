using System.ComponentModel.DataAnnotations;

namespace interp_lib.Interp
{
    public partial class InterpParser
    {
        public List<Instr> Code = new List<Instr>();
        // string literal -> index
        public Dictionary<string, int> StoI = new Dictionary<string, int>();
        // index -> string literal
        public Dictionary<int, string> ItoS = new Dictionary<int, string>();
        // variable -> ident id
        public Dictionary<Variable, int> VtoI = new Dictionary<Variable, int>();
        // ident id -> variable
        public Dictionary<int, Variable> ItoV = new Dictionary<int, Variable>();

        internal int labelno = FIRST_LABEL;

        public InterpParser() : base(null) { }

        public void Reset()
        {
            labelno = FIRST_LABEL;
            Code = new List<Instr>();
            StoI = new Dictionary<string, int>();
            ItoS = new Dictionary<int, string>();
            VtoI = new Dictionary<Variable, int>();
            ItoV = new Dictionary<int, Variable>();
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
                case Token.NUMBER_LITERAL:
                    GenCode(Op.PushN, n);
                    break;
                case Token.STRING_LITERAL:
                    GenCode(Op.PushS, n);
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
                case Op.PushS:
                    instr = new Instr(op, StringLiteralId(n.S));
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
                    if (n.NodeType == NodeType.S)
                    {
                        instr = new Instr(op, IdentId(n.S));
                    }
                    else
                    {
                        instr = new Instr(op, n.I);
                    }
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
        public int NewLabel()
        {
            return labelno++;
        }

        public int StringLiteralId(string s)
        {
            return PoolStringLiteral(s);
        }

        public int PoolStringLiteral(string s)
        {
            if (StoI.ContainsKey(s))
            {
                return StoI[s];
            }
            int n = StoI.Count + 1;
            StoI[s] = n;
            ItoS[n] = s;
            return n;
        }

        public int IdentId(string s)
        {
            return PoolIdent(s);
        }

        public int PoolIdent(string s, VariableType vt = VariableType.INT)
        {
            Variable v = new Variable(vt, s);
            if (VtoI.ContainsKey(v))
            {
                return VtoI[v];
            }
            int n = VtoI.Count + 1;
            VtoI[v] = n;
            ItoV[n] = v;
            return n;
        }

        public int UpdateIdent(string s, Token token)
        {
            int n;
            VariableType vt;
            switch (token)
            {
                case Token.INT:
                    vt = VariableType.INT;
                    break;
                case Token.STRING:
                    vt = VariableType.STRING;
                    break;
                default:
                    throw new Exception($"Token {token} not implemented yet");
            }
            if (token == Token.STRING) vt = VariableType.STRING;

            Variable v = new Variable(vt, s);
            if (VtoI.ContainsKey(v))
            {
                n = VtoI[v];
                ItoV[n] = v;
                return VtoI[v];
            }
            n = VtoI.Count + 1;
            VtoI[v] = n;
            ItoV[n] = v;
            return n;
        }
    }
}
