namespace interp_lib.Interp
{
    internal partial class InterpScanner
    {
        public override void yyerror(string format, params object[] args)
        {
            base.yyerror(format, args);

            string msg = string.Format(format, args);
            throw new Exception(string.Format("{0}:{1} {2}", yyline, yycol, msg));
        }
    }
}
