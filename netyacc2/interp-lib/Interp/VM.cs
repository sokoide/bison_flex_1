namespace interp_lib.Interp
{
    public class VM
    {
        const int STACK_SIZE = 1024;

        private int pc;
        private int sp;
        private int[] stack;
        private int[] g;

        public VM()
        {
            Reset();
        }

        public void Execute(List<Instr> code)
        {
            System.Console.WriteLine("* Executing...");
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
                        if (++sp >= STACK_SIZE)
                        {
                            throw new Exception("stack overflow");
                        }
                        stack[sp] = sub;
                        break;
                    case Op.PutI:
                        Console.WriteLine("{0}", g[sub]);
                        break;
                    case Op.PutN:
                        Console.WriteLine("{0}", sub);
                        break;
                    default:
                        throw new NotImplementedException(op.ToString());
                }
                pc++;
            }
            System.Console.WriteLine("* Done");
        }

        public void Reset()
        {
            stack = new int[STACK_SIZE];
            g = new int[26];
            pc = 0;
            sp = -1;
        }
    }
}