using System;
using System.Collections.Generic;
using System.IO;
using System.Text;

namespace netyacc1.Calc
{
    internal partial class CalcParser
    {
        public CalcParser() : base(null) { }

        public void Parse(string s)
        {
            byte[] inputBuffer = System.Text.Encoding.Default.GetBytes(s);
            MemoryStream stream = new MemoryStream(inputBuffer);
            this.Scanner = new CalcScanner(stream);
            this.Parse();
        }
    }
}
