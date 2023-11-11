namespace interp_lib.Interp
{
    public enum Op
    {
        Pop,
        PushI,
        PushN,
        Calc,
        Label,
        Jump,
        JumpF,
        Get,
        PutI,
        PutN,
        PutS,
    }

    public class Instr
    {
        public Op Op;
        public int Sub;

        public Instr(Op op, int sub)
        {
            this.Op = op;
            this.Sub = sub;
        }

        public override string ToString()
        {
            switch (Op)
            {
                case Op.Calc:
                    Token t = (Token)Enum.ToObject(typeof(Token), Sub);
                    return $"{Op} {t}";
                default:
                    return $"{Op} {Sub}";
            }
        }
    }
}