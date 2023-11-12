using System.Runtime.CompilerServices;
[assembly: InternalsVisibleTo("interp_test")]

namespace interp_lib.Interp
{
    internal partial class InterpScanner
    {
        public Dictionary<string, int> StoI = new Dictionary<string, int>();
        public Dictionary<int, string> ItoS = new Dictionary<int, string>();

        public void Reset()
        {
            StoI = new Dictionary<string, int>();
            ItoS = new Dictionary<int, string>();
        }

        public override void yyerror(string format, params object[] args)
        {
            base.yyerror(format, args);

            string msg = string.Format(format, args);
            throw new Exception(string.Format("{0}:{1} {2}", yyline, yycol, msg));
        }

        public int Pool(string s)
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
    }
}
