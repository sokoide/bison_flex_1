using System;
using System.Collections.Generic;
using System.IO;
using System.Text;

namespace netyacc2.Interp
{
    internal partial class InterpParser
    {
        public InterpParser() : base(null) { }

        public void Parse(string s)
        {
            byte[] inputBuffer = System.Text.Encoding.Default.GetBytes(s);
            MemoryStream stream = new MemoryStream(inputBuffer);
            this.Scanner = new InterpScanner(stream);
            this.Parse();
        }
    }
}
