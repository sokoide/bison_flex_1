using System;
using System.Collections.Generic;
using System.Text;

namespace interp_lib.Interp
{
    internal partial class InterpScanner
    {

        void GetNumber()
        {
            yylval.s = yytext;
            yylval.n = int.Parse(yytext);
        }

        void GetString()
        {
            yylval.s = yytext;
        }

        public override void yyerror(string format, params object[] args)
        {
            base.yyerror(format, args);
            throw new Exception(string.Format(format, args));
            // throw new Exception(string.Format("line: {0}", yylineno);
        }
    }
}
