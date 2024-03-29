using System.ComponentModel.DataAnnotations;

namespace interp_lib.Interp
{
    public class VM
    {
        const int STACK_SIZE = 1024;

        private int pc;
        private int sp;
        private int[] stack = { };
        private int[] g = { };
        private List<int> labels = new List<int>();

        public VM()
        {
            Reset();
        }

        public int Execute(List<Instr> code, Dictionary<int, string> ItoS, Dictionary<int, Variable> ItoV)
        {
            Op op;
            int sub;

            while (pc < code.Count)
            {
                op = code[pc].Op;
                sub = code[pc].Sub;
                switch (op)
                {
                    case Op.Pop:
                        g[sub] = stack[sp--];
                        break;
                    case Op.PushI:
                        if (++sp >= STACK_SIZE)
                        {
                            throw new Exception("stack overflow");
                        }
                        stack[sp] = g[sub];
                        break;
                    case Op.PushN:
                    case Op.PushS:
                        if (++sp >= STACK_SIZE)
                        {
                            throw new Exception("stack overflow");
                        }
                        stack[sp] = sub;
                        break;
                    case Op.Calc:
                        switch (sub)
                        {
                            case (int)Token.ADD:
                                sp--;
                                stack[sp] = stack[sp] + stack[sp + 1];
                                break;
                            case (int)Token.SUB:
                                sp--;
                                stack[sp] = stack[sp] - stack[sp + 1];
                                break;
                            case (int)Token.MINUS:
                                stack[sp] = -stack[sp];
                                break;
                            case (int)Token.MUL:
                                sp--;
                                stack[sp] = stack[sp] * stack[sp + 1];
                                break;
                            case (int)Token.DIV:
                                sp--;
                                stack[sp] = stack[sp] / stack[sp + 1];
                                break;
                            case (int)Token.EQOP:
                                sp--;
                                stack[sp] = stack[sp] == stack[sp + 1] ? 1 : 0;
                                break;
                            case (int)Token.GTOP:
                                sp--;
                                stack[sp] = stack[sp] > stack[sp + 1] ? 1 : 0;
                                break;
                            case (int)Token.GEOP:
                                sp--;
                                stack[sp] = stack[sp] >= stack[sp + 1] ? 1 : 0;
                                break;
                            case (int)Token.LTOP:
                                sp--;
                                stack[sp] = stack[sp] < stack[sp + 1] ? 1 : 0;
                                break;
                            case (int)Token.LEOP:
                                sp--;
                                stack[sp] = stack[sp] <= stack[sp + 1] ? 1 : 0;
                                break;
                            case (int)Token.NEOP:
                                sp--;
                                stack[sp] = stack[sp] != stack[sp + 1] ? 1 : 0;
                                break;
                            default:
                                throw new NotImplementedException(string.Format("Instr: {0}", code[pc]));
                        }
                        break;
                    case Op.Label:
                        // do nothing
                        break;
                    case Op.Jump:
                        pc = sub;
                        break;
                    case Op.JumpF:
                        if (stack[sp--] == 0)
                        {
                            pc = sub;
                        }
                        break;
                    case Op.PutI:
                        Variable v = ItoV[sub];
                        if (v.Vt == VariableType.INT)
                        {
                            Console.Write("{0}", g[sub]);
                        }
                        else if (v.Vt == VariableType.STRING)
                        {
                            Console.Write("{0}", ItoS[g[sub]]);
                        }
                        else
                        {
                            throw new Exception($"VariableType {v.Vt} not supported");
                        }
                        break;
                    case Op.PutN:
                        Console.Write("{0}", sub);
                        break;
                    case Op.PutS:
                        Console.Write("{0}", ItoS[sub]);
                        break;
                    case Op.ReturnI:
                        return g[sub];
                    case Op.ReturnN:
                        return sub;
                    default:
                        throw new NotImplementedException(op.ToString());
                }
                pc++;
            }
            return 0;
        }

        public List<Instr> ResoleLabels(List<Instr> code)
        {
            List<Instr> resolvedCode = new List<Instr>(code.Count);
            Debug.Print("start");

            // resolve labels
            for (int i = 0; i < code.Count; i++)
            {
                // make a copy
                Instr instr = code[i];
                resolvedCode.Add(new Instr(instr.Op, instr.Sub));

                if (instr.Op == Op.Label)
                {
                    Debug.Print("adding i:{0} {1}", i, instr);
                    // i == pc
                    labels.Add(i);
                }
            }
            // update Jump/JumpF
            for (int i = 0; i < resolvedCode.Count; i++)
            {
                Instr instr = resolvedCode[i];
                if (instr.Op == Op.Jump || instr.Op == Op.JumpF)
                {
                    Debug.Print("resolving {0}", instr);
                    resolvedCode[i].Sub = labels[instr.Sub - InterpParser.FIRST_LABEL];
                }
            }

            Debug.Print("end");
            return resolvedCode;
        }

        public void Dump(List<Instr> code)
        {
            int line = 0;
            foreach (Instr i in code)
            {
                Console.WriteLine("[{0:D4}] {1}", line++, i);
            }
        }

        public void DumpStringTable(Dictionary<int, string> ItoS)
        {
            foreach (var item in ItoS)
            {
                Console.WriteLine("[{0:D4}] {1}", item.Key, item.Value.Replace("\n", "\\n"));
            }
        }

        public void DumpVariableTable(Dictionary<Variable, int> VtoI)
        {
            foreach (var item in VtoI)
            {
                Console.WriteLine("{0}, Index: {1}", item.Key, item.Value);
            }
        }

        public void Reset()
        {
            stack = new int[STACK_SIZE];
            g = new int[26];
            labels = new List<int>();
            pc = 0;
            sp = -1;
        }
    }
}